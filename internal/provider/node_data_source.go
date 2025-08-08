package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-pbs/internal/pbsclient"
)

func NewPbsNodeDataSource() datasource.DataSource {
	return &pbsNodeDataSource{}
}

type pbsNodeDataSource struct {
	client *pbsclient.PbsClient
}

func (d *pbsNodeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_node"
}

func (d *pbsNodeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for this node. This is the same as the name.",
			},
			"comment": schema.StringAttribute{
				Computed:    true,
				Description: "Information about this vnode.  This attribute may be set by the manager to any string to inform users of any information relating to the node. If this attribute is not explicitly set, the PBS server will use the attribute to pass information about the node status, specifically why the node is down. If the attribute is explicitly set by the manager, it will not be modified by the server.",
			},
			"current_aoe": schema.StringAttribute{
				Computed:    true,
				Description: "The AOE currently instantiated on this vnode.  Case-sensitive.  Cannot be set on server's host.",
			},
			"current_eoe": schema.StringAttribute{
				Computed:    true,
				Description: "Current value of eoe on this vnode. We do not recommend setting this attribute manually.",
			},
			"in_multi_node_host": schema.Int32Attribute{
				Computed:    true,
				Description: "Specifies whether a vnode is part of a multi-vnoded host.  Used internally.  Do not set.",
			},
			"mom": schema.StringAttribute{
				Computed:    true,
				Description: "host.  By default the server queries the canonicalized name of the MoM host, unless you set this attribute when you create the vnode.  Can be explicitly set by Manager only via qmgr, and only at vnode creation.  The server can set this to the FQDN of the host on which MoM runs, if the vnode name is the same as the hostname.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of this vnode. Must be resolvable to an IP address.  Must be unique within the server.",
			},
			"no_multinode_jobs": schema.BoolAttribute{
				Computed:    true,
				Description: " Controls whether jobs which request more than one chunk are allowed to execute on this vnode.  Used for cycle harvesting.",
			},
			"partition": schema.StringAttribute{
				Computed:    true,
				Description: "Name of partition to which this vnode is assigned.  A vnode can be assigned to at most one partition.",
			},
			"p_names": schema.StringAttribute{
				Computed:    true,
				Description: "The list of resources being used for placement sets.  Not used for scheduling; advisory only.",
			},
			"port": schema.Int32Attribute{
				Computed:    true,
				Description: "Port number on which MoM daemon listens. Can be explicitly set only via qmgr, and only at vnode creation.",
			},
			"poweroff_eligible": schema.BoolAttribute{
				Computed:    true,
				Description: "Enables powering this vnode up and down by PBS.",
			},
			"power_provisioning": schema.BoolAttribute{
				Computed:    true,
				Description: "Specifies whether this node is eligible to have its power managed by PBS, including whether it can use power profiles.",
			},
			"priority": schema.Int32Attribute{
				Computed:    true,
				Description: "The priority of this vnode compared with other vnodes.",
				Validators: []validator.Int32{
					int32validator.Between(-1024, 1023),
				},
			},
			"provision_enable": schema.BoolAttribute{
				Computed:    true,
				Description: "Controls whether this vnode can be provisioned.  Cannot be set on server's host.",
			},
			"queue": schema.StringAttribute{
				Computed:           true,
				Description:        "Deprecated.  The queue with which this vnode is associated.  Each vnode can be associated with at most 1 queue.  Queues can be associated with multiple vnodes.  Any jobs in a queue that has associated vnodes can run only on those vnodes.  If a vnode has an associated queue, only jobs in that queue can run on that vnode.",
				DeprecationMessage: "This attribute is deprecated",
			},
			"resources_available": schema.MapAttribute{
				Computed:    true,
				Description: "The list of resources and the amounts available on this vnode. If not explicitly set, the amount shown is that reported by the pbs_mom running on this vnode. If a resource value is explicitly set, that value is retained across restarts.",
				ElementType: types.StringType,
			},
			"resv_enable": schema.BoolAttribute{
				Computed:    true,
				Description: "Controls whether the vnode can be used for advance and standing reservations.  Reservations are incompatible with cycle harvesting. ",
			},
		},
	}
}

func (d *pbsNodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	sourceData := pbsNodeModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &sourceData)...)

	resultData, err := d.client.GetNode(sourceData.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to connect to PBS server and get hook information", err.Error())
		return
	}

	model := createPbsNodeModel(resultData)

	diag := resp.State.Set(ctx, &model)
	resp.Diagnostics.Append(diag...)
}

func (d *pbsNodeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*pbsclient.PbsClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *PbsClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
