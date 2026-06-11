#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

export GOCACHE="${GOCACHE:-/tmp/clusterforge-go-cache}"
export GOPATH="${GOPATH:-/tmp/clusterforge-go}"

cleanup() {
  rm -f cli/cf
}

trap cleanup EXIT

echo "==> Go module download"
(
  cd cli
  go mod download
)

echo "==> gofmt check"
unformatted="$(find cli -name '*.go' -print0 | xargs -0 gofmt -l)"
if [[ -n "${unformatted}" ]]; then
  echo "The following Go files need gofmt:"
  echo "${unformatted}"
  exit 1
fi

echo "==> Go tests"
(
  cd cli
  go test ./...
)

echo "==> Go build"
(
  cd cli
  go build -o cf .
)
