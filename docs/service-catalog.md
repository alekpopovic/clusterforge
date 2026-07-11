# Service catalog

`service-catalog.yaml` is a local metadata source for service owners, tier,
lifecycle, source/image locations, environment URLs, dependencies, dashboards
and runbook names. It contains no credentials and calls no remote API.

Use `cf service list`, `show`, `validate`, `export --format json|markdown`, and
`graph --format dot`. App manifests may set `service: api` to associate a
workload with an entry. Backstage generators and runbook tooling can consume the
same ownership and operational links.

Catalog URLs and repository/image identifiers are references only. Never store
tokens, passwords, connection strings or secret query parameters. Review owner
and lifecycle changes like code, and validate dependency graphs before using
them for operational decisions.
