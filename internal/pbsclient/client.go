package pbsclient

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

type PbsClient struct {
	SshClientConfig *ssh.ClientConfig
	Address         string
}

func runSshCommand(sshClient *ssh.Client, cmd string) ([]byte, []byte, error) {
	session, err := sshClient.NewSession()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create SSH session %s: %s", err.Error(), cmd)
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to attach stdout pipe so cannot process results %s: %s", err.Error(), cmd)
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to attach stdout pipe so cannot process results %s: %s", err.Error(), cmd)
	}

	if err := session.Start(cmd); err != nil {
		return nil, nil, fmt.Errorf("failed to create command %s: %s", err.Error(), cmd)
	}

	cmdOutput, err := io.ReadAll(stdout)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read text from stdout %s: %s", err.Error(), cmd)
	}
	stdErrOutput, err := io.ReadAll(stderr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read text from stderr %s: %s", err.Error(), cmd)
	}
	if err := session.Wait(); err != nil {
		return cmdOutput, stdErrOutput, fmt.Errorf("failed to execute command against PBS server %s: %s", err.Error(), cmd)
	}

	return cmdOutput, stdErrOutput, nil
}

func (client *PbsClient) runCommands(commands []string) ([][]byte, [][]byte, error) {
	sshClient, err := ssh.Dial("tcp", client.Address, client.SshClientConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to connect to server with SSH config provided %s", err.Error())
	}
	defer sshClient.Close()

	var output [][]byte
	var errOutput [][]byte
	for _, cmd := range commands {
		cmdOutput, stdErrOutput, err := runSshCommand(sshClient, cmd)
		output = append(output, cmdOutput)
		errOutput = append(errOutput, stdErrOutput)
		if err != nil {
			return output, errOutput, err
		}
	}

	return output, errOutput, nil
}

func (client *PbsClient) runCommand(cmd string) ([]byte, []byte, error) {
	output, errOutput, err := client.runCommands([]string{cmd})
	if err != nil {
		maybeOutput := []byte{}
		if len(output) > 0 {
			maybeOutput = output[0]
		}
		maybeErrOutput := []byte{}
		if len(errOutput) > 0 {
			maybeErrOutput = errOutput[0]
		}
		return maybeOutput, maybeErrOutput, err
	}

	return output[0], errOutput[0], nil
}

func generateUpdateBoolAttributeCommand(obj string, name string, attribute string, oldValue *bool, newValue *bool) []string {
	if oldValue == nil {
		if newValue != nil {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", obj, name, attribute, strconv.FormatBool(*newValue))}
		}
	} else {
		if newValue == nil {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset %s %s %s'", obj, name, attribute)}
		} else if *oldValue != *newValue {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", obj, name, attribute, strconv.FormatBool(*newValue))}
		}
	}

	return []string{}
}

func generateUpdateInt32AttributeCommand(obj string, name string, attribute string, oldValue *int32, newValue *int32) []string {
	if oldValue == nil {
		if newValue != nil {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", obj, name, attribute, strconv.Itoa(int(*newValue)))}
		}
	} else {
		if newValue == nil {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset %s %s %s'", obj, name, attribute)}
		} else if *oldValue != *newValue {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", obj, name, attribute, strconv.Itoa(int(*newValue)))}
		}
	}

	return []string{}
}

func generateUpdateInt64AttributeCommand(obj string, name string, attribute string, oldValue *int64, newValue *int64) []string {
	if oldValue == nil {
		if newValue != nil {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", obj, name, attribute, strconv.FormatInt(*newValue, 10))}
		}
	} else {
		if newValue == nil {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset %s %s %s'", obj, name, attribute)}
		} else if *oldValue != *newValue {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", obj, name, attribute, strconv.FormatInt(*newValue, 10))}
		}
	}

	return []string{}
}

func generateUpdateStringAttributeCommand(obj string, name string, attribute string, oldValue *string, newValue *string) []string {
	if oldValue == nil {
		if newValue != nil {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=\"%s\"'", obj, name, attribute, *newValue)}
		}
	} else {
		if newValue == nil {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset %s %s %s'", obj, name, attribute)}
		} else if *oldValue != *newValue {
			return []string{fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=\"%s\"'", obj, name, attribute, *newValue)}
		}
	}

	return []string{}
}

var (
	nameRegex              = regexp.MustCompile(`^(\w+)\s+(\w+)$`)
	attributeRegex         = regexp.MustCompile(`^    (\w+)\s+=\s+(.+)$`)
	dotAttributeRegex      = regexp.MustCompile(`^    (\w+)\.(\w+)\s+=\s+(.+)$`)
	continueAttributeRegex = regexp.MustCompile(`^        (.+)$`)
)

type qmgrResult struct {
	objType    string
	name       string
	attributes map[string]any
}

func parseGenericQmgrOutput(output string) []qmgrResult {
	lines := strings.Split(output, "\n")
	results := make([]qmgrResult, 0)
	current := qmgrResult{}
	var prevAttribute string
	for _, line := range lines {
		if nameRegex.MatchString(line) {
			if current.name != "" {
				results = append(results, current)
			}
			current = qmgrResult{
				objType:    nameRegex.FindStringSubmatch(line)[1],
				name:       nameRegex.FindStringSubmatch(line)[2],
				attributes: make(map[string]any, 0),
			}
		} else if attributeRegex.MatchString(line) {
			subMatch := attributeRegex.FindStringSubmatch(line)
			attribute := subMatch[1]
			value := subMatch[2]
			current.attributes[attribute] = value
			prevAttribute = attribute
		} else if dotAttributeRegex.MatchString(line) {
			subMatch := dotAttributeRegex.FindStringSubmatch(line)
			attribute := subMatch[1]
			subAttribute := subMatch[2]
			value := subMatch[3]
			if _, ok := current.attributes[attribute]; ok {
				if attrMap, ok := current.attributes[attribute].(map[string]string); ok {
					attrMap[subAttribute] = value
				} else {
					attrMap := make(map[string]string)
					attrMap[subAttribute] = value
					current.attributes[attribute] = attrMap
				}
			} else {
				attrMap := make(map[string]string)
				attrMap[subAttribute] = value
				current.attributes[attribute] = attrMap
			}
			prevAttribute = attribute
		} else if continueAttributeRegex.MatchString(line) {
			if prevAttribute != "" {
				current.attributes[prevAttribute] = current.attributes[prevAttribute].(string) + continueAttributeRegex.FindStringSubmatch(line)[1]
			}
		}
	}

	if current.name != "" {
		results = append(results, current)
	}

	return results
}

func generateCreateCommands(new any, qmgrObjectType string, qmgrObjectName string, qmgrAttribute string) ([]string, error) {
	commands := []string{}
	switch new := new.(type) {
	case *bool:
		if new != nil {
			commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", qmgrObjectType, qmgrObjectName, qmgrAttribute, strconv.FormatBool(*new)))
		}
	case *int32:
		if new != nil {
			commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", qmgrObjectType, qmgrObjectName, qmgrAttribute, strconv.Itoa(int(*new))))
		}
	case *int64:
		if new != nil {
			commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", qmgrObjectType, qmgrObjectName, qmgrAttribute, strconv.FormatInt(*new, 10)))
		}
	case *string:
		if new != nil {
			commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=\"%s\"'", qmgrObjectType, qmgrObjectName, qmgrAttribute, *new))
		}
	case map[string]string:
		for k, subval := range new {
			commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s.%s=%s'", qmgrObjectType, qmgrObjectName, qmgrAttribute, k, subval))
		}
	default:
		return commands, fmt.Errorf("unsupported type %T", new)
	}

	return commands, nil
}

func generateUpdateAttributeCommand(old any, new any, qmgrObjectType string, qmgrObjectName string, qmgrAttribute string) ([]string, error) {
	switch old := old.(type) {
	case bool:
		newValue := new.(bool)
		return generateUpdateBoolAttributeCommand(qmgrObjectType, qmgrObjectName, qmgrAttribute, &old, &newValue), nil
	case *bool:
		return generateUpdateBoolAttributeCommand(qmgrObjectType, qmgrObjectName, qmgrAttribute, old, new.(*bool)), nil
	case int32:
		newValue := new.(int32)
		return generateUpdateInt32AttributeCommand(qmgrObjectType, qmgrObjectName, qmgrAttribute, &old, &newValue), nil
	case *int32:
		return generateUpdateInt32AttributeCommand(qmgrObjectType, qmgrObjectName, qmgrAttribute, old, new.(*int32)), nil
	case *int64:
		return generateUpdateInt64AttributeCommand(qmgrObjectType, qmgrObjectName, qmgrAttribute, old, new.(*int64)), nil
	case string:
		newValue := new.(string)
		return generateUpdateStringAttributeCommand(qmgrObjectType, qmgrObjectName, qmgrAttribute, &old, &newValue), nil
	case *string:
		return generateUpdateStringAttributeCommand(qmgrObjectType, qmgrObjectName, qmgrAttribute, old, new.(*string)), nil
	case map[string]string:
		newValue := new.(map[string]string)
		commands := []string{}
		for k, oldAttrVal := range old {
			newAttrVal, ok := newValue[k]
			if !ok {
				commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset %s %s %s.%s'", qmgrObjectType, qmgrObjectName, qmgrAttribute, k))

			} else if oldAttrVal != newAttrVal {
				commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s.%s=%s'", qmgrObjectType, qmgrObjectName, qmgrAttribute, k, newAttrVal))
			}
		}
		for k, newAttrVal := range newValue {
			if _, ok := old[k]; !ok {
				commands = append(commands, fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s.%s=%s'", qmgrObjectType, qmgrObjectName, qmgrAttribute, k, newAttrVal))
			}
		}

		return commands, nil

	default:
		return nil, fmt.Errorf("unsupported type %T", old)
	}
}
