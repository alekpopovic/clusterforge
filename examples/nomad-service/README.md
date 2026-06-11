# nomad-service

Example root module for `modules/workloads/nomad/service`.

This example renders a Nomad service job using the Docker driver and optional
service registration metadata.

## Provider

The Nomad provider uses `var.nomad_address`, defaulting to:

```text
http://127.0.0.1:4646
```

Run against a local or test Nomad cluster:

```bash
terraform init
terraform validate
terraform plan
```

Service discovery usually requires Consul or another integration configured in
your Nomad environment. This example does not create Consul or Nomad servers.
