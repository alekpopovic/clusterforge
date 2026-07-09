## Prompt 35 — ECS ALB module i povezivanje sa ECS servisom

```text
Implement ECS Application Load Balancer support.

Create module:
- modules/platform/ecs/alb

Purpose:
Create an AWS Application Load Balancer, listeners, target groups, and security groups for ECS/Fargate services.

Inputs:
- name: string
- environment: string
- vpc_id: string
- public_subnet_ids: list(string)
- allowed_cidr_blocks: list(string), default ["0.0.0.0/0"]
- certificate_arn: string, default ""
- enable_http: bool, default true
- enable_https: bool, default false
- target_groups: map(object({
    port = number
    protocol = optional(string, "HTTP")
    health_check_path = optional(string, "/")
    health_check_matcher = optional(string, "200-399")
  }))
- tags: map(string), default {}

Resources:
- aws_security_group for ALB
- aws_lb
- aws_lb_target_group for each target group
- aws_lb_listener for HTTP if enabled
- aws_lb_listener for HTTPS if enabled and certificate_arn provided
- Optional HTTP to HTTPS redirect if both enabled

Outputs:
- alb_arn
- alb_dns_name
- alb_zone_id
- security_group_id
- target_group_arns
- listener_arns

Update:
- modules/workloads/ecs/service README
  Show how to pass target_group_arn.

Create example:
- examples/ecs-fargate-app-with-alb

Rules:
- Do not create Route53 DNS records yet unless already existing DNS module supports it.
- Do not hardcode cert ARNs.
- HTTPS must require certificate_arn.
- Keep target groups generic.

Run:
- terraform fmt -recursive
- validation where possible

Final response:
- Summarize ALB support.
- Show short usage snippet.
- Mention DNS/certificate assumptions.
```

---
