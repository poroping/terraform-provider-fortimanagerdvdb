---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bfd6"
description: |-
  Configure IPv6 BFD.
---

## fortios_router_bfd6
Configure IPv6 BFD.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `deviceparam` - Specifies the device to which the resource will be applied to on the FortiManager. If you want to inherit the configuration from the provider, do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `neighbor` - Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `interface` - Interface to the BFD neighbor. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip6_address` - IPv6 address of the BFD neighbor.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_bfd6 can be imported using:
```sh
terraform import fortios_router_bfd6.labelname {{mkey}}
```
