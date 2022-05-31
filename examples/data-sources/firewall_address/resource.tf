data "fortimanagerdvdb_firewall_address" "example" {
  vdomparam = "root"

  name = "all"
}

output "example" {
  value = data.fortimanagerdvdb_firewall_address.example
}