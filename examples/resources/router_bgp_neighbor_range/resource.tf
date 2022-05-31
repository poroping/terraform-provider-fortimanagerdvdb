resource "fortimanagerdvdb_router_bgp_neighbor_group" "example" {
  name      = "foobarrange"
  remote_as = 65000
}

resource "fortimanagerdvdb_router_bgp_neighbor_range" "example" {
  prefix         = "192.168.1.0/24"
  neighbor_group = fortimanagerdvdb_router_bgp_neighbor_group.example.name
}