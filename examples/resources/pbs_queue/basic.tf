# Create a PBS execution queue
resource "pbs_queue" "this" {
  name       = "work"
  queue_type = "Execution"

  # Optional settings
  enabled = true
}

# Import existing queue:
# terraform import pbs_queue.this work
