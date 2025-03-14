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
			"comment": schema.StringAttribute{
				Optional:    true,
				Description: "Information about this vnode.  This attribute may be set by the manager to any string to inform users of any information relating to the node. If this attribute is not explicitly set, the PBS server will use the attribute to pass information about the node status, specifically why the node is down. If the attribute is explicitly set by the manager, it will not be modified by the server.",
			},
			"current_aoe": schema.StringAttribute{
				Optional:    true,
				Description: "The AOE currently instantiated on this vnode.  Case-sensitive.  Cannot be set on server's host.",
			},
			"current_eoe": schema.StringAttribute{
				Optional:    true,
				Description: "Current value of eoe on this vnode. We do not recommend setting this attribute manually.",
			},
			"in_multi_node_host": schema.Int32Attribute{
				Optional:    true,
				Description: "Specifies whether a vnode is part of a multi-vnoded host.  Used internally.  Do not set.",
			},
			"mom": schema.StringAttribute{
				Optional:    true,
				Description: "Hostname where server queries for MoM host.  By default the server queries the canonicalized name of the MoM host, unless you set this attribute when you create the vnode.  Can be explicitly set by Manager only via qmgr, and only at vnode creation.  The server can set this to the FQDN of the host on which MoM runs, if the vnode name is the same as the hostname.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of this vnode. Must be resolvable to an IP address.  Must be unique within the server.",
			},
			"no_multinode_jobs": schema.BoolAttribute{
				Optional:    true,
				Description: " Controls whether jobs which request more than one chunk are allowed to execute on this vnode.  Used for cycle harvesting.",
			},
			"partition": schema.StringAttribute{
				Optional:    true,
				Description: "Name of partition to which this vnode is assigned.  A vnode can be assigned to at most one partition.",
			},
			"p_names": schema.StringAttribute{
				Optional:    true,
				Description: "The list of resources being used for placement sets.  Not used for scheduling; advisory only.",
			},
			"port": schema.Int32Attribute{
				Optional:    true,
				Description: "Port number on which MoM daemon listens. Can be explicitly set only via qmgr, and only at vnode creation.",
				PlanModifiers: []planmodifier.Int32{
					int32planmodifier.RequiresReplace(),
				},
			},
			"power_off_eligible": schema.BoolAttribute{
				Optional:    true,
				Description: "Enables powering this vnode up and down by PBS.",
			},
			"power_provisioning": schema.BoolAttribute{
				Optional:    true,
				Description: "Specifies whether this node is eligible to have its power managed by PBS, including whether it can use power profiles.",
			},
			"priority": schema.Int32Attribute{
				Optional:    true,
				Description: "The priority of this vnode compared with other vnodes.",
				Validators: []validator.Int32{
					int32validator.Between(-1024, 1023),
				},
			},
			"provision_enable": schema.BoolAttribute{
				Optional:    true,
				Description: "Controls whether this vnode can be provisioned.  Cannot be set on server's host.",
			},
			"queue": schema.StringAttribute{
				Optional:           true,
				Description:        "Deprecated.  The queue with which this vnode is associated.  Each vnode can be associated with at most 1 queue.  Queues can be associated with multiple vnodes.  Any jobs in a queue that has associated vnodes can run only on those vnodes.  If a vnode has an associated queue, only jobs in that queue can run on that vnode.",
				DeprecationMessage: "This attribute is deprecated",
			},
			"resources_available": schema.MapAttribute{
				Optional:    true,
				Description: "The list of resources and the amounts available on this vnode. If not explicitly set, the amount shown is that reported by the pbs_mom running on this vnode. If a resource value is explicitly set, that value is retained across restarts.",
				ElementType: types.StringType,
				Validators: []validator.Map{
					mapvalidator.KeysAre(stringvalidator.NoneOf("host", "vnode")),
				},
			},
			"resv_enable": schema.BoolAttribute{
				Required:    true,
				Description: "Controls whether the vnode can be used for advance and standing reservations.  Reservations are incompatible with cycle harvesting. ",
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

	pbsNode, err := r.client.GetNode(data.Name.ValueString())
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

	_, err := r.client.UpdateNode(data.ToPbsNode())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to update the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Error: "+err.Error(),
		)

		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
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
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}
