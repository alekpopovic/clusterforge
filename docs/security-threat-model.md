# ClusterForge security threat model

## Scope and trust boundaries

ClusterForge generates readable Terraform/OpenTofu, invokes infrastructure
tools, and packages reusable modules and policies. Important trust boundaries
exist between the operator workstation, source repository, CI runners, module
and provider registries, remote state, cloud APIs, Kubernetes APIs, and release
distribution. This model is guidance, not a compliance claim or proof that a
deployment is secure.

## Assets

| Asset | Why it matters | Typical boundary |
| --- | --- | --- |
| Terraform state | May contain infrastructure metadata and sensitive values | workstation/CI to remote backend |
| Cloud credentials | Can authorize infrastructure changes | identity provider to workstation/CI |
| Kubeconfig files | Can grant cluster API access | workstation/CI to Kubernetes API |
| Generated Terraform | Defines privileged and networked resources | source/generator to plan and apply |
| App manifests | Control workload images, permissions, and exposure | application source to generators |
| CLI audit logs | Record operations and may expose paths or metadata | CLI to local log storage |
| Release artifacts | Executed with operator privileges | release workflow to operator |
| CI credentials | Can publish releases or change infrastructure | GitHub runner to cloud/registry |
| Module sources | Execute declarative provider operations | registry/source control to Terraform |
| Template packs | Influence generated roots and workloads | external source to generator |
| Policy packs | Decide whether risky plans are accepted | policy source to CI/operator |

## Threats and mitigations

| Threat | Impact | Existing controls | Residual risk / planned work |
| --- | --- | --- | --- |
| Leaked Terraform state | Credential or topology disclosure | `.gitignore`; remote backend guidance for encryption, locking, and access control | Backend security is operator-configured; automated backend posture checks are planned |
| Committed secrets | Account or cluster compromise | ignore rules, pre-commit hooks, Gitleaks workflow, contributor redaction rules | Scanners can miss encoded or novel secrets; history cleanup remains manual |
| Malicious template pack | Generated backdoor, unsafe provider/module source | Template packs are opt-in and output stays reviewable | Signing, provenance, sandboxing, and an allowlist are planned; treat packs as code today |
| Compromised CLI binary | Arbitrary execution with operator credentials | Release checksums and release workflow security scans | Artifact signing and complete provenance verification are planned; build from reviewed source for high assurance |
| Provider/module supply chain compromise | Arbitrary provider code execution or hostile infrastructure changes | Version constraints, dependency review, lock-file review guidance, checksums/SBOM release work | Constraints are not immutability; operators must review lock changes and trusted sources |
| Over-permissive IAM | Cloud privilege escalation or broad blast radius | Focused IAM modules, least-privilege guidance, IRSA/workload identity support | Generated roles still require environment-specific review; continuous entitlement analysis is not implemented |
| Accidental production destroy | Availability or data loss | Existing plan file required for production apply, delete override, production destroy block and confirmation | Direct Terraform use bypasses CLI protections; backend protection and peer review remain necessary |
| Drift hides malicious changes | Unreviewed infrastructure persists | Read-only drift command, scheduled drift workflow templates, audit logging | No automatic remediation or guaranteed alert delivery; operators must schedule and triage checks |
| Kubernetes privilege escalation | Cluster takeover or tenant escape | Pod security, NetworkPolicy, RBAC/service-account, Kyverno and Gatekeeper modules | Policy enforcement is optional; admission coverage and chart permissions need cluster-specific testing |
| Public ingress exposure | Data or service exposure | Explicit ingress inputs, NetworkPolicy modules, policy checks | Cloud load balancer and firewall composition can still expose services; default-deny is not universal |
| DNS takeover or misconfiguration | Traffic interception or outage | Route53 modules and scoped external-dns IAM patterns | Ownership records, delegation, deletion protection, and stale records require operational review |

## Control status

### Implemented in this repository

- Ignore rules for common state, plan, credential, key, kubeconfig, and local
  artifact files.
- Secret scanning workflows and optional pre-commit checks.
- Policy-pack structure and advisory plan-risk checks.
- Production CLI requirements for reviewed plan files and explicit destructive
  overrides.
- Remote-state modules and encryption guidance.
- Reusable least-privilege and workload-identity building blocks, including
  EKS IRSA.
- Release checksum generation and SBOM workflow support.
- Credential-free CI jobs where practical and GitHub workflow permission
  scoping.
- Local CLI audit logging and read-only drift/fleet inspection commands.

These controls reduce risk only when enabled and correctly configured. Users
can invoke Terraform directly, select unsafe module inputs, install untrusted
providers, or weaken cloud and Kubernetes controls outside ClusterForge.

### Planned or incomplete

- Signed releases and verifiable end-to-end build provenance.
- Signed, allowlisted, or sandboxed template and policy packs.
- Automated remote-state, IAM entitlement, DNS ownership, and public-exposure
  posture validation.
- Broad real-cloud and Kubernetes-version security regression coverage.
- Formal vulnerability-response SLA, external audit, or compliance mapping.
- Automatic remediation of drift or security findings (explicitly outside the
  v0.3 scope).

## Review triggers

Update this model when a new execution path, provider, cloud, template/plugin
mechanism, credential flow, release channel, or production mutation command is
introduced. Review it before each minor release and after a relevant incident.
