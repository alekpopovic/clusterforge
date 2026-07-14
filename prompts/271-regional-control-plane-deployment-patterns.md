# Prompt 271 — Regional Control Plane deployment patterns

```text
Create regional Control Plane deployment patterns.

Goal:
Document deployment topologies for single-region, multi-region, and disaster recovery.

Create:
- docs/control-plane/regional-deployment.md

Topologies:
1. Single-region self-hosted
   - simplest
   - one database
   - one artifact backend
   - one runner pool

2. Single-region HA
   - multiple API replicas
   - HA database
   - S3-compatible artifact storage
   - multiple runners

3. Multi-region active/passive
   - primary Control Plane
   - standby region
   - database backup/restore
   - artifact replication
   - runner pools per region

4. Multi-region active/active
   - not recommended for MVP
   - requires strong consistency decisions
   - tenant routing
   - data residency concerns

5. Per-region runner pools with central API
   - common enterprise pattern
   - low complexity
   - good separation

Include:
- diagrams
- pros/cons
- failure modes
- backup implications
- data residency implications
- recommended v0.7 pattern

No code changes required.
```
