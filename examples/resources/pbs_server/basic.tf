# Import and manage an existing PBS server
resource "pbs_server" "this" {
  name = "pbs"
  # Example managed attributes
  acl_users = "admin,staff"
}

# Read server values via data source
data "pbs_server" "this" {
  name = "pbs"
}

output "acl_hosts_normalized" {
  value = data.pbs_server.this.acl_hosts_normalized
}
