#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

root="modules/platform/kubernetes"
failures=0
warnings=0

fail() {
  echo "ERROR: $*" >&2
  failures=$((failures + 1))
}

warn() {
  echo "WARN: $*" >&2
  warnings=$((warnings + 1))
}

require_file() {
  local path="$1"
  [[ -f "${path}" ]] || fail "${path} is missing"
}

require_grep() {
  local pattern="$1"
  local path="$2"
  local message="$3"
  if ! grep -Eq "${pattern}" "${path}"; then
    fail "${message}"
  fi
}

check_helm_module() {
  local module="$1"
  local chart="$2"
  local dir="${root}/${module}"

  if [[ ! -d "${dir}" ]]; then
    fail "${dir} is missing"
    return
  fi

  for file in main.tf variables.tf outputs.tf README.md; do
    require_file "${dir}/${file}"
  done

  require_grep 'repository[[:space:]]*=' "${dir}/main.tf" "${module}: Helm repository must be set or documented"
  require_grep "chart[[:space:]]*=[[:space:]]*\"${chart}\"" "${dir}/main.tf" "${module}: expected chart ${chart}"
  require_grep 'variable "namespace"' "${dir}/variables.tf" "${module}: namespace input is missing"
  require_grep 'variable "chart_version"' "${dir}/variables.tf" "${module}: chart_version input is missing"
  require_grep 'variable "values"' "${dir}/variables.tf" "${module}: values input is missing"
  require_grep 'variable "create_namespace"' "${dir}/variables.tf" "${module}: create_namespace input is missing"
  require_grep 'output "release_name"' "${dir}/outputs.tf" "${module}: release_name output is missing"
  require_grep 'output "namespace"' "${dir}/outputs.tf" "${module}: namespace output is missing"
  require_grep '(^## Usage|module ")' "${dir}/README.md" "${module}: README usage example is missing"
}

check_bootstrap_module() {
  local module="$1"
  local slug="${module//-/_}"
  local bootstrap="${root}/bootstrap/main.tf"
  local variables="${root}/bootstrap/variables.tf"

  require_grep "module \"${slug}\"" "${bootstrap}" "bootstrap: module ${slug} is not wired"
  require_grep "source[[:space:]]*=[[:space:]]*\"\\.\\./${module}\"" "${bootstrap}" "bootstrap: ${slug} source is wrong"
  require_grep "enable_${slug}" "${variables}" "bootstrap: enable_${slug} input is missing"
  require_grep "namespace[[:space:]]*=[[:space:]]*local\\.namespaces\\.${slug}" "${bootstrap}" "bootstrap: ${slug} namespace is not wired"

  if grep -q "variable \"${slug}_chart_version\"" "${variables}"; then
    require_grep "chart_version[[:space:]]*=[[:space:]]*var\\.${slug}_chart_version" "${bootstrap}" "bootstrap: ${slug} chart_version input is not wired"
  else
    warn "bootstrap: ${slug}_chart_version pass-through is not declared"
  fi

  if grep -q "variable \"${slug}_values\"" "${variables}"; then
    require_grep "values[[:space:]]*=[[:space:]]*var\\.${slug}_values" "${bootstrap}" "bootstrap: ${slug} values input is not wired"
  else
    warn "bootstrap: ${slug}_values pass-through is not declared"
  fi
}

declare -A charts=(
  [ingress-nginx]="ingress-nginx"
  [cert-manager]="cert-manager"
  [external-dns]="external-dns"
  [external-secrets]="external-secrets"
  [metrics-server]="metrics-server"
  [prometheus-stack]="kube-prometheus-stack"
  [loki]="loki"
  [argocd]="argo-cd"
)

if [[ -d "${root}/karpenter" ]]; then
  charts[karpenter]="karpenter"
fi

for module in "${!charts[@]}"; do
  check_helm_module "${module}" "${charts[${module}]}"
  check_bootstrap_module "${module}"
done

if ((failures > 0)); then
  echo "platform module conformance failed: ${failures} error(s), ${warnings} warning(s)" >&2
  exit 1
fi

echo "platform module conformance passed: ${warnings} warning(s)"
