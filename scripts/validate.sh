#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

TERRAFORM_BIN="${TERRAFORM_BIN:-terraform}"

echo "==> Terraform fmt"
"${TERRAFORM_BIN}" fmt -recursive

echo "==> Terraform validation"
echo "Using ${TERRAFORM_BIN}"

# This validates all roots/modules that declare versions.tf. Provider init uses
# -backend=false and no real cloud credentials. Some examples still cannot be
# safely planned without live endpoints, so this script validates syntax and
# provider schemas rather than applying or refreshing remote infrastructure.
find . -name versions.tf -not -path '*/.terraform/*' -print | sort | while IFS= read -r versions_file; do
  dir="$(dirname "${versions_file}")"
  echo "==> ${dir}"
  (
    cd "${dir}"
    "${TERRAFORM_BIN}" init -backend=false -input=false -no-color >/dev/null
    "${TERRAFORM_BIN}" validate -no-color
  )
done

find modules -path '*/tests/*.tftest.hcl' -print | sort | while IFS= read -r test_file; do
  dir="$(dirname "$(dirname "${test_file}")")"
  echo "==> terraform test ${dir}"
  (
    cd "${dir}"
    "${TERRAFORM_BIN}" test -no-color
  )
done
