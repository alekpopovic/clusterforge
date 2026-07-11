# Kubernetes Security Policy Pack

Advisory intent:

- namespace pod security labels recommended
- network policy recommended
- privileged containers flagged
- LoadBalancer services flagged unless explicitly allowed

Use with Kubernetes provider plans and static manifests where available.

The extended production profile is opt-in and starts in audit/warn mode. It
covers image tags/registries/digests, pod privileges and host access, non-root
execution and capabilities, resource requests/limits, namespace NetworkPolicy,
secret-like environment literals, and approval annotations for public ingress.
See `policies.yaml` for stable rule identifiers and rollout settings.
