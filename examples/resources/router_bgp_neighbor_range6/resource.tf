resource "fortimanagerdvdb_router_bgp_neighbor_group" "example" {
  name      = "foobarrange6"
  remote_as = 65000
}

resource "fortimanagerdvdb_router_bgp_neighbor_range6" "example" {
  prefix6        = "2003::/48"
  neighbor_group = fortimanagerdvdb_router_bgp_neighbor_group.example.name
}