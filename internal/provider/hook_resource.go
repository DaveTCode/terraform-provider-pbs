package provider

import (
	"context"
	"fmt"
	"strings"
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	_ resource.Resource                = &pbsHookResource{}
	_ resource.ResourceWithConfigure   = &pbsHookResource{}
	_ resource.ResourceWithImportState = &pbsHookResource{}
)

func NewPbsHookResource() resource.Resource {
	return &pbsHookResource{}
}

type pbsHookResource struct {
	client *pbsclient.PbsClient
}

func (r *pbsHookResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hook"
}

func (r *pbsHookResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{ // TODO - How to avoid duplication of this schema with data source?
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for this hook. This is the same as the name.",
			},
			"alarm": schema.Int32Attribute{
				Description: "Specifies the number of seconds to allow a hook to run before the hook times out.",
				Optional:    true,
				Validators: []validator.Int32{
					int32validator.AtLeast(0),
				},
			},
			"debug": schema.BoolAttribute{
				Description: "debugging files under PBS_HOME/server_priv/hooks/tmp or PBS_HOME/mom_priv/hooks/tmp.  Files are named hook_<hook event>_<hook name>_<unique ID>.in, .data, and .out",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: " Determines whether or not a hook  is  run when its triggering event occurs.",
				Optional:    true,
			},
			"event": schema.StringAttribute{
				Description: "List of events that trigger the hook. The provision event cannot be combined with any other events.",
				Optional:    true,
			},
			"fail_action": schema.StringAttribute{
				MarkdownDescription: "Specifies the action to be taken when hook fails due to alarm call or unhandled exception, or to an internal error such as not enough disk space or memory.  Can also specify a subsequent action to be taken when hook runs successfully.  Value can be either `none` or one or more of `offline_vnodes`, `clear_vnodes_upon_recovery`, and `scheduler_restart_cycle`. If this attribute is set to multiple values, scheduler restart happens last.",
				Optional:            true,
			},
			"freq": schema.Int32Attribute{
				MarkdownDescription: " Number of seconds between `periodic` or `exechost_periodic` triggers.",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				Description: "The unique name of the hook on the server",
				Required:    true,
			},
			"order": schema.Int32Attribute{
				Description: "Indicates relative order of hook execution, for hooks of the same type sharing a trigger.  Hooks with lower order values execute before those with higher values. Does not apply to periodic or exechost_periodic hooks. ",
				Optional:    true,
				Validators: []validator.Int32{
					int32validator.Between(-1000, 2000), // Range: built-in hooks: [-1000, 2000] site hooks: [1,1000] but we are not enforcing this
				},
			},
			"type": schema.StringAttribute{
				Description: "The type of the hook. Cannot be set for a built-in hook.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOf("site", "pbs"),
				},
			},
			"user": schema.StringAttribute{
				Description: "Specifies who executes the hook.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOf("pbsadmin", "pbsuser"),
				},
			},
		},
	}
}

func (r *pbsHookResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*pbsclient.PbsClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *pbsclient.PbsClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *pbsHookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model pbsHookModel
	var pbsHook pbsclient.PbsHook
	diags := req.Plan.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	pbsHook, err := r.client.CreateHook(model.ToPbsHook())
	if err != nil {
		resp.Diagnostics.AddError("Error creating resource", "Could not create hook, unexpected error: "+err.Error())
		return
	}

	model = createPbsHookModel(pbsHook)

	diags = resp.State.Set(ctx, model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *pbsHookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state pbsHookModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// For import, use ID if name is not set
	hookName := state.Name.ValueString()
	if hookName == "" && !state.ID.IsNull() {
		hookName = state.ID.ValueString()
	}

	pbsHook, err := r.client.GetHook(hookName)
	if err != nil {
		// Check if hook doesn't exist and remove from state
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "does not exist") || strings.Contains(err.Error(), "Unknown hook") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read hook, got error: %s", err))
		return
	}

	// If the resource is not found, remove it from the state
	if pbsHook.Name == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	// Update state with current values, preserving plan-only values
	updatedState := createPbsHookModel(pbsHook)
	// Preserve the name from the original state to avoid unnecessary changes,
	// but only if it's not empty (during import, state.Name will be empty)
	if !state.Name.IsNull() && state.Name.ValueString() != "" {
		updatedState.Name = state.Name
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedState)...)
}

func (r *pbsHookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data pbsHookModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	updatedHook, err := r.client.UpdateHook(data.ToPbsHook())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to update the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Error: "+err.Error(),
		)

		return
	}

	// Create the model from the updated hook to ensure all fields including ID are properly set
	updatedModel := createPbsHookModel(updatedHook)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedModel)...)
}

func (r *pbsHookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data pbsHookModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteHook(data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete hook, got error: %s", err))
		return
	}
}

func (r *pbsHookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Use the standard passthrough for ID, which will set both id and trigger a Read
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
