.PHONY: help lint fmt tools test t

help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "%-10s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

lint: ## Run golangci-lint
	golangci-lint run

fmt: ## Format and fix lint
	gofumpt -w .
	golangci-lint run --fix

tools: ## Install dev tools
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest

test: ## Run tests
	go test ./...

t: test ## Run tests
