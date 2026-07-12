# Prompt 172 — Apply job execution MVP

```text
Implement apply job execution MVP with strict safeguards.

Goal:
Allow approved apply requests to be executed by a trusted runner.

Preconditions:
- successful plan request exists
- apply request approved
- policy checks passed
- environment is not blocked
- runner is allowed to execute apply
- apply job type explicitly enabled in runner config

Control Plane:
- create apply job after approval
- track status:
  - pending
  - claimed
  - running
  - succeeded
  - failed
  - canceled
- store sanitized logs
- store summary
- audit every state transition

Runner:
- support apply job only when:
  allowed_job_types includes apply
- run apply using plan file only if artifact strategy exists
- otherwise first implementation may run:
  terraform plan -out
  terraform apply planfile
  in same runner workspace
- never use auto-approve unless explicitly applying a saved plan file
- cleanup workspace after job

CLI:
- cf apply status <apply-id>
- cf apply logs <apply-id>

Safety:
- prod apply requires approval
- destroy operations in plan require allow_destroy approval flag
- failed apply must be clearly reported
- no retry by default

Tests:
- unapproved apply not executed
- approved apply creates job
- runner blocks apply if not allowed
- destroy plan blocked without allow flag
- audit events emitted

Docs:
- docs/control-plane-apply-workflow.md

Rules:
- Be conservative.
- Do not make apply button easy to misuse.
- Do not hide Terraform output, but redact secrets.
```
