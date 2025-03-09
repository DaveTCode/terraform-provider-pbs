terraform {
  required_providers {
    pbs = {
      source = "registry.terraform.io/hashicorp/pbs"
    }
  }
}

provider "pbs" {
  server   = "localhost"
  sshport  = 2222
  username = "root"
  password = "pbs"
}

data "pbs_queue" "workq_queue" {
  name = "workq"
}

resource "pbs_queue" "newq" {
  name        = "newq"
  queue_type  = "Execution"
  enabled     = true
  started     = true
  priority    = 200
  max_running = 11
}

import {
  id = "test"
  to = pbs_queue.test
}

resource "pbs_queue" "test" {
  name       = "test"
  queue_type = "Execution"
  enabled    = true
  started    = true
  resources_default = {
    ncpus   = 1
    nodect  = 1
    nodes   = 1
    walltime = "02:00:00"
  }
}

locals {
  hni_resources = [
    "hni_pkts_recv_by_tc_0",
    "hni_pkts_sent_by_tc_0",
  ]
}

resource "pbs_resource" "hni_pkts_resource" {
  for_each = toset(local.hni_resources)
  name = each.value
  type = "size"
  flag = "hf"
}

output "wq" {
  value = data.pbs_queue.workq_queue.priority
}
