package pbsclient

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// hookFieldDefinition represents a hook field with its attribute name, execution order, and value extractor.
type hookFieldDefinition struct {
	attribute string
	order     int                    // Lower numbers execute first
	getValue  func(hook PbsHook) any // Function to extract the value from a PbsHook
}

// getHookFieldDefinitions returns the ordered list of hook field definitions.
// This ensures consistent ordering across create and update operations.
func getHookFieldDefinitions() []hookFieldDefinition {
	return []hookFieldDefinition{
		{"alarm", 10, func(h PbsHook) any { return h.Alarm }},
		{"debug", 10, func(h PbsHook) any { return h.Debug }},
		{"event", 10, func(h PbsHook) any { return h.Event }},
		{"fail_action", 10, func(h PbsHook) any { return h.FailAction }},
		{"freq", 10, func(h PbsHook) any { return h.Freq }},
		{"order", 10, func(h PbsHook) any { return h.Order }},
		{"type", 10, func(h PbsHook) any { return h.Type }},
		{"user", 10, func(h PbsHook) any { return h.User }},
		// Enable the hook last - similar to queue pattern
		{"enabled", 90, func(h PbsHook) any { return h.Enabled }},
	}
}

type PbsHook struct {
	Alarm      *int32
	Debug      *bool
	Enabled    *bool
	Event      *string
	FailAction *string
	Freq       *int32
	Name       string
	Order      *int32
	Type       *string
	User       *string
}

func parseHookOutput(output []byte) ([]PbsHook, error) {
	parsedOutput := parseGenericQmgrOutput(string(output))
	var hooks []PbsHook

	for _, r := range parsedOutput {
		if r.objType == "Hook" {
			current := PbsHook{
				Name: r.name,
			}

			for k, v := range r.attributes {
				if s, ok := v.(string); ok {
					switch strings.ToLower(k) {
					case "alarm":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert alarm value to int %s", err.Error())
						}
						i32Value := int32(intValue)
						current.Alarm = &i32Value

					case "debug":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert debug value to bool %s", err.Error())
						}
						current.Debug = &boolValue
					case "enabled":
						boolValue, err := strconv.ParseBool(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert enabled value to bool %s", err.Error())
						}
						current.Enabled = &boolValue
					case "event":
						current.Event = &s
					case "fail_action":
						current.FailAction = &s
					case "freq":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert freq value to int %s", err.Error())
						}
						i32Value := int32(intValue)
						current.Freq = &i32Value
					case "order":
						intValue, err := strconv.Atoi(s)
						if err != nil {
							return nil, fmt.Errorf("failed to convert order value to int %s", err.Error())
						}
						i32Value := int32(intValue)
						current.Order = &i32Value
					case "type":
						current.Type = &s
					case "user":
						current.User = &s
					}
				} else {
					return nil, fmt.Errorf("WHAT IS THIS TYPE %T", v)
				}
			}

			hooks = append(hooks, current)
		}
	}

	return hooks, nil
}

func (c *PbsClient) GetHook(name string) (PbsHook, error) {
	all, err := c.GetHooks()
	if err != nil {
		return PbsHook{}, err
	}

	for _, r := range all {
		if r.Name == name {
			return r, nil
		}
	}

	return PbsHook{}, nil
}

func (c *PbsClient) GetHooks() ([]PbsHook, error) {
	out, errOutput, err := c.runCommand("/opt/pbs/bin/qmgr -c 'list hook @default'")
	if err != nil {
		return nil, fmt.Errorf("%s %s", err, errOutput)
	}

	return parseHookOutput(out)
}

func (c *PbsClient) CreateHook(newHook PbsHook) (PbsHook, error) {
	var commands = []string{
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create hook %s'", newHook.Name),
	}

	// Get field definitions and sort by order
	fieldDefs := getHookFieldDefinitions()
	sort.Slice(fieldDefs, func(i, j int) bool {
		return fieldDefs[i].order < fieldDefs[j].order
	})

	// Process fields in order
	for _, fieldDef := range fieldDefs {
		value := fieldDef.getValue(newHook)
		c, err := generateCreateCommands(value, "hook", newHook.Name, fieldDef.attribute)
		if err != nil {
			return PbsHook{}, err
		}
		commands = append(commands, c...)
	}

	_, errOutput, err := c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return PbsHook{}, fmt.Errorf("%s %s %s", err, completeErrOutput, strings.Join(commands, ","))
	}

	return c.GetHook(newHook.Name)
}

func (c *PbsClient) UpdateHook(newHook PbsHook) (PbsHook, error) {
	oldHook, err := c.GetHook(newHook.Name)
	if err != nil {
		return oldHook, err
	}

	var commands = []string{}

	// Get field definitions and sort by order
	fieldDefs := getHookFieldDefinitions()
	sort.Slice(fieldDefs, func(i, j int) bool {
		return fieldDefs[i].order < fieldDefs[j].order
	})

	// Process fields in order
	for _, fieldDef := range fieldDefs {
		oldValue := fieldDef.getValue(oldHook)
		newValue := fieldDef.getValue(newHook)
		newCommands, err := generateUpdateAttributeCommand(oldValue, newValue, "hook", newHook.Name, fieldDef.attribute)
		if err != nil {
			return oldHook, err
		}
		commands = append(commands, newCommands...)
	}

	_, errOutput, err := c.runCommands(commands) // TODO - Reject bad chars to avoid command injection
	if err != nil {
		completeErrOutput := ""
		for _, e := range errOutput {
			completeErrOutput += string(e)
		}
		return oldHook, fmt.Errorf("%s %s %s", err, completeErrOutput, strings.Join(commands, ","))
	}

	return c.GetHook(oldHook.Name)
}

func (c *PbsClient) DeleteHook(name string) error {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete hook %s'", name)
	_, errOutput, err := c.runCommand(cmd)
	if err != nil {
		return fmt.Errorf("%s %s", err, errOutput)
	}

	return nil
}
