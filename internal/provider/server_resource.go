package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"terraform-provider-pbs/internal/pbsclient"
	validators "terraform-provider-pbs/internal/provider/validators"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ resource.Resource                = &serverResource{}
	_ resource.ResourceWithConfigure   = &serverResource{}
	_ resource.ResourceWithImportState = &serverResource{}
	_ resource.ResourceWithModifyPlan  = &serverResource{}
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
			"unset_attributes": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: DescServerUnsetAttributes,
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

	// unset_attributes is provider-only metadata (not a PBS attribute); carry it
	// forward from prior state so a refresh does not clear it. State otherwise
	// reflects the real PBS values (private state tracks completed one-shot unsets).
	updatedState.UnsetAttributes = currentState.UnsetAttributes

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
	updatedServer, err := r.client.UpdatePbsServer(server)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to update the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Error: "+err.Error(),
		)

		return
	}

	// UpdatePbsServer already returns the re-read server (including computed
	// fields), so use it directly rather than issuing a second read that could
	// fail after the update itself succeeded.
	updatedData := createServerModel(updatedServer)

	// Preserve user-provided ACL formats from plan where possible
	preserveUserServerAclFormat(&planData, &updatedData)

	// unset_attributes is provider-only metadata; persist the configured value.
	updatedData.UnsetAttributes = planData.UnsetAttributes

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Record the applied unsets in private state so a persistently-configured
	// unset_attributes is not repeatedly re-issued on future plans. The recorded
	// set is exactly the current unset_attributes, so removing a name re-arms it.
	var appliedNames []string
	if !planData.UnsetAttributes.IsNull() && !planData.UnsetAttributes.IsUnknown() {
		resp.Diagnostics.Append(planData.UnsetAttributes.ElementsAs(ctx, &appliedNames, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	resp.Diagnostics.Append(writeAppliedServerUnsets(ctx, resp.Private, appliedNames)...)
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

// ModifyPlan implements explicit attribute removal through the unset_attributes set.
// Because every server attribute is Optional+Computed with UseStateForUnknown,
// omitting an attribute preserves its imported value. To actively remove an
// attribute, list its name in unset_attributes: a scalar's planned value is marked
// unknown (PBS computes the value after the reset) and a map's planned value is
// marked null (its entries are removed), so Update emits the appropriate qmgr unset.
//
// unset_attributes describes a one-shot reset. Once an attribute's unset has been
// applied it is recorded in private state and skipped on subsequent plans, so a
// persistently-configured unset_attributes does not repeatedly re-issue the reset
// and Terraform state keeps reflecting the real PBS value. The set and its elements
// must be known at plan time: an unknown or null value is rejected rather than
// guessed at, which keeps plan and apply consistent.
func (r *serverResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// No modification on destroy (null plan) or on create/import (null prior state).
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

	var unsetAttributes types.Set
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("unset_attributes"), &unsetAttributes)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if unsetAttributes.IsNull() {
		return
	}

	// The exact attribute names must be determinable at plan time (similar to
	// for_each keys). Reject an unknown set or unknown/null elements rather than
	// guessing, which would risk an inconsistent apply or an unintended bulk unset.
	if unsetAttributes.IsUnknown() {
		resp.Diagnostics.AddAttributeError(path.Root("unset_attributes"),
			"unset_attributes must be known at plan time",
			"unset_attributes cannot be derived from a value that is unknown until apply. Use literal attribute names.")
		return
	}
	for _, el := range unsetAttributes.Elements() {
		if el.IsUnknown() {
			resp.Diagnostics.AddAttributeError(path.Root("unset_attributes"),
				"unset_attributes must be known at plan time",
				"unset_attributes contains a value that is unknown until apply. Use literal attribute names.")
			return
		}
		if el.IsNull() {
			resp.Diagnostics.AddAttributeError(path.Root("unset_attributes"),
				"unset_attributes must not contain null values",
				"unset_attributes contains a null value. Provide only literal attribute names.")
			return
		}
	}

	var names []string
	resp.Diagnostics.Append(unsetAttributes.ElementsAs(ctx, &names, false)...)
	if resp.Diagnostics.HasError() || len(names) == 0 {
		return
	}

	applied, appliedDiags := readAppliedServerUnsets(ctx, req.Private)
	resp.Diagnostics.Append(appliedDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	planObj := map[string]tftypes.Value{}
	if err := req.Plan.Raw.As(&planObj); err != nil {
		resp.Diagnostics.AddError("Unable to read plan", err.Error())
		return
	}
	configObj := map[string]tftypes.Value{}
	if err := req.Config.Raw.As(&configObj); err != nil {
		resp.Diagnostics.AddError("Unable to read configuration", err.Error())
		return
	}
	stateObj := map[string]tftypes.Value{}
	if err := req.State.Raw.As(&stateObj); err != nil {
		resp.Diagnostics.AddError("Unable to read state", err.Error())
		return
	}

	modified := false
	for _, name := range names {
		planVal, ok := planObj[name]
		if !ok {
			resp.Diagnostics.AddAttributeError(path.Root("unset_attributes"),
				"Unknown server attribute",
				fmt.Sprintf("%q is not a pbs_server attribute and cannot be unset.", name))
			continue
		}
		if !isUnsettableServerAttribute(name) {
			resp.Diagnostics.AddAttributeError(path.Root("unset_attributes"),
				"Attribute cannot be unset",
				fmt.Sprintf("%q cannot be listed in unset_attributes. Identity, computed (normalized) and PBS read-only attributes cannot be unset.", name))
			continue
		}
		// Only a conflict when the attribute is configured with a known, non-null
		// value; an unknown config value may still resolve to null at apply.
		if cfgVal, ok := configObj[name]; ok && cfgVal.IsKnown() && !cfgVal.IsNull() {
			resp.Diagnostics.AddAttributeError(path.Root("unset_attributes"),
				"Attribute both set and unset",
				fmt.Sprintf("%q is set in configuration and also listed in unset_attributes. Remove it from one of them.", name))
			continue
		}
		// One-shot: the reset has already been applied, so leave the attribute
		// pinned to its real PBS value instead of re-issuing the unset.
		if _, ok := applied[name]; ok {
			continue
		}
		// Nothing to unset if the attribute is already absent from state.
		if stateVal, ok := stateObj[name]; !ok || stateVal.IsNull() {
			continue
		}
		// Mark the attribute unknown: PBS decides the value after the reset (some
		// attributes vanish, others revert to a default), so it is only known after
		// apply. Update converts an unknown value to a qmgr unset.
		planObj[name] = tftypes.NewValue(planVal.Type(), tftypes.UnknownValue)
		modified = true
	}

	if resp.Diagnostics.HasError() || !modified {
		return
	}

	resp.Plan.Raw = tftypes.NewValue(req.Plan.Raw.Type(), planObj)
}

// serverUnsetAppliedPrivateKey is the private-state key tracking which
// unset_attributes entries have already been applied (a one-shot reset).
const serverUnsetAppliedPrivateKey = "unset_applied"

// privateStateGetter and privateStateSetter abstract the framework's private state
// (whose concrete type is internal) so the helpers stay testable.
type privateStateGetter interface {
	GetKey(ctx context.Context, key string) ([]byte, diag.Diagnostics)
}

type privateStateSetter interface {
	SetKey(ctx context.Context, key string, value []byte) diag.Diagnostics
}

// readAppliedServerUnsets returns the set of attribute names whose unset has already
// been applied, decoded from private state.
func readAppliedServerUnsets(ctx context.Context, private privateStateGetter) (map[string]struct{}, diag.Diagnostics) {
	applied := map[string]struct{}{}
	if private == nil {
		return applied, nil
	}

	raw, diags := private.GetKey(ctx, serverUnsetAppliedPrivateKey)
	if diags.HasError() || len(raw) == 0 {
		return applied, diags
	}

	var names []string
	if err := json.Unmarshal(raw, &names); err != nil {
		diags.AddError("Unable to read private state",
			fmt.Sprintf("Could not decode unset_attributes tracking data: %s", err))
		return applied, diags
	}
	for _, name := range names {
		applied[name] = struct{}{}
	}
	return applied, diags
}

// writeAppliedServerUnsets records the given attribute names as applied unsets in
// private state, replacing any previous value so that names dropped from
// unset_attributes are re-armed for a future unset.
func writeAppliedServerUnsets(ctx context.Context, private privateStateSetter, names []string) diag.Diagnostics {
	var diags diag.Diagnostics
	if private == nil {
		return diags
	}

	sorted := append([]string(nil), names...)
	sort.Strings(sorted)
	raw, err := json.Marshal(sorted)
	if err != nil {
		diags.AddError("Unable to write private state",
			fmt.Sprintf("Could not encode unset_attributes tracking data: %s", err))
		return diags
	}
	return private.SetKey(ctx, serverUnsetAppliedPrivateKey, raw)
}

// isUnsettableServerAttribute reports whether an attribute may be listed in
// unset_attributes. Identity, computed (normalized) and PBS read-only attributes
// cannot be unset; all other settable attributes (scalars and maps) qualify.
func isUnsettableServerAttribute(name string) bool {
	switch name {
	case "id", "name", "unset_attributes":
		return false
	// power_provisioning is PBS-managed and rejects manager writes, so a qmgr
	// unset of it always fails.
	case "power_provisioning":
		return false
	}
	return !strings.HasSuffix(name, "_normalized")
}
