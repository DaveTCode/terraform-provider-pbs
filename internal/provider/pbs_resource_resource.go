package provider

import (
	"context"
	"fmt"
	"regexp"
	"terraform-provider-pbs/internal/pbsclient"
	validators "terraform-provider-pbs/internal/provider/validators"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	_ resource.Resource                = &pbsResourceResource{}
	_ resource.ResourceWithConfigure   = &pbsResourceResource{}
	_ resource.ResourceWithImportState = &pbsResourceResource{}
)

func NewPbsResourceResource() resource.Resource {
	return &pbsResourceResource{}
}

type pbsResourceResource struct {
	client *pbsclient.PbsClient
}

func (r *pbsResourceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (r *pbsResourceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescPbsResourceID,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: DescPbsResourceName,
				Required:            true,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: DescPbsResourceType,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile("^(boolean|string|long|size|float|string_array)$"),
						"resource type must be one of boolean|string|long|size|float|string_array",
					),
				},
			},
			"flag": schema.StringAttribute{
				MarkdownDescription: DescPbsResourceFlag,
				Optional:            true,
				// TODO - Validators for the flags when I understand them better
			},
		},
	}
}

func (r *pbsResourceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *pbsResourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var resourceModel pbsResourceModel
	var pbsResource pbsclient.PbsResource
	diags := req.Plan.Get(ctx, &resourceModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	pbsResource, err := r.client.CreateResource(resourceModel.ToPbsResource())
	if err != nil {
		resp.Diagnostics.AddError("Error creating resource", "Could not create resource, unexpected error: "+err.Error())
		return
	}

	resourceModel = createPbsResoureModel(pbsResource)

	diags = resp.State.Set(ctx, resourceModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *pbsResourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data pbsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// For import, use ID if name is not set
	resourceName := data.Name.ValueString()
	if resourceName == "" && !data.ID.IsNull() {
		resourceName = data.ID.ValueString()
	}

	pbsResource, err := r.client.GetResource(resourceName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read resources, got error: %s", err))
		return
	}

	// If the resource is not found, remove it from the state
	if pbsResource.Name == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	rModel := createPbsResoureModel(pbsResource)

	resp.Diagnostics.Append(resp.State.Set(ctx, &rModel)...)
}

func (r *pbsResourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data pbsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	updatedResource, err := r.client.UpdateResource(data.ToPbsResource())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to update the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Error: "+err.Error(),
		)

		return
	}

	// Create the model from the updated resource to ensure all fields including ID are properly set
	updatedModel := createPbsResoureModel(updatedResource)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedModel)...)
}

func (r *pbsResourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data pbsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteResource(data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete resource, got error: %s", err))
		return
	}
}

func (r *pbsResourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Use the standard passthrough for ID, which will set both id and trigger a Read
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
