GO := go

.PHONY: go.mod.tidy
go.mod.tidy:
	@$(GO) mod tidy

.PHONY: go.lint
go.lint:
	@echo "===========> Run golangci to lint source codes"
	@golangci-lint run $(ROOT_DIR)/...

.PHONY: go.run
go.run:
	@$(GO) run $(ROOT_DIR)/cmd/apiserver

.PHONY: go.test
go.test:
	@echo "===========> Run unit test"
	@$(GO) test -v $(ROOT_DIR)/$(TEST_DIR)/

.PHONY: go.create.test
go.create.test:
	@echo $(TEST_DIR)
	@gotests --all -w $(TEST_DIR)
