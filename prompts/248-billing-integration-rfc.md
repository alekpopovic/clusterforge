# Prompt 248 — Billing integration RFC

```text
Create billing integration RFC for future SaaS or managed offering.

Goal:
Design billing boundaries without implementing payments.

Create:
- docs/rfcs/027-billing-integration.md
- docs/control-plane/billing.md

Cover:

1. Non-goals for current project
   - no payment processor implementation
   - no invoicing
   - no subscription enforcement
   - no public SaaS launch

2. Possible future billing dimensions
   - active projects
   - active environments
   - active runners
   - job volume
   - artifact storage
   - audit retention
   - preview environments
   - enterprise features

3. Usage metering dependency
   - usage_events
   - daily rollups
   - monthly rollups

4. Billing account model
   - organization billing account
   - plan tier
   - quota mapping
   - trial mode
   - suspended mode

5. Safety
   - never block emergency access unexpectedly
   - support grace periods
   - read-only mode instead of destructive disablement
   - clear admin notifications

6. Data privacy
   - no secrets in usage data
   - minimal metadata
   - exportable reports
   - retention settings

7. Future APIs
   - GET /api/v1/billing/account
   - GET /api/v1/billing/usage
   - GET /api/v1/billing/invoices
   - POST /api/v1/billing/portal

Do not implement code.
Update roadmap and backlog.
```
