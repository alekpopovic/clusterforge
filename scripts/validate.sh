#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

TERRAFORM_BIN="${TERRAFORM_BIN:-terraform}"

validated=()
skipped=()

skip() {
  local dir="$1"
  local reason="$2"
  skipped+=("${dir}|${reason}")
  echo "==> ${dir}"
  echo "SKIP: ${reason}"
}

validate_module_contracts() {
  echo "==> Terraform module contract checks"

  local missing=0
  while IFS= read -r module_dir; do
    [[ -f "${module_dir}/main.tf" ]] || continue

    for required_file in main.tf variables.tf outputs.tf versions.tf README.md; do
      if [[ ! -f "${module_dir}/${required_file}" ]]; then
        echo "ERROR: ${module_dir} is missing ${required_file}" >&2
        missing=1
      fi
    done
  done < <(find modules -type d -not -path '*/.terraform/*' | sort)

  if [[ "${missing}" -ne 0 ]]; then
    return 1
  fi
}

should_skip_root() {
  local dir="$1"

  if [[ ! -f "${dir}/versions.tf" ]]; then
    echo "missing versions.tf; not treated as a complete Terraform root"
    return 0
  fi

  if grep -R -n -E '^[[:space:]]*cloud[[:space:]]*\{' "${dir}"/*.tf >/dev/null 2>&1; then
    echo "uses Terraform Cloud remote operations; validate manually with workspace access"
    return 0
  fi

  if grep -R -n -E '^[[:space:]]*backend[[:space:]]+"remote"' "${dir}"/*.tf >/dev/null 2>&1; then
    echo "uses a remote backend; validate manually with backend access"
    return 0
  fi

  return 1
}

echo "==> Terraform fmt check"
"${TERRAFORM_BIN}" fmt -check -recursive

validate_module_contracts

echo "==> Terraform validation"
echo "Using ${TERRAFORM_BIN}"

while IFS= read -r dir; do
  if reason="$(should_skip_root "${dir}")"; then
    skip "${dir}" "${reason}"
    continue
  fi

  echo "==> ${dir}"
  (
    cd "${dir}"
    "${TERRAFORM_BIN}" init -backend=false -input=false -no-color >/dev/null
    "${TERRAFORM_BIN}" validate -no-color
  )
  validated+=("${dir}")
done < <(./scripts/list-terraform-roots.sh)

find modules -path '*/tests/*.tftest.hcl' -print | sort | while IFS= read -r test_file; do
  dir="$(dirname "$(dirname "${test_file}")")"
  echo "==> terraform test ${dir}"
  (
    cd "${dir}"
    "${TERRAFORM_BIN}" test -no-color
  )
done

echo "==> Validation summary"
echo "Validated directories: ${#validated[@]}"
for dir in "${validated[@]}"; do
  echo "  - ${dir}"
done

if [[ "${#skipped[@]}" -eq 0 ]]; then
  echo "Skipped directories: 0"
else
  echo "Skipped directories: ${#skipped[@]}"
  for entry in "${skipped[@]}"; do
    dir="${entry%%|*}"
    reason="${entry#*|}"
    echo "  - ${dir}: ${reason}"
  done
fi
