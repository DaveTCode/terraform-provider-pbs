# Terraform Provider for Altair PBS

This repository contains a custom Terraform provider for provisioning aspects of a PBS system. Essentially providing an infrastructure as code interface
on top of `qmgr`

## Status

The following table illustrates the various resources that this provider eventually expects to include and the state of each

| Resource | Create | Read | Update | Delete | Data Source |
|----------|--------|------|--------|--------|-------------|
| Queue    | x      | x    | x      | x      | y           |
| vNode    | x      | x    | x      | x      | x           |
| Custom Resource | x | x | x | x |