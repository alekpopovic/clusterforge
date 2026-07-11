## Prompt 150 — Compliance mapping packs

```text
Create compliance mapping documentation and policy packs.

Goal:
Map ClusterForge controls to common compliance frameworks without claiming certification.

Create:
- docs/compliance/
  index.md
  soc2-mapping.md
  iso27001-mapping.md
  cis-kubernetes-mapping.md
  cis-aws-foundations-mapping.md
  nsa-cisa-kubernetes-hardening.md

Create policy pack metadata:
- policies/packs/compliance-soc2/
- policies/packs/compliance-cis-kubernetes/
- policies/packs/compliance-aws-foundations/

Each mapping must include:
- Control area
- ClusterForge feature or policy
- Status:
  - implemented
  - partial
  - documented
  - planned
  - not covered
- Evidence source:
  - policy result
  - audit log
  - Terraform config
  - CI result
  - manual procedure
- Limitations

Important:
- Do not claim SOC 2, ISO 27001, CIS, or NSA/CISA compliance.
- Say these are control mappings and implementation aids.
- Add legal/compliance disclaimer.

CLI:
- cf compliance list
- cf compliance report --pack cis-kubernetes
- cf compliance report --format markdown|json

Tests:
- Compliance report renders.
- Unknown pack fails.
- Markdown output.

Run:
- gofmt
- go test ./...
```


---
