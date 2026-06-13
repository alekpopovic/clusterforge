---
title: ClusterForge Documentation
permalink: /
---

<section class="hero">
  <div class="eyebrow">Terraform/OpenTofu framework</div>
  <h1>ClusterForge</h1>
  <p class="lead">
    Opinionated, readable Infrastructure-as-Code for Kubernetes, ECS, Nomad,
    and Docker-based container platforms.
  </p>
</section>

## Start Here

ClusterForge standardizes container-platform infrastructure without hiding the
Terraform. The CLI can generate files and guide safe workflows, but the
infrastructure logic remains visible in root environments and reusable modules.

<div class="grid">
  <div class="card">
    <h3>Architecture</h3>
    <p>Understand the foundation, orchestrator, platform, and workload layers.</p>
    <p><a href="{{ '/architecture/' | relative_url }}">Read architecture</a></p>
  </div>
  <div class="card">
    <h3>CLI</h3>
    <p>Generate projects, environments, app manifests, plans, and policy checks.</p>
    <p><a href="{{ '/cli/' | relative_url }}">Use the CLI</a></p>
  </div>
  <div class="card">
    <h3>Environments</h3>
    <p>Choose simple or stacked Terraform roots for each environment.</p>
    <p><a href="{{ '/environments/' | relative_url }}">Plan environments</a></p>
  </div>
  <div class="card">
    <h3>Backends</h3>
    <p>Generate local or remote state backend configuration safely.</p>
    <p><a href="{{ '/backends/' | relative_url }}">Configure backends</a></p>
  </div>
  <div class="card">
    <h3>Modules</h3>
    <p>Follow module conventions for provider placement, inputs, outputs, and docs.</p>
    <p><a href="{{ '/module-conventions/' | relative_url }}">Build modules</a></p>
  </div>
  <div class="card">
    <h3>Autoscaling</h3>
    <p>Install Karpenter for EKS node autoscaling while keeping bootstrap capacity.</p>
    <p><a href="{{ '/autoscaling/' | relative_url }}">Review autoscaling</a></p>
  </div>
  <div class="card">
    <h3>Security</h3>
    <p>Review production safety rules, secret handling, and static scanning.</p>
    <p><a href="{{ '/security/' | relative_url }}">Review security</a></p>
  </div>
</div>

## Quickstart

```bash
cd cli
go build -o cf .
cd ..

./cli/cf project init demo
./cli/cf env create dev --cloud aws --orchestrator eks --region eu-central-1
./cli/cf generate dev
./cli/cf init dev
./cli/cf plan dev
```

## Local Development

```bash
make help
make fmt-check
make test
make security
```

## Documentation Map

- [Architecture]({{ '/architecture/' | relative_url }})
- [CLI]({{ '/cli/' | relative_url }})
- [Environments]({{ '/environments/' | relative_url }})
- [Backends]({{ '/backends/' | relative_url }})
- [App manifest]({{ '/app-manifest/' | relative_url }})
- [Module conventions]({{ '/module-conventions/' | relative_url }})
- [General conventions]({{ '/conventions/' | relative_url }})
- [Validation]({{ '/validation/' | relative_url }})
- [GitOps]({{ '/gitops/' | relative_url }})
- [Node autoscaling]({{ '/autoscaling/' | relative_url }})
- [Security]({{ '/security/' | relative_url }})
- [Security scanning]({{ '/security-scanning/' | relative_url }})
- [Roadmap]({{ '/roadmap/' | relative_url }})
