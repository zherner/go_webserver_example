.PHONY: build

help: ## Display this help text
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Runs Go tests
	go test -v

build: test ## Builds the Go binary
	go build ./...

run: ## Runs the app in container
	docker-compose up

clean: ## Cleanup the binary.
	@rm -rf go_webserver_example