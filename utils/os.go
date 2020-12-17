package utils

import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"errors"
	"strings"
)

type OS struct {
	DistributionID string `hcl:"distributionID"`
	Version	string	`hcl:"version,optional"`
}

//GetRemoteOS - Function to determine the operating system of the remote system
func GetRemoteOS(host string, port string, config *ssh.ClientConfig) (string, error) {
	command := NewCommand("cat", "/etc/*-release")

	endpoint := fmt.Sprintf("%s:%s", host, port)

	client, err := ssh.Dial("tcp", endpoint, config)

	if err != nil {
		return "", errors.New("Could not establish a connection to the specified host. Error: " + err.Error())
	}

	err = command.RunRemote(client)

	if err != nil {
		return "", errors.New("Could not properly run the remote command specified. Error: " + err.Error())
	}

	output := command.Stdout.String()

	details := strings.Split(output, "\n")

	distrib := strings.TrimSpace(strings.Split(details[0], "=")[1])

	version := strings.TrimSpace(strings.Split(details[1], "=")[1])

	output = distrib + ":" + version

	return output, nil
}