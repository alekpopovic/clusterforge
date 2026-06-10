# cloud/aws/network

Creates the AWS networking foundation for a container platform:

- VPC
- Public and private subnets
- Internet gateway
- Optional single NAT gateway
- Public and private route tables

Provider configuration must be declared in the root module.

## Example

```hcl
module "network" {
  source = "../../../../modules/cloud/aws/network"

  name     = "clusterforge-dev"
  vpc_cidr = "10.40.0.0/16"
  tags     = module.tags.tags
}
```
