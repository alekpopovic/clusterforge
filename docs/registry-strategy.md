# Registry Strategy

Recommended path for the next release: local paths for contributors, Git
sources pinned to repository tags for early users, and Terraform registry
publishing later.

## Local Development

```hcl
source = "../../../modules/cloud/aws/network"
```

## Git Source

```hcl
source = "git::https://github.com/<org>/clusterforge.git//modules/cloud/aws/network?ref=v0.1.0"
```

Do not use `main` in production. Pin module versions with tags.

## Public Registry Option

Terraform Registry publishing needs stable module interfaces, examples, README
docs, version tags, and registry naming conventions. A monorepo is convenient
now, but split repos can be considered if registry workflows become painful.

## Private Registry Option

Teams can use Terraform Cloud/Enterprise private registry or internal Git
sources. Private registry governance should require owners, review rules,
version pinning, and deprecation notes.

## Migration Plan

1. Use local paths during active development.
2. Use Git tag sources for first adopters.
3. Publish registry modules only after stable examples and smoke evidence.
