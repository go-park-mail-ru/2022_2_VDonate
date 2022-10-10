PROJECT_PATH = ./cmd/api/main.go

.PHONY: test
test: ## Run all the tests
	echo 'mode: atomic' > coverage.out && go test -covermode=atomic -coverprofile=coverage.out -v -race -timeout=30s ./...

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
	go tool cover -func=coverage.out
	got tool cover -html=coverage.out

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: local_build
local_build: ## Build locally
	go build ${PROJECT_PATH}

.PHONY: clean
clean: ## Remove temporary files
	rm -f main
	go clean

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := local_build