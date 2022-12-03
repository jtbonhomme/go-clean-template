# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

linter-golangci: ### check by golangci linter
	golangci-lint -v --deadline 100s --skip-dirs docs run ./...
.PHONY: linter-golangci

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test

mock: ### run mockery
	mockery --all -r --case snake
.PHONY: mock
