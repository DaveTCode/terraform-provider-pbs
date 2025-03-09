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
	ToplogyInfo         *string
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
					case "toplogy_info":
						current.ToplogyInfo = &s
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
	out, err := c.runCommand("/opt/pbs/bin/qmgr -c 'list node @default'")
	if err != nil {
		return nil, err
	}

	return parseNodeOutput(out)
}

func (c *PbsClient) CreateNode(new PbsNode) (PbsNode, error) {
	var commands = []string{
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create node %s", new.Name),
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
		{"last_state_change_time", new.LastStateChangeTime},
		{"last_used_time", new.LastUsedTime},
		{"license", new.License},
		{"license_info", new.LicenseInfo},
		{"lic_type", new.LicType},
		{"maintenance_jobs", new.MaintenanceJobs},
		{"mom", new.Mom},
		{"no_multinode_jobs", new.NoMultinodeJobs},
		{"n_type", new.NType},
		{"partition", new.Partition},
		{"pbs_version", new.PbsVersion},
		{"p_cpus", new.PCpus},
		{"p_names", new.PNames},
		{"port", new.Port},
		{"power_off_eligible", new.PowerOffEligible},
		{"power_provisioning", new.PowerProvisioning},
		{"priority", new.Priority},
		{"provision_enable", new.ProvisionEnable},
		{"queue", new.Queue},
		{"resources_assigned", new.ResourcesAssigned},
		{"resources_available", new.ResourcesAvailable},
		{"resv", new.Resv},
		{"resv_enable", new.ResvEnable},
		{"sharing", new.Sharing},
		{"state", new.State},
		{"toplogy_info", new.ToplogyInfo},
	}
	for _, v := range fields {
		command := ""
		switch v.new.(type) {
		case *bool:
			b := v.new.(*bool)
			if b != nil {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set node %s %s=%s'", new.Name, v.attribute, strconv.FormatBool(*b))
			}
		case *int32:
			i := v.new.(*int32)
			if i != nil {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set node %s %s=%s'", new.Name, v.attribute, strconv.Itoa(int(*i)))
			}
		case *string:
			s := v.new.(*string)
			if s != nil {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set node %s %s=%s'", new.Name, v.attribute, *s)
			}
		default:
			return new, fmt.Errorf("unsupported type %T", v.new)
		}

		if command != "" {
			commands = append(commands, command)
		}
	}

	_, err := c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		return PbsNode{}, err
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
		{"jobs", old.Jobs, new.Jobs},
		{"last_state_change_time", old.LastStateChangeTime, new.LastStateChangeTime},
		{"last_used_time", old.LastUsedTime, new.LastUsedTime},
		{"license", old.License, new.License},
		{"license_info", old.LicenseInfo, new.LicenseInfo},
		{"lic_type", old.LicType, new.LicType},
		{"maintenance_jobs", old.MaintenanceJobs, new.MaintenanceJobs},
		{"mom", old.Mom, new.Mom},
		{"no_multinode_jobs", old.NoMultinodeJobs, new.NoMultinodeJobs},
		{"n_type", old.NType, new.NType},
		{"partition", old.Partition, new.Partition},
		{"pbs_version", old.PbsVersion, new.PbsVersion},
		{"p_cpus", old.PCpus, new.PCpus},
		{"p_names", old.PNames, new.PNames},
		{"port", old.Port, new.Port},
		{"power_off_eligible", old.PowerOffEligible, new.PowerOffEligible},
		{"power_provisioning", old.PowerProvisioning, new.PowerProvisioning},
		{"priority", old.Priority, new.Priority},
		{"provision_enable", old.ProvisionEnable, new.ProvisionEnable},
		{"queue", old.Queue, new.Queue},
		{"resources_assigned", old.ResourcesAssigned, new.ResourcesAssigned},
		{"resources_available", old.ResourcesAvailable, new.ResourcesAvailable},
		{"resv", old.Resv, new.Resv},
		{"resv_enable", old.ResvEnable, new.ResvEnable},
		{"sharing", old.Sharing, new.Sharing},
		{"state", old.State, new.State},
		{"toplogy_info", old.ToplogyInfo, new.ToplogyInfo},
	}
	for _, v := range fields {
		command := ""
		switch v.old.(type) {
		case *bool:
			command = generateUpdateBoolAttributeCommand("node", new.Name, v.attribute, v.old.(*bool), v.new.(*bool))
		case *int32:
			command = generateUpdateInt32AttributeCommand("node", new.Name, v.attribute, v.old.(*int32), v.new.(*int32))
		case *string:
			command = generateUpdateStringAttributeCommand("node", new.Name, v.attribute, v.old.(*string), v.new.(*string))
		default:
			return old, fmt.Errorf("unsupported type %T", v.old)
		}

		if command != "" {
			commands = append(commands, command)
		}

	}

	_, err = c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		return old, err
	}

	old, err = c.GetNode(old.Name)
	if err != nil {
		return old, err
	}

	return old, nil
}

func (c *PbsClient) DeleteNode(name string) error {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete node %s'", name)
	_, err := c.runCommand(cmd)
	if err != nil {
		return err
	}

	return nil
}
