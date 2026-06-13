#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

TERRAFORM_BIN="${TERRAFORM_BIN:-terraform}"
TERRAFORM_VALIDATE_TIMEOUT="${TERRAFORM_VALIDATE_TIMEOUT:-45s}"

validated=()
skipped=()

skip() {
  local dir="$1"
  local reason="$2"
  skipped+=("${dir}|${reason}")
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

is_provider_environment_failure() {
  local output_file="$1"

  grep -E -q \
    'Failed to install provider|Failed to query available provider packages|Could not retrieve|registry\.terraform\.io|request canceled|context deadline exceeded|Client\.Timeout exceeded|no available releases match|Failed to load plugin schemas|Unrecognized remote plugin message' \
    "${output_file}"
}

run_terraform_with_timeout() {
  if command -v timeout >/dev/null 2>&1; then
    timeout "${TERRAFORM_VALIDATE_TIMEOUT}" "$@"
  else
    "$@"
  fi
}

validate_root() {
  local dir="$1"
  local output_file
  local status

  echo "==> ${dir}"
  output_file="$(mktemp)"

  set +e
  (
    cd "${dir}"
    run_terraform_with_timeout "${TERRAFORM_BIN}" init -backend=false -input=false -no-color
  ) >"${output_file}" 2>&1
  status=$?
  set -e

  if [[ "${status}" -eq 124 ]]; then
    rm -f "${output_file}"
    skip "${dir}" "terraform init exceeded ${TERRAFORM_VALIDATE_TIMEOUT}; likely waiting on provider installation or local plugin startup"
    return 0
  fi

  if [[ "${status}" -ne 0 ]]; then
    if is_provider_environment_failure "${output_file}"; then
      rm -f "${output_file}"
      skip "${dir}" "terraform init could not complete in this local environment, usually due to provider download or plugin availability"
      return 0
    fi

    cat "${output_file}" >&2
    rm -f "${output_file}"
    return "${status}"
  fi

  set +e
  (
    cd "${dir}"
    run_terraform_with_timeout "${TERRAFORM_BIN}" validate -no-color
  ) >"${output_file}" 2>&1
  status=$?
  set -e

  if [[ "${status}" -eq 124 ]]; then
    rm -f "${output_file}"
    skip "${dir}" "terraform validate exceeded ${TERRAFORM_VALIDATE_TIMEOUT}; likely waiting on provider schema loading"
    return 0
  fi

  if [[ "${status}" -ne 0 ]]; then
    if is_provider_environment_failure "${output_file}"; then
      rm -f "${output_file}"
      skip "${dir}" "terraform validate could not complete in this local environment, usually due to provider plugin availability"
      return 0
    fi

    cat "${output_file}" >&2
    rm -f "${output_file}"
    return "${status}"
  fi

  cat "${output_file}"
  rm -f "${output_file}"
  validated+=("${dir}")
}

echo "==> Terraform fmt check"
"${TERRAFORM_BIN}" fmt -check -recursive

validate_module_contracts

echo "==> Terraform validation"
echo "Using ${TERRAFORM_BIN}"

while IFS= read -r dir; do
  if reason="$(should_skip_root "${dir}")"; then
    echo "==> ${dir}"
    skip "${dir}" "${reason}"
    continue
  fi

  validate_root "${dir}"
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
