# Import Strategy

Future commands:

```bash
cf adopt aws-vpc
cf adopt eks
cf adopt ecs
cf adopt route53-zone
cf import plan
cf import generate
```

For now, use Terraform import blocks in a reviewed branch:

```hcl
import {
  to = module.network.aws_vpc.this
  id = "vpc-xxxxxxxx"
}
```

Start with read-only data sources when possible. Generate import plans, inspect
all replacements, and avoid mixing multiple ownership models for the same
resource.
