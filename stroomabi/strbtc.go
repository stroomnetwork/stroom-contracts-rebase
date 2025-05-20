// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stroomabi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StrBTCMintInvoice is an auto generated low-level Go binding around an user-defined struct.
type StrBTCMintInvoice struct {
	BtcDepositId [32]byte
	Recipient    common.Address
	Amount       *big.Int
}

// StrBtcMetaData contains all meta data concerning the StrBtc contract.
var StrBtcMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONVERTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DUST_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_MINT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_UPDATE_TOTAL_SUPPLY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addConverter\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"btcDepositIds\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"converterBurn\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"converterMint\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeInvoice\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"encodeTotalSupplyUpdate\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMintInvoiceHash\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPooledBTCByShares\",\"inputs\":[{\"name\":\"sharesAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShares\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSharesByPooledBTC\",\"inputs\":[{\"name\":\"btcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalSupplyUpdateHash\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_network\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"},{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRewardTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxRewardPercent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minTimeBetweenRewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintRewards\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"network\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"BTCAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeConverter\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxRewardPercent\",\"inputs\":[{\"name\":\"_maxRewardPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinTimeBetweenRewards\",\"inputs\":[{\"name\":\"_minTimeBetweenRewards\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWithdrawAmount\",\"inputs\":[{\"name\":\"_minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupplyUpdateNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"px\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"s\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"m\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConverterBurn\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConverterMint\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintBtcEvent\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_btcDepositId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedeemBtcEvent\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_BTCAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalSupplyUpdatedEvent\",\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalBTCSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AmountBelowMinWithdraw\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotBurnFromZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotMintToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DeltaIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBTCAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSharesOrPooledBTC\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSupplyNonce\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxRewardPercentTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinTimeBetweenRewardsTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinWithdrawTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAlreadyProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountTooBig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintToContractAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RewardTooBig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardTooFrequent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UpdateAlreadyProcessed\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557612baa908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714611b8057508063026034f014611b4f57806306fdde0314611a8d578063095ea7b314611a0c5780630fde6e55146119db57806313776a8d146119b257806318160ddd1461199457806321172f3b1461178757806323b872dd146116a4578063248a9ca31461168657806324b76fd51461151857806325e0a33f146114dd5780632792949d146114be5780632f2ff15d1461148d5780632f4cfd4714611407578063313ce567146113eb5780633357325e146113655780633395b75a1461134757806336568abe146113015780633a98ef39146112e35780633f4ba83a146112605780634219347314611243578063457e1a49146112255780635187599d14610cd457806355573c7714610c7f5780635abdb0dc14610c445780635c975abb14610c14578063603c6a6714610bbe5780636739afca14610b935780636ce1c4dc14610b5f57806370a0823114610b3c578063715018a614610ad257806380239a0414610ab95780638456cb5914610a455780638c659bf214610a2b5780638da5cb5b146109f557806391d148541461099b57806395d89b41146108b45780639cee9492146108795780639e76a007146108455780639fd9568714610654578063a217fddf14610638578063a7ce45651461061a578063a9059cbb146105e9578063af7f131f146105cb578063b0390325146105ad578063c4d66de814610551578063c87b7a231461052b578063d547741f146104f5578063dd62ed3e146104ac578063e07fbd00146103d4578063e0a04610146103a4578063e737557f14610343578063f04da65b14610309578063f2fde38b146102de578063f376ebbb146102b55763f8077fae1461029257600080fd5b346102b05760003660031901126102b0576020600854604051908152f35b600080fd5b346102b05760003660031901126102b0576000546040516001600160a01b039091168152602090f35b346102b05760203660031901126102b0576103076102fa611c1b565b6103026128e5565b6121f1565b005b346102b05760203660031901126102b0576001600160a01b0361032a611c1b565b16600052600a6020526020604060002054604051908152f35b346102b05760003660031901126102b0576103a0604051610365604082611c75565b601a8152795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b6020820152604051918291602083526020830190611bf6565b0390f35b346102b05760603660031901126102b0576103a06103c061219a565b604051918291602083526020830190611bf6565b346102b05760603660031901126102b057600061044a602060018060a01b0383541660405190610405604083611c75565b60138252725354524f4f4d5f4d494e545f494e564f49434560681b8383015261042c61219a565b604051635e50560360e11b8152948593849283929160048401612088565b03915afa9081156104a1578291610467575b602082604051908152f35b90506020813d602011610499575b8161048260209383611c75565b810103126104955760209150518261045c565b5080fd5b3d9150610475565b6040513d84823e3d90fd5b346102b05760403660031901126102b0576104c5611c1b565b6104d66104d0611c31565b91612019565b9060018060a01b03166000526020526020604060002054604051908152f35b346102b05760403660031901126102b057610307600435610514611c31565b9061052661052182611ff8565b6124dc565b612845565b346102b05760203660031901126102b0576020610549600435612169565b604051908152f35b346102b05760203660031901126102b05761056a611c1b565b61057261291b565b61057a61291b565b61058261291b565b61058b336121f1565b600080546001600160a01b0319166001600160a01b0392909216919091179055005b346102b05760003660031901126102b0576020600754604051908152f35b346102b05760003660031901126102b0576020600354604051908152f35b346102b05760403660031901126102b05761060f610605611c1b565b60243590336122de565b602060405160018152f35b346102b05760003660031901126102b0576020600254604051908152f35b346102b05760003660031901126102b057602060405160008152f35b346102b0573660031901608081126102b0576060136102b05760643567ffffffffffffffff81116102b05761068f6020913690600401611c47565b91906106996122b4565b6040516106a7604082611c75565b60138152725354524f4f4d5f4d494e545f494e564f49434560681b838201526106ce61219a565b60018060a01b0360005416916106fa604051968795869485946311ee58a960e01b865260048601611f34565b03915afa9081156108395760009161080a575b50156107f9576024356001600160a01b03811690604435908281036102b0576004359180156107e857660775f05a0740008110156107d7573084146107c65782600052600960205260ff604060002054166107b557826107a8826040947fb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa46996600052600960205285600020600160ff19825416179055612712565b82519182526020820152a2005b637bdb87d160e01b60005260046000fd5b63603acaa960e11b60005260046000fd5b636006659960e11b60005260046000fd5b631cebf66f60e11b60005260046000fd5b636227817160e01b60005260046000fd5b61082c915060203d602011610832575b6108248183611c75565b810190611efb565b8161070d565b503d61081a565b6040513d6000823e3d90fd5b346102b05760203660031901126102b057610307610861611c1b565b61086961243b565b610874610521611faf565b6127af565b346102b05760203660031901126102b05760043561089561243b565b6103e881116108a357600655005b630737692960e31b60005260046000fd5b346102b05760003660031901126102b0576040516000600080516020612ab5833981519152546108e381611cd0565b8084529060018116908115610977575060011461090b575b6103a0836103c081850382611c75565b600080516020612ab583398151915260009081527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa939250905b80821061095d575090915081016020016103c06108fb565b919260018160209254838588010152019101909291610945565b60ff191660208086019190915291151560051b840190910191506103c090506108fb565b346102b05760403660031901126102b0576109b4611c31565b600435600052600080516020612af583398151915260205260406000209060018060a01b0316600052602052602060ff604060002054166040519015158152f35b346102b05760003660031901126102b057600080516020612ad5833981519152546040516001600160a01b039091168152602090f35b346102b0576103a06103c0610a3f36611cba565b90612142565b346102b05760003660031901126102b057610a5e61243b565b610a666122b4565b600160ff19600080516020612b15833981519152541617600080516020612b15833981519152557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346102b0576020610549610acc36611cba565b906120ad565b346102b05760003660031901126102b057610aeb6128e5565b600080516020612ad583398151915280546001600160a01b031981169091556000906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346102b05760203660031901126102b0576020610549610b5a611c1b565b612052565b346102b05760203660031901126102b057610307610b7b611c1b565b610b8361243b565b610b8e610521611faf565b6125d8565b346102b05760003660031901126102b057602060ff60005460a01c16610bbc6040518092611c97565bf35b346102b05760003660031901126102b0576103a0604051610be0604082611c75565b60138152725354524f4f4d5f4d494e545f494e564f49434560681b6020820152604051918291602083526020830190611bf6565b346102b05760003660031901126102b057602060ff600080516020612b1583398151915254166040519015158152f35b346102b05760203660031901126102b057600435610c6061243b565b6102228110610c6e57600155005b6343a87e9960e11b60005260046000fd5b346102b05760203660031901126102b05760055480158015610cca575b610cb957610549602091610cb4600454600435611f7c565b611f8f565b6340a3daff60e11b60005260046000fd5b5060045415610c9c565b346102b05760403660031901126102b05760043560048110156102b057610cf9611c31565b90600080516020612b358339815191525460ff8160401c16159067ffffffffffffffff81168015908161121d575b6001149081611213575b15908161120a575b506111f95767ffffffffffffffff198116600117600080516020612b3583398151915255816111cc575b5060405192610d73604085611c75565b600e84526d29ba3937b7b6902134ba31b7b4b760911b602085015260405193610d9d604086611c75565b600685526573747242544360d01b6020860152610db861291b565b610dc061291b565b80519067ffffffffffffffff82116110ac578190610dec600080516020612a9583398151915254611cd0565b601f811161114f575b50602090601f83116001146110cd576000926110c2575b50508160011b916000199060031b1c191617600080516020612a95833981519152555b835167ffffffffffffffff81116110ac57610e58600080516020612ab583398151915254611cd0565b601f811161103a575b50602094601f8211600114610fb957948192939495600092610fae575b50508160011b916000199060031b1c191617600080516020612ab5833981519152555b610ea961291b565b610eb161291b565b60ff19600080516020612b158339815191525416600080516020612b1583398151915255610edd61291b565b610ee561291b565b610eed61291b565b610ef561291b565b610efe336121f1565b60008054611b586001556001600160a81b0319166001600160a01b039092169190911760a09390931b60ff60a01b16929092178255610f3c33612526565b5060646006556201518060075581600855610f545780f35b68ff000000000000000019600080516020612b358339815191525416600080516020612b35833981519152557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b015190508580610e7e565b601f19821695600080516020612ab5833981519152600052806000209160005b88811061102257508360019596979810611009575b505050811b01600080516020612ab583398151915255610ea1565b015160001960f88460031b161c19169055858080610fee565b91926020600181928685015181550194019201610fd9565b600080516020612ab58339815191526000527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa601f830160051c810191602084106110a2575b601f0160051c01905b8181106110965750610e61565b60008155600101611089565b9091508190611080565b634e487b7160e01b600052604160045260246000fd5b015190508680610e0c565b600080516020612a9583398151915260009081528281209350601f198516905b818110611137575090846001959493921061111e575b505050811b01600080516020612a9583398151915255610e2f565b015160001960f88460031b161c19169055868080611103565b929360206001819287860151815501950193016110ed565b600080516020612a958339815191526000529091507f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0601f840160051c810191602085106111c2575b90601f859493920160051c01905b8181106111b35750610df5565b600081558493506001016111a6565b9091508190611198565b68ffffffffffffffffff19166801000000000000000117600080516020612b358339815191525583610d63565b63f92ee8a960e01b60005260046000fd5b90501585610d39565b303b159150610d31565b839150610d27565b346102b05760003660031901126102b0576020600154604051908152f35b346102b05760003660031901126102b05760206040516102228152f35b346102b05760003660031901126102b05761127961243b565b600080516020612b158339815191525460ff8116156112d25760ff1916600080516020612b15833981519152557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b60005260046000fd5b346102b05760003660031901126102b0576020600554604051908152f35b346102b05760403660031901126102b05761131a611c31565b336001600160a01b038216036113365761030790600435612845565b63334bd91960e11b60005260046000fd5b346102b05760003660031901126102b0576020600654604051908152f35b346102b05760403660031901126102b05761137e611c1b565b6024356113896122b4565b61139161248e565b6001600160a01b0382169182156113da57816113ac916123a4565b6040519081527f25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad56760203392a3005b63a30d2d8760e01b60005260046000fd5b346102b05760003660031901126102b057602060405160088152f35b346102b05760403660031901126102b057611420611c1b565b60243561142b6122b4565b61143361248e565b6001600160a01b03821691821561147c578161144e91612712565b6040519081527f76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e60203392a3005b6389a4ea1960e01b60005260046000fd5b346102b05760403660031901126102b0576103076004356114ac611c31565b906114b961052182611ff8565b612670565b346102b05760003660031901126102b05760206040516305f5e1008152f35b346102b05760203660031901126102b0576004356114f961243b565b610e10811061150757600755005b63a7bbe41d60e01b60005260046000fd5b346102b05760403660031901126102b05760043560243567ffffffffffffffff81116102b05761154c903690600401611c47565b6115546122b4565b600154831061167557600054604051632ce1e0f560e01b81529061158290600483019060a01c60ff16611c97565b604060248201526020818061159b604482018688611f13565b038173__$113774c24a411e022512c75ffd4c4add5e$__5af490811561083957600091611656575b5015611645576115d383336123a4565b6002546001810180911161162f5761161c7f83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e938260025560405193606085526060850191611f13565b93602083015260408201528033930390a2005b634e487b7160e01b600052601160045260246000fd5b6312d58f2760e21b60005260046000fd5b61166f915060203d602011610832576108248183611c75565b846115c3565b633813eacd60e21b60005260046000fd5b346102b05760203660031901126102b0576020610549600435611ff8565b346102b05760603660031901126102b0576116bd611c1b565b6116c5611c31565b604435906116d283612019565b336000908152602091909152604090205492600184016116f7575b61060f93506122de565b82841061176a576001600160a01b0381161561175457331561173e5761060f9361172082612019565b60018060a01b033316600052602052836040600020910390556116ed565b634a1406b160e11b600052600060045260246000fd5b63e602df0560e01b600052600060045260246000fd5b8284637dc7a0d960e11b6000523360045260245260445260646000fd5b346102b05760603660031901126102b05760043560243560443567ffffffffffffffff81116102b0576117c06020913690600401611c47565b91906117ca6122b4565b6040516117d8604082611c75565b601a8152795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b838201526118088587612142565b60018060a01b036000541691611834604051968795869485946311ee58a960e01b865260048601611f34565b03915afa90811561083957600091611975575b50156107f957600354808303611964576001810180911161162f5760035580156119535761187a60085460075490611f6f565b42106119435761271061189260045460065490611f7c565b048111611932576118a381836120ad565b9182600052600960205260ff60406000205416611921576119026060927f339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b934409460005260096020526040600020600160ff19825416179055600454611f6f565b80600455426008556005549060405192835260208301526040820152a1005b63ce199d1760e01b60005260046000fd5b630779a2a360e31b60005260046000fd5b6283cb6b60e51b60005260046000fd5b630aa7fcbb60e21b60005260046000fd5b6336b118d760e11b60005260046000fd5b61198e915060203d602011610832576108248183611c75565b83611847565b346102b05760003660031901126102b0576020600454604051908152f35b346102b05760003660031901126102b0576020604051600080516020612b558339815191528152f35b346102b05760803660031901126102b0576020611a02606435604435602435600435611d17565b6040519015158152f35b346102b05760403660031901126102b057611a25611c1b565b602435903315611754576001600160a01b031690811561173e57611a4833612019565b82600052602052806040600020556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b346102b05760003660031901126102b0576040516000600080516020612a9583398151915254611abc81611cd0565b80845290600181169081156109775750600114611ae3576103a0836103c081850382611c75565b600080516020612a9583398151915260009081527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0939250905b808210611b35575090915081016020016103c06108fb565b919260018160209254838588010152019101909291611b1d565b346102b05760203660031901126102b0576004356000526009602052602060ff604060002054166040519015158152f35b346102b05760203660031901126102b0576004359063ffffffff60e01b82168092036102b057602091637965db0b60e01b8114908115611bc2575b5015158152f35b6301ffc9a760e01b14905083611bbb565b60005b838110611be65750506000910152565b8181015183820152602001611bd6565b90602091611c0f81518092818552858086019101611bd3565b601f01601f1916010190565b600435906001600160a01b03821682036102b057565b602435906001600160a01b03821682036102b057565b9181601f840112156102b05782359167ffffffffffffffff83116102b057602083818601950101116102b057565b90601f8019910116810190811067ffffffffffffffff8211176110ac57604052565b906004821015611ca45752565b634e487b7160e01b600052602160045260246000fd5b60409060031901126102b0576004359060243590565b90600182811c92168015611d00575b6020831014611cea57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611cdf565b9190820391821161162f57565b90926401000003d0198210801590611eeb575b8015611ecf575b611ec657611d3e84612267565b94909415611ebc57611dd060009160209360405190858201927f7bb52d7a9fef58323eb1bf7a407db382d2f3f2d81bb1224f49fe518f6d48d37c84527f7bb52d7a9fef58323eb1bf7a407db382d2f3f2d81bb1224f49fe518f6d48d37c6040840152606083015286608083015260a082015260a08152611dbf60c082611c75565b604051928392839251928391611bd3565b8101039060025afa156108395770014551231950b75fc4402da1732fc9bebe1960005106918170014551231950b75fc4402da1732fc9bebe19910970014551231950b75fc4402da1732fc9bebe19039170014551231950b75fc4402da1732fc9bebe19831161162f578170014551231950b75fc4402da1732fc9bebe19910970014551231950b75fc4402da1732fc9bebe19039070014551231950b75fc4402da1732fc9bebe19821161162f57602092600092608092604051928352601b868401526040830152606082015282805260015afa15610839576000516001600160a01b0390811691161490565b5050505050600090565b50505050600090565b5070014551231950b75fc4402da1732fc9bebe19831015611d31565b506401000003d019841015611d2a565b908160209103126102b0575180151581036102b05790565b908060209392818452848401376000828201840152601f01601f1916010190565b9290611f6c9492611f50611f5e92606087526060870190611bf6565b908582036020870152611bf6565b926040818503910152611f13565b90565b9190820180921161162f57565b8181029291811591840414171561162f57565b8115611f99570490565b634e487b7160e01b600052601260045260246000fd5b600080516020612b55833981519152600052600080516020612af58339815191526020527f87e0fa9230cb9c98f35bc37ac101b2496acf1ea71ed1f9e77de6b9cabae5f6385490565b600052600080516020612af583398151915260205260016040600020015490565b6001600160a01b031660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b600554801561208157611f6c9160018060a01b0316600052600a602052610cb460406000205460045490611f7c565b5050600090565b909161209f611f6c93604084526040840190611bf6565b916020818403910152611bf6565b6020906120fd9261042c60018060a01b036000541691604051936120d2604086611c75565b601a8552795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b86860152612142565b03915afa90811561083957600091612113575090565b90506020813d60201161213a575b8161212e60209383611c75565b810103126102b0575190565b3d9150612121565b9060405191602083015260408201523060601b606082015260548152611f6c607482611c75565b6005549081158015612190575b610cb957611f6c9161218791611f7c565b60045490611f8f565b5060045415612176565b6024356001600160a01b03811681036102b057604051906bffffffffffffffffffffffff199060601b166020820152604435603482015260043560548201523060601b607482015260688152611f6c608882611c75565b6001600160a01b0316801561225157600080516020612ad583398151915280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b631e4fbdf760e01b600052600060045260246000fd5b61227081612949565b919091156122aa576040519160208301918252604083015260408252612297606083611c75565b905190206001600160a01b031690600190565b5050600090600090565b60ff600080516020612b1583398151915254166122cd57565b63d93c066560e01b60005260046000fd5b916001600160a01b03831691821561238e576001600160a01b0316928315612378578161230c600092612052565b10612369578160209160406123417fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef95612169565b91868152600a8552818120838154039055878152600a855220908154019055604051908152a3565b631e9acf1760e31b8152600490fd5b63ec442f0560e01b600052600060045260246000fd5b634b637e8f60e11b600052600060045260246000fd5b906001600160a01b03821690811561238e57806123c2600094612052565b1061242c576020816123f47fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef93612169565b61240082600454611d0a565b60045561240f81600554611d0a565b600555848652600a835260408620908154039055604051908152a3565b631e9acf1760e31b8352600483fd5b3360009081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff161561247457565b63e2517d3f60e01b60005233600452600060245260446000fd5b336000908152600080516020612a75833981519152602052604090205460ff16156124b557565b63e2517d3f60e01b60005233600452600080516020612b5583398151915260245260446000fd5b6000818152600080516020612af58339815191526020908152604080832033845290915290205460ff161561250e5750565b63e2517d3f60e01b6000523360045260245260446000fd5b6001600160a01b03811660009081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff166125d2576001600160a01b031660008181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b6001600160a01b0381166000908152600080516020612a75833981519152602052604090205460ff166125d2576001600160a01b03166000818152600080516020612a7583398151915260205260408120805460ff19166001179055339190600080516020612b55833981519152907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b6000818152600080516020612af5833981519152602090815260408083206001600160a01b038616845290915290205460ff16612081576000818152600080516020612af5833981519152602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b6001600160a01b0316908115612378577fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020600092600554801580156127a5575b8514612792576127758280925b61276d85600454611f6f565b600455611f6f565b600555858552600a835260408520908154019055604051908152a3565b61277561279e83612169565b8092612761565b5060045415612754565b6001600160a01b0381166000908152600080516020612a75833981519152602052604090205460ff16156125d2576001600160a01b03166000818152600080516020612a7583398151915260205260408120805460ff19169055339190600080516020612b55833981519152907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6000818152600080516020612af5833981519152602090815260408083206001600160a01b038616845290915290205460ff1615612081576000818152600080516020612af5833981519152602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b600080516020612ad5833981519152546001600160a01b0316330361290657565b63118cdaa760e01b6000523360045260246000fd5b60ff600080516020612b358339815191525460401c161561293857565b631afcd79f60e31b60005260046000fd5b6401000003d0198110156129bc57612981906401000003d01990816007816000840908906401000003d01990818180099009086129c5565b60018116612990575b90600190565b6401000003d019036401000003d01981111561298a57634e487b7160e01b600052601160045260246000fd5b50600090600090565b80156125d257600190600160ff1b90815b6129df57505090565b90916401000003d0199063400000f4600160fe1b0384161515830a9082908009096401000003d0199063400000f4600160fe1b03600185901c161515830a9082908009096401000003d0199063400000f4600160fe1b03600285901c161515830a9082908009096401000003d0199063400000f4600160fe1b03600385901c161515830a9082908009099160041c90816129d656fe87e0fa9230cb9c98f35bc37ac101b2496acf1ea71ed1f9e77de6b9cabae5f63752c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace049016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a001cf336fddcc7dc48127faf7a5b80ee54fce73ef647eecd31c24bb6cce3ac3eefa2646970667358221220e4b6febad6314449900f18ec1d7bffef76ab332b4eb657314acd416053d2a0e364736f6c634300081b0033",
}

// StrBtcABI is the input ABI used to generate the binding from.
// Deprecated: Use StrBtcMetaData.ABI instead.
var StrBtcABI = StrBtcMetaData.ABI

// StrBtcBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrBtcMetaData.Bin instead.
var StrBtcBin = StrBtcMetaData.Bin

// DeployStrBtc deploys a new Ethereum contract, binding an instance of StrBtc to it.
func DeployStrBtc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StrBtc, error) {
	parsed, err := StrBtcMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrBtcBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrBtc{StrBtcCaller: StrBtcCaller{contract: contract}, StrBtcTransactor: StrBtcTransactor{contract: contract}, StrBtcFilterer: StrBtcFilterer{contract: contract}}, nil
}

// StrBtc is an auto generated Go binding around an Ethereum contract.
type StrBtc struct {
	StrBtcCaller     // Read-only binding to the contract
	StrBtcTransactor // Write-only binding to the contract
	StrBtcFilterer   // Log filterer for contract events
}

// StrBtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrBtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrBtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrBtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrBtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrBtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrBtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrBtcSession struct {
	Contract     *StrBtc           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrBtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrBtcCallerSession struct {
	Contract *StrBtcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StrBtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrBtcTransactorSession struct {
	Contract     *StrBtcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrBtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrBtcRaw struct {
	Contract *StrBtc // Generic contract binding to access the raw methods on
}

// StrBtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrBtcCallerRaw struct {
	Contract *StrBtcCaller // Generic read-only contract binding to access the raw methods on
}

// StrBtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrBtcTransactorRaw struct {
	Contract *StrBtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrBtc creates a new instance of StrBtc, bound to a specific deployed contract.
func NewStrBtc(address common.Address, backend bind.ContractBackend) (*StrBtc, error) {
	contract, err := bindStrBtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrBtc{StrBtcCaller: StrBtcCaller{contract: contract}, StrBtcTransactor: StrBtcTransactor{contract: contract}, StrBtcFilterer: StrBtcFilterer{contract: contract}}, nil
}

// NewStrBtcCaller creates a new read-only instance of StrBtc, bound to a specific deployed contract.
func NewStrBtcCaller(address common.Address, caller bind.ContractCaller) (*StrBtcCaller, error) {
	contract, err := bindStrBtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrBtcCaller{contract: contract}, nil
}

// NewStrBtcTransactor creates a new write-only instance of StrBtc, bound to a specific deployed contract.
func NewStrBtcTransactor(address common.Address, transactor bind.ContractTransactor) (*StrBtcTransactor, error) {
	contract, err := bindStrBtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrBtcTransactor{contract: contract}, nil
}

// NewStrBtcFilterer creates a new log filterer instance of StrBtc, bound to a specific deployed contract.
func NewStrBtcFilterer(address common.Address, filterer bind.ContractFilterer) (*StrBtcFilterer, error) {
	contract, err := bindStrBtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrBtcFilterer{contract: contract}, nil
}

// bindStrBtc binds a generic wrapper to an already deployed contract.
func bindStrBtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrBtcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrBtc *StrBtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrBtc.Contract.StrBtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrBtc *StrBtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrBtc.Contract.StrBtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrBtc *StrBtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrBtc.Contract.StrBtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrBtc *StrBtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrBtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrBtc *StrBtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrBtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrBtc *StrBtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrBtc.Contract.contract.Transact(opts, method, params...)
}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_StrBtc *StrBtcCaller) BTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "BTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_StrBtc *StrBtcSession) BTC() (*big.Int, error) {
	return _StrBtc.Contract.BTC(&_StrBtc.CallOpts)
}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) BTC() (*big.Int, error) {
	return _StrBtc.Contract.BTC(&_StrBtc.CallOpts)
}

// CONVERTERROLE is a free data retrieval call binding the contract method 0x13776a8d.
//
// Solidity: function CONVERTER_ROLE() view returns(bytes32)
func (_StrBtc *StrBtcCaller) CONVERTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "CONVERTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONVERTERROLE is a free data retrieval call binding the contract method 0x13776a8d.
//
// Solidity: function CONVERTER_ROLE() view returns(bytes32)
func (_StrBtc *StrBtcSession) CONVERTERROLE() ([32]byte, error) {
	return _StrBtc.Contract.CONVERTERROLE(&_StrBtc.CallOpts)
}

// CONVERTERROLE is a free data retrieval call binding the contract method 0x13776a8d.
//
// Solidity: function CONVERTER_ROLE() view returns(bytes32)
func (_StrBtc *StrBtcCallerSession) CONVERTERROLE() ([32]byte, error) {
	return _StrBtc.Contract.CONVERTERROLE(&_StrBtc.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StrBtc *StrBtcCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StrBtc *StrBtcSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StrBtc.Contract.DEFAULTADMINROLE(&_StrBtc.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StrBtc *StrBtcCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StrBtc.Contract.DEFAULTADMINROLE(&_StrBtc.CallOpts)
}

// DUSTLIMIT is a free data retrieval call binding the contract method 0x42193473.
//
// Solidity: function DUST_LIMIT() view returns(uint256)
func (_StrBtc *StrBtcCaller) DUSTLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "DUST_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DUSTLIMIT is a free data retrieval call binding the contract method 0x42193473.
//
// Solidity: function DUST_LIMIT() view returns(uint256)
func (_StrBtc *StrBtcSession) DUSTLIMIT() (*big.Int, error) {
	return _StrBtc.Contract.DUSTLIMIT(&_StrBtc.CallOpts)
}

// DUSTLIMIT is a free data retrieval call binding the contract method 0x42193473.
//
// Solidity: function DUST_LIMIT() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) DUSTLIMIT() (*big.Int, error) {
	return _StrBtc.Contract.DUSTLIMIT(&_StrBtc.CallOpts)
}

// MESSAGEMINT is a free data retrieval call binding the contract method 0x603c6a67.
//
// Solidity: function MESSAGE_MINT() view returns(bytes)
func (_StrBtc *StrBtcCaller) MESSAGEMINT(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "MESSAGE_MINT")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MESSAGEMINT is a free data retrieval call binding the contract method 0x603c6a67.
//
// Solidity: function MESSAGE_MINT() view returns(bytes)
func (_StrBtc *StrBtcSession) MESSAGEMINT() ([]byte, error) {
	return _StrBtc.Contract.MESSAGEMINT(&_StrBtc.CallOpts)
}

// MESSAGEMINT is a free data retrieval call binding the contract method 0x603c6a67.
//
// Solidity: function MESSAGE_MINT() view returns(bytes)
func (_StrBtc *StrBtcCallerSession) MESSAGEMINT() ([]byte, error) {
	return _StrBtc.Contract.MESSAGEMINT(&_StrBtc.CallOpts)
}

// MESSAGEUPDATETOTALSUPPLY is a free data retrieval call binding the contract method 0xe737557f.
//
// Solidity: function MESSAGE_UPDATE_TOTAL_SUPPLY() view returns(bytes)
func (_StrBtc *StrBtcCaller) MESSAGEUPDATETOTALSUPPLY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "MESSAGE_UPDATE_TOTAL_SUPPLY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MESSAGEUPDATETOTALSUPPLY is a free data retrieval call binding the contract method 0xe737557f.
//
// Solidity: function MESSAGE_UPDATE_TOTAL_SUPPLY() view returns(bytes)
func (_StrBtc *StrBtcSession) MESSAGEUPDATETOTALSUPPLY() ([]byte, error) {
	return _StrBtc.Contract.MESSAGEUPDATETOTALSUPPLY(&_StrBtc.CallOpts)
}

// MESSAGEUPDATETOTALSUPPLY is a free data retrieval call binding the contract method 0xe737557f.
//
// Solidity: function MESSAGE_UPDATE_TOTAL_SUPPLY() view returns(bytes)
func (_StrBtc *StrBtcCallerSession) MESSAGEUPDATETOTALSUPPLY() ([]byte, error) {
	return _StrBtc.Contract.MESSAGEUPDATETOTALSUPPLY(&_StrBtc.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StrBtc *StrBtcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StrBtc *StrBtcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StrBtc.Contract.Allowance(&_StrBtc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StrBtc *StrBtcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StrBtc.Contract.Allowance(&_StrBtc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StrBtc *StrBtcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StrBtc *StrBtcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StrBtc.Contract.BalanceOf(&_StrBtc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StrBtc *StrBtcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StrBtc.Contract.BalanceOf(&_StrBtc.CallOpts, account)
}

// BtcDepositIds is a free data retrieval call binding the contract method 0x026034f0.
//
// Solidity: function btcDepositIds(bytes32 ) view returns(bool)
func (_StrBtc *StrBtcCaller) BtcDepositIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "btcDepositIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BtcDepositIds is a free data retrieval call binding the contract method 0x026034f0.
//
// Solidity: function btcDepositIds(bytes32 ) view returns(bool)
func (_StrBtc *StrBtcSession) BtcDepositIds(arg0 [32]byte) (bool, error) {
	return _StrBtc.Contract.BtcDepositIds(&_StrBtc.CallOpts, arg0)
}

// BtcDepositIds is a free data retrieval call binding the contract method 0x026034f0.
//
// Solidity: function btcDepositIds(bytes32 ) view returns(bool)
func (_StrBtc *StrBtcCallerSession) BtcDepositIds(arg0 [32]byte) (bool, error) {
	return _StrBtc.Contract.BtcDepositIds(&_StrBtc.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StrBtc *StrBtcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StrBtc *StrBtcSession) Decimals() (uint8, error) {
	return _StrBtc.Contract.Decimals(&_StrBtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StrBtc *StrBtcCallerSession) Decimals() (uint8, error) {
	return _StrBtc.Contract.Decimals(&_StrBtc.CallOpts)
}

// EncodeInvoice is a free data retrieval call binding the contract method 0xe0a04610.
//
// Solidity: function encodeInvoice((bytes32,address,uint256) invoice) view returns(bytes)
func (_StrBtc *StrBtcCaller) EncodeInvoice(opts *bind.CallOpts, invoice StrBTCMintInvoice) ([]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "encodeInvoice", invoice)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeInvoice is a free data retrieval call binding the contract method 0xe0a04610.
//
// Solidity: function encodeInvoice((bytes32,address,uint256) invoice) view returns(bytes)
func (_StrBtc *StrBtcSession) EncodeInvoice(invoice StrBTCMintInvoice) ([]byte, error) {
	return _StrBtc.Contract.EncodeInvoice(&_StrBtc.CallOpts, invoice)
}

// EncodeInvoice is a free data retrieval call binding the contract method 0xe0a04610.
//
// Solidity: function encodeInvoice((bytes32,address,uint256) invoice) view returns(bytes)
func (_StrBtc *StrBtcCallerSession) EncodeInvoice(invoice StrBTCMintInvoice) ([]byte, error) {
	return _StrBtc.Contract.EncodeInvoice(&_StrBtc.CallOpts, invoice)
}

// EncodeTotalSupplyUpdate is a free data retrieval call binding the contract method 0x8c659bf2.
//
// Solidity: function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) view returns(bytes)
func (_StrBtc *StrBtcCaller) EncodeTotalSupplyUpdate(opts *bind.CallOpts, nonce *big.Int, delta *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "encodeTotalSupplyUpdate", nonce, delta)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTotalSupplyUpdate is a free data retrieval call binding the contract method 0x8c659bf2.
//
// Solidity: function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) view returns(bytes)
func (_StrBtc *StrBtcSession) EncodeTotalSupplyUpdate(nonce *big.Int, delta *big.Int) ([]byte, error) {
	return _StrBtc.Contract.EncodeTotalSupplyUpdate(&_StrBtc.CallOpts, nonce, delta)
}

// EncodeTotalSupplyUpdate is a free data retrieval call binding the contract method 0x8c659bf2.
//
// Solidity: function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) view returns(bytes)
func (_StrBtc *StrBtcCallerSession) EncodeTotalSupplyUpdate(nonce *big.Int, delta *big.Int) ([]byte, error) {
	return _StrBtc.Contract.EncodeTotalSupplyUpdate(&_StrBtc.CallOpts, nonce, delta)
}

// GetMintInvoiceHash is a free data retrieval call binding the contract method 0xe07fbd00.
//
// Solidity: function getMintInvoiceHash((bytes32,address,uint256) invoice) view returns(bytes32)
func (_StrBtc *StrBtcCaller) GetMintInvoiceHash(opts *bind.CallOpts, invoice StrBTCMintInvoice) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getMintInvoiceHash", invoice)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMintInvoiceHash is a free data retrieval call binding the contract method 0xe07fbd00.
//
// Solidity: function getMintInvoiceHash((bytes32,address,uint256) invoice) view returns(bytes32)
func (_StrBtc *StrBtcSession) GetMintInvoiceHash(invoice StrBTCMintInvoice) ([32]byte, error) {
	return _StrBtc.Contract.GetMintInvoiceHash(&_StrBtc.CallOpts, invoice)
}

// GetMintInvoiceHash is a free data retrieval call binding the contract method 0xe07fbd00.
//
// Solidity: function getMintInvoiceHash((bytes32,address,uint256) invoice) view returns(bytes32)
func (_StrBtc *StrBtcCallerSession) GetMintInvoiceHash(invoice StrBTCMintInvoice) ([32]byte, error) {
	return _StrBtc.Contract.GetMintInvoiceHash(&_StrBtc.CallOpts, invoice)
}

// GetPooledBTCByShares is a free data retrieval call binding the contract method 0x55573c77.
//
// Solidity: function getPooledBTCByShares(uint256 sharesAmount) view returns(uint256)
func (_StrBtc *StrBtcCaller) GetPooledBTCByShares(opts *bind.CallOpts, sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getPooledBTCByShares", sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledBTCByShares is a free data retrieval call binding the contract method 0x55573c77.
//
// Solidity: function getPooledBTCByShares(uint256 sharesAmount) view returns(uint256)
func (_StrBtc *StrBtcSession) GetPooledBTCByShares(sharesAmount *big.Int) (*big.Int, error) {
	return _StrBtc.Contract.GetPooledBTCByShares(&_StrBtc.CallOpts, sharesAmount)
}

// GetPooledBTCByShares is a free data retrieval call binding the contract method 0x55573c77.
//
// Solidity: function getPooledBTCByShares(uint256 sharesAmount) view returns(uint256)
func (_StrBtc *StrBtcCallerSession) GetPooledBTCByShares(sharesAmount *big.Int) (*big.Int, error) {
	return _StrBtc.Contract.GetPooledBTCByShares(&_StrBtc.CallOpts, sharesAmount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StrBtc *StrBtcCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StrBtc *StrBtcSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StrBtc.Contract.GetRoleAdmin(&_StrBtc.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StrBtc *StrBtcCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StrBtc.Contract.GetRoleAdmin(&_StrBtc.CallOpts, role)
}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address account) view returns(uint256)
func (_StrBtc *StrBtcCaller) GetShares(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getShares", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address account) view returns(uint256)
func (_StrBtc *StrBtcSession) GetShares(account common.Address) (*big.Int, error) {
	return _StrBtc.Contract.GetShares(&_StrBtc.CallOpts, account)
}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address account) view returns(uint256)
func (_StrBtc *StrBtcCallerSession) GetShares(account common.Address) (*big.Int, error) {
	return _StrBtc.Contract.GetShares(&_StrBtc.CallOpts, account)
}

// GetSharesByPooledBTC is a free data retrieval call binding the contract method 0xc87b7a23.
//
// Solidity: function getSharesByPooledBTC(uint256 btcAmount) view returns(uint256)
func (_StrBtc *StrBtcCaller) GetSharesByPooledBTC(opts *bind.CallOpts, btcAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getSharesByPooledBTC", btcAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledBTC is a free data retrieval call binding the contract method 0xc87b7a23.
//
// Solidity: function getSharesByPooledBTC(uint256 btcAmount) view returns(uint256)
func (_StrBtc *StrBtcSession) GetSharesByPooledBTC(btcAmount *big.Int) (*big.Int, error) {
	return _StrBtc.Contract.GetSharesByPooledBTC(&_StrBtc.CallOpts, btcAmount)
}

// GetSharesByPooledBTC is a free data retrieval call binding the contract method 0xc87b7a23.
//
// Solidity: function getSharesByPooledBTC(uint256 btcAmount) view returns(uint256)
func (_StrBtc *StrBtcCallerSession) GetSharesByPooledBTC(btcAmount *big.Int) (*big.Int, error) {
	return _StrBtc.Contract.GetSharesByPooledBTC(&_StrBtc.CallOpts, btcAmount)
}

// GetTotalSupplyUpdateHash is a free data retrieval call binding the contract method 0x80239a04.
//
// Solidity: function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) view returns(bytes32)
func (_StrBtc *StrBtcCaller) GetTotalSupplyUpdateHash(opts *bind.CallOpts, nonce *big.Int, delta *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getTotalSupplyUpdateHash", nonce, delta)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTotalSupplyUpdateHash is a free data retrieval call binding the contract method 0x80239a04.
//
// Solidity: function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) view returns(bytes32)
func (_StrBtc *StrBtcSession) GetTotalSupplyUpdateHash(nonce *big.Int, delta *big.Int) ([32]byte, error) {
	return _StrBtc.Contract.GetTotalSupplyUpdateHash(&_StrBtc.CallOpts, nonce, delta)
}

// GetTotalSupplyUpdateHash is a free data retrieval call binding the contract method 0x80239a04.
//
// Solidity: function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) view returns(bytes32)
func (_StrBtc *StrBtcCallerSession) GetTotalSupplyUpdateHash(nonce *big.Int, delta *big.Int) ([32]byte, error) {
	return _StrBtc.Contract.GetTotalSupplyUpdateHash(&_StrBtc.CallOpts, nonce, delta)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StrBtc *StrBtcCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StrBtc *StrBtcSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StrBtc.Contract.HasRole(&_StrBtc.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StrBtc *StrBtcCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StrBtc.Contract.HasRole(&_StrBtc.CallOpts, role, account)
}

// LastRewardTimestamp is a free data retrieval call binding the contract method 0xf8077fae.
//
// Solidity: function lastRewardTimestamp() view returns(uint256)
func (_StrBtc *StrBtcCaller) LastRewardTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "lastRewardTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardTimestamp is a free data retrieval call binding the contract method 0xf8077fae.
//
// Solidity: function lastRewardTimestamp() view returns(uint256)
func (_StrBtc *StrBtcSession) LastRewardTimestamp() (*big.Int, error) {
	return _StrBtc.Contract.LastRewardTimestamp(&_StrBtc.CallOpts)
}

// LastRewardTimestamp is a free data retrieval call binding the contract method 0xf8077fae.
//
// Solidity: function lastRewardTimestamp() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) LastRewardTimestamp() (*big.Int, error) {
	return _StrBtc.Contract.LastRewardTimestamp(&_StrBtc.CallOpts)
}

// MaxRewardPercent is a free data retrieval call binding the contract method 0x3395b75a.
//
// Solidity: function maxRewardPercent() view returns(uint256)
func (_StrBtc *StrBtcCaller) MaxRewardPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "maxRewardPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRewardPercent is a free data retrieval call binding the contract method 0x3395b75a.
//
// Solidity: function maxRewardPercent() view returns(uint256)
func (_StrBtc *StrBtcSession) MaxRewardPercent() (*big.Int, error) {
	return _StrBtc.Contract.MaxRewardPercent(&_StrBtc.CallOpts)
}

// MaxRewardPercent is a free data retrieval call binding the contract method 0x3395b75a.
//
// Solidity: function maxRewardPercent() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) MaxRewardPercent() (*big.Int, error) {
	return _StrBtc.Contract.MaxRewardPercent(&_StrBtc.CallOpts)
}

// MinTimeBetweenRewards is a free data retrieval call binding the contract method 0xb0390325.
//
// Solidity: function minTimeBetweenRewards() view returns(uint256)
func (_StrBtc *StrBtcCaller) MinTimeBetweenRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "minTimeBetweenRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTimeBetweenRewards is a free data retrieval call binding the contract method 0xb0390325.
//
// Solidity: function minTimeBetweenRewards() view returns(uint256)
func (_StrBtc *StrBtcSession) MinTimeBetweenRewards() (*big.Int, error) {
	return _StrBtc.Contract.MinTimeBetweenRewards(&_StrBtc.CallOpts)
}

// MinTimeBetweenRewards is a free data retrieval call binding the contract method 0xb0390325.
//
// Solidity: function minTimeBetweenRewards() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) MinTimeBetweenRewards() (*big.Int, error) {
	return _StrBtc.Contract.MinTimeBetweenRewards(&_StrBtc.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_StrBtc *StrBtcCaller) MinWithdrawAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "minWithdrawAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_StrBtc *StrBtcSession) MinWithdrawAmount() (*big.Int, error) {
	return _StrBtc.Contract.MinWithdrawAmount(&_StrBtc.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) MinWithdrawAmount() (*big.Int, error) {
	return _StrBtc.Contract.MinWithdrawAmount(&_StrBtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrBtc *StrBtcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrBtc *StrBtcSession) Name() (string, error) {
	return _StrBtc.Contract.Name(&_StrBtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrBtc *StrBtcCallerSession) Name() (string, error) {
	return _StrBtc.Contract.Name(&_StrBtc.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint8)
func (_StrBtc *StrBtcCaller) Network(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "network")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint8)
func (_StrBtc *StrBtcSession) Network() (uint8, error) {
	return _StrBtc.Contract.Network(&_StrBtc.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint8)
func (_StrBtc *StrBtcCallerSession) Network() (uint8, error) {
	return _StrBtc.Contract.Network(&_StrBtc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrBtc *StrBtcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrBtc *StrBtcSession) Owner() (common.Address, error) {
	return _StrBtc.Contract.Owner(&_StrBtc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrBtc *StrBtcCallerSession) Owner() (common.Address, error) {
	return _StrBtc.Contract.Owner(&_StrBtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StrBtc *StrBtcCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StrBtc *StrBtcSession) Paused() (bool, error) {
	return _StrBtc.Contract.Paused(&_StrBtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StrBtc *StrBtcCallerSession) Paused() (bool, error) {
	return _StrBtc.Contract.Paused(&_StrBtc.CallOpts)
}

// RedeemCounter is a free data retrieval call binding the contract method 0xa7ce4565.
//
// Solidity: function redeemCounter() view returns(uint256)
func (_StrBtc *StrBtcCaller) RedeemCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "redeemCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemCounter is a free data retrieval call binding the contract method 0xa7ce4565.
//
// Solidity: function redeemCounter() view returns(uint256)
func (_StrBtc *StrBtcSession) RedeemCounter() (*big.Int, error) {
	return _StrBtc.Contract.RedeemCounter(&_StrBtc.CallOpts)
}

// RedeemCounter is a free data retrieval call binding the contract method 0xa7ce4565.
//
// Solidity: function redeemCounter() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) RedeemCounter() (*big.Int, error) {
	return _StrBtc.Contract.RedeemCounter(&_StrBtc.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StrBtc *StrBtcCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StrBtc *StrBtcSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StrBtc.Contract.SupportsInterface(&_StrBtc.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StrBtc *StrBtcCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StrBtc.Contract.SupportsInterface(&_StrBtc.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StrBtc *StrBtcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StrBtc *StrBtcSession) Symbol() (string, error) {
	return _StrBtc.Contract.Symbol(&_StrBtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StrBtc *StrBtcCallerSession) Symbol() (string, error) {
	return _StrBtc.Contract.Symbol(&_StrBtc.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrBtc *StrBtcCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrBtc *StrBtcSession) TotalShares() (*big.Int, error) {
	return _StrBtc.Contract.TotalShares(&_StrBtc.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) TotalShares() (*big.Int, error) {
	return _StrBtc.Contract.TotalShares(&_StrBtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StrBtc *StrBtcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StrBtc *StrBtcSession) TotalSupply() (*big.Int, error) {
	return _StrBtc.Contract.TotalSupply(&_StrBtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) TotalSupply() (*big.Int, error) {
	return _StrBtc.Contract.TotalSupply(&_StrBtc.CallOpts)
}

// TotalSupplyUpdateNonce is a free data retrieval call binding the contract method 0xaf7f131f.
//
// Solidity: function totalSupplyUpdateNonce() view returns(uint256)
func (_StrBtc *StrBtcCaller) TotalSupplyUpdateNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "totalSupplyUpdateNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyUpdateNonce is a free data retrieval call binding the contract method 0xaf7f131f.
//
// Solidity: function totalSupplyUpdateNonce() view returns(uint256)
func (_StrBtc *StrBtcSession) TotalSupplyUpdateNonce() (*big.Int, error) {
	return _StrBtc.Contract.TotalSupplyUpdateNonce(&_StrBtc.CallOpts)
}

// TotalSupplyUpdateNonce is a free data retrieval call binding the contract method 0xaf7f131f.
//
// Solidity: function totalSupplyUpdateNonce() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) TotalSupplyUpdateNonce() (*big.Int, error) {
	return _StrBtc.Contract.TotalSupplyUpdateNonce(&_StrBtc.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_StrBtc *StrBtcCaller) ValidatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "validatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_StrBtc *StrBtcSession) ValidatorRegistry() (common.Address, error) {
	return _StrBtc.Contract.ValidatorRegistry(&_StrBtc.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_StrBtc *StrBtcCallerSession) ValidatorRegistry() (common.Address, error) {
	return _StrBtc.Contract.ValidatorRegistry(&_StrBtc.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x0fde6e55.
//
// Solidity: function verify(uint256 px, uint256 rx, uint256 s, bytes32 m) pure returns(bool)
func (_StrBtc *StrBtcCaller) Verify(opts *bind.CallOpts, px *big.Int, rx *big.Int, s *big.Int, m [32]byte) (bool, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "verify", px, rx, s, m)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x0fde6e55.
//
// Solidity: function verify(uint256 px, uint256 rx, uint256 s, bytes32 m) pure returns(bool)
func (_StrBtc *StrBtcSession) Verify(px *big.Int, rx *big.Int, s *big.Int, m [32]byte) (bool, error) {
	return _StrBtc.Contract.Verify(&_StrBtc.CallOpts, px, rx, s, m)
}

// Verify is a free data retrieval call binding the contract method 0x0fde6e55.
//
// Solidity: function verify(uint256 px, uint256 rx, uint256 s, bytes32 m) pure returns(bool)
func (_StrBtc *StrBtcCallerSession) Verify(px *big.Int, rx *big.Int, s *big.Int, m [32]byte) (bool, error) {
	return _StrBtc.Contract.Verify(&_StrBtc.CallOpts, px, rx, s, m)
}

// AddConverter is a paid mutator transaction binding the contract method 0x6ce1c4dc.
//
// Solidity: function addConverter(address converter) returns()
func (_StrBtc *StrBtcTransactor) AddConverter(opts *bind.TransactOpts, converter common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "addConverter", converter)
}

// AddConverter is a paid mutator transaction binding the contract method 0x6ce1c4dc.
//
// Solidity: function addConverter(address converter) returns()
func (_StrBtc *StrBtcSession) AddConverter(converter common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.AddConverter(&_StrBtc.TransactOpts, converter)
}

// AddConverter is a paid mutator transaction binding the contract method 0x6ce1c4dc.
//
// Solidity: function addConverter(address converter) returns()
func (_StrBtc *StrBtcTransactorSession) AddConverter(converter common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.AddConverter(&_StrBtc.TransactOpts, converter)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StrBtc *StrBtcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StrBtc *StrBtcSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.Approve(&_StrBtc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StrBtc *StrBtcTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.Approve(&_StrBtc.TransactOpts, spender, value)
}

// ConverterBurn is a paid mutator transaction binding the contract method 0x3357325e.
//
// Solidity: function converterBurn(address from, uint256 amount) returns()
func (_StrBtc *StrBtcTransactor) ConverterBurn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "converterBurn", from, amount)
}

// ConverterBurn is a paid mutator transaction binding the contract method 0x3357325e.
//
// Solidity: function converterBurn(address from, uint256 amount) returns()
func (_StrBtc *StrBtcSession) ConverterBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.ConverterBurn(&_StrBtc.TransactOpts, from, amount)
}

// ConverterBurn is a paid mutator transaction binding the contract method 0x3357325e.
//
// Solidity: function converterBurn(address from, uint256 amount) returns()
func (_StrBtc *StrBtcTransactorSession) ConverterBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.ConverterBurn(&_StrBtc.TransactOpts, from, amount)
}

// ConverterMint is a paid mutator transaction binding the contract method 0x2f4cfd47.
//
// Solidity: function converterMint(address recipient, uint256 amount) returns()
func (_StrBtc *StrBtcTransactor) ConverterMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "converterMint", recipient, amount)
}

// ConverterMint is a paid mutator transaction binding the contract method 0x2f4cfd47.
//
// Solidity: function converterMint(address recipient, uint256 amount) returns()
func (_StrBtc *StrBtcSession) ConverterMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.ConverterMint(&_StrBtc.TransactOpts, recipient, amount)
}

// ConverterMint is a paid mutator transaction binding the contract method 0x2f4cfd47.
//
// Solidity: function converterMint(address recipient, uint256 amount) returns()
func (_StrBtc *StrBtcTransactorSession) ConverterMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.ConverterMint(&_StrBtc.TransactOpts, recipient, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StrBtc *StrBtcTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StrBtc *StrBtcSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.GrantRole(&_StrBtc.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StrBtc *StrBtcTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.GrantRole(&_StrBtc.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry) returns()
func (_StrBtc *StrBtcTransactor) Initialize(opts *bind.TransactOpts, _network uint8, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "initialize", _network, _validatorRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry) returns()
func (_StrBtc *StrBtcSession) Initialize(_network uint8, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.Initialize(&_StrBtc.TransactOpts, _network, _validatorRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry) returns()
func (_StrBtc *StrBtcTransactorSession) Initialize(_network uint8, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.Initialize(&_StrBtc.TransactOpts, _network, _validatorRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorRegistry) returns()
func (_StrBtc *StrBtcTransactor) Initialize0(opts *bind.TransactOpts, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "initialize0", _validatorRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorRegistry) returns()
func (_StrBtc *StrBtcSession) Initialize0(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.Initialize0(&_StrBtc.TransactOpts, _validatorRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorRegistry) returns()
func (_StrBtc *StrBtcTransactorSession) Initialize0(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.Initialize0(&_StrBtc.TransactOpts, _validatorRegistry)
}

// Mint is a paid mutator transaction binding the contract method 0x9fd95687.
//
// Solidity: function mint((bytes32,address,uint256) invoice, bytes signature) returns()
func (_StrBtc *StrBtcTransactor) Mint(opts *bind.TransactOpts, invoice StrBTCMintInvoice, signature []byte) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "mint", invoice, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x9fd95687.
//
// Solidity: function mint((bytes32,address,uint256) invoice, bytes signature) returns()
func (_StrBtc *StrBtcSession) Mint(invoice StrBTCMintInvoice, signature []byte) (*types.Transaction, error) {
	return _StrBtc.Contract.Mint(&_StrBtc.TransactOpts, invoice, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x9fd95687.
//
// Solidity: function mint((bytes32,address,uint256) invoice, bytes signature) returns()
func (_StrBtc *StrBtcTransactorSession) Mint(invoice StrBTCMintInvoice, signature []byte) (*types.Transaction, error) {
	return _StrBtc.Contract.Mint(&_StrBtc.TransactOpts, invoice, signature)
}

// MintRewards is a paid mutator transaction binding the contract method 0x21172f3b.
//
// Solidity: function mintRewards(uint256 nonce, uint256 delta, bytes signature) returns()
func (_StrBtc *StrBtcTransactor) MintRewards(opts *bind.TransactOpts, nonce *big.Int, delta *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "mintRewards", nonce, delta, signature)
}

// MintRewards is a paid mutator transaction binding the contract method 0x21172f3b.
//
// Solidity: function mintRewards(uint256 nonce, uint256 delta, bytes signature) returns()
func (_StrBtc *StrBtcSession) MintRewards(nonce *big.Int, delta *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrBtc.Contract.MintRewards(&_StrBtc.TransactOpts, nonce, delta, signature)
}

// MintRewards is a paid mutator transaction binding the contract method 0x21172f3b.
//
// Solidity: function mintRewards(uint256 nonce, uint256 delta, bytes signature) returns()
func (_StrBtc *StrBtcTransactorSession) MintRewards(nonce *big.Int, delta *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrBtc.Contract.MintRewards(&_StrBtc.TransactOpts, nonce, delta, signature)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StrBtc *StrBtcTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StrBtc *StrBtcSession) Pause() (*types.Transaction, error) {
	return _StrBtc.Contract.Pause(&_StrBtc.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StrBtc *StrBtcTransactorSession) Pause() (*types.Transaction, error) {
	return _StrBtc.Contract.Pause(&_StrBtc.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 _amount, string BTCAddress) returns()
func (_StrBtc *StrBtcTransactor) Redeem(opts *bind.TransactOpts, _amount *big.Int, BTCAddress string) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "redeem", _amount, BTCAddress)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 _amount, string BTCAddress) returns()
func (_StrBtc *StrBtcSession) Redeem(_amount *big.Int, BTCAddress string) (*types.Transaction, error) {
	return _StrBtc.Contract.Redeem(&_StrBtc.TransactOpts, _amount, BTCAddress)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 _amount, string BTCAddress) returns()
func (_StrBtc *StrBtcTransactorSession) Redeem(_amount *big.Int, BTCAddress string) (*types.Transaction, error) {
	return _StrBtc.Contract.Redeem(&_StrBtc.TransactOpts, _amount, BTCAddress)
}

// RemoveConverter is a paid mutator transaction binding the contract method 0x9e76a007.
//
// Solidity: function removeConverter(address converter) returns()
func (_StrBtc *StrBtcTransactor) RemoveConverter(opts *bind.TransactOpts, converter common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "removeConverter", converter)
}

// RemoveConverter is a paid mutator transaction binding the contract method 0x9e76a007.
//
// Solidity: function removeConverter(address converter) returns()
func (_StrBtc *StrBtcSession) RemoveConverter(converter common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.RemoveConverter(&_StrBtc.TransactOpts, converter)
}

// RemoveConverter is a paid mutator transaction binding the contract method 0x9e76a007.
//
// Solidity: function removeConverter(address converter) returns()
func (_StrBtc *StrBtcTransactorSession) RemoveConverter(converter common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.RemoveConverter(&_StrBtc.TransactOpts, converter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrBtc *StrBtcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrBtc *StrBtcSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrBtc.Contract.RenounceOwnership(&_StrBtc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrBtc *StrBtcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrBtc.Contract.RenounceOwnership(&_StrBtc.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StrBtc *StrBtcTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StrBtc *StrBtcSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.RenounceRole(&_StrBtc.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StrBtc *StrBtcTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.RenounceRole(&_StrBtc.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StrBtc *StrBtcTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StrBtc *StrBtcSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.RevokeRole(&_StrBtc.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StrBtc *StrBtcTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.RevokeRole(&_StrBtc.TransactOpts, role, account)
}

// SetMaxRewardPercent is a paid mutator transaction binding the contract method 0x9cee9492.
//
// Solidity: function setMaxRewardPercent(uint256 _maxRewardPercent) returns()
func (_StrBtc *StrBtcTransactor) SetMaxRewardPercent(opts *bind.TransactOpts, _maxRewardPercent *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "setMaxRewardPercent", _maxRewardPercent)
}

// SetMaxRewardPercent is a paid mutator transaction binding the contract method 0x9cee9492.
//
// Solidity: function setMaxRewardPercent(uint256 _maxRewardPercent) returns()
func (_StrBtc *StrBtcSession) SetMaxRewardPercent(_maxRewardPercent *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetMaxRewardPercent(&_StrBtc.TransactOpts, _maxRewardPercent)
}

// SetMaxRewardPercent is a paid mutator transaction binding the contract method 0x9cee9492.
//
// Solidity: function setMaxRewardPercent(uint256 _maxRewardPercent) returns()
func (_StrBtc *StrBtcTransactorSession) SetMaxRewardPercent(_maxRewardPercent *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetMaxRewardPercent(&_StrBtc.TransactOpts, _maxRewardPercent)
}

// SetMinTimeBetweenRewards is a paid mutator transaction binding the contract method 0x25e0a33f.
//
// Solidity: function setMinTimeBetweenRewards(uint256 _minTimeBetweenRewards) returns()
func (_StrBtc *StrBtcTransactor) SetMinTimeBetweenRewards(opts *bind.TransactOpts, _minTimeBetweenRewards *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "setMinTimeBetweenRewards", _minTimeBetweenRewards)
}

// SetMinTimeBetweenRewards is a paid mutator transaction binding the contract method 0x25e0a33f.
//
// Solidity: function setMinTimeBetweenRewards(uint256 _minTimeBetweenRewards) returns()
func (_StrBtc *StrBtcSession) SetMinTimeBetweenRewards(_minTimeBetweenRewards *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetMinTimeBetweenRewards(&_StrBtc.TransactOpts, _minTimeBetweenRewards)
}

// SetMinTimeBetweenRewards is a paid mutator transaction binding the contract method 0x25e0a33f.
//
// Solidity: function setMinTimeBetweenRewards(uint256 _minTimeBetweenRewards) returns()
func (_StrBtc *StrBtcTransactorSession) SetMinTimeBetweenRewards(_minTimeBetweenRewards *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetMinTimeBetweenRewards(&_StrBtc.TransactOpts, _minTimeBetweenRewards)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x5abdb0dc.
//
// Solidity: function setMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StrBtc *StrBtcTransactor) SetMinWithdrawAmount(opts *bind.TransactOpts, _minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "setMinWithdrawAmount", _minWithdrawAmount)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x5abdb0dc.
//
// Solidity: function setMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StrBtc *StrBtcSession) SetMinWithdrawAmount(_minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetMinWithdrawAmount(&_StrBtc.TransactOpts, _minWithdrawAmount)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x5abdb0dc.
//
// Solidity: function setMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StrBtc *StrBtcTransactorSession) SetMinWithdrawAmount(_minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetMinWithdrawAmount(&_StrBtc.TransactOpts, _minWithdrawAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StrBtc *StrBtcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StrBtc *StrBtcSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.Transfer(&_StrBtc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StrBtc *StrBtcTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.Transfer(&_StrBtc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StrBtc *StrBtcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StrBtc *StrBtcSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.TransferFrom(&_StrBtc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StrBtc *StrBtcTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.TransferFrom(&_StrBtc.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrBtc *StrBtcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrBtc *StrBtcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.TransferOwnership(&_StrBtc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrBtc *StrBtcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.TransferOwnership(&_StrBtc.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StrBtc *StrBtcTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StrBtc *StrBtcSession) Unpause() (*types.Transaction, error) {
	return _StrBtc.Contract.Unpause(&_StrBtc.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StrBtc *StrBtcTransactorSession) Unpause() (*types.Transaction, error) {
	return _StrBtc.Contract.Unpause(&_StrBtc.TransactOpts)
}

// StrBtcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StrBtc contract.
type StrBtcApprovalIterator struct {
	Event *StrBtcApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcApproval represents a Approval event raised by the StrBtc contract.
type StrBtcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StrBtc *StrBtcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StrBtcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcApprovalIterator{contract: _StrBtc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StrBtc *StrBtcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StrBtcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcApproval)
				if err := _StrBtc.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StrBtc *StrBtcFilterer) ParseApproval(log types.Log) (*StrBtcApproval, error) {
	event := new(StrBtcApproval)
	if err := _StrBtc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcConverterBurnIterator is returned from FilterConverterBurn and is used to iterate over the raw logs and unpacked data for ConverterBurn events raised by the StrBtc contract.
type StrBtcConverterBurnIterator struct {
	Event *StrBtcConverterBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcConverterBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcConverterBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcConverterBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcConverterBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcConverterBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcConverterBurn represents a ConverterBurn event raised by the StrBtc contract.
type StrBtcConverterBurn struct {
	Converter common.Address
	From      common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConverterBurn is a free log retrieval operation binding the contract event 0x25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad567.
//
// Solidity: event ConverterBurn(address indexed converter, address indexed from, uint256 amount)
func (_StrBtc *StrBtcFilterer) FilterConverterBurn(opts *bind.FilterOpts, converter []common.Address, from []common.Address) (*StrBtcConverterBurnIterator, error) {

	var converterRule []interface{}
	for _, converterItem := range converter {
		converterRule = append(converterRule, converterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "ConverterBurn", converterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcConverterBurnIterator{contract: _StrBtc.contract, event: "ConverterBurn", logs: logs, sub: sub}, nil
}

// WatchConverterBurn is a free log subscription operation binding the contract event 0x25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad567.
//
// Solidity: event ConverterBurn(address indexed converter, address indexed from, uint256 amount)
func (_StrBtc *StrBtcFilterer) WatchConverterBurn(opts *bind.WatchOpts, sink chan<- *StrBtcConverterBurn, converter []common.Address, from []common.Address) (event.Subscription, error) {

	var converterRule []interface{}
	for _, converterItem := range converter {
		converterRule = append(converterRule, converterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "ConverterBurn", converterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcConverterBurn)
				if err := _StrBtc.contract.UnpackLog(event, "ConverterBurn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConverterBurn is a log parse operation binding the contract event 0x25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad567.
//
// Solidity: event ConverterBurn(address indexed converter, address indexed from, uint256 amount)
func (_StrBtc *StrBtcFilterer) ParseConverterBurn(log types.Log) (*StrBtcConverterBurn, error) {
	event := new(StrBtcConverterBurn)
	if err := _StrBtc.contract.UnpackLog(event, "ConverterBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcConverterMintIterator is returned from FilterConverterMint and is used to iterate over the raw logs and unpacked data for ConverterMint events raised by the StrBtc contract.
type StrBtcConverterMintIterator struct {
	Event *StrBtcConverterMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcConverterMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcConverterMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcConverterMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcConverterMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcConverterMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcConverterMint represents a ConverterMint event raised by the StrBtc contract.
type StrBtcConverterMint struct {
	Converter common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConverterMint is a free log retrieval operation binding the contract event 0x76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e.
//
// Solidity: event ConverterMint(address indexed converter, address indexed recipient, uint256 amount)
func (_StrBtc *StrBtcFilterer) FilterConverterMint(opts *bind.FilterOpts, converter []common.Address, recipient []common.Address) (*StrBtcConverterMintIterator, error) {

	var converterRule []interface{}
	for _, converterItem := range converter {
		converterRule = append(converterRule, converterItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "ConverterMint", converterRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcConverterMintIterator{contract: _StrBtc.contract, event: "ConverterMint", logs: logs, sub: sub}, nil
}

// WatchConverterMint is a free log subscription operation binding the contract event 0x76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e.
//
// Solidity: event ConverterMint(address indexed converter, address indexed recipient, uint256 amount)
func (_StrBtc *StrBtcFilterer) WatchConverterMint(opts *bind.WatchOpts, sink chan<- *StrBtcConverterMint, converter []common.Address, recipient []common.Address) (event.Subscription, error) {

	var converterRule []interface{}
	for _, converterItem := range converter {
		converterRule = append(converterRule, converterItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "ConverterMint", converterRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcConverterMint)
				if err := _StrBtc.contract.UnpackLog(event, "ConverterMint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConverterMint is a log parse operation binding the contract event 0x76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e.
//
// Solidity: event ConverterMint(address indexed converter, address indexed recipient, uint256 amount)
func (_StrBtc *StrBtcFilterer) ParseConverterMint(log types.Log) (*StrBtcConverterMint, error) {
	event := new(StrBtcConverterMint)
	if err := _StrBtc.contract.UnpackLog(event, "ConverterMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrBtc contract.
type StrBtcInitializedIterator struct {
	Event *StrBtcInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcInitialized represents a Initialized event raised by the StrBtc contract.
type StrBtcInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrBtc *StrBtcFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrBtcInitializedIterator, error) {

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrBtcInitializedIterator{contract: _StrBtc.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrBtc *StrBtcFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrBtcInitialized) (event.Subscription, error) {

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcInitialized)
				if err := _StrBtc.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrBtc *StrBtcFilterer) ParseInitialized(log types.Log) (*StrBtcInitialized, error) {
	event := new(StrBtcInitialized)
	if err := _StrBtc.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcMintBtcEventIterator is returned from FilterMintBtcEvent and is used to iterate over the raw logs and unpacked data for MintBtcEvent events raised by the StrBtc contract.
type StrBtcMintBtcEventIterator struct {
	Event *StrBtcMintBtcEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcMintBtcEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcMintBtcEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcMintBtcEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcMintBtcEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcMintBtcEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcMintBtcEvent represents a MintBtcEvent event raised by the StrBtc contract.
type StrBtcMintBtcEvent struct {
	To           common.Address
	Value        *big.Int
	BtcDepositId [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintBtcEvent is a free log retrieval operation binding the contract event 0xb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa469.
//
// Solidity: event MintBtcEvent(address indexed _to, uint256 _value, bytes32 _btcDepositId)
func (_StrBtc *StrBtcFilterer) FilterMintBtcEvent(opts *bind.FilterOpts, _to []common.Address) (*StrBtcMintBtcEventIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "MintBtcEvent", _toRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcMintBtcEventIterator{contract: _StrBtc.contract, event: "MintBtcEvent", logs: logs, sub: sub}, nil
}

// WatchMintBtcEvent is a free log subscription operation binding the contract event 0xb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa469.
//
// Solidity: event MintBtcEvent(address indexed _to, uint256 _value, bytes32 _btcDepositId)
func (_StrBtc *StrBtcFilterer) WatchMintBtcEvent(opts *bind.WatchOpts, sink chan<- *StrBtcMintBtcEvent, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "MintBtcEvent", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcMintBtcEvent)
				if err := _StrBtc.contract.UnpackLog(event, "MintBtcEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintBtcEvent is a log parse operation binding the contract event 0xb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa469.
//
// Solidity: event MintBtcEvent(address indexed _to, uint256 _value, bytes32 _btcDepositId)
func (_StrBtc *StrBtcFilterer) ParseMintBtcEvent(log types.Log) (*StrBtcMintBtcEvent, error) {
	event := new(StrBtcMintBtcEvent)
	if err := _StrBtc.contract.UnpackLog(event, "MintBtcEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StrBtc contract.
type StrBtcOwnershipTransferredIterator struct {
	Event *StrBtcOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcOwnershipTransferred represents a OwnershipTransferred event raised by the StrBtc contract.
type StrBtcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrBtc *StrBtcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StrBtcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcOwnershipTransferredIterator{contract: _StrBtc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrBtc *StrBtcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StrBtcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcOwnershipTransferred)
				if err := _StrBtc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrBtc *StrBtcFilterer) ParseOwnershipTransferred(log types.Log) (*StrBtcOwnershipTransferred, error) {
	event := new(StrBtcOwnershipTransferred)
	if err := _StrBtc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StrBtc contract.
type StrBtcPausedIterator struct {
	Event *StrBtcPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcPaused represents a Paused event raised by the StrBtc contract.
type StrBtcPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StrBtc *StrBtcFilterer) FilterPaused(opts *bind.FilterOpts) (*StrBtcPausedIterator, error) {

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StrBtcPausedIterator{contract: _StrBtc.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StrBtc *StrBtcFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StrBtcPaused) (event.Subscription, error) {

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcPaused)
				if err := _StrBtc.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StrBtc *StrBtcFilterer) ParsePaused(log types.Log) (*StrBtcPaused, error) {
	event := new(StrBtcPaused)
	if err := _StrBtc.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcRedeemBtcEventIterator is returned from FilterRedeemBtcEvent and is used to iterate over the raw logs and unpacked data for RedeemBtcEvent events raised by the StrBtc contract.
type StrBtcRedeemBtcEventIterator struct {
	Event *StrBtcRedeemBtcEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcRedeemBtcEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcRedeemBtcEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcRedeemBtcEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcRedeemBtcEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcRedeemBtcEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcRedeemBtcEvent represents a RedeemBtcEvent event raised by the StrBtc contract.
type StrBtcRedeemBtcEvent struct {
	From       common.Address
	BTCAddress string
	Value      *big.Int
	Id         *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedeemBtcEvent is a free log retrieval operation binding the contract event 0x83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e.
//
// Solidity: event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id)
func (_StrBtc *StrBtcFilterer) FilterRedeemBtcEvent(opts *bind.FilterOpts, _from []common.Address) (*StrBtcRedeemBtcEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "RedeemBtcEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcRedeemBtcEventIterator{contract: _StrBtc.contract, event: "RedeemBtcEvent", logs: logs, sub: sub}, nil
}

// WatchRedeemBtcEvent is a free log subscription operation binding the contract event 0x83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e.
//
// Solidity: event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id)
func (_StrBtc *StrBtcFilterer) WatchRedeemBtcEvent(opts *bind.WatchOpts, sink chan<- *StrBtcRedeemBtcEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "RedeemBtcEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcRedeemBtcEvent)
				if err := _StrBtc.contract.UnpackLog(event, "RedeemBtcEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemBtcEvent is a log parse operation binding the contract event 0x83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e.
//
// Solidity: event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id)
func (_StrBtc *StrBtcFilterer) ParseRedeemBtcEvent(log types.Log) (*StrBtcRedeemBtcEvent, error) {
	event := new(StrBtcRedeemBtcEvent)
	if err := _StrBtc.contract.UnpackLog(event, "RedeemBtcEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StrBtc contract.
type StrBtcRoleAdminChangedIterator struct {
	Event *StrBtcRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcRoleAdminChanged represents a RoleAdminChanged event raised by the StrBtc contract.
type StrBtcRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StrBtc *StrBtcFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StrBtcRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcRoleAdminChangedIterator{contract: _StrBtc.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StrBtc *StrBtcFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StrBtcRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcRoleAdminChanged)
				if err := _StrBtc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StrBtc *StrBtcFilterer) ParseRoleAdminChanged(log types.Log) (*StrBtcRoleAdminChanged, error) {
	event := new(StrBtcRoleAdminChanged)
	if err := _StrBtc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StrBtc contract.
type StrBtcRoleGrantedIterator struct {
	Event *StrBtcRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcRoleGranted represents a RoleGranted event raised by the StrBtc contract.
type StrBtcRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StrBtc *StrBtcFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StrBtcRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcRoleGrantedIterator{contract: _StrBtc.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StrBtc *StrBtcFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StrBtcRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcRoleGranted)
				if err := _StrBtc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StrBtc *StrBtcFilterer) ParseRoleGranted(log types.Log) (*StrBtcRoleGranted, error) {
	event := new(StrBtcRoleGranted)
	if err := _StrBtc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StrBtc contract.
type StrBtcRoleRevokedIterator struct {
	Event *StrBtcRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcRoleRevoked represents a RoleRevoked event raised by the StrBtc contract.
type StrBtcRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StrBtc *StrBtcFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StrBtcRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcRoleRevokedIterator{contract: _StrBtc.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StrBtc *StrBtcFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StrBtcRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcRoleRevoked)
				if err := _StrBtc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StrBtc *StrBtcFilterer) ParseRoleRevoked(log types.Log) (*StrBtcRoleRevoked, error) {
	event := new(StrBtcRoleRevoked)
	if err := _StrBtc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcTotalSupplyUpdatedEventIterator is returned from FilterTotalSupplyUpdatedEvent and is used to iterate over the raw logs and unpacked data for TotalSupplyUpdatedEvent events raised by the StrBtc contract.
type StrBtcTotalSupplyUpdatedEventIterator struct {
	Event *StrBtcTotalSupplyUpdatedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcTotalSupplyUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcTotalSupplyUpdatedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcTotalSupplyUpdatedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcTotalSupplyUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcTotalSupplyUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcTotalSupplyUpdatedEvent represents a TotalSupplyUpdatedEvent event raised by the StrBtc contract.
type StrBtcTotalSupplyUpdatedEvent struct {
	Nonce          *big.Int
	TotalBTCSupply *big.Int
	TotalShares    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalSupplyUpdatedEvent is a free log retrieval operation binding the contract event 0x339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b93440.
//
// Solidity: event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares)
func (_StrBtc *StrBtcFilterer) FilterTotalSupplyUpdatedEvent(opts *bind.FilterOpts) (*StrBtcTotalSupplyUpdatedEventIterator, error) {

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "TotalSupplyUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return &StrBtcTotalSupplyUpdatedEventIterator{contract: _StrBtc.contract, event: "TotalSupplyUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchTotalSupplyUpdatedEvent is a free log subscription operation binding the contract event 0x339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b93440.
//
// Solidity: event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares)
func (_StrBtc *StrBtcFilterer) WatchTotalSupplyUpdatedEvent(opts *bind.WatchOpts, sink chan<- *StrBtcTotalSupplyUpdatedEvent) (event.Subscription, error) {

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "TotalSupplyUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcTotalSupplyUpdatedEvent)
				if err := _StrBtc.contract.UnpackLog(event, "TotalSupplyUpdatedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTotalSupplyUpdatedEvent is a log parse operation binding the contract event 0x339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b93440.
//
// Solidity: event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares)
func (_StrBtc *StrBtcFilterer) ParseTotalSupplyUpdatedEvent(log types.Log) (*StrBtcTotalSupplyUpdatedEvent, error) {
	event := new(StrBtcTotalSupplyUpdatedEvent)
	if err := _StrBtc.contract.UnpackLog(event, "TotalSupplyUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StrBtc contract.
type StrBtcTransferIterator struct {
	Event *StrBtcTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcTransfer represents a Transfer event raised by the StrBtc contract.
type StrBtcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StrBtc *StrBtcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StrBtcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StrBtcTransferIterator{contract: _StrBtc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StrBtc *StrBtcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StrBtcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcTransfer)
				if err := _StrBtc.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StrBtc *StrBtcFilterer) ParseTransfer(log types.Log) (*StrBtcTransfer, error) {
	event := new(StrBtcTransfer)
	if err := _StrBtc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StrBtc contract.
type StrBtcUnpausedIterator struct {
	Event *StrBtcUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StrBtcUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StrBtcUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StrBtcUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcUnpaused represents a Unpaused event raised by the StrBtc contract.
type StrBtcUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StrBtc *StrBtcFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StrBtcUnpausedIterator, error) {

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StrBtcUnpausedIterator{contract: _StrBtc.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StrBtc *StrBtcFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StrBtcUnpaused) (event.Subscription, error) {

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcUnpaused)
				if err := _StrBtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StrBtc *StrBtcFilterer) ParseUnpaused(log types.Log) (*StrBtcUnpaused, error) {
	event := new(StrBtcUnpaused)
	if err := _StrBtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
