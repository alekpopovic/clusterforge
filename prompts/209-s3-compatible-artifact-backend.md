# Prompt 209 — S3-compatible artifact backend

```text
Add S3-compatible artifact storage backend.

Goal:
Support production artifact storage outside the database/local filesystem.

Config:
artifacts:
  backend: s3
  s3:
    bucket: clusterforge-artifacts
    region: eu-central-1
    prefix: clusterforge
    endpoint: ""
    force_path_style: false
    kms_key_id: ""
    presign_downloads: true
    presign_ttl: 15m

Features:
- upload artifact
- download artifact
- delete artifact
- checksum verification
- optional KMS encryption metadata
- object key scoped by organization/project/environment
- no public access

API:
- same artifact API as filesystem backend
- backend abstraction handles storage

Security:
- do not store AWS credentials in database
- credentials come from environment, IAM role, workload identity, or runner environment
- signed URLs require RBAC and audit
- raw plan upload remains disabled by default

Tests:
- mock S3 client or local fake
- object key generation
- upload/download/delete
- presigned URL path if implemented
- credentials not logged

Docs:
- docs/control-plane/artifacts-s3.md

Terraform:
- optional module:
  modules/cloud/aws/clusterforge-artifacts-bucket

Resources:
- S3 bucket
- versioning optional
- encryption
- public access block
- lifecycle expiration

Run:
- go test ./...
- terraform fmt -recursive
```
