package pbsclient

import (
	"fmt"
	"regexp"
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
	ResourcesAssigned      *string
	ResourcesAvailable     *string
	ResourcesDefault       *string
	ResourcesMax           *string
	ResourcesMin           *string
	RouteDestinations      *string
	RouteHeldJobs          *bool
	RouteLifetime          *int32
	RouteRetryTime         *int32
	RouteWaitingJobs       *bool
	Started                bool
	StateCount             string
	TotalJobs              int32
}

var (
	queueNameRegex = regexp.MustCompile(`Queue\s+(\w+)`)
	attributeRegex = regexp.MustCompile(`\s+(\w+)\s*=\s*(.*)`)
)

func parseQueueOutput(output string) ([]PbsQueue, error) {
	var currentQueue PbsQueue
	var queues []PbsQueue
	for line := range strings.SplitSeq(string(output), "\n") {
		if queueNameRegex.MatchString(line) {
			if currentQueue.Name != "" { // Is there a queue currently being processed? If so add it to the completed list
				queues = append(queues, currentQueue)
			}

			currentQueue = PbsQueue{
				Name: queueNameRegex.FindStringSubmatch(line)[1],
			}
		} else if attributeRegex.MatchString(line) {
			subMatch := attributeRegex.FindStringSubmatch(line)
			attribute := subMatch[1]
			value := subMatch[2]

			switch strings.ToLower(attribute) {
			case "acl_group_enable":
				boolValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to bool %s", attribute, err.Error())
				}
				currentQueue.AclGroupEnable = &boolValue
			case "acl_groups":
				currentQueue.AclGroups = &value
			case "acl_host_enable":
				boolValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to bool %s", attribute, err.Error())
				}
				currentQueue.AclHostEnable = &boolValue
			case "acl_hosts":
				currentQueue.AclHosts = &value
			case "acl_user_enable":
				boolValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to bool %s", attribute, err.Error())
				}
				currentQueue.AclUserEnable = &boolValue
			case "acl_users":
				currentQueue.AclUsers = &value
			case "alt_router":
				currentQueue.AltRouter = &value
			case "backfill_depth":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.BackfillDepth = &i32Value
			case "checkpoint_min":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.CheckpointMin = &i32Value
			case "default_chunk":
				currentQueue.DefaultChunk = &value
			case "enabled": //                bool
				boolValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to bool %s", attribute, err.Error())
				}
				currentQueue.Enabled = boolValue
			case "from_route_only":
				boolValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to bool %s", attribute, err.Error())
				}
				currentQueue.FromRouteOnly = &boolValue
			case "kill_delay":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.KillDelay = &i32Value
			case "max_array_size":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.MaxArraySize = &i32Value
			case "max_group_res":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.MaxGroupRes = &i32Value
			case "max_group_res_soft":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.MaxGroupResSoft = &i32Value
			case "max_group_run":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.MaxGroupRun = &i32Value
			case "max_group_run_soft":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.MaxGroupRunSoft = &i32Value
			case "max_queuable":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.MaxQueuable = &i32Value
			case "max_queued":
				currentQueue.MaxQueued = &value
			case "max_queued_res":
				currentQueue.MaxQueuedRes = &value
			case "max_run":
				currentQueue.MaxRun = &value
			case "max_run_res":
				currentQueue.MaxRunRes = &value
			case "max_run_res_soft":
				currentQueue.MaxRunResSoft = &value
			case "max_run_soft":
				currentQueue.MaxRunSoft = &value
			case "max_running":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.MaxRunning = &i32Value
			case "max_user_res":
				currentQueue.MaxUserRes = &value
			case "max_user_res_soft":
				currentQueue.MaxUserResSoft = &value
			case "max_user_run":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.MaxUserRun = &i32Value
			case "max_user_run_soft":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.MaxUserRunSoft = &i32Value
			case "node_group_key":
				currentQueue.NodeGroupKey = &value
			case "partition":
				currentQueue.Partition = &value
			case "priority":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.Priority = &i32Value
			case "queued_jobs_threshold":
				currentQueue.QueuedJobsThreshold = &value
			case "queued_jobs_threshold_res":
				currentQueue.QueuedJobsThresholdRes = &value
			case "queue_type":
				currentQueue.QueueType = value
			case "resources_assigned":
				currentQueue.ResourcesAssigned = &value
			case "resources_available":
				currentQueue.ResourcesAvailable = &value
			case "resources_default":
				currentQueue.ResourcesDefault = &value
			case "resources_max":
				currentQueue.ResourcesMax = &value
			case "resources_min":
				currentQueue.ResourcesMin = &value
			case "route_destinations":
				currentQueue.RouteDestinations = &value
			case "route_held_jobs":
				boolValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to bool %s", attribute, err.Error())
				}
				currentQueue.RouteHeldJobs = &boolValue
			case "route_lifetime":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.RouteLifetime = &i32Value
			case "route_retry_time":
				intValue, err := strconv.Atoi(value)
				i32Value := int32(intValue)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.RouteRetryTime = &i32Value
			case "route_waiting_jobs":
				boolValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to bool %s", attribute, err.Error())
				}
				currentQueue.RouteWaitingJobs = &boolValue
			case "started":
				boolValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to bool %s", attribute, err.Error())
				}
				currentQueue.Started = boolValue
			case "state_count":
				currentQueue.StateCount = value
			case "total_jobs":
				intValue, err := strconv.Atoi(value)
				if err != nil {
					return nil, fmt.Errorf("failed to convert %s value to int %s", attribute, err.Error())
				}
				currentQueue.TotalJobs = int32(intValue)
			default:
				// TODO - What to do with attributes we don't recognise?
			}
		}
	}

	if currentQueue.Name != "" {
		queues = append(queues, currentQueue)
	}

	return queues, nil
}

// GetQueue returns a single queue by name
func (client *PbsClient) GetQueue(name string) (PbsQueue, error) {
	var queue PbsQueue
	queueOutput, err := client.runCommand(fmt.Sprintf(GET_QUEUE_QMGR_CMD, name))
	if err != nil {
		return queue, err
	}

	queues, err := parseQueueOutput(string(queueOutput))
	if err != nil {
		return queue, err
	}
	if len(queues) != 1 {
		return queue, fmt.Errorf("queue %s not found", name)
	}

	return queues[0], nil
}

// GetQueues returns all queues configured on the PBS server
func (client *PbsClient) GetQueues() ([]PbsQueue, error) {
	queueOutput, err := client.runCommand(GET_QUEUES_QMGR_CMD)
	if err != nil {
		return nil, err
	}

	queues, err := parseQueueOutput(string(queueOutput))
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
