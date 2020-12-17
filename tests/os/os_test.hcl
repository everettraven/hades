unittest "OS Check Test without Version" {
    image = "bpalmer/ssh_test"
    port = "9091"
    containerName = "hades-os-test"

    run {
        os {
            distributionID = "ubuntu"
        }
    }
}

unittest "OS Check Test with Version" {
    image = "bpalmer/ssh_test"
    port = "9091"
    containerName = "hades-os-test-version"

    run {
        os {
            distributionID = "ubuntu"
            version = "20.04"
        }
    }
}