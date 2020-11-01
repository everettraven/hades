test "SSH Command Test" {
    image = "bpalmer/ssh_test"
    port = "9090"
    containerName = "hades-ssh-test"

    command {
        name = "echo"
        args = ["Hello World!"]
    }

    expectedOutput = "Hello World!"

}