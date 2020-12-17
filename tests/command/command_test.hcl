unittest "Command Test" {
    image = "bpalmer/ssh_test"
    port = "9090"
    containerName = "hades-command-test"

    run {
        command {
            name = "echo"
            args = ["Hello World!"]
            expectedOutput = "Hello World!"
        }

        command {
            name = "service"
            args = ["ssh", "status"]
            expectedOutput = "* sshd is running"
        }
    }
}

unittest "Command Test 2" {
    image = "bpalmer/ssh_test"
    port = "9090"
    containerName = "hades-command-test"

    run {
        command {
            name = "echo"
            args = ["World Hello!"]
            expectedOutput = "World Hello!"
        }
    }
}