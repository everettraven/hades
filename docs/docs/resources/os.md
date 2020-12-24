---
id: os
title: os
---

The os resource is used to test the operating system of the machine. There are two parameters:

**distributionID** - This parameter expects the distributor ID of the operating system you are expecting the machine to have. Not case sensitive.

**version _(optional)_** - This parameter will be used to check the version of the operating system to this specified version. Not case sensitive.

## Examples

Without version:
```hcl
os {
    distributionID = "Ubuntu"
}
```

This test will pass with any operating system with the distribution ID of Ubuntu

With version:
```hcl
os {
    distributionID = "Ubuntu"
    version = "20.04"
}
```

This test will pass only with the operating system Ubuntu version 20.04.

