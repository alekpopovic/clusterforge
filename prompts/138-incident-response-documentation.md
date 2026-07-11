## Prompt 138 — Incident response documentation

```text
Create incident response documentation for ClusterForge-managed platforms.

Create:
- docs/incident-response/
  index.md
  kubernetes-incident.md
  aws-eks-incident.md
  aws-ecs-incident.md
  terraform-state-incident.md
  secret-leak-incident.md
  dns-incident.md
  failed-deployment.md
  cluster-outage.md

Each runbook must include:
- Severity classification
- Symptoms
- Initial checks
- Containment
- Diagnosis
- Remediation
- Rollback
- Communication notes
- Evidence collection
- Postmortem checklist

Specific incidents:
1. Terraform apply failed midway
2. Accidental resource destroy
3. Kubernetes nodes NotReady
4. Ingress unreachable
5. Certificate expired
6. Secret leaked
7. DNS changed incorrectly
8. Image pull failures
9. ECS service unhealthy
10. Argo CD sync failure

Rules:
- No fake guarantees.
- No real credentials.
- Keep runbooks actionable.
- Include commands where useful, but mark destructive commands clearly.
```


---
