# Stroom Smart Contracts

## strBTC

strBTC is a interest-bearing rebasing token wich is pegged at 1-to-1 ratio to BTC by redemption mechanism. Hence, every strBTC token can be permissionlessly converted to native BTC on Bitcoin blockchain. Each user's balance is automatically updated during a rebase to reflect pool growth (e.g., from earned rewards).

The rebase mechanism is implemented through the relationship between the user's shares and the total amount of staked BTC in the pool (`totalPooledBTC`). This ensures automatic distribution of rewards among tokens without requiring explicit balance updates for each user.

### Features

- **Rebase Token:** All token holders automatically accrue rewards through the rebase mechanism.
- **Validator Approved Updates:** Rebase can be triggered by validators via messages signed by validators, ensuring security, transparency and validity of updates.
- **Validator Approved Mints:** Minting strBTC is triggered by user via submitting the validator signed message to the contract
- **ValidatorMessageReceiver:** Base contract validates the BIP340 Schnorr signatures for secp256k1 elliptic curve produced by validators using threshold signing algorithm [FROST](https://www.cryptohopper.com/news/frost-flexible-round-optimized-schnorr-threshold-signatures-4820)

### Key Functions

#### Mint

- Generates new tokens upon receiving confirmed BTC deposits.
- Requires a validator signature for transaction confirmation.

#### Mint Rewards

- Adds rewards to the pool via signed messages.
- Updates `totalSupply` without changing `totalShares`.
- Supports the rebase mechanism, which automatically distributes new rewards among token holders.

#### Redeem

- Converts strBTC tokens into an equivalent amount of BTC by burning strBTC and specifying Bitcoin redemption address in the burning transactions.
- Contract checks the validity of Bitcoin addresses to avoid locking BTC unspendable.

---

## wstrBTC

**wstrBTC** is an ERC-4626 Tokenized Vault wrapper for the rebase token **strBTC**. It provides a fixed-balance representation of strBTC shares, ensuring full compatibility with DeFi protocols that do not support rebase tokens.

### Standards Compliance

- **ERC-4626:** Standardized yield-bearing vault with `deposit()`, `redeem()`, and rate conversion functions
- **ERC-20:** Standard token functionality with 8 decimal places (matching strBTC)
- **ERC-20 Permit (EIP-2612):** Gasless approvals through cryptographic signatures

### Key Functions

#### Standard ERC-4626 Methods

- **`deposit(assets, receiver)`** — Converts `strBTC` to `wstrBTC` shares
- **`redeem(shares, receiver, owner)`** — Converts `wstrBTC` shares back to `strBTC`
- **`convertToShares(assets)`** — Preview how many shares you'll get for assets
- **`convertToAssets(shares)`** — Preview how many assets you'll get for shares
- **`totalAssets()`** — Total strBTC held by the vault
- **`asset()`** — Returns the underlying strBTC token address

#### ERC-20 Permit Methods

- **`permit(owner, spender, value, deadline, v, r, s)`** — Gasless approval via signature
- Enables one-transaction DeFi interactions without separate approve calls

---

## Implementation Details

### strBTC

- **`totalSupply`** — The total amount of BTC staked in the pool.
- **`totalShares`** — The total number of shares in the pool.
- Balances are calculated as:

    ```solidity
    balanceOf(account) = (shares[account] * totalSupply) / totalShares
    ```

### wstrBTC (ERC-4626 Vault)

- Uses standard ERC-4626 conversion formulas with built-in rounding protection
- **Assets:** strBTC tokens (underlying rebase token)
- **Shares:** wstrBTC tokens (fixed-balance vault tokens)
- Conversion rates automatically reflect strBTC rebase rewards:

    ```solidity
    // Standard ERC-4626 conversions
    shares = convertToShares(assets) = assets * totalSupply() / totalAssets()
    assets = convertToAssets(shares) = shares * totalAssets() / totalSupply()
    ```

## Reward Minting Process for the Rebase Token

1. **Validator Signature**  
   A validator signs a message to mint new `strBTC` with the specified delta amount.

2. **Increase Total Supply**  
   The `strBTC.mintRewards()` function increases the `totalSupply` of `strBTC` by the delta amount.

3. **Rewards Distribution**  
   The newly minted `strBTC` is added to the pool, automatically distributed to token holders through the rebase mechanism.

4. **wstrBTC Rate Update**  
   The `wstrBTC` vault automatically reflects new rewards through increased `convertToAssets()` rate.

5. **Shares Remain Constant**  
   Both strBTC total shares (`totalShares`) and wstrBTC shares remain constant, ensuring ownership percentages don't change.

6. **Rebase Mechanism**  
   The rebase mechanism ensures that all token holders' balances increase proportionally, reflected in both strBTC balances and wstrBTC conversion rates.

### Build & Export

```shell
make go-gen
```

The above goal produces `abistroom/strbtc.go` with Go bindings for the compiled contract ABI.

Changing the contract should cause releasing a new version by publishing a new tag.
Tag should be a valid [Semantic Versioning](https://semver.org/) number. 


## Contract Deployment

### Local Deployment

To deploy contracts to a local network (e.g., Anvil):

```shell
make deploy-local
```

- Uses the local RPC (http://localhost:8545)
- Requires a private key in `.env` (`PRIVATE_KEY_LOCAL`)
- Make sure your local node is running (e.g., with `anvil`) before deploying

To start a local Ethereum node using Anvil run this command in a separate terminal window:

```shell
anvil
```

### Sepolia Deployment

To deploy contracts to the Sepolia testnet:

```shell
make deploy-sepolia
```

- Uses environment variables from `.env`:
  - `SEPOLIA_RPC_URL` — Sepolia RPC endpoint
  - `PRIVATE_KEY` — deployer's private key
  - `ETHERSCAN_API_KEY` — for contract verification
- Contracts are automatically verified on Etherscan

### Additional Scripts

- Deploy Timelock:
  ```shell
  make script-deploy-timelock
  ```
- Initialize seed (joint pubkey, taproot):
  ```shell
  make script-set-seed-sepolia
  ```
  
## Foundry

**Foundry is a blazing fast, portable and modular toolkit for Ethereum application development written in Rust.**

Foundry consists of:

-   **Forge**: Ethereum testing framework (like Truffle, Hardhat and DappTools).
-   **Cast**: Swiss army knife for interacting with EVM smart contracts, sending transactions and getting chain data.
-   **Anvil**: Local Ethereum node, akin to Ganache, Hardhat Network.
-   **Chisel**: Fast, utilitarian, and verbose solidity REPL.

For additional details see [the official Foundry documentation](https://book.getfoundry.sh/).

## Usage

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Format

```shell
$ forge fmt
```

### Gas Snapshots

```shell
$ forge snapshot
```