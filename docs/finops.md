# FinOps and cost review

Built-in `cf cost scan|estimate|report <env> --plan-file <plan>` checks resource
categories heuristically and requires no cloud credentials. It flags NAT
gateways, load balancers, EKS/node groups, RDS, ElastiCache, CloudWatch logs,
persistent volumes, multi-region resources and related ongoing-cost risks. It
does not calculate prices and never blocks apply by itself.

Use `--infracost` for an installed/configured Infracost breakdown or
`cf cost diff <env>`. Keep `INFRACOST_API_KEY` outside Git and review pricing
assumptions. Add cost diffs to PRs, budgets/alerts per account, consistent
project/environment/owner/cost-center allocation tags and test-environment TTL
cleanup.

Review NAT gateway alternatives such as endpoints or architecture changes,
right-size node pools/databases/caches, make Multi-AZ/replication decisions from
reliability requirements, and set intentional log retention. Optimization must
not silently weaken availability, security or recovery objectives.
