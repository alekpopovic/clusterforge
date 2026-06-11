# cloud/aws/network

## Purpose

Creates a production-friendly AWS VPC network suitable for EKS and ECS roots.
The module creates a VPC, internet gateway, public and private subnets, route
tables, optional NAT gateways, and default routes.

Provider configuration belongs in the root module. This module declares the
AWS provider requirement but does not configure the provider.

## Kubernetes Compatibility

Public subnets receive `kubernetes.io/role/elb = "1"` by default. Private
subnets receive `kubernetes.io/role/internal-elb = "1"` by default. Pass
cluster-specific subnet discovery tags with `public_subnet_tags` and
`private_subnet_tags`, for example `kubernetes.io/cluster/my-cluster`.

## NAT Cost

NAT gateways are useful for private subnet egress, but they have hourly and
data processing costs. `single_nat_gateway = true` creates one shared NAT
gateway to reduce cost. Set it to `false` for one NAT gateway per availability
zone when higher availability is worth the additional cost.

## Inputs

| Name | Type | Default | Description |
| --- | --- | --- | --- |
| `name` | `string` | n/a | Name prefix for VPC networking resources. |
| `environment` | `string` | n/a | Environment name for tagging and resource names. |
| `cidr_block` | `string` | n/a | CIDR block for the VPC. |
| `availability_zones` | `list(string)` | n/a | Availability zones used for public and private subnets. |
| `public_subnet_cidrs` | `list(string)` | n/a | Public subnet CIDRs matching availability zone order. |
| `private_subnet_cidrs` | `list(string)` | n/a | Private subnet CIDRs matching availability zone order. |
| `enable_nat_gateway` | `bool` | `true` | Whether to create NAT gateway resources and private default routes. |
| `single_nat_gateway` | `bool` | `true` | Whether to create one shared NAT gateway. |
| `enable_dns_hostnames` | `bool` | `true` | Whether DNS hostnames are enabled in the VPC. |
| `enable_dns_support` | `bool` | `true` | Whether DNS resolution is enabled in the VPC. |
| `tags` | `map(string)` | `{}` | Tags applied to all supported AWS resources. |
| `public_subnet_tags` | `map(string)` | `{}` | Additional tags applied to public subnets. |
| `private_subnet_tags` | `map(string)` | `{}` | Additional tags applied to private subnets. |

## Outputs

| Name | Description |
| --- | --- |
| `vpc_id` | ID of the VPC. |
| `vpc_cidr_block` | CIDR block of the VPC. |
| `public_subnet_ids` | IDs of the public subnets, ordered by availability zones. |
| `private_subnet_ids` | IDs of the private subnets, ordered by availability zones. |
| `public_route_table_ids` | IDs of public route tables. |
| `private_route_table_ids` | IDs of private route tables, ordered by availability zones. |
| `nat_gateway_ids` | IDs of NAT gateways created by this module. |
| `internet_gateway_id` | ID of the internet gateway. |

## EKS Usage

```hcl
provider "aws" {
  region = "us-east-1"
}

module "network" {
  source = "../../../modules/cloud/aws/network"

  name               = "clusterforge-dev"
  environment        = "dev"
  cidr_block         = "10.40.0.0/16"
  availability_zones = ["us-east-1a", "us-east-1b"]

  public_subnet_cidrs  = ["10.40.0.0/20", "10.40.16.0/20"]
  private_subnet_cidrs = ["10.40.128.0/20", "10.40.144.0/20"]

  private_subnet_tags = {
    "kubernetes.io/cluster/clusterforge-dev" = "shared"
  }
}
```

## ECS Usage

```hcl
provider "aws" {
  region = "us-east-1"
}

module "network" {
  source = "../../../modules/cloud/aws/network"

  name               = "clusterforge-ecs-dev"
  environment        = "dev"
  cidr_block         = "10.50.0.0/16"
  availability_zones = ["us-east-1a", "us-east-1b"]

  public_subnet_cidrs  = ["10.50.0.0/20", "10.50.16.0/20"]
  private_subnet_cidrs = ["10.50.128.0/20", "10.50.144.0/20"]

  tags = {
    Project = "clusterforge"
  }
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
