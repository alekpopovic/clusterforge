---
title: CLI JSON Output
permalink: /cli-json/
---

# CLI JSON Output

Selected ClusterForge commands support stable JSON output for scripts and CI.
Human-readable output remains the default.

Errors are still returned through normal command failures and stderr messages.
JSON output does not include secret values.

## Environment List

```bash
cf env list --json
```

```json
{
  "environments": [
    {
      "name": "dev",
      "cloud": "aws",
      "orchestrator": "eks",
      "region": "eu-central-1",
      "path": "live/dev/aws-eks"
    }
  ]
}
```

## App List

```bash
cf app list --json
```

```json
{
  "apps": ["api", "worker"]
}
```

## Doctor

```bash
cf doctor --json
```

```json
{
  "version": "dev",
  "commit": "unknown",
  "date": "unknown",
  "status": "warn",
  "checks": [
    {
      "name": "terraform binary",
      "status": "pass",
      "message": "terraform found"
    }
  ]
}
```

`status` is one of `pass`, `warn`, or `fail`.

## Policy Check

```bash
cf policy check prod --json
cf policy check prod --plan-file .cf/plans/prod.tfplan --json
```

```json
{
  "environment": "prod",
  "policies": {
    "require_plan_file_for_apply": true,
    "block_destroy_in_prod": true
  },
  "messages": ["production apply requires --plan-file"],
  "risk": "MEDIUM",
  "policy": "apply allowed only with plan file",
  "summary": {
    "creates": 1,
    "updates": 0,
    "deletes": 0,
    "replacements": 0,
    "no_ops": 0,
    "addresses": ["module.example"]
  }
}
```

## Plan Risk Summary

```bash
cf plan dev --risk-summary --json
cf plan dev --stack network --risk-summary --json
```

Terraform/OpenTofu plan output is written to stderr in JSON mode so stdout
remains parseable.

```json
{
  "environment": "dev",
  "stacks": [
    {
      "path": "live/dev/aws-eks",
      "plan_file": ".cf/plans/dev.tfplan",
      "risk": "LOW",
      "policy": "apply allowed",
      "summary": {
        "creates": 1,
        "updates": 0,
        "deletes": 0,
        "replacements": 0,
        "no_ops": 0,
        "addresses": ["module.example"]
      }
    }
  ]
}
```
