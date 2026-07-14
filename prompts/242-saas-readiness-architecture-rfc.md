# Prompt 242 — SaaS-readiness architecture RFC

```text
Create the SaaS-readiness architecture RFC for ClusterForge.

Goal:
Assess what would be required to evolve the self-hosted Control Plane into a SaaS-capable architecture later, without implementing SaaS now.

Create:
- docs/rfcs/025-saas-readiness.md
- docs/control-plane/saas-readiness.md

Cover:

1. Goals
   - strong tenant isolation
   - scalable API
   - scalable runners
   - usage metering
   - organization onboarding
   - audit evidence export
   - regional deployment
   - secure artifact storage
   - customer data deletion
   - data residency model

2. Non-goals
   - public SaaS launch
   - billing implementation
   - marketplace
   - storing customer cloud credentials
   - running untrusted arbitrary code
   - automatic remediation by default

3. Tenant model
   - organization
   - workspace
   - project
   - environment
   - runner pool
   - artifact namespace
   - audit namespace

4. Trust boundaries
   - user browser
   - CLI
   - API
   - database
   - artifact storage
   - runner
   - Git provider
   - cloud provider
   - Kubernetes clusters

5. Data classification
   - public docs
   - non-sensitive metadata
   - sensitive metadata
   - audit events
   - plan summaries
   - raw plan files
   - logs
   - tokens
   - secrets, which must not be stored

6. SaaS blockers
   - multi-tenant authorization bugs
   - runner isolation
   - artifact sensitivity
   - customer-managed credentials
   - data deletion
   - regional storage
   - operational monitoring
   - incident response
   - compliance evidence

7. Recommended path
   Phase 1:
   - self-hosted enterprise
   Phase 2:
   - single-tenant managed
   Phase 3:
   - limited multi-tenant private SaaS
   Phase 4:
   - broader SaaS only after security review

8. Required future features
   - tenant isolation test suite
   - per-tenant quotas
   - usage metering
   - SCIM
   - SSO hardening
   - artifact encryption
   - immutable audit log
   - data residency
   - customer deletion workflow
   - abuse prevention
   - support tooling

Include:
- Mermaid deployment diagram
- Mermaid trust boundary diagram
- risk matrix
- migration considerations

Do not implement code in this prompt.
Update:
- ROADMAP_V0.7.md if it exists
- docs/roadmap.md
```
