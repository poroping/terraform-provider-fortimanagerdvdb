resource "fortimanagerdvdb_system_interface" "example" {
  vdom = "BUTT"
  type = "loopback"
  name = "foobar123"
  ip   = "5.5.5.0/31"
}

resource "fortimanagerdvdb_router_ospf_ospf_interface" "example" {
  name         = fortimanagerdvdb_system_interface.example.name
  interface    = fortimanagerdvdb_system_interface.example.name
  network_type = "non-broadcast"
}

resource "fortimanagerdvdb_router_ospf_neighbor" "example" {
  ip = "5.5.5.1"

  depends_on = [
    fortimanagerdvdb_router_ospf_ospf_interface.example
  ]
}