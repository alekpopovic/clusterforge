# Air-gapped deployment preparation

ClusterForge's initial air-gapped support creates a checksummed manifest bundle;
it does not download providers, images, charts, operating-system packages, or
cloud APIs. Build the bundle in a connected staging environment, review it, fetch
and verify approved artifacts through your supply-chain process, scan them, then
transfer them using the organization's controlled media workflow.

```bash
cf bundle create --env prod --output clusterforge-bundle
cf bundle inspect clusterforge-bundle
cf bundle verify clusterforge-bundle
```

At the restricted site, verify `SHA256SUMS` before use. Configure Terraform's
provider installation mirror, module source paths, Helm registry/repository
mirror, and container registry mirror explicitly. Disable unintended internet
fallback and test DNS, CA trust, time synchronization, revocation information,
identity bootstrap, upgrades, rollback and recovery while disconnected.

Checksums detect accidental or post-creation changes; they do not establish
publisher identity. Use signed upstream releases, image signatures/digests, an
approved internal artifact repository, malware scanning, SBOM review and a
separate signed transfer manifest for higher assurance. Never add credentials,
kubeconfigs, secret values, Terraform state, or plan files to the bundle.
