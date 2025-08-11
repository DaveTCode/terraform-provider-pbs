# Create a PBS hook
resource "pbs_hook" "this" {
  name = "my_hook"
  type = "site"

  # Optional settings
  enabled = true
}

# Import existing hook:
# terraform import pbs_hook.this my_hook
