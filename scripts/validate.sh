#!/usr/bin/env sh
set -eu

TERRAFORM_BIN="${TERRAFORM_BIN:-terraform}"

find . -name versions.tf -not -path '*/.terraform/*' -print | while IFS= read -r versions_file; do
  dir="$(dirname "${versions_file}")"
  echo "==> ${dir}"
  (
    cd "${dir}"
    "${TERRAFORM_BIN}" init -backend=false -input=false >/dev/null
    "${TERRAFORM_BIN}" validate
  )
done
