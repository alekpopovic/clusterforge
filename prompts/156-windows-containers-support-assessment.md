## Prompt 156 — Windows containers support assessment

```text
Create Windows containers support assessment.

Create:
- docs/rfcs/015-windows-containers.md

Goal:
Decide whether ClusterForge should support Windows container workloads.

Cover:
1. Kubernetes Windows nodes:
   - workload constraints
   - node selectors
   - taints/tolerations
   - networking caveats
   - image requirements

2. ECS Windows tasks:
   - launch type support
   - task definition differences
   - logging
   - IAM
   - cost implications

3. Module impact:
   - workloads/kubernetes/app
   - workloads/ecs/service
   - node group modules
   - app manifest schema

4. Proposed MVP:
   - documentation first
   - workload node_selector/tolerations support
   - ECS task runtime platform input if practical

5. Non-goals:
   - full Windows node lifecycle automation in first version

If implementation is straightforward:
- Add runtime_platform support to ECS service module.
- Add docs examples.
- Add app manifest field:
  platform:
    os: linux|windows
    architecture: amd64|arm64

Rules:
- Do not claim Windows support unless tested.
- Mark experimental.
```


---
