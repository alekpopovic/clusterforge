# Reproducible development environment

The optional devcontainer supplies a consistent ClusterForge toolchain without
copying cloud credentials or host-specific configuration into the image.

## Open in a devcontainer

1. Install Docker, Visual Studio Code, and the Dev Containers extension.
2. Clone the repository and open its root in Visual Studio Code.
3. Run **Dev Containers: Reopen in Container** from the command palette.

The first build downloads pinned tools and can take several minutes. Rebuild
the container after changing `.devcontainer/Dockerfile`.

## Included tools

The image contains Go, Terraform, OpenTofu, terraform-docs, TFLint, Trivy,
kubectl, Helm, make, and Git. Versions are pinned in the Dockerfile and stay
within the support ranges in [`rel/VERSION_MATRIX.md`](../rel/VERSION_MATRIX.md).

Checkov is intentionally not baked into the image because its Python dependency
set adds substantial image weight. Install it only when needed:

```bash
python3 -m venv .venv-checkov
. .venv-checkov/bin/activate
pip install checkov
checkov --version
```

If `python3-venv` is unavailable on the host or base image, install Checkov
using your platform's trusted Python tooling instead of baking credentials or
user-specific package indexes into this repository.

Do not mount or copy cloud credentials into a shared image. When a real-cloud
test explicitly requires credentials, use a short-lived host credential flow
and remove the mount or environment variables afterwards.

## Without a devcontainer

The container is optional. Install supported tools locally and use the same
Make targets:

```bash
make fmt-check
make test
make validate
```

See [tool versions](tool-versions.md) for optional asdf setup.
