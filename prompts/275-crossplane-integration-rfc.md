# Prompt 275 — Crossplane integration RFC

```text
Create Crossplane integration RFC.

Goal:
Evaluate Crossplane as an optional backend for platform resource provisioning.

Create:
- docs/rfcs/036-crossplane-integration.md
- docs/crossplane.md

Cover:

1. Goals
   - Kubernetes-native cloud resource provisioning
   - platform APIs through compositions
   - GitOps-compatible infrastructure
   - developer self-service resources

2. Non-goals
   - replacing Terraform entirely
   - mandatory Crossplane dependency
   - managing every cloud resource through Crossplane

3. Use cases
   - app teams request database/cache/queue through claims
   - platform team owns compositions
   - GitOps reconciles claims
   - Control Plane tracks claims/status

4. Terraform boundary
   - Terraform creates cluster and installs Crossplane
   - Crossplane provisions app-level cloud resources
   - Terraform can still provision foundation resources

5. Proposed modules
   - modules/platform/kubernetes/crossplane
   - modules/platform/kubernetes/crossplane-provider-aws
   - modules/platform/kubernetes/crossplane-provider-azure
   - modules/platform/kubernetes/crossplane-provider-gcp

6. CLI future
   - cf resource claim create postgres
   - cf resource claim list
   - cf resource claim status

7. Risks
   - ownership conflicts with Terraform
   - provider credentials
   - debugging complexity
   - CRD lifecycle
   - state split across systems

Do not implement code.
Update roadmap.
```
