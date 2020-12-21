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