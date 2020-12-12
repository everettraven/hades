package utils

import (
	"context"

	"io"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// RunBlock is a struct to hold the data of a run block of a test
type RunBlock struct {
	Cmd Command `hcl:"command,block"`
}

// UnitTestUtil is a struct to hold our test data.
type UnitTestUtil struct {
	Name           string   `hcl:"name,label"`
	Image          string   `hcl:"image"`
	Port           string   `hcl:"port"`
	ContainerName  string   `hcl:"containerName"`
	ExpectedOutput string   `hcl:"expectedOutput"`
	Run            RunBlock `hcl:"run,block"`
}

// NewUnitTestUtil Returns a new UnitTestUtil
func NewUnitTestUtil(name string, image string, port string, containerName string) *UnitTestUtil {
	UnitTestUtil := new(UnitTestUtil)
	UnitTestUtil.Name = name
	UnitTestUtil.Image = image
	UnitTestUtil.Port = port
	UnitTestUtil.ContainerName = containerName
	return UnitTestUtil
}

// GetImage will get the specified image for the test
func (test *UnitTestUtil) GetImage() error {
	// Get a new client with API Version Negotiation to make sure the client will work even with newest versions of Docker
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	// Make sure we didn't hit an error
	if err != nil {
		return err
	}

	// Pull the Image requested
	reader, err := cli.ImagePull(context.Background(), test.Image, types.ImagePullOptions{})

	io.Copy(ioutil.Discard, reader)

	defer reader.Close()

	// Make sure we didn't hit an error
	if err != nil {
		return err
	}

	// Now that the Image is pulled, close out our client
	err = cli.Close()

	// Make sure we didn't hit an error
	if err != nil {
		return err
	}

	// If we made it here we never hit an error so return nil
	return nil
}

// RunImage runs the pulled test image in a container
func (test *UnitTestUtil) RunImage() error {
	// Get a new client with API Version Negotiation to make sure the client will work even with newest versions of Docker
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())

	// Set the context. We want everything to run in the background.
	ctx := context.Background()

	// Ensure no errors were encountered
	if err != nil {
		return err
	}

	// We want to make sure that the container we start has the SSH port open so we can test SSH and running commands
	containerConfig := &container.Config{
		Image: test.Image,
		ExposedPorts: nat.PortSet{
			nat.Port("22/tcp"): {},
		},
	}

	// Set up the config for the host. Make sure that we open up the ports specified in the test being ran
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("22/tcp"): []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: test.Port}},
		},
	}

	// Create the container with our settings
	container, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, test.ContainerName)

	// Make sure no errors in creating the container
	if err != nil {
		return err
	}

	// Start the container we created. Also check for errors at the same time...
	if err := cli.ContainerStart(ctx, container.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	// Made it here so no errors
	return nil
}

// StopContainer - Function to stop a test container
func (test *UnitTestUtil) StopContainer() error {
	// Get a new client with API Version Negotiation to make sure the client will work even with newest versions of Docker
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())

	// Make sure no errors
	if err != nil {
		return err
	}

	// Run this in the background
	ctx := context.Background()

	// Stop the container
	if err := cli.ContainerStop(ctx, test.ContainerName, nil); err != nil {
		return err
	}

	// Made it here so no errors
	return nil
}

// RemoveContainer - Function to remove a Docker Container
func (test *UnitTestUtil) RemoveContainer() error {
	// Get a new client with API Version Negotiation to make sure the client will work even with newest versions of Docker
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())

	// Make sure we don't hit any errors
	if err != nil {
		return err
	}

	// Set the context to run it in the background
	ctx := context.Background()

	// Set some removal options
	rmOpts := &types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	// Attempt to remove the container
	if err := cli.ContainerRemove(ctx, test.ContainerName, *rmOpts); err != nil {
		return err
	}

	// If we made it here no errors
	return nil
}

// RemoveImage - Function to remove a Docker Image from the system
func (test *UnitTestUtil) RemoveImage() error {
	// Get a new client with API Version Negotiation to make sure the client will work even with newest versions of Docker
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())

	// Make sure we didn't hit an error
	if err != nil {
		return err
	}

	// Run in the background
	ctx := context.Background()

	// Set up the removal options
	rmOpts := &types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	}

	// Attempt to remove the image
	if _, err := cli.ImageRemove(ctx, test.Image, *rmOpts); err != nil {
		return err
	}

	// If we made it here there are no errors
	return nil
}
