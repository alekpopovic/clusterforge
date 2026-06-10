#!/usr/bin/env sh
set -eu

TERRAFORM_BIN="${TERRAFORM_BIN:-terraform}"

"${TERRAFORM_BIN}" fmt -recursive -check -diff .
