unittest "Command Test - Ubuntu" {
    image = "bpalmer/ubuntu-base-ssh"
    port = "9090"
    containerName = "hades-command-test-ubuntu"

    run {
        command {
            cmd = "echo"
            args = ["Hello World!"]
            expectedOutput = "Hello World!"
        }
    }
}

unittest "Command Test - Fedora" {
    image = "bpalmer/fedora-base-ssh"
    port = "9090"
    containerName = "hades-command-test-fedora"

    run {
        command {
            cmd = "echo"
            args = ["World Hello!"]
            expectedOutput = "World Hello!"
        }
    }
}

unittest "Command Test - CentOS" {
    image = "bpalmer/centos-base-ssh"
    port = "9090"
    containerName = "hades-command-test-centos"

    run {
        command {
            cmd = "echo"
            args = ["World Hello!"]
            expectedOutput = "World Hello!"
        }
    }
}

unittest "Command Test - Alpine" {
    image = "bpalmer/alpine-base-ssh"
    port = "9090"
    containerName = "hades-command-test-alpine"

    run {
        command {
            cmd = "echo"
            args = ["World Hello!"]
            expectedOutput = "World Hello!"
        }
    }
}