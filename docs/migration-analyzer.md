# Terraform migration analyzer

The migration analyzer statically inventories an existing Terraform repository
to estimate ClusterForge adoption work. It is read-only: it does not initialize
Terraform, load state contents, contact providers/cloud APIs, generate imports, or
change files.

```bash
cf migrate analyze --path ../infrastructure
cf migrate analyze --path ../infrastructure --json
cf migrate report --path ../infrastructure > migration-assessment.md
```

The report lists providers, roots, module sources, backend declarations, tfvars
and state filenames, environment/module layout, and counts for VPC, EKS, ECS,
Kubernetes and Helm resources. It suggests ClusterForge module families, risks,
migration stages and import/adoption notes.

Secret-like assignments are reported only as file, line, key and `[REDACTED]`;
their values are never included. State files are detected by name and never read.
Nevertheless, run the analyzer only where you are authorized to inspect source.
Treat paths, module sources and architecture output as sensitive.

Static matching is an estimate, not a Terraform parser, semantic equivalence
proof, plan, security scan, or migration guarantee. Dynamic expressions, generated
configuration, wrappers, remote state and cloud-side drift can be missed. Before
adoption, back up state, document ownership, pin providers/modules, compare saved
plans in non-production, and migrate small resource groups using reviewed
Terraform `import` or `moved` workflows. Never copy/edit state as a shortcut.
