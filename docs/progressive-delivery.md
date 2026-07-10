# Progressive delivery with Argo Rollouts

Argo Rollouts is optional. The standard ClusterForge app module continues to use
Kubernetes Deployment, while `rollout-app` creates an Argo Rollout for workloads
that explicitly need progressive delivery.

## Deployment and Rollout

A Deployment provides rolling replacement with readiness-based availability.
A Rollout adds staged promotion, pauses, blue/green services, analysis runs, and
integrations for weighted traffic. This capability adds a controller, CRDs, more
Services, and additional operational decisions.

## Canary

Canary steps gradually increase the new ReplicaSet weight and may pause for
manual review. Without a traffic router, weight is approximated through replica
counts. Precise percentages require a supported ingress controller or service
mesh. Start with a few observable steps and an explicit manual pause.

## Blue/green

Blue/green keeps active and preview Services. The new version can be tested on
the preview Service before promotion switches the active Service. Automatic
promotion is disabled by default in the module so operators retain a review gate.
Capacity planning must include both versions during the rollout.

## Metrics and rollback

Automated analysis requires an AnalysisTemplate and a configured metrics source
such as Prometheus, Datadog, a web check, or a Kubernetes Job. The first workload
module intentionally does not invent success thresholds. Define service-level
signals, test failure behavior, and protect credentials used by metrics providers.

Use the kubectl Argo Rollouts plugin to inspect, promote, abort, retry, and undo a
Rollout. An abort stops progression; rollback still requires selecting or applying
the known-good image/configuration. Keep immutable images and normal Terraform
review so desired state remains clear after emergency actions.

## Rollout plan

1. Pin and install the controller; confirm CRDs and controller health.
2. Deploy one non-critical app with manual pauses and no automated analysis.
3. Validate stable/canary or active/preview Service routing.
4. Add metrics-based analysis only after dashboards and thresholds are trusted.
5. Exercise promotion, abort, rollback, controller outage, and cleanup.

The dashboard ingress is disabled by default. Prefer local access through the
kubectl plugin unless authentication and exposure are explicitly designed.

Official references:

- [Argo Rollouts overview](https://argo-rollouts.readthedocs.io/en/latest/)
- [Controller installation](https://argo-rollouts.readthedocs.io/en/release-1.8/installation/)
- [Dashboard](https://argo-rollouts.readthedocs.io/en/release-1.8/dashboard/)
