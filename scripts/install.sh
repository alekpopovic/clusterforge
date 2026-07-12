#!/usr/bin/env bash
set -euo pipefail

REPO="${CLUSTERFORGE_REPO:-alekpopovic/clusterforge}"
VERSION="${VERSION:-latest}"
INSTALL_DIR="${INSTALL_DIR:-${HOME}/.local/bin}"
BINARY_NAME="${BINARY_NAME:-cf}"

fail() {
  echo "ERROR: $*" >&2
  exit 1
}

command -v curl >/dev/null 2>&1 || fail "curl is required"

case "$(uname -s)" in
  Linux) os="linux" ;;
  Darwin) os="darwin" ;;
  *) fail "unsupported operating system: $(uname -s); use the release assets manually" ;;
esac

case "$(uname -m)" in
  x86_64 | amd64) arch="amd64" ;;
  arm64 | aarch64) arch="arm64" ;;
  *) fail "unsupported architecture: $(uname -m)" ;;
esac

artifact="cf-${os}-${arch}"
if [[ "${VERSION}" == "latest" ]]; then
  base_url="https://github.com/${REPO}/releases/latest/download"
else
  case "${VERSION}" in
    v*) ;;
    *) VERSION="v${VERSION}" ;;
  esac
  base_url="https://github.com/${REPO}/releases/download/${VERSION}"
fi
base_url="${CLUSTERFORGE_DOWNLOAD_BASE_URL:-${base_url}}"

tmpdir="$(mktemp -d)"
cleanup() {
  rm -rf "${tmpdir}"
}
trap cleanup EXIT

echo "==> Downloading ${artifact} (${VERSION})"
curl --fail --location --silent --show-error \
  --output "${tmpdir}/${artifact}" "${base_url}/${artifact}"
curl --fail --location --silent --show-error \
  --output "${tmpdir}/${artifact}.sha256" "${base_url}/${artifact}.sha256"

expected="$(awk 'NF { print $1; exit }' "${tmpdir}/${artifact}.sha256")"
[[ "${expected}" =~ ^[0-9a-fA-F]{64}$ ]] || fail "invalid checksum file"

if command -v sha256sum >/dev/null 2>&1; then
  actual="$(sha256sum "${tmpdir}/${artifact}" | awk '{print $1}')"
elif command -v shasum >/dev/null 2>&1; then
  actual="$(shasum -a 256 "${tmpdir}/${artifact}" | awk '{print $1}')"
else
  fail "sha256sum or shasum is required to verify the download"
fi

[[ "${actual}" == "${expected}" ]] || fail "checksum mismatch for ${artifact}"

mkdir -p "${INSTALL_DIR}" 2>/dev/null || true
install_path="${INSTALL_DIR}/${BINARY_NAME}"
if [[ -w "${INSTALL_DIR}" ]]; then
  install -m 0755 "${tmpdir}/${artifact}" "${install_path}"
elif command -v sudo >/dev/null 2>&1; then
  echo "==> ${INSTALL_DIR} is not writable; using sudo"
  sudo mkdir -p "${INSTALL_DIR}"
  sudo install -m 0755 "${tmpdir}/${artifact}" "${install_path}"
else
  fail "${INSTALL_DIR} is not writable; set INSTALL_DIR to a writable directory"
fi

echo "==> Installed ${install_path}"
"${install_path}" version

case ":${PATH}:" in
  *:"${INSTALL_DIR}":*) ;;
  *) echo "NOTE: add ${INSTALL_DIR} to PATH." ;;
esac
