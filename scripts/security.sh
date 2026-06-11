#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

ran=0

if command -v checkov >/dev/null 2>&1; then
  ran=1
  echo "==> Checkov"
  checkov --config-file .checkov.yml -d .
else
  echo "WARN: checkov is not installed; skipping Checkov IaC scan." >&2
fi

if command -v trivy >/dev/null 2>&1; then
  ran=1
  echo "==> Trivy config scan"
  trivy config --config trivy.yaml .
else
  echo "WARN: trivy is not installed; skipping Trivy config scan." >&2
fi

if [[ "${ran}" -eq 0 ]]; then
  echo "WARN: no security scanners were available; install checkov and/or trivy for local security checks." >&2
fi
