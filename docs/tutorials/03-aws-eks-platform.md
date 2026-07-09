# Tutorial 03: AWS EKS Platform

## Goal

Generate an EKS environment, apply with a saved plan, and bootstrap ingress.

## Prerequisites

- AWS CLI configured outside the repository
- `terraform` or `tofu`
- `kubectl`
- `helm`
- ClusterForge `cf`

This tutorial can create billable AWS resources.

## Commands

```bash
cf project init eks-demo
cf env create dev --cloud aws --orchestrator eks --region us-east-2
cf generate dev
cp live/dev/aws-eks/terraform.tfvars.example live/dev/aws-eks/terraform.tfvars
cf init dev
cf plan dev --out .cf/plans/dev.tfplan --risk-summary
cf apply dev --plan-file .cf/plans/dev.tfplan
```

## Generated Files

- `live/dev/aws-eks/*.tf`
- local ignored tfvars and plan files

## What Terraform Creates

- VPC and subnets
- NAT Gateway if enabled
- EKS control plane
- managed node group
- CloudWatch logs

## Validate

```bash
aws eks update-kubeconfig --region us-east-2 --name <cluster-name> --alias dev
kubectl get nodes
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx --namespace ingress-nginx --create-namespace
kubectl -n ingress-nginx get service ingress-nginx-controller
```

## Cleanup

```bash
helm uninstall ingress-nginx -n ingress-nginx || true
kubectl delete namespace ingress-nginx --ignore-not-found
cf destroy dev --allow-destroy
```

## Troubleshooting

Review IAM permissions, EKS version availability, node group events, security
groups, load balancers, ENIs, NAT gateways, and CloudWatch logs.
