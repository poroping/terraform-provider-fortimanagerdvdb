resource "fortimanagerdvdb_system_interface" "example" {
  allow_append = true
  vdomparam    = "BUTT"

  name = "TESTACCINT"
  type = "loopback"
  vdom = "BUTT"
}

resource "fortimanagerdvdb_vpnssl_settings" "example" {
  vdomparam = "BUTT"

  default_portal = "full-access"

  source_address {
    name = "all"
  }

  source_interface {
    name = fortimanagerdvdb_system_interface.example.name
  }
}