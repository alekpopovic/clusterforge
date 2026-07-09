# Kubernetes Platform Module Conventions

Helm-based platform modules under `modules/platform/kubernetes/` should expose
a consistent interface so the bootstrap module can compose them predictably.

Required module inputs:

- `namespace`
- `chart_version`
- `values`
- `create_namespace`

Required outputs:

- `namespace`
- `release_name`

Each module README needs a usage example. Chart versions are intentionally not
pinned globally; production roots should pin versions through environment
configuration after review.

Run the conformance check without a Kubernetes cluster:

```bash
scripts/check-platform-modules.sh
```

The script checks the target Helm chart names, repository configuration,
required inputs and outputs, README examples, and basic bootstrap wiring. It
does not install charts or contact a Kubernetes API.

Bootstrap should pass through add-on-specific `chart_version` and `values`
inputs when those controls are declared. Missing pass-through controls are
reported as warnings so maintainers can add them deliberately without blocking
unrelated module work.
