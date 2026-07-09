#!/usr/bin/env bash
set -euo pipefail

if ! command -v mkdocs >/dev/null 2>&1; then
  echo "ERROR: mkdocs is required. Install mkdocs-material to build the docs site." >&2
  exit 1
fi

mkdocs build
