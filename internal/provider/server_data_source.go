package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-pbs/internal/pbsclient"
)

func NewServerDataSource() datasource.DataSource {
	return &serverDataSource{}
}

type serverDataSource struct {
	client *pbsclient.PbsClient
}

func (d *serverDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (d *serverDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerID,
			},
			"acl_host_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclHostEnable,
			},
			"acl_host_moms_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclHostsMomsEnable,
			},
			"acl_hosts": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclHosts,
			},
			"acl_hosts_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclHostsNormalized,
			},
			"acl_resv_group_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvGroupEnable,
			},
			"acl_resv_groups": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvGroups,
			},
			"acl_resv_groups_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvGroupsNormalized,
			},
			"acl_resv_host_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvHostEnable,
			},
			"acl_resv_hosts": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvHosts,
			},
			"acl_resv_hosts_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvHostsNormalized,
			},
			"acl_resv_user_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvUserEnable,
			},
			"acl_resv_users": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvUsers,
			},
			"acl_resv_users_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvUsersNormalized,
			},
			"acl_roots": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclRoots,
			},
			"acl_roots_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclRootsNormalized,
			},
			"acl_user_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclUserEnable,
			},
			"acl_users": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclUsers,
			},
			"acl_users_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclUsersNormalized,
			},
			"backfill_depth": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerBackfillDepth,
				Validators: []validator.Int32{
					int32validator.AtLeast(0),
				},
			},
			"comment": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerComment,
			},
			"default_chunk": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerDefaultChunk,
			},
			"default_qdel_arguments": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerDefaultQdelArguments,
			},
			"default_qsub_arguments": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerDefaultQsubArguments,
			},
			"default_queue": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerDefaultQueue,
			},
			"eligible_time_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerEligibleTimeEnable,
			},
			"elim_on_subjobs": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerElimOnSubjobs,
			},
			"flatuid": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerFlatuid,
			},
			"job_history_duration": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerJobHistoryDuration,
			},
			"job_history_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerJobHistoryEnable,
			},
			"job_requeue_timeout": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerJobRequeueTimeout,
			},
			"job_sort_formula": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerJobSortFormula,
			},
			"jobscript_max_size": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerJobscriptMaxSize,
			},
			"log_events": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerLogEvents,
			},
			"mailer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerMailer,
			},
			"mail_from": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerMailFrom,
			},
			"managers": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerManagers,
			},
			"max_array_size": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerMaxArraySize,
			},
			"max_concurrent_provision": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerMaxConcurrentProvision,
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"max_group_res": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxGroupRes,
			},
			"max_group_res_soft": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxGroupResSoft,
			},
			"max_group_run": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerMaxGroupRun,
			},
			"max_group_run_soft": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerMaxGroupRunSoft,
			},
			"max_job_sequence_id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: DescServerMaxJobSequenceId,
				Validators: []validator.Int64{
					int64validator.Between(9999999, 999999999999),
				},
			},
			"max_queued": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxQueued,
			},
			"max_queued_res": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxQueuedRes,
			},
			"max_run": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxRun,
			},
			"max_run_res": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxRunRes,
			},
			"max_run_res_soft": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxRunResSoft,
			},
			"max_run_soft": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxRunSoft,
			},
			"max_running": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerMaxRunning,
			},
			"max_user_res": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxUserRes,
			},
			"max_user_res_soft": schema.MapAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxUserResSoft,
			},
			"max_user_run": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerMaxUserRun,
			},
			"max_user_run_soft": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerMaxUserRunSoft,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: DescServerName,
			},
			"node_fail_requeue": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerNodeFailRequeue,
			},
			"node_group_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerNodeGroupEnable,
			},
			"node_group_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerNodeGroupKey,
			},
			"operators": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerOperators,
			},
			"pbs_license_info": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerPbsLicenseInfo,
			},
			"pbs_license_linger_time": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerPbsLicenseLingerTime,
			},
			"pbs_license_max": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerPbsLicenseMax,
			},
			"pbs_license_min": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerPbsLicenseMin,
			},
			"power_provisioning": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerPowerProvisioning,
			},
			"python_gc_min_interval": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerPythonGcCollectMinInterval,
			},
			"python_restart_max_hooks": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerPythonRestartMaxHooks,
			},
			"python_restart_max_objects": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerPythonRestartMaxObjects,
			},
			"python_restart_min_interval": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerPythonRestartMinInterval,
			},
			"query_other_jobs": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerQueryOtherJobs,
			},
			"queued_jobs_threshold": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerQueuedJobsThreshold,
			},
			"queued_jobs_threshold_res": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerQueuedJobsThresholdRes,
			},
			"reserve_retry_init": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerReserveRetryInit,
				DeprecationMessage:  "Deprecated",
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"reserve_retry_time": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerReserveRetryTime,
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"resources_available": schema.MapAttribute{
				Computed:            true,
				MarkdownDescription: DescServerResourcesAvailable,
				ElementType:         types.StringType,
			},
			"resources_default": schema.MapAttribute{
				Computed:            true,
				MarkdownDescription: DescServerResourcesDefault,
				ElementType:         types.StringType,
			},
			"resources_max": schema.MapAttribute{
				Computed:            true,
				MarkdownDescription: DescServerResourcesMax,
				ElementType:         types.StringType,
			},
			"restrict_res_to_release_on_suspend": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerRestrictResToReleaseOnSuspend,
			},
			"resv_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerResvEnable,
			},
			"resv_post_processing_time": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerResvPostProcessingTime,
			},
			"rpp_highwater": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerRppHighwater,
			},
			"rpp_max_pkt_check": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerRppMaxPktCheck,
			},
			"rpp_retry": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerRppRetry,
			},
			"scheduler_iteration": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescServerSchedulerIteration,
			},
			"webapi_auth_issuers": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerWebapiAuthIssuers,
			},
			"webapi_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescServerWebapiEnable,
			},
			"webapi_oidc_clientid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerWebapiOidcClientid,
			},
			"webapi_oidc_provider_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerWebapiOidcProviderUrl,
			},
		},
	}
}

func (d *serverDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	sourceData := serverModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &sourceData)...)

	resultData, err := d.client.GetPbsServer(sourceData.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to connect to PBS server and get information", err.Error())
		return
	}

	serverModel := createServerModel(resultData)

	diag := resp.State.Set(ctx, &serverModel)
	resp.Diagnostics.Append(diag...)
}

func (d *serverDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
