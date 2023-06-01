.DEFAULT_GOAL := help

.PHONY: help download wasm build

help:
	@# 20s is the width of the first column
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

download: ## download dependencies
	@echo Download go.mod dependencies
	@go mod download

build: ## build from current commit
	@rm -f ./wasm/*.wasm
	@make antlr
	@make wasm
	@echo Build parser
	@goreleaser build --snapshot --clean
	@rm -f ./sql_grammar/go.mod

wasm: ## build wasm binary
	@echo Build wasm binary
	@go generate ./wasm/wasm.go

antlr: ## build Antlr code
	@# need to generate go.mod file for antlr
	@echo Generate antlr code
	@cd ./sql_grammar/ && rm -f go.mod && go mod init sql_grammar && ./generate.sh

git-sync: ## sync submodule
	@git submodule update --remote

git-init: ## init submodule
	@git submodule update --init