---
title: ClusterForge Documentation
description: Readable Terraform and OpenTofu platform engineering for Kubernetes, ECS, Nomad, and Docker.
permalink: /
---

<section class="hero-shell">
  <div class="relative z-10">
    <span class="status-pill">● v0.4.2 available</span>
    <p class="hero-kicker mt-6">Visible infrastructure. Safer operations.</p>
    <h1>Forge consistent container platforms without hiding Terraform.</h1>
    <p class="hero-lead">ClusterForge combines readable Terraform/OpenTofu modules with a Go CLI for generation, policy, fleet visibility, upgrades, compliance mappings, GitOps, and operational evidence.</p>
    <div class="mt-7 flex flex-wrap gap-3">
      <a href="{{ '/installation/' | relative_url }}" class="rounded-xl bg-forge-600 px-5 py-3 text-sm font-bold text-white no-underline shadow-forge hover:bg-forge-700">Install the CLI</a>
      <a href="{{ '/getting-started.html' | relative_url }}" class="rounded-xl border border-slate-300 bg-white/70 px-5 py-3 text-sm font-bold text-slate-800 no-underline hover:border-forge-400 dark:border-slate-600 dark:bg-slate-900/60 dark:text-white">Start building</a>
      <a href="https://github.com/alekpopovic/clusterforge" class="rounded-xl px-5 py-3 text-sm font-bold no-underline">View on GitHub →</a>
    </div>
  </div>
</section>

## What ClusterForge is today

ClusterForge is a multi-target platform engineering framework. Provider
configuration stays in environment roots, reusable modules stay focused, and the
CLI generates files you can inspect, edit, review, and run with Terraform or
OpenTofu directly.

For an evidence-based capability and maturity matrix, see
[Project status]({{ '/project-status/' | relative_url }}).

<div class="docs-grid">
  <a class="docs-card no-underline" href="{{ '/architecture/' | relative_url }}"><h3>Four-layer architecture</h3><p>Foundation, orchestrator, platform, and workload modules with visible provider boundaries.</p></a>
  <a class="docs-card no-underline" href="{{ '/cli/' | relative_url }}"><h3>Operations-aware CLI</h3><p>Scaffold, generate, plan, inspect drift, export inventory, evaluate policy, and prepare upgrades.</p></a>
  <a class="docs-card no-underline" href="{{ '/fleet-operations.html' | relative_url }}"><h3>Fleet visibility</h3><p>Read-only health, inventory, graph, dashboard-data, cost, and multi-cluster metadata workflows.</p></a>
  <a class="docs-card no-underline" href="{{ '/policy-engine.html' | relative_url }}"><h3>Policy &amp; compliance</h3><p>Versioned policy packs, admission guidance, security checks, and non-certifying control mappings.</p></a>
  <a class="docs-card no-underline" href="{{ '/gitops-multi-cluster.html' | relative_url }}"><h3>GitOps handoff</h3><p>Render Argo CD app-of-apps structures while cluster registration and credentials stay external.</p></a>
  <a class="docs-card no-underline" href="{{ '/air-gapped.html' | relative_url }}"><h3>Restricted environments</h3><p>Checksummed offline manifest bundles, dependency acquisition lists, and explicit secret exclusions.</p></a>
</div>

## Supported targets and maturity

| Area | Current state |
|---|---|
| AWS EKS and ECS | Implemented modules, examples, generation, production safety guidance, and credential-free validation. Account-specific production review and real-cloud evidence are still required. |
| Existing Kubernetes | Implemented attachment, platform bootstrap, workload rendering, policy, backup, and operational documentation without owning cluster lifecycle. |
| Local Kubernetes | K3s/RKE2 and local example paths are available for development; lightweight/edge production lifecycle remains experimental. |
| Azure AKS and Google GKE | Modules and examples validate, but production maturity remains experimental until current real-cloud hardening evidence exists. |
| Nomad and Docker | Focused cluster/workload patterns exist; host lifecycle, networking, secrets, and production hardening remain operator-owned. |
| Windows containers | ECS runtime metadata and app platform intent are experimental. Production Windows and Kubernetes Windows lifecycle are not claimed. |

## Install and create a project

Linux or macOS:

```bash
curl -fsSL \
  https://github.com/alekpopovic/clusterforge/releases/download/v0.4.2/install.sh \
  | VERSION=v0.4.2 bash

cf wizard --defaults --non-interactive
cf generate dev
cf plan dev
```

The installer downloads the matching release binary and verifies its SHA256
checksum before installation. See [Install the CLI]({{ '/installation/' | relative_url }})
for Linux, macOS, Windows PowerShell, manual verification, upgrades, and uninstall.

## Enterprise and operations workflows

<div class="docs-grid">
  <a class="docs-card no-underline" href="{{ '/plugins.html' | relative_url }}"><h3>Plugins &amp; templates</h3><p>Local plugins are disabled by default. Versioned template packs keep generation extensible and reviewable.</p></a>
  <a class="docs-card no-underline" href="{{ '/aws-multi-account.html' | relative_url }}"><h3>Organization model</h3><p>Teams, workspaces, AWS accounts, regions, execution profiles, and Terraform Cloud metadata.</p></a>
  <a class="docs-card no-underline" href="{{ '/platform-upgrades.html' | relative_url }}"><h3>Upgrade planning</h3><p>Read-only Kubernetes and platform add-on readiness with explicit rollback limitations.</p></a>
  <a class="docs-card no-underline" href="{{ '/backup-validation.html' | relative_url }}"><h3>Operational evidence</h3><p>Audit export, backup/restore-test evidence, incident runbooks, SLOs, secret rotation, and drift checks.</p></a>
  <a class="docs-card no-underline" href="{{ '/service-catalog.html' | relative_url }}"><h3>Catalog integrations</h3><p>Service metadata, Backstage YAML, runbook indexes, and stable dashboard export schemas.</p></a>
  <a class="docs-card no-underline" href="{{ '/migration-analyzer.html' | relative_url }}"><h3>Adoption analysis</h3><p>Read-only Terraform repository analysis with provider/resource mapping and redacted risk findings.</p></a>
</div>

## Safety boundary

ClusterForge is a wrapper and generator—not a hidden control plane. Production
apply requires a saved plan, production destroy is blocked by default, secret
values do not belong in generated files, and read-only commands do not call cloud
APIs unless explicitly documented. Compliance documents are implementation aids,
not certifications.

The current stable release has passing Terraform validation and CLI tests. The
release assessment still tracks lint, Checkov, dependency, and production-cloud
evidence; stable versioning does not imply every target is production-proven. Read
[the roadmap]({{ '/roadmap/' | relative_url }}) and
[release process]({{ '/release-process.html' | relative_url }}) for the exact boundary.

## Documentation paths

- New users: [installation]({{ '/installation/' | relative_url }}), [getting started]({{ '/getting-started.html' | relative_url }}), [wizard]({{ '/wizard.html' | relative_url }}), [CLI]({{ '/cli/' | relative_url }})
- Platform teams: [architecture]({{ '/architecture/' | relative_url }}), [modules]({{ '/modules.html' | relative_url }}), [environments]({{ '/environments/' | relative_url }}), [backends]({{ '/backends/' | relative_url }})
- Operators: [operations]({{ '/operations.html' | relative_url }}), [fleet]({{ '/fleet-operations.html' | relative_url }}), [upgrades]({{ '/upgrades.html' | relative_url }}), [incident response]({{ '/incident-response/' | relative_url }})
- Security teams: [security]({{ '/security/' | relative_url }}), [admission policy]({{ '/kubernetes-admission-security.html' | relative_url }}), [compliance mappings]({{ '/compliance/' | relative_url }}), [supply chain]({{ '/supply-chain-security.html' | relative_url }})
