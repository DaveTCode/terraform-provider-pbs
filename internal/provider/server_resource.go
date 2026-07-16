package provider

import (
	"context"
	"fmt"
	"terraform-provider-pbs/internal/pbsclient"
	validators "terraform-provider-pbs/internal/provider/validators"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &serverResource{}
	_ resource.ResourceWithConfigure   = &serverResource{}
	_ resource.ResourceWithImportState = &serverResource{}
)

func NewServerResource() resource.Resource {
	return &serverResource{}
}

type serverResource struct {
	client *pbsclient.PbsClient
}

func (r *serverResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (r *serverResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerID,
			},
			"acl_host_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclHostEnable,
			},
			"acl_host_moms_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclHostsMomsEnable,
			},
			"acl_hosts": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclHosts,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"acl_hosts_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclHostsNormalized,
			},
			"acl_resv_group_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclResvGroupEnable,
			},
			"acl_resv_groups": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclResvGroups,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"acl_resv_groups_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvGroupsNormalized,
			},
			"acl_resv_host_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclResvHostEnable,
			},
			"acl_resv_hosts": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclResvHosts,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"acl_resv_hosts_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvHostsNormalized,
			},
			"acl_resv_user_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclResvUserEnable,
			},
			"acl_resv_users": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclResvUsers,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"acl_resv_users_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclResvUsersNormalized,
			},
			"acl_roots": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclRoots,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"acl_roots_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclRootsNormalized,
			},
			"acl_user_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclUserEnable,
			},
			"acl_users": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerAclUsers,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"acl_users_normalized": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescServerAclUsersNormalized,
			},
			"backfill_depth": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerBackfillDepth,
				Validators: []validator.Int32{
					int32validator.AtLeast(0),
				},
			},
			"comment": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerComment,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"default_chunk": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:         types.StringType,
				MarkdownDescription: DescServerDefaultChunk,
			},
			"default_qdel_arguments": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerDefaultQdelArguments,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"default_qsub_arguments": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerDefaultQsubArguments,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"default_queue": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerDefaultQueue,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"eligible_time_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerEligibleTimeEnable,
			},
			"elim_on_subjobs": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerElimOnSubjobs,
			},
			"flatuid": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerFlatuid,
			},
			"job_history_duration": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerJobHistoryDuration,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"job_history_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerJobHistoryEnable,
			},
			"job_requeue_timeout": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerJobRequeueTimeout,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"job_sort_formula": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerJobSortFormula,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"jobscript_max_size": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerJobscriptMaxSize,
				Validators: []validator.String{
					validators.PbsSize(),
				},
			},
			"log_events": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerLogEvents,
			},
			"mailer": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMailer,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"mail_from": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMailFrom,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"managers": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerManagers,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"max_array_size": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxArraySize,
			},
			"max_concurrent_provision": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxConcurrentProvision,
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"max_group_res": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxGroupRes,
			},
			"max_group_res_soft": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxGroupResSoft,
			},
			"max_group_run": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxGroupRun,
			},
			"max_group_run_soft": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxGroupRunSoft,
			},
			"max_job_sequence_id": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxJobSequenceId,
				Validators: []validator.Int64{
					int64validator.Between(9999999, 999999999999),
				},
			},
			"max_queued": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxQueued,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"max_queued_res": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxQueuedRes,
			},
			"max_run": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxRun,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"max_run_res": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxRunRes,
			},
			"max_run_res_soft": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxRunResSoft,
			},
			"max_run_soft": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxRunSoft,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"max_running": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxRunning,
			},
			"max_user_res": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxUserRes,
			},
			"max_user_res_soft": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:         types.StringType,
				MarkdownDescription: DescServerMaxUserResSoft,
			},
			"max_user_run": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxUserRun,
			},
			"max_user_run_soft": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerMaxUserRunSoft,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: DescServerName,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"node_fail_requeue": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerNodeFailRequeue,
			},
			"node_group_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerNodeGroupEnable,
			},
			"node_group_key": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerNodeGroupKey,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"operators": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerOperators,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"pbs_license_info": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerPbsLicenseInfo,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"pbs_license_linger_time": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerPbsLicenseLingerTime,
			},
			"pbs_license_max": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerPbsLicenseMax,
			},
			"pbs_license_min": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerPbsLicenseMin,
			},
			"power_provisioning": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerPowerProvisioning,
			},
			"python_gc_min_interval": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerPythonGcCollectMinInterval,
			},
			"python_restart_max_hooks": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerPythonRestartMaxHooks,
			},
			"python_restart_max_objects": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerPythonRestartMaxObjects,
			},
			"python_restart_min_interval": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerPythonRestartMinInterval,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"query_other_jobs": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerQueryOtherJobs,
			},
			"queued_jobs_threshold": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerQueuedJobsThreshold,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"queued_jobs_threshold_res": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerQueuedJobsThresholdRes,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"reserve_retry_init": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerReserveRetryInit,
				DeprecationMessage:  "Deprecated",
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"reserve_retry_time": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerReserveRetryTime,
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"resources_available": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerResourcesAvailable,
				ElementType:         types.StringType,
			},
			"resources_default": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerResourcesDefault,
				ElementType:         types.StringType,
			},
			"resources_max": schema.MapAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerResourcesMax,
				ElementType:         types.StringType,
			},
			"restrict_res_to_release_on_suspend": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerRestrictResToReleaseOnSuspend,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"resv_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerResvEnable,
			},
			"resv_post_processing_time": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerResvPostProcessingTime,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"rpp_highwater": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerRppHighwater,
			},
			"rpp_max_pkt_check": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerRppMaxPktCheck,
			},
			"rpp_retry": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerRppRetry,
			},
			"scheduler_iteration": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int32{int32planmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerSchedulerIteration,
			},
			"webapi_auth_issuers": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerWebapiAuthIssuers,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"webapi_enable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerWebapiEnable,
			},
			"webapi_oidc_clientid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerWebapiOidcClientid,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
			"webapi_oidc_provider_url": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				MarkdownDescription: DescServerWebapiOidcProviderUrl,
				Validators: []validator.String{
					validators.PbsString(),
				},
			},
		},
	}
}

func (r *serverResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *serverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// PBS server resources cannot be created - they must be imported
	resp.Diagnostics.AddError(
		"Server Resource Cannot Be Created",
		"PBS server resources cannot be created through Terraform as there is exactly one server per PBS cluster that already exists. "+
			"Please use 'terraform import' to import the existing server configuration. "+
			"Example: terraform import pbs_server.example server",
	)
}

func (r *serverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var currentState serverModel

	resp.Diagnostics.Append(req.State.Get(ctx, &currentState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// For import, use ID if name is not set
	serverName := currentState.Name.ValueString()
	if serverName == "" && !currentState.ID.IsNull() {
		serverName = currentState.ID.ValueString()
	}

	q, err := r.client.GetPbsServer(serverName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read servers, got error: %s", err))
		return
	}

	// If the server doesn't exist then remove from state
	if q.Name == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	updatedState := createServerModel(q)

	// Preserve user-provided ACL formats when semantically equivalent.
	preserveUserServerAclFormatFromState(&currentState, &updatedState)

	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedState)...)
}

func (r *serverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var planData, stateData serverModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	server := planData.ToPbsServer(ctx)
	_, err := r.client.UpdatePbsServer(server)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to update the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Error: "+err.Error(),
		)

		return
	}

	// Read the updated server to get the actual state including computed fields
	serverName := planData.Name.ValueString()
	if serverName == "" && !planData.ID.IsNull() {
		serverName = planData.ID.ValueString()
	}

	updatedServer, err := r.client.GetPbsServer(serverName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read updated server, got error: %s", err))
		return
	}

	// Create model from the actual server state
	updatedData := createServerModel(updatedServer)

	// Preserve user-provided ACL formats from plan where possible
	preserveUserServerAclFormat(&planData, &updatedData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedData)...)
}

func (r *serverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// PBS server resources cannot be deleted - just remove from Terraform state
	// The actual PBS server continues to exist in the cluster
	resp.Diagnostics.AddWarning(
		"Server Resource Not Deleted",
		"The PBS server resource has been removed from Terraform state but the actual PBS server configuration remains unchanged. "+
			"PBS server resources cannot be deleted as there is exactly one server per PBS cluster.",
	)
	// State is automatically removed by the framework
}

func (r *serverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Use the standard passthrough for ID, which will set both id and trigger a Read
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
