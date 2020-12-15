unittest "SSH Command Test" {
    image = "bpalmer/ssh_test"
    port = "9090"
    containerName = "hades-ssh-test"

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