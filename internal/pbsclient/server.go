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
	MaxGroupRes                   *string
	MaxGroupResSoft               *string
	MaxGroupRun                   *int32
	MaxGroupRunSoft               *int32
	MaxJobSequenceId              *int64
	MaxQueued                     *string
	MaxQueuedRes                  *string
	MaxRun                        *string
	MaxRunRes                     *string
	MaxRunResSoft                 *string
	MaxRunSoft                    *string
	MaxRunning                    *int32
	MaxUserRes                    *string
	MaxUserResSoft                *string
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
						current.MaxGroupRes = &s
					case "max_group_res_soft":
						current.MaxGroupResSoft = &s
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
						current.MaxQueued = &s
					case "max_queued_res":
						current.MaxQueuedRes = &s
					case "max_run":
						current.MaxRun = &s
					case "max_run_res":
						current.MaxRunRes = &s
					case "max_run_res_soft":
						current.MaxRunResSoft = &s
					case "max_run_soft":
						current.MaxRunSoft = &s
					case "max_running":
						intValue, err := strconv.ParseInt(s, 10, 32)
						if err != nil {
							return nil, fmt.Errorf("failed to convert max_running value to int32 %s", err.Error())
						}
						int32Value := int32(intValue)
						current.MaxRunning = &int32Value
					case "max_user_res":
						current.MaxUserRes = &s
					case "max_user_res_soft":
						current.MaxUserResSoft = &s
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

func (c *PbsClient) CreatePbsServer(new PbsServer) (PbsServer, error) {
	var commands = []string{
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create server %s'", new.Name),
	}

	fields := []struct {
		attribute string
		new       any
	}{
		{"acl_host_enable", new.AclHostEnable},
		{"acl_host_moms_enable", new.AclHostMomsEnable},
		{"acl_hosts", new.AclHosts},
		{"acl_resv_group_enable", new.AclResvGroupEnable},
		{"acl_resv_groups", new.AclResvGroups},
		{"acl_resv_host_enable", new.AclResvHostEnable},
		{"acl_resv_hosts", new.AclResvHosts},
		{"acl_resv_user_enable", new.AclResvUserEnable},
		{"acl_resv_users", new.AclResvUsers},
		{"acl_roots", new.AclRoots},
		{"acl_user_enable", new.AclUserEnable},
		{"acl_users", new.AclUsers},
		{"backfill_depth", new.BackfillDepth},
		{"comment", new.Comment},
		{"default_chunk", new.DefaultChunk},
		{"default_qdel_arguments", new.DefaultQdelArguments},
		{"default_qsub_arguments", new.DefaultQsubArguments},
		{"default_queue", new.DefaultQueue},
		{"eligible_time_enable", new.EligibleTimeEnable},
		{"elim_on_subjobs", new.ElimOnSubjobs},
		{"flatuid", new.Flatuid},
		{"job_history_duration", new.JobHistoryDuration},
		{"job_history_enable", new.JobHistoryEnable},
		{"job_requeue_timeout", new.JobRequeueTimeout},
		{"job_sort_formula", new.JobSortFormula},
		{"jobscript_max_size", new.JobscriptMaxSize},
		{"log_events", new.LogEvents},
		{"mailer", new.Mailer},
		{"mail_from", new.MailFrom},
		{"managers", new.Managers},
		{"max_array_size", new.MaxArraySize},
		{"max_concurrent_provision", new.MaxConcurrentProvision},
		{"max_group_res", new.MaxGroupRes},
		{"max_group_res_soft", new.MaxGroupResSoft},
		{"max_group_run", new.MaxGroupRun},
		{"max_group_run_soft", new.MaxGroupRunSoft},
		{"max_job_sequence_id", new.MaxJobSequenceId},
		{"max_queued", new.MaxQueued},
		{"max_queued_res", new.MaxQueuedRes},
		{"max_run", new.MaxRun},
		{"max_run_res", new.MaxRunRes},
		{"max_run_res_soft", new.MaxRunResSoft},
		{"max_run_soft", new.MaxRunSoft},
		{"max_running", new.MaxRunning},
		{"max_user_res", new.MaxUserRes},
		{"max_user_res_soft", new.MaxUserResSoft},
		{"max_user_run", new.MaxUserRun},
		{"max_user_run_soft", new.MaxUserRunSoft},
		{"node_fail_requeue", new.NodeFailRequeue},
		{"node_group_enable", new.NodeGroupEnable},
		{"node_group_key", new.NodeGroupKey},
		{"operators", new.Operators},
		{"pbs_license_info", new.PbsLicenseInfo},
		{"pbs_license_linger_time", new.PbsLicenseLingerTime},
		{"pbs_license_max", new.PbsLicenseMax},
		{"pbs_license_min", new.PbsLicenseMin},
		{"power_provisioning", new.PowerProvisioning},
		{"python_gc_min_interval", new.PythonGcMinInterval},
		{"python_restart_max_pbs_servers", new.PythonRestartMaxPbsServers},
		{"python_restart_max_objects", new.PythonRestartMaxObjects},
		{"python_restart_min_interval", new.PythonRestartMinInterval},
		{"query_other_jobs", new.QueryOtherJobs},
		{"queued_jobs_threshold", new.QueuedJobsThreshold},
		{"queued_jobs_threshold_res", new.QueuedJobsThresholdRes},
		{"reserve_retry_init", new.ReserveRetryInit},
		{"reserve_retry_time", new.ReserveRetryTime},
		{"resources_available", new.ResourcesAvailable},
		{"resources_default", new.ResourcesDefault},
		{"resources_max", new.ResourcesMax},
		{"restrict_res_to_release_on_suspend", new.RestrictResToReleaseOnSuspend},
		{"resv_enable", new.ResvEnable},
		{"resv_post_processing_time", new.ResvPostProcessingTime},
		{"rpp_highwater", new.RppHighwater},
		{"rpp_max_pkt_check", new.RppMaxPktCheck},
		{"rpp_retry", new.RppRetry},
		{"scheduler_iteration", new.SchedulerIteration},
		{"webapi_auth_issuers", new.WebapiAuthIssuers},
		{"webapi_enable", new.WebapiEnable},
		{"webapi_oidc_clientid", new.WebapiOidcClientid},
		{"webapi_oidc_provider_url", new.WebapiOidcProviderUrl},
	}
	for _, v := range fields {
		c, err := generateCreateCommands(v.new, "server", new.Name, v.attribute)
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

	return c.GetPbsServer(new.Name)
}

func (c *PbsClient) UpdatePbsServer(new PbsServer) (PbsServer, error) {
	old, err := c.GetPbsServer(new.Name)
	if err != nil {
		return old, err
	}

	var commands = []string{}
	fields := []struct {
		attribute string
		old       any
		new       any
	}{
		{"acl_host_enable", old.AclHostEnable, new.AclHostEnable},
		{"acl_host_moms_enable", old.AclHostMomsEnable, new.AclHostMomsEnable},
		{"acl_hosts", old.AclHosts, new.AclHosts},
		{"acl_resv_group_enable", old.AclResvGroupEnable, new.AclResvGroupEnable},
		{"acl_resv_groups", old.AclResvGroups, new.AclResvGroups},
		{"acl_resv_host_enable", old.AclResvHostEnable, new.AclResvHostEnable},
		{"acl_resv_hosts", old.AclResvHosts, new.AclResvHosts},
		{"acl_resv_user_enable", old.AclResvUserEnable, new.AclResvUserEnable},
		{"acl_resv_users", old.AclResvUsers, new.AclResvUsers},
		{"acl_roots", old.AclRoots, new.AclRoots},
		{"acl_user_enable", old.AclUserEnable, new.AclUserEnable},
		{"acl_users", old.AclUsers, new.AclUsers},
		{"backfill_depth", old.BackfillDepth, new.BackfillDepth},
		{"comment", old.Comment, new.Comment},
		{"default_chunk", old.DefaultChunk, new.DefaultChunk},
		{"default_qdel_arguments", old.DefaultQdelArguments, new.DefaultQdelArguments},
		{"default_qsub_arguments", old.DefaultQsubArguments, new.DefaultQsubArguments},
		{"default_queue", old.DefaultQueue, new.DefaultQueue},
		{"eligible_time_enable", old.EligibleTimeEnable, new.EligibleTimeEnable},
		{"elim_on_subjobs", old.ElimOnSubjobs, new.ElimOnSubjobs},
		{"flatuid", old.Flatuid, new.Flatuid},
		{"job_history_duration", old.JobHistoryDuration, new.JobHistoryDuration},
		{"job_history_enable", old.JobHistoryEnable, new.JobHistoryEnable},
		{"job_requeue_timeout", old.JobRequeueTimeout, new.JobRequeueTimeout},
		{"job_sort_formula", old.JobSortFormula, new.JobSortFormula},
		{"jobscript_max_size", old.JobscriptMaxSize, new.JobscriptMaxSize},
		{"log_events", old.LogEvents, new.LogEvents},
		{"mailer", old.Mailer, new.Mailer},
		{"mail_from", old.MailFrom, new.MailFrom},
		{"managers", old.Managers, new.Managers},
		{"max_array_size", old.MaxArraySize, new.MaxArraySize},
		{"max_concurrent_provision", old.MaxConcurrentProvision, new.MaxConcurrentProvision},
		{"max_group_res", old.MaxGroupRes, new.MaxGroupRes},
		{"max_group_res_soft", old.MaxGroupResSoft, new.MaxGroupResSoft},
		{"max_group_run", old.MaxGroupRun, new.MaxGroupRun},
		{"max_group_run_soft", old.MaxGroupRunSoft, new.MaxGroupRunSoft},
		{"max_job_sequence_id", old.MaxJobSequenceId, new.MaxJobSequenceId},
		{"max_queued", old.MaxQueued, new.MaxQueued},
		{"max_queued_res", old.MaxQueuedRes, new.MaxQueuedRes},
		{"max_run", old.MaxRun, new.MaxRun},
		{"max_run_res", old.MaxRunRes, new.MaxRunRes},
		{"max_run_res_soft", old.MaxRunResSoft, new.MaxRunResSoft},
		{"max_run_soft", old.MaxRunSoft, new.MaxRunSoft},
		{"max_running", old.MaxRunning, new.MaxRunning},
		{"max_user_res", old.MaxUserRes, new.MaxUserRes},
		{"max_user_res_soft", old.MaxUserResSoft, new.MaxUserResSoft},
		{"max_user_run", old.MaxUserRun, new.MaxUserRun},
		{"max_user_run_soft", old.MaxUserRunSoft, new.MaxUserRunSoft},
		{"node_fail_requeue", old.NodeFailRequeue, new.NodeFailRequeue},
		{"node_group_enable", old.NodeGroupEnable, new.NodeGroupEnable},
		{"node_group_key", old.NodeGroupKey, new.NodeGroupKey},
		{"operators", old.Operators, new.Operators},
		{"pbs_license_info", old.PbsLicenseInfo, new.PbsLicenseInfo},
		{"pbs_license_linger_time", old.PbsLicenseLingerTime, new.PbsLicenseLingerTime},
		{"pbs_license_max", old.PbsLicenseMax, new.PbsLicenseMax},
		{"pbs_license_min", old.PbsLicenseMin, new.PbsLicenseMin},
		{"power_provisioning", old.PowerProvisioning, new.PowerProvisioning},
		{"python_gc_min_interval", old.PythonGcMinInterval, new.PythonGcMinInterval},
		{"python_restart_max_pbs_servers", old.PythonRestartMaxPbsServers, new.PythonRestartMaxPbsServers},
		{"python_restart_max_objects", old.PythonRestartMaxObjects, new.PythonRestartMaxObjects},
		{"python_restart_min_interval", old.PythonRestartMinInterval, new.PythonRestartMinInterval},
		{"query_other_jobs", old.QueryOtherJobs, new.QueryOtherJobs},
		{"queued_jobs_threshold", old.QueuedJobsThreshold, new.QueuedJobsThreshold},
		{"queued_jobs_threshold_res", old.QueuedJobsThresholdRes, new.QueuedJobsThresholdRes},
		{"reserve_retry_init", old.ReserveRetryInit, new.ReserveRetryInit},
		{"reserve_retry_time", old.ReserveRetryTime, new.ReserveRetryTime},
		{"resources_available", old.ResourcesAvailable, new.ResourcesAvailable},
		{"resources_default", old.ResourcesDefault, new.ResourcesDefault},
		{"resources_max", old.ResourcesMax, new.ResourcesMax},
		{"restrict_res_to_release_on_suspend", old.RestrictResToReleaseOnSuspend, new.RestrictResToReleaseOnSuspend},
		{"resv_enable", old.ResvEnable, new.ResvEnable},
		{"resv_post_processing_time", old.ResvPostProcessingTime, new.ResvPostProcessingTime},
		{"rpp_highwater", old.RppHighwater, new.RppHighwater},
		{"rpp_max_pkt_check", old.RppMaxPktCheck, new.RppMaxPktCheck},
		{"rpp_retry", old.RppRetry, new.RppRetry},
		{"scheduler_iteration", old.SchedulerIteration, new.SchedulerIteration},
		{"webapi_auth_issuers", old.WebapiAuthIssuers, new.WebapiAuthIssuers},
		{"webapi_enable", old.WebapiEnable, new.WebapiEnable},
		{"webapi_oidc_clientid", old.WebapiOidcClientid, new.WebapiOidcClientid},
		{"webapi_oidc_provider_url", old.WebapiOidcProviderUrl, new.WebapiOidcProviderUrl},
	}
	for _, v := range fields {
		newCommands, err := generateUpdateAttributeCommand(v.old, v.new, "server", new.Name, v.attribute)
		if err != nil {
			return old, err
		}
		commands = append(commands, newCommands...)
	}

	_, errOutput, err := c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return old, fmt.Errorf("%s %s %s", err, completeErrOutput, strings.Join(commands, ","))
	}

	return c.GetPbsServer(old.Name)
}

func (c *PbsClient) DeletePbsServer(name string) error {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete server %s'", name)
	_, errOutput, err := c.runCommand(cmd)
	if err != nil {
		return fmt.Errorf("%s %s", err, errOutput)
	}

	return nil
}
