SHELL := /bin/bash
.SHELLFLAGS := --noprofile --norc -euo pipefail -c

TERRAFORM_BIN ?= terraform
GO ?= go
GOCACHE ?= /tmp/clusterforge-go-cache
GOPATH ?= /tmp/clusterforge-go
CLI_VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
CLI_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
CLI_DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
CLI_LDFLAGS := -s -w -X github.com/textracta/clusterforge/cli/cmd.Version=$(CLI_VERSION) -X github.com/textracta/clusterforge/cli/cmd.Commit=$(CLI_COMMIT) -X github.com/textracta/clusterforge/cli/cmd.Date=$(CLI_DATE)

.PHONY: help fmt fmt-check validate lint test test-cli test-terraform build-cli security docs clean ci

help: ## Show available targets.
	@awk 'BEGIN {FS = ":.*## "; printf "ClusterForge developer targets:\n\n"} /^[a-zA-Z0-9_-]+:.*## / {printf "  %-14s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

fmt: ## Format Terraform and Go files.
	$(TERRAFORM_BIN) fmt -recursive
	if [[ -d cli ]]; then find cli -name '*.go' -print0 | xargs -0 gofmt -w; fi

fmt-check: ## Check Terraform and Go formatting without modifying files.
	$(TERRAFORM_BIN) fmt -check -recursive
	if [[ -d cli ]]; then \
		unformatted="$$(find cli -name '*.go' -print0 | xargs -0 gofmt -l)"; \
		if [[ -n "$${unformatted}" ]]; then \
			echo "The following Go files need gofmt:"; \
			echo "$${unformatted}"; \
			exit 1; \
		fi; \
	fi

validate: ## Validate Terraform roots/modules and lightweight Terraform tests.
	TERRAFORM_BIN=$(TERRAFORM_BIN) ./scripts/validate.sh

lint: ## Run repository lint checks.
	./scripts/lint.sh

test: test-cli validate ## Run CLI tests and lightweight Terraform validation/tests.

test-cli: ## Run CLI unit tests and build check.
	GOCACHE=$(GOCACHE) GOPATH=$(GOPATH) ./scripts/test-cli.sh

test-terraform: ## Run Terraform native tests for safe provider-free modules.
	for module in modules/core/naming modules/core/tags modules/core/labels; do \
		echo "==> terraform test $${module}"; \
		$(TERRAFORM_BIN) -chdir=$${module} test -no-color; \
	done

build-cli: ## Build the ClusterForge CLI at cli/cf.
	cd cli && GOCACHE=$(GOCACHE) GOPATH=$(GOPATH) $(GO) build -trimpath -ldflags "$(CLI_LDFLAGS)" -o cf .

security: ## Run installed security scanners when available.
	./scripts/security.sh

docs: ## Generate module documentation when terraform-docs is installed.
	./scripts/docs.sh

clean: ## Remove local build artifacts and temporary generated files.
	rm -f cli/cf cli/clusterforge
	rm -rf cli/bin cli/dist cli/coverage
	rm -rf bin dist coverage
	rm -rf .cf/plans
	rm -rf /tmp/clusterforge-go-cache /tmp/clusterforge-go
	find . -name '*.tmp' -not -path './.git/*' -delete

ci: fmt-check test-cli validate ## Run the default credential-free CI checks.
