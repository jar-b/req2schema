.DEFAULT_GOAL:=help

.PHONY: build
build: clean ## Build binaries
	@go build

.PHONY: clean
clean: ## Clean up binaries
	@rm -f req2schema coverage.txt

.PHONY: cover
cover: test ## Display test coverage report
	@go tool cover -func=coverage.txt

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Run unit tests
	@go test -v -coverprofile=coverage.txt ./...
