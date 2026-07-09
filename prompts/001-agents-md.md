## Prompt 1 — AGENTS.md

```text
Create an AGENTS.md file for this repository.

The file should define coding and infrastructure rules for the ClusterForge project.

Include these sections:

1. Project overview
   - ClusterForge is a Terraform/OpenTofu framework for container orchestrators.
   - It supports Kubernetes, ECS, Nomad, and Docker targets.
   - The CLI is a wrapper/generator, not a replacement for Terraform.

2. Repository conventions
   - Terraform modules live under modules/.
   - Real environment compositions live under live/.
   - CLI source code lives under cli/.
   - Examples live under examples/.
   - Policies live under policies/.
   - Scripts live under scripts/.

3. Terraform module rules
   - Each module must have main.tf, variables.tf, outputs.tf, versions.tf, README.md.
   - Use typed variables.
   - Use descriptions on all variables and outputs.
   - Use validation blocks for important inputs.
   - Do not hardcode environment-specific values.
   - Do not configure providers inside reusable child modules unless unavoidable.
   - Use locals for naming and common labels/tags.
   - Do not store secrets in state unless absolutely unavoidable.
   - Avoid excessive dynamic blocks unless they improve readability.

4. Provider rules
   - Provider configuration belongs in live environment roots.
   - Every root configuration must declare required providers.
   - Pin provider versions using safe constraints.
   - Prefer explicit provider aliases when multiple clusters/cloud accounts are involved.

5. CLI rules
   - CLI is written in Go.
   - Use Cobra for command structure.
   - Use clear package boundaries under cli/internal.
   - CLI commands must not perform destructive actions without confirmation.
   - Production apply must require an existing plan file.
   - Destroy in production must be blocked by default.
   - CLI must generate readable Terraform files.

6. Testing and validation
   - Run terraform fmt -recursive.
   - Run terraform validate for examples where possible.
   - Run gofmt and go test for CLI.
   - Add shell scripts for repeatable validation.
   - Add GitHub Actions for CI.

7. Documentation rules
   - Every module needs README.md with usage example.
   - Root README must explain architecture, quickstart, and roadmap.
   - Examples must be copy-paste friendly.

8. Security rules
   - Never commit credentials, kubeconfig files, private keys, or tfstate files.
   - Do not generate real secrets into tfvars.
   - Use references to external secret stores where possible.
   - Production operations require explicit confirmation.

After creating AGENTS.md, also update README.md with a short note that contributors and AI agents must follow AGENTS.md.
```

---
