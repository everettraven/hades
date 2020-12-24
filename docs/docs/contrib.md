---
id: contrib
title: Contributing to hades
---

Thank you for your interest in contributing to the hades project! Please follow the following guidelines when contributing to the project.

# Set Up Your Development Environment

## Install Go
hades is being developed in the Go programming language, so it is important to have it installed if you plan to contribute to this project. The instructions to install Go can be found at https://golang.org/

## Install Docker
hades uses Docker to run tests that simulate connecting to an external server and running remote commands over SSH. In order to contribute to hades, it is important to run the tests. To run the tests you will need to have Docker installed and running on your machine. The instructions on how to install Docker can be found at https://docs.docker.com/get-docker/

## Fork the Repository
Now that you have both Docker and Go, you need to fork the hades repository. This will ensure that you are making changes to your very own copy of hades.

## Clone Your Repository
Now that you have your own fork of hades, it is time to clone the repository to your local filesystem so you can make changes. Make sure to change directories to where you want the repository to exist on your filesystem and then run `git clone <repository-link>` to clone your repository.

# Developing for hades

## Start Developing
Now that you have everything cloned and set up, you can start developing any features, tests, or components you think would improve this project. Whenever you are developing any new features, be sure that a test exists for that feature or that you are creating a test before the development of that feature. hades follows the test driven development workflow which means a feature shouldn't be developed before the test. 

## Run Tests
Before your contribution is accepted it will have to have passed all the tests corresponding to that feature, so make sure to run the tests before committing your changes! To run all the tests you can run:
```
go test ./...
```

To run a specific test package:
```
go test ./<package>
```

Example to run the command test package:
```
go test ./tests/command
```

The Go testing framework will cache test results for later use. If you want to clear this cache you can run:

```
go clean -testcache
```

One last thing to note: when running the tests Docker images corresponding to the test are pulled from DockerHub. These images are not removed by the tests so that it is faster to run them in the future. If you want to remove them, you can use the `docker image list` command to list what images you currently have pulled and use `docker image rm <image name>` to remove them from your system.

## Adding dependencies
If you need to add any dependencies when developing for hades, make sure that the changes are made in the go.mod file. To ensure you are getting the necessary dependencies use `go get` from the root of the hades directory as so:

```
go get github.com/everettraven/hades
```

and then run 

```
go mod tidy
```

## Make a Pull Request
Now that you have finished your contribution, it is time to make a Pull Request. When making your PR make sure to follow the template to describe what type of changes you made and the reason for these changes.