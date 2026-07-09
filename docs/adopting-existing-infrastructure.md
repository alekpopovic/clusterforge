# Adopting Existing Infrastructure

ClusterForge adoption should start read-only. Prefer data sources and generated
configuration review before importing production resources.

Adoption targets:

- existing VPCs
- existing EKS clusters
- existing Kubernetes clusters
- existing ECS clusters
- existing Route53 zones
- existing Terraform state
- non-Terraform infrastructure

Risks include accidental replacement, naming mismatch, missing tags, state
ownership conflicts, and provider diff noise. Never import production without
state backup and reviewed plans.
