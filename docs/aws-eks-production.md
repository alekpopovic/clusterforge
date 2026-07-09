# AWS EKS Production Hardening

Use `modules/orchestrators/kubernetes/eks` as a readable Terraform module, then
make production intent explicit in the root configuration.

## API Endpoint Access

Private endpoint only:

```hcl
endpoint_public_access  = false
endpoint_private_access = true
public_access_cidrs     = []
```

Restricted public access:

```hcl
endpoint_public_access  = true
endpoint_private_access = true
public_access_cidrs     = ["203.0.113.0/24"] # Replace with approved operator CIDRs.
```

## Secrets Encryption

Create a module-owned KMS key:

```hcl
enable_cluster_encryption = true
create_kms_key            = true
```

Use an existing key:

```hcl
enable_cluster_encryption = true
kms_key_arn               = var.eks_secrets_kms_key_arn
```

You can create that key with `modules/cloud/aws/kms-key`:

```hcl
module "eks_secrets_key" {
  source = "../modules/cloud/aws/kms-key"

  name        = "clusterforge-prod-eks-secrets"
  environment = "prod"
  alias_name  = "clusterforge-prod-eks-secrets"
}
```

## Control Plane Logs

Enable the full control plane log set and retain logs deliberately:

```hcl
enabled_cluster_log_types  = ["api", "audit", "authenticator", "controllerManager", "scheduler"]
cluster_log_retention_days = 90
```

## Managed Node Groups

SSH remote access is disabled by default. Prefer SSM Session Manager or
Kubernetes-native debugging. If SSH is required, set `node_group_remote_access`
with a specific key and source security groups.

Rollout controls:

```hcl
node_group_update_config = {
  max_unavailable_percentage = 25
}
```

AMI controls:

```hcl
node_group_ami_type             = "AL2023_x86_64_STANDARD"
node_group_release_version      = null
node_group_force_update_version = false
```

## Deletion Protection

This module does not hide cluster deletion behind a custom shortcut. Protect
production deletion in root workflows with remote state locking, protected
branches, required plan files, and the ClusterForge CLI production destroy
guard.

See `examples/aws-eks-production-hardened` for a credential-free validation
root that demonstrates the hardening options.
