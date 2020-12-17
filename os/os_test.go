package os

import (
	"testing"
	"time"

	"github.com/everettraven/hades/utils"
	"golang.org/x/crypto/ssh"

	"strings"
)

func TestOS(t *testing.T) {
	t.Log("Parsing test HCL file")
	// Parse the tests from the HCL file
	unitTests, err := utils.ParseUnitTests("os_test.hcl")

	// Make sure we didn't hit any errors while parsing the HCL file
	if err != nil {
		t.Fatal(err)
	}

	// Set up the SSH config for testing
	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("root")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	// Loop through all the tests we parsed
	for i := 0; i < len(unitTests.UnitTests); i++ {
		// Get the current test
		curTest := unitTests.UnitTests[i]

		// Print out the test that we are running
		t.Logf("Running Test: %s", curTest.Name)

		t.Logf("Pulling Docker Image %s", curTest.Image)
		// Attempt to pull the Docker Image corresponding to the test
		err := curTest.GetImage()

		// Check for errors
		if err != nil {
			// Set the fail message
			t.Fatalf("Could not pull the Docker Image %s - %s", curTest.Image, err.Error())
		}

		t.Logf("Running the Docker Image %s as Container %s", curTest.Image, curTest.ContainerName)
		// Attempt to run the image corresponding to the current test
		err = curTest.RunImage()

		// Check for errors
		if err != nil {
			// Set the fail message
			t.Fatalf("Could not run the Docker Image %s with Container Name %s - %s", curTest.Image, curTest.ContainerName, err.Error())
		}

		t.Log("Waiting until SSH is running in the container")
		running, err := curTest.SSHDRunning()

		if err != nil {
			t.Fatalf("Encountered an issue while checking status of container %s - %s", curTest.ContainerName, err.Error())
		}

		// Loop and check status of the container (only continue once the status is running)
		for running != true {
			running, err = curTest.SSHDRunning()

			if err != nil {
				t.Fatalf("Encountered an issue while checking status of container %s - %s", curTest.ContainerName, err.Error())
				break
			}
		}

		t.Logf("SSH is running in the container")

		// Actual test implementation
		//-----------------------------------------------------------------------------------------------------

		// Get the details from the operating system on the remote machine
		osDetails, err := utils.GetRemoteOS("127.0.0.1", curTest.Port, config)

		// Check for any errors
		if err != nil {
			t.Error(err.Error())
		}

		// Parse the returned OS info for the distro and version
		details := strings.Split(osDetails, ":")
		distro := details[0]
		version := details[1]

		// test to see if it matches the expected distro
		if strings.ToLower(curTest.Run.Os.DistributionID) != strings.ToLower(distro) {
			t.Errorf("OS of the remote machine does not match the expected value: Remote OS: %s", distro)
		}

		// test to see if it matches the expected version
		if curTest.Run.Os.Version != "" {
			if strings.ToLower(curTest.Run.Os.Version) != strings.ToLower(version) {
				t.Errorf("OS version of the remote machine does not match the expected value: Remote OS Version: %s", version)
			}
		}

		//-----------------------------------------------------------------------------------------------------

		t.Logf("Stopping Docker Container %s", curTest.ContainerName)
		if err = curTest.StopContainer(); err != nil {
			t.Fatalf("Error encountered while attempting to stop container %s: %s", curTest.ContainerName, err.Error())
		}

		t.Logf("Removing Docker Container %s", curTest.ContainerName)
		if err = curTest.RemoveContainer(); err != nil {
			t.Fatalf("Error encountered while attempting to remove container %s: %s", curTest.ContainerName, err.Error())
		}

		if len(unitTests.UnitTests) > i+1 {
			if unitTests.UnitTests[i+1].Image != curTest.Image {
				t.Logf("Removing Docker Image %s", curTest.Image)
				if err = curTest.RemoveImage(); err != nil {
					t.Fatalf("Error encountered while attempting to remove image %s: %s", curTest.Image, err.Error())
				}
			}
		} else {
			t.Logf("Removing Docker Image %s", curTest.Image)
			if err = curTest.RemoveImage(); err != nil {
				t.Fatalf("Error encountered while attempting to remove image %s: %s", curTest.Image, err.Error())
			}
		}

		// If it failed any of the command tests lets fail now before we continue to loop
		if t.Failed() {
			t.FailNow()
		}

	}

}
