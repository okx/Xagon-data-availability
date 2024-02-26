STOP := docker compose down

.PHONY: run
run: ## Runs a full node for e2e
	docker compose up -d xlayer-state-db
	docker compose up -d xlayer-pool-db
	docker compose up -d xlayer-event-db
	docker compose up -d xlayer-data-availability-db
	docker compose up -d xlayer-mock-l1-network
	sleep 1
	docker compose up -d xlayer-prover
	docker compose up -d xlayer-approve
	sleep 3
	docker compose up -d xlayer-sync
	sleep 2
	docker compose up -d xlayer-eth-tx-manager
	docker compose up -d xlayer-sequencer
	docker compose up -d xlayer-sequence-sender
	docker compose up -d xlayer-l2gaspricer
	docker compose up -d xlayer-aggregator
	docker compose up -d xlayer-json-rpc
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
	docker compose up -d xlayer-eth-tx-manager
	docker compose up -d xlayer-sync
	sleep 2
	docker compose up -d xlayer-sequencer
	docker compose up -d xlayer-sequence-sender
	docker compose up -d xlayer-l2gaspricer
	docker compose up -d xlayer-aggregator
	docker compose up -d xlayer-json-rpc

.PHONY: stop-node
stop-node: ## Stops the node
	docker compose stop xlayer-sequencer && docker compose rm -f xlayer-sequencer
	docker compose stop xlayer-sequence-sender && docker compose rm -f xlayer-sequence-sender
	docker compose stop xlayer-json-rpc && docker compose rm -f xlayer-json-rpc
	docker compose stop xlayer-l2gaspricer && docker compose rm -f xlayer-l2gaspricer
	docker compose stop xlayer-aggregator && docker compose rm -f xlayer-aggregator
	docker compose stop xlayer-sync && docker compose rm -f xlayer-sync
	docker compose stop xlayer-eth-tx-manager && docker compose rm -f xlayer-eth-tx-manager

.PHONY: run-network
run-network: ## Runs the l1 network
	docker compose up -d xlayer-mock-l1-network

.PHONY: stop-network
stop-network: ## Stops the l1 network
	docker compose stop xlayer-mock-l1-network && docker compose rm -f xlayer-mock-l1-network

.PHONY: run-approve-matic
run-approve-matic: ## Runs approve in node container
	docker compose up -d xlayer-approve

.PHONY: stop-approve-matic
stop-approve-matic: ## Stops approve in node container
	docker compose stop xlayer-approve && docker compose rm -f xlayer-approve

### manual utility to stop dac nodes that were started by e2e, but not stopped by e2e ###
.PHONY: stop-dac-nodes
stop-dac-nodes:
	for i in 0 1 2 3 4 ; do \
  		(docker kill x1-data-availability-$$i || true) && (docker rm x1-data-availability-$$i || true) ; \
  	done

.PHONY: stop-dac-dbs
stop-dac-dbs:
	for i in 0 1 2 3 4 ; do \
  		(docker kill xlayer-data-node-db-$$i || true) && (docker rm xlayer-data-node-db-$$i || true); \
  	done

.PHONY: stop-dacs
stop-dacs: stop-dac-nodes stop-dac-dbs
