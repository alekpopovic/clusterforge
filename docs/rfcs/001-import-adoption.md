# RFC 001: Import and Adoption

Goal: help teams adopt existing AWS, Kubernetes, ECS, Route53, and Terraform
resources without unsafe replacement.

Initial implementation is documentation only. Future CLI commands should read
metadata, generate candidate Terraform, and produce import plans without
applying changes.

Non-goals: automatic production import, state editing, or destructive
remediation.
