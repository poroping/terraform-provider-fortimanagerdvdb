resource "fortimanagerdvdb_json_generic_api" "example" {
  path   = "/api/v2/cmdb/firewall/address"
  method = "GET"
}

output "example" {
  value = jsondecode(fortimanagerdvdb_json_generic_api.example.response)["results"][0]
}