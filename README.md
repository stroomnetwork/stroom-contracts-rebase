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

**wstrBTC** is an ERC-20 wrapper for the rebase token **strBTC**, which locks user shares. This ensures compatibility with DeFi protocols that do not support rebase tokens.

### Key Functions

#### Wrap

- Converts `strBTC` to `wstrBTC`.
- user's `wstrBTC` balance is stable, but conversion rate to `strBTC` is changing

#### Unwrap

- Converts `wstrBTC` back to `strBTC`.
- Allows users to retrieve tokens with updated balances after a reward accrual rebases.

---

## Implementation Details

### strBTC

- **`totalSupply`** — The total amount of BTC staked in the pool.
- **`totalShares`** — The total number of shares in the pool.
- Balances are calculated as:

    ```solidity
    balanceOf(account) = (shares[account] * totalSupply) / totalShares
    ```

### wstrBTC

- Locks user shares to provide fixed balances.
- Compatible with DeFi protocols that do not support rebase tokens.
- Calculates shares and balances using the following formulas:

    ```solidity
    wstrBTC = (strBTCAmount * totalShares) / totalSupply
    strBTC = (sharesAmount * totalSupply) / totalShares
    ```

## Reward Minting Process for the Rebase Token

1. **Validator Signature**  
   A validator signs a message to mint new `strBTC` with the specified delta amount.

2. **Increase Total Supply**  
   The `strBTC.mintRewards()` function increases the `totalSupply` of `strBTC` by the delta amount.

3. **Rewards Distribution**  
   The newly minted `strBTC` is added to the pool, automatically distributed to token holders through the rebase mechanism.

4. **Shares Remain Constant**  
   The total shares (`totalShares`) remain constant, ensuring that ownership percentages do not change.

5. **Rebase Mechanism**  
   The rebase mechanism ensures that all token holders' balances increase proportionally, reflecting the pool growth.


### Build & Export

```shell
make go-gen
```

The above goal produces `abistroom/strbtc.go` with Go bindings for the compiled contract ABI.

Changing the contract should cause releasing a new version by publishing a new tag.
Tag should be a valid [Semantic Versioning](https://semver.org/) number. 

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
