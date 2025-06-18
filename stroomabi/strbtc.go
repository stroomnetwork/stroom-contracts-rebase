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
	ABI: "[{\"type\":\"function\",\"name\":\"BTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONVERTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DUST_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_MINT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_UPDATE_TOTAL_SUPPLY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addConverter\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPauser\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"btcDepositIds\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"converterBurn\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"converterMint\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeInvoice\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"encodeTotalSupplyUpdate\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMintInvoiceHash\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPooledBTCByShares\",\"inputs\":[{\"name\":\"sharesAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShares\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSharesByPooledBTC\",\"inputs\":[{\"name\":\"btcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalSupplyUpdateHash\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_network\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"},{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"},{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRewardTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxRewardPercent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minTimeBetweenRewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintRewards\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"network\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"BTCAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeConverter\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePauser\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxRewardPercent\",\"inputs\":[{\"name\":\"_maxRewardPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinTimeBetweenRewards\",\"inputs\":[{\"name\":\"_minTimeBetweenRewards\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWithdrawAmount\",\"inputs\":[{\"name\":\"_minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidatorRegistry\",\"inputs\":[{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupplyUpdateNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConverterBurn\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConverterMint\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintBtcEvent\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_btcDepositId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedeemBtcEvent\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_BTCAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalSupplyUpdatedEvent\",\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalBTCSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AmountBelowMinWithdraw\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotBurnFromZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotMintToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DeltaIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBTCAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSharesOrPooledBTC\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSupplyNonce\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxRewardPercentTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinTimeBetweenRewardsTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinWithdrawTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAlreadyProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountTooBig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintToContractAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardTooBig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardTooFrequent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UpdateAlreadyProcessed\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b506123b58061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610372575f3560e01c80636739afca116101d4578063a7ce456511610109578063dd62ed3e116100a9578063e737557f11610079578063e737557f14610742578063f04da65b1461077b578063f376ebbb146107a3578063f8077fae146107cd575f5ffd5b8063dd62ed3e146106f5578063e07fbd0014610708578063e0a046101461071b578063e63ab1e91461072e575f5ffd5b8063b0390325116100e4578063b0390325146106b3578063c4d66de8146106bc578063c87b7a23146106cf578063d547741f146106e2575f5ffd5b8063a7ce45651461068e578063a9059cbb14610697578063af7f131f146106aa575f5ffd5b80638c659bf2116101745780639cee94921161014f5780639cee94921461064e5780639e76a007146106615780639fd9568714610674578063a217fddf14610687575f5ffd5b80638c659bf21461062057806391d148541461063357806395d89b4114610646575f5ffd5b806370a08231116101af57806370a08231146105df57806380239a04146105f257806382dc1ec4146106055780638456cb5914610618575f5ffd5b80636739afca146105995780636b2c0f55146105b95780636ce1c4dc146105cc575f5ffd5b8063313ce567116102aa578063457e1a491161024a5780635abdb0dc116102255780635abdb0dc1461052a5780635c975abb1461053d578063603c6a6714610554578063607375a114610586575f5ffd5b8063457e1a49146104fb578063497730501461050457806355573c7714610517575f5ffd5b806336568abe1161028557806336568abe146104cf5780633a98ef39146104e25780633f4ba83a146104ea57806342193473146104f2575f5ffd5b8063313ce567146104a45780633357325e146104b35780633395b75a146104c6575f5ffd5b806323b872dd1161031557806325e0a33f116102f057806325e0a33f146104605780632792949d146104735780632f2ff15d1461047e5780632f4cfd4714610491575f5ffd5b806323b872dd14610427578063248a9ca31461043a57806324b76fd51461044d575f5ffd5b8063095ea7b311610350578063095ea7b3146103d557806313776a8d146103e857806318160ddd1461040a57806321172f3b14610412575f5ffd5b806301ffc9a714610376578063026034f01461039e57806306fdde03146103c0575b5f5ffd5b610389610384366004611ca5565b6107d6565b60405190151581526020015b60405180910390f35b6103896103ac366004611ccc565b60096020525f908152604090205460ff1681565b6103c861080c565b6040516103959190611d11565b6103896103e3366004611d37565b6108cc565b6103fc5f5160206123605f395f51905f5281565b604051908152602001610395565b6004546103fc565b610425610420366004611da6565b6108e3565b005b610389610435366004611df5565b610b48565b6103fc610448366004611ccc565b610b6b565b61042561045b366004611e33565b610b8b565b61042561046e366004611ccc565b610cd0565b6103fc6305f5e10081565b61042561048c366004611e7b565b610d03565b61042561049f366004611d37565b610d25565b60405160088152602001610395565b6104256104c1366004611d37565b610dbb565b6103fc60065481565b6104256104dd366004611e7b565b610e48565b6005546103fc565b610425610e80565b6103fc61022281565b6103fc60015481565b610425610512366004611ea9565b610ea2565b6103fc610525366004611ccc565b610ef5565b610425610538366004611ccc565b610f3e565b5f5160206123405f395f51905f525460ff16610389565b6103c8604051806040016040528060138152602001725354524f4f4d5f4d494e545f494e564f49434560681b81525081565b610425610594366004611ec4565b610f71565b5f546105ac90600160a01b900460ff1681565b6040516103959190611f54565b6104256105c7366004611ea9565b611142565b6104256105da366004611ea9565b611167565b6103fc6105ed366004611ea9565b611188565b6103fc610600366004611f62565b6111c1565b610425610613366004611ea9565b611274565b610425611295565b6103c861062e366004611f62565b6112b4565b610389610641366004611e7b565b6112f9565b6103c861132f565b61042561065c366004611ccc565b61136d565b61042561066f366004611ea9565b61139f565b610425610682366004611f98565b6113c0565b6103fc5f81565b6103fc60025481565b6103896106a5366004611d37565b6113d9565b6103fc60035481565b6103fc60075481565b6104256106ca366004611ea9565b6113e6565b6103fc6106dd366004611ccc565b611407565b6104256106f0366004611e7b565b611446565b6103fc610703366004611fcf565b611462565b6103fc610716366004611ffb565b6114ab565b6103c8610729366004611ffb565b61154f565b6103fc5f5160206123005f395f51905f5281565b6103c86040518060400160405280601a8152602001795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b81525081565b6103fc610789366004611ea9565b6001600160a01b03165f908152600a602052604090205490565b5f546107b5906001600160a01b031681565b6040516001600160a01b039091168152602001610395565b6103fc60085481565b5f6001600160e01b03198216637965db0b60e01b148061080657506301ffc9a760e01b6001600160e01b03198316145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060915f5160206122e05f395f51905f529161084a90612015565b80601f016020809104026020016040519081016040528092919081815260200182805461087690612015565b80156108c15780601f10610898576101008083540402835291602001916108c1565b820191905f5260205f20905b8154815290600101906020018083116108a457829003601f168201915b505050505091505090565b5f336108d98185856115b2565b5060019392505050565b6108eb6115bf565b6040518060400160405280601a8152602001795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b81525061092885856112b4565b5f546040516311ee58a960e01b8152859185916001600160a01b03909116906311ee58a99061096190879087908790879060040161206f565b602060405180830381865afa15801561097c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109a091906120b3565b6109bd57604051636227817160e01b815260040160405180910390fd5b60035488146109df576040516336b118d760e11b815260040160405180910390fd5b600160035f8282546109f191906120e6565b90915550505f879003610a1757604051630aa7fcbb60e21b815260040160405180910390fd5b600754600854610a2791906120e6565b421015610a46576040516283cb6b60e51b815260040160405180910390fd5b5f612710600654600454610a5a91906120f9565b610a649190612110565b905080881115610a8757604051630779a2a360e31b815260040160405180910390fd5b5f610a928a8a6111c1565b5f8181526009602052604090205490915060ff1615610ac45760405163ce199d1760e01b815260040160405180910390fd5b5f818152600960205260408120805460ff19166001179055600480548b9290610aee9084906120e6565b909155505042600855600454600554604080518d815260208101939093528201527f339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b934409060600160405180910390a150505050505050505050565b5f33610b558582856115f1565b610b60858585611653565b506001949350505050565b5f9081525f5160206123205f395f51905f52602052604090206001015490565b610b936115bf565b600154831015610bb657604051633813eacd60e21b815260040160405180910390fd5b5f54600160a01b900460ff166003811115610bd357610bd3611f20565b604051632ce1e0f560e01b815273__$113774c24a411e022512c75ffd4c4add5e$__91632ce1e0f591610c0d91908690869060040161212f565b602060405180830381865af4158015610c28573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c4c91906120b3565b610c69576040516312d58f2760e21b815260040160405180910390fd5b610c7333846116b0565b600160025f828254610c8591906120e6565b909155505060025460405133917f83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e91610cc391869186918991612158565b60405180910390a2505050565b5f610cda816116e4565b61a8c0821015610cfd5760405163a7bbe41d60e01b815260040160405180910390fd5b50600755565b610d0c82610b6b565b610d15816116e4565b610d1f83836116ee565b50505050565b610d2d6115bf565b5f5160206123605f395f51905f52610d44816116e4565b6001600160a01b038316610d6b576040516389a4ea1960e01b815260040160405180910390fd5b610d75838361178f565b6040518281526001600160a01b0384169033907f76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e906020015b60405180910390a3505050565b610dc36115bf565b5f5160206123605f395f51905f52610dda816116e4565b6001600160a01b038316610e015760405163a30d2d8760e01b815260040160405180910390fd5b610e0b83836116b0565b6040518281526001600160a01b0384169033907f25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad56790602001610dae565b6001600160a01b0381163314610e715760405163334bd91960e11b815260040160405180910390fd5b610e7b82826117c3565b505050565b5f5160206123005f395f51905f52610e97816116e4565b610e9f61183c565b50565b5f610eac816116e4565b6001600160a01b038216610ed35760405163e6c4247b60e01b815260040160405180910390fd5b505f80546001600160a01b0319166001600160a01b0392909216919091179055565b5f6005545f1480610f065750600454155b15610f24576040516340a3daff60e11b815260040160405180910390fd5b600554600454610f3490846120f9565b6108069190612110565b5f610f48816116e4565b610222821015610f6b576040516343a87e9960e11b815260040160405180910390fd5b50600155565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f81158015610fb65750825b90505f8267ffffffffffffffff166001148015610fd25750303b155b905081158015610fe0575080155b15610ffe5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561102857845460ff60401b1916600160401b1785555b6110766040518060400160405280600e81526020016d29ba3937b7b6902134ba31b7b4b760911b8152506040518060400160405280600681526020016573747242544360d01b81525061189b565b61107e6118ad565b6110866118bd565b61108f886113e6565b611b586001555f80548a919060ff60a01b1916600160a01b8360038111156110b9576110b9611f20565b02179055506110c85f886116ee565b506110e05f5160206123005f395f51905f52876116ee565b506064600655620151806007555f600855831561113757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f61114c816116e4565b6111635f5160206123005f395f51905f5283611446565b5050565b5f611171816116e4565b6111635f5160206123605f395f51905f5283610d03565b5f6005545f0361119957505f919050565b6005546004546001600160a01b0384165f908152600a6020526040902054610f3491906120f9565b5f805460408051808201909152601a8152795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b60208201526001600160a01b039091169063bca0ac069061121186866112b4565b6040518363ffffffff1660e01b815260040161122e92919061217e565b602060405180830381865afa158015611249573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061126d91906121a2565b9392505050565b5f61127e816116e4565b6111635f5160206123005f395f51905f5283610d03565b5f5160206123005f395f51905f526112ac816116e4565b610e9f6118c5565b604080516020810184905290810182905230606090811b6bffffffffffffffffffffffff19168183015290607401604051602081830303815290604052905092915050565b5f9182525f5160206123205f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060915f5160206122e05f395f51905f529161084a90612015565b5f611377816116e4565b606482111561139957604051630737692960e31b815260040160405180910390fd5b50600655565b5f6113a9816116e4565b6111635f5160206123605f395f51905f5283611446565b6040516336b118d760e11b815260040160405180910390fd5b5f336108d9818585611653565b5f80546001600160a01b0319166001600160a01b0392909216919091179055565b5f6005545f14806114185750600454155b15611436576040516340a3daff60e11b815260040160405180910390fd5b600454600554610f3490846120f9565b61144f82610b6b565b611458816116e4565b610d1f83836117c3565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b5f80546040805180820190915260138152725354524f4f4d5f4d494e545f494e564f49434560681b60208201526001600160a01b039091169063bca0ac06906114f38561154f565b6040518363ffffffff1660e01b815260040161151092919061217e565b602060405180830381865afa15801561152b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080691906121a2565b60606115616040830160208401611ea9565b604080516bffffffffffffffffffffffff19606093841b81166020830152918501356034820152843560548201523090921b1660748201526088016040516020818303038152906040529050919050565b610e7b838383600161190d565b5f5160206123405f395f51905f525460ff16156115ef5760405163d93c066560e01b815260040160405180910390fd5b565b5f6115fc8484611462565b90505f198114610d1f578181101561164557604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b610d1f84848484035f61190d565b6001600160a01b03831661167c57604051634b637e8f60e11b81525f600482015260240161163c565b6001600160a01b0382166116a55760405163ec442f0560e01b81525f600482015260240161163c565b610e7b8383836119f1565b6001600160a01b0382166116d957604051634b637e8f60e11b81525f600482015260240161163c565b611163825f836119f1565b610e9f8133611b84565b5f5f5160206123205f395f51905f5261170784846112f9565b611786575f848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561173c3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610806565b5f915050610806565b6001600160a01b0382166117b85760405163ec442f0560e01b81525f600482015260240161163c565b6111635f83836119f1565b5f5f5160206123205f395f51905f526117dc84846112f9565b15611786575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610806565b611844611bbd565b5f5160206123405f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b6118a3611bec565b6111638282611c35565b6118b5611bec565b6115ef611c85565b6115ef611bec565b6118cd6115bf565b5f5160206123405f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361187d565b5f5160206122e05f395f51905f526001600160a01b0385166119445760405163e602df0560e01b81525f600482015260240161163c565b6001600160a01b03841661196d57604051634a1406b160e11b81525f600482015260240161163c565b6001600160a01b038086165f908152600183016020908152604080832093881683529290522083905581156119ea57836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516119e191815260200190565b60405180910390a35b5050505050565b6001600160a01b038316611a77575f6005545f1480611a105750600454155b611a2257611a1d82611407565b611a24565b815b90508160045f828254611a3791906120e6565b925050819055508060055f828254611a4f91906120e6565b90915550506001600160a01b0383165f908152600a6020526040902080549091019055611b3f565b5f611a8184611188565b905081811015611aa457604051631e9acf1760e31b815260040160405180910390fd5b5f611aae83611407565b90506001600160a01b038416611b0f578260045f828254611acf91906121b9565b925050819055508060055f828254611ae791906121b9565b90915550506001600160a01b0385165f908152600a6020526040902080548290039055611b3c565b6001600160a01b038086165f908152600a6020526040808220805485900390559186168152208054820190555b50505b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610dae91815260200190565b611b8e82826112f9565b6111635760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161163c565b5f5160206123405f395f51905f525460ff166115ef57604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166115ef57604051631afcd79f60e31b815260040160405180910390fd5b611c3d611bec565b5f5160206122e05f395f51905f527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03611c768482612224565b5060048101610d1f8382612224565b611c8d611bec565b5f5160206123405f395f51905f52805460ff19169055565b5f60208284031215611cb5575f5ffd5b81356001600160e01b03198116811461126d575f5ffd5b5f60208284031215611cdc575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61126d6020830184611ce3565b6001600160a01b0381168114610e9f575f5ffd5b5f5f60408385031215611d48575f5ffd5b8235611d5381611d23565b946020939093013593505050565b5f5f83601f840112611d71575f5ffd5b50813567ffffffffffffffff811115611d88575f5ffd5b602083019150836020828501011115611d9f575f5ffd5b9250929050565b5f5f5f5f60608587031215611db9575f5ffd5b8435935060208501359250604085013567ffffffffffffffff811115611ddd575f5ffd5b611de987828801611d61565b95989497509550505050565b5f5f5f60608486031215611e07575f5ffd5b8335611e1281611d23565b92506020840135611e2281611d23565b929592945050506040919091013590565b5f5f5f60408486031215611e45575f5ffd5b83359250602084013567ffffffffffffffff811115611e62575f5ffd5b611e6e86828701611d61565b9497909650939450505050565b5f5f60408385031215611e8c575f5ffd5b823591506020830135611e9e81611d23565b809150509250929050565b5f60208284031215611eb9575f5ffd5b813561126d81611d23565b5f5f5f5f60808587031215611ed7575f5ffd5b843560048110611ee5575f5ffd5b93506020850135611ef581611d23565b92506040850135611f0581611d23565b91506060850135611f1581611d23565b939692955090935050565b634e487b7160e01b5f52602160045260245ffd5b60048110611f5057634e487b7160e01b5f52602160045260245ffd5b9052565b602081016108068284611f34565b5f5f60408385031215611f73575f5ffd5b50508035926020909101359150565b5f60608284031215611f92575f5ffd5b50919050565b5f5f5f60808486031215611faa575f5ffd5b611fb48585611f82565b9250606084013567ffffffffffffffff811115611e62575f5ffd5b5f5f60408385031215611fe0575f5ffd5b8235611feb81611d23565b91506020830135611e9e81611d23565b5f6060828403121561200b575f5ffd5b61126d8383611f82565b600181811c9082168061202957607f821691505b602082108103611f9257634e487b7160e01b5f52602260045260245ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606081525f6120816060830187611ce3565b82810360208401526120938187611ce3565b905082810360408401526120a8818587612047565b979650505050505050565b5f602082840312156120c3575f5ffd5b8151801515811461126d575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115610806576108066120d2565b8082028115828204841417610806576108066120d2565b5f8261212a57634e487b7160e01b5f52601260045260245ffd5b500490565b6121398185611f34565b604060208201525f61214f604083018486612047565b95945050505050565b606081525f61216b606083018688612047565b6020830194909452506040015292915050565b604081525f6121906040830185611ce3565b828103602084015261214f8185611ce3565b5f602082840312156121b2575f5ffd5b5051919050565b81810381811115610806576108066120d2565b634e487b7160e01b5f52604160045260245ffd5b601f821115610e7b57805f5260205f20601f840160051c810160208510156122055750805b601f840160051c820191505b818110156119ea575f8155600101612211565b815167ffffffffffffffff81111561223e5761223e6121cc565b6122528161224c8454612015565b846121e0565b6020601f821160018114612284575f831561226d5750848201515b5f19600385901b1c1916600184901b1784556119ea565b5f84815260208120601f198516915b828110156122b35787850151825560209485019460019092019101612293565b50848210156122d057868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0065d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033001cf336fddcc7dc48127faf7a5b80ee54fce73ef647eecd31c24bb6cce3ac3eefa26469706673582212204b2801669b0edab9b6561197792de9dccf34133f0c3bad52cb36c50723ab378e64736f6c634300081b0033",
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

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StrBtc *StrBtcCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StrBtc *StrBtcSession) PAUSERROLE() ([32]byte, error) {
	return _StrBtc.Contract.PAUSERROLE(&_StrBtc.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StrBtc *StrBtcCallerSession) PAUSERROLE() ([32]byte, error) {
	return _StrBtc.Contract.PAUSERROLE(&_StrBtc.CallOpts)
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

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address pauser) returns()
func (_StrBtc *StrBtcTransactor) AddPauser(opts *bind.TransactOpts, pauser common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "addPauser", pauser)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address pauser) returns()
func (_StrBtc *StrBtcSession) AddPauser(pauser common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.AddPauser(&_StrBtc.TransactOpts, pauser)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address pauser) returns()
func (_StrBtc *StrBtcTransactorSession) AddPauser(pauser common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.AddPauser(&_StrBtc.TransactOpts, pauser)
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

// Initialize is a paid mutator transaction binding the contract method 0x607375a1.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry, address _admin, address _pauser) returns()
func (_StrBtc *StrBtcTransactor) Initialize(opts *bind.TransactOpts, _network uint8, _validatorRegistry common.Address, _admin common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "initialize", _network, _validatorRegistry, _admin, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0x607375a1.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry, address _admin, address _pauser) returns()
func (_StrBtc *StrBtcSession) Initialize(_network uint8, _validatorRegistry common.Address, _admin common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.Initialize(&_StrBtc.TransactOpts, _network, _validatorRegistry, _admin, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0x607375a1.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry, address _admin, address _pauser) returns()
func (_StrBtc *StrBtcTransactorSession) Initialize(_network uint8, _validatorRegistry common.Address, _admin common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.Initialize(&_StrBtc.TransactOpts, _network, _validatorRegistry, _admin, _pauser)
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

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address pauser) returns()
func (_StrBtc *StrBtcTransactor) RemovePauser(opts *bind.TransactOpts, pauser common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "removePauser", pauser)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address pauser) returns()
func (_StrBtc *StrBtcSession) RemovePauser(pauser common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.RemovePauser(&_StrBtc.TransactOpts, pauser)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address pauser) returns()
func (_StrBtc *StrBtcTransactorSession) RemovePauser(pauser common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.RemovePauser(&_StrBtc.TransactOpts, pauser)
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

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_StrBtc *StrBtcTransactor) SetValidatorRegistry(opts *bind.TransactOpts, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "setValidatorRegistry", _validatorRegistry)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_StrBtc *StrBtcSession) SetValidatorRegistry(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.SetValidatorRegistry(&_StrBtc.TransactOpts, _validatorRegistry)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_StrBtc *StrBtcTransactorSession) SetValidatorRegistry(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc.Contract.SetValidatorRegistry(&_StrBtc.TransactOpts, _validatorRegistry)
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
