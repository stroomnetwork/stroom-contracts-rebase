# include .env file and export its env vars.
-include .env
export

GO_ABI_DIR := stroomabi

ifeq ($(shell uname), Darwin) # macOS
	SED_CMD = sed -i ''
else # Linux and others
	SED_CMD = sed -i
endif

.PHONY: all
all: help

## Solidity
.PHONY: build
build: ## Build Solidity contracts using forge
	forge build

.PHONY: test
test: ## Test Solidity contracts using forge
	forge test

.PHONY: coverage
coverage: build ## Run coverage report
	forge coverage --report lcov
	genhtml ./lcov.info -o coverage --ignore-errors category --branch-coverage
	open ./coverage/index.html

.PHONY: go-init
go-init: ## Initialize Go module
	@if [ ! -f go.mod ]; then \
		echo "Initializing Go module..."; \
		go mod init stroom-contracts-rebase; \
	fi
	go mod tidy

.PHONY: export-abi
export-abi: build ## Export contracts ABI and bytecode for strBTC, wstrBTC
	cat ./out/strBTC.sol/strBTC.json | jq -r '.abi' > ./out/strBTC.abi.json
	cat ./out/strBTC.sol/strBTC.json | jq -r '.bytecode.object' | sed 's/0x//' > ./out/strBTC.bin
	cat ./out/wstrBTC.sol/wstrBTC.json | jq -r '.abi' > ./out/wstrBTC.abi.json
	cat ./out/wstrBTC.sol/wstrBTC.json | jq -r '.bytecode.object' | sed 's/0x//' > ./out/wstrBTC.bin
	cat ./out/UserActivator.sol/UserActivator.json | jq -r '.abi' > ./out/UserActivator.abi.json
	cat ./out/UserActivator.sol/UserActivator.json | jq -r '.bytecode.object' | sed 's/0x//' > ./out/UserActivator.bin
	cat ./out/ValidatorRegistry.sol/ValidatorRegistry.json | jq -r '.abi' > ./out/ValidatorRegistry.abi.json
	cat ./out/ValidatorRegistry.sol/ValidatorRegistry.json | jq -r '.bytecode.object' | sed 's/0x//' > ./out/ValidatorRegistry.bin
	cat ./out/TransparentUpgradeableProxy.sol/TransparentUpgradeableProxy.json | jq -r '.abi' > ./out/TransparentUpgradeableProxy.abi.json
	cat ./out/TransparentUpgradeableProxy.sol/TransparentUpgradeableProxy.json | jq -r '.bytecode.object' | sed 's/0x//' > ./out/TransparentUpgradeableProxy.bin

## Go Bindings
.PHONY: go-gen
go-gen: build export-abi ## Generate Go bindings
	abigen --abi ./out/strBTC.abi.json --bin ./out/strBTC.bin --pkg ${GO_ABI_DIR} --type StrBtc --out ./${GO_ABI_DIR}/strbtc.go
	abigen --abi ./out/wstrBTC.abi.json --bin ./out/wstrBTC.bin --pkg ${GO_ABI_DIR} --type WStrBtc --out ./${GO_ABI_DIR}/wstrbtc.go
	abigen --abi ./out/UserActivator.abi.json --bin ./out/UserActivator.bin --pkg ${GO_ABI_DIR} --type UserActivator --out ./${GO_ABI_DIR}/user_activator.go
	abigen --abi ./out/ValidatorRegistry.abi.json --bin ./out/ValidatorRegistry.bin --pkg ${GO_ABI_DIR} --type ValidatorRegistry --out ./${GO_ABI_DIR}/validator_registry.go
	abigen --abi ./out/TransparentUpgradeableProxy.abi.json --bin ./out/TransparentUpgradeableProxy.bin --pkg ${GO_ABI_DIR} --type TransparentUpgradeableProxy --out ./${GO_ABI_DIR}/proxy.go

.PHONY: go-gen-test
go-gen-test: build export-abi ## Verify Go bindings
	mkdir -p ./build/tmp/${GO_ABI_DIR}
	abigen --abi ./out/strBTC.abi.json --bin ./out/strBTC.bin --pkg ${GO_ABI_DIR} --type StrBtc --out ./build/tmp/${GO_ABI_DIR}/strbtc.go
	abigen --abi ./out/wstrBTC.abi.json --bin ./out/wstrBTC.bin --pkg ${GO_ABI_DIR} --type WStrBtc --out ./build/tmp/${GO_ABI_DIR}/wstrbtc.go
	abigen --abi ./out/UserActivator.abi.json --bin ./out/UserActivator.bin --pkg ${GO_ABI_DIR} --type UserActivator --out ./build/tmp/${GO_ABI_DIR}/user_activator.go
	abigen --abi ./out/ValidatorRegistry.abi.json --bin ./out/ValidatorRegistry.bin --pkg ${GO_ABI_DIR} --type ValidatorRegistry --out ./build/tmp/${GO_ABI_DIR}/validator_registry.go
	abigen --abi ./out/TransparentUpgradeableProxy.abi.json --bin ./out/TransparentUpgradeableProxy.bin --pkg ${GO_ABI_DIR} --type TransparentUpgradeableProxy --out ./build/tmp/${GO_ABI_DIR}/proxy.go
	go test -v ./${GO_ABI_DIR}

## Script & Deploy. Uses environment variables that can be specified by '.env' file
.PHONY: deploy-local
deploy-local: export BITCOIN_NETWORK=2
deploy-local: ## Deploy contracts locally
	forge script script/Deploy.s.sol:DeployScript --rpc-url http://localhost:8545 --broadcast --private-key ${PRIVATE_KEY_LOCAL}

.PHONY: deploy-sepolia
deploy-sepolia: export BITCOIN_NETWORK=1 
deploy-sepolia: ## Deploy contracts to Sepolia
	forge script script/Deploy.s.sol:DeployScript --rpc-url ${SEPOLIA_RPC_URL} --private-key ${PRIVATE_KEY} --broadcast --verify --etherscan-api-key ${ETHERSCAN_API_KEY}

.PHONY: deploy-mainnet
deploy-mainnet: export BITCOIN_NETWORK=0
deploy-mainnet: ## Deploy contracts to Mainnet
	forge script script/Deploy.s.sol:DeployScript --rpc-url ${MAINNET_RPC_URL} --private-key ${PRIVATE_KEY} --broadcast --verify --etherscan-api-key ${ETHERSCAN_API_KEY}

.PHONY: verify-proxy-admin
verify-proxy-admin: ## Verify proxy contract address
	forge verify-contract --compiler-version 0.8.27 --constructor-args $(cast abi-encode "constructor(address)" ${OWNER_ADDRESS}) --etherscan-api-key ${ETHERSCAN_API_KEY} ${PROXY_ADMIN_ADDRESS} ProxyAdmin --rpc-url ${SEPOLIA_RPC_URL}

.PHONY: script-set-seed-sepolia
script-set-seed-sepolia: export BITCOIN_NETWORK=1 
script-set-seed-sepolia: ## Sets joint public key and taproot addresses for the contracts
	forge script script/SetSeed.s.sol:SetSeedScript --rpc-url ${SEPOLIA_RPC_URL} --private-key ${PRIVATE_KEY} --broadcast --verify --etherscan-api-key ${ETHERSCAN_API_KEY}

.PHONY: script-deploy-timelock
script-deploy-timelock: export BITCOIN_NETWORK=1 
script-deploy-timelock: ## Deploy timelock contract
	forge script script/DeployTimelock.s.sol:DeployTimelockScript --rpc-url ${SEPOLIA_RPC_URL} --private-key ${PRIVATE_KEY} --broadcast --verify --etherscan-api-key ${ETHERSCAN_API_KEY}

.PHONY: script-clean
script-clean: ## Clean up forge artifacts
	forge clean

.PHONY: clean
clean: script-clean ## Clean all build artifacts and caches
	rm -rf ./out
	rm -rf ./build
	rm -rf ./cache

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

## Help:
.PHONY: help
help: ## Show this help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-30s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
