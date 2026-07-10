#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

mode="repository"
if [[ "${1:-}" == "--staged" ]]; then
  mode="staged"
elif [[ "$#" -gt 0 ]]; then
  echo "Usage: $0 [--staged]" >&2
  exit 2
fi

prohibited=0
check_path() {
  local path="$1"
  local base="${path##*/}"
  case "$path" in
    *.tfstate|*.tfstate.*|*/.terraform/*|.terraform/*|*/.kube/config|.kube/config)
      echo "ERROR: prohibited sensitive artifact is tracked or staged: $path" >&2
      prohibited=1
      ;;
    kubeconfig|kubeconfig.*|*/kubeconfig|*/kubeconfig.*|*.kubeconfig|*.kubeconfig.yaml)
      echo "ERROR: kubeconfig artifact is tracked or staged: $path" >&2
      prohibited=1
      ;;
  esac
  if [[ "$base" == ".env" || ( "$base" == .env.* && "$base" != ".env.example" ) ]]; then
    echo "ERROR: environment secret file is tracked or staged: $path" >&2
    prohibited=1
  fi
}

if [[ "$mode" == "staged" ]]; then
  while IFS= read -r -d '' path; do check_path "$path"; done < <(git diff --cached --name-only --diff-filter=ACMR -z)
else
  while IFS= read -r -d '' path; do check_path "$path"; done < <(git ls-files -z)
fi

if [[ "$prohibited" -ne 0 ]]; then
  exit 1
fi

if ! command -v gitleaks >/dev/null 2>&1; then
  if [[ "${REQUIRE_GITLEAKS:-false}" == "true" ]]; then
    echo "ERROR: gitleaks is required but not installed." >&2
    exit 1
  fi
  echo "WARN: gitleaks is not installed; prohibited-file checks passed, content scan skipped." >&2
  exit 0
fi

echo "==> Gitleaks ${mode} scan"
if [[ "$mode" == "staged" ]]; then
  exec gitleaks git --staged --redact --config .gitleaks.toml .
fi
exec gitleaks git --redact --config .gitleaks.toml .
