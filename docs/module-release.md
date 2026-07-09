# Module Release

ClusterForge modules are distributed from this monorepo for the next release.
Use local paths during development and Git sources pinned to tags for early
consumers.

```hcl
source = "../../../modules/cloud/aws/network"
```

```hcl
source = "git::https://github.com/<org>/clusterforge.git//modules/cloud/aws/network?ref=v0.1.0"
```

Future registry-style examples will use split or generated registry metadata:

```hcl
source  = "<org>/network/aws"
version = "~> 0.1"
```

Release rules:

- tag the whole repository, such as `v0.1.0`
- do not introduce breaking module input/output changes without a pre-1.0
  minor bump
- document module-specific compatibility notes in release notes
- stable modules need README, versions, examples, and tests
- placeholder modules must be clearly marked
