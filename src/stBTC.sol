// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

import "blockchain-tools/src/BitcoinUtils.sol";
import "blockchain-tools/src/BitcoinNetworkEncoder.sol";
 
import "./lib/ValidatorMessageReceiver.sol";
import "./lib/ValidatorRegistry.sol";

import "./IStBTC.sol";

contract stBTC is 
    ERC20Upgradeable, 
    ValidatorMessageReceiver, 
    BitcoinUtils, 
    PausableUpgradeable, 
    IStBTC 
{
    BitcoinNetworkEncoder.Network public network;

    uint256 public constant DUST_LIMIT = 546; // sat
    uint256 public constant BTC = 1e8; // sat

    bytes public constant MESSAGE_MINT = "STROOM_MINT_INVOICE";
    bytes public constant MESSAGE_UPDATE_TOTAL_SUPPLY = "STROOM_UPDATE_TOTAL_SUPPLY";

    uint256 public minWithdrawAmount; // 0.00002 BTC
    uint256 public redeemCounter;
    uint256 public totalSupplyUpdateNonce;

    address public minter;

    uint256 private totalPooledBTC;
    uint256 private totalShares;    

    mapping(bytes32 => bool) public btcDepositIds;
    mapping(address => uint256) private shares; 

    function initialize(
        BitcoinNetworkEncoder.Network _network,
        ValidatorRegistry _validatorRegistry
    ) public initializer {
        ERC20Upgradeable.__ERC20_init("Stroom Bitcoin", "stBTC");
        PausableUpgradeable.__Pausable_init();
        ValidatorMessageReceiver.initialize(_validatorRegistry);

        minWithdrawAmount = 7_000; // 0.00007 BTC

        redeemCounter = 0;

        network = _network;
    }

    function decimals() public pure override returns (uint8) {
        return 8;
    }

    function totalSupply() public view override returns (uint256) {
        return totalPooledBTC;
    }

    function balanceOf(address account) public view override returns (uint256) {
        if (totalShares == 0) {
            return 0;
        }
        return (shares[account] * totalPooledBTC) / totalShares;
    }

    function getShares(address account) external view returns (uint256) {
        return shares[account];
    }

    // ========= Minting Signature ======

    /**
     * @dev Calculates the data of an invoice.
     * @return The data of the invoice.
     */
    function getMintInvoiceHash(
        MintInvoice calldata invoice
    ) public view returns (bytes32) {
        return validatorRegistry.getMessageHash(
            MESSAGE_MINT, 
            encodeInvoice(invoice)
        );
    }

    /**
     * @dev Calculates the data of the total supply update.
     * @return The data of the total supply update.
     */
    function getTotalSupplyUpdateHash(
        uint256 nonce,
        uint256 delta
    ) public view returns (bytes32) {
        return validatorRegistry.getMessageHash(
            MESSAGE_UPDATE_TOTAL_SUPPLY,
            encodeTotalSupplyUpdate(nonce, delta)
        );
    }

    /**
    * @dev Calculates the inner data of an invoice.
    * @return The bytes data of the invoice.
    */
    function encodeInvoice(
        MintInvoice calldata invoice
    ) public view returns (bytes memory) {
        return
            abi.encodePacked(
                invoice.recipient,
                invoice.amount,
                invoice.btcDepositId,
                address(this)
            );
    }

    /**
    * @dev Calculates the inner data of total supply update.
    * @return The bytes data of the total supply update.
    */
    function encodeTotalSupplyUpdate(
        uint256 nonce,
        uint256 delta
    ) public view returns (bytes memory) {
        return 
            abi.encodePacked(
                nonce, 
                delta,
                address(this)
            );
    }

    // ========= Owner-only ========

    /**
     * @dev Function to stop the contract (Pausable pattern).
     */
    function pause() public onlyOwner {
        _pause();
    }

    /**
     * @dev Function to resume the contract (Pausable pattern).
     */
    function unpause() public onlyOwner {
        _unpause();
    }

    /**
     * @dev Sets the minimum withdrawal amount.
     * @param _minWithdrawAmount The new minimum withdrawal amount.
     * @notice Only the contract owner can call this function.
     */
    function setMinWithdrawAmount(uint256 _minWithdrawAmount) public onlyOwner {
        require(
            _minWithdrawAmount >= DUST_LIMIT,
            "Min withdraw amount should be greater or equal to dust limit"
        );

        minWithdrawAmount = _minWithdrawAmount;
    }

    /**
     * @dev Mint new tokens.
     * Only the owner can call this function.
     * @param _amount The amount of tokens to mint.
     * @param _recipient The address that will receive the minted tokens.
     * @param _btcDepositId The id of the BTC deposit = keccak256(txHash, vout)
     */
    function mint(
        uint256 _amount,
        address _recipient,
        bytes32 _btcDepositId
    ) public whenNotPaused onlyOwner{
        _mint(_amount, _recipient, _btcDepositId);
    }

    /**
     * @dev Mints `_amount` of BTC to `_recipient` with a signature.
     * Anyone can use if they have valid signature from the owner.
     */
    function mint(
        IStBTC.MintInvoice calldata invoice,
        bytes calldata signature
    )
        public
        whenNotPaused
        onlyValidator(
            MESSAGE_MINT,
            encodeInvoice(invoice),
            signature
        )
    {
        _mint(invoice.amount, invoice.recipient, invoice.btcDepositId);
    }

    /**
     * @dev Mint new tokens.
     * @param _amount The amount of tokens to mint.
     * @param _recipient The address that will receive the minted tokens.
     * @param _btcDepositId The id of the BTC deposit = keccak256(txHash, vout)
     */
    function _mint(
        uint256 _amount,
        address _recipient,
        bytes32 _btcDepositId
    ) internal {
        require(_amount > 0, "MINT_AMOUNT_ZERO");
        require(_amount < 21_000_000 * BTC, "MINT_AMOUNT_TOO_BIG");

        require(_recipient != address(this), "MINT_TO_THE_CONTRACT_ADDRESS");

        require(
            btcDepositIds[_btcDepositId] == false,
            "MINT_ALREADY_PROCESSED"
        );
        btcDepositIds[_btcDepositId] = true;

        uint256 sharesToMint;
        if (totalShares == 0 || totalPooledBTC == 0) {
            sharesToMint = _amount;
        } else {
            sharesToMint = (_amount * totalShares) / totalPooledBTC;
        }

        totalPooledBTC += _amount;
        totalShares += sharesToMint;

        shares[_recipient] += sharesToMint;

        emit Rebase(totalPooledBTC, totalShares);
        emit MintBtcEvent(_recipient, _amount, _btcDepositId);
    }

    function mintRewards(
        uint nonce,
        uint delta
    ) external whenNotPaused onlyOwner {
        require(
            nonce == totalSupplyUpdateNonce,
            "Invalid update total supply nonce"
        );
        totalSupplyUpdateNonce += 1;

        if (delta == 0) {
            revert("delta is 0");
        }

        bytes32 rewardId = getTotalSupplyUpdateHash(nonce, delta);
        require(btcDepositIds[rewardId] == false, "UPDATE_ALREADY_PROCESSED");
        btcDepositIds[rewardId] = true;

        totalPooledBTC += delta;

        emit TotalSupplyUpdatedEvent(nonce, totalPooledBTC, totalShares);
    }

    function mintRewards(
        uint256 nonce,
        uint256 delta,
        bytes calldata signature
    )
        external
        whenNotPaused
        onlyValidator(
            MESSAGE_UPDATE_TOTAL_SUPPLY,
            encodeTotalSupplyUpdate(nonce, delta),
            signature
        )
    {
        require(
            nonce == totalSupplyUpdateNonce,
            "Invalid update total supply nonce"
        );
        totalSupplyUpdateNonce += 1;

        if (delta == 0) {
            revert("delta is 0");
        }

        bytes32 rewardId = getTotalSupplyUpdateHash(nonce, delta);
        require(btcDepositIds[rewardId] == false, "UPDATE_ALREADY_PROCESSED");
        btcDepositIds[rewardId] = true;

        totalPooledBTC += delta;
        
        emit TotalSupplyUpdatedEvent(nonce, totalPooledBTC, totalShares);
    }

    function redeem(
        uint256 _amount,
        string calldata BTCAddress
    ) public whenNotPaused {
        require(
            _amount >= minWithdrawAmount,
            "The sent value must be greater or equal to min withdraw amount"
        );
        require(
            validateBitcoinAddress(network, BTCAddress),
            "The sent BTC address is not valid"
        );

        // balance check in the following function
        _burn(_amount);

        redeemCounter += 1;
        emit RedeemBtcEvent(msg.sender, BTCAddress, _amount, redeemCounter);
    }

    function _burn(uint256 _amount) internal {
        uint256 userBalance = balanceOf(msg.sender);

        require(_amount <= userBalance, "BURN_AMOUNT_EXCEEDS_BALANCE");

        uint256 sharesToBurn = (_amount * totalShares) / totalPooledBTC;

        totalPooledBTC -= _amount;
        totalShares -= sharesToBurn;

        shares[msg.sender] -= sharesToBurn;
        
        emit Rebase(totalPooledBTC, totalShares);     
        emit Transfer(msg.sender, address(0), _amount);
    }
}
