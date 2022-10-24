MAIN_PATH = ./cmd/api/main.go
MOCKS_DESTINATION = internal/mocks
INTERNAL_PATH = internal
ACTIVE_PACKAGES = $(shell go list ./... | grep -v "/mocks/" | tr '\n' ',')

.PHONY: test
test: ## Run all the tests
	go test -coverpkg=$(ACTIVE_PACKAGES) -coverprofile=c.out ./...

.PHONY: cover_out
cover_out: test ## Run all the tests and opens the coverage report
	go tool cover -func=c.out

.PHONY: cover_html
cover_html: test ## Run all the tests and opens the coverage report in HTML
	go tool cover -html=c.out

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: local_build
local_build: ## Build locally
	go build -o bin/ ${MAIN_PATH}

.PHONY: mocks
mocks: ## Generate mocks
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@mockgen -source=internal/domain/auth.go -destination=$(MOCKS_DESTINATION)/domain/auth.go
	@mockgen -source=internal/domain/posts.go -destination=$(MOCKS_DESTINATION)/domain/posts.go
	@mockgen -source=internal/domain/users.go -destination=$(MOCKS_DESTINATION)/domain/users.go
	@mockgen -source=internal/domain/subscribers.go -destination=$(MOCKS_DESTINATION)/domain/subscribers.go
	@mockgen -source=internal/domain/subscriptions.go -destination=$(MOCKS_DESTINATION)/domain/subscriptions.go
	@mockgen -source=internal/domain/repository.go -destination=$(MOCKS_DESTINATION)/domain/repository.go

.PHONY: clean
clean: ## Remove temporary files
	rm -f main
	go clean

.PHONY: dev
dev: ## Start containers
	# Clearing all stopped containers
	docker container prune -f
    # UP backend docker compose
	docker-compose -f ../deployments/dev/docker-compose.yaml up -d

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := local_build