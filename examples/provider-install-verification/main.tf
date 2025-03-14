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

import {
  id = "pbs"
  to = pbs_server.pbs
}

resource "pbs_server" "pbs" {
  name = "pbs"
  default_chunk = {
    "ncpus" = 1
  }
  default_queue            = pbs_queue.newq.name
  eligible_time_enable     = false
  log_events               = 511
  mail_from                = "adm"
  mailer                   = "/usr/sbin/sendmail"
  max_array_size           = 10000
  max_concurrent_provision = 5
  max_job_sequence_id      = 9999999
  node_fail_requeue        = 310
  pbs_license_linger_time  = 31536000
  pbs_license_max          = 2147483647
  pbs_license_min          = 0
  power_provisioning       = false
  query_other_jobs         = true
  resources_default = {
    "ncpus" = 1
  }
  resv_enable         = true
  scheduler_iteration = 600
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
    ncpus    = 1
    nodect   = 1
    nodes    = 1
    walltime = "02:00:00"
  }
}

# import {
#   id = "pbs"
#   to = pbs_node.pbs
# }

resource "pbs_node" "pbs" {
  name = "pbs"
  port = 15002
  resources_available = {
    mem                   = "10mb"
    hni_pkts_recv_by_tc_0 = "10mb"
  }
  resv_enable = true
}

locals {
  hni_resources = [
    "hni_pkts_recv_by_tc_0",
    "hni_pkts_sent_by_tc_0",
  ]
}

resource "pbs_resource" "hni_pkts_resource" {
  for_each = toset(local.hni_resources)
  name     = each.value
  type     = "size"
  flag     = "hf"
}

resource "pbs_hook" "ahook" {
  debug       = false
  enabled     = true
  name        = "ahook"
  user        = "pbsadmin"
  event       = "execjob_begin"
  order       = 1
  type        = "site"
  alarm       = 30
  fail_action = "none"
}

output "wq" {
  value = data.pbs_queue.workq_queue.priority
}
