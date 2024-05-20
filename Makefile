golangci_lint_cmd=golangci-lint
golangci_version=v1.57.2
gofumpt_cmd=gofumpt
gofumpt_version=v0.6.0

default: help

.PHONY: help
## help: Prints this help message
help: Makefile 
	@echo
	@echo "Available make commands:"
	@echo 
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: lint
## lint: Lint the repository
lint:
	@echo "--> Running linter"
	@if ! $(golangci_lint_cmd) --version 2>/dev/null | grep -q $(golangci_version); then \
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version); \
    fi
	@$(golangci_lint_cmd) run ./... --timeout 15m

.PHONY: lint-fix
## lint-fix: Lint the repository and fix warnings (if applicable)
lint-fix: 
	@echo "--> Running linter and fixing issues"
	@if ! $(golangci_lint_cmd) --version 2>/dev/null | grep -q $(golangci_version); then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version); \
	fi
	@$(golangci_lint_cmd) run ./... --fix --timeout 15m

.PHONY: gofumpt
## gofumpt: Format the code with gofumpt
gofumpt:
	@echo "--> Running gofumpt"
	@if ! $(gofumpt_cmd) -version 2>/dev/null | grep -q $(gofumpt_version); then \
		go install mvdan.cc/gofumpt@$(gofumpt_version); \
	fi
	@gofumpt -l -w .