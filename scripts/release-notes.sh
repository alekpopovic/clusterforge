#!/usr/bin/env bash
set -euo pipefail

tag="${1:-}"
if [[ -z "${tag}" ]]; then
  echo "Usage: $0 <tag>" >&2
  exit 2
fi

echo "# ClusterForge ${tag}"
echo
if [[ -f CHANGELOG.md ]]; then
  awk -v tag="${tag}" '
    $0 ~ "^## " { if (seen) exit; if (index($0, tag) > 0) seen=1 }
    seen { print }
  ' CHANGELOG.md
else
  echo "See commit history for changes."
fi
echo
echo "## Artifacts"
echo
echo "CLI binaries and SHA256SUMS are attached to this release."
echo
echo "## Supply Chain"
echo
echo "- Checksums are generated for release artifacts."
echo "- SBOM and signing are planned; use docs/supply-chain-security.md for manual SBOM guidance."
