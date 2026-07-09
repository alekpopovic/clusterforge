---
title: Smoke Test Cleanup
permalink: /smoke-tests/cleanup/
---

# Smoke Test Cleanup

Use this checklist after every real-cloud or real-cluster smoke test. A smoke
test is not complete until Terraform/OpenTofu resources and manually created
test resources have been removed or explicitly documented as stuck.

Never commit credentials, kubeconfigs, state files, plan files, private keys, or
unredacted cloud identifiers while collecting cleanup evidence.

## General Cleanup Order

1. Delete demo applications and workload namespaces.
2. Delete platform bootstrap components such as ingress controllers,
   cert-manager, metrics, logging, and GitOps tools.
3. Destroy workload roots before cluster or network roots.
4. Destroy cluster roots before shared network roots.
5. Verify the cloud console or CLI shows no unexpected residual resources.
6. Remove local kubeconfig contexts and local generated test files.
7. Record redacted cleanup evidence.

## Kubernetes Cleanup

```bash
kubectl delete namespace clusterforge-smoke --ignore-not-found
helm list -A
helm uninstall ingress-nginx -n ingress-nginx || true
kubectl delete namespace ingress-nginx --ignore-not-found
kubectl get all -A
```

If namespace deletion is stuck, inspect finalizers:

```bash
kubectl get namespace <namespace> -o yaml
kubectl get apiservices
```

Document stuck resources before manual finalizer changes. Do not force-remove
finalizers in shared clusters without an owner review.

## AWS EKS Cleanup

Destroy through ClusterForge first:

```bash
cf destroy "<smoke-env>" --allow-destroy
```

Then check for residual resources in the smoke test region:

```bash
aws eks list-clusters --region "<region>"
aws elbv2 describe-load-balancers --region "<region>"
aws elbv2 describe-target-groups --region "<region>"
aws ec2 describe-network-interfaces --region "<region>" \
  --filters "Name=description,Values=*ELB*"
aws ec2 describe-nat-gateways --region "<region>" \
  --filter "Name=state,Values=available,pending,deleting"
aws ec2 describe-addresses --region "<region>"
aws logs describe-log-groups --region "<region>" \
  --log-group-name-prefix "/aws/eks/"
```

Manual checklist for stuck resources:

- Load balancers created by Kubernetes services or ingress controllers
- Target groups and listeners
- Security groups attached to load balancers or node groups
- ENIs attached to load balancers, NAT gateways, nodes, or VPC endpoints
- NAT gateways and Elastic IPs
- EBS volumes created by test workloads
- CloudWatch log groups
- IAM roles and policies created for the smoke cluster
- Route53 records, if DNS was tested

## AWS ECS Cleanup

Destroy optional workload roots first. Then destroy the ECS environment:

```bash
cf destroy "<smoke-env>" --allow-destroy
```

Check for residual resources:

```bash
aws ecs list-clusters --region "<region>"
aws ecs list-services --region "<region>" --cluster "<cluster-name>"
aws elbv2 describe-load-balancers --region "<region>"
aws ec2 describe-vpcs --region "<region>" \
  --filters "Name=tag:Environment,Values=<smoke-env>"
aws logs describe-log-groups --region "<region>"
```

Manual checklist for stuck resources:

- ECS services that still have desired tasks
- running or stopped ECS tasks
- ALBs, target groups, listeners, and listener rules
- security groups referenced by ECS services or ALBs
- ENIs attached to Fargate tasks or load balancers
- NAT gateways and Elastic IPs
- CloudWatch log groups
- IAM roles and policies created for task execution
- Route53 records, if DNS was tested

## Local Cleanup

Remove local files that should never be committed:

```bash
rm -f live/<env>/aws-eks/terraform.tfvars
rm -f live/<env>/aws-ecs/terraform.tfvars
find live/<env> -name .terraform.lock.hcl -type f -delete
find live/<env> -name .terraform -type d -prune -exec rm -rf {} +
find live/<env> -name .cf -type d -prune -exec rm -rf {} +
kubectl config delete-context "<context-alias>" || true
```

Do not delete evidence that maintainers intentionally kept in an ignored or
external evidence folder. Evidence committed to the repository must be redacted
and should not include account IDs, credentials, kubeconfigs, plan files, or
state.

## Cleanup Evidence

| Item | Evidence path or note |
| --- | --- |
| Demo app deleted | `<path>` |
| Platform components deleted | `<path>` |
| Cluster or ECS environment destroyed | `<path>` |
| VPC deleted | `<path>` |
| Load balancers checked | `<path>` |
| ENIs checked | `<path>` |
| NAT gateways and Elastic IPs checked | `<path>` |
| CloudWatch logs checked | `<path>` |
| Route53 checked, if applicable | `<path or not tested>` |
| Local files removed | `<path>` |
| Cleanup owner | `<name or handle>` |
| Cleanup date | `<YYYY-MM-DD>` |
