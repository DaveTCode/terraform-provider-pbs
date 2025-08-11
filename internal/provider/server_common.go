package provider

import (
	"context"
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type serverModel struct {
	ID                            types.String            `tfsdk:"id"`
	AclHostEnable                 types.Bool              `tfsdk:"acl_host_enable"`
	AclHostMomsEnable             types.Bool              `tfsdk:"acl_host_moms_enable"`
	AclHosts                      types.String            `tfsdk:"acl_hosts"`
	AclHostsNormalized            types.String            `tfsdk:"acl_hosts_normalized"`
	AclResvGroupEnable            types.Bool              `tfsdk:"acl_resv_group_enable"`
	AclResvGroups                 types.String            `tfsdk:"acl_resv_groups"`
	AclResvGroupsNormalized       types.String            `tfsdk:"acl_resv_groups_normalized"`
	AclResvHostEnable             types.Bool              `tfsdk:"acl_resv_host_enable"`
	AclResvHosts                  types.String            `tfsdk:"acl_resv_hosts"`
	AclResvHostsNormalized        types.String            `tfsdk:"acl_resv_hosts_normalized"`
	AclResvUserEnable             types.Bool              `tfsdk:"acl_resv_user_enable"`
	AclResvUsers                  types.String            `tfsdk:"acl_resv_users"`
	AclResvUsersNormalized        types.String            `tfsdk:"acl_resv_users_normalized"`
	AclRoots                      types.String            `tfsdk:"acl_roots"`
	AclRootsNormalized            types.String            `tfsdk:"acl_roots_normalized"`
	AclUserEnable                 types.Bool              `tfsdk:"acl_user_enable"`
	AclUsers                      types.String            `tfsdk:"acl_users"`
	AclUsersNormalized            types.String            `tfsdk:"acl_users_normalized"`
	BackfillDepth                 types.Int32             `tfsdk:"backfill_depth"`
	Comment                       types.String            `tfsdk:"comment"`
	DefaultChunk                  map[string]types.String `tfsdk:"default_chunk"`
	DefaultQdelArguments          types.String            `tfsdk:"default_qdel_arguments"`
	DefaultQsubArguments          types.String            `tfsdk:"default_qsub_arguments"`
	DefaultQueue                  types.String            `tfsdk:"default_queue"`
	EligibleTimeEnable            types.Bool              `tfsdk:"eligible_time_enable"`
	ElimOnSubjobs                 types.Bool              `tfsdk:"elim_on_subjobs"`
	Flatuid                       types.Bool              `tfsdk:"flatuid"`
	JobHistoryDuration            types.String            `tfsdk:"job_history_duration"`
	JobHistoryEnable              types.Bool              `tfsdk:"job_history_enable"`
	JobRequeueTimeout             types.String            `tfsdk:"job_requeue_timeout"`
	JobSortFormula                types.String            `tfsdk:"job_sort_formula"`
	JobscriptMaxSize              types.String            `tfsdk:"jobscript_max_size"`
	LogEvents                     types.Int32             `tfsdk:"log_events"`
	Mailer                        types.String            `tfsdk:"mailer"`
	MailFrom                      types.String            `tfsdk:"mail_from"`
	Managers                      types.String            `tfsdk:"managers"`
	MaxArraySize                  types.Int32             `tfsdk:"max_array_size"`
	MaxConcurrentProvision        types.Int32             `tfsdk:"max_concurrent_provision"`
	MaxGroupRes                   map[string]types.String `tfsdk:"max_group_res"`
	MaxGroupResSoft               map[string]types.String `tfsdk:"max_group_res_soft"`
	MaxGroupRun                   types.Int32             `tfsdk:"max_group_run"`
	MaxGroupRunSoft               types.Int32             `tfsdk:"max_group_run_soft"`
	MaxJobSequenceId              types.Int64             `tfsdk:"max_job_sequence_id"`
	MaxQueued                     map[string]types.String `tfsdk:"max_queued"`
	MaxQueuedRes                  map[string]types.String `tfsdk:"max_queued_res"`
	MaxRun                        map[string]types.String `tfsdk:"max_run"`
	MaxRunRes                     map[string]types.String `tfsdk:"max_run_res"`
	MaxRunResSoft                 map[string]types.String `tfsdk:"max_run_res_soft"`
	MaxRunSoft                    map[string]types.String `tfsdk:"max_run_soft"`
	MaxRunning                    types.Int32             `tfsdk:"max_running"`
	MaxUserRes                    map[string]types.String `tfsdk:"max_user_res"`
	MaxUserResSoft                map[string]types.String `tfsdk:"max_user_res_soft"`
	MaxUserRun                    types.Int32             `tfsdk:"max_user_run"`
	MaxUserRunSoft                types.Int32             `tfsdk:"max_user_run_soft"`
	Name                          types.String            `tfsdk:"name"`
	NodeFailRequeue               types.Int32             `tfsdk:"node_fail_requeue"`
	NodeGroupEnable               types.Bool              `tfsdk:"node_group_enable"`
	NodeGroupKey                  types.String            `tfsdk:"node_group_key"`
	Operators                     types.String            `tfsdk:"operators"`
	PbsLicenseInfo                types.String            `tfsdk:"pbs_license_info"`
	PbsLicenseLingerTime          types.Int32             `tfsdk:"pbs_license_linger_time"`
	PbsLicenseMax                 types.Int32             `tfsdk:"pbs_license_max"`
	PbsLicenseMin                 types.Int32             `tfsdk:"pbs_license_min"`
	PowerProvisioning             types.Bool              `tfsdk:"power_provisioning"`
	PythonGcMinInterval           types.Int32             `tfsdk:"python_gc_min_interval"`
	PythonRestartMaxHooks         types.Int32             `tfsdk:"python_restart_max_hooks"`
	PythonRestartMaxObjects       types.Int32             `tfsdk:"python_restart_max_objects"`
	PythonRestartMinInterval      types.String            `tfsdk:"python_restart_min_interval"`
	QueryOtherJobs                types.Bool              `tfsdk:"query_other_jobs"`
	QueuedJobsThreshold           types.String            `tfsdk:"queued_jobs_threshold"`
	QueuedJobsThresholdRes        types.String            `tfsdk:"queued_jobs_threshold_res"`
	ReserveRetryInit              types.Int32             `tfsdk:"reserve_retry_init"`
	ReserveRetryTime              types.Int32             `tfsdk:"reserve_retry_time"`
	ResourcesAvailable            map[string]types.String `tfsdk:"resources_available"`
	ResourcesDefault              map[string]types.String `tfsdk:"resources_default"`
	ResourcesMax                  map[string]types.String `tfsdk:"resources_max"`
	RestrictResToReleaseOnSuspend types.String            `tfsdk:"restrict_res_to_release_on_suspend"`
	ResvEnable                    types.Bool              `tfsdk:"resv_enable"`
	ResvPostProcessingTime        types.String            `tfsdk:"resv_post_processing_time"`
	RppHighwater                  types.Int32             `tfsdk:"rpp_highwater"`
	RppMaxPktCheck                types.Int32             `tfsdk:"rpp_max_pkt_check"`
	RppRetry                      types.Int32             `tfsdk:"rpp_retry"`
	SchedulerIteration            types.Int32             `tfsdk:"scheduler_iteration"`
	WebapiAuthIssuers             types.String            `tfsdk:"webapi_auth_issuers"`
	WebapiEnable                  types.Bool              `tfsdk:"webapi_enable"`
	WebapiOidcClientid            types.String            `tfsdk:"webapi_oidc_clientid"`
	WebapiOidcProviderUrl         types.String            `tfsdk:"webapi_oidc_provider_url"`
}

func (m serverModel) ToPbsServer(ctx context.Context) pbsclient.PbsServer {
	server := pbsclient.PbsServer{
		Name: m.Name.ValueString(),
	}

	// Set pointer fields using utility functions for null checking
	SetBoolPointerIfNotNull(m.AclHostEnable, &server.AclHostEnable)
	SetBoolPointerIfNotNull(m.AclHostMomsEnable, &server.AclHostMomsEnable)
	SetStringPointerIfNotNull(m.AclHosts, &server.AclHosts)
	SetBoolPointerIfNotNull(m.AclResvGroupEnable, &server.AclResvGroupEnable)
	SetStringPointerIfNotNull(m.AclResvGroups, &server.AclResvGroups)
	SetBoolPointerIfNotNull(m.AclResvHostEnable, &server.AclResvHostEnable)
	SetStringPointerIfNotNull(m.AclResvHosts, &server.AclResvHosts)
	SetBoolPointerIfNotNull(m.AclResvUserEnable, &server.AclResvUserEnable)
	SetStringPointerIfNotNull(m.AclResvUsers, &server.AclResvUsers)
	SetStringPointerIfNotNull(m.AclRoots, &server.AclRoots)
	SetBoolPointerIfNotNull(m.AclUserEnable, &server.AclUserEnable)
	SetStringPointerIfNotNull(m.AclUsers, &server.AclUsers)
	SetInt32PointerIfNotNull(m.BackfillDepth, &server.BackfillDepth)
	SetStringPointerIfNotNull(m.Comment, &server.Comment)
	SetStringPointerIfNotNull(m.DefaultQdelArguments, &server.DefaultQdelArguments)
	SetStringPointerIfNotNull(m.DefaultQsubArguments, &server.DefaultQsubArguments)
	SetStringPointerIfNotNull(m.DefaultQueue, &server.DefaultQueue)
	SetBoolPointerIfNotNull(m.EligibleTimeEnable, &server.EligibleTimeEnable)
	SetBoolPointerIfNotNull(m.ElimOnSubjobs, &server.ElimOnSubjobs)
	SetBoolPointerIfNotNull(m.Flatuid, &server.Flatuid)
	SetStringPointerIfNotNull(m.JobHistoryDuration, &server.JobHistoryDuration)
	SetBoolPointerIfNotNull(m.JobHistoryEnable, &server.JobHistoryEnable)
	SetStringPointerIfNotNull(m.JobRequeueTimeout, &server.JobRequeueTimeout)
	SetStringPointerIfNotNull(m.JobSortFormula, &server.JobSortFormula)
	SetStringPointerIfNotNull(m.JobscriptMaxSize, &server.JobscriptMaxSize)
	SetInt32PointerIfNotNull(m.LogEvents, &server.LogEvents)
	SetStringPointerIfNotNull(m.Mailer, &server.Mailer)
	SetStringPointerIfNotNull(m.MailFrom, &server.MailFrom)
	SetStringPointerIfNotNull(m.Managers, &server.Managers)
	SetInt32PointerIfNotNull(m.MaxArraySize, &server.MaxArraySize)
	SetInt32PointerIfNotNull(m.MaxConcurrentProvision, &server.MaxConcurrentProvision)
	SetInt32PointerIfNotNull(m.MaxGroupRun, &server.MaxGroupRun)
	SetInt32PointerIfNotNull(m.MaxGroupRunSoft, &server.MaxGroupRunSoft)
	SetInt64PointerIfNotNull(m.MaxJobSequenceId, &server.MaxJobSequenceId)
	SetInt32PointerIfNotNull(m.MaxRunning, &server.MaxRunning)
	SetInt32PointerIfNotNull(m.MaxUserRun, &server.MaxUserRun)
	SetInt32PointerIfNotNull(m.MaxUserRunSoft, &server.MaxUserRunSoft)
	SetInt32PointerIfNotNull(m.NodeFailRequeue, &server.NodeFailRequeue)
	SetBoolPointerIfNotNull(m.NodeGroupEnable, &server.NodeGroupEnable)
	SetStringPointerIfNotNull(m.NodeGroupKey, &server.NodeGroupKey)
	SetStringPointerIfNotNull(m.Operators, &server.Operators)
	SetStringPointerIfNotNull(m.PbsLicenseInfo, &server.PbsLicenseInfo)
	SetInt32PointerIfNotNull(m.PbsLicenseLingerTime, &server.PbsLicenseLingerTime)
	SetInt32PointerIfNotNull(m.PbsLicenseMax, &server.PbsLicenseMax)
	SetInt32PointerIfNotNull(m.PbsLicenseMin, &server.PbsLicenseMin)
	SetBoolPointerIfNotNull(m.PowerProvisioning, &server.PowerProvisioning)
	SetInt32PointerIfNotNull(m.PythonGcMinInterval, &server.PythonGcMinInterval)
	SetInt32PointerIfNotNull(m.PythonRestartMaxHooks, &server.PythonRestartMaxPbsServers)
	SetInt32PointerIfNotNull(m.PythonRestartMaxObjects, &server.PythonRestartMaxObjects)
	SetStringPointerIfNotNull(m.PythonRestartMinInterval, &server.PythonRestartMinInterval)
	SetBoolPointerIfNotNull(m.QueryOtherJobs, &server.QueryOtherJobs)
	SetStringPointerIfNotNull(m.QueuedJobsThreshold, &server.QueuedJobsThreshold)
	SetStringPointerIfNotNull(m.QueuedJobsThresholdRes, &server.QueuedJobsThresholdRes)
	SetInt32PointerIfNotNull(m.ReserveRetryInit, &server.ReserveRetryInit)
	SetInt32PointerIfNotNull(m.ReserveRetryTime, &server.ReserveRetryTime)
	SetStringPointerIfNotNull(m.RestrictResToReleaseOnSuspend, &server.RestrictResToReleaseOnSuspend)
	SetBoolPointerIfNotNull(m.ResvEnable, &server.ResvEnable)
	SetStringPointerIfNotNull(m.ResvPostProcessingTime, &server.ResvPostProcessingTime)
	SetInt32PointerIfNotNull(m.RppHighwater, &server.RppHighwater)
	SetInt32PointerIfNotNull(m.RppMaxPktCheck, &server.RppMaxPktCheck)
	SetInt32PointerIfNotNull(m.RppRetry, &server.RppRetry)
	SetInt32PointerIfNotNull(m.SchedulerIteration, &server.SchedulerIteration)
	SetStringPointerIfNotNull(m.WebapiAuthIssuers, &server.WebapiAuthIssuers)
	SetBoolPointerIfNotNull(m.WebapiEnable, &server.WebapiEnable)
	SetStringPointerIfNotNull(m.WebapiOidcClientid, &server.WebapiOidcClientid)
	SetStringPointerIfNotNull(m.WebapiOidcProviderUrl, &server.WebapiOidcProviderUrl)

	// Convert map fields using utility functions
	server.DefaultChunk = ConvertTypesStringMap(m.DefaultChunk)
	server.ResourcesAvailable = ConvertTypesStringMap(m.ResourcesAvailable)
	server.ResourcesDefault = ConvertTypesStringMap(m.ResourcesDefault)
	server.ResourcesMax = ConvertTypesStringMap(m.ResourcesMax)

	// Convert limit attribute maps from Terraform types to Go maps (only if not empty)
	ConvertTypesStringMapIfNotEmpty(m.MaxGroupRes, &server.MaxGroupRes)
	ConvertTypesStringMapIfNotEmpty(m.MaxGroupResSoft, &server.MaxGroupResSoft)
	ConvertTypesStringMapIfNotEmpty(m.MaxQueued, &server.MaxQueued)
	ConvertTypesStringMapIfNotEmpty(m.MaxQueuedRes, &server.MaxQueuedRes)
	ConvertTypesStringMapIfNotEmpty(m.MaxRun, &server.MaxRun)
	ConvertTypesStringMapIfNotEmpty(m.MaxRunRes, &server.MaxRunRes)
	ConvertTypesStringMapIfNotEmpty(m.MaxRunResSoft, &server.MaxRunResSoft)
	ConvertTypesStringMapIfNotEmpty(m.MaxRunSoft, &server.MaxRunSoft)
	ConvertTypesStringMapIfNotEmpty(m.MaxUserRes, &server.MaxUserRes)
	ConvertTypesStringMapIfNotEmpty(m.MaxUserResSoft, &server.MaxUserResSoft)

	return server
}

func createServerModel(server pbsclient.PbsServer) serverModel {
	model := serverModel{
		ID:   types.StringValue(server.Name), // Use name as ID
		Name: types.StringValue(server.Name),
	}

	model.AclHostEnable = types.BoolPointerValue(server.AclHostEnable)
	model.AclHostMomsEnable = types.BoolPointerValue(server.AclHostMomsEnable)
	model.AclHosts = types.StringPointerValue(server.AclHosts)
	addNormalizedAclField(server.AclHosts, &model.AclHostsNormalized)
	model.AclResvGroupEnable = types.BoolPointerValue(server.AclResvGroupEnable)
	model.AclResvGroups = types.StringPointerValue(server.AclResvGroups)
	addNormalizedAclField(server.AclResvGroups, &model.AclResvGroupsNormalized)
	model.AclResvHostEnable = types.BoolPointerValue(server.AclResvHostEnable)
	model.AclResvHosts = types.StringPointerValue(server.AclResvHosts)
	addNormalizedAclField(server.AclResvHosts, &model.AclResvHostsNormalized)
	model.AclResvUserEnable = types.BoolPointerValue(server.AclResvUserEnable)
	model.AclResvUsers = types.StringPointerValue(server.AclResvUsers)
	addNormalizedAclField(server.AclResvUsers, &model.AclResvUsersNormalized)
	model.AclRoots = types.StringPointerValue(server.AclRoots)
	addNormalizedAclField(server.AclRoots, &model.AclRootsNormalized)
	model.AclUserEnable = types.BoolPointerValue(server.AclUserEnable)
	model.AclUsers = types.StringPointerValue(server.AclUsers)
	addNormalizedAclField(server.AclUsers, &model.AclUsersNormalized)
	model.BackfillDepth = types.Int32PointerValue(server.BackfillDepth)
	model.Comment = types.StringPointerValue(server.Comment)
	model.DefaultQdelArguments = types.StringPointerValue(server.DefaultQdelArguments)
	model.DefaultQsubArguments = types.StringPointerValue(server.DefaultQsubArguments)
	model.DefaultQueue = types.StringPointerValue(server.DefaultQueue)
	model.EligibleTimeEnable = types.BoolPointerValue(server.EligibleTimeEnable)
	model.ElimOnSubjobs = types.BoolPointerValue(server.ElimOnSubjobs)
	model.Flatuid = types.BoolPointerValue(server.Flatuid)
	model.JobHistoryDuration = types.StringPointerValue(server.JobHistoryDuration)
	model.JobHistoryEnable = types.BoolPointerValue(server.JobHistoryEnable)
	model.JobRequeueTimeout = types.StringPointerValue(server.JobRequeueTimeout)
	model.JobSortFormula = types.StringPointerValue(server.JobSortFormula)
	model.JobscriptMaxSize = types.StringPointerValue(server.JobscriptMaxSize)
	model.LogEvents = types.Int32PointerValue(server.LogEvents)
	model.Mailer = types.StringPointerValue(server.Mailer)
	model.MailFrom = types.StringPointerValue(server.MailFrom)
	model.Managers = types.StringPointerValue(server.Managers)
	model.MaxArraySize = types.Int32PointerValue(server.MaxArraySize)
	model.MaxConcurrentProvision = types.Int32PointerValue(server.MaxConcurrentProvision)
	// Convert map attributes for limit settings
	if server.MaxGroupRes != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxGroupRes {
			elements[k] = types.StringValue(v)
		}
		model.MaxGroupRes = elements
	}
	if server.MaxGroupResSoft != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxGroupResSoft {
			elements[k] = types.StringValue(v)
		}
		model.MaxGroupResSoft = elements
	}
	model.MaxGroupRun = types.Int32PointerValue(server.MaxGroupRun)
	model.MaxGroupRunSoft = types.Int32PointerValue(server.MaxGroupRunSoft)
	model.MaxJobSequenceId = types.Int64PointerValue(server.MaxJobSequenceId)
	if server.MaxQueued != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxQueued {
			elements[k] = types.StringValue(v)
		}
		model.MaxQueued = elements
	}
	if server.MaxQueuedRes != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxQueuedRes {
			elements[k] = types.StringValue(v)
		}
		model.MaxQueuedRes = elements
	}
	if server.MaxRun != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxRun {
			elements[k] = types.StringValue(v)
		}
		model.MaxRun = elements
	}
	if server.MaxRunRes != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxRunRes {
			elements[k] = types.StringValue(v)
		}
		model.MaxRunRes = elements
	}
	if server.MaxRunResSoft != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxRunResSoft {
			elements[k] = types.StringValue(v)
		}
		model.MaxRunResSoft = elements
	}
	if server.MaxRunSoft != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxRunSoft {
			elements[k] = types.StringValue(v)
		}
		model.MaxRunSoft = elements
	}
	model.MaxRunning = types.Int32PointerValue(server.MaxRunning)
	if server.MaxUserRes != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxUserRes {
			elements[k] = types.StringValue(v)
		}
		model.MaxUserRes = elements
	}
	if server.MaxUserResSoft != nil {
		elements := make(map[string]types.String)
		for k, v := range server.MaxUserResSoft {
			elements[k] = types.StringValue(v)
		}
		model.MaxUserResSoft = elements
	}
	model.MaxUserRun = types.Int32PointerValue(server.MaxUserRun)
	model.MaxUserRunSoft = types.Int32PointerValue(server.MaxUserRunSoft)
	model.NodeFailRequeue = types.Int32PointerValue(server.NodeFailRequeue)
	model.NodeGroupEnable = types.BoolPointerValue(server.NodeGroupEnable)
	model.NodeGroupKey = types.StringPointerValue(server.NodeGroupKey)
	model.Operators = types.StringPointerValue(server.Operators)
	model.PbsLicenseInfo = types.StringPointerValue(server.PbsLicenseInfo)
	model.PbsLicenseLingerTime = types.Int32PointerValue(server.PbsLicenseLingerTime)
	model.PbsLicenseMax = types.Int32PointerValue(server.PbsLicenseMax)
	model.PbsLicenseMin = types.Int32PointerValue(server.PbsLicenseMin)
	model.PowerProvisioning = types.BoolPointerValue(server.PowerProvisioning)
	model.PythonGcMinInterval = types.Int32PointerValue(server.PythonGcMinInterval)
	model.PythonRestartMaxHooks = types.Int32PointerValue(server.PythonRestartMaxPbsServers)
	model.PythonRestartMaxObjects = types.Int32PointerValue(server.PythonRestartMaxObjects)
	model.PythonRestartMinInterval = types.StringPointerValue(server.PythonRestartMinInterval)
	model.QueryOtherJobs = types.BoolPointerValue(server.QueryOtherJobs)
	model.QueuedJobsThreshold = types.StringPointerValue(server.QueuedJobsThreshold)
	model.QueuedJobsThresholdRes = types.StringPointerValue(server.QueuedJobsThresholdRes)
	model.ReserveRetryInit = types.Int32PointerValue(server.ReserveRetryInit)
	model.ReserveRetryTime = types.Int32PointerValue(server.ReserveRetryTime)
	model.RestrictResToReleaseOnSuspend = types.StringPointerValue(server.RestrictResToReleaseOnSuspend)
	model.ResvEnable = types.BoolPointerValue(server.ResvEnable)
	model.ResvPostProcessingTime = types.StringPointerValue(server.ResvPostProcessingTime)
	model.RppHighwater = types.Int32PointerValue(server.RppHighwater)
	model.RppMaxPktCheck = types.Int32PointerValue(server.RppMaxPktCheck)
	model.RppRetry = types.Int32PointerValue(server.RppRetry)
	model.SchedulerIteration = types.Int32PointerValue(server.SchedulerIteration)
	model.WebapiAuthIssuers = types.StringPointerValue(server.WebapiAuthIssuers)
	model.WebapiEnable = types.BoolPointerValue(server.WebapiEnable)
	model.WebapiOidcClientid = types.StringPointerValue(server.WebapiOidcClientid)
	model.WebapiOidcProviderUrl = types.StringPointerValue(server.WebapiOidcProviderUrl)

	if server.DefaultChunk != nil {
		model.DefaultChunk = convertStringMapToTypesStringMap(server.DefaultChunk)
	}
	if server.ResourcesAvailable != nil {
		model.ResourcesAvailable = convertStringMapToTypesStringMap(server.ResourcesAvailable)
	}
	if server.ResourcesDefault != nil {
		model.ResourcesDefault = convertStringMapToTypesStringMap(server.ResourcesDefault)
	}
	if server.ResourcesMax != nil {
		model.ResourcesMax = convertStringMapToTypesStringMap(server.ResourcesMax)
	}

	return model
}

// preserveUserServerAclFormat preserves user-provided ACL field formats from plan in the result model.
func preserveUserServerAclFormat(planModel, resultModel *serverModel) {
	planFields := []AclFieldPair{
		{UserField: planModel.AclHosts, NormalizedField: planModel.AclHostsNormalized},
		{UserField: planModel.AclResvGroups, NormalizedField: planModel.AclResvGroupsNormalized},
		{UserField: planModel.AclResvHosts, NormalizedField: planModel.AclResvHostsNormalized},
		{UserField: planModel.AclResvUsers, NormalizedField: planModel.AclResvUsersNormalized},
		{UserField: planModel.AclRoots, NormalizedField: planModel.AclRootsNormalized},
		{UserField: planModel.AclUsers, NormalizedField: planModel.AclUsersNormalized},
	}

	resultFields := []AclFieldPair{
		{UserField: resultModel.AclHosts, NormalizedField: resultModel.AclHostsNormalized},
		{UserField: resultModel.AclResvGroups, NormalizedField: resultModel.AclResvGroupsNormalized},
		{UserField: resultModel.AclResvHosts, NormalizedField: resultModel.AclResvHostsNormalized},
		{UserField: resultModel.AclResvUsers, NormalizedField: resultModel.AclResvUsersNormalized},
		{UserField: resultModel.AclRoots, NormalizedField: resultModel.AclRootsNormalized},
		{UserField: resultModel.AclUsers, NormalizedField: resultModel.AclUsersNormalized},
	}

	preserveUserAclFormats(planFields, resultFields)

	// Update the result model with preserved values.
	resultModel.AclHosts = resultFields[0].UserField
	resultModel.AclResvGroups = resultFields[1].UserField
	resultModel.AclResvHosts = resultFields[2].UserField
	resultModel.AclResvUsers = resultFields[3].UserField
	resultModel.AclRoots = resultFields[4].UserField
	resultModel.AclUsers = resultFields[5].UserField
}

// preserveUserServerAclFormatFromState preserves user-provided ACL field formats from state when semantically equivalent.
func preserveUserServerAclFormatFromState(state, updatedState *serverModel) {
	stateFields := []AclFieldPair{
		{UserField: state.AclHosts, NormalizedField: state.AclHostsNormalized},
		{UserField: state.AclResvGroups, NormalizedField: state.AclResvGroupsNormalized},
		{UserField: state.AclResvHosts, NormalizedField: state.AclResvHostsNormalized},
		{UserField: state.AclResvUsers, NormalizedField: state.AclResvUsersNormalized},
		{UserField: state.AclRoots, NormalizedField: state.AclRootsNormalized},
		{UserField: state.AclUsers, NormalizedField: state.AclUsersNormalized},
	}

	updatedFields := []AclFieldPair{
		{UserField: updatedState.AclHosts, NormalizedField: updatedState.AclHostsNormalized},
		{UserField: updatedState.AclResvGroups, NormalizedField: updatedState.AclResvGroupsNormalized},
		{UserField: updatedState.AclResvHosts, NormalizedField: updatedState.AclResvHostsNormalized},
		{UserField: updatedState.AclResvUsers, NormalizedField: updatedState.AclResvUsersNormalized},
		{UserField: updatedState.AclRoots, NormalizedField: updatedState.AclRootsNormalized},
		{UserField: updatedState.AclUsers, NormalizedField: updatedState.AclUsersNormalized},
	}

	preserveUserAclFormatsFromState(stateFields, updatedFields)

	// Update the updated state with preserved values.
	updatedState.AclHosts = updatedFields[0].UserField
	updatedState.AclResvGroups = updatedFields[1].UserField
	updatedState.AclResvHosts = updatedFields[2].UserField
	updatedState.AclResvUsers = updatedFields[3].UserField
	updatedState.AclRoots = updatedFields[4].UserField
	updatedState.AclUsers = updatedFields[5].UserField
}
