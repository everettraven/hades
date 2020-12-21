title = "Hosts file test"

command {
    cmd = "echo"
    args = ["Hello Infrastructure Testing World!"]
    expectedOutput = "Hello Infrastructure Testing World!"
}

os {
    distributionID = "Ubuntu"
    version = "20.04"
}