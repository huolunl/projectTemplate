mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
GO := go

.PHONY: all
all: install-swagger tidy
include scripts/make-rules/golang.mk
include scripts/make-rules/swagger.mk

.PHONY: tidy
tidy:
	@$(MAKE) go.mod.tidy

## swagger: Generate swagger document.
.PHONY: swagger
swagger:
	@$(MAKE) swagger.run

## serve-swagger: Serve swagger spec and docs.
.PHONY: swagger.serve
serve-swagger:
	@$(MAKE) swagger.serve

## lint: Check syntax and styling of go sources.
.PHONY: lint
lint:
	@$(MAKE) go.lint

## format: Gofmt (reformat) package sources (exclude vendor dir if existed).
.PHONY: format
format:
	@echo "===========> Formating codes"
	@gofmt ./

## format: Gofmt (reformat) package sources (exclude vendor dir if existed).
.PHONY: run
run:
	@$(MAKE) go.run


.PHONY: cover
cover:
	@go test -race -cover  -coverprofile=./coverage.out -timeout=10m -short -v ./...
	@go tool cover -func ./coverage.out

## test: Run unit test.
.PHONY: test
test:
	@$(MAKE) go.test

## test: Run unit test.
.PHONY: create.test
create-test:
	@$(MAKE) go.create.test

.PHONY: install.swagger
install-swagger:
	@$(GO) get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: install.golangci-lint
install-golangci-lint:
	@$(GO) get -u github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1

.PHONY: install.gotests
install-gotests:
	@$(GO) get -u github.com/cweill/gotests/...

.PHONY: help
help:
	@echo "make run 'Start the application'"
	@echo "make lint 'Static detection'"
	@echo "make swagger 'Generate the swagger document"
	@echo "make serve-swagger 'Start the swagger service"
	@echo "make format 'Format the code'"
	@echo "make create-test TEST_DIR=<cmd/apiserver> 'Create the unit tests'"
	@echo "make test TEST_DIR=<cmd/apiserver> 'Run the unit test'"
	@echo "make cover 'Test unit test cover rate'"
	@echo "make install-swagger 'Install the go-swagger'"
	@echo "make install-golangci-lint 'Install the golangci-lint'"
	@echo "make install-gotests 'Install the gotests'"

