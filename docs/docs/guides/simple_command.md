---
id: simple_command
title: Simple Command
---

This guide shows how to use hades with a simple test file that runs a very simple command. All source code for this example can be found [here](https://github.com/everettraven/hades/tree/main/examples/simple_command)

This guide uses Docker to create a container that would simulate a remote machine running on your local machine. If you have an existing remote machine feel free to skip the Docker setup steps.

This guide can also be run locally by passing the `--local` flag when running hades. It is important to note that this guide was made to run with the Docker container so some changes may need to be made in order for all the tests to pass when running them locally. If you are running it locally on a Windows system, changes will need to be made to the tests as the test commands on Windows results in a different output than the Ubuntu 20.04 Docker container.

## Docker Setup

1. Make sure that you have docker installed by running:
```
docker -v
```
If you do not have Docker installed you can install it by following the offical Docker installation instructions: https://docs.docker.com/get-docker/

2. Pull the Docker Image we are going to use by running:
```
docker pull bpalmer/ssh_test
```
3. Run the Docker Container by running:
```
docker run --name simple_command -d -p 5000:22 bpalmer/ssh_test
```
This will run the Docker Container in the background with your localhost port 5000 mapped to port 22 (the standard SSH port) on the Docker Container. It also names the container 'simple_command' for easy cleanup when we are done with it. Feel free to play with these values as you see fit, but make sure to adjust where these values are used in the future pieces of the guide.

For reference:

Both the username and password for this Docker Container is **root**

## Create the Test File

Make sure you are in the directory you would like to store the source code of this guide in and make your test file. We are going to name it **simple_command.hcl**. The .hcl file extension is required and signifies that is a Hashicorp Configuration Language file. Open the file in your IDE or text editor and add the following to it:
```hcl
title = "Simple Command Test"

command {
    cmd = "echo"
    args = ["Hello Infrastructure Testing World!"]
    expectedOutput = "Hello Infrastructure Testing World!"
}
```

You can also do the following to get rid of the optional args parameter in the command block if you would like:

```hcl
title = "Simple Command Test"

command {
    cmd = "echo Hello Infrastructure Testing World!"
    expectedOutput = "Hello Infrastructure Testing World!"
}
```

## Run hades
Now that we have a simple test file created we can run hades to test the remote system. In this case we will run it on our Docker Container we created for this guide by running:
```
hades --hosts 127.0.0.1:5000 --user root --pass root simple_command.hcl
```

Run it locally (does not require authentication):
```
hades --local simple_command.hcl
```

It is not recommended to run hades by passing the `--pass` flag to the command as this puts the password for the remote machine in your command line history as plaintext. We recommend running hades using an SSH key. If you are interested in doing so we first need to do some setup:

1. Create an SSH key by running:
```
ssh-keygen
```
2. Send the SSH key to the remote system (or Docker Container in this case):

Windows (Powershell):
```powershell
type $env:USERPROFILE\.ssh\id_rsa.pub | ssh root@127.0.0.1 -p 5000 "cat >> .ssh/authorized_keys"
```

Unix:
```
ssh-copy-id -i ~/.ssh/id_rsa root@127.0.0.1 -p 5000
```

Now that the SSH Key has been sent to the remote system we can run hades with the SSH key. By default hades attempts to get the SSH key from `~/.ssh/id_rsa` so if you created the key to a different directory you can use the flag `--key-file` followed by the path to the SSH key.

Running hades with the SSH Key:
```
hades --hosts 127.0.0.1:5000 --user root simple_command.hcl
```

Running hades with a non-default SSH Key path:
```
hades --hosts 127.0.0.1:5000 --user root --key-file [key path] simple_command.hcl
```

## Docker Cleanup
All the guides use Docker Containers so if you plan to continue with the rest of the guides you can return to this step when you are finished with the guides you would like to go through.

If you would like to cleanup the Docker Container now you can run the following commands to stop and remove the Docker Container:
```
docker container stop simple_command
```

and

```
docker container rm simple_command
```

All the guides will use the same Docker Image for their Containers, but if you would like to remove the image as well you can run:
```
docker image rm bpalmer/ssh_test
```