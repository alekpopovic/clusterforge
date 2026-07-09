## Prompt 23 — Dokumentacija

```text
Improve documentation for ClusterForge.

Update root README.md with:

1. Project title and one-line pitch:
   ClusterForge is a Terraform/OpenTofu framework and CLI for Kubernetes, ECS, Nomad and Docker-based container platforms.

2. Why it exists:
   - Standardize infrastructure layout.
   - Avoid copy-paste Terraform chaos.
   - Keep Terraform readable.
   - Support multiple container orchestrators through adapters.
   - Provide a CLI for project generation and safe workflows.

3. Architecture:
   - Foundation layer
   - Orchestrator layer
   - Platform layer
   - Workload layer

4. Repository layout:
   Include tree with modules/, live/, examples/, cli/, policies/, scripts/.

5. Quickstart:
   - Install Terraform or OpenTofu.
   - Build CLI:
     cd cli && go build -o cf .
   - Initialize project:
     cf project init demo
   - Create environment:
     cf env create dev --cloud aws --orchestrator eks --region eu-central-1
   - Generate:
     cf generate dev
   - Terraform:
     cf init dev
     cf plan dev

6. Example workflows:
   - AWS EKS
   - AWS ECS
   - Kubernetes app
   - App manifest render

7. Safety model:
   - No auto-apply in production.
   - Plan file required for production apply.
   - Destroy blocked by default in production.
   - Secrets not stored in tfvars.

8. Module development guide:
   - How to add a new module.
   - Required files.
   - README requirements.
   - Inputs and outputs.

9. CLI development guide:
   - Go structure.
   - Tests.
   - Commands.

10. Roadmap:
   Phase 1: AWS EKS
   Phase 2: Kubernetes platform
   Phase 3: Kubernetes app workloads
   Phase 4: ECS
   Phase 5: CLI hardening
   Phase 6: Nomad
   Phase 7: Docker Swarm
   Phase 8: AKS/GKE/K3s/RKE2

Also create docs/:
- docs/architecture.md
- docs/module-conventions.md
- docs/cli.md
- docs/security.md
- docs/roadmap.md

Keep docs practical and not too marketing-heavy.
```

---
