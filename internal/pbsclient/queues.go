package pbsclient

import (
	"fmt"
	"strconv"
	"strings"
)

const GET_QUEUE_QMGR_CMD = "/opt/pbs/bin/qmgr -c 'list queue %s'"
const GET_QUEUES_QMGR_CMD = "/opt/pbs/bin/qmgr -c 'list queue @default'"

type PbsQueue struct {
	AclGroupEnable         *bool
	AclGroups              *string
	AclHostEnable          *bool
	AclHosts               *string
	AclUserEnable          *bool
	AclUsers               *string
	AltRouter              *string
	BackfillDepth          *int32
	CheckpointMin          *int32
	DefaultChunk           *string
	Enabled                bool
	FromRouteOnly          *bool
	KillDelay              *int32
	MaxArraySize           *int32
	MaxGroupRes            map[string]string
	MaxGroupResSoft        map[string]string
	MaxGroupRun            *int32
	MaxGroupRunSoft        *int32
	MaxQueuable            *int32
	MaxQueued              map[string]string
	MaxQueuedRes           map[string]string
	MaxRun                 map[string]string
	MaxRunRes              map[string]string
	MaxRunResSoft          map[string]string
	MaxRunSoft             map[string]string
	MaxRunning             *int32
	MaxUserRes             map[string]string
	MaxUserResSoft         map[string]string
	MaxUserRun             *int32
	MaxUserRunSoft         *int32
	Name                   string
	NodeGroupKey           *string
	Partition              *string
	Priority               *int32
	QueuedJobsThreshold    *string
	QueuedJobsThresholdRes *string
	QueueType              string
	ResourcesAssigned      map[string]string
	ResourcesAvailable     map[string]string
	ResourcesDefault       map[string]string
	ResourcesMax           map[string]string
	ResourcesMin           map[string]string
	RouteDestinations      *string
	RouteHeldJobs          *bool
	RouteLifetime          *int32
	RouteRetryTime         *int32
	RouteWaitingJobs       *bool
	Started                bool
}

func parseQueueOutput(output []byte) ([]PbsQueue, error) {
	parsedOutput := parseGenericQmgrOutput(string(output))
	var queues []PbsQueue

	for _, r := range parsedOutput {
		if r.objType == "Queue" {
			current := PbsQueue{
				Name: r.name,
			}

			for k, value := range r.attributes {
				if s, ok := value.(string); ok {
					switch strings.ToLower(k) {
					case "acl_group_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.AclGroupEnable = &boolValue
					case "acl_groups":
						current.AclGroups = &s
					case "acl_host_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.AclHostEnable = &boolValue
					case "acl_hosts":
						current.AclHosts = &s
					case "acl_user_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.AclUserEnable = &boolValue
					case "acl_users":
						current.AclUsers = &s
					case "alt_router":
						current.AltRouter = &s
					case "backfill_depth":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.BackfillDepth = &i32Value
					case "checkpoint_min":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.CheckpointMin = &i32Value
					case "default_chunk":
						current.DefaultChunk = &s
					case "enabled": //                bool
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.Enabled = boolValue
					case "from_route_only":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.FromRouteOnly = &boolValue
					case "kill_delay":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.KillDelay = &i32Value
					case "max_array_size":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxArraySize = &i32Value
					case "max_group_run":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxGroupRun = &i32Value
					case "max_group_run_soft":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxGroupRunSoft = &i32Value
					case "max_queuable":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxQueuable = &i32Value
					case "max_running":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxRunning = &i32Value
					case "max_user_run":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxUserRun = &i32Value
					case "max_user_run_soft":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxUserRunSoft = &i32Value
					case "node_group_key":
						current.NodeGroupKey = &s
					case "partition":
						current.Partition = &s
					case "priority":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.Priority = &i32Value
					case "queued_jobs_threshold":
						current.QueuedJobsThreshold = &s
					case "queued_jobs_threshold_res":
						current.QueuedJobsThresholdRes = &s
					case "queue_type":
						current.QueueType = s
					case "route_destinations":
						current.RouteDestinations = &s
					case "route_held_jobs":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.RouteHeldJobs = &boolValue
					case "route_lifetime":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.RouteLifetime = &i32Value
					case "route_retry_time":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.RouteRetryTime = &i32Value
					case "route_waiting_jobs":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.RouteWaitingJobs = &boolValue
					case "started":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.Started = boolValue
					default:
						// TODO - What to do with attributes we don't recognise?
					}
				} else if a, ok := value.(map[string]string); ok {
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
					case "resources_assigned":
						current.ResourcesAssigned = a
					case "resources_available":
						current.ResourcesAvailable = a
					case "resources_default":
						current.ResourcesDefault = a
					case "resources_max":
						current.ResourcesMax = a
					case "resources_min":
						current.ResourcesMin = a
					}
				}
			}

			queues = append(queues, current)
		}
	}

	return queues, nil
}

// GetQueue returns a single queue by name.
func (client *PbsClient) GetQueue(name string) (PbsQueue, error) {
	all, err := client.GetQueues()
	if err != nil {
		return PbsQueue{}, err
	}

	for _, r := range all {
		if r.Name == name {
			return r, nil
		}
	}

	return PbsQueue{}, nil
}

// GetQueues returns all queues configured on the PBS server.
func (client *PbsClient) GetQueues() ([]PbsQueue, error) {
	queueOutput, errOutput, err := client.runCommand(GET_QUEUES_QMGR_CMD)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command against PBS server %s: %s", err.Error(), errOutput)
	}

	queues, err := parseQueueOutput(queueOutput)
	if err != nil {
		return nil, err
	}

	return queues, nil
}

func (client *PbsClient) UpdateQueue(newQueue PbsQueue) (PbsQueue, error) {
	oldQueue, err := client.GetQueue(newQueue.Name)
	if err != nil {
		return oldQueue, err
	}

	var commands = []string{}
	fields := []struct {
		attribute string
		oldAttr   any
		newAttr   any
	}{
		{"acl_group_enable", oldQueue.AclGroupEnable, newQueue.AclGroupEnable},
		{"acl_groups", oldQueue.AclGroups, newQueue.AclGroups},
		{"acl_host_enable", oldQueue.AclHostEnable, newQueue.AclHostEnable},
		{"acl_hosts", oldQueue.AclHosts, newQueue.AclHosts},
		{"acl_user_enable", oldQueue.AclUserEnable, newQueue.AclUserEnable},
		{"acl_users", oldQueue.AclUsers, newQueue.AclUsers},
		{"alt_router", oldQueue.AltRouter, newQueue.AltRouter},
		{"backfill_depth", oldQueue.BackfillDepth, newQueue.BackfillDepth},
		{"checkpoint_min", oldQueue.CheckpointMin, newQueue.CheckpointMin},
		{"default_chunk", oldQueue.DefaultChunk, newQueue.DefaultChunk},
		{"enabled", oldQueue.Enabled, newQueue.Enabled},
		{"from_route_only", oldQueue.FromRouteOnly, newQueue.FromRouteOnly},
		{"kill_delay", oldQueue.KillDelay, newQueue.KillDelay},
		{"max_array_size", oldQueue.MaxArraySize, newQueue.MaxArraySize},
		{"max_group_res", oldQueue.MaxGroupRes, newQueue.MaxGroupRes},
		{"max_group_res_soft", oldQueue.MaxGroupResSoft, newQueue.MaxGroupResSoft},
		{"max_group_run", oldQueue.MaxGroupRun, newQueue.MaxGroupRun},
		{"max_group_run_soft", oldQueue.MaxGroupRunSoft, newQueue.MaxGroupRunSoft},
		{"max_queuable", oldQueue.MaxQueuable, newQueue.MaxQueuable},
		{"max_queued", oldQueue.MaxQueued, newQueue.MaxQueued},
		{"max_queued_res", oldQueue.MaxQueuedRes, newQueue.MaxQueuedRes},
		{"max_run", oldQueue.MaxRun, newQueue.MaxRun},
		{"max_run_res", oldQueue.MaxRunRes, newQueue.MaxRunRes},
		{"max_run_res_soft", oldQueue.MaxRunResSoft, newQueue.MaxRunResSoft},
		{"max_run_soft", oldQueue.MaxRunSoft, newQueue.MaxRunSoft},
		{"max_running", oldQueue.MaxRunning, newQueue.MaxRunning},
		{"max_user_res", oldQueue.MaxUserRes, newQueue.MaxUserRes},
		{"max_user_res_soft", oldQueue.MaxUserResSoft, newQueue.MaxUserResSoft},
		{"max_user_run", oldQueue.MaxUserRun, newQueue.MaxUserRun},
		{"max_user_run_soft", oldQueue.MaxUserRunSoft, newQueue.MaxUserRunSoft},
		{"node_group_key", oldQueue.NodeGroupKey, newQueue.NodeGroupKey},
		{"partition", oldQueue.Partition, newQueue.Partition},
		{"priority", oldQueue.Priority, newQueue.Priority},
		{"queued_jobs_threshold", oldQueue.QueuedJobsThreshold, newQueue.QueuedJobsThreshold},
		{"queued_jobs_threshold_res", oldQueue.QueuedJobsThresholdRes, newQueue.QueuedJobsThresholdRes},
		{"queue_type", oldQueue.QueueType, newQueue.QueueType},
		{"resources_assigned", oldQueue.ResourcesAssigned, newQueue.ResourcesAssigned},
		{"resources_available", oldQueue.ResourcesAvailable, newQueue.ResourcesAvailable},
		{"resources_default", oldQueue.ResourcesDefault, newQueue.ResourcesDefault},
		{"resources_max", oldQueue.ResourcesMax, newQueue.ResourcesMax},
		{"resources_min", oldQueue.ResourcesMin, newQueue.ResourcesMin},
		{"route_destinations", oldQueue.RouteDestinations, newQueue.RouteDestinations},
		{"route_held_jobs", oldQueue.RouteHeldJobs, newQueue.RouteHeldJobs},
		{"route_lifetime", oldQueue.RouteLifetime, newQueue.RouteLifetime},
		{"route_retry_time", oldQueue.RouteRetryTime, newQueue.RouteRetryTime},
		{"route_waiting_jobs", oldQueue.RouteWaitingJobs, newQueue.RouteWaitingJobs},
		{"started", oldQueue.Started, newQueue.Started},
	}
	for _, v := range fields {
		newCommands, err := generateUpdateAttributeCommand(v.oldAttr, v.newAttr, "queue", newQueue.Name, v.attribute)
		if err != nil {
			return oldQueue, err
		}
		commands = append(commands, newCommands...)
	}

	_, errOutput, err := client.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return oldQueue, fmt.Errorf("%s, %s, %s", err, strings.Join(commands, ","), completeErrOutput)
	}

	oldQueue, err = client.GetQueue(oldQueue.Name)
	if err != nil {
		return oldQueue, err
	}

	return oldQueue, nil
}

func (client *PbsClient) CreateQueue(newQueue PbsQueue) (PbsQueue, error) {
	var commands = []string{
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create queue %s queue_type=%s'", newQueue.Name, newQueue.QueueType),
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s enabled=%s'", newQueue.Name, strconv.FormatBool(newQueue.Enabled)),
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s started=%s'", newQueue.Name, strconv.FormatBool(newQueue.Started)),
	}

	fields := []struct {
		attribute string
		newAttr   any
	}{
		{"acl_group_enable", newQueue.AclGroupEnable},
		{"acl_groups", newQueue.AclGroups},
		{"acl_host_enable", newQueue.AclHostEnable},
		{"acl_hosts", newQueue.AclHosts},
		{"acl_user_enable", newQueue.AclUserEnable},
		{"acl_users", newQueue.AclUsers},
		{"alt_router", newQueue.AltRouter},
		{"backfill_depth", newQueue.BackfillDepth},
		{"checkpoint_min", newQueue.CheckpointMin},
		{"default_chunk", newQueue.DefaultChunk},
		{"from_route_only", newQueue.FromRouteOnly},
		{"kill_delay", newQueue.KillDelay},
		{"max_array_size", newQueue.MaxArraySize},
		{"max_group_res", newQueue.MaxGroupRes},
		{"max_group_res_soft", newQueue.MaxGroupResSoft},
		{"max_group_run", newQueue.MaxGroupRun},
		{"max_group_run_soft", newQueue.MaxGroupRunSoft},
		{"max_queuable", newQueue.MaxQueuable},
		{"max_queued", newQueue.MaxQueued},
		{"max_queued_res", newQueue.MaxQueuedRes},
		{"max_run", newQueue.MaxRun},
		{"max_run_res", newQueue.MaxRunRes},
		{"max_run_res_soft", newQueue.MaxRunResSoft},
		{"max_run_soft", newQueue.MaxRunSoft},
		{"max_running", newQueue.MaxRunning},
		{"max_user_res", newQueue.MaxUserRes},
		{"max_user_res_soft", newQueue.MaxUserResSoft},
		{"max_user_run", newQueue.MaxUserRun},
		{"max_user_run_soft", newQueue.MaxUserRunSoft},
		{"node_group_key", newQueue.NodeGroupKey},
		{"partition", newQueue.Partition},
		{"priority", newQueue.Priority},
		{"queued_jobs_threshold", newQueue.QueuedJobsThreshold},
		{"queued_jobs_threshold_res", newQueue.QueuedJobsThresholdRes},
		{"resources_assigned", newQueue.ResourcesAssigned},
		{"resources_available", newQueue.ResourcesAvailable},
		{"resources_default", newQueue.ResourcesDefault},
		{"resources_max", newQueue.ResourcesMax},
		{"resources_min", newQueue.ResourcesMin},
		{"route_destinations", newQueue.RouteDestinations},
		{"route_held_jobs", newQueue.RouteHeldJobs},
		{"route_lifetime", newQueue.RouteLifetime},
		{"route_retry_time", newQueue.RouteRetryTime},
		{"route_waiting_jobs", newQueue.RouteWaitingJobs},
	}
	for _, v := range fields {
		c, err := generateCreateCommands(v.newAttr, "queue", newQueue.Name, v.attribute)
		if err != nil {
			return PbsQueue{}, err
		}
		commands = append(commands, c...)
	}

	_, errOutput, err := client.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return PbsQueue{}, fmt.Errorf("%s, %s, %s", err, strings.Join(commands, ","), completeErrOutput)
	}

	newQueue, err = client.GetQueue(newQueue.Name)
	if err != nil {
		return newQueue, err
	}

	return newQueue, nil
}

func (client *PbsClient) DeleteQueue(name string) error {
	_, errOutput, err := client.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete queue %s'", name))
	if err != nil {
		return fmt.Errorf("%s %s", err, errOutput)
	}

	return nil
}
