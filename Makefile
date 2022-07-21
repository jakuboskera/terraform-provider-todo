.ONESHELL:
.SHELL := /bin/bash
.DEFAULT_GOAL := help

help:
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

pre-commit-install: ## Install pre-commit into your git hooks. After that pre-commit will now run on every commit.
	pre-commit install

pre-commit-all: ## Manually run all pre-commit hooks on a repository (all files).
	pre-commit run --all-files

build: ## Build terraform provider
	mkdir -p ~/.terraform.d/plugins/terraform.local/local/todo/0.1.0/darwin_amd64/ \
	&& go build -o ~/.terraform.d/plugins/terraform.local/local/todo/0.1.0/darwin_amd64/terraform-provider-todo_v0.1.0

tf-init: ## Initialize a Terraform project
	terraform -chdir=develop init

tf-apply: ## Create resources from develop/main.tf manifests
	terraform -chdir=develop apply -auto-approve

tf-destroy: ## Destroy resources from develop/main.tf manifests
	terraform -chdir=develop destroy -auto-approve

clean: ## Clean built things
	rm -rf ~/.terraform.d/plugins/terraform.local/local/todo

# New targets here
