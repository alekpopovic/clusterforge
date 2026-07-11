## Summary

Describe the user-visible outcome and the reason for the change.

## Validation

List commands run and their results. Note any checks that could not run.

## Checklist

- [ ] Tests were added or updated where behavior changed.
- [ ] Documentation was updated where user-facing behavior changed.
- [ ] `terraform fmt -recursive` was run for Terraform changes.
- [ ] `go test ./...` was run from `cli/` for Go changes.
- [ ] No credentials, secrets, kubeconfigs, state, plan files, or private data are included.
- [ ] Production safety and destructive-operation behavior were considered.
- [ ] Each changed Terraform module README was updated.
- [ ] Generated Terraform remains readable and reviewable.

## Production impact

Describe rollout, rollback, compatibility, IAM, state, and cost implications, or write `None`.
