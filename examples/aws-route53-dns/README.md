# aws-route53-dns

## Purpose

Example Route53 DNS composition using `modules/cloud/aws/dns`.

This example uses `example.com` and placeholder Route53 identifiers. Replace
them with real values before applying.

## Usage

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

The example includes fake AWS credentials for local syntax planning with
`-refresh=false`. Use real AWS credentials before applying.

## ALB Alias

In a real ECS composition, pass ALB outputs directly:

```hcl
records = {
  app = {
    name = "app.example.com"
    type = "A"

    alias = {
      name    = module.alb.alb_dns_name
      zone_id = module.alb.alb_zone_id
    }
  }
}
```

DNS changes can be disruptive. Review hosted zone IDs, record names, and
planned deletions carefully before applying.
