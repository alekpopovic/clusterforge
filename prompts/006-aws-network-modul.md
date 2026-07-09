## Prompt 6 — AWS network modul

```text
Implement modules/cloud/aws/network.

Purpose:
Create a production-friendly AWS VPC network suitable for EKS and ECS.

Provider:
- Use hashicorp/aws.
- Do not configure the provider inside the module.

Inputs:
- name: string
- environment: string
- cidr_block: string
- availability_zones: list(string)
- public_subnet_cidrs: list(string)
- private_subnet_cidrs: list(string)
- enable_nat_gateway: bool, default true
- single_nat_gateway: bool, default true
- enable_dns_hostnames: bool, default true
- enable_dns_support: bool, default true
- tags: map(string), default {}
- public_subnet_tags: map(string), default {}
- private_subnet_tags: map(string), default {}

Validation:
- availability_zones length must be > 0.
- public_subnet_cidrs length must equal availability_zones length.
- private_subnet_cidrs length must equal availability_zones length.
- cidr_block must not be empty.
- name and environment must not be empty.

Resources:
- aws_vpc
- aws_internet_gateway
- aws_subnet public
- aws_subnet private
- aws_route_table public
- aws_route_table private
- aws_route_table_association public/private
- aws_eip for NAT when enabled
- aws_nat_gateway when enabled
- routes:
  - public default route to IGW
  - private default route to NAT when NAT enabled

Kubernetes compatibility:
- Add optional/default tags suitable for Kubernetes load balancer discovery:
  public subnet:
    kubernetes.io/role/elb = "1"
  private subnet:
    kubernetes.io/role/internal-elb = "1"
- Allow user to pass cluster-specific subnet tags through inputs.

Outputs:
- vpc_id
- vpc_cidr_block
- public_subnet_ids
- private_subnet_ids
- public_route_table_ids
- private_route_table_ids
- nat_gateway_ids
- internet_gateway_id

README:
- Explain module purpose.
- Add EKS usage example.
- Add ECS usage example.
- Mention NAT cost implications.
- Mention that provider is configured in root module.

Also create examples/aws-network with:
- providers.tf
- main.tf
- variables.tf
- outputs.tf
- versions.tf
- README.md

Run terraform fmt -recursive.
```

---
