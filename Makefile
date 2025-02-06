BUILD_DIR := bin
TOOLS_DIR := tools

.DEFAULT_GOAL:=help
.PHONY: all clean lint build mock test run docker-up docker-down help

all: clean lint test build run ## Run all tests, then build and run

.PHONY: $(BUILD_DIR)/server
$(BUILD_DIR)/server:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BUILD_DIR)/server ./cmd/server

build: $(BUILD_DIR)/server ## Build the binary

clean: ## Clean up, i.e. remove build artifacts
	rm -rf $(BUILD_DIR)
	rm -rf $(TOOLS_DIR)
	@go mod tidy

run: build ## Run the binary
	$(BUILD_DIR)/server

docker-up: ## Run in a container using docker
	docker-compose up --build --force-recreate

docker-down: ## Stop the running container
	docker-compose down -v --remove-orphans

tools/golangci-lint/golangci-lint:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b tools/golangci-lint latest

tools/mockery:
	go install github.com/vektra/mockery/v2@latest

lint: $(TOOLS_DIR)/golangci-lint/golangci-lint ## Run linters
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run ./...

.PHONY: test ## Run tests
test: ## Run tests
	go test -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

mock: tools/mockery ## Generate mocks using mockery
	mockery --name=Repository --dir=internal/api/device --output=internal/api/device/mock --case=underscore

help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
