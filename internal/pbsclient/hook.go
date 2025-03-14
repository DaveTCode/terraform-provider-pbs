package pbsclient

import (
	"fmt"
	"strconv"
	"strings"
)

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

	fields := []struct {
		attribute string
		newAttr   any
	}{
		{"alarm", newHook.Alarm},
		{"debug", newHook.Debug},
		{"enabled", newHook.Enabled},
		{"event", newHook.Event},
		{"fail_action", newHook.FailAction},
		{"freq", newHook.Freq},
		{"order", newHook.Order},
		{"type", newHook.Type},
		{"user", newHook.User},
	}
	for _, v := range fields {
		c, err := generateCreateCommands(v.newAttr, "hook", newHook.Name, v.attribute)
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
	fields := []struct {
		attribute string
		oldAttr   any
		newAttr   any
	}{
		{"alarm", oldHook.Alarm, newHook.Alarm},
		{"debug", oldHook.Debug, newHook.Debug},
		{"enabled", oldHook.Enabled, newHook.Enabled},
		{"event", oldHook.Event, newHook.Event},
		{"fail_action", oldHook.FailAction, newHook.FailAction},
		{"freq", oldHook.Freq, newHook.Freq},
		{"order", oldHook.Order, newHook.Order},
		{"type", oldHook.Type, newHook.Type},
		{"user", oldHook.User, newHook.User},
	}
	for _, v := range fields {
		newCommands, err := generateUpdateAttributeCommand(v.oldAttr, v.newAttr, "hook", newHook.Name, v.attribute)
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
