# Nomad cluster configuration

Experimental configuration-rendering module for operator-managed Nomad hosts.
It outputs server/client JSON and optional client cloud-init, but provisions no
VMs, networking, TLS, ACLs, storage, upgrades, or quorum lifecycle.

```hcl
module "nomad" {
  source = "../../modules/orchestrators/nomad/cluster"
  name = "platform"
  environment = "prod"
  server_addresses = ["10.0.0.10:4647"]
}
```
