package pbsclient

import (
	"fmt"
	"io"
	"strconv"

	"golang.org/x/crypto/ssh"
)

type PbsClient struct {
	SshClientConfig *ssh.ClientConfig
	Address         string
}

func runSshCommand(sshClient *ssh.Client, cmd string) ([]byte, error) {
	session, err := sshClient.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create SSH session %s", err.Error())
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to attach stdout pipe so cannot process results %s", err.Error())
	}

	if err := session.Start(cmd); err != nil {
		return nil, fmt.Errorf("failed to create command to run qmgr command to get queues %s", err.Error())
	}

	cmdOutput, err := io.ReadAll(stdout)
	if err != nil {
		return nil, fmt.Errorf("failed to read text from stdout %s", err.Error())
	}
	if err := session.Wait(); err != nil {
		return nil, fmt.Errorf("failed to execute command against PBS server %s", err.Error())
	}

	return cmdOutput, nil
}

func (client *PbsClient) runCommands(commands []string) ([][]byte, error) {
	sshClient, err := ssh.Dial("tcp", client.Address, client.SshClientConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to server with SSH config provided %s", err.Error())
	}
	defer sshClient.Close()

	var output [][]byte
	for _, cmd := range commands {
		cmdOutput, err := runSshCommand(sshClient, cmd)
		if err != nil {
			return nil, err
		}
		output = append(output, cmdOutput)
	}

	return output, nil
}

func (client *PbsClient) runCommand(cmd string) ([]byte, error) {
	output, err := client.runCommands([]string{cmd})
	if err != nil {
		return nil, err
	}
	if len(output) != 1 {
		return nil, fmt.Errorf("expected 1 output from running ssh command but got %d", len(output))
	}

	return output[0], nil
}

func generateUpdateBoolAttributeCommand(obj string, name string, attribute string, oldValue *bool, newValue *bool) string {
	if oldValue != newValue {
		if newValue != nil {
			return fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", obj, name, attribute, strconv.FormatBool(*newValue))
		} else {
			return fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset %s %s %s'", obj, name, attribute)
		}
	}

	return ""
}

func generateUpdateInt32AttributeCommand(obj string, name string, attribute string, oldValue *int32, newValue *int32) string {
	if oldValue != newValue {
		if newValue != nil {
			return fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", obj, name, attribute, strconv.Itoa(int(*newValue)))
		} else {
			return fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset %s %s %s'", obj, name, attribute)
		}
	}

	return ""
}

func generateUpdateStringAttributeCommand(obj string, name string, attribute string, oldValue *string, newValue *string) string {
	if oldValue != newValue {
		if newValue != nil {
			return fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set %s %s %s=%s'", obj, name, attribute, *newValue)
		} else {
			return fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset %s %s %s'", obj, name, attribute)
		}
	}

	return ""
}
