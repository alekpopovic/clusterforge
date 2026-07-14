# Prompt 251 — Customer-managed encryption keys RFC

```text
Create customer-managed encryption keys RFC.

Goal:
Design how ClusterForge could support organization-specific encryption keys for sensitive metadata and artifacts.

Create:
- docs/rfcs/029-customer-managed-keys.md
- docs/control-plane/customer-managed-keys.md

Cover:

1. Goals
   - per-organization encryption boundary
   - artifact encryption
   - token metadata encryption where useful
   - audit metadata protection
   - future SaaS readiness
   - key rotation

2. Non-goals
   - storing cloud provider root credentials
   - implementing every KMS provider immediately
   - encrypting Terraform state inside Control Plane, because state should not be stored there

3. Key providers
   - local development key
   - AWS KMS
   - Azure Key Vault future
   - GCP KMS future
   - HashiCorp Vault future

4. Data to encrypt
   - artifact payloads
   - sensitive metadata fields
   - webhook URLs if stored
   - notification secrets should use env references where possible
   - tokens are hashed, not encrypted

5. Data not to store
   - cloud credentials
   - secret values
   - kubeconfigs
   - Terraform state

6. Key hierarchy
   - platform master key
   - organization data key
   - artifact data key
   - envelope encryption

7. Rotation
   - create new key version
   - re-encrypt new artifacts
   - background re-encryption for old artifacts
   - key disabled/deleted behavior

8. Failure modes
   - key unavailable
   - access denied
   - key deleted
   - rotation partially complete

Do not implement code in this prompt.
Update roadmap.
```
