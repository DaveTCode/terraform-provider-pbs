# Configuration-based authentication with password
provider "pbs" {
  server   = "10.10.10.10"
  sshport  = 22
  username = "root"
  password = "password"
}

# Configuration-based authentication with SSH private key (using file function)
provider "pbs" {
  server          = "10.10.10.10"
  sshport         = 22
  username        = "root"
  ssh_private_key = file("~/.ssh/id_rsa")
}

# Configuration-based authentication with both methods (fallback)
provider "pbs" {
  server          = "10.10.10.10"
  sshport         = 22
  username        = "root"
  password        = var.pbs_password
  ssh_private_key = file(var.ssh_key_path)
}

# Environment-based authentication (recommended for production)
provider "pbs" {
  # All configuration read from environment variables:
  # PBS_SERVER, PBS_SSH_PORT, PBS_USERNAME
  # PBS_PASSWORD (for password auth) or PBS_SSH_PRIVATE_KEY (for key auth)
}

# Using variables for all configuration
provider "pbs" {
  server          = var.pbs_server
  sshport         = var.pbs_ssh_port
  username        = var.pbs_username
  password        = var.pbs_password        # Optional: for password auth
  ssh_private_key = var.pbs_ssh_private_key # Optional: for SSH key auth
}