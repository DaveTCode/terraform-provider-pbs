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
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueID,
			},
			"acl_group_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAclGroupEnable,
			},
			"acl_groups": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAclGroups,
			},
			"acl_groups_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAclGroupsNormalized,
			},
			"acl_host_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAclHostEnable,
			},
			"acl_hosts": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAclHosts,
			},
			"acl_hosts_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAclHostsNormalized,
			},
			"acl_user_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAclUserEnable,
			},
			"acl_users": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAclUsers,
			},
			"acl_users_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAclUsersNormalized,
			},
			"alt_router": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueAltRouter,
			},
			"backfill_depth": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueBackfillDepth,
			},
			"checkpoint_min": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueCheckpointMin,
			},
			"default_chunk": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueDefaultChunk,
			},
			"enabled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueEnabled,
			},
			"from_route_only": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueFromRouteOnly,
			},
			"kill_delay": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueKillDelay,
			},
			"max_array_size": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxArraySize,
			},
			"max_group_res": schema.MapAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: DescQueueMaxGroupRes,
			},
			"max_group_res_soft": schema.MapAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: DescQueueMaxGroupResSoft,
			},
			"max_group_run": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxGroupRun,
			},
			"max_group_run_soft": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxGroupRunSoft,
			},
			"max_queuable": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxQueuable,
			},
			"max_queued": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxQueued,
			},
			"max_queued_res": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueMaxQueuedRes,
			},
			"max_run": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxRun,
			},
			"max_run_res": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueMaxRunRes,
			},
			"max_run_res_soft": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueMaxRunResSoft,
			},
			"max_run_soft": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxRunSoft,
			},
			"max_running": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxRunning,
			},
			"max_user_res": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueMaxUserRes,
			},
			"max_user_res_soft": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueMaxUserResSoft,
			},
			"max_user_run": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxUserRun,
			},
			"max_user_run_soft": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueMaxUserRunSoft,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: DescQueueName,
			},
			"node_group_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueNodeGroupKey,
			},
			"partition": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueuePartition,
			},
			"priority": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueuePriority,
			},
			"queued_jobs_threshold": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueQueuedJobsThreshold,
			},
			"queued_jobs_threshold_res": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueQueuedJobsThresholdRes,
			},
			"queue_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueQtype,
			},
			"resources_assigned": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueResourcesAssigned,
			},
			"resources_available": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueResourcesAvailable,
			},
			"resources_default": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueResourcesDefault,
			},
			"resources_max": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueResourcesMax,
			},
			"resources_min": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescQueueResourcesMin,
			},
			"route_destinations": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueRouteDestinations,
			},
			"route_held_jobs": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueRouteHeldJobs,
			},
			"route_lifetime": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueRouteLifetime,
			},
			"route_retry_time": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescQueueRouteRetryTime,
			},
			"route_waiting_jobs": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueRouteWaitingJobs,
			},
			"started": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescQueueStarted,
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

	queueModel := createQueueModel(resultData)

	diag := resp.State.Set(ctx, &queueModel)
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
