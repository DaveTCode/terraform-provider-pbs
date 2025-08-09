package pbsclient

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type PbsServer struct {
	AclHostEnable                 *bool
	AclHostMomsEnable             *bool
	AclHosts                      *string
	AclResvGroupEnable            *bool
	AclResvGroups                 *string
	AclResvHostEnable             *bool
	AclResvHosts                  *string
	AclResvUserEnable             *bool
	AclResvUsers                  *string
	AclRoots                      *string
	AclUserEnable                 *bool
	AclUsers                      *string
	BackfillDepth                 *int32
	Comment                       *string
	DefaultChunk                  map[string]string
	DefaultQdelArguments          *string
	DefaultQsubArguments          *string
	DefaultQueue                  *string
	EligibleTimeEnable            *bool
	ElimOnSubjobs                 *bool
	Flatuid                       *bool
	JobHistoryDuration            *string
	JobHistoryEnable              *bool
	JobRequeueTimeout             *string
	JobSortFormula                *string
	JobscriptMaxSize              *string
	LogEvents                     *int32
	Mailer                        *string
	MailFrom                      *string
	Managers                      *string
	MaxArraySize                  *int32
	MaxConcurrentProvision        *int32
	MaxGroupRes                   map[string]string
	MaxGroupResSoft               map[string]string
	MaxGroupRun                   *int32
	MaxGroupRunSoft               *int32
	MaxJobSequenceId              *int64
	MaxQueued                     map[string]string
	MaxQueuedRes                  map[string]string
	MaxRun                        map[string]string
	MaxRunRes                     map[string]string
	MaxRunResSoft                 map[string]string
	MaxRunSoft                    map[string]string
	MaxRunning                    *int32
	MaxUserRes                    map[string]string
	MaxUserResSoft                map[string]string
	MaxUserRun                    *int32
	MaxUserRunSoft                *int32
	Name                          string
	NodeFailRequeue               *int32
	NodeGroupEnable               *bool
	NodeGroupKey                  *string
	Operators                     *string
	PbsLicenseInfo                *string
	PbsLicenseLingerTime          *int32
	PbsLicenseMax                 *int32
	PbsLicenseMin                 *int32
	PowerProvisioning             *bool
	PythonGcMinInterval           *int32
	PythonRestartMaxPbsServers    *int32
	PythonRestartMaxObjects       *int32
	PythonRestartMinInterval      *string
	QueryOtherJobs                *bool
	QueuedJobsThreshold           *string
	QueuedJobsThresholdRes        *string
	ReserveRetryInit              *int32
	ReserveRetryTime              *int32
	ResourcesAvailable            map[string]string
	ResourcesDefault              map[string]string
	ResourcesMax                  map[string]string
	RestrictResToReleaseOnSuspend *string
	ResvEnable                    *bool
	ResvPostProcessingTime        *string
	RppHighwater                  *int32
	RppMaxPktCheck                *int32
	RppRetry                      *int32
	SchedulerIteration            *int32
	WebapiAuthIssuers             *string
	WebapiEnable                  *bool
	WebapiOidcClientid            *string
	WebapiOidcProviderUrl         *string
}

// serverFieldDefinition represents a server field with its attribute name and execution order
type serverFieldDefinition struct {
	attribute string
	order     int                        // Lower numbers execute first
	getValue  func(server PbsServer) any // Function to extract the value from a PbsServer
}

// getServerFieldDefinitions returns the ordered list of server field definitions
// This ensures consistent ordering across create and update operations
func getServerFieldDefinitions() []serverFieldDefinition {
	return []serverFieldDefinition{
		{"acl_host_enable", 10, func(q PbsServer) any { return q.AclHostEnable }},
		{"acl_host_moms_enable", 10, func(q PbsServer) any { return q.AclHostMomsEnable }},
		{"acl_hosts", 10, func(q PbsServer) any { return q.AclHosts }},
		{"acl_resv_group_enable", 10, func(q PbsServer) any { return q.AclResvGroupEnable }},
		{"acl_resv_groups", 10, func(q PbsServer) any { return q.AclResvGroups }},
		{"acl_resv_host_enable", 10, func(q PbsServer) any { return q.AclResvHostEnable }},
		{"acl_resv_hosts", 10, func(q PbsServer) any { return q.AclResvHosts }},
		{"acl_resv_user_enable", 10, func(q PbsServer) any { return q.AclResvUserEnable }},
		{"acl_resv_users", 10, func(q PbsServer) any { return q.AclResvUsers }},
		{"acl_roots", 10, func(q PbsServer) any { return q.AclRoots }},
		{"acl_user_enable", 10, func(q PbsServer) any { return q.AclUserEnable }},
		{"acl_users", 10, func(q PbsServer) any { return q.AclUsers }},
		{"backfill_depth", 10, func(q PbsServer) any { return q.BackfillDepth }},
		{"comment", 10, func(q PbsServer) any { return q.Comment }},
		{"default_chunk", 10, func(q PbsServer) any { return q.DefaultChunk }},
		{"default_qdel_arguments", 10, func(q PbsServer) any { return q.DefaultQdelArguments }},
		{"default_qsub_arguments", 10, func(q PbsServer) any { return q.DefaultQsubArguments }},
		{"default_queue", 10, func(q PbsServer) any { return q.DefaultQueue }},
		{"eligible_time_enable", 10, func(q PbsServer) any { return q.EligibleTimeEnable }},
		{"elim_on_subjobs", 10, func(q PbsServer) any { return q.ElimOnSubjobs }},
		{"flatuid", 10, func(q PbsServer) any { return q.Flatuid }},
		{"job_history_duration", 10, func(q PbsServer) any { return q.JobHistoryDuration }},
		{"job_history_enable", 10, func(q PbsServer) any { return q.JobHistoryEnable }},
		{"job_requeue_timeout", 10, func(q PbsServer) any { return q.JobRequeueTimeout }},
		{"job_sort_formula", 10, func(q PbsServer) any { return q.JobSortFormula }},
		{"jobscript_max_size", 10, func(q PbsServer) any { return q.JobscriptMaxSize }},
		{"log_events", 10, func(q PbsServer) any { return q.LogEvents }},
		{"mailer", 10, func(q PbsServer) any { return q.Mailer }},
		{"mail_from", 10, func(q PbsServer) any { return q.MailFrom }},
		{"managers", 10, func(q PbsServer) any { return q.Managers }},
		{"max_array_size", 10, func(q PbsServer) any { return q.MaxArraySize }},
		{"max_concurrent_provision", 10, func(q PbsServer) any { return q.MaxConcurrentProvision }},
		{"max_group_res", 10, func(q PbsServer) any { return q.MaxGroupRes }},
		{"max_group_res_soft", 10, func(q PbsServer) any { return q.MaxGroupResSoft }},
		{"max_group_run", 10, func(q PbsServer) any { return q.MaxGroupRun }},
		{"max_group_run_soft", 10, func(q PbsServer) any { return q.MaxGroupRunSoft }},
		{"max_job_sequence_id", 10, func(q PbsServer) any { return q.MaxJobSequenceId }},
		{"max_queued", 10, func(q PbsServer) any { return q.MaxQueued }},
		{"max_queued_res", 10, func(q PbsServer) any { return q.MaxQueuedRes }},
		{"max_run", 10, func(q PbsServer) any { return q.MaxRun }},
		{"max_run_res", 10, func(q PbsServer) any { return q.MaxRunRes }},
		{"max_run_res_soft", 10, func(q PbsServer) any { return q.MaxRunResSoft }},
		{"max_run_soft", 10, func(q PbsServer) any { return q.MaxRunSoft }},
		{"max_running", 10, func(q PbsServer) any { return q.MaxRunning }},
		{"max_user_res", 10, func(q PbsServer) any { return q.MaxUserRes }},
		{"max_user_res_soft", 10, func(q PbsServer) any { return q.MaxUserResSoft }},
		{"max_user_run", 10, func(q PbsServer) any { return q.MaxUserRun }},
		{"max_user_run_soft", 10, func(q PbsServer) any { return q.MaxUserRunSoft }},
		{"node_fail_requeue", 10, func(q PbsServer) any { return q.NodeFailRequeue }},
		{"node_group_enable", 10, func(q PbsServer) any { return q.NodeGroupEnable }},
		{"node_group_key", 10, func(q PbsServer) any { return q.NodeGroupKey }},
		{"operators", 10, func(q PbsServer) any { return q.Operators }},
		{"pbs_license_info", 10, func(q PbsServer) any { return q.PbsLicenseInfo }},
		{"pbs_license_linger_time", 10, func(q PbsServer) any { return q.PbsLicenseLingerTime }},
		{"pbs_license_max", 10, func(q PbsServer) any { return q.PbsLicenseMax }},
		{"pbs_license_min", 10, func(q PbsServer) any { return q.PbsLicenseMin }},
		{"power_provisioning", 10, func(q PbsServer) any { return q.PowerProvisioning }},
		{"python_gc_min_interval", 10, func(q PbsServer) any { return q.PythonGcMinInterval }},
		{"python_restart_max_pbs_servers", 10, func(q PbsServer) any { return q.PythonRestartMaxPbsServers }},
		{"python_restart_max_objects", 10, func(q PbsServer) any { return q.PythonRestartMaxObjects }},
		{"python_restart_min_interval", 10, func(q PbsServer) any { return q.PythonRestartMinInterval }},
		{"query_other_jobs", 10, func(q PbsServer) any { return q.QueryOtherJobs }},
		{"queued_jobs_threshold", 10, func(q PbsServer) any { return q.QueuedJobsThreshold }},
		{"queued_jobs_threshold_res", 10, func(q PbsServer) any { return q.QueuedJobsThresholdRes }},
		{"reserve_retry_init", 10, func(q PbsServer) any { return q.ReserveRetryInit }},
		{"reserve_retry_time", 10, func(q PbsServer) any { return q.ReserveRetryTime }},
		{"resources_available", 10, func(q PbsServer) any { return q.ResourcesAvailable }},
		{"resources_default", 10, func(q PbsServer) any { return q.ResourcesDefault }},
		{"resources_max", 10, func(q PbsServer) any { return q.ResourcesMax }},
		{"restrict_res_to_release_on_suspend", 10, func(q PbsServer) any { return q.RestrictResToReleaseOnSuspend }},
		{"resv_enable", 10, func(q PbsServer) any { return q.ResvEnable }},
		{"resv_post_processing_time", 10, func(q PbsServer) any { return q.ResvPostProcessingTime }},
		{"rpp_highwater", 10, func(q PbsServer) any { return q.RppHighwater }},
		{"rpp_max_pkt_check", 10, func(q PbsServer) any { return q.RppMaxPktCheck }},
		{"rpp_retry", 10, func(q PbsServer) any { return q.RppRetry }},
		{"scheduler_iteration", 10, func(q PbsServer) any { return q.SchedulerIteration }},
		{"webapi_auth_issuers", 10, func(q PbsServer) any { return q.WebapiAuthIssuers }},
		{"webapi_enable", 10, func(q PbsServer) any { return q.WebapiEnable }},
		{"webapi_oidc_clientid", 10, func(q PbsServer) any { return q.WebapiOidcClientid }},
		{"webapi_oidc_provider_url", 10, func(q PbsServer) any { return q.WebapiOidcProviderUrl }},
	}
}

func parseServerOutput(output []byte) ([]PbsServer, error) {
	parsedOutput := parseGenericQmgrOutput(string(output))
	var servers []PbsServer

	for _, r := range parsedOutput {
		if r.objType == "Server" {
			current := PbsServer{
				Name: r.name,
			}

			for k, v := range r.attributes {
				if s, ok := v.(string); ok {
					switch strings.ToLower(k) {
					case "acl_host_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert acl_host_enable value to bool %s", err.Error())
						}
						current.AclHostEnable = &boolValue
					case "acl_host_moms_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert acl_host_moms_enable value to bool %s", err.Error())
						}
						current.AclHostMomsEnable = &boolValue
					case "acl_hosts":
						current.AclHosts = &s
					case "acl_resv_group_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert acl_resv_group_enable value to bool %s", err.Error())
						}
						current.AclResvGroupEnable = &boolValue
					case "acl_resv_groups":
						current.AclResvGroups = &s
					case "acl_resv_host_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert acl_resv_host_enable value to bool %s", err.Error())
						}
						current.AclResvHostEnable = &boolValue
					case "acl_resv_hosts":
						current.AclResvHosts = &s
					case "acl_resv_user_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert acl_resv_user_enable value to bool %s", err.Error())
						}
						current.AclResvUserEnable = &boolValue
					case "acl_resv_users":
						current.AclResvUsers = &s
					case "acl_roots":
						current.AclRoots = &s
					case "acl_user_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert acl_user_enable value to bool %s", err.Error())
						}
						current.AclUserEnable = &boolValue
					case "acl_users":
						current.AclUsers = &s
					case "backfill_depth":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert backfill_depth value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.BackfillDepth = &int32Value
					case "comment":
						current.Comment = &s
					case "default_qdel_arguments":
						current.DefaultQdelArguments = &s
					case "default_qsub_arguments":
						current.DefaultQsubArguments = &s
					case "default_queue":
						current.DefaultQueue = &s
					case "eligible_time_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert eligible_time_enable value to bool %s", err.Error())
						}
						current.EligibleTimeEnable = &boolValue
					case "elim_on_subjobs":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert elim_on_subjobs value to bool %s", err.Error())
						}
						current.ElimOnSubjobs = &boolValue
					case "flatuid":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert flatuid value to bool %s", err.Error())
						}
						current.Flatuid = &boolValue
					case "job_history_duration":
						current.JobHistoryDuration = &s
					case "job_history_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert job_history_enable value to bool %s", err.Error())
						}
						current.JobHistoryEnable = &boolValue
					case "job_requeue_timeout":
						current.JobRequeueTimeout = &s
					case "job_sort_formula":
						current.JobSortFormula = &s
					case "jobscript_max_size":
						current.JobscriptMaxSize = &s
					case "log_events":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_concurrent_provision value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.LogEvents = &int32Value
					case "mailer":
						current.Mailer = &s
					case "mail_from":
						current.MailFrom = &s
					case "managers":
						current.Managers = &s
					case "max_array_size":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_array_size value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.MaxArraySize = &int32Value
					case "max_concurrent_provision":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_concurrent_provision value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.MaxConcurrentProvision = &int32Value
					case "max_group_res":
						// Skip individual parsing - will be handled by map parsing below
					case "max_group_res_soft":
						// Skip individual parsing - will be handled by map parsing below
					case "max_group_run":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_group_run value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.MaxGroupRun = &int32Value
					case "max_group_run_soft":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_group_run_soft value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.MaxGroupRunSoft = &int32Value
					case "max_job_sequence_id":
						intValue, err := strconv.ParseInt(s, 10, 64)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_job_sequence_id value to int64 %s", err.Error())
						}
						current.MaxJobSequenceId = &intValue
					case "max_queued":
						// Skip individual parsing - will be handled by map parsing below
					case "max_queued_res":
						// Skip individual parsing - will be handled by map parsing below
					case "max_run":
						// Skip individual parsing - will be handled by map parsing below
					case "max_run_res":
						// Skip individual parsing - will be handled by map parsing below
					case "max_run_res_soft":
						// Skip individual parsing - will be handled by map parsing below
					case "max_run_soft":
						// Skip individual parsing - will be handled by map parsing below
					case "max_running":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_running value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.MaxRunning = &int32Value
					case "max_user_res":
						// Skip individual parsing - will be handled by map parsing below
					case "max_user_res_soft":
						// Skip individual parsing - will be handled by map parsing below
					case "max_user_run":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_user_run value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.MaxUserRun = &int32Value
					case "max_user_run_soft":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_user_run_soft value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.MaxUserRunSoft = &int32Value
					case "node_fail_requeue":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert node_fail_requeue value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.NodeFailRequeue = &int32Value
					case "node_group_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert node_group_enable value to bool %s", err.Error())
						}
						current.NodeGroupEnable = &boolValue
					case "node_group_key":
						current.NodeGroupKey = &s
					case "operators":
						current.Operators = &s
					case "pbs_license_info":
						current.PbsLicenseInfo = &s
					case "pbs_license_linger_time":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert pbs_license_linger_time value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.PbsLicenseLingerTime = &int32Value
					case "pbs_license_max":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert pbs_license_max value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.PbsLicenseMax = &int32Value
					case "pbs_license_min":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert pbs_license_min value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.PbsLicenseMin = &int32Value
					case "power_provisioning":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert power_provisioning value to bool %s", err.Error())
						}
						current.PowerProvisioning = &boolValue
					case "python_gc_min_interval":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert python_gc_min_interval value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.PythonGcMinInterval = &int32Value
					case "python_restart_max_pbs_servers":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert python_restart_max_pbs_servers value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.PythonRestartMaxPbsServers = &int32Value
					case "python_restart_max_objects":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert python_restart_max_objects value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.PythonRestartMaxObjects = &int32Value
					case "python_restart_min_interval":
						current.PythonRestartMinInterval = &s
					case "query_other_jobs":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert query_other_jobs value to bool %s", err.Error())
						}
						current.QueryOtherJobs = &boolValue
					case "queued_jobs_threshold":
						current.QueuedJobsThreshold = &s
					case "queued_jobs_threshold_res":
						current.QueuedJobsThresholdRes = &s
					case "reserve_retry_init":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert reserve_retry_init value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.ReserveRetryInit = &int32Value
					case "reserve_retry_time":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert reserve_retry_time value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.ReserveRetryTime = &int32Value
					case "restrict_res_to_release_on_suspend":
						current.RestrictResToReleaseOnSuspend = &s
					case "resv_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert resv_enable value to bool %s", err.Error())
						}
						current.ResvEnable = &boolValue
					case "resv_post_processing_time":
						current.ResvPostProcessingTime = &s
					case "rpp_highwater":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert rpp_highwater value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.RppHighwater = &int32Value
					case "rpp_max_pkt_check":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert rpp_max_pkt_check value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.RppMaxPktCheck = &int32Value
					case "rpp_retry":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert rpp_retry value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.RppRetry = &int32Value
					case "scheduler_iteration":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert scheduler_iteration value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.SchedulerIteration = &int32Value
					case "webapi_auth_issuers":
						current.WebapiAuthIssuers = &s
					case "webapi_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert webapi_enable value to bool %s", err.Error())
						}
						current.WebapiEnable = &boolValue
					case "webapi_oidc_clientid":
						current.WebapiOidcClientid = &s
					case "webapi_oidc_provider_url":
						current.WebapiOidcProviderUrl = &s
					}
				} else if a, ok := v.(map[string]string); ok {
					switch strings.ToLower(k) {
					case "default_chunk":
						current.DefaultChunk = a
					case "resources_available":
						current.ResourcesAvailable = a
					case "resources_default":
						current.ResourcesDefault = a
					case "resources_max":
						current.ResourcesMax = a
					}
				}

				// Handle map attributes (like max_queued.ncpus)
				if a, ok := v.(map[string]string); ok {
					switch strings.ToLower(k) {
					case "max_group_res":
						current.MaxGroupRes = a
					case "max_group_res_soft":
						current.MaxGroupResSoft = a
					case "max_queued":
						current.MaxQueued = a
					case "max_queued_res":
						current.MaxQueuedRes = a
					case "max_run":
						current.MaxRun = a
					case "max_run_res":
						current.MaxRunRes = a
					case "max_run_res_soft":
						current.MaxRunResSoft = a
					case "max_run_soft":
						current.MaxRunSoft = a
					case "max_user_res":
						current.MaxUserRes = a
					case "max_user_res_soft":
						current.MaxUserResSoft = a
					}
				}
			}

			servers = append(servers, current)
		}
	}

	return servers, nil
}

func (c *PbsClient) GetPbsServer(name string) (PbsServer, error) {
	all, err := c.GetPbsServers()
	if err != nil {
		return PbsServer{}, err
	}

	for _, r := range all {
		if r.Name == name {
			return r, nil
		}
	}

	return PbsServer{}, nil
}

func (c *PbsClient) GetPbsServers() ([]PbsServer, error) {
	out, errOutput, err := c.runCommand("/opt/pbs/bin/qmgr -c 'list server @default'")
	if err != nil {
		return nil, fmt.Errorf("%s %s", err, errOutput)
	}

	return parseServerOutput(out)
}

func (c *PbsClient) CreatePbsServer(newServer PbsServer) (PbsServer, error) {
	var commands = []string{
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create server %s'", newServer.Name),
	}

	// Get field definitions and sort by order
	fieldDefs := getServerFieldDefinitions()
	sort.Slice(fieldDefs, func(i, j int) bool {
		return fieldDefs[i].order < fieldDefs[j].order
	})

	// Process fields in order
	for _, fieldDef := range fieldDefs {
		value := fieldDef.getValue(newServer)
		c, err := generateCreateCommands(value, "server", newServer.Name, fieldDef.attribute)
		if err != nil {
			return PbsServer{}, err
		}
		commands = append(commands, c...)
	}

	_, errOutput, err := c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return PbsServer{}, fmt.Errorf("%s %s %s", err, completeErrOutput, strings.Join(commands, ","))
	}

	return c.GetPbsServer(newServer.Name)
}

func (c *PbsClient) UpdatePbsServer(newServer PbsServer) (PbsServer, error) {
	oldServer, err := c.GetPbsServer(newServer.Name)
	if err != nil {
		return oldServer, err
	}

	var commands = []string{}

	// Get field definitions and sort by order
	fieldDefs := getServerFieldDefinitions()
	sort.Slice(fieldDefs, func(i, j int) bool {
		return fieldDefs[i].order < fieldDefs[j].order
	})

	// Process fields in order
	for _, fieldDef := range fieldDefs {
		oldValue := fieldDef.getValue(oldServer)
		newValue := fieldDef.getValue(newServer)
		newCommands, err := generateUpdateAttributeCommand(oldValue, newValue, "server", newServer.Name, fieldDef.attribute)
		if err != nil {
			return oldServer, err
		}
		commands = append(commands, newCommands...)
	}

	_, errOutput, err := c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return oldServer, fmt.Errorf("%s %s %s", err, completeErrOutput, strings.Join(commands, ","))
	}

	return c.GetPbsServer(oldServer.Name)
}

func (c *PbsClient) DeletePbsServer(name string) error {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete server %s'", name)
	_, errOutput, err := c.runCommand(cmd)
	if err != nil {
		return fmt.Errorf("%s %s", err, errOutput)
	}

	return nil
}
