.DEFAULT_GOAL := help

COVER_FILE ?= coverage.out

get-install-dependences:
	go get firebase.google.com/go/v4@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files


install-migrate:
	curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
	echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
	apt-get update
	apt-get install -y migrate


run:
	go run cmd/main.go 

swag:
	swag fmt
	swag init -g ./cmd/main.go -o cmd/docs

cover:
	go tool cover

tools: ## Install all needed tools, e.g.
	@go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2



migrate-init:
	migrate create -dir ./migrations/test -ext sql test


.PHONY: test
test: ## Run all tests.
	@go test -race -count=1 -coverprofile=$(COVER_FILE) ./...
	@go tool cover -func=$(COVER_FILE) | grep ^total | tr -s '\t'

.PHONY: lint
lint: tools ## Check the project with lint.
	@golangci-lint run --fix ./...

.PHONY: build
build: ## Build a command to quickly check compiles.
	@go build ./...

.PHONY: check
check: lint build test ## Runs all necessary code checks.

kill-dockers:
	sudo docker stop $(sudo docker ps -aq)

help: ## Show help for each of the Makefile targets.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


