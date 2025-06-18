// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

import "blockchain-tools/src/BitcoinUtils.sol";
import "blockchain-tools/src/BitcoinNetworkEncoder.sol";

import "./lib/ValidatorMessageReceiver.sol";
import "./lib/ValidatorRegistry.sol";

contract strBTC is ERC20Upgradeable, ValidatorMessageReceiver, PausableUpgradeable, AccessControlUpgradeable {
    error InsufficientBalance();
    error MinWithdrawTooLow();
    error MintAmountZero();
    error MintAmountTooBig();
    error MintToContractAddress();
    error MintAlreadyProcessed();
    error InvalidTotalSupplyNonce();
    error DeltaIsZero();
    error UpdateAlreadyProcessed();
    error AmountBelowMinWithdraw();
    error InvalidBTCAddress();
    error InvalidTotalSharesOrPooledBTC();
    error CannotMintToZeroAddress();
    error CannotBurnFromZeroAddress();
    error RewardTooFrequent();
    error RewardTooBig();
    error MaxRewardPercentTooHigh();
    error MinTimeBetweenRewardsTooLow();
    error InvalidAddress();

    using BitcoinUtils for BitcoinNetworkEncoder.Network;

    BitcoinNetworkEncoder.Network public network;

    struct MintInvoice {
        bytes32 btcDepositId;
        address recipient;
        uint256 amount;
    }

    uint256 public constant DUST_LIMIT = 546; // sat
    uint256 public constant BTC = 1e8; // sat

    bytes public constant MESSAGE_MINT = "STROOM_MINT_INVOICE";
    bytes public constant MESSAGE_UPDATE_TOTAL_SUPPLY = "STROOM_UPDATE_TOTAL_SUPPLY";
    bytes32 public constant CONVERTER_ROLE = keccak256("CONVERTER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    uint256 public minWithdrawAmount;
    uint256 public redeemCounter;
    uint256 public totalSupplyUpdateNonce;

    uint256 private _totalPooledBTC;
    uint256 private _totalShares;

    uint256 public maxRewardPercent; // in basis points, 1% = 100
    uint256 public minTimeBetweenRewards; // in seconds
    uint256 public lastRewardTimestamp;

    mapping(bytes32 => bool) public btcDepositIds;
    mapping(address => uint256) private shares;

    event MintBtcEvent(address indexed _to, uint256 _value, bytes32 _btcDepositId);
    event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id);
    event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares);
    event ConverterMint(address indexed converter, address indexed recipient, uint256 amount);
    event ConverterBurn(address indexed converter, address indexed from, uint256 amount);

    function initialize(
        BitcoinNetworkEncoder.Network _network,
        ValidatorRegistry _validatorRegistry,
        address _admin,
        address _pauser
    ) public initializer {
        ERC20Upgradeable.__ERC20_init("Stroom Bitcoin", "strBTC");
        PausableUpgradeable.__Pausable_init();
        AccessControlUpgradeable.__AccessControl_init();
        ValidatorMessageReceiver.initialize(_validatorRegistry);

        minWithdrawAmount = 7_000; // 0.00007 BTC

        network = _network;

        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
        _grantRole(PAUSER_ROLE, _pauser);

        maxRewardPercent = 100; // 1% of the total supply
        minTimeBetweenRewards = 1 days;
        lastRewardTimestamp = 0;
    }

    // ========= Override Functions ======

    function decimals() public pure override returns (uint8) {
        return 8;
    }

    function totalSupply() public view override returns (uint256) {
        return _totalPooledBTC;
    }

    function totalShares() public view returns (uint256) {
        return _totalShares;
    }

    function balanceOf(address account) public view override returns (uint256) {
        if (_totalShares == 0) {
            return 0;
        }
        return (shares[account] * _totalPooledBTC) / _totalShares;
    }

    function getShares(address account) external view returns (uint256) {
        return shares[account];
    }

    /**
     * @dev Internal function for updating the balance of shares and totalShares during mint, burn, transfer operations.
     * @param from The sender's address (or address(0) for mint).
     * @param to The recipient's address (or address(0) for burn).
     * @param value The number of tokens that are transferred/exchanged/burned.
     */
    function _update(address from, address to, uint256 value) internal override {
        // mint
        if (from == address(0)) {
            uint256 sharesToMint = (_totalShares == 0 || _totalPooledBTC == 0) ? value : getSharesByPooledBTC(value);

            _totalPooledBTC += value;
            _totalShares += sharesToMint;

            unchecked {
                shares[to] += sharesToMint;
            }
        } else {
            uint256 fromBalance = balanceOf(from);
            if (fromBalance < value) revert InsufficientBalance();

            uint256 sharesToTransfer = getSharesByPooledBTC(value);

            // burn
            if (to == address(0)) {
                _totalPooledBTC -= value;
                _totalShares -= sharesToTransfer;

                unchecked {
                    shares[from] -= sharesToTransfer;
                }
            } else {
                // transfer
                unchecked {
                    shares[from] -= sharesToTransfer;
                    shares[to] += sharesToTransfer;
                }
            }
        }
        emit Transfer(from, to, value);
    }

    // ========= Minting Signature ======

    /**
     * @dev Calculates the data of an invoice.
     * @return The data of the invoice.
     */
    function getMintInvoiceHash(MintInvoice calldata invoice) public view returns (bytes32) {
        return validatorRegistry.getMessageHash(MESSAGE_MINT, encodeInvoice(invoice));
    }

    /**
     * @dev Calculates the data of the total supply update.
     * @return The data of the total supply update.
     */
    function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) public view returns (bytes32) {
        return validatorRegistry.getMessageHash(MESSAGE_UPDATE_TOTAL_SUPPLY, encodeTotalSupplyUpdate(nonce, delta));
    }

    /**
     * @dev Calculates the inner data of an invoice.
     * @return The bytes data of the invoice.
     */
    function encodeInvoice(MintInvoice calldata invoice) public view returns (bytes memory) {
        return abi.encodePacked(invoice.recipient, invoice.amount, invoice.btcDepositId, address(this));
    }

    /**
     * @dev Calculates the inner data of total supply update.
     * @return The bytes data of the total supply update.
     */
    function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) public view returns (bytes memory) {
        return abi.encodePacked(nonce, delta, address(this));
    }

    // ========= Access Control ========

    /**
     * @dev Function to stop the contract (Pausable pattern).
     */
    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @dev Function to resume the contract (Pausable pattern).
     */
    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @dev Sets the minimum withdrawal amount.
     * @param _minWithdrawAmount The new minimum withdrawal amount.
     * @notice Only the contract owner can call this function.
     */
    function setMinWithdrawAmount(uint256 _minWithdrawAmount) public onlyRole(DEFAULT_ADMIN_ROLE) {
        if (_minWithdrawAmount < DUST_LIMIT) revert MinWithdrawTooLow();
        minWithdrawAmount = _minWithdrawAmount;
    }

    /**
     * @notice Sets the maximum reward percentage
     * @param _maxRewardPercent The new maximum reward percentage
     * @dev Only the contract owner can call this function
     */
    function setMaxRewardPercent(uint256 _maxRewardPercent) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (_maxRewardPercent > 100) revert MaxRewardPercentTooHigh(); // Maximum 1%
        maxRewardPercent = _maxRewardPercent;
    }

    /**
     * @notice Sets the minimum time between rewards
     * @param _minTimeBetweenRewards The new minimum time between rewards
     * @dev Only the contract owner can call this function
     */
    function setMinTimeBetweenRewards(uint256 _minTimeBetweenRewards) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (_minTimeBetweenRewards < 12 hours) revert MinTimeBetweenRewardsTooLow(); // Minimum 12 hour
        minTimeBetweenRewards = _minTimeBetweenRewards;
    }

    /**
     * @notice Adds converter to the whitelist
     * @param converter Converter address
     */
    function addConverter(address converter) external onlyRole(DEFAULT_ADMIN_ROLE) {
        grantRole(CONVERTER_ROLE, converter);
    }

    /**
     * @notice Removes converter from the whitelist
     * @param converter Converter address
     */
    function removeConverter(address converter) external onlyRole(DEFAULT_ADMIN_ROLE) {
        revokeRole(CONVERTER_ROLE, converter);
    }

    /**
     * @notice Adds pauser role
     * @param pauser Pauser address
     */
    function addPauser(address pauser) external onlyRole(DEFAULT_ADMIN_ROLE) {
        grantRole(PAUSER_ROLE, pauser);
    }

    /**
     * @notice Removes pauser role
     * @param pauser Pauser address
     */
    function removePauser(address pauser) external onlyRole(DEFAULT_ADMIN_ROLE) {
        revokeRole(PAUSER_ROLE, pauser);
    }

    /**
     * @notice Change the validator registry contract address
     * @param _validatorRegistry The new validator registry contract address
     */
    function setValidatorRegistry(address _validatorRegistry) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (_validatorRegistry == address(0)) revert InvalidAddress();
        validatorRegistry = ValidatorRegistry(_validatorRegistry);
    }

    // ========= Converter-only ========

    /**
     * @notice Safe mint strBTC for wBTC converter
     * @param recipient strBTC recipient
     * @param amount strBTC amount to mint
     * @dev Can only be called by authorized converters
     */
    function converterMint(address recipient, uint256 amount) external whenNotPaused onlyRole(CONVERTER_ROLE) {
        if (recipient == address(0)) revert CannotMintToZeroAddress();

        _mint(recipient, amount);

        emit ConverterMint(msg.sender, recipient, amount);
    }

    /**
     * @notice Safe burning strBTC for wBTC converter
     * @param from strBTC address to burn
     * @param amount strBTC amount to burn
     * @dev Can only be called by authorized converters
     */
    function converterBurn(address from, uint256 amount) external whenNotPaused onlyRole(CONVERTER_ROLE) {
        if (from == address(0)) revert CannotBurnFromZeroAddress();

        _burn(from, amount);

        emit ConverterBurn(msg.sender, from, amount);
    }

    // ========= Validators-only ========

    /**
     * @dev Mints `_amount` of BTC to `_recipient` with a signature.
     * Anyone can use if they have valid signature from the owner.
     */
    function mint(MintInvoice calldata invoice, bytes calldata signature)
        public
//        whenNotPaused
//        onlyValidator(MESSAGE_MINT, encodeInvoice(invoice), signature)
    {
        revert InvalidTotalSupplyNonce();

        _mint(invoice.amount, invoice.recipient, invoice.btcDepositId);
    }

    /**
     * @notice Adds rewards to the total pooled BTC based on a signed validator message and updates the supply state.
     * @dev This function allows secure minting of new rewards by requiring a valid validator signature.
     * @param nonce The current nonce for the total supply update, ensuring the correct order of updates.
     * @param delta The amount of BTC to be added as rewards to the total pooled BTC.
     * @param signature The signed message from the validator set, validating the minting operation.
     */
    function mintRewards(uint256 nonce, uint256 delta, bytes calldata signature)
        external
        whenNotPaused
        onlyValidator(MESSAGE_UPDATE_TOTAL_SUPPLY, encodeTotalSupplyUpdate(nonce, delta), signature)
    {
        if (nonce != totalSupplyUpdateNonce) revert InvalidTotalSupplyNonce();
        totalSupplyUpdateNonce += 1;

        if (delta == 0) revert DeltaIsZero();

        if (block.timestamp < lastRewardTimestamp + minTimeBetweenRewards) {
            revert RewardTooFrequent();
        }

        uint256 maxDelta = (_totalPooledBTC * maxRewardPercent) / 10000;
        if (delta > maxDelta) {
            revert RewardTooBig();
        }

        bytes32 rewardId = getTotalSupplyUpdateHash(nonce, delta);
        if (btcDepositIds[rewardId]) revert UpdateAlreadyProcessed();
        btcDepositIds[rewardId] = true;

        _totalPooledBTC += delta;
        lastRewardTimestamp = block.timestamp;

        emit TotalSupplyUpdatedEvent(nonce, _totalPooledBTC, _totalShares);
    }

    // ========= Internal ========

    /**
     * @dev Mint new tokens.
     * @param _amount The amount of tokens to mint.
     * @param _recipient The address that will receive the minted tokens.
     * @param _btcDepositId The id of the BTC deposit = keccak256(txHash, vout)
     */
    function _mint(uint256 _amount, address _recipient, bytes32 _btcDepositId) internal {
        if (_amount == 0) revert MintAmountZero();
        if (_amount >= 21_000_000 * BTC) revert MintAmountTooBig();
        if (_recipient == address(this)) revert MintToContractAddress();
        if (btcDepositIds[_btcDepositId]) revert MintAlreadyProcessed();

        btcDepositIds[_btcDepositId] = true;

        _mint(_recipient, _amount);

        emit MintBtcEvent(_recipient, _amount, _btcDepositId);
    }

    // ========= Public ========

    /**
     * @notice Redeems strBTC for its underlying Bitcoin by burning the specified amount of tokens.
     * @dev This function allows users to convert their strBTC holdings back into Bitcoin.
     * @param _amount The amount of strBTC to redeem for Bitcoin.
     * @param BTCAddress The Bitcoin address to receive the redeemed BTC.
     */
    function redeem(uint256 _amount, string calldata BTCAddress) public whenNotPaused {
        if (_amount < minWithdrawAmount) revert AmountBelowMinWithdraw();
        if (!network.validateBitcoinAddress(BTCAddress)) revert InvalidBTCAddress();

        _burn(msg.sender, _amount);

        redeemCounter += 1;
        emit RedeemBtcEvent(msg.sender, BTCAddress, _amount, redeemCounter);
    }

    /**
     * @notice Calculates the number of shares corresponding to a given amount of staked BTC (strBTC).
     * @param btcAmount The amount of strBTC to convert to shares.
     * @return The number of shares equivalent to the given strBTC amount.
     */
    function getSharesByPooledBTC(uint256 btcAmount) public view returns (uint256) {
        if (_totalShares == 0 || _totalPooledBTC == 0) revert InvalidTotalSharesOrPooledBTC();
        return (btcAmount * _totalShares) / _totalPooledBTC;
    }

    /**
     * @notice Calculates the amount of strBTC corresponding to a given number of shares.
     * @param sharesAmount The number of shares to convert to strBTC.
     * @return The amount of strBTC equivalent to the given shares.
     */
    function getPooledBTCByShares(uint256 sharesAmount) public view returns (uint256) {
        if (_totalShares == 0 || _totalPooledBTC == 0) revert InvalidTotalSharesOrPooledBTC();
        return (sharesAmount * _totalPooledBTC) / _totalShares;
    }
}
