# Kubernetes With Ingress Example

Demonstrates creating a cert-manager `ClusterIssuer` for an ingress-based
Kubernetes setup.

This example assumes:

- A Kubernetes provider is configured by the caller or local environment.
- cert-manager is already installed.
- ingress-nginx is installed and uses the `nginx` ingress class.

## Usage

```bash
terraform init
terraform validate
terraform plan -var='acme_email=platform@example.com'
```

The example does not store ACME private keys manually. cert-manager manages the
ACME account key in the configured Kubernetes Secret.
