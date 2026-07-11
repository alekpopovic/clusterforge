# Offline manifest bundles

`cf bundle create` produces a local directory containing generated/environment
Terraform files, the reusable module snapshot, sanitized app manifests, policy and
template packs, dependency lists, a runbook index, metadata and `SHA256SUMS`.
App literal environment variables are omitted; external secret references may be
retained because they contain identifiers, not values.

The lists are acquisition inputs:

- `images.txt` lists app container images.
- `helm-charts.txt` lists statically detectable chart/repository declarations.
- `providers.txt` lists providers found in copied lock/configuration files.
- `runbooks.txt` lists relevant repository runbooks without copying their content.

Static discovery can miss dynamically composed sources and transitive artifacts.
Provider lock files guide acquisition and carry provider checksums, but the MVP
does not populate a provider mirror. Helm and image lists likewise do not download
content. Resolve every transitive dependency and license before declaring an
offline deployment ready.

`cf bundle inspect <directory>` shows metadata and the file inventory.
`cf bundle verify <directory>` recalculates every checksum, rejects modified files,
and rejects unlisted additions. Creation refuses to overwrite an existing output
directory. Treat operational topology and source snapshots as sensitive even
after secret exclusion, and transfer/retain them under least privilege.
