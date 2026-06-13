# platform/kubernetes/network-policy-baseline

## Purpose

Creates baseline Kubernetes NetworkPolicy resources for a namespace.

## Status

Implemented.

## Usage

```hcl
module "network_policy_baseline" {
  source = "../../../modules/platform/kubernetes/network-policy-baseline"

  namespace = "apps"

  default_deny_ingress = true
  default_deny_egress  = false
}
```

## Default Deny Tradeoffs

Default deny ingress is a strong baseline, but it can break traffic until
workloads define explicit allow policies. Default deny egress is even more
disruptive because applications often need DNS, cloud APIs, databases, queues,
and observability endpoints.

This module does not enable itself globally through bootstrap defaults. Use it
intentionally namespace by namespace.

## DNS Egress

When `default_deny_egress = true` and `allow_dns_egress = true`, this module
adds an egress allow policy for DNS traffic to pods labeled `k8s-app=kube-dns`
in the `kube-system` namespace. Some clusters use different DNS labels or
NodeLocal DNSCache, so review and override with explicit policies where needed.

## CNI Requirement

NetworkPolicy enforcement requires a CNI plugin that supports NetworkPolicy.
Some managed Kubernetes defaults do not enforce NetworkPolicy unless an
appropriate CNI or policy engine is installed.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
