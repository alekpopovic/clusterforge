#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

TERRAFORM_BIN="${TERRAFORM_BIN:-terraform}"

echo "==> Terraform fmt"
"${TERRAFORM_BIN}" fmt -recursive -check -diff .
