# Kubernetes incident

## Severity classification
SEV-1 for broad production outage/privilege compromise; SEV-2 for major node, ingress or workload degradation; SEV-3 otherwise.

## Symptoms
Nodes NotReady, ingress unreachable, expired certificates, image pull failures, crash loops, or Argo CD sync failure.

## Initial checks
Run read-only `kubectl get nodes,pods -A`, `kubectl get events -A --sort-by=.lastTimestamp`, inspect ingress/certificates and `argocd app get <app>` when authorized.

## Containment
Pause GitOps auto-sync or rollout only with incident approval; isolate a compromised namespace/workload using reviewed policy. Do not delete nodes/pods as first response.

## Diagnosis
Check node conditions, CNI/DNS, quotas, image registry auth/digest, ingress endpoints, CertificateRequest/Challenge, controller logs and Argo desired/live diff.

## Remediation
Repair the failed dependency or apply a reviewed manifest. Rotate registry/secret credentials through the source secret store.

## Rollback
Revert to a known image/config commit through GitOps. **Destructive/high-risk:** node replacement, namespace deletion and forced finalizer removal require explicit approval.

## Communication notes
State affected clusters/namespaces/workloads and whether ingress, certificates or deployment reconciliation is impaired.

## Evidence collection
Capture redacted events, object YAML, controller logs, image digest, Git commit, chart versions and audit references.

## Postmortem checklist
Timeline; node/ingress/certificate/image/Argo cause; detection and capacity gaps; tested corrective actions; owners and dates.
