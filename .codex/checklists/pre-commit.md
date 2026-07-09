# Pre-Commit Checklist

- `git status --short --branch` reviewed.
- Scope matches the user request.
- No unrelated user changes reverted.
- Formatting run where relevant.
- Tests or validation run, or skip reason documented.
- No credentials, kubeconfigs, tfstate, tfplan, private keys, or real tfvars.
- Commit message summarizes the change plainly.
