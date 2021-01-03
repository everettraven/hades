unittest "OS Check Test without Version - Ubuntu" {
    image = "bpalmer/ubuntu-base-ssh"
    port = "9091"
    containerName = "hades-os-test-ubuntu"

    run {
        os {
            distributionID = "ubuntu"
        }
    }
}

unittest "OS Check Test with Version - Ubuntu" {
    image = "bpalmer/ubuntu-base-ssh"
    port = "9091"
    containerName = "hades-os-test-version-ubuntu"

    run {
        os {
            distributionID = "ubuntu"
            version = "20.04"
        }
    }
}

unittest "OS Check Test without Version - Fedora" {
    image = "bpalmer/fedora-base-ssh"
    port = "9091"
    containerName = "hades-os-test-fedora"

    run {
        os {
            distributionID = "fedora"
        }
    }
}

unittest "OS Check Test with Version - Fedora" {
    image = "bpalmer/fedora-base-ssh"
    port = "9091"
    containerName = "hades-os-test-version-fedora"

    run {
        os {
            distributionID = "fedora"
            version = "33"
        }
    }
}

unittest "OS Check Test without Version - CentOS" {
    image = "bpalmer/centos-base-ssh"
    port = "9091"
    containerName = "hades-os-test-centos"

    run {
        os {
            distributionID = "centos linux"
        }
    }
}

unittest "OS Check Test with Version - CentOS" {
    image = "bpalmer/centos-base-ssh"
    port = "9091"
    containerName = "hades-os-test-version-centos"

    run {
        os {
            distributionID = "centos linux"
            version = "8"
        }
    }
}

unittest "OS Check Test without Version - Alpine" {
    image = "bpalmer/alpine-base-ssh"
    port = "9091"
    containerName = "hades-os-test-alpine"

    run {
        os {
            distributionID = "alpine linux"
        }
    }
}

unittest "OS Check Test with Version - Alpine" {
    image = "bpalmer/alpine-base-ssh"
    port = "9091"
    containerName = "hades-os-test-version-alpine"

    run {
        os {
            distributionID = "alpine linux"
            version = "3.12.3"
        }
    }
}