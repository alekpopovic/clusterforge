# Real Cloud Test Workflow

Real cloud tests are manual and opt-in. They are never part of default CI.

1. Read `docs/smoke-tests/` or `docs/testing-integration.md`.
2. Confirm the account or cluster is disposable.
3. Confirm credentials are configured outside the repository.
4. Set required opt-in variables.
5. Save redacted evidence only.
6. Run cleanup before marking the test complete.

Never commit tfvars, tfstate, tfplan, kubeconfigs, credentials, or unredacted
cloud identifiers.
