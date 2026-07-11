# AWS EKS incident

## Severity classification
SEV-1 for cluster/API compromise or regional production outage; SEV-2 for unavailable control plane/node fleet; SEV-3 for partial degradation.

## Symptoms
API timeouts, nodes NotReady, add-on/CNI failure, failed scaling, load balancer or identity errors.

## Initial checks
Confirm AWS identity/account/region; read EKS status, node groups, CloudWatch control-plane logs, EC2/ASG health and Kubernetes events.

## Containment
Freeze applies and autoscaler changes; revoke compromised IAM sessions/roles; restrict public API access only through an approved network change.

## Diagnosis
Correlate EKS events, VPC/subnet IP capacity, security groups, NAT/endpoints, IAM/OIDC, CNI/CoreDNS/kube-proxy and Karpenter/ASG events.

## Remediation
Use a reviewed plan or provider-supported repair; restore add-on/node configuration and validate workloads.

## Rollback
Roll workloads/add-ons forward or revert reviewed configuration. EKS control-plane downgrade is not a rollback. **Destructive:** node group replacement requires drain, capacity and approval.

## Communication notes
Report account/region/cluster (redacted as needed), affected AZs, API/data-plane status and customer impact.

## Evidence collection
Preserve CloudTrail, EKS/CloudWatch logs, events, plan summaries, IAM changes, node diagnostics and exact versions.

## Postmortem checklist
Timeline; AWS/Kubernetes cause; IAM/network/capacity gaps; restore/cleanup proof; actions and owners.
