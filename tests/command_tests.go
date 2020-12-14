package tests

import (
	"errors"
	"fmt"
	"time"

	"github.com/everettraven/hades/utils"
	"golang.org/x/crypto/ssh"
)

// TestLocalCommand - Function for testing execution of a local command
func TestLocalCommand(name string, args ...string) *utils.Command {
	command := utils.NewCommand(name, args...)

	err := command.RunLocal()
	if err != nil {
		fmt.Println("\tFailed")
	}

	fmt.Println("\tPassed")

	return command
}

// TestRemoteCommand - Function for testing execution of a remote command
func TestRemoteCommand(host string, port string, name string, args ...string) (*utils.Command, error) {
	command := utils.NewCommand(name, args...)

	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("root")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	endpoint := fmt.Sprintf("%s:%s", host, port)

	client, err := ssh.Dial("tcp", endpoint, config)

	if err != nil {
		return command, errors.New("Could not establish a connection to the specified host. Error: " + err.Error())
	}

	err = command.RunRemote(client)

	if err != nil {
		return command, errors.New("Could not properly run the remote command specified. Error: " + err.Error())
	}

	return command, nil
}
