# Nomad production MVP

Nomad support is experimental. ClusterForge renders server/client
configuration and cloud-init, connects to an existing Nomad API, and manages
service/batch Docker jobs. It does not provision or operate Nomad servers.

Production operators must own odd-sized server quorum, TLS, gossip encryption,
ACL bootstrap/rotation, host hardening, persistent data, snapshots, upgrades,
drain/rollback, networking and disaster recovery. Consul and ingress modules
only render integration metadata; they do not install those systems.

Use short-lived Nomad tokens through provider environment variables, keep
provider configuration in the root, pin images, define resources, test service
registration and restore procedures, and never put ACL tokens in tfvars/state.
