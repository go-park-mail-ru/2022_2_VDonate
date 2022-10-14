PROJECT_PATH = ./cmd/api/main.go
MOCKS_DESTINATION = internal/mocks
INTERNAL_PATH = internal

.PHONY: test
test: ## Run all the tests
	go test -race -coverpkg=./... -coverprofile=c.out ./...

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
	go tool cover -func=c.out

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: local_build
local_build: ## Build locally
	go build ${PROJECT_PATH}

.PHONY: mocks
mocks: internal/auth/usecase/auth_usecase.go internal/posts/usecase/posts_usecase.go internal/users/usecase/user_usecase.go ## Generate mocks
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done

.PHONY: clean
clean: ## Remove temporary files
	rm -f main
	go clean

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := local_build