# ClusterForge v0.4.0 release candidate notes

Candidate: `v0.4.0-rc.1`

v0.4 focuses on enterprise extensibility, fleet visibility, policy/compliance,
and operational maturity. It adds local plugins and template packs, policy engine
v2, multi-account/region metadata, Terraform Cloud configuration, upgrade and
fleet reporting, service/Backstage catalogs, compliance mappings, checksummed
offline manifests, backup/audit/secret workflows, and migration analysis.

All operational exporters and analyzers are local/read-only by default. Plugins
are disabled by default. Admission enforcement, production mutation and provider
credentials remain explicit operator choices.

## Upgrade notes

- Review `clusterforge.yaml` with `cf upgrade check` and `cf upgrade plan` before
  using new organization/account/region/profile/Terraform Cloud fields.
- Review generated Terraform and golden diffs; do not overwrite existing roots or
  state. Use saved production plans and documented import/moved workflows.
- Pin and review plugin/template sources. Existing projects do not enable plugins
  automatically.
- New app `platform` metadata defaults to Linux/amd64; Windows remains experimental.

## Excluded and experimental

No SaaS, automatic remediation/failover/global scheduler, plugin marketplace,
certification claim, federation, edge lifecycle automation, or production Windows
support is included. AKS/GKE/K3s/RKE2 and cloud-specific production readiness stay
limited to their documented evidence.

## Candidate warning

This is not a final release. Production cloud smoke tests have not been run for
this candidate. Review `RELEASE_CANDIDATE_V0.4.md` for actual local results,
skips, blockers and the release recommendation.
