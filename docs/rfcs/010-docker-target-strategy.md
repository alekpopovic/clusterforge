# RFC 010: Docker target strategy

Decision: retain Docker Engine and Docker Swarm as **experimental** targets.
They serve local experiments, simple self-hosted installations, and
migration/legacy workloads. They are not recommended for large production
platforms.

The modules remain available and the CLI warns on environment creation and for
production doctor checks. Graduation would require tested host lifecycle,
multi-node networking, secrets, upgrades, rollback, backup and security
evidence. Deprecation remains an option if provider maintenance or user demand
does not justify that work.
