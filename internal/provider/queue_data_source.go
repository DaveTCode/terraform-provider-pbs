package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-pbs/internal/pbsclient"
)

func NewQueueDataSource() datasource.DataSource {
	return &queueDataSource{}
}

type queueDataSource struct {
	client *pbsclient.PbsClient
}

func (d *queueDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_queue"
}

func (d *queueDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"acl_group_enable": schema.BoolAttribute{
				Computed: true,
			},
			"acl_groups": schema.StringAttribute{
				Computed: true,
			},
			"acl_host_enable": schema.BoolAttribute{
				Computed: true,
			},
			"acl_hosts": schema.StringAttribute{
				Computed: true,
			},
			"acl_user_enable": schema.BoolAttribute{
				Computed: true,
			},
			"acl_users": schema.StringAttribute{
				Computed: true,
			},
			"alt_router": schema.StringAttribute{
				Computed: true,
			},
			"backfill_depth": schema.Int32Attribute{
				Computed: true,
			},
			"checkpoint_min": schema.Int32Attribute{
				Computed: true,
			},
			"default_chunk": schema.StringAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"from_route_only": schema.BoolAttribute{
				Computed: true,
			},
			"kill_delay": schema.Int32Attribute{
				Computed: true,
			},
			"max_array_size": schema.Int32Attribute{
				Computed: true,
			},
			"max_group_res": schema.Int32Attribute{
				Computed: true,
			},
			"max_group_res_soft": schema.Int32Attribute{
				Computed: true,
			},
			"max_group_run": schema.Int32Attribute{
				Computed: true,
			},
			"max_group_run_soft": schema.Int32Attribute{
				Computed: true,
			},
			"max_queuable": schema.Int32Attribute{
				Computed: true,
			},
			"max_queued": schema.StringAttribute{
				Computed: true,
			},
			"max_queued_res": schema.StringAttribute{
				Computed: true,
			},
			"max_run": schema.StringAttribute{
				Computed: true,
			},
			"max_run_res": schema.StringAttribute{
				Computed: true,
			},
			"max_run_res_soft": schema.StringAttribute{
				Computed: true,
			},
			"max_run_soft": schema.StringAttribute{
				Computed: true,
			},
			"max_running": schema.Int32Attribute{
				Computed: true,
			},
			"max_user_res": schema.StringAttribute{
				Computed: true,
			},
			"max_user_res_soft": schema.StringAttribute{
				Computed: true,
			},
			"max_user_run": schema.Int32Attribute{
				Computed: true,
			},
			"max_user_run_soft": schema.Int32Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"node_group_key": schema.StringAttribute{
				Computed: true,
			},
			"partition": schema.StringAttribute{
				Computed: true,
			},
			"priority": schema.Int32Attribute{
				Computed: true,
			},
			"queued_jobs_threshold": schema.StringAttribute{
				Computed: true,
			},
			"queued_jobs_threshold_res": schema.StringAttribute{
				Computed: true,
			},
			"queue_type": schema.StringAttribute{
				Computed: true,
			},
			"resources_assigned": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"resources_available": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"resources_default": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"resources_max": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"resources_min": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"route_destinations": schema.StringAttribute{
				Computed: true,
			},
			"route_held_jobs": schema.BoolAttribute{
				Computed: true,
			},
			"route_lifetime": schema.Int32Attribute{
				Computed: true,
			},
			"route_retry_time": schema.Int32Attribute{
				Computed: true,
			},
			"route_waiting_jobs": schema.BoolAttribute{
				Computed: true,
			},
			"started": schema.BoolAttribute{
				Computed: true,
			},
		},
	}
}

func (d *queueDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	sourceData := queueModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &sourceData)...)

	resultData, err := d.client.GetQueue(sourceData.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to connect to PBS server and get hook information", err.Error())
		return
	}

	queueModel, diag := createQueueModel(resultData)
	resp.Diagnostics.Append(diag...)

	diag = resp.State.Set(ctx, &queueModel)
	resp.Diagnostics.Append(diag...)
}

func (d *queueDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
