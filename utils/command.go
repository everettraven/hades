package utils

import (
	"bytes"
	"os/exec"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Command struct {
	Name           string   `hcl:"name"`
	Arguments      []string `hcl:"args"`
	Cmd            *exec.Cmd
	Stdout, Stderr bytes.Buffer
	ExitStatus     int
}

func NewCommand(name string, args ...string) *Command {
	command := new(Command)
	command.Name = name
	command.Arguments = args
	command.Cmd = exec.Command(name, args...)
	return command
}

func (c *Command) RunLocal() error {
	c.Cmd.Stdout = &c.Stdout
	c.Cmd.Stderr = &c.Stderr

	_, err := exec.LookPath(c.Name)

	if err != nil {
		return err
	}

	err = c.Cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func (c *Command) RunRemote(client *ssh.Client) error {
	session, err := client.NewSession()
	if err != nil {
		return err
	}

	session.Stdout = &c.Stdout
	session.Stderr = &c.Stderr

	command := c.Name + " " + strings.Join(c.Arguments, " ")

	err = session.Run(command)

	if err != nil {
		return err
	}

	return nil
}
