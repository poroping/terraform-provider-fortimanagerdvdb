data "fortimanagerdvdb_firewall_addresslist" "example" {
  vdomparam = "root"

  filter = "name!=fakeyjakey"
}

output "example" {
  value = data.fortimanagerdvdb_firewall_addresslist.example.namelist
}