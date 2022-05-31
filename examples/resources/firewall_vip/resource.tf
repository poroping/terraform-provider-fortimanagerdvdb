resource "fortimanagerdvdb_firewall_vip" "example" {
  name    = "foobar"
  extip   = "1.1.1.1"
  extintf = "any"

  mappedip {
    range = "192.168.1.1-192.168.1.1"
  }
}

resource "fortimanagerdvdb_firewall_address" "example2" {
  name = "example.com"
  type = "fqdn"
  fqdn = "example.com"  
}

resource "fortimanagerdvdb_firewall_vip" "example2" {
  extip       = "192.0.2.1/24"
  name        = "examplevip"
  extintf     = "any"
  mapped_addr = fortimanagerdvdb_firewall_address.example2.name
  type        = "fqdn"
}