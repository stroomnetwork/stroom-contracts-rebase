// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

interface IStBTC {
    struct MintInvoice {
        bytes32 btcDepositId;
        address recipient;
        uint256 amount;
    }

    function mint(uint256 amount, address to, bytes32 hash) external;

    function redeem(uint256 _amount, string calldata BTCAddress) external;
   
    event MintBtcEvent(
        address indexed _to,
        uint256 _value,
        bytes32 _btcDepositId
    );

    event RedeemBtcEvent(
        address indexed _from,
        string _BTCAddress,
        uint256 _value,
        uint256 _id
    );

    event TotalSupplyUpdatedEvent(
        uint256 _nonce,
        uint256 _totalBTCSupply,
        uint256 _totalShares
    );

    event Rebase(
        uint256 newTotalPooledBTC, 
        uint256 newTotalShares
    );
}
