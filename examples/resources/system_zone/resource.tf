resource "fortimanagerdvdb_system_interface" "example" {
  allow_append = true
  vdomparam    = "root"

  name = "TESTACCINT2"
  type = "aggregate"
  vdom = "root"
}

resource "fortimanagerdvdb_system_zone" "example" {
  vdomparam = fortimanagerdvdb_system_interface.example.vdom

  name = "TESTZONE"
  interface {
    interface_name = fortimanagerdvdb_system_interface.example.name
  }
}