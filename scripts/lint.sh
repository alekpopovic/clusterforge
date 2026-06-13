#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."
REPO_ROOT="$(pwd)"

TERRAFORM_BIN="${TERRAFORM_BIN:-terraform}"
GOCACHE="${GOCACHE:-/tmp/clusterforge-go-cache}"
GOPATH="${GOPATH:-/tmp/clusterforge-go}"

echo "==> Terraform fmt check"
"${TERRAFORM_BIN}" fmt -recursive -check -diff .

if command -v tflint >/dev/null 2>&1; then
  echo "==> TFLint"
  tflint_output="$(mktemp)"
  trap 'rm -f "${tflint_output}"' EXIT

  if ! tflint --init --config "${REPO_ROOT}/.tflint.hcl" 2>&1 | tee "${tflint_output}"; then
    if grep -q "Failed to initialize plugins" "${tflint_output}"; then
      echo "WARN: tflint plugin initialization failed; skipping Terraform static lint for this local run." >&2
    else
      exit 1
    fi
  elif ! tflint --recursive --config "${REPO_ROOT}/.tflint.hcl" 2>&1 | tee "${tflint_output}"; then
    if grep -q "Failed to initialize plugins" "${tflint_output}"; then
      echo "WARN: tflint plugin initialization failed; skipping Terraform static lint for this local run." >&2
    else
      exit 1
    fi
  fi
else
  echo "WARN: tflint is not installed; skipping Terraform static lint." >&2
fi

if [[ -d cli ]]; then
  echo "==> gofmt check"
  unformatted="$(find cli -name '*.go' -print0 | xargs -0 gofmt -l)"
  if [[ -n "${unformatted}" ]]; then
    echo "The following Go files need gofmt:"
    echo "${unformatted}"
    exit 1
  fi

  echo "==> go vet"
  (
    cd cli
    GOCACHE="${GOCACHE}" GOPATH="${GOPATH}" go vet ./...
  )
fi
