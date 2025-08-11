# Create a PBS vnode
resource "pbs_node" "this" {
  name = "node1.example.com"

  # Optional settings
  partition = "default"
}

# Import existing node:
# terraform import pbs_node.this node1.example.com
