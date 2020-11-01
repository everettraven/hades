package main

import (
	"fmt"
	"github.com/everettraven/hades/tests"
	"github.com/everettraven/hades/utils"
	"strings"
)

func main() {
	// Parse the tests from the HCL file
	unitTests, err := utils.ParseUnitTests("unittests.hcl")

	// Make sure we didn't hit any errors while parsing the HCL file
	if err != nil {
		fmt.Println(err)
		return
	}

	// Loop through all the tests we parsed
	for i := 0; i < len(unitTests.Tests); i++ {
		// Get the current test
		curTest := unitTests.Tests[i]

		// Set up some initial variables
		depFail := false
		testFail := false
		var failErr error

		// Print out the test that we are running
		fmt.Println("Running Test: " + curTest.Name)

		// Attempt to pull the Docker Image corresponding to the test
		err := curTest.GetImage()

		// Check for errors
		if err != nil {
			// Set the fail message
			failErr = fmt.Errorf("Could not pull the Docker Image %s - %s", curTest.Image, err.Error())
			// Set that a dependency failed
			depFail = true
		}

		// Don't run this portion if we already have a dependency failure
		if !depFail {
			// Attempt to run the image corresponding to the current test
			err = curTest.RunImage()

			// Check for errors
			if err != nil {
				// Set the fail message
				failErr = fmt.Errorf("Could not run the Docker Image %s with Container Name %s - %s", curTest.Image, curTest.ContainerName, err.Error())
				// Set that a dependency failed
				depFail = true
			}
		}

		// Don't run this portion if we already have a dependency failure
		if !depFail {
			// Attempt to run the "remote" command
			arguments := strings.Join(curTest.TestCommand.Arguments, " ")
			command, err := tests.TestRemoteCommand("127.0.0.1", curTest.Port, curTest.TestCommand.Name, arguments)

			// If we encountered any errors lets handle it.
			if err != nil {
				// Set the failure message
				failErr = err
				// Set that the test failed
				testFail = true
			} else {
				testOutput := strings.TrimSpace(command.Stdout.String())
				// Check to see if the output from running the test matches the expected output
				if testOutput != curTest.ExpectedOutput {
					// Test fails if it doesn't match
					testFail = true
					// Set the failure message
					failErr = fmt.Errorf("Output of the test did not match the Expected Output (Output: %s | Expected: %s)", testOutput, curTest.ExpectedOutput)
				}
			}
		}

		// Don't run this if there was a previous dependency failure
		if !depFail {
			// Attempt to stop the test container
			err = curTest.StopContainer()

			if err != nil {
				failErr = fmt.Errorf("Could not stop the Docker Container %s - %s", curTest.ContainerName, err.Error())
				depFail = true
			}
		}

		// Don't run this if there was a previous dependency failure
		if !depFail {
			// Attempt to remove the container
			err = curTest.RemoveContainer()
			if err != nil {
				failErr = fmt.Errorf("Could not remove the Docker Container %s - %s", curTest.ContainerName, err.Error())
				depFail = true
			}
		}

		// Don't run this if there was a previous dependency failure
		if !depFail {
			err = curTest.RemoveImage()
			if err != nil {
				failErr = fmt.Errorf("Could not remove the Docker Image %s - %s", curTest.Image, err.Error())
				depFail = true
			}
		}

		// If there was a dependency failure or the test failed the test will fail. Otherwise the test passes
		if depFail || testFail {
			fmt.Println("\tFailed - ", failErr)
			// Add a message to remind to clean up docker stuff if there is an issue in a dependency
			if depFail {
				fmt.Println("\tMake sure you clean up your Docker containers and images as an error was encountered with Docker while attempting to run the tests")
			}
		} else {
			fmt.Println("\tPassed")
		}
	}

}
