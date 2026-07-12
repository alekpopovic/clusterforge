---
title: Security scanner exceptions
description: Rules for narrow Checkov and Trivy exceptions in reusable ClusterForge modules.
---

# Security scanner exceptions

ClusterForge keeps Checkov and Trivy blocking for source infrastructure. The
repository does not use a global `skip-check`, severity reduction, or soft-fail
to obtain a green security scan.

Some static checks cannot determine the final value of a typed Terraform module
input or follow a security-group attachment through a caller. Other controls,
such as WAF, cross-region replication, DNS query logging, customer-managed keys,
and public/private load-balancer placement, require resources or decisions owned
by the consuming environment. Those cases may use an inline exception only when:

- it is attached to the exact Terraform or Docker resource reported;
- it includes the scanner rule ID;
- the reason states the operational boundary or static-analysis limitation;
- the module keeps the relevant choice visible to callers;
- production documentation or policy does not claim the exception is secure for
  every environment.

Inline exceptions must not be copied to unrelated resources. Remove an exception
when the module can satisfy the control without breaking compatibility or
silently creating operator-owned infrastructure. Every change must run:

```bash
make security
make test
```

Passing scanners means the configured checks have no unresolved findings. It is
not a certification and does not replace plan, IAM, network, data-protection,
cost, or production-readiness review.
