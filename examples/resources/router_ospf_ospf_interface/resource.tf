resource "fortimanagerdvdb_system_interface" "example" {
  vdom = "BUTT"
  type = "loopback"
  name = "foobar234"
}

resource "fortimanagerdvdb_router_ospf_ospf_interface" "example" {
  name      = fortimanagerdvdb_system_interface.example.name
  interface = fortimanagerdvdb_system_interface.example.name
}