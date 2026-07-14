# Optional tool version management

ClusterForge provides `.tool-versions` for contributors who use asdf. It is
optional: the Makefile continues to use tools available on `PATH` and also
supports `TERRAFORM_BIN=tofu`.

## Install with asdf

Install asdf using its official instructions, then add plugins for the entries
you intend to use. Plugin names and sources can differ by asdf installation, so
review a plugin before installing it. A typical workflow is:

```bash
asdf plugin add golang
asdf plugin add terraform
asdf plugin add opentofu
asdf plugin add kubectl
asdf plugin add helm
asdf install
```

Add plugins for terraform-docs, TFLint, Checkov, Trivy, and pre-commit when
available in your trusted plugin catalog. Otherwise install those tools using
their upstream instructions. Running `asdf current` shows the selected tools.

Versions are pinned for reproducibility and satisfy
[`rel/VERSION_MATRIX.md`](../rel/VERSION_MATRIX.md); they are not a claim that
every listed upstream version is the newest release. When updating a version,
update `.tool-versions`, the devcontainer where applicable, and CI together,
then run `make ci`.
