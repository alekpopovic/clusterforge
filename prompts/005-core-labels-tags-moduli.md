## Prompt 5 — Core labels/tags moduli

```text
Implement two core modules:

1. modules/core/tags
2. modules/core/labels

Goal:
Provide consistent metadata across cloud resources and Kubernetes resources.

modules/core/tags:
Inputs:
- project: string
- environment: string
- owner: string, default ""
- cost_center: string, default ""
- managed_by: string, default "terraform"
- component: string, default ""
- extra_tags: map(string), default {}

Behavior:
- Produce a map(string) of cloud tags.
- Include:
  Project
  Environment
  ManagedBy
  Component when non-empty
  Owner when non-empty
  CostCenter when non-empty
- Merge extra_tags last or first? Use extra_tags last only if we want override ability. Document the decision.
- Avoid null values.

Output:
- tags

modules/core/labels:
Inputs:
- project: string
- environment: string
- app: string, default ""
- component: string, default ""
- part_of: string, default ""
- managed_by: string, default "terraform"
- extra_labels: map(string), default {}

Behavior:
- Produce Kubernetes-compatible labels.
- Include recommended app.kubernetes.io labels where applicable:
  app.kubernetes.io/name
  app.kubernetes.io/part-of
  app.kubernetes.io/component
  app.kubernetes.io/managed-by
- Include clusterforge.io/project
- Include clusterforge.io/environment
- Merge extra_labels carefully.
- Ensure values are string-compatible and sane for Kubernetes labels.

Outputs:
- tags for tags module
- labels for labels module

README:
- Add examples for AWS tags and Kubernetes labels.
- Explain difference between cloud tags and Kubernetes labels.

Create examples/core-metadata showing both modules.

Run terraform fmt -recursive.
```

---
