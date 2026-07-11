# RFC 012: Web dashboard prototype

Status: proposal only. No web application or apply endpoint is implemented.

## Recommended MVP

A static, read-only UI imports one locally generated
`dashboard-data.json`. It shows environment inventory, cluster status, app and
service catalog, drift/policy summaries, heuristic cost warnings, audit events,
runbook links and release status. The CLI remains responsible for collecting
and redacting local data; a platform API may replace file import later.

Next.js static export is a reasonable option if component complexity warrants
it; a dependency-light static HTML/TypeScript prototype is preferable for the
first spike. The chosen UI must run without cloud credentials.

## Non-goals

- Direct apply/destroy/remediation buttons.
- Cloud credential, kubeconfig, state or secret storage.
- Replacing Git review, Terraform or existing CI workflows.
- Hosted multi-tenant SaaS, live fleet control or real-time guarantees.

## Data contract

The export should include schema version, generation timestamp and source
commit, followed by arrays/objects for environments, clusters, apps, services,
drift results, policy results, cost warnings, audit events, runbook metadata and
release status. Every record should include stable local identifiers, status,
source and freshness where applicable. Unknown/unavailable differs from healthy.

Fields containing raw Terraform values, credentials, tokens, state, plan
content, kubeconfigs, inline environment values, private log payloads or audit
details are prohibited. Export tests must scan serialized output for known
secret fixtures and enforce stable schema versioning.

## Security

Treat the JSON as potentially sensitive infrastructure metadata. Generate it
locally, use restrictive permissions, set short retention, never publish it as
an unrestricted Pages artifact and sanitize URLs/account/cluster identifiers
according to organizational policy. A browser must not fetch cloud APIs.

The first UI has no apply button because reliable mutation requires strong
authentication/authorization, plan/config digests, separation of duties,
backend locking, audit, cancellation and recovery. A superficial button would
bypass the safety model and imply guarantees the prototype cannot provide.

## Evolution

1. Define/test CLI export schema and fixtures.
2. Build local static import and empty/stale/error states.
3. Add accessible filtering and links to source/runbooks.
4. Evaluate authenticated read-only API from RFC 011.
5. Consider remote actions only through a separate security RFC.
