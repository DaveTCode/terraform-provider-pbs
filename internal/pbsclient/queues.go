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
	MaxGroupRes            *int32
	MaxGroupResSoft        *int32
	MaxGroupRun            *int32
	MaxGroupRunSoft        *int32
	MaxQueuable            *int32
	MaxQueued              *string
	MaxQueuedRes           *string
	MaxRun                 *string
	MaxRunRes              *string
	MaxRunResSoft          *string
	MaxRunSoft             *string
	MaxRunning             *int32
	MaxUserRes             *string
	MaxUserResSoft         *string
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
	StateCount             string
	TotalJobs              int32
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
					case "max_group_res":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxGroupRes = &i32Value
					case "max_group_res_soft":
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxGroupResSoft = &i32Value
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
						intValue, err := strconv.Atoi(s)
						i32Value := int32(intValue)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.MaxRunning = &i32Value
					case "max_user_res":
						current.MaxUserRes = &s
					case "max_user_res_soft":
						current.MaxUserResSoft = &s
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
					case "state_count":
						current.StateCount = s
					case "total_jobs":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						current.TotalJobs = int32(intValue)
					default:
						// TODO - What to do with attributes we don't recognise?
					}
				} else if a, ok := value.(map[string]string); ok {
					switch strings.ToLower(k) {
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

// GetQueue returns a single queue by name
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

// GetQueues returns all queues configured on the PBS server
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

func (client *PbsClient) UpdateQueue(updatedQueue PbsQueue) (PbsQueue, error) {
	oldQueue, err := client.GetQueue(updatedQueue.Name)
	if err != nil {
		return oldQueue, err
	}

	var commands = []string{}
	fields := []struct {
		attribute string
		old       any
		new       any
	}{
		{"acl_group_enable", oldQueue.AclGroupEnable, updatedQueue.AclGroupEnable},
		{"acl_groups", oldQueue.AclGroups, updatedQueue.AclGroups},
		{"acl_host_enable", oldQueue.AclHostEnable, updatedQueue.AclHostEnable},
		{"acl_hosts", oldQueue.AclHosts, updatedQueue.AclHosts},
		{"acl_user_enable", oldQueue.AclUserEnable, updatedQueue.AclUserEnable},
		{"acl_users", oldQueue.AclUsers, updatedQueue.AclUsers},
		{"alt_router", oldQueue.AltRouter, updatedQueue.AltRouter},
		{"backfill_depth", oldQueue.BackfillDepth, updatedQueue.BackfillDepth},
		{"checkpoint_min", oldQueue.CheckpointMin, updatedQueue.CheckpointMin},
		{"default_chunk", oldQueue.DefaultChunk, updatedQueue.DefaultChunk},
		{"enabled", oldQueue.Enabled, updatedQueue.Enabled},
		{"from_route_only", oldQueue.FromRouteOnly, updatedQueue.FromRouteOnly},
		{"kill_delay", oldQueue.KillDelay, updatedQueue.KillDelay},
		{"max_array_size", oldQueue.MaxArraySize, updatedQueue.MaxArraySize},
		{"max_group_res", oldQueue.MaxGroupRes, updatedQueue.MaxGroupRes},
		{"max_group_res_soft", oldQueue.MaxGroupResSoft, updatedQueue.MaxGroupResSoft},
		{"max_group_run", oldQueue.MaxGroupRun, updatedQueue.MaxGroupRun},
		{"max_group_run_soft", oldQueue.MaxGroupRunSoft, updatedQueue.MaxGroupRunSoft},
		{"max_queuable", oldQueue.MaxQueuable, updatedQueue.MaxQueuable},
		{"max_queued", oldQueue.MaxQueued, updatedQueue.MaxQueued},
		{"max_queued_res", oldQueue.MaxQueuedRes, updatedQueue.MaxQueuedRes},
		{"max_run", oldQueue.MaxRun, updatedQueue.MaxRun},
		{"max_run_res", oldQueue.MaxRunRes, updatedQueue.MaxRunRes},
		{"max_run_res_soft", oldQueue.MaxRunResSoft, updatedQueue.MaxRunResSoft},
		{"max_run_soft", oldQueue.MaxRunSoft, updatedQueue.MaxRunSoft},
		{"max_running", oldQueue.MaxRunning, updatedQueue.MaxRunning},
		{"max_user_res", oldQueue.MaxUserRes, updatedQueue.MaxUserRes},
		{"max_user_res_soft", oldQueue.MaxUserResSoft, updatedQueue.MaxUserResSoft},
		{"max_user_run", oldQueue.MaxUserRun, updatedQueue.MaxUserRun},
		{"max_user_run_soft", oldQueue.MaxUserRunSoft, updatedQueue.MaxUserRunSoft},
		{"node_group_key", oldQueue.NodeGroupKey, updatedQueue.NodeGroupKey},
		{"partition", oldQueue.Partition, updatedQueue.Partition},
		{"priority", oldQueue.Priority, updatedQueue.Priority},
		{"queued_jobs_threshold", oldQueue.QueuedJobsThreshold, updatedQueue.QueuedJobsThreshold},
		{"queued_jobs_threshold_res", oldQueue.QueuedJobsThresholdRes, updatedQueue.QueuedJobsThresholdRes},
		{"queue_type", oldQueue.QueueType, updatedQueue.QueueType},
		{"resources_assigned", oldQueue.ResourcesAssigned, updatedQueue.ResourcesAssigned},
		{"resources_available", oldQueue.ResourcesAvailable, updatedQueue.ResourcesAvailable},
		{"resources_default", oldQueue.ResourcesDefault, updatedQueue.ResourcesDefault},
		{"resources_max", oldQueue.ResourcesMax, updatedQueue.ResourcesMax},
		{"resources_min", oldQueue.ResourcesMin, updatedQueue.ResourcesMin},
		{"route_destinations", oldQueue.RouteDestinations, updatedQueue.RouteDestinations},
		{"route_held_jobs", oldQueue.RouteHeldJobs, updatedQueue.RouteHeldJobs},
		{"route_lifetime", oldQueue.RouteLifetime, updatedQueue.RouteLifetime},
		{"route_retry_time", oldQueue.RouteRetryTime, updatedQueue.RouteRetryTime},
		{"route_waiting_jobs", oldQueue.RouteWaitingJobs, updatedQueue.RouteWaitingJobs},
		{"started", oldQueue.Started, updatedQueue.Started},
	}
	for _, v := range fields {
		newCommands, err := generateUpdateAttributeCommand(v.old, v.new, "queue", updatedQueue.Name, v.attribute)
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

func (client *PbsClient) CreateQueue(new PbsQueue) (PbsQueue, error) {
	var commands = []string{
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create queue %s queue_type=%s'", new.Name, new.QueueType),
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s enabled=%s'", new.Name, strconv.FormatBool(new.Enabled)),
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s started=%s'", new.Name, strconv.FormatBool(new.Started)),
	}

	fields := []struct {
		attribute string
		new       any
	}{
		{"acl_group_enable", new.AclGroupEnable},
		{"acl_groups", new.AclGroups},
		{"acl_host_enable", new.AclHostEnable},
		{"acl_hosts", new.AclHosts},
		{"acl_user_enable", new.AclUserEnable},
		{"acl_users", new.AclUsers},
		{"alt_router", new.AltRouter},
		{"backfill_depth", new.BackfillDepth},
		{"checkpoint_min", new.CheckpointMin},
		{"default_chunk", new.DefaultChunk},
		{"from_route_only", new.FromRouteOnly},
		{"kill_delay", new.KillDelay},
		{"max_array_size", new.MaxArraySize},
		{"max_group_res", new.MaxGroupRes},
		{"max_group_res_soft", new.MaxGroupResSoft},
		{"max_group_run", new.MaxGroupRun},
		{"max_group_run_soft", new.MaxGroupRunSoft},
		{"max_queuable", new.MaxQueuable},
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
		{"node_group_key", new.NodeGroupKey},
		{"partition", new.Partition},
		{"priority", new.Priority},
		{"queued_jobs_threshold", new.QueuedJobsThreshold},
		{"queued_jobs_threshold_res", new.QueuedJobsThresholdRes},
		{"resources_assigned", new.ResourcesAssigned},
		{"resources_available", new.ResourcesAvailable},
		{"resources_default", new.ResourcesDefault},
		{"resources_max", new.ResourcesMax},
		{"resources_min", new.ResourcesMin},
		{"route_destinations", new.RouteDestinations},
		{"route_held_jobs", new.RouteHeldJobs},
		{"route_lifetime", new.RouteLifetime},
		{"route_retry_time", new.RouteRetryTime},
		{"route_waiting_jobs", new.RouteWaitingJobs},
	}
	for _, v := range fields {
		c, err := generateCreateCommands(v.new, "node", new.Name, v.attribute)
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

	new, err = client.GetQueue(new.Name)
	if err != nil {
		return new, err
	}

	return new, nil
}

func (client *PbsClient) DeleteQueue(name string) error {
	_, errOutput, err := client.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete queue %s'", name))
	if err != nil {
		return fmt.Errorf("%s %s", err, errOutput)
	}

	return nil
}
