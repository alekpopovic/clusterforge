# Read-only fleet operations

Fleet commands inspect filtered cluster inventory and intentionally expose no
apply, destroy, promotion, or remediation action.

```bash
cf fleet list --json
cf fleet status --cloud aws
cf fleet doctor --environment dev
cf fleet policy check --status production
cf fleet drift --orchestrator eks
```

All commands accept `--environment`, `--cloud`, `--orchestrator`, and `--status`.
Use `--json` for automation; human output is a compact tab-separated summary.

## Status and doctor

`fleet status` verifies inventory paths and summarizes target metadata.
`fleet doctor` aggregates local per-cluster checks, including path and kubeconfig
reference readiness. A failing cluster is reported without hiding results for the
rest of the fleet. These checks do not contact cloud APIs.

## Policy inspection

`fleet policy check` reports whether configured production safety gates are
present. It does not claim that live workloads comply with admission, cloud, or
organizational policy. Use environment-specific plan and policy checks for deeper
evidence.

## Drift

`fleet drift` runs Terraform/OpenTofu plans with `-detailed-exitcode` for each
selected cluster or each stack in a stacked environment. It writes plan files
under each root's ignored `.cf/plans/` directory, never applies them, and reports
drift as a warning.

Failures are collected and the command continues to later clusters. Pass
`--fail-fast` to stop and return an error on the first failed plan. Drift may need
backend or provider credentials even though it is read-only. Plan files and state
can contain sensitive data and must never be committed.

Fleet commands are synchronous and have no daemon, scheduler, background
remediation, or implicit cloud login.
