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
	AclResvGroupEnable            types.Bool              `tfsdk:"acl_resv_group_enable"`
	AclResvGroups                 types.String            `tfsdk:"acl_resv_groups"`
	AclResvHostEnable             types.Bool              `tfsdk:"acl_resv_host_enable"`
	AclResvHosts                  types.String            `tfsdk:"acl_resv_hosts"`
	AclResvUserEnable             types.Bool              `tfsdk:"acl_resv_user_enable"`
	AclResvUsers                  types.String            `tfsdk:"acl_resv_users"`
	AclRoots                      types.String            `tfsdk:"acl_roots"`
	AclUserEnable                 types.Bool              `tfsdk:"acl_user_enable"`
	AclUsers                      types.String            `tfsdk:"acl_users"`
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
	PythonRestartMaxPbsServers    types.Int32             `tfsdk:"python_restart_max_pbs_servers"`
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
		Name:                          m.Name.ValueString(),
		AclHostEnable:                 m.AclHostEnable.ValueBoolPointer(),
		AclHostMomsEnable:             m.AclHostMomsEnable.ValueBoolPointer(),
		AclHosts:                      m.AclHosts.ValueStringPointer(),
		AclResvGroupEnable:            m.AclResvGroupEnable.ValueBoolPointer(),
		AclResvGroups:                 m.AclResvGroups.ValueStringPointer(),
		AclResvHostEnable:             m.AclResvHostEnable.ValueBoolPointer(),
		AclResvHosts:                  m.AclResvHosts.ValueStringPointer(),
		AclResvUserEnable:             m.AclResvUserEnable.ValueBoolPointer(),
		AclResvUsers:                  m.AclResvUsers.ValueStringPointer(),
		AclRoots:                      m.AclRoots.ValueStringPointer(),
		AclUserEnable:                 m.AclUserEnable.ValueBoolPointer(),
		AclUsers:                      m.AclUsers.ValueStringPointer(),
		BackfillDepth:                 m.BackfillDepth.ValueInt32Pointer(),
		Comment:                       m.Comment.ValueStringPointer(),
		DefaultQdelArguments:          m.DefaultQdelArguments.ValueStringPointer(),
		DefaultQsubArguments:          m.DefaultQsubArguments.ValueStringPointer(),
		DefaultQueue:                  m.DefaultQueue.ValueStringPointer(),
		EligibleTimeEnable:            m.EligibleTimeEnable.ValueBoolPointer(),
		ElimOnSubjobs:                 m.ElimOnSubjobs.ValueBoolPointer(),
		Flatuid:                       m.Flatuid.ValueBoolPointer(),
		JobHistoryDuration:            m.JobHistoryDuration.ValueStringPointer(),
		JobHistoryEnable:              m.JobHistoryEnable.ValueBoolPointer(),
		JobRequeueTimeout:             m.JobRequeueTimeout.ValueStringPointer(),
		JobSortFormula:                m.JobSortFormula.ValueStringPointer(),
		JobscriptMaxSize:              m.JobscriptMaxSize.ValueStringPointer(),
		LogEvents:                     m.LogEvents.ValueInt32Pointer(),
		Mailer:                        m.Mailer.ValueStringPointer(),
		MailFrom:                      m.MailFrom.ValueStringPointer(),
		Managers:                      m.Managers.ValueStringPointer(),
		MaxArraySize:                  m.MaxArraySize.ValueInt32Pointer(),
		MaxConcurrentProvision:        m.MaxConcurrentProvision.ValueInt32Pointer(),
		MaxGroupRun:                   m.MaxGroupRun.ValueInt32Pointer(),
		MaxGroupRunSoft:               m.MaxGroupRunSoft.ValueInt32Pointer(),
		MaxJobSequenceId:              m.MaxJobSequenceId.ValueInt64Pointer(),
		MaxRunning:                    m.MaxRunning.ValueInt32Pointer(),
		MaxUserRun:                    m.MaxUserRun.ValueInt32Pointer(),
		MaxUserRunSoft:                m.MaxUserRunSoft.ValueInt32Pointer(),
		NodeFailRequeue:               m.NodeFailRequeue.ValueInt32Pointer(),
		NodeGroupEnable:               m.NodeGroupEnable.ValueBoolPointer(),
		NodeGroupKey:                  m.NodeGroupKey.ValueStringPointer(),
		Operators:                     m.Operators.ValueStringPointer(),
		PbsLicenseInfo:                m.PbsLicenseInfo.ValueStringPointer(),
		PbsLicenseLingerTime:          m.PbsLicenseLingerTime.ValueInt32Pointer(),
		PbsLicenseMax:                 m.PbsLicenseMax.ValueInt32Pointer(),
		PbsLicenseMin:                 m.PbsLicenseMin.ValueInt32Pointer(),
		PowerProvisioning:             m.PowerProvisioning.ValueBoolPointer(),
		PythonGcMinInterval:           m.PythonGcMinInterval.ValueInt32Pointer(),
		PythonRestartMaxPbsServers:    m.PythonRestartMaxPbsServers.ValueInt32Pointer(),
		PythonRestartMaxObjects:       m.PythonRestartMaxObjects.ValueInt32Pointer(),
		PythonRestartMinInterval:      m.PythonRestartMinInterval.ValueStringPointer(),
		QueryOtherJobs:                m.QueryOtherJobs.ValueBoolPointer(),
		QueuedJobsThreshold:           m.QueuedJobsThreshold.ValueStringPointer(),
		QueuedJobsThresholdRes:        m.QueuedJobsThresholdRes.ValueStringPointer(),
		ReserveRetryInit:              m.ReserveRetryInit.ValueInt32Pointer(),
		ReserveRetryTime:              m.ReserveRetryTime.ValueInt32Pointer(),
		RestrictResToReleaseOnSuspend: m.RestrictResToReleaseOnSuspend.ValueStringPointer(),
		ResvEnable:                    m.ResvEnable.ValueBoolPointer(),
		ResvPostProcessingTime:        m.ResvPostProcessingTime.ValueStringPointer(),
		RppHighwater:                  m.RppHighwater.ValueInt32Pointer(),
		RppMaxPktCheck:                m.RppMaxPktCheck.ValueInt32Pointer(),
		RppRetry:                      m.RppRetry.ValueInt32Pointer(),
		SchedulerIteration:            m.SchedulerIteration.ValueInt32Pointer(),
		WebapiAuthIssuers:             m.WebapiAuthIssuers.ValueStringPointer(),
		WebapiEnable:                  m.WebapiEnable.ValueBoolPointer(),
		WebapiOidcClientid:            m.WebapiOidcClientid.ValueStringPointer(),
		WebapiOidcProviderUrl:         m.WebapiOidcProviderUrl.ValueStringPointer(),
	}

	server.DefaultChunk = make(map[string]string)
	for k, v := range m.DefaultChunk {
		server.DefaultChunk[k] = v.ValueString()
	}
	server.ResourcesAvailable = make(map[string]string)
	for k, v := range m.ResourcesAvailable {
		server.ResourcesAvailable[k] = v.ValueString()
	}
	server.ResourcesDefault = make(map[string]string)
	for k, v := range m.ResourcesDefault {
		server.ResourcesDefault[k] = v.ValueString()
	}
	server.ResourcesMax = make(map[string]string)
	for k, v := range m.ResourcesMax {
		server.ResourcesMax[k] = v.ValueString()
	}

	// Convert limit attribute maps from Terraform types to Go maps
	if len(m.MaxGroupRes) > 0 {
		server.MaxGroupRes = make(map[string]string)
		for k, v := range m.MaxGroupRes {
			server.MaxGroupRes[k] = v.ValueString()
		}
	}
	if len(m.MaxGroupResSoft) > 0 {
		server.MaxGroupResSoft = make(map[string]string)
		for k, v := range m.MaxGroupResSoft {
			server.MaxGroupResSoft[k] = v.ValueString()
		}
	}
	if len(m.MaxQueued) > 0 {
		server.MaxQueued = make(map[string]string)
		for k, v := range m.MaxQueued {
			server.MaxQueued[k] = v.ValueString()
		}
	}
	if len(m.MaxQueuedRes) > 0 {
		server.MaxQueuedRes = make(map[string]string)
		for k, v := range m.MaxQueuedRes {
			server.MaxQueuedRes[k] = v.ValueString()
		}
	}
	if len(m.MaxRun) > 0 {
		server.MaxRun = make(map[string]string)
		for k, v := range m.MaxRun {
			server.MaxRun[k] = v.ValueString()
		}
	}
	if len(m.MaxRunRes) > 0 {
		server.MaxRunRes = make(map[string]string)
		for k, v := range m.MaxRunRes {
			server.MaxRunRes[k] = v.ValueString()
		}
	}
	if len(m.MaxRunResSoft) > 0 {
		server.MaxRunResSoft = make(map[string]string)
		for k, v := range m.MaxRunResSoft {
			server.MaxRunResSoft[k] = v.ValueString()
		}
	}
	if len(m.MaxRunSoft) > 0 {
		server.MaxRunSoft = make(map[string]string)
		for k, v := range m.MaxRunSoft {
			server.MaxRunSoft[k] = v.ValueString()
		}
	}
	if len(m.MaxUserRes) > 0 {
		server.MaxUserRes = make(map[string]string)
		for k, v := range m.MaxUserRes {
			server.MaxUserRes[k] = v.ValueString()
		}
	}
	if len(m.MaxUserResSoft) > 0 {
		server.MaxUserResSoft = make(map[string]string)
		for k, v := range m.MaxUserResSoft {
			server.MaxUserResSoft[k] = v.ValueString()
		}
	}

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
	model.AclResvGroupEnable = types.BoolPointerValue(server.AclResvGroupEnable)
	model.AclResvGroups = types.StringPointerValue(server.AclResvGroups)
	model.AclResvHostEnable = types.BoolPointerValue(server.AclResvHostEnable)
	model.AclResvHosts = types.StringPointerValue(server.AclResvHosts)
	model.AclResvUserEnable = types.BoolPointerValue(server.AclResvUserEnable)
	model.AclResvUsers = types.StringPointerValue(server.AclResvUsers)
	model.AclRoots = types.StringPointerValue(server.AclRoots)
	model.AclUserEnable = types.BoolPointerValue(server.AclUserEnable)
	model.AclUsers = types.StringPointerValue(server.AclUsers)
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
	model.PythonRestartMaxPbsServers = types.Int32PointerValue(server.PythonRestartMaxPbsServers)
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
		elements := make(map[string]types.String, 0)
		for k, v := range server.DefaultChunk {
			elements[k] = types.StringValue(v)
		}
		model.DefaultChunk = elements
	}
	if server.ResourcesAvailable != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range server.ResourcesAvailable {
			elements[k] = types.StringValue(v)
		}
		model.ResourcesAvailable = elements
	}
	if server.ResourcesDefault != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range server.ResourcesDefault {
			elements[k] = types.StringValue(v)
		}
		model.ResourcesDefault = elements
	}
	if server.ResourcesMax != nil {
		elements := make(map[string]types.String, 0)
		for k, v := range server.ResourcesMax {
			elements[k] = types.StringValue(v)
		}
		model.ResourcesMax = elements
	}

	return model
}
