package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/everettraven/hades/resources"
	"github.com/everettraven/hades/utils"
	"golang.org/x/crypto/ssh"
)

// Set up some variables to hold the command line arguments for now
var hosts string
var file string
var user string
var pass string
var bof bool
var dir string
var testDir string
var keyFile string
var local bool

func init() {
	// Create the hosts flag
	flag.StringVar(&hosts, "hosts", "", "The hosts flag can be used to set the host(s) that the program will run the designated tests on. Each host should be separated by a comma when used in the command line. EX: 127.0.0.1,127.1.1.1,...")

	// Create the file flag
	flag.StringVar(&file, "file", "", "The file flag can be used to set the location of the HCL file to use. If this is not set, hades will attempt to find a file to use on its own.")

	// Create the user flag
	flag.StringVar(&user, "user", "", "The user flag can be used to set the username to SSH into the remote system with. Overrides the user specified in a hosts.hcl file if one exists")

	// Create the password flag
	flag.StringVar(&pass, "pass", "", "The pass flag can be used to set the password to use when attempting to SSH into the remote system. If this flag is used hades will use the password given instead of using an SSH Key")

	// Create the bof flag
	flag.BoolVar(&bof, "bof", false, "The bof flag is used to stop running the test on any failure.")

	// Create the dir flag
	flag.StringVar(&dir, "dir", "", "The dir flag can be used to set the directory the program will use as the root directory when running.")

	// Create the test-dir flag
	flag.StringVar(&testDir, "test-dir", "tests", "The test-dir flag can be used to set the directory that hades will pull test files from. This defaults to 'tests'")

	// Create the key flag
	if runtime.GOOS == "windows" {
		flag.StringVar(&keyFile, "key-file", fmt.Sprintf("%s", os.Getenv("HOMEDRIVE")+os.Getenv("HOMEPATH")+"/.ssh/id_rsa"), "The key flag is used to set the path to the SSH Key to use when attempting to connect to a remote host.")
	} else {
		flag.StringVar(&keyFile, "key-file", fmt.Sprintf("%s", os.Getenv("HOME")+"/.ssh/id_rsa"), "The key flag is used to set the path to the SSH Key to use when attempting to connect to a remote host.")
	}

	flag.BoolVar(&local, "local", false, "The local flag is used to run hades on the local system rather than on a remote system.")
}

func main() {

	// Get the flags from the command line
	flag.Parse()

	failed := false
	_ = failed

	// List of test files
	var testFiles []string

	// List of hosts
	var hostList []utils.Host

	// See if it is meant to be run local
	if !local {
		// Check for a hosts HCL file
		if hosts == "" {
			var curDir string
			var err error
			if dir == "" {
				// Get the current working directory
				curDir, err = os.Getwd()
				if err != nil {
					fmt.Printf("Encountered an error when attempting to get the hosts file: %s\n", err.Error())
					os.Exit(1)
				}
			} else {
				// Set the working directory
				curDir = dir
			}

			// Look for the hosts.hcl file
			err = filepath.Walk(curDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if !info.IsDir() && info.Name() == "hosts.hcl" {
					temp, err := utils.Parse(path, utils.HostHCLUtil{})
					if err != nil {
						return err
					}

					hostList = temp.(utils.HostHCLUtil).Hosts
				}

				return nil
			})

			if err != nil {
				fmt.Printf("Encountered an error when attempting to get the hosts file: %s\n", err.Error())
				os.Exit(1)
			}

		} else {
			// get a list of the hosts from the hosts flag
			hostList = GetHosts(hosts)
		}
	} else {
		hostList = GetHosts("local")
	}

	// Check for the test HCL file being passed into the command
	if file == "" {
		// Get it from the last command line argument
		file = flag.Arg(0)
	}

	// Make sure that there is actually a file passed in. If not then lets look for the file/files to use
	if file == "" {
		if dir != "" {
			testDir = dir + testDir
		}

		if _, err := os.Stat(testDir); os.IsNotExist(err) {
			fmt.Println("No test files could be found to run.")
			os.Exit(1)
		}

		// Get the list of files from the test directory
		files, err := ioutil.ReadDir(testDir)

		if err != nil {
			fmt.Printf("Encountered an error while attempting to find the test files: %s\n", err.Error())
			os.Exit(1)
		}

		// Add all the files with the .hcl file ending to the test files list
		for _, f := range files {
			if strings.Contains(f.Name(), ".hcl") {
				testFiles = append(testFiles, testDir+"/"+f.Name())
			}
		}

	} else {
		// Check if multiple files are being passed in by looking for any commas
		testFiles = strings.Split(file, ",")
	}

	// Check to make sure that there are test files to be run
	if len(testFiles) <= 0 {
		fmt.Printf("No test files could be found to run.\n")
		os.Exit(1)
	}

	for i := 0; i < len(hostList); i++ {
		// Get the current host
		curHost := hostList[i]

		for j := 0; j < len(testFiles); j++ {
			testFile := testFiles[j]

			if dir != "" {
				if !strings.Contains(testFile, dir) {
					testFile = dir + testFile
				}
			}

			// Get the test data from the HCL file
			parseOut, err := utils.Parse(testFile, utils.Test{})

			test := parseOut.(utils.Test)

			if err != nil {
				fmt.Println("Encountered an error when attempting to read the HCL file: ", err.Error())
				os.Exit(1)
			}

			if local {
				// Print the current test running
				fmt.Printf("Running Test: %s on %s\n", test.Title, curHost.IP)
			} else {
				// Print the current test running
				fmt.Printf("Running Test: %s on %s\n", test.Title, curHost.IP+":"+curHost.Port)
			}

			// Get the ssh client
			var sshClient *ssh.Client
			var authMethod []ssh.AuthMethod

			// Check if a password is given to determine the ssh AuthMethod to use
			if pass != "" && !local {
				authMethod = []ssh.AuthMethod{ssh.Password(pass)}
			} else if pass == "" && !local {
				// Get the key from the keyFile
				key, err := ioutil.ReadFile(keyFile)

				if err != nil {
					fmt.Printf("Unable to read the private key: %s\n", err.Error())
					os.Exit(1)
				}

				// Get the signer from the private key
				signer, err := ssh.ParsePrivateKey(key)

				if err != nil {
					fmt.Printf("Unable to parse the private key: %s\n", err.Error())
					os.Exit(1)
				}

				// setup the auth method
				authMethod = []ssh.AuthMethod{ssh.PublicKeys(signer)}
			}

			// Get the SSH client based on the user command flag being passed in.
			if user != "" && !local {

				sshClient, err = GetSSHClient(curHost.IP, curHost.Port, user, authMethod)

				if err != nil {
					fmt.Println("Encountered an error when attempting to connect to the host specified: ", err.Error())
					os.Exit(2)
				}
			} else if curHost.User != "" && !local {
				sshClient, err = GetSSHClient(curHost.IP, curHost.Port, curHost.User, authMethod)

				if err != nil {
					fmt.Println("Encountered an error when attempting to connect to the host specified: ", err.Error())
					os.Exit(2)
				}
			} else if !local {
				fmt.Println("No valid user was specified for this host. Skipping this host.")
				continue
			}

			// Check if there are any command blocks that need to be run
			if len(test.Cmds) > 0 {
				failed = RunCommands(sshClient, test.Cmds, local)
			}

			// Check if there is an OS test block that needs to be run
			if test.Os != nil {
				failed = CheckOS(sshClient, *test.Os, local)
			}
		}

	}

}

//GetSSHClient - Function to setup the SSH client for use when dialing the remote machine
func GetSSHClient(host string, port string, user string, auth []ssh.AuthMethod) (*ssh.Client, error) {
	// Set up the SSH Config to use
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	// Get the ssh client
	client, err := utils.GetSSHClient(host, port, config)

	if err != nil {
		return client, err
	}

	return client, nil
}

//RunCommands - Function to run the command blocks in the test
func RunCommands(sshClient *ssh.Client, cmds []*resources.Command, local bool) bool {
	var command *resources.Command
	failed := false
	// Loop through all the command blocks in the test
	for j := 0; j < len(cmds); j++ {
		// Attempt to run the "remote" command
		arguments := strings.Join(cmds[j].Arguments, " ")

		if local {
			if runtime.GOOS == "windows" {
				command = resources.NewCommand("powershell", cmds[j].Name+" "+arguments)
			} else {
				command = resources.NewCommand(cmds[j].Name, arguments)
			}
		} else {
			command = resources.NewCommand(cmds[j].Name, arguments)
		}

		var err error

		if local {
			err = command.RunLocal()
		} else {
			err = command.RunRemote(sshClient)
		}

		// If we encountered any errors lets handle it.
		if err != nil {
			// Set the failure message
			fmt.Printf("\tCommand %s Failed - Encountered an Error while trying to run the command: %s\n ", cmds[j].Name, err.Error())
			failed = true
		} else {
			testOutput := strings.TrimSpace(command.Stdout.String())
			// Check to see if the output from running the command matches the expected output
			if testOutput != cmds[j].ExpectedOutput {
				// Set the failure message
				fmt.Printf("\tCommand \"%s\" Failed - Output of the test did not match the Expected Output (Output: %s | Expected: %s)\n", cmds[j].Name, testOutput, cmds[j].ExpectedOutput)
				failed = true
			} else {
				fmt.Printf("\tCommand \"%s\" - Passed\n", command.Name+" "+arguments)
			}
		}

		if bof && failed {
			os.Exit(3)
		}
	}

	return failed
}

//GetHosts - Function to parse a string and get the hosts from it
func GetHosts(hosts string) []utils.Host {

	// Set the output variable
	var out []utils.Host

	// Get the host data from splitting the strings by a comma
	hostData := strings.Split(hosts, ",")

	// Loop through each bit of host data
	for i := 0; i < len(hostData); i++ {
		// Get the current host
		curHost := hostData[i]
		// Check if the current host specifies a port by checking if it contains a ':'
		if strings.Contains(curHost, ":") {
			// Split the current host based on the colon to separate the ip and port
			hostSplit := strings.Split(curHost, ":")
			// Append a new host to the host list
			out = append(out, utils.NewHost(hostSplit[0], hostSplit[1]))
		} else {
			// Append a new host to the host list, but with port 22 since no specific port was given
			out = append(out, utils.NewHost(curHost, "22"))
		}
	}

	return out
}

// CheckOS - Function to check the OS of the system with the expected OS
func CheckOS(sshClient *ssh.Client, os resources.OS, local bool) bool {
	failed := false

	var distro string
	var version string
	var err error

	// Get the Remote OS details
	if local {
		distro, version, err = resources.GetLocalOS()
	} else {
		distro, version, err = resources.GetRemoteOS(sshClient)
	}

	if err != nil {
		fmt.Printf("\tOS Test Failed - An Error Occured While Attempting to Get the Remote Machine's OS: %s\n", err.Error())
		failed = true
	}

	// Check that the distros match
	if strings.ToLower(distro) != strings.ToLower(os.DistributionID) {
		fmt.Printf("\tOS Test Failed - The Distribution of the Remote OS did not match the expected Distribution (Output: %s | Expected: %s)\n", distro, os.DistributionID)
		failed = true
	}

	// Check that the version matches
	if os.Version != "" {
		if strings.ToLower(version) != strings.ToLower(os.Version) {
			fmt.Printf("\tOS Test Failed - The Version of the Remote OS did not match the expected Version (Output: %s | Expected: %s)\n", version, os.Version)
			failed = true
		}
	}

	// If the test hasn't failed at this point then state that it has passed
	if !failed {
		fmt.Printf("\tOS Test - Passed\n")
	}

	return failed
}
