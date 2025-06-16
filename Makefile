.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

# Passing args
RUN_ARGS := $(wordlist 2, $(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-38s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the application
	go build -o ./bin/chip8 ./cmd/main.go

.PHONY: test
test: ## Run the tests
	go test ./...