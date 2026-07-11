# Nomad Consul integration

Renders Nomad's Consul integration settings. It does not install Consul or put
ACL tokens in Terraform state; use an environment-injected token reference.

```hcl
module "consul_config" {
  source  = "../../modules/platform/nomad/consul"
  address = "127.0.0.1:8500"
}
```
