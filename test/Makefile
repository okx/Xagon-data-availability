STOP := docker compose down

.PHONY: run
run: stop ## Runs a full node for e2e
	docker compose up -d x1-state-db
	docker compose up -d x1-pool-db
	docker compose up -d x1-data-availability-db
	docker compose up -d x1-mock-l1-network
	sleep 1
	docker compose up -d x1-data-availability
	docker compose up -d x1-prover
	docker compose up -d x1-node

.PHONY: stop
stop: ## Stop a full data node
	make stop-dacs
	$(STOP)

.PHONY: stop-dacs
stop-dacs:
	@./stop-dacs > /dev/null 2>&1

.PHONY: test-e2e
test-e2e: run ## Runs the E2E tests
	trap '$(STOP)' EXIT; MallocNanoZone=0 go test -count=1 -race -v -p 1 -timeout 600s ./e2e/...

.PHONY: test-unit
test-unit: ## Runs the unit tests, skips e2e
	go test -coverprofile coverage.out `go list ../... | grep -v /test`

## Help display.
## Pulls comments from beside commands and prints a nicely formatted
## display with the commands and their usage information.
.DEFAULT_GOAL := help

.PHONY: help
help: ## Prints this help
		@grep -h -E '^[a-zA-Z0-9_-]*:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

### node commands called by testing operations ###

.PHONY: run-node
run-node: ## Runs the node
	docker compose up -d x1-node

.PHONY: stop-node
stop-node: ## Stops the node
	docker compose stop x1-node && docker compose rm -f x1-node

.PHONY: run-network
run-network: ## Runs the x1-mock-l1-network network
	docker compose up -d x1-mock-l1-network

.PHONY: stop-network
stop-network: ## Stops the x1-mock-l1-network network
	docker compose stop x1-mock-l1-network && docker compose rm -f x1-mock-l1-network

.PHONY: run-local
run-local: run-network ## Run network and db for local instance
	docker compose up -d x1-data-availability-db

.PHONY: stop-local
stop-local: stop-network ## Run network and db for local instance
	docker compose stop x1-data-availability-db && docker compose rm -f x1-data-availability-db


