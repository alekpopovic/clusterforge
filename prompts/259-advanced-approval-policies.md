# Prompt 259 — Advanced approval policies

```text
Implement advanced approval policies.

Goal:
Allow organizations to define approval requirements based on risk, environment, resource changes, and policy results.

Policy examples:
approval_policies:
  prod:
    default_required_approvals: 1
    two_person_rule: true
    require_approval_for_destroy: true
    require_approval_for_replacements: true
    require_security_approval_for_public_ingress: true
    require_database_owner_approval_for_rds_changes: true
    require_platform_approval_for_cluster_changes: true

Approval dimensions:
- environment
- stack
- resource type
- risk level
- destroy operations
- replacement operations
- policy severity
- cost warning severity
- incident mode
- break-glass use

Roles:
- approver
- security_approver
- platform_approver
- database_approver
- incident_commander

Behavior:
- approval request includes required approval types
- approval cannot be completed until all required approvals exist
- self-approval blocked when two-person rule is enabled
- approval expires after configurable time

API:
- GET /api/v1/approval-policies
- POST /api/v1/approval-policies
- GET /api/v1/apply-requests/{id}/approval-requirements

CLI:
- cf approval requirements <apply-id>
- cf approval policy list
- cf approval policy validate

Tests:
- destroy requires extra approval
- database change requires database approver
- self-approval blocked
- expired approval invalid
- all requirements satisfied allows apply

Docs:
- docs/control-plane/advanced-approvals.md

Rules:
- Deny by default when approval requirements cannot be evaluated.
- Do not weaken existing prod approval behavior.
```
