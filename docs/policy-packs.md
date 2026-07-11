# Policy packs

Policy engine v2 combines built-in checks with pack intent under
`policies/packs/`. Select stricter production behavior explicitly:

```bash
cf policy check prod --pack production
```

The baseline pack keeps uncertain infrastructure matches as warnings. The
production pack blocks deterministic production safety violations such as a
local backend or disabled plan/destroy safeguards. Actions can be overridden
to `advisory`, `warn`, or `block` by integrations using the policy-engine API.

The existing baseline, production, Kubernetes-security, and AWS-security pack
directories document intended controls. README-only rules are guidance until
they are represented by executable findings. Packs and output do not claim
regulatory compliance, and false-positive suppressions require review.
