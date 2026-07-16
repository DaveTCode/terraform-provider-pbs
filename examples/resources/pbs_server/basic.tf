# Import and manage an existing PBS server
resource "pbs_server" "this" {
  name = "pbs"
  # Example managed attributes
  acl_users = "admin,staff"

  # Attributes omitted from configuration keep their imported values (they are
  # not implicitly unset). To actively reset a scalar attribute back to its PBS
  # default, list its name here instead of removing it from configuration:
  # unset_attributes = ["default_queue"]
}

# Read server values via data source
data "pbs_server" "this" {
  name = "pbs"
}

output "acl_hosts_normalized" {
  value = data.pbs_server.this.acl_hosts_normalized
}
