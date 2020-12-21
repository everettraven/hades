# Example - Simple Command
This example showcases a host file and multiple test files to test multiple things on a remote server.

# Recreate This Example
This example is geared to use a Docker container to simulate accessing a remote server. This allows you to do everything locally without having to have an actual remote server. If you do have a remote server, feel free to play around and make this example work with a remote server!

If you don't have Docker installed -> https://docs.docker.com/get-docker/

## Clone the Repo
In order to run this example, clone this repo by running:
```
git clone https://github.com/everettraven/hades.git
```

and then change directory into the cloned repo by running:
```
cd hades
```

## Pull the Docker Image
The Docker Image we are going to use for this example is `bpalmer/ssh_test`. In order to use it we need to make sure that we have the image pulled to our local machine by running:
```
docker pull bpalmer/ssh_test
```

## Run the Docker Container
We are going to run the Docker container in the background and on port 5000. Do this by running:
```
docker run --name multiple_tests -d -p 5000:22 bpalmer/ssh_test
```

## Change Directory
Change your directory into the multiple_tests directory by running:
```
cd examples/multiple_tests
```

## Run hades
If you followed the build instructions to install hades onto your system you can run:
```
hades --pass root
```

If you didn't follow the build instructions you can also run this example by running:
```
go run ../../main.go --pass root
```

If you would like to use an SSH Key (preferred) over a password as shown in the previous command:

Create an SSH Key by running `ssh-keygen`

If you are using Windows you can share the key to the Docker container using:
```
type $env:USERPROFILE\.ssh\id_rsa.pub | ssh root@127.0.0.1 -p 5000 "cat >> .ssh/authorized_keys"
```

If you are using a UNIX based system you can share the key to the Docker container using:
```
ssh-copy-id -i ~/.ssh/id_rsa root@127.0.0.1 -p 5000
```

Now you can run hades using the SSH Key by running:

```
hades
```
or
```
go run ../../main.go
```

You should see that the test runs properly and passes all internal tests

## Clean Up Docker
Once you are done with the example it is important to clean up the running container. In order to stop and remove the container this example used, run:
```
docker container stop multiple_tests
```

and then
```
docker container rm multiple_tests
```

All examples use the same `bpalmer/ssh_test` Docker Image, but if you would like to remove it you can run:
```
docker image rm bpalmer/ssh_test
```