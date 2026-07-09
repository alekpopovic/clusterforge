# Version Support

ClusterForge supports currently maintained Kubernetes minor versions unless a
module explicitly documents otherwise. Tested versions are narrower than
supported versions and are tracked in `VERSION_MATRIX.md`.

Provider constraints must be pinned with compatible ranges. Examples should
state tested versions, and release notes must call out version support changes.

Status meanings:

- supported: intended to work within documented constraints
- tested: verified by CI, local validation, or a recorded smoke test
- experimental: early support, APIs may change
- deprecated: still present but planned for removal or replacement
