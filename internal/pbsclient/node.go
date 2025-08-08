package pbsclient

import (
	"fmt"
	"strconv"
	"strings"
)

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

	fields := []struct {
		attribute string
		newAttr   any
	}{
		{"comment", newNode.Comment},
		{"current_aoe", newNode.CurrentAoe},
		{"current_eoe", newNode.CurrentEoe},
		{"in_multi_node_host", newNode.InMultiNodeHost},
		{"jobs", newNode.Jobs},
		{"no_multinode_jobs", newNode.NoMultinodeJobs},
		{"partition", newNode.Partition},
		{"p_names", newNode.PNames},
		{"poweroff_eligible", newNode.PowerOffEligible},
		{"power_provisioning", newNode.PowerProvisioning},
		{"priority", newNode.Priority},
		{"provision_enable", newNode.ProvisionEnable},
		{"queue", newNode.Queue},
		{"resources_available", newNode.ResourcesAvailable},
		{"resv", newNode.Resv},
		{"resv_enable", newNode.ResvEnable},
	}
	for _, v := range fields {
		c, err := generateCreateCommands(v.newAttr, "node", newNode.Name, v.attribute)
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
	fields := []struct {
		attribute string
		oldAttr   any
		newAttr   any
	}{
		{"comment", oldNode.Comment, newNode.Comment},
		{"current_aoe", oldNode.CurrentAoe, newNode.CurrentAoe},
		{"current_eoe", oldNode.CurrentEoe, newNode.CurrentEoe},
		{"in_multi_node_host", oldNode.InMultiNodeHost, newNode.InMultiNodeHost},
		{"no_multinode_jobs", oldNode.NoMultinodeJobs, newNode.NoMultinodeJobs},
		{"partition", oldNode.Partition, newNode.Partition},
		{"p_names", oldNode.PNames, newNode.PNames},
		{"poweroff_eligible", oldNode.PowerOffEligible, newNode.PowerOffEligible},
		{"power_provisioning", oldNode.PowerProvisioning, newNode.PowerProvisioning},
		{"priority", oldNode.Priority, newNode.Priority},
		{"provision_enable", oldNode.ProvisionEnable, newNode.ProvisionEnable},
		{"queue", oldNode.Queue, newNode.Queue},
		{"resources_available", oldNode.ResourcesAvailable, newNode.ResourcesAvailable},
		{"resv_enable", oldNode.ResvEnable, newNode.ResvEnable},
	}

	// Horrible hack because host and vnode are properties in the resources_available map but actually set by the MoM not the user
	delete(oldNode.ResourcesAvailable, "host")
	delete(oldNode.ResourcesAvailable, "vnode")
	delete(newNode.ResourcesAvailable, "host")
	delete(newNode.ResourcesAvailable, "vnode")
	for _, v := range fields {
		newCommands, err := generateUpdateAttributeCommand(v.oldAttr, v.newAttr, "node", newNode.Name, v.attribute)
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
