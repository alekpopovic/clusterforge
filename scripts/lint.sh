#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

TERRAFORM_BIN="${TERRAFORM_BIN:-terraform}"
GOCACHE="${GOCACHE:-/tmp/clusterforge-go-cache}"
GOPATH="${GOPATH:-/tmp/clusterforge-go}"

echo "==> Terraform fmt check"
"${TERRAFORM_BIN}" fmt -recursive -check -diff .

if command -v tflint >/dev/null 2>&1; then
  echo "==> TFLint"
  tflint --init --config .tflint.hcl
  tflint --recursive --config .tflint.hcl
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
