---
layout: "fortimanager-devicedb"
page_title: "Provider: FortiManager-DeviceDB"
sidebar_current: "docs-fortimanager-devicedb-index"
description: |-
  The fortimanager-devicedb provider interacts with a FortiManager Device Database.
---

# fortimanager-devicedb Provider

## Why?
Official FortiManager provider doesn't give easy option to do device level configuration. This Provider is a spin off of my FortiOS provider with a wrapper on the SDK to interact with the FortiManager.

This is a side project and a WIP. There are bugs.

### Example Usage

```hcl
provider "fortimanagerdvdb" {
  hostname       = "192.168.1.99"
  username       = "admin"
  password       = "admin"
  vdom           = "root"
  insecure       = "true"
  default_device = "tftest"
}

# Create a Static Route
resource "fortimanagerdvdb_router_static" "example" {
  dst       = "110.2.2.122/32"
  blackhole = "enable"
  # …
}
```

### VDOM

If the FortiGate unit is running in VDOM mode, the `vdom` configuration needs to be added.

Usage:

```hcl
provider "fortimanagerdvdb" {
  hostname        = "192.168.1.99"
  username        = "admin"
  password        = "admin"
  vdom            = "root"
  insecure        = "true"
  default_device = "tftest"
}

resource "fortimanagerdvdb_router_static" "test1" {
  dst       = "120.2.2.122/32"
  gateway   = "2.2.2.2"
  blackhole = "disable"
  distance  = "22"
  weight    = "3"
  priority  = "3"
  device    = "lbforvdomtest"
  comment   = "Terraform test"
}
```

By default, each resource inherits the provider's global vdom settings, but it can also set its own vdom through the `vdomparam` of each resource. See the `vdomparam` argument of each resource for details.

### Device

To target a specific device on the FortiManager, the `deviceparam` needs to be added to the resource.

Usage:

```hcl
provider "fortimanagerdvdb" {
  hostname     = "192.168.1.99"
  username     = "admin"
  password     = "admin"
  vdom         = "root"
  insecure     = "true"
}

resource "fortimanagerdvdb_router_static" "test1" {
  deviceparam = "fortigate1"

  dst       = "120.2.2.122/32"
  gateway   = "2.2.2.2"
  blackhole = "disable"
  distance  = "22"
  weight    = "3"
  priority  = "3"
  device    = "lbforvdomtest"
  comment   = "Terraform test"
}
```

By default, each resource inherits the provider's global device setting `default_device`.

### Argument Reference

The following arguments are supported:

* `hostname` - The hostname/IP address of the FortiOS to be connected. If omitted, `TF_FMG_DEVICEDB_ACCESS_HOSTNAME` environment variable is used.

* `token` - (Optional) The FortiOS API access token. If omitted, `TF_FMG_DEVICEDB_ACCESS_TOKEN` environment variable is used.

* `insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `TF_FMG_DEVICEDB_INSECURE` environment variable is used. Default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. If omitted, `TF_FMG_DEVICEDB_CA_CABUNDLE` environment variable is used.

* `vdom` - (Optional) Default VDOM for API calls. If omitted, `TF_FMG_DEVICEDB_VDOM` environment variable is used.

* `default_device` - (Optional) Default device to target on FortiManager. If omitted, `TF_FMG_DEVICEDB_DEFAULT_DEVICE` environment variable is used.


## Found a bug or want to Contribute?
Help is welcome!

See `Guides` -> `Debugging Provider` for information on how to provide good debug logs.

[GitHub](https://github.com/poroping/terraform-provider-fortimanagerdvdb/issues).

## Versioning

The provider was generated from schemas covering 6.2, 6.4 and 7.0. Acceptance testing is performed with the latest GA release.