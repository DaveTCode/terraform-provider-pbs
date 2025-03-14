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
					case "power_off_eligible":
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

func (c *PbsClient) CreateNode(new PbsNode) (PbsNode, error) {
	var extraSettingsOnBaseCmd string
	if new.Mom != nil {
		extraSettingsOnBaseCmd += fmt.Sprintf("mom=%s ", *new.Mom)
	}
	if new.Port != nil {
		extraSettingsOnBaseCmd += fmt.Sprintf("port=%d ", *new.Port)
	}

	var commands = []string{
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create node %s %s'", new.Name, extraSettingsOnBaseCmd),
	}

	fields := []struct {
		attribute string
		new       any
	}{
		{"comment", new.Comment},
		{"current_aoe", new.CurrentAoe},
		{"current_eoe", new.CurrentEoe},
		{"in_multi_node_host", new.InMultiNodeHost},
		{"jobs", new.Jobs},
		{"no_multinode_jobs", new.NoMultinodeJobs},
		{"partition", new.Partition},
		{"p_names", new.PNames},
		{"power_off_eligible", new.PowerOffEligible},
		{"power_provisioning", new.PowerProvisioning},
		{"priority", new.Priority},
		{"provision_enable", new.ProvisionEnable},
		{"queue", new.Queue},
		{"resources_available", new.ResourcesAvailable},
		{"resv", new.Resv},
		{"resv_enable", new.ResvEnable},
	}
	for _, v := range fields {
		c, err := generateCreateCommands(v.new, "node", new.Name, v.attribute)
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

	return c.GetNode(new.Name)
}

func (c *PbsClient) UpdateNode(new PbsNode) (PbsNode, error) {
	old, err := c.GetNode(new.Name)
	if err != nil {
		return old, err
	}

	var commands = []string{}
	fields := []struct {
		attribute string
		old       any
		new       any
	}{
		{"comment", old.Comment, new.Comment},
		{"current_aoe", old.CurrentAoe, new.CurrentAoe},
		{"current_eoe", old.CurrentEoe, new.CurrentEoe},
		{"in_multi_node_host", old.InMultiNodeHost, new.InMultiNodeHost},
		{"no_multinode_jobs", old.NoMultinodeJobs, new.NoMultinodeJobs},
		{"partition", old.Partition, new.Partition},
		{"p_names", old.PNames, new.PNames},
		{"power_off_eligible", old.PowerOffEligible, new.PowerOffEligible},
		{"power_provisioning", old.PowerProvisioning, new.PowerProvisioning},
		{"priority", old.Priority, new.Priority},
		{"provision_enable", old.ProvisionEnable, new.ProvisionEnable},
		{"queue", old.Queue, new.Queue},
		{"resources_available", old.ResourcesAvailable, new.ResourcesAvailable},
		{"resv_enable", old.ResvEnable, new.ResvEnable},
	}
	for _, v := range fields {
		switch v.old.(type) {
		case *bool:
			commands = append(commands, generateUpdateBoolAttributeCommand("node", new.Name, v.attribute, v.old.(*bool), v.new.(*bool))...)
		case *int32:
			commands = append(commands, generateUpdateInt32AttributeCommand("node", new.Name, v.attribute, v.old.(*int32), v.new.(*int32))...)
		case *string:
			commands = append(commands, generateUpdateStringAttributeCommand("node", new.Name, v.attribute, v.old.(*string), v.new.(*string))...)
		case map[string]string:
			oldValue := v.old.(map[string]string)
			newValue := v.new.(map[string]string)
			for k, oldAttrVal := range oldValue {
				// Special case because host/vnode are set by mom on the node not by the user
				if k == "host" || k == "vnode" {
					continue
				}
				newAttrVal, ok := newValue[k]
				if !ok {
					commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset node %s %s.%s'", new.Name, v.attribute, k))
				} else if oldAttrVal != newAttrVal {
					commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set node %s %s.%s=%s'", new.Name, v.attribute, k, newAttrVal))
				}
			}
			for k, newAttrVal := range newValue {
				if _, ok := oldValue[k]; !ok {
					commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set node %s %s.%s=%s'", new.Name, v.attribute, k, newAttrVal))
				}
			}
		default:
			return old, fmt.Errorf("unsupported type %T", v.old)
		}
	}

	_, errOutput, err := c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return old, fmt.Errorf("%s, %s, %s", err, strings.Join(commands, ","), completeErrOutput)
	}

	old, err = c.GetNode(old.Name)
	if err != nil {
		return old, err
	}

	return old, nil
}

func (c *PbsClient) DeleteNode(name string) error {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete node %s'", name)
	_, errOutput, err := c.runCommand(cmd)
	if err != nil {
		return fmt.Errorf("%s %s", err, errOutput)
	}

	return nil
}
