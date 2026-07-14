#trivy:ignore:DS-0002
FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

ARG TERRAFORM_VERSION=1.8.5
ARG OPENTOFU_VERSION=1.8.2
ARG TERRAFORM_DOCS_VERSION=0.18.0
ARG TFLINT_VERSION=0.52.0
ARG KUBECTL_VERSION=1.30.3
ARG HELM_VERSION=3.15.3
ARG TRIVY_VERSION=0.72.0

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN rm -f /etc/apt/sources.list.d/yarn.list \
    && apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates curl git make unzip \
    && architecture="$(dpkg --print-architecture)" \
    && case "${architecture}" in amd64) zip_arch=amd64; trivy_arch=64bit ;; arm64) zip_arch=arm64; trivy_arch=ARM64 ;; *) exit 1 ;; esac \
    && curl -fsSLo /tmp/terraform.zip "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_${zip_arch}.zip" \
    && unzip /tmp/terraform.zip -d /usr/local/bin \
    && curl -fsSLo /tmp/tofu.tar.gz "https://github.com/opentofu/opentofu/releases/download/v${OPENTOFU_VERSION}/tofu_${OPENTOFU_VERSION}_linux_${zip_arch}.tar.gz" \
    && tar -xzf /tmp/tofu.tar.gz -C /usr/local/bin tofu \
    && curl -fsSLo /tmp/terraform-docs.tar.gz "https://github.com/terraform-docs/terraform-docs/releases/download/v${TERRAFORM_DOCS_VERSION}/terraform-docs-v${TERRAFORM_DOCS_VERSION}-linux-${zip_arch}.tar.gz" \
    && tar -xzf /tmp/terraform-docs.tar.gz -C /usr/local/bin terraform-docs \
    && curl -fsSLo /tmp/tflint.zip "https://github.com/terraform-linters/tflint/releases/download/v${TFLINT_VERSION}/tflint_linux_${zip_arch}.zip" \
    && unzip /tmp/tflint.zip -d /usr/local/bin \
    && curl -fsSLo /usr/local/bin/kubectl "https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/${zip_arch}/kubectl" \
    && chmod +x /usr/local/bin/kubectl \
    && curl -fsSLo /tmp/helm.tar.gz "https://get.helm.sh/helm-v${HELM_VERSION}-linux-${zip_arch}.tar.gz" \
    && tar -xzf /tmp/helm.tar.gz -C /tmp \
    && mv "/tmp/linux-${zip_arch}/helm" /usr/local/bin/helm \
    && curl -fsSLo /tmp/trivy.tar.gz "https://github.com/aquasecurity/trivy/releases/download/v${TRIVY_VERSION}/trivy_${TRIVY_VERSION}_Linux-${trivy_arch}.tar.gz" \
    && tar -xzf /tmp/trivy.tar.gz -C /usr/local/bin trivy \
    && rm -rf /var/lib/apt/lists/* /tmp/*

WORKDIR /workspaces/clusterforge

USER vscode
