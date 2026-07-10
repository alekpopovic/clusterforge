#!/usr/bin/env bash
set -euo pipefail

if ! command -v gitleaks >/dev/null 2>&1; then
  echo "SKIP: gitleaks is not installed"
  exit 0
fi

if gitleaks help git >/dev/null 2>&1; then
  exec gitleaks git --staged --redact .
fi

exec gitleaks protect --staged --redact
