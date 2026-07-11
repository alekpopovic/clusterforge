# Platform add-on upgrades

Configure desired chart versions under `platform_versions`, then use
`cf platform versions <env>`, `cf platform upgrade plan <env>`, or `check`.
Add `--json` for automation.

The planner scans local generated Terraform for configured chart versions and
reports differences. It knows ingress-nginx, cert-manager, external-dns,
external-secrets, metrics-server, prometheus-stack, Loki, Argo CD, Kyverno,
Gatekeeper, Velero, and Argo Rollouts. Unknown components produce warnings.

CRD-bearing components are explicitly flagged because CRD schema conversion,
ordering, rollback, and stored-version compatibility require manual review.
The planner uses no internet lookup and never runs Helm, Terraform apply, or an
automatic upgrade.
