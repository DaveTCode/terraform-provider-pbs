package provider

import (
	"context"
	"fmt"
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &pbsNodeResource{}
	_ resource.ResourceWithConfigure   = &pbsNodeResource{}
	_ resource.ResourceWithImportState = &pbsNodeResource{}
)

func NewPbsNodeResource() resource.Resource {
	return &pbsNodeResource{}
}

type pbsNodeResource struct {
	client *pbsclient.PbsClient
}

func (r *pbsNodeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_node"
}

func (r *pbsNodeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{ // TODO - How to avoid duplication of this schema with data source?
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeID,
			},
			"comment": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: DescNodeComment,
			},
			"current_aoe": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: DescNodeCurrentAoe,
			},
			"current_eoe": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: DescNodeCurrentEoe,
			},
			"in_multi_node_host": schema.Int32Attribute{
				Optional:            true,
				MarkdownDescription: DescNodeInMultiNodeHost,
			},
			"mom": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: DescNodeMom,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: DescNodeName,
			},
			"no_multinode_jobs": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: DescNodeNoMultinodeJobs,
			},
			"partition": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: DescNodePartition,
			},
			"p_names": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: DescNodePNames,
			},
			"port": schema.Int32Attribute{
				Optional:            true,
				MarkdownDescription: DescNodePort,
				PlanModifiers: []planmodifier.Int32{
					int32planmodifier.RequiresReplace(),
				},
			},
			"poweroff_eligible": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: DescNodePoweroffEligible,
			},
			"power_provisioning": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: DescNodePowerProvisioning,
			},
			"priority": schema.Int32Attribute{
				Optional:            true,
				MarkdownDescription: DescNodePriority,
				Validators: []validator.Int32{
					int32validator.Between(-1024, 1023),
				},
			},
			"provision_enable": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: DescNodeProvisionEnable,
			},
			"queue": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: DescNodeQueue,
				DeprecationMessage:  "This attribute is deprecated",
			},
			"resources_available": schema.MapAttribute{
				Optional:            true,
				MarkdownDescription: DescNodeResourcesAvailable,
				ElementType:         types.StringType,
				Validators: []validator.Map{
					mapvalidator.KeysAre(stringvalidator.NoneOf("host", "vnode")),
				},
			},
			"resv_enable": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: DescNodeResvEnable,
			},
		},
	}
}

func (r *pbsNodeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *pbsNodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model pbsNodeModel
	var pbsNode pbsclient.PbsNode
	diags := req.Plan.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	pbsNode, err := r.client.CreateNode(model.ToPbsNode())
	if err != nil {
		resp.Diagnostics.AddError("Error creating resource", "Could not create node, unexpected error: "+err.Error())
		return
	}

	model = createPbsNodeModel(pbsNode)

	diags = resp.State.Set(ctx, model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *pbsNodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data pbsNodeModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// For import, use ID if name is not set
	nodeName := data.Name.ValueString()
	if nodeName == "" && !data.ID.IsNull() {
		nodeName = data.ID.ValueString()
	}

	pbsNode, err := r.client.GetNode(nodeName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read resources, got error: %s", err))
		return
	}

	// If the resource is not found, remove it from the state
	if pbsNode.Name == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	rModel := createPbsNodeModel(pbsNode)

	resp.Diagnostics.Append(resp.State.Set(ctx, &rModel)...)
}

func (r *pbsNodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data pbsNodeModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	updatedNode, err := r.client.UpdateNode(data.ToPbsNode())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to update the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Error: "+err.Error(),
		)

		return
	}

	// Create the model from the updated node to ensure all fields including ID are properly set
	updatedModel := createPbsNodeModel(updatedNode)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedModel)...)
}

func (r *pbsNodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data pbsNodeModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteNode(data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete node, got error: %s", err))
		return
	}
}

func (r *pbsNodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Use the standard passthrough for ID, which will set both id and trigger a Read
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
