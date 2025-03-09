# Terraform Provider for Altair PBS

This repository contains a custom Terraform provider for provisioning aspects of a PBS system. Essentially providing an infrastructure as code interface
on top of `qmgr`

## Status

The following table illustrates the various resources that this provider eventually expects to include and the state of each

| Resource             | Create | Read | Update | Delete | Data Source |
|----------------------|--------|------|--------|--------|-------------|
| Queue                | y      | y    | y      | y      | y           |
| vNode                | x      | x    | x      | x      | x           |
| Custom Resource      | y      | y    | y      | y      | y           |
| Hooks                | y      | y    | y      | y      | y           |
| Server Attributes    | x      | x    | x      | x      | x           |
| Scheduler Attributes | x      | x    | x      | x      | x           |
| Hook files           | x      | x    | x      | x      | x           |

This repository will probably never provision jobs/reservations etc as those are deemed outside of the general "configuration of PBS" steps.

This repository will also never provision the VM/containers required to actually run PBS. That's typically handled by another layer of automation, 
whether cloud based or manual.