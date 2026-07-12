# Project scaffolding wizard

The v2 wizard creates a readable starter `clusterforge.yaml`, one environment
directory, and an optional demo app. It prints a complete summary before any
write and asks for confirmation. It never requests secret values, contacts a
cloud, runs Terraform, or applies infrastructure.

```bash
cf wizard
cf init
cf project init --wizard
cf env create --wizard
cf app add --wizard
```

`cf init <env>` retains its existing meaning and runs Terraform/OpenTofu init for
an existing environment; only argument-free `cf init` starts scaffolding. Env and
app commands already prompt interactively, and `--wizard` makes that intent
explicit.

For a quick local demo or CI fixture:

```bash
cf wizard --defaults --non-interactive
cf init --defaults --non-interactive
```

Non-interactive scaffolding without `--defaults` fails instead of guessing. The
default creates a local-kind dev project, local backend, and pinned demo image.
The generated environment uses the existing-Kubernetes templates because kind
is a local Kubernetes cluster, while cluster lifecycle remains an explicit
separate operation:

```bash
cf local create kind
cf generate dev
cf init dev
cf plan dev
```

If kind already exists, omit `cf local create kind`. Generation and planning do
not create or delete a local cluster.
Cloud targets generate configuration only; review account, region, identity,
networking, backend, add-ons, and policies before generation.

Interactive choices cover project/owner, AWS EKS/ECS, existing Kubernetes, local
kind, experimental AKS/GKE, dev/staging/prod, local/S3/Terraform Cloud backend,
add-on intent, demo app, and safe production policies. Add-ons are shown in the
summary but remain explicit follow-up module choices; the wizard does not install
them. Backend placeholders are non-secret and must be replaced with reviewed
identifiers. CI must always use `--non-interactive`.
