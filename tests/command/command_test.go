package command

import (
	"strings"

	"testing"

	"github.com/everettraven/hades/utils"
)

//TestCommand - Unit test the command function
func TestCommand(t *testing.T) {
	t.Log("Parsing test HCL file")
	// Parse the tests from the HCL file
	unitTests, err := utils.ParseUnitTests("command_test.hcl")

	// Make sure we didn't hit any errors while parsing the HCL file
	if err != nil {
		t.Fatal(err)
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
		//--------------------------------------------------------------------------------------

		// Loop through all the command blocks in the test
		for j := 0; j < len(curTest.Run.Cmd); j++ {
			// Attempt to run the "remote" command
			arguments := strings.Join(curTest.Run.Cmd[j].Arguments, " ")
			command, err := TestRemoteCommand("127.0.0.1", curTest.Port, curTest.Run.Cmd[j].Name, arguments)

			// If we encountered any errors lets handle it.
			if err != nil {
				// Set the failure message
				t.Fatalf(err.Error())
			} else {
				testOutput := strings.TrimSpace(command.Stdout.String())
				// Check to see if the output from running the command matches the expected output
				if testOutput != curTest.Run.Cmd[j].ExpectedOutput {
					// Set the failure message
					t.Errorf("Output of the test did not match the Expected Output (Output: %s | Expected: %s)", testOutput, curTest.Run.Cmd[j].ExpectedOutput)
				} else {
					t.Logf("Command \"%s\" - Passed", command.Name+" "+arguments)
				}
			}
		}

		//--------------------------------------------------------------------------------------

		t.Logf("Stopping Docker Container %s", curTest.ContainerName)
		if err = curTest.StopContainer(); err != nil {
			t.Fatalf("Error encountered while attempting to stop container %s: %s", curTest.ContainerName, err.Error())
		}

		t.Logf("Removing Docker Container %s", curTest.ContainerName)
		if err = curTest.RemoveContainer(); err != nil {
			t.Fatalf("Error encountered while attempting to remove container %s: %s", curTest.ContainerName, err.Error())
		}

		// If it failed any of the command tests lets fail now before we continue to loop
		if t.Failed() {
			t.FailNow()
		}

	}

}
