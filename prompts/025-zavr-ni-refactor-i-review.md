## Prompt 25 — Završni refactor i review

```text
Perform a full repository review and hardening pass.

Review these areas:

1. Terraform module quality
   - Every module has main.tf, variables.tf, outputs.tf, versions.tf, README.md.
   - Variables have descriptions and types.
   - Outputs have descriptions.
   - Provider configs are not hidden in child modules.
   - No fake resources.
   - No hardcoded environment-specific values.
   - No secrets in examples.
   - terraform fmt -recursive passes.

2. CLI quality
   - gofmt passes.
   - go test ./... passes.
   - Commands have useful help text.
   - Errors are actionable.
   - Destructive commands require explicit confirmation or flags.
   - Production safety rules are enforced.
   - Generated Terraform is readable.

3. Docs quality
   - Root README explains quickstart.
   - docs/architecture.md matches current repo.
   - docs/cli.md documents commands.
   - docs/security.md documents secret and production rules.
   - Each implemented module has a usage example.

4. CI
   - Workflows are syntactically valid.
   - CI does not require real cloud credentials for basic checks.
   - Scripts use bash strict mode.

5. Security
   - .gitignore excludes tfstate, plans, secrets, env files and kubeconfigs.
   - No committed credentials.
   - No sample secret values that look real.
   - Production destroy blocked by default.

6. Developer experience
   - A new user can understand the repo from README.
   - A new user can run CLI tests.
   - A new user can generate dev AWS EKS files.
   - A new user can inspect generated Terraform.

Make improvements directly.

At the end, produce a report:
- What was fixed
- What remains TODO
- Which commands were run
- Which commands passed or failed
- Any risks or assumptions
```

---

## Bonus prompt — Implementation plan pre kodiranja

```text
Before writing code, create a detailed implementation plan for ClusterForge.

Use the existing repository state.

Plan must include:

1. Current repository assessment
   - What exists
   - What is missing
   - Any risky assumptions

2. Proposed milestone plan
   Milestone 1:
   - Repo skeleton
   - AGENTS.md
   - README
   - scripts

   Milestone 2:
   - Core Terraform modules
   - naming
   - tags
   - labels

   Milestone 3:
   - AWS network
   - EKS module
   - live/dev/aws-eks

   Milestone 4:
   - Kubernetes platform Helm modules
   - bootstrap

   Milestone 5:
   - Kubernetes workload app and cronjob modules

   Milestone 6:
   - ECS cluster and service modules

   Milestone 7:
   - Nomad and Docker modules

   Milestone 8:
   - Go CLI

   Milestone 9:
   - App manifest generator

   Milestone 10:
   - Policy/risk summary

   Milestone 11:
   - CI and docs

3. File-by-file change plan
   - Mention which directories and files will be created or edited.

4. Validation plan
   - Terraform fmt
   - Terraform validate where possible
   - Go tests
   - Shell scripts
   - CI workflows

5. Risks
   - Cloud credentials not available
   - Terraform validation limitations without initialized providers
   - EKS IAM/OIDC details may need iterative refinement
   - Helm chart versions should be pinned before production

Do not change files yet. Only produce the plan.
```

---

## Bonus prompt — Mala kontrolisana promena

```text
Task:
Implement only [SPECIFIC MODULE OR FEATURE].

Scope:
- Modify only these paths:
  - [path 1]
  - [path 2]
  - [path 3]

Do not modify:
- unrelated modules
- unrelated docs
- CI workflows
- generated examples unless required

Requirements:
- Follow AGENTS.md.
- Keep Terraform readable.
- Do not hide provider configuration in child modules.
- Do not introduce secrets.
- Add or update README for changed module.
- Add example if useful.
- Run formatting and relevant tests.

Definition of done:
- Files are created/updated.
- Formatting passes.
- Tests or validation attempted.
- Final response includes:
  - changed files
  - commands run
  - pass/fail result
  - TODOs
```

---
