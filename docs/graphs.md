# Environment dependency graphs

ClusterForge generates DOT text and never requires Graphviz or renders images.
The default logical graph is available before Terraform initialization:

```bash
cf graph dev
cf graph dev --stack platform
cf graph dev --format text
cf graph dev --output graphs/dev.dot
```

The logical graph models the known stack order:

```text
network -> cluster -> platform -> apps
```

Selecting a stack includes its upstream dependencies. This graph explains
ClusterForge orchestration order; it does not claim to discover every Terraform
resource or module edge.

Use Terraform/OpenTofu's native graph explicitly when the working directory is
initialized and providers/modules are available:

```bash
cf graph dev --terraform --format dot --output dev-terraform.dot
cf graph prod --terraform --stack platform --output prod-platform.dot
```

For stacked environments, `--terraform` requires exactly one `--stack` because
each stack has a separate Terraform root. The CLI runs only `terraform graph` or
`tofu graph` in that directory; it never plans or applies. DOT files can be
reviewed directly or rendered later with an independently installed Graphviz.
