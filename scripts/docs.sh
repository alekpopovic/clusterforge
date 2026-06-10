#!/usr/bin/env sh
set -eu

if ! command -v terraform-docs >/dev/null 2>&1; then
  echo "terraform-docs is required to generate module documentation" >&2
  exit 1
fi

find modules -name versions.tf -print | while IFS= read -r versions_file; do
  dir="$(dirname "${versions_file}")"
  terraform-docs markdown table --output-file README.md --output-mode inject "${dir}"
done
