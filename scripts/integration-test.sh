#!/usr/bin/env bash
set -euo pipefail

target="${1:-}"

if [[ "${CLUSTERFORGE_RUN_INTEGRATION_TESTS:-}" != "true" ]]; then
  echo "ERROR: integration tests are opt-in. Set CLUSTERFORGE_RUN_INTEGRATION_TESTS=true." >&2
  exit 2
fi

case "${target}" in
  aws-eks|aws-ecs|existing-kubernetes) ;;
  *)
    echo "Usage: CLUSTERFORGE_RUN_INTEGRATION_TESTS=true $0 aws-eks|aws-ecs|existing-kubernetes" >&2
    exit 2
    ;;
esac

echo "WARNING: integration tests may create billable infrastructure."
echo "Target: ${target}"

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
tmpdir="$(mktemp -d "${TMPDIR:-/tmp}/clusterforge-it-${target}.XXXXXX")"
env_name="it-${target}-$(date -u +%Y%m%d%H%M%S)"
cleanup_needed=false

cleanup() {
  status=$?
  set +e
  if [[ "${cleanup_needed}" == "true" ]]; then
    echo "Attempting cleanup for ${env_name}"
    (cd "${tmpdir}" && "${repo_root}/cli/cf" destroy "${env_name}" --allow-destroy || true)
  fi
  rm -rf "${tmpdir}"
  exit "${status}"
}
trap cleanup EXIT

cd "${repo_root}/cli"
go build -o cf .

cd "${tmpdir}"
"${repo_root}/cli/cf" project init clusterforge-integration

case "${target}" in
  aws-eks)
    : "${AWS_REGION:?AWS_REGION is required}"
    "${repo_root}/cli/cf" env create "${env_name}" --cloud aws --orchestrator eks --region "${AWS_REGION}"
    "${repo_root}/cli/cf" generate "${env_name}"
    cp "live/${env_name}/aws-eks/terraform.tfvars.example" "live/${env_name}/aws-eks/terraform.tfvars"
    "${repo_root}/cli/cf" init "${env_name}"
    "${repo_root}/cli/cf" plan "${env_name}" --out ".cf/plans/${env_name}.tfplan" --risk-summary
    cleanup_needed=true
    "${repo_root}/cli/cf" apply "${env_name}" --plan-file ".cf/plans/${env_name}.tfplan"
    "${repo_root}/cli/cf" output "${env_name}" || true
    if command -v kubectl >/dev/null 2>&1; then kubectl get nodes || true; fi
    ;;
  aws-ecs)
    : "${AWS_REGION:?AWS_REGION is required}"
    "${repo_root}/cli/cf" env create "${env_name}" --cloud aws --orchestrator ecs --region "${AWS_REGION}"
    "${repo_root}/cli/cf" generate "${env_name}"
    cp "live/${env_name}/aws-ecs/terraform.tfvars.example" "live/${env_name}/aws-ecs/terraform.tfvars"
    "${repo_root}/cli/cf" init "${env_name}"
    "${repo_root}/cli/cf" plan "${env_name}" --out ".cf/plans/${env_name}.tfplan" --risk-summary
    cleanup_needed=true
    "${repo_root}/cli/cf" apply "${env_name}" --plan-file ".cf/plans/${env_name}.tfplan"
    "${repo_root}/cli/cf" output "${env_name}" || true
    ;;
  existing-kubernetes)
    : "${KUBECONFIG:?KUBECONFIG is required}"
    kubectl create namespace clusterforge-integration
    trap 'kubectl delete namespace clusterforge-integration --ignore-not-found; cleanup' EXIT
    terraform -chdir="${repo_root}/examples/kubernetes-basic-app" init
    terraform -chdir="${repo_root}/examples/kubernetes-basic-app" plan -out "${tmpdir}/k8s.tfplan" -var='namespace=clusterforge-integration'
    terraform -chdir="${repo_root}/examples/kubernetes-basic-app" apply "${tmpdir}/k8s.tfplan"
    kubectl -n clusterforge-integration get deployments,services,pods
    terraform -chdir="${repo_root}/examples/kubernetes-basic-app" destroy -auto-approve -var='namespace=clusterforge-integration'
    ;;
esac

echo "Integration target ${target} completed. Cleanup will run now."
