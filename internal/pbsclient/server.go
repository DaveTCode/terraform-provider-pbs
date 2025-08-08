package pbsclient

import (
	"fmt"
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

	fields := []struct {
		attribute string
		newAttr   any
	}{
		{"acl_host_enable", newServer.AclHostEnable},
		{"acl_host_moms_enable", newServer.AclHostMomsEnable},
		{"acl_hosts", newServer.AclHosts},
		{"acl_resv_group_enable", newServer.AclResvGroupEnable},
		{"acl_resv_groups", newServer.AclResvGroups},
		{"acl_resv_host_enable", newServer.AclResvHostEnable},
		{"acl_resv_hosts", newServer.AclResvHosts},
		{"acl_resv_user_enable", newServer.AclResvUserEnable},
		{"acl_resv_users", newServer.AclResvUsers},
		{"acl_roots", newServer.AclRoots},
		{"acl_user_enable", newServer.AclUserEnable},
		{"acl_users", newServer.AclUsers},
		{"backfill_depth", newServer.BackfillDepth},
		{"comment", newServer.Comment},
		{"default_chunk", newServer.DefaultChunk},
		{"default_qdel_arguments", newServer.DefaultQdelArguments},
		{"default_qsub_arguments", newServer.DefaultQsubArguments},
		{"default_queue", newServer.DefaultQueue},
		{"eligible_time_enable", newServer.EligibleTimeEnable},
		{"elim_on_subjobs", newServer.ElimOnSubjobs},
		{"flatuid", newServer.Flatuid},
		{"job_history_duration", newServer.JobHistoryDuration},
		{"job_history_enable", newServer.JobHistoryEnable},
		{"job_requeue_timeout", newServer.JobRequeueTimeout},
		{"job_sort_formula", newServer.JobSortFormula},
		{"jobscript_max_size", newServer.JobscriptMaxSize},
		{"log_events", newServer.LogEvents},
		{"mailer", newServer.Mailer},
		{"mail_from", newServer.MailFrom},
		{"managers", newServer.Managers},
		{"max_array_size", newServer.MaxArraySize},
		{"max_concurrent_provision", newServer.MaxConcurrentProvision},
		{"max_group_res", newServer.MaxGroupRes},
		{"max_group_res_soft", newServer.MaxGroupResSoft},
		{"max_group_run", newServer.MaxGroupRun},
		{"max_group_run_soft", newServer.MaxGroupRunSoft},
		{"max_job_sequence_id", newServer.MaxJobSequenceId},
		{"max_queued", newServer.MaxQueued},
		{"max_queued_res", newServer.MaxQueuedRes},
		{"max_run", newServer.MaxRun},
		{"max_run_res", newServer.MaxRunRes},
		{"max_run_res_soft", newServer.MaxRunResSoft},
		{"max_run_soft", newServer.MaxRunSoft},
		{"max_running", newServer.MaxRunning},
		{"max_user_res", newServer.MaxUserRes},
		{"max_user_res_soft", newServer.MaxUserResSoft},
		{"max_user_run", newServer.MaxUserRun},
		{"max_user_run_soft", newServer.MaxUserRunSoft},
		{"node_fail_requeue", newServer.NodeFailRequeue},
		{"node_group_enable", newServer.NodeGroupEnable},
		{"node_group_key", newServer.NodeGroupKey},
		{"operators", newServer.Operators},
		{"pbs_license_info", newServer.PbsLicenseInfo},
		{"pbs_license_linger_time", newServer.PbsLicenseLingerTime},
		{"pbs_license_max", newServer.PbsLicenseMax},
		{"pbs_license_min", newServer.PbsLicenseMin},
		{"power_provisioning", newServer.PowerProvisioning},
		{"python_gc_min_interval", newServer.PythonGcMinInterval},
		{"python_restart_max_pbs_servers", newServer.PythonRestartMaxPbsServers},
		{"python_restart_max_objects", newServer.PythonRestartMaxObjects},
		{"python_restart_min_interval", newServer.PythonRestartMinInterval},
		{"query_other_jobs", newServer.QueryOtherJobs},
		{"queued_jobs_threshold", newServer.QueuedJobsThreshold},
		{"queued_jobs_threshold_res", newServer.QueuedJobsThresholdRes},
		{"reserve_retry_init", newServer.ReserveRetryInit},
		{"reserve_retry_time", newServer.ReserveRetryTime},
		{"resources_available", newServer.ResourcesAvailable},
		{"resources_default", newServer.ResourcesDefault},
		{"resources_max", newServer.ResourcesMax},
		{"restrict_res_to_release_on_suspend", newServer.RestrictResToReleaseOnSuspend},
		{"resv_enable", newServer.ResvEnable},
		{"resv_post_processing_time", newServer.ResvPostProcessingTime},
		{"rpp_highwater", newServer.RppHighwater},
		{"rpp_max_pkt_check", newServer.RppMaxPktCheck},
		{"rpp_retry", newServer.RppRetry},
		{"scheduler_iteration", newServer.SchedulerIteration},
		{"webapi_auth_issuers", newServer.WebapiAuthIssuers},
		{"webapi_enable", newServer.WebapiEnable},
		{"webapi_oidc_clientid", newServer.WebapiOidcClientid},
		{"webapi_oidc_provider_url", newServer.WebapiOidcProviderUrl},
	}
	for _, v := range fields {
		c, err := generateCreateCommands(v.newAttr, "server", newServer.Name, v.attribute)
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
	fields := []struct {
		attribute string
		oldAttr   any
		newAttr   any
	}{
		{"acl_host_enable", oldServer.AclHostEnable, newServer.AclHostEnable},
		{"acl_host_moms_enable", oldServer.AclHostMomsEnable, newServer.AclHostMomsEnable},
		{"acl_hosts", oldServer.AclHosts, newServer.AclHosts},
		{"acl_resv_group_enable", oldServer.AclResvGroupEnable, newServer.AclResvGroupEnable},
		{"acl_resv_groups", oldServer.AclResvGroups, newServer.AclResvGroups},
		{"acl_resv_host_enable", oldServer.AclResvHostEnable, newServer.AclResvHostEnable},
		{"acl_resv_hosts", oldServer.AclResvHosts, newServer.AclResvHosts},
		{"acl_resv_user_enable", oldServer.AclResvUserEnable, newServer.AclResvUserEnable},
		{"acl_resv_users", oldServer.AclResvUsers, newServer.AclResvUsers},
		{"acl_roots", oldServer.AclRoots, newServer.AclRoots},
		{"acl_user_enable", oldServer.AclUserEnable, newServer.AclUserEnable},
		{"acl_users", oldServer.AclUsers, newServer.AclUsers},
		{"backfill_depth", oldServer.BackfillDepth, newServer.BackfillDepth},
		{"comment", oldServer.Comment, newServer.Comment},
		{"default_chunk", oldServer.DefaultChunk, newServer.DefaultChunk},
		{"default_qdel_arguments", oldServer.DefaultQdelArguments, newServer.DefaultQdelArguments},
		{"default_qsub_arguments", oldServer.DefaultQsubArguments, newServer.DefaultQsubArguments},
		{"default_queue", oldServer.DefaultQueue, newServer.DefaultQueue},
		{"eligible_time_enable", oldServer.EligibleTimeEnable, newServer.EligibleTimeEnable},
		{"elim_on_subjobs", oldServer.ElimOnSubjobs, newServer.ElimOnSubjobs},
		{"flatuid", oldServer.Flatuid, newServer.Flatuid},
		{"job_history_duration", oldServer.JobHistoryDuration, newServer.JobHistoryDuration},
		{"job_history_enable", oldServer.JobHistoryEnable, newServer.JobHistoryEnable},
		{"job_requeue_timeout", oldServer.JobRequeueTimeout, newServer.JobRequeueTimeout},
		{"job_sort_formula", oldServer.JobSortFormula, newServer.JobSortFormula},
		{"jobscript_max_size", oldServer.JobscriptMaxSize, newServer.JobscriptMaxSize},
		{"log_events", oldServer.LogEvents, newServer.LogEvents},
		{"mailer", oldServer.Mailer, newServer.Mailer},
		{"mail_from", oldServer.MailFrom, newServer.MailFrom},
		{"managers", oldServer.Managers, newServer.Managers},
		{"max_array_size", oldServer.MaxArraySize, newServer.MaxArraySize},
		{"max_concurrent_provision", oldServer.MaxConcurrentProvision, newServer.MaxConcurrentProvision},
		{"max_group_res", oldServer.MaxGroupRes, newServer.MaxGroupRes},
		{"max_group_res_soft", oldServer.MaxGroupResSoft, newServer.MaxGroupResSoft},
		{"max_group_run", oldServer.MaxGroupRun, newServer.MaxGroupRun},
		{"max_group_run_soft", oldServer.MaxGroupRunSoft, newServer.MaxGroupRunSoft},
		{"max_job_sequence_id", oldServer.MaxJobSequenceId, newServer.MaxJobSequenceId},
		{"max_queued", oldServer.MaxQueued, newServer.MaxQueued},
		{"max_queued_res", oldServer.MaxQueuedRes, newServer.MaxQueuedRes},
		{"max_run", oldServer.MaxRun, newServer.MaxRun},
		{"max_run_res", oldServer.MaxRunRes, newServer.MaxRunRes},
		{"max_run_res_soft", oldServer.MaxRunResSoft, newServer.MaxRunResSoft},
		{"max_run_soft", oldServer.MaxRunSoft, newServer.MaxRunSoft},
		{"max_running", oldServer.MaxRunning, newServer.MaxRunning},
		{"max_user_res", oldServer.MaxUserRes, newServer.MaxUserRes},
		{"max_user_res_soft", oldServer.MaxUserResSoft, newServer.MaxUserResSoft},
		{"max_user_run", oldServer.MaxUserRun, newServer.MaxUserRun},
		{"max_user_run_soft", oldServer.MaxUserRunSoft, newServer.MaxUserRunSoft},
		{"node_fail_requeue", oldServer.NodeFailRequeue, newServer.NodeFailRequeue},
		{"node_group_enable", oldServer.NodeGroupEnable, newServer.NodeGroupEnable},
		{"node_group_key", oldServer.NodeGroupKey, newServer.NodeGroupKey},
		{"operators", oldServer.Operators, newServer.Operators},
		{"pbs_license_info", oldServer.PbsLicenseInfo, newServer.PbsLicenseInfo},
		{"pbs_license_linger_time", oldServer.PbsLicenseLingerTime, newServer.PbsLicenseLingerTime},
		{"pbs_license_max", oldServer.PbsLicenseMax, newServer.PbsLicenseMax},
		{"pbs_license_min", oldServer.PbsLicenseMin, newServer.PbsLicenseMin},
		{"power_provisioning", oldServer.PowerProvisioning, newServer.PowerProvisioning},
		{"python_gc_min_interval", oldServer.PythonGcMinInterval, newServer.PythonGcMinInterval},
		{"python_restart_max_pbs_servers", oldServer.PythonRestartMaxPbsServers, newServer.PythonRestartMaxPbsServers},
		{"python_restart_max_objects", oldServer.PythonRestartMaxObjects, newServer.PythonRestartMaxObjects},
		{"python_restart_min_interval", oldServer.PythonRestartMinInterval, newServer.PythonRestartMinInterval},
		{"query_other_jobs", oldServer.QueryOtherJobs, newServer.QueryOtherJobs},
		{"queued_jobs_threshold", oldServer.QueuedJobsThreshold, newServer.QueuedJobsThreshold},
		{"queued_jobs_threshold_res", oldServer.QueuedJobsThresholdRes, newServer.QueuedJobsThresholdRes},
		{"reserve_retry_init", oldServer.ReserveRetryInit, newServer.ReserveRetryInit},
		{"reserve_retry_time", oldServer.ReserveRetryTime, newServer.ReserveRetryTime},
		{"resources_available", oldServer.ResourcesAvailable, newServer.ResourcesAvailable},
		{"resources_default", oldServer.ResourcesDefault, newServer.ResourcesDefault},
		{"resources_max", oldServer.ResourcesMax, newServer.ResourcesMax},
		{"restrict_res_to_release_on_suspend", oldServer.RestrictResToReleaseOnSuspend, newServer.RestrictResToReleaseOnSuspend},
		{"resv_enable", oldServer.ResvEnable, newServer.ResvEnable},
		{"resv_post_processing_time", oldServer.ResvPostProcessingTime, newServer.ResvPostProcessingTime},
		{"rpp_highwater", oldServer.RppHighwater, newServer.RppHighwater},
		{"rpp_max_pkt_check", oldServer.RppMaxPktCheck, newServer.RppMaxPktCheck},
		{"rpp_retry", oldServer.RppRetry, newServer.RppRetry},
		{"scheduler_iteration", oldServer.SchedulerIteration, newServer.SchedulerIteration},
		{"webapi_auth_issuers", oldServer.WebapiAuthIssuers, newServer.WebapiAuthIssuers},
		{"webapi_enable", oldServer.WebapiEnable, newServer.WebapiEnable},
		{"webapi_oidc_clientid", oldServer.WebapiOidcClientid, newServer.WebapiOidcClientid},
		{"webapi_oidc_provider_url", oldServer.WebapiOidcProviderUrl, newServer.WebapiOidcProviderUrl},
	}
	for _, v := range fields {
		newCommands, err := generateUpdateAttributeCommand(v.oldAttr, v.newAttr, "server", newServer.Name, v.attribute)
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
