
GOEXE ?= go

help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
.PHONY: help

.PHONY: openapi
openapi: ## generate openapi's go code for type, server and client
	@./scripts/openapi-http.sh blog src/app/internal/blog/http blog


.PHONY: tests
tests: ## build and run pomodoro example
	$(GOEXE) test ./internal/product/domain