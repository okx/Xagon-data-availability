STOP := docker compose down

.PHONY: run
run: ## Runs a full node for e2e
	docker compose up -d x1-state-db
	docker compose up -d x1-pool-db
	docker compose up -d x1-event-db
	docker compose up -d x1-data-availability-db
	docker compose up -d x1-mock-l1-network
	sleep 1
	docker compose up -d x1-prover
	docker compose up -d x1-approve
	sleep 3
	docker compose up -d x1-sync
	sleep 2
	docker compose up -d x1-eth-tx-manager
	docker compose up -d x1-sequencer
	docker compose up -d x1-sequence-sender
	docker compose up -d x1-l2gaspricer
	docker compose up -d x1-aggregator
	docker compose up -d x1-json-rpc
	sleep 2
	docker compose up -d x1-data-availability

.PHONY: stop
stop: ## Stop a full data node
	$(STOP)

.PHONY: test-e2e
test-e2e: run ## Runs the E2E tests
	trap '$(STOP)' EXIT; MallocNanoZone=0 go test -count=1 -race -v -p 1 -timeout 600s ./e2e/...

.PHONY: test-unit
test-unit: ## Runs the unit tests, skips e2e
	go test `go list ../... | grep -v /e2e`

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
	docker compose up -d x1-eth-tx-manager
	docker compose up -d x1-sync
	sleep 2
	docker compose up -d x1-sequencer
	docker compose up -d x1-sequence-sender
	docker compose up -d x1-l2gaspricer
	docker compose up -d x1-aggregator
	docker compose up -d x1-json-rpc

.PHONY: stop-node
stop-node: ## Stops the node
	docker compose stop x1-sequencer && docker compose rm -f x1-sequencer
	docker compose stop x1-sequence-sender && docker compose rm -f x1-sequence-sender
	docker compose stop x1-json-rpc && docker compose rm -f x1-json-rpc
	docker compose stop x1-l2gaspricer && docker compose rm -f x1-l2gaspricer
	docker compose stop x1-aggregator && docker compose rm -f x1-aggregator
	docker compose stop x1-sync && docker compose rm -f x1-sync
	docker compose stop x1-eth-tx-manager && docker compose rm -f x1-eth-tx-manager

.PHONY: run-network
run-network: ## Runs the l1 network
	docker compose up -d x1-mock-l1-network

.PHONY: stop-network
stop-network: ## Stops the l1 network
	docker compose stop x1-mock-l1-network && docker compose rm -f x1-mock-l1-network

.PHONY: run-approve-matic
run-approve-matic: ## Runs approve in node container
	docker compose up -d x1-approve

.PHONY: stop-approve-matic
stop-approve-matic: ## Stops approve in node container
	docker compose stop x1-approve && docker compose rm -f x1-approve

### manual utility to stop dac nodes that were started by e2e, but not stopped by e2e ###
.PHONY: stop-dac-nodes
stop-dac-nodes:
	for i in 0 1 2 3 4 ; do \
  		(docker kill x1-data-availability-$$i || true) && (docker rm x1-data-availability-$$i || true) ; \
  	done

.PHONY: stop-dac-dbs
stop-dac-dbs:
	for i in 0 1 2 3 4 ; do \
  		(docker kill x1-data-node-db-$$i || true) && (docker rm x1-data-node-db-$$i || true); \
  	done

.PHONY: stop-dacs
stop-dacs: stop-dac-nodes stop-dac-dbs
