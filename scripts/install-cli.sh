#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

GO_BIN="${GO:-go}"
INSTALL_DIR="${INSTALL_DIR:-${1:-/usr/local/bin}}"
BINARY_NAME="${BINARY_NAME:-cf}"
VERSION="${VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo dev)}"
COMMIT="${COMMIT:-$(git rev-parse --short HEAD 2>/dev/null || echo unknown)}"
DATE="${DATE:-$(date -u +%Y-%m-%dT%H:%M:%SZ)}"
GOCACHE="${GOCACHE:-/tmp/clusterforge-go-cache}"
GOPATH="${GOPATH:-/tmp/clusterforge-go}"

LDFLAGS="-s -w -X github.com/alekpopovic/clusterforge/cli/cmd.Version=${VERSION} -X github.com/alekpopovic/clusterforge/cli/cmd.Commit=${COMMIT} -X github.com/alekpopovic/clusterforge/cli/cmd.Date=${DATE}"

tmpdir="$(mktemp -d)"
cleanup() {
  rm -rf "${tmpdir}"
}
trap cleanup EXIT

echo "==> Building ${BINARY_NAME}"
(
  cd cli
  GOCACHE="${GOCACHE}" GOPATH="${GOPATH}" "${GO_BIN}" build -trimpath -ldflags "${LDFLAGS}" -o "${tmpdir}/${BINARY_NAME}" .
)

mkdir -p "${INSTALL_DIR}" 2>/dev/null || true

install_path="${INSTALL_DIR}/${BINARY_NAME}"
if [[ -w "${INSTALL_DIR}" ]]; then
  install -m 0755 "${tmpdir}/${BINARY_NAME}" "${install_path}"
elif command -v sudo >/dev/null 2>&1; then
  echo "==> ${INSTALL_DIR} is not writable; using sudo for install"
  sudo mkdir -p "${INSTALL_DIR}"
  sudo install -m 0755 "${tmpdir}/${BINARY_NAME}" "${install_path}"
else
  echo "ERROR: ${INSTALL_DIR} is not writable and sudo is not available." >&2
  echo "Set INSTALL_DIR to a writable directory, for example:" >&2
  echo "  INSTALL_DIR=\"${HOME}/.local/bin\" ./scripts/install-cli.sh" >&2
  exit 1
fi

echo "==> Installed ${install_path}"
"${install_path}" version
