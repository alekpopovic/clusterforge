## Prompt 11 — Kubernetes cronjob modul

```text
Implement modules/workloads/kubernetes/cronjob.

Purpose:
Deploy a Kubernetes CronJob using Terraform.

Inputs:
- name
- namespace
- create_namespace
- image
- schedule
- command
- args
- env
- secret_env
- labels
- annotations
- concurrency_policy default "Forbid"
- successful_jobs_history_limit default 3
- failed_jobs_history_limit default 3
- restart_policy default "OnFailure"
- resources object same style as app module
- image_pull_policy default "IfNotPresent"
- image_pull_secrets list(string) default []

Resource:
- kubernetes_namespace_v1 optional
- kubernetes_cron_job_v1

Validation:
- name, namespace, image, schedule non-empty.
- concurrency_policy must be Allow, Forbid, or Replace.
- restart_policy must be OnFailure or Never.

Outputs:
- name
- namespace
- cronjob_name
- labels

README:
- Basic example.
- Example with env and secret refs.
- Explain secret handling.

Run terraform fmt -recursive.
```

---
