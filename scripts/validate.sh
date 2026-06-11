#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

TERRAFORM_BIN="${TERRAFORM_BIN:-terraform}"

echo "==> Terraform validation"
echo "Using ${TERRAFORM_BIN}"

find . -name versions.tf -not -path '*/.terraform/*' -print | sort | while IFS= read -r versions_file; do
  dir="$(dirname "${versions_file}")"
  echo "==> ${dir}"
  (
    cd "${dir}"
    "${TERRAFORM_BIN}" init -backend=false -input=false -no-color >/dev/null
    "${TERRAFORM_BIN}" validate -no-color
  )
done
