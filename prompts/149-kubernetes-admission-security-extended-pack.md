## Prompt 149 — Kubernetes admission security extended pack

```text
Extend Kubernetes admission security support.

Goal:
Provide stronger but opt-in Kubernetes admission policy modules and docs.

Enhance:
- Kyverno module
- Gatekeeper module
- policy packs

Policy categories:
1. Images:
   - disallow latest tag
   - require registry allowlist
   - require digest in prod optional

2. Pod security:
   - disallow privileged
   - disallow hostPath
   - disallow hostNetwork
   - require runAsNonRoot
   - restrict capabilities

3. Resources:
   - require requests/limits advisory or enforce

4. Networking:
   - require network policy per namespace advisory

5. Secrets:
   - block env values matching secret-like patterns advisory

6. Ingress:
   - public ingress requires approval annotation

Docs:
- docs/kubernetes-admission-security.md

Examples:
- examples/kubernetes-kyverno-production-pack
- examples/kubernetes-gatekeeper-production-pack

Rules:
- Default mode must be audit/warn.
- Enforcement must be opt-in.
- Do not break workloads silently.
- Document rollout strategy.

Run:
- terraform fmt -recursive
```


---
