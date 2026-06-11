#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

if ! command -v terraform-docs >/dev/null 2>&1; then
  echo "ERROR: terraform-docs is required to generate module documentation." >&2
  echo "Install it from https://terraform-docs.io/ and rerun ./scripts/docs.sh." >&2
  exit 1
fi

updated=()
unchanged=()

while IFS= read -r versions_file; do
  module_dir="$(dirname "${versions_file}")"
  readme="${module_dir}/README.md"

  if [[ ! -f "${readme}" ]]; then
    echo "ERROR: ${module_dir} is missing README.md" >&2
    exit 1
  fi

  before="$(cksum "${readme}")"
  terraform-docs --config .terraform-docs.yml markdown table "${module_dir}"
  after="$(cksum "${readme}")"

  if [[ "${before}" != "${after}" ]]; then
    updated+=("${module_dir}")
    echo "UPDATED: ${module_dir}"
  else
    unchanged+=("${module_dir}")
    echo "UNCHANGED: ${module_dir}"
  fi
done < <(find modules -name versions.tf -not -path '*/.terraform/*' -print | sort)

echo "==> terraform-docs summary"
echo "Updated modules: ${#updated[@]}"
for module_dir in "${updated[@]}"; do
  echo "  - ${module_dir}"
done
echo "Unchanged modules: ${#unchanged[@]}"
