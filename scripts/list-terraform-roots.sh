#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

find . \
  -type f \
  -name '*.tf' \
  -not -path '*/.terraform/*' \
  -not -path './.git/*' \
  -not -path './.cf/*' \
  -not -path './tmp/*' \
  -not -path './.tmp/*' \
  -not -path './dist/*' \
  -not -path './bin/*' \
  -print \
  | while IFS= read -r terraform_file; do dirname "${terraform_file}"; done \
  | sort -u
