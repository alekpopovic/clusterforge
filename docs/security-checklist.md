# Security checklist

Use this checklist for changes and release gates. A checked item means evidence
was reviewed for the specific environment; it is not a compliance attestation.

## Source and dependencies

- [ ] `git status` contains no credentials, state, plans, kubeconfigs, keys, or private data.
- [ ] New provider, module, chart, action, image, and asdf sources are trusted and version constrained.
- [ ] Dependency and lock-file changes are reviewed separately from generated noise.
- [ ] Secret scanning and IaC/config scanning completed, with exceptions documented.
- [ ] Generated Terraform and external template or policy packs received human review.

## Identity and state

- [ ] CI and operator credentials are short-lived and least privilege.
- [ ] Workloads use workload identity instead of node-wide or static credentials.
- [ ] Remote state is encrypted, access-controlled, versioned, and locked where supported.
- [ ] State and plan artifacts are excluded from source control and retained only as required.
- [ ] IAM policies, trust relationships, service accounts, and RBAC have no unexplained wildcards.

## Network and Kubernetes

- [ ] Public ingress, load balancers, security groups, firewall rules, and DNS records are intentional.
- [ ] TLS and DNS ownership/delegation are verified before exposure.
- [ ] Namespace, Pod Security, ResourceQuota/LimitRange, RBAC, and NetworkPolicy baselines are evaluated.
- [ ] Images use immutable references where practical and pass the required vulnerability policy.
- [ ] Admission controllers and Helm chart permissions are tested against the target Kubernetes version.

## Plan and production operations

- [ ] The saved production plan matches the reviewed commit and configuration.
- [ ] Deletes, replacements, privilege increases, and public exposure are explicitly reviewed.
- [ ] Production apply uses the reviewed plan file; destroy remains blocked unless deliberately authorized.
- [ ] Backups and restore runbooks exist for stateful or critical resources and have current test evidence.
- [ ] Drift checks and audit-log retention/permissions are configured without recording secrets.

## CI and releases

- [ ] Workflow permissions are minimal and untrusted pull requests cannot access deployment secrets.
- [ ] Release builds come from a reviewed tag and produce checksums and an SBOM.
- [ ] Published artifacts and documented installation paths were verified.
- [ ] Known unimplemented controls and failed/skipped tests are stated in release evidence.
- [ ] `SECURITY.md`, the threat model, and supported-version statements remain accurate.
