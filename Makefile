# Inspired by:
#   https://dev.to/brpaz/building-a-basic-ci-cd-pipeline-for-a-golang-application-using-github-actions-icj
#   https://github.com/brpaz/github-actions-demo-go/blob/master/Makefile
PROJECT_NAME := "go-griddlersolver"
PKG := "github.com/yinonavraham/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/...)
GO_FILES := $(shell find . -name '*.go' | grep -v _test.go)
TEST_ARGS := $(TEST_ARGS)

.PHONY: all dep lint vet test test-coverage build clean

all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@go get -u golang.org/x/lint/golint
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test: ## Run tests
	@go test ${TEST_ARGS} ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@mkdir -p out
	@go test -coverprofile out/coverage.out -covermode=atomic ${PKG_LIST}

build: dep ## Build the binary file
	@go build ./...

clean: ## Remove previous build
	@rm -rf out

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
