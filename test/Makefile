STOP := docker compose down

.PHONY: run
run: ## Runs a full node for e2e
	docker compose up -d xgon-state-db
	docker compose up -d xgon-pool-db
	docker compose up -d xgon-event-db
	docker compose up -d xgon-data-availability-db
	docker compose up -d xgon-mock-l1-network
	sleep 1
	docker compose up -d xgon-prover
	docker compose up -d xgon-approve
	sleep 3
	docker compose up -d xgon-sync
	sleep 2
	docker compose up -d xgon-eth-tx-manager
	docker compose up -d xgon-sequencer
	docker compose up -d xgon-sequence-sender
	docker compose up -d xgon-l2gaspricer
	docker compose up -d xgon-aggregator
	docker compose up -d xgon-json-rpc
	sleep 2
	docker compose up -d xgon-data-availability

.PHONY: stop
stop: ## Stop a full data node
	$(STOP)

.PHONY: test-e2e
test-e2e: run ## Runs the E2E tests
	go test -v ./e2e/...
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
	docker compose up -d xgon-eth-tx-manager
	docker compose up -d xgon-sync
	sleep 2
	docker compose up -d xgon-sequencer
	docker compose up -d xgon-sequence-sender
	docker compose up -d xgon-l2gaspricer
	docker compose up -d xgon-aggregator
	docker compose up -d xgon-json-rpc

.PHONY: stop-node
stop-node: ## Stops the node
	docker compose stop xgon-sequencer && docker compose rm -f xgon-sequencer
	docker compose stop xgon-sequence-sender && docker compose rm -f xgon-sequence-sender
	docker compose stop xgon-json-rpc && docker compose rm -f xgon-json-rpc
	docker compose stop xgon-l2gaspricer && docker compose rm -f xgon-l2gaspricer
	docker compose stop xgon-aggregator && docker compose rm -f xgon-aggregator
	docker compose stop xgon-sync && docker compose rm -f xgon-sync
	docker compose stop xgon-eth-tx-manager && docker compose rm -f xgon-eth-tx-manager

.PHONY: run-network
run-network: ## Runs the l1 network
	docker compose up -d xgon-mock-l1-network

.PHONY: stop-network
stop-network: ## Stops the l1 network
	docker compose stop xgon-mock-l1-network && docker compose rm -f xgon-mock-l1-network

.PHONY: run-approve-matic
run-approve-matic: ## Runs approve in node container
	docker compose up -d xgon-approve

.PHONY: stop-approve-matic
stop-approve-matic: ## Stops approve in node container
	docker compose stop xgon-approve && docker compose rm -f xgon-approve

### manual utility to stop dac nodes that were started by e2e, but not stopped by e2e ###
.PHONY: stop-dac-nodes
stop-dac-nodes:
	for i in 0 1 2 3 4 ; do \
  		(docker kill xgon-data-availability-$$i || true) && (docker rm xgon-data-availability-$$i || true) ; \
  	done

.PHONY: stop-dac-dbs
stop-dac-dbs:
	for i in 0 1 2 3 4 ; do \
  		(docker kill xgon-data-node-db-$$i || true) && (docker rm xgon-data-node-db-$$i || true); \
  	done

.PHONY: stop-dacs
stop-dacs: stop-dac-nodes stop-dac-dbs
