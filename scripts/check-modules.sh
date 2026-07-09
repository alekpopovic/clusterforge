#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

required=(main.tf variables.tf outputs.tf versions.tf README.md)

while IFS= read -r module_dir; do
  for file in "${required[@]}"; do
    if [[ ! -f "${module_dir}/${file}" ]]; then
      echo "ERROR: ${module_dir} missing ${file}" >&2
      exit 1
    fi
  done
  if grep -qi "placeholder" "${module_dir}/README.md"; then
    if ! grep -qi "status" "${module_dir}/README.md"; then
      echo "ERROR: ${module_dir} placeholder README must include status" >&2
      exit 1
    fi
  fi
done < <(find modules -mindepth 3 -name versions.tf -not -path '*/.terraform/*' -print | sed 's#/versions.tf$##' | sort)

echo "module contract checks passed"
