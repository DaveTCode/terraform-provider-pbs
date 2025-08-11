# Create a PBS custom resource
resource "pbs_resource" "this" {
  name = "myres"
  type = "string"

  # Optional settings
  flag = "q"
}

# Import existing resource:
# terraform import pbs_resource.this myres
