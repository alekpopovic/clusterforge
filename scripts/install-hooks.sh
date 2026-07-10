#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

if ! command -v pre-commit >/dev/null 2>&1; then
  echo "ERROR: pre-commit is not installed." >&2
  echo "Install it from https://pre-commit.com/ and rerun this script." >&2
  exit 1
fi

pre-commit install --install-hooks
echo "ClusterForge pre-commit hook installed."
