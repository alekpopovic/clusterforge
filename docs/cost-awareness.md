# Cost Awareness

ClusterForge includes heuristic cost warnings for saved plan files:

```bash
cf cost scan dev --plan-file .cf/plans/dev.tfplan
cf cost estimate dev --plan-file .cf/plans/dev.tfplan --json
```

Built-in checks look for expensive categories such as NAT Gateway, EKS control
planes, managed node groups, load balancers, persistent volumes, and CloudWatch
logs. They do not use pricing data and do not claim exact cost.

Install Infracost for pricing-backed estimates when needed.
