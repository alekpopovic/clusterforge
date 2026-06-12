# cloud/aws/dns

## Purpose

Manages Route53 hosted zone creation or lookup and DNS records for
ClusterForge environments.

Provider configuration belongs in the root module. This module declares the
AWS provider requirement but does not configure the provider.

## Status

Implemented.

## Hosted Zone Creation

```hcl
module "dns" {
  source = "../../../modules/cloud/aws/dns"

  create_zone = true
  zone_name   = "example.com"

  tags = module.tags.tags
}
```

Creating a public hosted zone can affect delegation and live DNS. Review name
servers and registrar delegation before using this in production.

## Existing Zone Lookup

```hcl
module "dns" {
  source = "../../../modules/cloud/aws/dns"

  create_zone = false
  zone_name   = "example.com"

  records = {
    txt = {
      name    = "clusterforge.example.com"
      type    = "TXT"
      ttl     = 300
      records = ["\"managed-by=clusterforge\""]
    }
  }
}
```

When `create_zone = false` and `zone_id` is empty, the module looks up a public
hosted zone by `zone_name`.

## ALB Alias Record

```hcl
module "dns" {
  source = "../../../modules/cloud/aws/dns"

  zone_name = "example.com"

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
}
```

Alias records must not set `ttl`. Route53 uses the target service behavior.

## Kubernetes Ingress And External DNS

For Kubernetes ingress records, prefer ExternalDNS when DNS should follow
Ingress or Service annotations. Use this module for explicit environment DNS
records such as ALB aliases, validation records, and stable platform records.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
