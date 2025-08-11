terraform {
  required_providers {
    pbs = {
      source = "registry.terraform.io/hashicorp/pbs"
    }
  }
}

provider "pbs" {
  server   = var.pbs_server
  sshport  = var.pbs_port
  username = var.pbs_username
  password = var.pbs_password
}

variable "pbs_server" {
  description = "PBS server hostname"
  type        = string
  default     = "localhost"
}

variable "pbs_port" {
  description = "PBS SSH port"
  type        = string
  default     = "2222"
}

variable "pbs_username" {
  description = "PBS SSH username"
  type        = string
  default     = "root"
}

variable "pbs_password" {
  description = "PBS SSH password"
  type        = string
  default     = "pbs"
  sensitive   = true
}

variable "resource_prefix" {
  description = "Prefix for resource names to ensure uniqueness"
  type        = string
  default     = "test"
}
