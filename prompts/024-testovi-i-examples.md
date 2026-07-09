## Prompt 24 — Testovi i examples

```text
Add practical examples and validation coverage.

Examples to create or improve:

1. examples/aws-network
   - Uses modules/cloud/aws/network
   - Safe example variables
   - README with terraform commands

2. examples/aws-eks-minimal
   - Uses core tags, AWS network, EKS module
   - No real backend
   - README explains required AWS credentials

3. examples/kubernetes-basic-app
   - Uses workloads/kubernetes/app
   - Assumes Kubernetes provider is configured
   - README explains usage against existing cluster

4. examples/ecs-fargate-app
   - Uses ECS cluster and ECS service modules
   - Uses AWS network module
   - README explains what must exist or be created

5. examples/nomad-service
   - Uses Nomad service module
   - README explains Nomad provider endpoint

6. examples/docker-swarm-service
   - Uses Docker Swarm service module
   - README explains Docker provider requirements

Validation:
- Update scripts/validate.sh to:
  - run terraform fmt -recursive
  - run terraform validate in examples where possible
  - skip examples that require live credentials unless initialized properly
- Add comments explaining limitations.

Testing:
- For CLI:
  - Add unit tests for config.
  - Add unit tests for generator.
  - Add unit tests for policy parser.
- For Terraform:
  - If terraform test is practical, add basic tests for core modules.
  - Otherwise add simple examples that can be validated manually.

Do not add real cloud credentials, account IDs, or secrets.
```

---
