# Release Checklist

Use this checklist for ClusterForge release candidates.

For v0.4 candidates, also review `ROADMAP_V0.4.md`, `RELEASE_PLAN_V0.4.md`,
`BACKLOG_V0.4.md`, `RELEASE_NOTES_V0.4.md`, and
`RELEASE_CANDIDATE_V0.4.md`.

## Version and Changelog

- Confirm `VERSION` contains the intended release version.
- Update `CHANGELOG.md` with Added, Changed, Security, and Known Limitations.
- Confirm release notes do not overstate cloud validation.
- Confirm known limitations are still accurate.
- Confirm `VERSION_MATRIX.md` reflects supported, tested, experimental, and
  deprecated versions.
- Confirm `SMOKE_TEST_MATRIX.md` does not claim unrun tests passed.

## Documentation

- Review `README.md` quickstart, safety model, roadmap, and local commands.
- Review required docs:
  - `docs/architecture.md`
  - `docs/cli.md`
  - `docs/app-manifest.md`
  - `docs/environments.md`
  - `docs/backends.md`
  - `docs/security.md`
  - `docs/gitops.md`
  - `docs/roadmap.md`
  - `docs/plugins.md`
  - `docs/template-pack-registry.md`
  - `docs/policy-engine.md`
  - `docs/aws-multi-account.md`
  - `docs/compliance/index.md`
  - `docs/air-gapped.md`
- Confirm module READMEs include useful examples.
- Review `MODULE_CATALOG.md` and confirm stability labels are honest.
- Confirm no example contains real credentials, account IDs, private keys, or
  secret values.

## Local Checks

Run from the repository root:

```bash
make fmt-check
make lint
make test
make validate
make security
make check-modules
cd cli && go build -o cf .
./cf version
./cf doctor
```

Record each command and its result in `FINAL_MVP_REPORT.md` or the release
notes.

## CLI Smoke Checks

When practical, run CLI smoke checks in a temporary directory:

```bash
./cli/cf version
./cli/cf project init demo
./cli/cf env create dev --cloud aws --orchestrator eks --region eu-central-1
./cli/cf generate dev
./cli/cf app add api --image ghcr.io/example/api:1.0.0 --port 8080
./cli/cf app validate api
./cli/cf app render api --env dev
./cli/cf doctor
```

Do not run `terraform apply` as a release smoke check.

## Terraform Review

- Confirm `terraform fmt -recursive` is clean.
- Confirm credential-free validation passes where possible.
- Confirm skip reasons are explicit for roots that require credentials,
  initialized providers, or remote backends.
- Confirm child modules do not configure providers.
- Confirm production examples do not include hardcoded backend buckets,
  credentials, or account-specific identifiers.

## Security Review

- Confirm `.gitignore` excludes Terraform state, plan files, kubeconfigs,
  private keys, `.env` files, and local build artifacts.
- Confirm `git status` does not show generated state, lock, cache, or secret
  files.
- Confirm Checkov, Trivy, and TFLint results are reviewed when installed.
- Confirm release artifacts have SHA256 checksums.
- Generate or explicitly defer SBOM evidence according to
  `docs/supply-chain-security.md`.
- Confirm production apply/destroy safety rules are documented.

## GitHub Release

- Confirm CI workflows pass on the release branch or tag.
- Create and push the tag only after local checks are reviewed:

```bash
git tag v0.4.1
git push origin v0.4.1
```

- Verify GitHub release artifacts are uploaded by the release workflow.
- Verify generated artifacts match the expected target matrix:
  - Linux amd64
  - Linux arm64
  - macOS amd64
  - macOS arm64
  - Windows amd64
- Verify `install.sh`, every per-binary `.sha256`, and `SHA256SUMS` are attached.
- Run the installer against the published pinned tag in a clean Linux/macOS
  environment and confirm `cf version` matches the tag.

## Post-Release

- Open follow-up issues for known limitations.
- Prioritize real AWS EKS/ECS integration validation.
- Pin recommended Helm chart versions for production examples.
- Expand provider-specific hardening guidance.
