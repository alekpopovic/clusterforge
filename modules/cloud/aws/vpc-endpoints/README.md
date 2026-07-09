# cloud/aws/vpc-endpoints

## Purpose

Creates AWS gateway and interface VPC endpoints for private EKS and ECS
environments. VPC endpoints let workloads reach supported AWS services through
private networking instead of routing all traffic through NAT gateways.

## Status

Implemented.

## Gateway vs Interface Endpoints

Gateway endpoints, such as `s3`, attach to route tables and do not use endpoint
network interfaces. Interface endpoints, such as `ecr.api`, `ecr.dkr`, `logs`,
`sts`, `ec2`, `eks`, `secretsmanager`, and `ssm`, create elastic network
interfaces in subnets and require security groups.

Service names are constructed as:

```text
com.amazonaws.<region>.<service>
```

For example, in `us-east-1`, `ecr.api` becomes
`com.amazonaws.us-east-1.ecr.api`.

## Private Cluster Use

Private EKS nodes commonly need these endpoints when pulling images from ECR
without internet egress:

- `s3` gateway endpoint for ECR image layer storage.
- `ecr.api` interface endpoint for ECR API calls.
- `ecr.dkr` interface endpoint for Docker registry traffic.
- `logs` interface endpoint when workloads or agents write CloudWatch logs.
- `sts` interface endpoint for IAM role token exchange.

Add `ec2`, `eks`, `secretsmanager`, and `ssm` when workloads or platform
add-ons use those services privately.

## NAT Cost Reduction

VPC endpoints can reduce NAT Gateway data processing charges for AWS service
traffic. They do not replace all internet egress. Keep NAT or another egress
path when workloads need public internet access.

## Usage

```hcl
module "vpc_endpoints" {
  source = "../../../modules/cloud/aws/vpc-endpoints"

  name            = "clusterforge-prod"
  environment     = "prod"
  vpc_id          = module.network.vpc_id
  subnet_ids      = module.network.private_subnet_ids
  route_table_ids = module.network.private_route_table_ids

  allowed_security_group_ids = [module.eks.cluster_security_group_id]

  gateway_endpoints = ["s3"]
  interface_endpoints = [
    "ecr.api",
    "ecr.dkr",
    "logs",
    "sts",
  ]
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
