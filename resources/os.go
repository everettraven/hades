package resources

import (
	"errors"
	"runtime"
	"strings"

	"golang.org/x/crypto/ssh"
)

//OS - Struct to hold the os block testing information
type OS struct {
	DistributionID string `hcl:"distributionID"`
	Version        string `hcl:"version,optional"`
}

//GetRemoteOS - Function to determine the operating system of the remote system
func GetRemoteOS(client *ssh.Client) (string, string, error) {
	// Get the OS from the os-release file and try to get only the name and version id
	command := NewCommand("cat", "/etc/os-release | grep -w \"NAME\\|VERSION_ID\"")

	// Run the remote command
	err := command.RunRemote(client)

	// Check for errors
	if err != nil {
		command = NewCommand("powershell", "Write-Output \"OSNAME=$(Get-ComputerInfo | % {$_.OsName})\";Write-Output \"OSVERSION=$(Get-ComputerInfo | % {$_.OsVersion})\"")
		err = command.RunRemote(client)

		if err != nil {
			return "", "", errors.New("Error: " + err.Error())
		}
	}

	// Get the output from the command
	output := command.Stdout.String()

	// split on the newline
	details := strings.Split(output, "\n")

	//Split on the equals sign and then trim off the quotes for both the name and version id values
	distrib := strings.TrimSpace(strings.Split(details[0], "=")[1])

	distrib = strings.Trim(distrib, "\"")

	version := strings.TrimSpace(strings.Split(details[1], "=")[1])

	version = strings.Trim(version, "\"")

	return distrib, version, nil
}

//GetLocalOS - Get the local
func GetLocalOS() (string, string, error) {
	var command *Command
	// Get the OS from the os-release file and try to get only the name and version id
	if runtime.GOOS == "windows" {
		command = NewCommand("powershell", "Write-Output \"OSNAME=$(Get-ComputerInfo | % {$_.OsName})\";Write-Output \"OSVERSION=$(Get-ComputerInfo | % {$_.OsVersion})\"")
	} else {
		command = NewCommand("bash", "-c", "cat /etc/os-release | grep -w \"NAME\\|VERSION_ID\"")
	}

	// Run the command
	err := command.RunLocal()

	// Check for errors
	if err != nil {
		return "", "", errors.New("Error: " + err.Error())
	}

	// Get the output from the command
	output := command.Stdout.String()

	// split on the newline
	details := strings.Split(output, "\n")

	//Split on the equals sign and then trim off the quotes for both the name and version id values
	distrib := strings.TrimSpace(strings.Split(details[0], "=")[1])

	distrib = strings.Trim(distrib, "\"")

	version := strings.TrimSpace(strings.Split(details[1], "=")[1])

	version = strings.Trim(version, "\"")

	return distrib, version, nil
}
