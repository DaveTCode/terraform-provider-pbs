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
				}

				hooks = append(hooks, current)
			}
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
	out, err := c.runCommand("/opt/pbs/bin/qmgr -c 'list hook @default'")
	if err != nil {
		return nil, err
	}

	return parseHookOutput(out)
}

func (c *PbsClient) CreateHook(new PbsHook) (PbsHook, error) {
	var commands = []string{
		fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create hook %s", new.Name),
	}

	fields := []struct {
		attribute string
		new       any
	}{
		{"alarm", new.Alarm},
		{"debug", new.Debug},
		{"enabled", new.Enabled},
		{"event", new.Event},
		{"fail_action", new.FailAction},
		{"freq", new.Freq},
		{"order", new.Order},
		{"type", new.Type},
		{"user", new.User},
	}
	for _, v := range fields {
		command := ""
		switch v.new.(type) {
		case *bool:
			b := v.new.(*bool)
			if b != nil {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set hook %s %s=%s'", new.Name, v.attribute, strconv.FormatBool(*b))
			}
		case *int32:
			i := v.new.(*int32)
			if i != nil {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set hook %s %s=%s'", new.Name, v.attribute, strconv.Itoa(int(*i)))
			}
		case *string:
			s := v.new.(*string)
			if s != nil {
				command = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set hook %s %s=%s'", new.Name, v.attribute, *s)
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
		return PbsHook{}, err
	}

	return c.GetHook(new.Name)
}

func (c *PbsClient) UpdateHook(new PbsHook) (PbsHook, error) {
	old, err := c.GetHook(new.Name)
	if err != nil {
		return old, err
	}

	var commands = []string{}
	fields := []struct {
		attribute string
		old       any
		new       any
	}{
		{"alarm", old.Alarm, new.Alarm},
		{"debug", old.Debug, new.Debug},
		{"enabled", old.Enabled, new.Enabled},
		{"event", old.Event, new.Event},
		{"fail_action", old.FailAction, new.FailAction},
		{"freq", old.Freq, new.Freq},
		{"order", old.Order, new.Order},
		{"type", old.Type, new.Type},
		{"user", old.User, new.User},
	}
	for _, v := range fields {
		command := ""
		switch v.old.(type) {
		case *bool:
			command = generateUpdateBoolAttributeCommand("hook", new.Name, v.attribute, v.old.(*bool), v.new.(*bool))
		case *int32:
			command = generateUpdateInt32AttributeCommand("hook", new.Name, v.attribute, v.old.(*int32), v.new.(*int32))
		case *string:
			command = generateUpdateStringAttributeCommand("hook", new.Name, v.attribute, v.old.(*string), v.new.(*string))
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

	old, err = c.GetHook(old.Name)
	if err != nil {
		return old, err
	}

	return old, nil
}

func (c *PbsClient) DeleteHook(name string) error {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete hook %s'", name)
	_, err := c.runCommand(cmd)
	if err != nil {
		return err
	}

	return nil
}
