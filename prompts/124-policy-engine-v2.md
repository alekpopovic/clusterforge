## Prompt 124 — Policy engine v2

```text
Implement ClusterForge policy engine v2.

Goal:
Unify built-in policies, policy packs, plan JSON checks, app manifest checks, and repository checks.

Create package:
- cli/internal/policyengine

Policy sources:
1. built-in policies
2. policy packs from policies/packs
3. project-local policies
4. external policy plugin results

Policy object:
- id
- title
- description
- severity
- category
- scope
- default_action:
  - advisory
  - warn
  - block
- remediation
- references optional
- enabled

Scopes:
- repository
- environment
- stack
- plan
- app
- module
- image
- state

CLI:
- cf policy list
- cf policy show <id>
- cf policy check
- cf policy check <env>
- cf policy check <env> --stack <stack>
- cf policy check --app <name>
- cf policy check --json
- cf policy check --format table|json|sarif

Output:
- human table
- JSON
- SARIF if practical for GitHub code scanning

Policies to implement:
- no tfstate committed
- no .env committed
- no plaintext secret-looking values in app manifests
- no latest image tag in prod
- prod apply requires plan file
- prod destroy blocked
- prod backend must not be local
- module source should not use main in prod
- public ingress requires explicit approval annotation
- LoadBalancer service in prod requires approval
- wildcard IAM policy warns or blocks based on policy pack

Tests:
- Built-in policy evaluation.
- JSON output.
- SARIF output if implemented.
- Production pack blocks expected issues.
- Advisory policies do not block.

Docs:
- docs/policy-engine.md
- docs/policy-packs.md

Rules:
- Do not overclaim compliance.
- Make severity and blocking behavior configurable.
- Keep false positives as warnings where uncertain.

Run:
- gofmt
- go test ./...
```


---
