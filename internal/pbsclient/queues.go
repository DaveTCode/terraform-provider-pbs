package pbsclient

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const GET_QUEUE_QMGR_CMD = "/opt/pbs/bin/qmgr -c 'list queue %s'"
const GET_QUEUES_QMGR_CMD = "/opt/pbs/bin/qmgr -c 'list queue @default'"

// queueFieldDefinition represents a queue field with its attribute name, execution order, and value extractor.
type queueFieldDefinition struct {
	attribute string
	order     int                      // Lower numbers execute first
	getValue  func(queue PbsQueue) any // Function to extract the value from a PbsQueue
}

// getQueueFieldDefinitions returns the ordered list of queue field definitions.
// This ensures consistent ordering across create and update operations.
func getQueueFieldDefinitions() []queueFieldDefinition {
	return []queueFieldDefinition{
		{"acl_group_enable", 10, func(q PbsQueue) any { return q.AclGroupEnable }},
		{"acl_groups", 10, func(q PbsQueue) any { return q.AclGroups }},
		{"acl_host_enable", 10, func(q PbsQueue) any { return q.AclHostEnable }},
		{"acl_hosts", 10, func(q PbsQueue) any { return q.AclHosts }},
		{"acl_user_enable", 10, func(q PbsQueue) any { return q.AclUserEnable }},
		{"acl_users", 10, func(q PbsQueue) any { return q.AclUsers }},
		{"alt_router", 10, func(q PbsQueue) any { return q.AltRouter }},
		{"backfill_depth", 10, func(q PbsQueue) any { return q.BackfillDepth }},
		{"checkpoint_min", 10, func(q PbsQueue) any { return q.CheckpointMin }},
		{"default_chunk", 10, func(q PbsQueue) any { return q.DefaultChunk }},
		{"from_route_only", 10, func(q PbsQueue) any { return q.FromRouteOnly }},
		{"kill_delay", 10, func(q PbsQueue) any { return q.KillDelay }},
		{"max_array_size", 10, func(q PbsQueue) any { return q.MaxArraySize }},
		{"max_group_res", 10, func(q PbsQueue) any { return q.MaxGroupRes }},
		{"max_group_res_soft", 10, func(q PbsQueue) any { return q.MaxGroupResSoft }},
		{"max_group_run", 10, func(q PbsQueue) any { return q.MaxGroupRun }},
		{"max_group_run_soft", 10, func(q PbsQueue) any { return q.MaxGroupRunSoft }},
		{"max_queuable", 10, func(q PbsQueue) any { return q.MaxQueuable }},
		{"max_queued", 10, func(q PbsQueue) any { return q.MaxQueued }},
		{"max_queued_res", 10, func(q PbsQueue) any { return q.MaxQueuedRes }},
		{"max_run", 10, func(q PbsQueue) any { return q.MaxRun }},
		{"max_run_res", 10, func(q PbsQueue) any { return q.MaxRunRes }},
		{"max_run_res_soft", 10, func(q PbsQueue) any { return q.MaxRunResSoft }},
		{"max_run_soft", 10, func(q PbsQueue) any { return q.MaxRunSoft }},
		{"max_running", 10, func(q PbsQueue) any { return q.MaxRunning }},
		{"max_user_res", 10, func(q PbsQueue) any { return q.MaxUserRes }},
		{"max_user_res_soft", 10, func(q PbsQueue) any { return q.MaxUserResSoft }},
		{"max_user_run", 10, func(q PbsQueue) any { return q.MaxUserRun }},
		{"max_user_run_soft", 10, func(q PbsQueue) any { return q.MaxUserRunSoft }},
		{"node_group_key", 10, func(q PbsQueue) any { return q.NodeGroupKey }},
		{"partition", 10, func(q PbsQueue) any { return q.Partition }},
		{"priority", 10, func(q PbsQueue) any { return q.Priority }},
		{"queued_jobs_threshold", 10, func(q PbsQueue) any { return q.QueuedJobsThreshold }},
		{"queued_jobs_threshold_res", 10, func(q PbsQueue) any { return q.QueuedJobsThresholdRes }},
		{"queue_type", 10, func(q PbsQueue) any {
			if q.QueueType != "" {
				return &q.QueueType
			}
			return (*string)(nil)
		}},
		{"resources_assigned", 10, func(q PbsQueue) any { return q.ResourcesAssigned }},
		{"resources_available", 10, func(q PbsQueue) any { return q.ResourcesAvailable }},
		{"resources_default", 10, func(q PbsQueue) any { return q.ResourcesDefault }},
		{"resources_max", 10, func(q PbsQueue) any { return q.ResourcesMax }},
		{"resources_min", 10, func(q PbsQueue) any { return q.ResourcesMin }},
		{"route_destinations", 10, func(q PbsQueue) any { return q.RouteDestinations }},
		{"route_held_jobs", 10, func(q PbsQueue) any { return q.RouteHeldJobs }},
		{"route_lifetime", 10, func(q PbsQueue) any { return q.RouteLifetime }},
		{"route_retry_time", 10, func(q PbsQueue) any { return q.RouteRetryTime }},
		{"route_waiting_jobs", 10, func(q PbsQueue) any { return q.RouteWaitingJobs }},
		// Enable the queue second to last (order 90)
		{"enabled", 90, func(q PbsQueue) any { return &q.Enabled }},
		// Start the queue last (order 100) - this must happen after all other configuration
		{"started", 100, func(q PbsQueue) any { return &q.Started }},
	}
}

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

	// Get field definitions and sort by order
	fieldDefs := getQueueFieldDefinitions()
	sort.Slice(fieldDefs, func(i, j int) bool {
		return fieldDefs[i].order < fieldDefs[j].order
	})

	// Process fields in order
	for _, fieldDef := range fieldDefs {
		oldValue := fieldDef.getValue(oldQueue)
		newValue := fieldDef.getValue(newQueue)
		newCommands, err := generateUpdateAttributeCommand(oldValue, newValue, "queue", newQueue.Name, fieldDef.attribute)
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
	}

	// Get field definitions and sort by order
	fieldDefs := getQueueFieldDefinitions()
	sort.Slice(fieldDefs, func(i, j int) bool {
		return fieldDefs[i].order < fieldDefs[j].order
	})

	// Process fields in order
	for _, fieldDef := range fieldDefs {
		value := fieldDef.getValue(newQueue)
		c, err := generateCreateCommands(value, "queue", newQueue.Name, fieldDef.attribute)
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
