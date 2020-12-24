---
id: multiple_tests
title: Multiple Tests
---

This guide shows how to use hades with a simple test file that runs a very simple command. All source code for this example can be found [here](https://github.com/everettraven/hades/tree/main/examples/multiple_tests)

This guide uses Docker to create a container that would simulate a remote machine running on your local machine. If you have an existing remote machine feel free to skip the Docker setup steps.

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
docker run --name multiple_tests -d -p 5000:22 bpalmer/ssh_test
```
This will run the Docker Container in the background with your localhost port 5000 mapped to port 22 (the standard SSH port) on the Docker Container. It also names the container 'multiple_tests' for easy cleanup when we are done with it. Feel free to play with these values as you see fit, but make sure to adjust where these values are used in the future pieces of the guide.

For reference:

Both the username and password for this Docker Container is **root**

## Create the Test Files

Make sure you are in the directory you would like to store the source code of this guide in and make a directory named **tests**.

hades will by default look for the **tests** folder and run all the tests within that folder but you can name this directory whatever you would like and pass the `--test-dir` flag followed by the name of your test directory and hades will use that folder.

In the **tests** folder we are going to create 2 test files named **test1.hcl** and **test2.hcl**.

In **test1.hcl** put:
```hcl
title = "Multiple Tests - Test #1"

command {
    cmd = "echo"
    args = ["Multiple"]
    expectedOutput = "Multiple"
}

command {
    cmd = "ls"
    args = ["/usr"]
    expectedOutput = "bin\ngames\ninclude\nlib\nlib32\nlib64\nlibx32\nlocal\nsbin\nshare\nsrc"
}

os {
    distributionID = "ubuntu"
}
```

You can also do the following to get rid of the optional args parameter in the command block if you would like:

```hcl
title = "Multiple Tests - Test #1"

command {
    cmd = "echo Multiple"
    expectedOutput = "Multiple"
}

command {
    cmd = "ls /usr"
    expectedOutput = "bin\ngames\ninclude\nlib\nlib32\nlib64\nlibx32\nlocal\nsbin\nshare\nsrc"
}

os {
    distributionID = "ubuntu"
}
```

In **test2.hcl** put:
```hcl
title = "Multiple Tests - Test #2"

command {
    cmd = "echo"
    args = ["hades is working!"]
    expectedOutput = "hades is working!"
}
```

You can also do the following to get rid of the optional args parameter in the command block if you would like:

```hcl
title = "Multiple Tests - Test #2"

command {
    cmd = "echo hades is working!"
    expectedOutput = "hades is working!"
}
```

We recommend that you create a hosts folder for better organization, but it is not necessary when creating the **hosts.hcl** file as hades will look for it as long as it is in the current directory or a sub-directory of the current directory. In our case we are going to create the hosts folder and then within that folder create the **hosts.hcl** file. In the **hosts.hcl** file we are going to put the following:

```hcl
host {
    ip = "127.0.0.1"
    port = "5000"
    user = "root"
}
```
If you wanted to run the tests on multiple hosts you can place multiple host blocks in the hosts file. hades will run all the tests on each of the hosts.

## Run hades on the Remote System
Now that we have a simple test file created we can run hades to test the remote system. In this case we will run it on our Docker Container we created for this guide by running:
```
hades --pass root
```

If you did not name the test directory **tests** you can let hades know what directory to use for getting the tests by running:
```
hades --test-dir [directory] --pass root
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
hades
```

If you did not name the test directory **tests** you can let hades know what directory to use for getting the tests by running:
```
hades --test-dir [directory]
```

Running hades with a non-default SSH Key path:
```
hades --key-file [key path]
```

If you did not name the test directory **tests** you can let hades know what directory to use for getting the tests by running:
```
hades --test-dir [directory] --key-file [key path]
```

## Docker Cleanup
All the guides use Docker Containers so if you plan to continue with the rest of the guides you can return to this step when you are finished with the guides you would like to go through.

If you would like to cleanup the Docker Container now you can run the following commands to stop and remove the Docker Container:
```
docker container stop multiple_tests
```

and

```
docker container rm multiple_tests
```

All the guides will use the same Docker Image for their Containers, but if you would like to remove the image as well you can run:
```
docker image rm bpalmer/ssh_test
```