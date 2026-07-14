# Prompt 257 — Runner autoscaling RFC

```text
Create runner autoscaling RFC.

Goal:
Design how ClusterForge could scale runners based on queued jobs.

Create:
- docs/rfcs/031-runner-autoscaling.md
- docs/control-plane/runner-autoscaling.md

Autoscaling signals:
- queued jobs
- queue wait time
- jobs by type
- environment priority
- runner pool labels
- max concurrent jobs
- failure rate
- cost constraints

Execution environments:
1. Kubernetes Deployment scale
2. Kubernetes Job-per-runner
3. VM auto scaling group
4. CI runner integration
5. manual scaling

Control Plane responsibilities:
- expose queue metrics
- define desired runner capacity
- enforce pool limits
- avoid dev runner processing prod jobs

Non-goals:
- building a full Kubernetes autoscaler
- cloud-specific autoscaling implementation in first phase
- automatic apply scaling without safety controls

Future commands:
- cf runner autoscaling status
- cf runner autoscaling configure
- cf runner autoscaling simulate

Do not implement code.
Update roadmap.
```
