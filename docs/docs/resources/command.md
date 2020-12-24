---
id: command
title: command
---

The command testing resource allows you to run any command on the system being tested. The command resource has three different parameters:

**cmd** - This parameter specifies the actual command you would like to run. Case sensitive.

**args _(optional)_** - This parameter is used to pass arguments to the command specified in the cmd parameter, but is optional as you can pass arguments in the cmd parameter as well.

**expectedOutput** - This parameter is used to test the output from running the command specified in the cmd parameter to ensure it matches the expected output. The text in this is case sensitive and accepts standard string special characters such as `\n`.

## Examples

Simple echo command with args parameter:
```hcl
command {
    cmd = "echo"
    args = ["hello world"]
    expectedOutput = "hello world"
}
```

Multiple arguments in args parameter:
```hcl
command {
    cmd = "head"
    args = ["-1", "/etc/os-release"]
    expectedOutput = "NAME=\"Ubuntu\""
}
```

Without args parameter:
```hcl
command {
    cmd = "lsb_release -a | grep -i Release"
    expectedOutput = "Release:\t20.04"
}
```