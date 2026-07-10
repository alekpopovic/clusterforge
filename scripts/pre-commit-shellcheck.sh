#!/usr/bin/env bash
set -euo pipefail

if ! command -v shellcheck >/dev/null 2>&1; then
  echo "SKIP: shellcheck is not installed"
  exit 0
fi

if [[ "$#" -eq 0 ]]; then
  exit 0
fi

exec shellcheck "$@"
