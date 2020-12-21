package resources

import (
	"errors"
	"strings"

	"golang.org/x/crypto/ssh"
)

type OS struct {
	DistributionID string `hcl:"distributionID"`
	Version        string `hcl:"version,optional"`
}

//GetRemoteOS - Function to determine the operating system of the remote system
func GetRemoteOS(client *ssh.Client) (string, string, error) {
	command := NewCommand("cat", "/etc/*-release")

	err := command.RunRemote(client)

	if err != nil {
		return "", "", errors.New("Could not properly run the remote command specified. Error: " + err.Error())
	}

	output := command.Stdout.String()

	details := strings.Split(output, "\n")

	distrib := strings.TrimSpace(strings.Split(details[0], "=")[1])

	version := strings.TrimSpace(strings.Split(details[1], "=")[1])

	return distrib, version, nil
}
