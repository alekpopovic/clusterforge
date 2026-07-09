## Prompt 91 — AWS VPC endpoints module

```text
Create AWS VPC endpoints module.

Path:
- modules/cloud/aws/vpc-endpoints

Purpose:
Support private EKS/ECS environments by creating gateway and interface VPC endpoints.

Inputs:
- name
- environment
- vpc_id
- subnet_ids
- route_table_ids
- security_group_ids optional
- create_security_group default true
- allowed_security_group_ids list(string)
- gateway_endpoints list(string), default ["s3"]
- interface_endpoints list(string), default []
- private_dns_enabled default true
- tags map(string)

Resources:
- aws_vpc_endpoint for gateway endpoints
- aws_vpc_endpoint for interface endpoints
- aws_security_group optional for interface endpoints

Common interface endpoint examples:
- ecr.api
- ecr.dkr
- logs
- sts
- ec2
- eks
- secretsmanager
- ssm

Outputs:
- endpoint_ids
- endpoint_dns_entries
- security_group_id

Docs:
- Explain private cluster use case.
- Explain NAT Gateway cost reduction.
- Explain required endpoints for EKS nodes pulling images from ECR.
- Explain gateway vs interface endpoints.

Example:
- examples/aws-vpc-endpoints-private-eks

Rules:
- Do not enable every endpoint by default.
- Keep default minimal.
- Document regional service name construction.

Run:
- terraform fmt -recursive
```

---
