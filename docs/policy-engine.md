# Policy engine v2

The ClusterForge policy engine provides a common finding model for repository,
environment, stack, plan, app, module, image, and state checks. Each policy has
an ID, severity, category, scope, action, remediation, and enabled state.

```bash
cf policy list
cf policy show CF-PROD-003
cf policy check
cf policy check prod --pack production
cf policy check prod --stack platform --format json
cf policy check --app api --format sarif
```

Actions are `advisory`, `warn`, or `block`. A blocking result exits with code 2
after writing findings. Table output is intended for people, JSON for
automation, and SARIF 2.1.0 for code-scanning ingestion.

Built-in checks cover tracked state and `.env` files, secret-looking app
literals, mutable production images, production plan/destroy configuration,
local production state, mutable module refs, unapproved public ingress and
LoadBalancer services, and wildcard IAM policy text. Static matching can
produce false positives; ambiguous network, module, and IAM findings therefore
default to warnings.

Policy results are guardrails, not a compliance certification. They do not
replace plan review, cloud IAM analysis, admission control, secret scanning, or
real-environment validation. External policy plugins may provide findings using
the same model, but the core engine does not automatically run plugins.
