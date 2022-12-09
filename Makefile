MAIN_PATH = ./cmd/api/main.go
MOCKS_DESTINATION = internal/mocks
INTERNAL_PATH = internal
ACTIVE_PACKAGES = $(shell go list ./... | grep -Ev "mocks|protobuf" | tr '\n' ',')
PROTO_FILES = $(shell find . -iname '*.proto')
GEN_PROTO_FILES = $(shell find . -iname "*.pb.go")

.PHONY: test
test: ## Run all the tests
	go test -coverpkg=$(ACTIVE_PACKAGES) -coverprofile=c.out ./...

.PHONY: cover_out
cover_out: test ## Run all the tests and opens the coverage report
	cat c.out | grep -v "cmd" > tmp.out
	go tool cover -func=tmp.out

.PHONY: cover_html
cover_html: test ## Run all the tests and opens the coverage report in HTML
	cat c.out | grep -v "cmd" > tmp.out
	go tool cover -html=tmp.out

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: local_build
local_build: ## Build locally
	go build -o bin/ ${MAIN_PATH}

.PHONY: docs
docs: ## Make swagger docs
	swag fmt
	swag init --parseDependency --parseInternal -g cmd/api/main.go

.PHONY: proto
proto: ## Make protobuf files
	protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO_FILES)

.PHONY: mocks
mocks: ## Generate mocks
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@mockgen -source=internal/domain/auth.go -destination=$(MOCKS_DESTINATION)/domain/auth.go
	@mockgen -source=internal/domain/posts.go -destination=$(MOCKS_DESTINATION)/domain/posts.go
	@mockgen -source=internal/domain/users.go -destination=$(MOCKS_DESTINATION)/domain/users.go
	@mockgen -source=internal/domain/subscribers.go -destination=$(MOCKS_DESTINATION)/domain/subscribers.go
	@mockgen -source=internal/domain/subscriptions.go -destination=$(MOCKS_DESTINATION)/domain/subscriptions.go
	@mockgen -source=internal/domain/images.go -destination=$(MOCKS_DESTINATION)/domain/images.go
	@mockgen -source=internal/domain/donates.go -destination=$(MOCKS_DESTINATION)/domain/donates.go
	@mockgen -source=internal/domain/repository.go -destination=$(MOCKS_DESTINATION)/domain/repository.go
	@mockgen -source=internal/microservices/auth/protobuf/auth_grpc.pb.go -destination=$(MOCKS_DESTINATION)/domain/auth_client.go
	@mockgen -source=internal/microservices/users/protobuf/users_grpc.pb.go -destination=$(MOCKS_DESTINATION)/domain/users_client.go
	@mockgen -source=internal/microservices/donates/protobuf/donates_grpc.pb.go -destination=$(MOCKS_DESTINATION)/domain/donates_client.go
	@mockgen -source=internal/microservices/images/protobuf/images_grpc.pb.go -destination=$(MOCKS_DESTINATION)/domain/images_client.go
	@mockgen -source=internal/microservices/post/protobuf/posts_grpc.pb.go -destination=$(MOCKS_DESTINATION)/domain/post_client.go
	@mockgen -source=internal/microservices/subscribers/protobuf/subscribers_grpc.pb.go -destination=$(MOCKS_DESTINATION)/domain/subscribers_client.go
	@mockgen -source=internal/microservices/subscriptions/protobuf/subscriptions_grpc.pb.go -destination=$(MOCKS_DESTINATION)/domain/subscriptions_client.go
	@mockgen -source=internal/microservices/users/protobuf/users_grpc.pb.go -destination=$(MOCKS_DESTINATION)/domain/users_client.go
	@echo "OK"

.PHONY: lint
lint: ## Make linters
	@golangci-lint run -c configs/.golangci.yaml

.PHONY: clean
clean: ## Remove temporary files
	rm -f main
	go clean

.PHONY: dev
dev: ## Start containers
	# Clearing all stopped containers
	docker container prune -f
    # UP backend docker compose
	docker-compose -f deployments/docker-compose.yaml up -d

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := local_build