package command

import (
	"errors"
	"fmt"
	"time"

	"github.com/everettraven/hades/resources"
	"github.com/everettraven/hades/utils"
	"golang.org/x/crypto/ssh"
)

// TestLocalCommand - Function for testing execution of a local command
func TestLocalCommand(name string, args ...string) *resources.Command {
	command := resources.NewCommand(name, args...)

	err := command.RunLocal()
	if err != nil {
		fmt.Println("\tFailed")
	}

	fmt.Println("\tPassed")

	return command
}

// TestRemoteCommand - Function for testing execution of a remote command
func TestRemoteCommand(host string, port string, name string, args ...string) (*resources.Command, error) {
	command := resources.NewCommand(name, args...)

	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("root")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	client, err := utils.GetSSHClient(host, port, config)

	if err != nil {
		return command, err
	}

	err = command.RunRemote(client)

	if err != nil {
		return command, errors.New("Could not properly run the remote command specified. Error: " + err.Error())
	}

	return command, nil
}
