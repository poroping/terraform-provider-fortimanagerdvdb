---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicecustom"
description: |-
  Get information on a fortios Configure custom Internet Services.
---

# Data Source: fortios_firewall_internetservicecustom
Use this data source to get information on a fortios Configure custom Internet Services.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Internet Service name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `deviceparam` - Specifies the device to which the resource will be applied to on the FortiManager. If you want to inherit the configuration from the provider, do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `name` - Internet Service name.
* `reputation` - Reputation level of the custom Internet Service.
* `entry` - Entries added to the Internet Service database and custom database.The structure of `entry` block is documented below.

The `entry` block contains:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `dst` - Destination address or address group name.The structure of `dst` block is documented below.

The `dst` block contains:

* `name` - Select the destination address or address group object from available options.
* `port_range` - Port ranges in the custom entry.The structure of `port_range` block is documented below.

The `port_range` block contains:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
