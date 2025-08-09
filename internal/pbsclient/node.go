package pbsclient

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// nodeFieldDefinition represents a node field with its attribute name, execution order, and value extractor
type nodeFieldDefinition struct {
	attribute string
	order     int                    // Lower numbers execute first
	getValue  func(node PbsNode) any // Function to extract the value from a PbsNode
}

// getNodeFieldDefinitions returns the ordered list of node field definitions
// This ensures consistent ordering across create and update operations
func getNodeFieldDefinitions() []nodeFieldDefinition {
	return []nodeFieldDefinition{
		{"comment", 10, func(n PbsNode) any { return n.Comment }},
		{"current_aoe", 10, func(n PbsNode) any { return n.CurrentAoe }},
		{"current_eoe", 10, func(n PbsNode) any { return n.CurrentEoe }},
		{"in_multi_node_host", 10, func(n PbsNode) any { return n.InMultiNodeHost }},
		{"jobs", 10, func(n PbsNode) any { return n.Jobs }},
		{"no_multinode_jobs", 10, func(n PbsNode) any { return n.NoMultinodeJobs }},
		{"partition", 10, func(n PbsNode) any { return n.Partition }},
		{"p_names", 10, func(n PbsNode) any { return n.PNames }},
		{"poweroff_eligible", 10, func(n PbsNode) any { return n.PowerOffEligible }},
		{"power_provisioning", 10, func(n PbsNode) any { return n.PowerProvisioning }},
		{"priority", 10, func(n PbsNode) any { return n.Priority }},
		{"provision_enable", 10, func(n PbsNode) any { return n.ProvisionEnable }},
		{"queue", 10, func(n PbsNode) any { return n.Queue }},
		{"resources_available", 10, func(n PbsNode) any { return n.ResourcesAvailable }},
		{"resv", 10, func(n PbsNode) any { return n.Resv }},
		{"resv_enable", 10, func(n PbsNode) any { return n.ResvEnable }},
	}
}

type PbsNode struct {
	Comment             *string
	CurrentAoe          *string
	CurrentEoe          *string
	InMultiNodeHost     *int32
	Jobs                *string
	LastStateChangeTime *int32
	LastUsedTime        *int32
	License             *string
	LicenseInfo         *int32
	LicType             *string
	MaintenanceJobs     *string
	Mom                 *string // Requires replace
	Name                string
	NoMultinodeJobs     *bool
	NType               *string
	Partition           *string
	PbsVersion          *string
	PCpus               *int32
	PNames              *string
	Port                *int32 // Requires replace
	PowerOffEligible    *bool
	PowerProvisioning   *bool
	Priority            *int32
	ProvisionEnable     *bool
	Queue               *string
	ResourcesAssigned   map[string]string
	ResourcesAvailable  map[string]string
	Resv                *string
	ResvEnable          *bool
	Sharing             *string
	State               *string
}

func parseNodeOutput(output []byte) ([]PbsNode, error) {
	parsedOutput := parseGenericQmgrOutput(string(output))
	var nodes []PbsNode

	for _, r := range parsedOutput {
		if r.objType == "Node" {
			current := PbsNode{
				Name: r.name,
			}

			for k, v := range r.attributes {
				if s, ok := v.(string); ok {
					switch strings.ToLower(k) {
					case "comment":
						current.Comment = &s
					case "current_aoe":
						current.CurrentAoe = &s
					case "current_eoe":
						current.CurrentEoe = &s
					case "in_multinode_host":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						i32Value := int32(intValue)
						current.InMultiNodeHost = &i32Value
					case "jobs":
						current.Jobs = &s
					case "last_state_change_time":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						i32Value := int32(intValue)
						current.LastStateChangeTime = &i32Value
					case "last_used_time":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						i32Value := int32(intValue)
						current.LastUsedTime = &i32Value
					case "license":
						current.License = &s
					case "license_info":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						i32Value := int32(intValue)
						current.LicenseInfo = &i32Value
					case "lic_type":
						current.LicType = &s
					case "maintenance_jobs":
						current.MaintenanceJobs = &s
					case "Mom":
						current.Mom = &s
					case "no_multinode_jobs":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.NoMultinodeJobs = &boolValue
					case "ntype":
						current.NType = &s
					case "partition":
						current.Partition = &s
					case "pbs_version":
						current.PbsVersion = &s
					case "pcpus":
						intValue, err := strconv.Atoi(s)
						if err != nil {

							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						i32Value := int32(intValue)
						current.PCpus = &i32Value
					case "pnames":
						current.PNames = &s
					case "port":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						i32Value := int32(intValue)
						current.Port = &i32Value
					case "poweroff_eligible":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {

							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.PowerOffEligible = &boolValue
					case "power_provisioning":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.PowerProvisioning = &boolValue
					case "priority":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to int %s", k, err.Error())
						}
						i32Value := int32(intValue)
						current.Priority = &i32Value
					case "provision_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.ProvisionEnable = &boolValue
					case "queue":
						current.Queue = &s
					case "resv":
						current.Resv = &s
					case "resv_enable":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert %s value to bool %s", k, err.Error())
						}
						current.ResvEnable = &boolValue
					case "sharing":
						current.Sharing = &s
					case "state":
						current.State = &s
					default:
						// TODO - Not sure whether we care about unsupported attributes
					}
				} else if a, ok := v.(map[string]string); ok {
					switch strings.ToLower(k) {
					case "resources_assigned":
						current.ResourcesAssigned = a
					case "resources_available":
						current.ResourcesAvailable = a
					default:
						// TODO - Not sure whether we care about unsupported attributes
					}
				}
			}

			nodes = append(nodes, current)
		}
	}

	return nodes, nil
}

func (c *PbsClient) GetNode(name string) (PbsNode, error) {
	all, err := c.GetNodes()
	if err != nil {
		return PbsNode{}, err
	}

	for _, r := range all {
		if r.Name == name {
			return r, nil
		}
	}

	return PbsNode{}, nil
}

func (c *PbsClient) GetNodes() ([]PbsNode, error) {
	out, errOutput, err := c.runCommand("/opt/pbs/bin/qmgr -c 'list node @default'")
	if err != nil {
		// PBS returns an error when checking for a list of nodes if there aren't any.
		if strings.Contains(string(errOutput), "Server has no node list") {
			return []PbsNode{}, nil
		}
		return nil, fmt.Errorf("%s %s %s", err, errOutput, out)
	}

	return parseNodeOutput(out)
}

func (c *PbsClient) CreateNode(newNode PbsNode) (PbsNode, error) {
	var extraSettingsOnBaseCmd string
	if newNode.Mom != nil {
		extraSettingsOnBaseCmd += fmt.Sprintf("mom=%s ", *newNode.Mom)
	}
	if newNode.Port != nil {
		extraSettingsOnBaseCmd += fmt.Sprintf("port=%d ", *newNode.Port)
	}

	var commands = []string{
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create node %s %s'", newNode.Name, extraSettingsOnBaseCmd),
	}

	// Get field definitions and sort by order
	fieldDefs := getNodeFieldDefinitions()
	sort.Slice(fieldDefs, func(i, j int) bool {
		return fieldDefs[i].order < fieldDefs[j].order
	})

	// Process fields in order
	for _, fieldDef := range fieldDefs {
		value := fieldDef.getValue(newNode)
		c, err := generateCreateCommands(value, "node", newNode.Name, fieldDef.attribute)
		if err != nil {
			return PbsNode{}, err
		}
		commands = append(commands, c...)
	}

	_, errOutput, err := c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return PbsNode{}, fmt.Errorf("%s, %s, %s", err, strings.Join(commands, ","), completeErrOutput)
	}

	return c.GetNode(newNode.Name)
}

func (c *PbsClient) UpdateNode(newNode PbsNode) (PbsNode, error) {
	oldNode, err := c.GetNode(newNode.Name)
	if err != nil {
		return oldNode, err
	}

	var commands = []string{}

	// Get field definitions and sort by order
	fieldDefs := getNodeFieldDefinitions()
	sort.Slice(fieldDefs, func(i, j int) bool {
		return fieldDefs[i].order < fieldDefs[j].order
	})

	// Horrible hack because host and vnode are properties in the resources_available map but actually set by the MoM not the user
	delete(oldNode.ResourcesAvailable, "host")
	delete(oldNode.ResourcesAvailable, "vnode")
	delete(newNode.ResourcesAvailable, "host")
	delete(newNode.ResourcesAvailable, "vnode")

	// Process fields in order
	for _, fieldDef := range fieldDefs {
		oldValue := fieldDef.getValue(oldNode)
		newValue := fieldDef.getValue(newNode)
		newCommands, err := generateUpdateAttributeCommand(oldValue, newValue, "node", newNode.Name, fieldDef.attribute)
		if err != nil {
			return oldNode, err
		}
		commands = append(commands, newCommands...)
	}

	_, errOutput, err := c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return oldNode, fmt.Errorf("%s, %s, %s", err, strings.Join(commands, ","), completeErrOutput)
	}

	oldNode, err = c.GetNode(oldNode.Name)
	if err != nil {
		return oldNode, err
	}

	return oldNode, nil
}

func (c *PbsClient) DeleteNode(name string) error {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete node %s'", name)
	_, errOutput, err := c.runCommand(cmd)
	if err != nil {
		return fmt.Errorf("%s %s", err, errOutput)
	}

	return nil
}
