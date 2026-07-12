# ClusterForge promptovi 000–200

Ovaj direktorijum sadrži kompletan, numerisan niz promptova za razvoj
ClusterForge projekta. Promptove je preporučljivo izvršavati redom jer kasniji
koraci često pretpostavljaju rezultate ranijih promptova.

Svaki unos ispod vodi direktno na odgovarajući Markdown fajl na GitHub `main`
grani. Pre izvršavanja proveri trenutno stanje repozitorijuma: već završene
zahteve ne treba ponavljati, a svaki prompt mora poštovati aktuelni
[`AGENTS.md`](https://github.com/alekpopovic/clusterforge/blob/main/AGENTS.md).

## Osnova i MVP (000–025)

Početna arhitektura repozitorijuma, Terraform moduli, CLI, CI/CD, dokumentacija i završni MVP pregled.

- [`000-master-prompt-za-ceo-projekat.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/000-master-prompt-za-ceo-projekat.md) — Master prompt za ceo projekat
- [`001-agents-md.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/001-agents-md.md) — Agents md
- [`002-repo-skeleton.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/002-repo-skeleton.md) — Repo skeleton
- [`003-terraform-standardi-i-module-templates.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/003-terraform-standardi-i-module-templates.md) — Terraform standardi i module templates
- [`004-core-naming-modul.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/004-core-naming-modul.md) — Core naming modul
- [`005-core-labels-tags-moduli.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/005-core-labels-tags-moduli.md) — Core labels tags moduli
- [`006-aws-network-modul.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/006-aws-network-modul.md) — AWS network modul
- [`007-eks-orchestrator-modul.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/007-eks-orchestrator-modul.md) — EKS orchestrator modul
- [`008-kubernetes-generic-provider-bootstrap-root-primer.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/008-kubernetes-generic-provider-bootstrap-root-primer.md) — Kubernetes generic provider bootstrap root primer
- [`009-kubernetes-platform-bootstrap.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/009-kubernetes-platform-bootstrap.md) — Kubernetes platform bootstrap
- [`010-kubernetes-workload-app-modul.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/010-kubernetes-workload-app-modul.md) — Kubernetes workload app modul
- [`011-kubernetes-cronjob-modul.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/011-kubernetes-cronjob-modul.md) — Kubernetes cronjob modul
- [`012-ecs-cluster-modul.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/012-ecs-cluster-modul.md) — ECS cluster modul
- [`013-ecs-service-workload-modul.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/013-ecs-service-workload-modul.md) — ECS service workload modul
- [`014-nomad-job-modul.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/014-nomad-job-modul.md) — Nomad job modul
- [`015-docker-container-i-swarm-service-moduli.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/015-docker-container-i-swarm-service-moduli.md) — Docker container i swarm service moduli
- [`016-cli-osnova-u-go-cobra.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/016-cli-osnova-u-go-cobra.md) — CLI osnova u go cobra
- [`017-cli-config-loader.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/017-cli-config-loader.md) — CLI config loader
- [`018-cli-terraform-runner.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/018-cli-terraform-runner.md) — CLI Terraform runner
- [`019-cli-generator-terraform-fajlova.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/019-cli-generator-terraform-fajlova.md) — CLI generator Terraform fajlova
- [`020-app-manifest-i-app-generator.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/020-app-manifest-i-app-generator.md) — App manifest i app generator
- [`021-policy-i-risk-summary.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/021-policy-i-risk-summary.md) — Policy i risk summary
- [`022-ci-cd-workflows.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/022-ci-cd-workflows.md) — CI CD workflows
- [`023-dokumentacija.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/023-dokumentacija.md) — Dokumentacija
- [`024-testovi-i-examples.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/024-testovi-i-examples.md) — Testovi i examples
- [`025-zavr-ni-refactor-i-review.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/025-zavr-ni-refactor-i-review.md) — Zavr ni refactor i review

## Stabilizacija i produkcioni temelji (026–050)

Validacija, security skeneri, hardening, GitOps, observability, CLI automatizacija i MVP pakovanje.

- [`026-audit-trenutnog-stanja-projekta.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/026-audit-trenutnog-stanja-projekta.md) — Audit trenutnog stanja projekta
- [`027-makefile-developer-workflow.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/027-makefile-developer-workflow.md) — Makefile developer workflow
- [`028-stabilizuj-terraform-validaciju.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/028-stabilizuj-terraform-validaciju.md) — Stabilizuj Terraform validaciju
- [`029-automatska-dokumentacija-za-terraform-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/029-automatska-dokumentacija-za-terraform-module.md) — Automatska dokumentacija za Terraform module
- [`030-tflint-checkov-i-trivy-config-scanning.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/030-tflint-checkov-i-trivy-config-scanning.md) — TFLint Checkov i Trivy config scanning
- [`031-eks-hardening-oidc-irsa-i-add-on-role-support.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/031-eks-hardening-oidc-irsa-i-add-on-role-support.md) — EKS hardening OIDC IRSA i add on role support
- [`032-external-secrets-operator-integracija.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/032-external-secrets-operator-integracija.md) — External secrets operator integracija
- [`033-argo-cd-gitops-module-i-app-of-apps.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/033-argo-cd-gitops-module-i-app-of-apps.md) — Argo CD gitops module i app of apps
- [`034-kubernetes-autoscaling-cluster-autoscaler-ili-karpenter.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/034-kubernetes-autoscaling-cluster-autoscaler-ili-karpenter.md) — Kubernetes autoscaling cluster autoscaler ili karpenter
- [`035-ecs-alb-module-i-povezivanje-sa-ecs-servisom.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/035-ecs-alb-module-i-povezivanje-sa-ecs-servisom.md) — ECS ALB module i povezivanje sa ECS servisom
- [`036-route53-dns-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/036-route53-dns-module.md) — Route 53 DNS module
- [`037-cert-manager-clusterissuer-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/037-cert-manager-clusterissuer-module.md) — Cert manager clusterissuer module
- [`038-cli-install-shell-completion-i-release-build.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/038-cli-install-shell-completion-i-release-build.md) — CLI install shell completion i release build
- [`039-cli-interactive-wizard.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/039-cli-interactive-wizard.md) — CLI interactive wizard
- [`040-app-manifest-schema-validation.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/040-app-manifest-schema-validation.md) — App manifest schema validation
- [`041-environment-manifest-i-multi-stack-layout.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/041-environment-manifest-i-multi-stack-layout.md) — Environment manifest i multi stack layout
- [`042-remote-backend-generator.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/042-remote-backend-generator.md) — Remote backend generator
- [`043-backend-bootstrap-module-za-aws-tfstate.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/043-backend-bootstrap-module-za-aws-tfstate.md) — Backend bootstrap module za AWS tfstate
- [`044-observability-stack-prometheus-loki-grafana-values.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/044-observability-stack-prometheus-loki-grafana-values.md) — Observability stack prometheus loki grafana values
- [`045-kubernetes-worker-workload-modul.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/045-kubernetes-worker-workload-modul.md) — Kubernetes worker workload modul
- [`046-workload-module-helm-app-wrapper.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/046-workload-module-helm-app-wrapper.md) — Workload module helm app wrapper
- [`047-cli-doctor-command-hardening.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/047-cli-doctor-command-hardening.md) — CLI doctor command hardening
- [`048-cli-json-output-za-automatizaciju.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/048-cli-json-output-za-automatizaciju.md) — CLI JSON output za automatizaciju
- [`049-pod-security-i-networkpolicy-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/049-pod-security-i-networkpolicy-module.md) — Pod security i networkpolicy module
- [`050-final-mvp-packaging.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/050-final-mvp-packaging.md) — Final MVP packaging

## Testiranje, distribucija i proširenje platformi (051–080)

Acceptance i cloud testovi, release/supply-chain proces, migracije, Azure/GCP/self-hosted podrška, plugin strategija i v0.2 plan.

- [`051-real-mvp-acceptance-test.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/051-real-mvp-acceptance-test.md) — Real MVP acceptance test
- [`052-terraform-native-tests-for-core-modules.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/052-terraform-native-tests-for-core-modules.md) — Terraform native tests for core modules
- [`053-plan-mode-tests-for-aws-modules.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/053-plan-mode-tests-for-aws-modules.md) — Plan mode tests for AWS modules
- [`054-real-cloud-smoke-test-runbook.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/054-real-cloud-smoke-test-runbook.md) — Real cloud smoke test runbook
- [`055-ephemeral-integration-test-harness.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/055-ephemeral-integration-test-harness.md) — Ephemeral integration test harness
- [`056-version-support-matrix.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/056-version-support-matrix.md) — Version support matrix
- [`057-module-release-packaging.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/057-module-release-packaging.md) — Module release packaging
- [`058-github-release-automation-for-v0-1-x.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/058-github-release-automation-for-v0-1-x.md) — GitHub release automation for v0 1 x
- [`059-supply-chain-security-baseline.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/059-supply-chain-security-baseline.md) — Supply chain security baseline
- [`060-public-private-module-registry-strategy.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/060-public-private-module-registry-strategy.md) — Public private module registry strategy
- [`061-drift-detection-command.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/061-drift-detection-command.md) — Drift detection command
- [`062-state-inspection-and-safety-helpers.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/062-state-inspection-and-safety-helpers.md) — State inspection and safety helpers
- [`063-upgrade-and-migration-framework.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/063-upgrade-and-migration-framework.md) — Upgrade and migration framework
- [`064-import-adopt-existing-infrastructure-strategy.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/064-import-adopt-existing-infrastructure-strategy.md) — Import adopt existing infrastructure strategy
- [`065-environment-promotion-workflow.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/065-environment-promotion-workflow.md) — Environment promotion workflow
- [`066-aks-module-design-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/066-aks-module-design-rfc.md) — AKS module design RFC
- [`067-implement-azure-network-and-aks-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/067-implement-azure-network-and-aks-mvp.md) — Implement azure network and AKS MVP
- [`068-gke-module-design-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/068-gke-module-design-rfc.md) — GKE module design RFC
- [`069-implement-gcp-network-and-gke-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/069-implement-gcp-network-and-gke-mvp.md) — Implement gcp network and GKE MVP
- [`070-k3s-and-rke2-self-hosted-kubernetes-support.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/070-k3s-and-rke2-self-hosted-kubernetes-support.md) — K3s and RKE2 self hosted Kubernetes support
- [`071-enterprise-policy-packs.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/071-enterprise-policy-packs.md) — Enterprise policy packs
- [`072-rbac-and-service-account-workload-support.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/072-rbac-and-service-account-workload-support.md) — RBAC and service account workload support
- [`073-advanced-kubernetes-workload-features.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/073-advanced-kubernetes-workload-features.md) — Advanced Kubernetes workload features
- [`074-ecs-blue-green-deployment-design.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/074-ecs-blue-green-deployment-design.md) — ECS blue green deployment design
- [`075-cost-estimation-hooks.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/075-cost-estimation-hooks.md) — Cost estimation hooks
- [`076-plugin-architecture-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/076-plugin-architecture-rfc.md) — Plugin architecture RFC
- [`077-template-pack-support.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/077-template-pack-support.md) — Template pack support
- [`078-product-website-docs-skeleton.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/078-product-website-docs-skeleton.md) — Product website docs skeleton
- [`079-user-onboarding-examples.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/079-user-onboarding-examples.md) — User onboarding examples
- [`080-v0-2-0-planning-and-milestone-board.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/080-v0-2-0-planning-and-milestone-board.md) — V0 2 0 planning and milestone board

## Conformance, bezbednost i napredne platforme (081–120)

Release gate, lokalni Kubernetes, compatibility i golden testovi, AWS/Kubernetes proširenja, fleet operacije, developer tooling i v0.3 plan.

- [`081-v0-2-release-gate-review.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/081-v0-2-release-gate-review.md) — V0 2 release gate review
- [`082-local-kubernetes-development-target-with-kind-or-k3d.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/082-local-kubernetes-development-target-with-kind-or-k3d.md) — Local Kubernetes development target with kind or k3d
- [`083-existing-kubernetes-environment-support.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/083-existing-kubernetes-environment-support.md) — Existing Kubernetes environment support
- [`084-provider-compatibility-matrix-ci.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/084-provider-compatibility-matrix-ci.md) — Provider compatibility matrix CI
- [`085-golden-tests-for-cli-generators.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/085-golden-tests-for-cli-generators.md) — Golden tests for CLI generators
- [`086-cli-end-to-end-non-cloud-tests.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/086-cli-end-to-end-non-cloud-tests.md) — CLI end to end non cloud tests
- [`087-module-conformance-checker.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/087-module-conformance-checker.md) — Module conformance checker
- [`088-platform-conformance-tests-for-kubernetes-add-ons.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/088-platform-conformance-tests-for-kubernetes-add-ons.md) — Platform conformance tests for Kubernetes add ons
- [`089-eks-production-hardening-options.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/089-eks-production-hardening-options.md) — EKS production hardening options
- [`090-aws-kms-reusable-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/090-aws-kms-reusable-module.md) — AWS KMS reusable module
- [`091-aws-vpc-endpoints-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/091-aws-vpc-endpoints-module.md) — AWS vpc endpoints module
- [`092-ecr-registry-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/092-ecr-registry-module.md) — ECR registry module
- [`093-container-image-security-workflow.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/093-container-image-security-workflow.md) — Container image security workflow
- [`094-velero-backup-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/094-velero-backup-module.md) — Velero backup module
- [`095-disaster-recovery-runbooks.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/095-disaster-recovery-runbooks.md) — Disaster recovery runbooks
- [`096-external-dns-production-hardening.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/096-external-dns-production-hardening.md) — External DNS production hardening
- [`097-cert-manager-route53-dns01-iam-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/097-cert-manager-route53-dns01-iam-module.md) — Cert manager Route 53 dns01 IAM module
- [`098-aws-rds-postgresql-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/098-aws-rds-postgresql-module.md) — AWS RDS postgresql module
- [`099-aws-elasticache-redis-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/099-aws-elasticache-redis-module.md) — AWS elasticache redis module
- [`100-aws-messaging-modules-sqs-and-sns.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/100-aws-messaging-modules-sqs-and-sns.md) — AWS messaging modules SQS and SNS
- [`101-workload-cloud-identity-abstraction.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/101-workload-cloud-identity-abstraction.md) — Workload cloud identity abstraction
- [`102-service-binding-pattern-for-apps.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/102-service-binding-pattern-for-apps.md) — Service binding pattern for apps
- [`103-kubernetes-tenant-model.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/103-kubernetes-tenant-model.md) — Kubernetes tenant model
- [`104-resourcequota-and-limitrange-baseline-modules.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/104-resourcequota-and-limitrange-baseline-modules.md) — Resourcequota and limitrange baseline modules
- [`105-kyverno-policy-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/105-kyverno-policy-module.md) — Kyverno policy module
- [`106-opa-gatekeeper-alternative-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/106-opa-gatekeeper-alternative-module.md) — OPA Gatekeeper alternative module
- [`107-progressive-delivery-with-argo-rollouts.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/107-progressive-delivery-with-argo-rollouts.md) — Progressive delivery with argo rollouts
- [`108-service-mesh-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/108-service-mesh-rfc.md) — Service mesh RFC
- [`109-multi-cluster-inventory-model.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/109-multi-cluster-inventory-model.md) — Multi cluster inventory model
- [`110-fleet-operations-cli.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/110-fleet-operations-cli.md) — Fleet operations CLI
- [`111-environment-graph-visualization.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/111-environment-graph-visualization.md) — Environment graph visualization
- [`112-scheduled-drift-check-workflow-templates.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/112-scheduled-drift-check-workflow-templates.md) — Scheduled drift check workflow templates
- [`113-cli-audit-log.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/113-cli-audit-log.md) — CLI audit log
- [`114-pre-commit-hooks.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/114-pre-commit-hooks.md) — Pre commit hooks
- [`115-secret-scanning-baseline.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/115-secret-scanning-baseline.md) — Secret scanning baseline
- [`116-devcontainer-and-reproducible-dev-environment.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/116-devcontainer-and-reproducible-dev-environment.md) — Devcontainer and reproducible dev environment
- [`117-nix-flake-or-asdf-tool-versions.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/117-nix-flake-or-asdf-tool-versions.md) — Nix flake or asdf tool versions
- [`118-github-issue-and-pr-templates.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/118-github-issue-and-pr-templates.md) — GitHub issue and pr templates
- [`119-security-threat-model.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/119-security-threat-model.md) — Security threat model
- [`120-v0-3-release-planning.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/120-v0-3-release-planning.md) — V0 3 release planning

## Enterprise i operativne funkcije (121–160)

Plugin i policy engine, organizacioni model, upgrade planeri, enterprise integracije, inventar, compliance, DR, edge/offline procene i v0.4 izdanje.

- [`121-v0-3-release-gate-review.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/121-v0-3-release-gate-review.md) — V0 3 release gate review
- [`122-cli-plugin-system-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/122-cli-plugin-system-mvp.md) — CLI plugin system MVP
- [`123-template-pack-registry-support.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/123-template-pack-registry-support.md) — Template pack registry support
- [`124-policy-engine-v2.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/124-policy-engine-v2.md) — Policy engine v2
- [`125-organization-and-workspace-model.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/125-organization-and-workspace-model.md) — Organization and workspace model
- [`126-aws-multi-account-strategy.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/126-aws-multi-account-strategy.md) — AWS multi account strategy
- [`127-multi-region-environment-strategy.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/127-multi-region-environment-strategy.md) — Multi region environment strategy
- [`128-kubernetes-upgrade-planner.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/128-kubernetes-upgrade-planner.md) — Kubernetes upgrade planner
- [`129-platform-add-on-upgrade-planner.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/129-platform-add-on-upgrade-planner.md) — Platform add on upgrade planner
- [`130-terraform-opentofu-execution-profiles.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/130-terraform-opentofu-execution-profiles.md) — Terraform OpenTofu execution profiles
- [`131-terraform-cloud-hcp-terraform-integration.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/131-terraform-cloud-hcp-terraform-integration.md) — Terraform cloud hcp Terraform integration
- [`132-gitlab-ci-templates.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/132-gitlab-ci-templates.md) — GitLab CI templates
- [`133-azure-and-gcp-production-hardening-docs.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/133-azure-and-gcp-production-hardening-docs.md) — Azure and gcp production hardening docs
- [`134-nomad-production-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/134-nomad-production-mvp.md) — Nomad production MVP
- [`135-docker-target-policy-and-lifecycle-decision.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/135-docker-target-policy-and-lifecycle-decision.md) — Docker target policy and lifecycle decision
- [`136-slo-and-platform-health-model.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/136-slo-and-platform-health-model.md) — SLO and platform health model
- [`137-opentelemetry-platform-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/137-opentelemetry-platform-module.md) — Opentelemetry platform module
- [`138-incident-response-documentation.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/138-incident-response-documentation.md) — Incident response documentation
- [`139-runbook-cli-scaffolding.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/139-runbook-cli-scaffolding.md) — Runbook CLI scaffolding
- [`140-finops-v2-with-infracost-integration.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/140-finops-v2-with-infracost-integration.md) — FinOps v2 with Infracost integration
- [`141-cloud-asset-inventory-export.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/141-cloud-asset-inventory-export.md) — Cloud asset inventory export
- [`142-backstage-integration.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/142-backstage-integration.md) — Backstage integration
- [`143-service-catalog-manifest.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/143-service-catalog-manifest.md) — Service catalog manifest
- [`144-platform-api-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/144-platform-api-rfc.md) — Platform API RFC
- [`145-web-dashboard-prototype-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/145-web-dashboard-prototype-rfc.md) — Web dashboard prototype RFC
- [`146-dashboard-data-export.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/146-dashboard-data-export.md) — Dashboard data export
- [`147-audit-event-export-and-siem-integration.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/147-audit-event-export-and-siem-integration.md) — Audit event export and SIEM integration
- [`148-secret-rotation-workflow.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/148-secret-rotation-workflow.md) — Secret rotation workflow
- [`149-kubernetes-admission-security-extended-pack.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/149-kubernetes-admission-security-extended-pack.md) — Kubernetes admission security extended pack
- [`150-compliance-mapping-packs.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/150-compliance-mapping-packs.md) — Compliance mapping packs
- [`151-backup-validation-tests.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/151-backup-validation-tests.md) — Backup validation tests
- [`152-cross-cluster-gitops-support.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/152-cross-cluster-gitops-support.md) — Cross cluster gitops support
- [`153-cluster-federation-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/153-cluster-federation-rfc.md) — Cluster federation RFC
- [`154-edge-deployment-support-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/154-edge-deployment-support-rfc.md) — Edge deployment support RFC
- [`155-air-gapped-and-offline-support.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/155-air-gapped-and-offline-support.md) — Air gapped and offline support
- [`156-windows-containers-support-assessment.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/156-windows-containers-support-assessment.md) — Windows containers support assessment
- [`157-migration-analyzer-for-existing-terraform-repos.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/157-migration-analyzer-for-existing-terraform-repos.md) — Migration analyzer for existing Terraform repos
- [`158-project-scaffolding-wizard-v2.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/158-project-scaffolding-wizard-v2.md) — Project scaffolding wizard v2
- [`159-v0-4-roadmap-and-scope.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/159-v0-4-roadmap-and-scope.md) — V0 4 roadmap and scope
- [`160-v0-4-release-candidate-packaging.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/160-v0-4-release-candidate-packaging.md) — V0 4 release candidate packaging

## Control Plane i v0.5 (161–200)

Finalni v0.4 gate, Control Plane API/runner/dashboard, Git integracije, bezbednost, deployment, observability, testiranje i v0.5 izdanje.

- [`161-v0-4-final-release-gate-review.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/161-v0-4-final-release-gate-review.md) — V0 4 final release gate review
- [`162-clusterforge-control-plane-architecture-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/162-clusterforge-control-plane-architecture-rfc.md) — Clusterforge control plane architecture RFC
- [`163-control-plane-api-server-scaffold.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/163-control-plane-api-server-scaffold.md) — Control plane API server scaffold
- [`164-control-plane-database-schema-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/164-control-plane-database-schema-mvp.md) — Control plane database schema MVP
- [`165-control-plane-rest-api-resources.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/165-control-plane-rest-api-resources.md) — Control plane rest API resources
- [`166-control-plane-authentication-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/166-control-plane-authentication-mvp.md) — Control plane authentication MVP
- [`167-cli-integration-with-control-plane.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/167-cli-integration-with-control-plane.md) — CLI integration with control plane
- [`168-runner-architecture-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/168-runner-architecture-rfc.md) — Runner architecture RFC
- [`169-runner-agent-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/169-runner-agent-mvp.md) — Runner agent MVP
- [`170-plan-request-workflow.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/170-plan-request-workflow.md) — Plan request workflow
- [`171-approval-workflow-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/171-approval-workflow-mvp.md) — Approval workflow MVP
- [`172-apply-job-execution-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/172-apply-job-execution-mvp.md) — Apply job execution MVP
- [`173-server-side-audit-trail.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/173-server-side-audit-trail.md) — Server side audit trail
- [`174-notification-system-mvp.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/174-notification-system-mvp.md) — Notification system MVP
- [`175-dashboard-mvp-scaffold.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/175-dashboard-mvp-scaffold.md) — Dashboard MVP scaffold
- [`176-dashboard-inventory-pages.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/176-dashboard-inventory-pages.md) — Dashboard inventory pages
- [`177-dashboard-operations-pages.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/177-dashboard-operations-pages.md) — Dashboard operations pages
- [`178-service-catalog-api-and-dashboard.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/178-service-catalog-api-and-dashboard.md) — Service catalog API and dashboard
- [`179-runbook-api-and-dashboard.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/179-runbook-api-and-dashboard.md) — Runbook API and dashboard
- [`180-git-provider-integration-rfc.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/180-git-provider-integration-rfc.md) — Git provider integration RFC
- [`181-github-pr-plan-comments.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/181-github-pr-plan-comments.md) — GitHub pr plan comments
- [`182-gitlab-merge-request-plan-comments.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/182-gitlab-merge-request-plan-comments.md) — GitLab merge request plan comments
- [`183-sarif-and-code-scanning-integration.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/183-sarif-and-code-scanning-integration.md) — SARIF and code scanning integration
- [`184-secrets-reference-inventory.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/184-secrets-reference-inventory.md) — Secrets reference inventory
- [`185-scheduled-drift-checks-in-control-plane.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/185-scheduled-drift-checks-in-control-plane.md) — Scheduled drift checks in control plane
- [`186-scheduled-cost-reports-in-control-plane.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/186-scheduled-cost-reports-in-control-plane.md) — Scheduled cost reports in control plane
- [`187-runner-deployment-on-kubernetes.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/187-runner-deployment-on-kubernetes.md) — Runner deployment on Kubernetes
- [`188-control-plane-helm-chart.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/188-control-plane-helm-chart.md) — Control plane helm chart
- [`189-terraform-module-for-control-plane-deployment.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/189-terraform-module-for-control-plane-deployment.md) — Terraform module for control plane deployment
- [`190-control-plane-external-database-module.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/190-control-plane-external-database-module.md) — Control plane external database module
- [`191-control-plane-observability.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/191-control-plane-observability.md) — Control plane observability
- [`192-control-plane-backup-and-restore.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/192-control-plane-backup-and-restore.md) — Control plane backup and restore
- [`193-control-plane-e2e-tests.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/193-control-plane-e2e-tests.md) — Control plane E2E tests
- [`194-control-plane-load-and-reliability-tests.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/194-control-plane-load-and-reliability-tests.md) — Control plane load and reliability tests
- [`195-control-plane-security-hardening.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/195-control-plane-security-hardening.md) — Control plane security hardening
- [`196-docker-images-and-container-release.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/196-docker-images-and-container-release.md) — Docker images and container release
- [`197-container-signing-and-provenance-plan.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/197-container-signing-and-provenance-plan.md) — Container signing and provenance plan
- [`198-control-plane-documentation-site-section.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/198-control-plane-documentation-site-section.md) — Control plane documentation site section
- [`199-v0-5-roadmap-and-release-plan.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/199-v0-5-roadmap-and-release-plan.md) — V0 5 roadmap and release plan
- [`200-v0-5-release-candidate-packaging.md`](https://github.com/alekpopovic/clusterforge/blob/main/prompts/200-v0-5-release-candidate-packaging.md) — V0 5 release candidate packaging

## Način korišćenja

1. Otvori sledeći prompt preko GitHub linka.
2. Proveri da li su njegovi preduslovi već implementirani.
3. Izvrši samo nedostajući opseg i zadrži postojeće korisničke izmene.
4. Pokreni provere propisane promptom i `AGENTS.md` dokumentom.
5. Evidentiraj rezultat, commit i preostale blokere pre prelaska na sledeći prompt.
