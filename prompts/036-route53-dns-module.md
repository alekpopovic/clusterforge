## Prompt 36 — Route53 DNS module

```text
Implement AWS Route53 DNS module.

Create module:
- modules/cloud/aws/dns

Purpose:
Manage Route53 hosted zone lookup/creation and DNS records for ClusterForge environments.

Inputs:
- create_zone: bool, default false
- zone_name: string
- zone_id: string, default ""
- records: map(object({
    name = string
    type = string
    ttl = optional(number)
    records = optional(list(string))
    alias = optional(object({
      name = string
      zone_id = string
      evaluate_target_health = optional(bool, true)
    }))
  })), default {}
- tags: map(string), default {}

Behavior:
- If create_zone=true, create aws_route53_zone.
- If create_zone=false and zone_id is empty, lookup zone by zone_name.
- Create aws_route53_record for records.
- Support normal records and alias records.
- Validate that either records or alias is provided per record.
- For alias records, ttl should not be used.

Outputs:
- zone_id
- zone_name
- name_servers
- record_fqdns

README:
- Hosted zone creation example.
- Existing zone lookup example.
- ALB alias record example.
- Ingress/external-dns note.

Rules:
- Do not create real domain values in examples.
- Use example.com for examples.
- Do not hardcode AWS account data.
- Be careful with destructive DNS changes in docs.

Create example:
- examples/aws-route53-dns

Run:
- terraform fmt -recursive
- validation where possible
```

---
