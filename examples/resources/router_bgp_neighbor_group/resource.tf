resource "fortimanagerdvdb_router_bgp_neighbor_group" "example" {
  name      = "foobar"
  remote_as = 65000
}