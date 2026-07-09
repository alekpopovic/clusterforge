## Prompt 60 — Public/private module registry strategy

```text
Create a module registry strategy for ClusterForge.

Goal:
Decide how modules will be consumed by users and teams.

Create:
- docs/registry-strategy.md

Cover:
1. Local development source:
   source = "../../../modules/..."

2. Git source:
   source = "git::https://github.com/<org>/clusterforge.git//modules/cloud/aws/network?ref=v0.1.0"

3. Terraform public registry option:
   - requirements
   - naming conventions
   - module repository layout considerations
   - pros and cons of monorepo vs split repos

4. Private registry option:
   - Terraform Cloud/Enterprise private registry
   - internal Git source usage
   - module governance

5. Versioning policy:
   - tag-based usage
   - do not use main branch in production
   - pin module versions

6. Migration plan:
   - local path during development
   - git source for first users
   - registry later

Update README:
- Add "Using modules from Git" section.

Rules:
- Do not restructure repository now.
- Do not publish anything.
- Make recommendation clear.

Final response:
- State recommended distribution path for next release.
```

---
