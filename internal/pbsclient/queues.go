package pbsclient

import (
	"fmt"
	"strconv"
	"strings"
)

const GET_QUEUE_QMGR_CMD = "/opt/pbs/bin/qmgr -c 'list queue %s'"
const GET_QUEUES_QMGR_CMD = "/opt/pbs/bin/qmgr -c 'list queue @default'"
const CREATE_QUEUE_QMGR_CMD = "/opt/pbs/bin/qmgr -c 'create queue %s queue_type=%s'"

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
	queueOutput, err := client.runCommand(GET_QUEUES_QMGR_CMD)
	if err != nil {
		return nil, err
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
		command := ""
		switch v.old.(type) {
		case bool:
			oldValue := v.old.(bool)
			newValue := v.new.(bool)
			command = generateUpdateBoolAttributeCommand("queue", updatedQueue.Name, v.attribute, &oldValue, &newValue)
		case *bool:
			command = generateUpdateBoolAttributeCommand("queue", updatedQueue.Name, v.attribute, v.old.(*bool), v.new.(*bool))
		case int32:
			oldValue := v.old.(int32)
			newValue := v.new.(int32)
			command = generateUpdateInt32AttributeCommand("queue", updatedQueue.Name, v.attribute, &oldValue, &newValue)
		case *int32:
			command = generateUpdateInt32AttributeCommand("queue", updatedQueue.Name, v.attribute, v.old.(*int32), v.new.(*int32))
		case string:
			oldValue := v.old.(string)
			newValue := v.new.(string)
			command = generateUpdateStringAttributeCommand("queue", updatedQueue.Name, v.attribute, &oldValue, &newValue)
		case *string:
			command = generateUpdateStringAttributeCommand("queue", updatedQueue.Name, v.attribute, v.old.(*string), v.new.(*string))
		case map[string]string:
			oldValue := v.old.(map[string]string)
			newValue := v.new.(map[string]string)
			for k, oldAttrVal := range oldValue {
				newAttrVal, ok := newValue[k]
				if !ok {
					command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset queue %s %s.%s'", updatedQueue.Name, v.attribute, k)
				} else if oldAttrVal != newAttrVal {
					command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s %s.%s=%s'", updatedQueue.Name, v.attribute, k, newAttrVal)
				}
			}
			for k, newAttrVal := range newValue {
				if _, ok := oldValue[k]; !ok {
					command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s %s.%s=%s'", updatedQueue.Name, v.attribute, k, newAttrVal)
				}
			}

		default:
			return oldQueue, fmt.Errorf("unsupported type %T", v.old)
		}

		if command != "" {
			commands = append(commands, command)
		}

	}

	_, err = client.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		return oldQueue, err
	}

	oldQueue, err = client.GetQueue(oldQueue.Name)
	if err != nil {
		return oldQueue, err
	}

	return oldQueue, nil
}

func (client *PbsClient) CreateQueue(newQueue PbsQueue) (PbsQueue, error) {
	var commands = []string{
		fmt.Sprintf(CREATE_QUEUE_QMGR_CMD, newQueue.Name, newQueue.QueueType),
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s enabled=%s'", newQueue.Name, strconv.FormatBool(newQueue.Enabled)),
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s started=%s'", newQueue.Name, strconv.FormatBool(newQueue.Started)),
	}

	fields := []struct {
		attribute string
		new       any
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
		command := ""
		switch v.new.(type) {
		case *bool:
			b := v.new.(*bool)
			if b != nil {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s %s=%s'", newQueue.Name, v.attribute, strconv.FormatBool(*b))
			}
		case *int32:
			i := v.new.(*int32)
			if i != nil {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s %s=%s'", newQueue.Name, v.attribute, strconv.Itoa(int(*i)))
			}
		case *string:
			s := v.new.(*string)
			if s != nil {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s %s=%s'", newQueue.Name, v.attribute, *s)
			}
		case map[string]string:
			m := v.new.(map[string]string)
			for k, subval := range m {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set queue %s %s.%s=%s'", newQueue.Name, v.attribute, k, subval)
			}
		default:
			return newQueue, fmt.Errorf("unsupported type %T", v.new)
		}

		if command != "" {
			commands = append(commands, command)
		}
	}

	_, err := client.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		var newQueue PbsQueue
		return newQueue, err
	}
	newQueue, err = client.GetQueue(newQueue.Name)
	if err != nil {
		return newQueue, err
	}

	return newQueue, nil
}

func (client *PbsClient) DeleteQueue(name string) error {
	_, err := client.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete queue %s'", name))
	if err != nil {
		return err
	}

	return nil
}
