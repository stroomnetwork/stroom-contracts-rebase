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
	ABI: "[{\"type\":\"function\",\"name\":\"BTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONVERTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DUST_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_MINT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_UPDATE_TOTAL_SUPPLY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addConverter\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPauser\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"btcDepositIds\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"converterBurn\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"converterMint\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dailyGlobalRedeemLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dailyRedeemLimitPerAccount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeInvoice\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"encodeTotalSupplyUpdate\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMintInvoiceHash\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPooledBTCByShares\",\"inputs\":[{\"name\":\"sharesAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemainingAccountRedeemLimit\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemainingGlobalRedeemLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShares\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSharesByPooledBTC\",\"inputs\":[{\"name\":\"btcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalSupplyUpdateHash\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUsedAccountRedeemAmount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUsedGlobalRedeemAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_network\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"},{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"},{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRewardTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxRewardPercent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minTimeBetweenRewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintRewards\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"network\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"BTCAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeemLimitsEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reinitializeV2\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reinitializeV3\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeConverter\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePauser\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDailyGlobalRedeemLimit\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDailyRedeemLimitPerAccount\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxRewardPercent\",\"inputs\":[{\"name\":\"_maxRewardPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinTimeBetweenRewards\",\"inputs\":[{\"name\":\"_minTimeBetweenRewards\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWithdrawAmount\",\"inputs\":[{\"name\":\"_minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRedeemLimitsEnabled\",\"inputs\":[{\"name\":\"_enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidatorRegistry\",\"inputs\":[{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupplyUpdateNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConverterBurn\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConverterMint\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DailyGlobalRedeemLimitUpdated\",\"inputs\":[{\"name\":\"newLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DailyRedeemLimitPerAccountUpdated\",\"inputs\":[{\"name\":\"newLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintBtcEvent\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_btcDepositId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedeemBtcEvent\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_BTCAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedeemLimitsEnabledUpdated\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalSupplyUpdatedEvent\",\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalBTCSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AmountBelowMinWithdraw\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotBurnFromZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotMintToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DailyRedeemLimitExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DeltaIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalDailyRedeemLimitExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBTCAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSharesOrPooledBTC\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSupplyNonce\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxRewardPercentTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinTimeBetweenRewardsTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinWithdrawTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAlreadyProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountTooBig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintToContractAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardTooBig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardTooFrequent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UpdateAlreadyProcessed\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50612b678061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106103eb575f3560e01c80636739afca1161020b578063b03903251161011f578063ddd0c0b0116100b4578063e737557f11610084578063e737557f14610839578063ef3e008414610872578063f04da65b1461087f578063f376ebbb146108a7578063f8077fae146108d1575f5ffd5b8063ddd0c0b0146107f6578063e07fbd00146107ff578063e0a0461014610812578063e63ab1e914610825575f5ffd5b8063c87b7a23116100ef578063c87b7a23146107b4578063cd897c71146107c7578063d547741f146107d0578063dd62ed3e146107e3575f5ffd5b8063b039032514610788578063bac22bb814610791578063bdf446de14610799578063c4115874146107ac575f5ffd5b806391d14854116101a05780639fd95687116101705780639fd9568714610749578063a217fddf1461075c578063a7ce456514610763578063a9059cbb1461076c578063af7f131f1461077f575f5ffd5b806391d148541461070857806395d89b411461071b5780639cee9492146107235780639e76a00714610736575f5ffd5b806380239a04116101db57806380239a04146106c757806382dc1ec4146106da5780638456cb59146106ed5780638c659bf2146106f5575f5ffd5b80636739afca1461066e5780636b2c0f551461068e5780636ce1c4dc146106a157806370a08231146106b4575f5ffd5b8063313ce567116103025780634977305011610297578063587234dd11610267578063587234dd146105ec5780635abdb0dc146105ff5780635c975abb14610612578063603c6a6714610629578063607375a11461065b575f5ffd5b806349773050146105ab5780634c5bb881146105be57806355573c77146105d157806357e90cab146105e4575f5ffd5b80633a98ef39116102d25780633a98ef39146105895780633f4ba83a146105915780634219347314610599578063457e1a49146105a2575f5ffd5b8063313ce5671461054b5780633357325e1461055a5780633395b75a1461056d57806336568abe14610576575f5ffd5b806321172f3b1161038357806325e0a33f1161035357806325e0a33f146104ff5780632792949d14610512578063279b91cb1461051d5780632f2ff15d146105255780632f4cfd4714610538575f5ffd5b806321172f3b146104b357806323b872dd146104c6578063248a9ca3146104d957806324b76fd5146104ec575f5ffd5b806313776a8d116103be57806313776a8d1461046157806318160ddd146104835780631a7b3f921461048b5780631e57686e146104a0575f5ffd5b806301ffc9a7146103ef578063026034f01461041757806306fdde0314610439578063095ea7b31461044e575b5f5ffd5b6104026103fd366004612413565b6108da565b60405190151581526020015b60405180910390f35b61040261042536600461243a565b60096020525f908152604090205460ff1681565b610441610910565b60405161040e919061247f565b61040261045c3660046124a5565b6109d0565b6104755f516020612b125f395f51905f5281565b60405190815260200161040e565b600454610475565b61049e61049936600461243a565b6109e7565b005b61049e6104ae3660046124dc565b610a2e565b61049e6104c136600461253c565b610a79565b6104026104d436600461258b565b610cde565b6104756104e736600461243a565b610d01565b61049e6104fa3660046125c9565b610d21565b61049e61050d36600461243a565b610e7b565b6104756305f5e10081565b610475610eae565b61049e610533366004612611565b610f0d565b61049e6105463660046124a5565b610f2f565b6040516008815260200161040e565b61049e6105683660046124a5565b610fc5565b61047560065481565b61049e610584366004612611565b611052565b600554610475565b61049e61108a565b61047561022281565b61047560015481565b61049e6105b936600461263f565b6110ac565b6104756105cc36600461263f565b6110ff565b6104756105df36600461243a565b611146565b61047561118f565b61049e6105fa36600461243a565b6111bc565b61049e61060d36600461243a565b6111fb565b5f516020612ad25f395f51905f525460ff16610402565b610441604051806040016040528060138152602001725354524f4f4d5f4d494e545f494e564f49434560681b81525081565b61049e61066936600461265a565b61122e565b5f5461068190600160a01b900460ff1681565b60405161040e91906126ea565b61049e61069c36600461263f565b6113ec565b61049e6106af36600461263f565b611411565b6104756106c236600461263f565b611432565b6104756106d53660046126f8565b61146b565b61049e6106e836600461263f565b61151e565b61049e61153f565b6104416107033660046126f8565b61155e565b610402610716366004612611565b6115a3565b6104416115d9565b61049e61073136600461243a565b611617565b61049e61074436600461263f565b611649565b61049e61075736600461272e565b61166a565b6104755f81565b61047560025481565b61040261077a3660046124a5565b611763565b61047560035481565b61047560075481565b61049e611770565b6104756107a736600461263f565b61183d565b61049e6118bb565b6104756107c236600461243a565b61196b565b610475600b5481565b61049e6107de366004612611565b6119aa565b6104756107f1366004612765565b6119c6565b610475600c5481565b61047561080d366004612791565b611a0f565b610441610820366004612791565b611ab3565b6104755f516020612a925f395f51905f5281565b6104416040518060400160405280601a8152602001795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b81525081565b600d546104029060ff1681565b61047561088d36600461263f565b6001600160a01b03165f908152600a602052604090205490565b5f546108b9906001600160a01b031681565b6040516001600160a01b03909116815260200161040e565b61047560085481565b5f6001600160e01b03198216637965db0b60e01b148061090a57506301ffc9a760e01b6001600160e01b03198316145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060915f516020612a725f395f51905f529161094e906127ab565b80601f016020809104026020016040519081016040528092919081815260200182805461097a906127ab565b80156109c55780601f1061099c576101008083540402835291602001916109c5565b820191905f5260205f20905b8154815290600101906020018083116109a857829003601f168201915b505050505091505090565b5f336109dd818585611b16565b5060019392505050565b5f6109f181611b23565b600b8290556040518281527fdd7c282f04e637ea678511c395b45078fae3d4226e171c4af2307546f2f3be43906020015b60405180910390a15050565b5f610a3881611b23565b600d805460ff19168315159081179091556040519081527fe55d1ef19fc4263b518e1eecbe36f3cfd4f3192c442fc2b28122e779727f97f190602001610a22565b610a81611b2d565b6040518060400160405280601a8152602001795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b815250610abe858561155e565b5f546040516311ee58a960e01b8152859185916001600160a01b03909116906311ee58a990610af7908790879087908790600401612805565b602060405180830381865afa158015610b12573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b369190612849565b610b5357604051636227817160e01b815260040160405180910390fd5b6003548814610b75576040516336b118d760e11b815260040160405180910390fd5b600160035f828254610b879190612878565b90915550505f879003610bad57604051630aa7fcbb60e21b815260040160405180910390fd5b600754600854610bbd9190612878565b421015610bdc576040516283cb6b60e51b815260040160405180910390fd5b5f612710600654600454610bf0919061288b565b610bfa91906128a2565b905080881115610c1d57604051630779a2a360e31b815260040160405180910390fd5b5f610c288a8a61146b565b5f8181526009602052604090205490915060ff1615610c5a5760405163ce199d1760e01b815260040160405180910390fd5b5f818152600960205260408120805460ff19166001179055600480548b9290610c84908490612878565b909155505042600855600454600554604080518d815260208101939093528201527f339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b934409060600160405180910390a150505050505050505050565b5f33610ceb858285611b5f565b610cf6858585611bc1565b506001949350505050565b5f9081525f516020612ab25f395f51905f52602052604090206001015490565b610d29611b2d565b600154831015610d4c57604051633813eacd60e21b815260040160405180910390fd5b5f54600160a01b900460ff166003811115610d6957610d696126b6565b604051632ce1e0f560e01b815273__$113774c24a411e022512c75ffd4c4add5e$__91632ce1e0f591610da39190869086906004016128c1565b602060405180830381865af4158015610dbe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610de29190612849565b610dff576040516312d58f2760e21b815260040160405180910390fd5b600d5460ff1615610e1457610e143384611c1e565b610e1e3384611d06565b600160025f828254610e309190612878565b909155505060025460405133917f83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e91610e6e918691869189916128ea565b60405180910390a2505050565b5f610e8581611b23565b61a8c0821015610ea85760405163a7bbe41d60e01b815260040160405180910390fd5b50600755565b600d545f9060ff16610ec057505f1990565b5f610ece62015180426128a2565b905080600e600101541015610ee5575050600c5490565b600c54600e5410610ef7575f91505090565b600e54600c54610f079190612910565b91505090565b610f1682610d01565b610f1f81611b23565b610f298383611d3a565b50505050565b610f37611b2d565b5f516020612b125f395f51905f52610f4e81611b23565b6001600160a01b038316610f75576040516389a4ea1960e01b815260040160405180910390fd5b610f7f8383611ddb565b6040518281526001600160a01b0384169033907f76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e906020015b60405180910390a3505050565b610fcd611b2d565b5f516020612b125f395f51905f52610fe481611b23565b6001600160a01b03831661100b5760405163a30d2d8760e01b815260040160405180910390fd5b6110158383611d06565b6040518281526001600160a01b0384169033907f25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad56790602001610fb8565b6001600160a01b038116331461107b5760405163334bd91960e11b815260040160405180910390fd5b6110858282611e0f565b505050565b5f516020612a925f395f51905f526110a181611b23565b6110a9611e88565b50565b5f6110b681611b23565b6001600160a01b0382166110dd5760405163e6c4247b60e01b815260040160405180910390fd5b505f80546001600160a01b0319166001600160a01b0392909216919091179055565b5f8061110e62015180426128a2565b6001600160a01b0384165f90815260106020526040902060018101549192509082111561113e57505f9392505050565b549392505050565b5f6005545f14806111575750600454155b15611175576040516340a3daff60e11b815260040160405180910390fd5b600554600454611185908461288b565b61090a91906128a2565b5f8061119e62015180426128a2565b905080600e6001015410156111b4575f91505090565b5050600e5490565b5f6111c681611b23565b600c8290556040518281527f087142b64e709fbcc7d1f7097cf49027c4c377dfc0c9f78d9a6cf2d2c627403490602001610a22565b5f61120581611b23565b610222821015611228576040516343a87e9960e11b815260040160405180910390fd5b50600155565b5f516020612af25f395f51905f528054600160401b810460ff16159067ffffffffffffffff165f811580156112605750825b90505f8267ffffffffffffffff16600114801561127c5750303b155b90508115801561128a575080155b156112a85760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156112d257845460ff60401b1916600160401b1785555b6113206040518060400160405280600e81526020016d29ba3937b7b6902134ba31b7b4b760911b8152506040518060400160405280600681526020016573747242544360d01b815250611ee7565b611328611ef9565b611330611f09565b61133988611f11565b611b586001555f80548a919060ff60a01b1916600160a01b836003811115611363576113636126b6565b02179055506113725f88611d3a565b5061138a5f516020612a925f395f51905f5287611d3a565b506064600655620151806007555f60085583156113e157845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f6113f681611b23565b61140d5f516020612a925f395f51905f52836119aa565b5050565b5f61141b81611b23565b61140d5f516020612b125f395f51905f5283610f0d565b5f6005545f0361144357505f919050565b6005546004546001600160a01b0384165f908152600a6020526040902054611185919061288b565b5f805460408051808201909152601a8152795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b60208201526001600160a01b039091169063bca0ac06906114bb868661155e565b6040518363ffffffff1660e01b81526004016114d8929190612923565b602060405180830381865afa1580156114f3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115179190612947565b9392505050565b5f61152881611b23565b61140d5f516020612a925f395f51905f5283610f0d565b5f516020612a925f395f51905f5261155681611b23565b6110a9611f3a565b604080516020810184905290810182905230606090811b6bffffffffffffffffffffffff19168183015290607401604051602081830303815290604052905092915050565b5f9182525f516020612ab25f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060915f516020612a725f395f51905f529161094e906127ab565b5f61162181611b23565b606482111561164357604051630737692960e31b815260040160405180910390fd5b50600655565b5f61165381611b23565b61140d5f516020612b125f395f51905f52836119aa565b611672611b2d565b604051806040016040528060138152602001725354524f4f4d5f4d494e545f494e564f49434560681b8152506116a784611ab3565b5f546040516311ee58a960e01b8152859185916001600160a01b03909116906311ee58a9906116e0908790879087908790600401612805565b602060405180830381865afa1580156116fb573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061171f9190612849565b61173c57604051636227817160e01b815260040160405180910390fd5b61175a604088018035906117539060208b0161263f565b8935611f82565b50505050505050565b5f336109dd818585611bc1565b5f516020612af25f395f51905f52805460039190600160401b900460ff16806117a75750805467ffffffffffffffff808416911610155b156117c55760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff191667ffffffffffffffff8316908117600160401b178255620f4240600b556305f5e100600c55600d805460ff19169055815460ff60401b191682556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602001610a22565b600d545f9060ff1661185157505f19919050565b5f61185f62015180426128a2565b6001600160a01b0384165f908152601060205260409020600181015491925090821115611891575050600b5492915050565b600b548154106118a457505f9392505050565b8054600b546118b39190612910565b949350505050565b5f516020612af25f395f51905f52805460029190600160401b900460ff16806118f25750805467ffffffffffffffff808416911610155b156119105760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff191667ffffffffffffffff8316908117600160401b1760ff60401b191682556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602001610a22565b5f6005545f148061197c5750600454155b1561199a576040516340a3daff60e11b815260040160405180910390fd5b600454600554611185908461288b565b6119b382610d01565b6119bc81611b23565b610f298383611e0f565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b5f80546040805180820190915260138152725354524f4f4d5f4d494e545f494e564f49434560681b60208201526001600160a01b039091169063bca0ac0690611a5785611ab3565b6040518363ffffffff1660e01b8152600401611a74929190612923565b602060405180830381865afa158015611a8f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061090a9190612947565b6060611ac5604083016020840161263f565b604080516bffffffffffffffffffffffff19606093841b81166020830152918501356034820152843560548201523090921b1660748201526088016040516020818303038152906040529050919050565b611085838383600161208e565b6110a98133612172565b5f516020612ad25f395f51905f525460ff1615611b5d5760405163d93c066560e01b815260040160405180910390fd5b565b5f611b6a84846119c6565b90505f198114610f295781811015611bb357604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b610f2984848484035f61208e565b6001600160a01b038316611bea57604051634b637e8f60e11b81525f6004820152602401611baa565b6001600160a01b038216611c135760405163ec442f0560e01b81525f6004820152602401611baa565b6110858383836121ab565b5f611c2c62015180426128a2565b6001600160a01b0384165f908152601060205260409020600181015491925090821115611c5e575f8155600181018290555b600b548154611c6e908590612878565b1115611c8d5760405163e55017a760e01b815260040160405180910390fd5b82815f015f828254611c9f9190612878565b9091555050600f54821115611cb8575f600e55600f8290555b600c54600e54611cc9908590612878565b1115611ce85760405163c998883360e01b815260040160405180910390fd5b82600e5f015f828254611cfb9190612878565b909155505050505050565b6001600160a01b038216611d2f57604051634b637e8f60e11b81525f6004820152602401611baa565b61140d825f836121ab565b5f5f516020612ab25f395f51905f52611d5384846115a3565b611dd2575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055611d883390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061090a565b5f91505061090a565b6001600160a01b038216611e045760405163ec442f0560e01b81525f6004820152602401611baa565b61140d5f83836121ab565b5f5f516020612ab25f395f51905f52611e2884846115a3565b15611dd2575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061090a565b611e9061233e565b5f516020612ad25f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b611eef61236d565b61140d82826123a3565b611f0161236d565b611b5d6123f3565b611b5d61236d565b611f1961236d565b5f80546001600160a01b0319166001600160a01b0392909216919091179055565b611f42611b2d565b5f516020612ad25f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611ec9565b825f03611fa257604051631cebf66f60e11b815260040160405180910390fd5b611fb46305f5e1006301406f4061288b565b8310611fd357604051636006659960e11b815260040160405180910390fd5b306001600160a01b03831603611ffc5760405163603acaa960e11b815260040160405180910390fd5b5f8181526009602052604090205460ff161561202b57604051637bdb87d160e01b815260040160405180910390fd5b5f818152600960205260409020805460ff1916600117905561204d8284611ddb565b60408051848152602081018390526001600160a01b038416917fb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa4699101610e6e565b5f516020612a725f395f51905f526001600160a01b0385166120c55760405163e602df0560e01b81525f6004820152602401611baa565b6001600160a01b0384166120ee57604051634a1406b160e11b81525f6004820152602401611baa565b6001600160a01b038086165f9081526001830160209081526040808320938816835292905220839055811561216b57836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161216291815260200190565b60405180910390a35b5050505050565b61217c82826115a3565b61140d5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401611baa565b6001600160a01b038316612231575f6005545f14806121ca5750600454155b6121dc576121d78261196b565b6121de565b815b90508160045f8282546121f19190612878565b925050819055508060055f8282546122099190612878565b90915550506001600160a01b0383165f908152600a60205260409020805490910190556122f9565b5f61223b84611432565b90508181101561225e57604051631e9acf1760e31b815260040160405180910390fd5b5f6122688361196b565b90506001600160a01b0384166122c9578260045f8282546122899190612910565b925050819055508060055f8282546122a19190612910565b90915550506001600160a01b0385165f908152600a60205260409020805482900390556122f6565b6001600160a01b038086165f908152600a6020526040808220805485900390559186168152208054820190555b50505b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610fb891815260200190565b5f516020612ad25f395f51905f525460ff16611b5d57604051638dfc202b60e01b815260040160405180910390fd5b5f516020612af25f395f51905f5254600160401b900460ff16611b5d57604051631afcd79f60e31b815260040160405180910390fd5b6123ab61236d565b5f516020612a725f395f51905f527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036123e484826129b6565b5060048101610f2983826129b6565b6123fb61236d565b5f516020612ad25f395f51905f52805460ff19169055565b5f60208284031215612423575f5ffd5b81356001600160e01b031981168114611517575f5ffd5b5f6020828403121561244a575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6115176020830184612451565b6001600160a01b03811681146110a9575f5ffd5b5f5f604083850312156124b6575f5ffd5b82356124c181612491565b946020939093013593505050565b80151581146110a9575f5ffd5b5f602082840312156124ec575f5ffd5b8135611517816124cf565b5f5f83601f840112612507575f5ffd5b50813567ffffffffffffffff81111561251e575f5ffd5b602083019150836020828501011115612535575f5ffd5b9250929050565b5f5f5f5f6060858703121561254f575f5ffd5b8435935060208501359250604085013567ffffffffffffffff811115612573575f5ffd5b61257f878288016124f7565b95989497509550505050565b5f5f5f6060848603121561259d575f5ffd5b83356125a881612491565b925060208401356125b881612491565b929592945050506040919091013590565b5f5f5f604084860312156125db575f5ffd5b83359250602084013567ffffffffffffffff8111156125f8575f5ffd5b612604868287016124f7565b9497909650939450505050565b5f5f60408385031215612622575f5ffd5b82359150602083013561263481612491565b809150509250929050565b5f6020828403121561264f575f5ffd5b813561151781612491565b5f5f5f5f6080858703121561266d575f5ffd5b84356004811061267b575f5ffd5b9350602085013561268b81612491565b9250604085013561269b81612491565b915060608501356126ab81612491565b939692955090935050565b634e487b7160e01b5f52602160045260245ffd5b600481106126e657634e487b7160e01b5f52602160045260245ffd5b9052565b6020810161090a82846126ca565b5f5f60408385031215612709575f5ffd5b50508035926020909101359150565b5f60608284031215612728575f5ffd5b50919050565b5f5f5f60808486031215612740575f5ffd5b61274a8585612718565b9250606084013567ffffffffffffffff8111156125f8575f5ffd5b5f5f60408385031215612776575f5ffd5b823561278181612491565b9150602083013561263481612491565b5f606082840312156127a1575f5ffd5b6115178383612718565b600181811c908216806127bf57607f821691505b60208210810361272857634e487b7160e01b5f52602260045260245ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606081525f6128176060830187612451565b82810360208401526128298187612451565b9050828103604084015261283e8185876127dd565b979650505050505050565b5f60208284031215612859575f5ffd5b8151611517816124cf565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561090a5761090a612864565b808202811582820484141761090a5761090a612864565b5f826128bc57634e487b7160e01b5f52601260045260245ffd5b500490565b6128cb81856126ca565b604060208201525f6128e16040830184866127dd565b95945050505050565b606081525f6128fd6060830186886127dd565b6020830194909452506040015292915050565b8181038181111561090a5761090a612864565b604081525f6129356040830185612451565b82810360208401526128e18185612451565b5f60208284031215612957575f5ffd5b5051919050565b634e487b7160e01b5f52604160045260245ffd5b601f82111561108557805f5260205f20601f840160051c810160208510156129975750805b601f840160051c820191505b8181101561216b575f81556001016129a3565b815167ffffffffffffffff8111156129d0576129d061295e565b6129e4816129de84546127ab565b84612972565b6020601f821160018114612a16575f83156129ff5750848201515b5f19600385901b1c1916600184901b17845561216b565b5f84815260208120601f198516915b82811015612a455787850151825560209485019460019092019101612a25565b5084821015612a6257868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0065d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a001cf336fddcc7dc48127faf7a5b80ee54fce73ef647eecd31c24bb6cce3ac3eefa2646970667358221220d68af52e75874b5eae4dbea400458b53a5145575544b5b7ee499d5133087aada64736f6c634300081b0033",
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

// DailyGlobalRedeemLimit is a free data retrieval call binding the contract method 0xddd0c0b0.
//
// Solidity: function dailyGlobalRedeemLimit() view returns(uint256)
func (_StrBtc *StrBtcCaller) DailyGlobalRedeemLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "dailyGlobalRedeemLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyGlobalRedeemLimit is a free data retrieval call binding the contract method 0xddd0c0b0.
//
// Solidity: function dailyGlobalRedeemLimit() view returns(uint256)
func (_StrBtc *StrBtcSession) DailyGlobalRedeemLimit() (*big.Int, error) {
	return _StrBtc.Contract.DailyGlobalRedeemLimit(&_StrBtc.CallOpts)
}

// DailyGlobalRedeemLimit is a free data retrieval call binding the contract method 0xddd0c0b0.
//
// Solidity: function dailyGlobalRedeemLimit() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) DailyGlobalRedeemLimit() (*big.Int, error) {
	return _StrBtc.Contract.DailyGlobalRedeemLimit(&_StrBtc.CallOpts)
}

// DailyRedeemLimitPerAccount is a free data retrieval call binding the contract method 0xcd897c71.
//
// Solidity: function dailyRedeemLimitPerAccount() view returns(uint256)
func (_StrBtc *StrBtcCaller) DailyRedeemLimitPerAccount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "dailyRedeemLimitPerAccount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyRedeemLimitPerAccount is a free data retrieval call binding the contract method 0xcd897c71.
//
// Solidity: function dailyRedeemLimitPerAccount() view returns(uint256)
func (_StrBtc *StrBtcSession) DailyRedeemLimitPerAccount() (*big.Int, error) {
	return _StrBtc.Contract.DailyRedeemLimitPerAccount(&_StrBtc.CallOpts)
}

// DailyRedeemLimitPerAccount is a free data retrieval call binding the contract method 0xcd897c71.
//
// Solidity: function dailyRedeemLimitPerAccount() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) DailyRedeemLimitPerAccount() (*big.Int, error) {
	return _StrBtc.Contract.DailyRedeemLimitPerAccount(&_StrBtc.CallOpts)
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

// GetRemainingAccountRedeemLimit is a free data retrieval call binding the contract method 0xbdf446de.
//
// Solidity: function getRemainingAccountRedeemLimit(address account) view returns(uint256)
func (_StrBtc *StrBtcCaller) GetRemainingAccountRedeemLimit(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getRemainingAccountRedeemLimit", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingAccountRedeemLimit is a free data retrieval call binding the contract method 0xbdf446de.
//
// Solidity: function getRemainingAccountRedeemLimit(address account) view returns(uint256)
func (_StrBtc *StrBtcSession) GetRemainingAccountRedeemLimit(account common.Address) (*big.Int, error) {
	return _StrBtc.Contract.GetRemainingAccountRedeemLimit(&_StrBtc.CallOpts, account)
}

// GetRemainingAccountRedeemLimit is a free data retrieval call binding the contract method 0xbdf446de.
//
// Solidity: function getRemainingAccountRedeemLimit(address account) view returns(uint256)
func (_StrBtc *StrBtcCallerSession) GetRemainingAccountRedeemLimit(account common.Address) (*big.Int, error) {
	return _StrBtc.Contract.GetRemainingAccountRedeemLimit(&_StrBtc.CallOpts, account)
}

// GetRemainingGlobalRedeemLimit is a free data retrieval call binding the contract method 0x279b91cb.
//
// Solidity: function getRemainingGlobalRedeemLimit() view returns(uint256)
func (_StrBtc *StrBtcCaller) GetRemainingGlobalRedeemLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getRemainingGlobalRedeemLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingGlobalRedeemLimit is a free data retrieval call binding the contract method 0x279b91cb.
//
// Solidity: function getRemainingGlobalRedeemLimit() view returns(uint256)
func (_StrBtc *StrBtcSession) GetRemainingGlobalRedeemLimit() (*big.Int, error) {
	return _StrBtc.Contract.GetRemainingGlobalRedeemLimit(&_StrBtc.CallOpts)
}

// GetRemainingGlobalRedeemLimit is a free data retrieval call binding the contract method 0x279b91cb.
//
// Solidity: function getRemainingGlobalRedeemLimit() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) GetRemainingGlobalRedeemLimit() (*big.Int, error) {
	return _StrBtc.Contract.GetRemainingGlobalRedeemLimit(&_StrBtc.CallOpts)
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

// GetUsedAccountRedeemAmount is a free data retrieval call binding the contract method 0x4c5bb881.
//
// Solidity: function getUsedAccountRedeemAmount(address account) view returns(uint256)
func (_StrBtc *StrBtcCaller) GetUsedAccountRedeemAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getUsedAccountRedeemAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsedAccountRedeemAmount is a free data retrieval call binding the contract method 0x4c5bb881.
//
// Solidity: function getUsedAccountRedeemAmount(address account) view returns(uint256)
func (_StrBtc *StrBtcSession) GetUsedAccountRedeemAmount(account common.Address) (*big.Int, error) {
	return _StrBtc.Contract.GetUsedAccountRedeemAmount(&_StrBtc.CallOpts, account)
}

// GetUsedAccountRedeemAmount is a free data retrieval call binding the contract method 0x4c5bb881.
//
// Solidity: function getUsedAccountRedeemAmount(address account) view returns(uint256)
func (_StrBtc *StrBtcCallerSession) GetUsedAccountRedeemAmount(account common.Address) (*big.Int, error) {
	return _StrBtc.Contract.GetUsedAccountRedeemAmount(&_StrBtc.CallOpts, account)
}

// GetUsedGlobalRedeemAmount is a free data retrieval call binding the contract method 0x57e90cab.
//
// Solidity: function getUsedGlobalRedeemAmount() view returns(uint256)
func (_StrBtc *StrBtcCaller) GetUsedGlobalRedeemAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "getUsedGlobalRedeemAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsedGlobalRedeemAmount is a free data retrieval call binding the contract method 0x57e90cab.
//
// Solidity: function getUsedGlobalRedeemAmount() view returns(uint256)
func (_StrBtc *StrBtcSession) GetUsedGlobalRedeemAmount() (*big.Int, error) {
	return _StrBtc.Contract.GetUsedGlobalRedeemAmount(&_StrBtc.CallOpts)
}

// GetUsedGlobalRedeemAmount is a free data retrieval call binding the contract method 0x57e90cab.
//
// Solidity: function getUsedGlobalRedeemAmount() view returns(uint256)
func (_StrBtc *StrBtcCallerSession) GetUsedGlobalRedeemAmount() (*big.Int, error) {
	return _StrBtc.Contract.GetUsedGlobalRedeemAmount(&_StrBtc.CallOpts)
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

// RedeemLimitsEnabled is a free data retrieval call binding the contract method 0xef3e0084.
//
// Solidity: function redeemLimitsEnabled() view returns(bool)
func (_StrBtc *StrBtcCaller) RedeemLimitsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrBtc.contract.Call(opts, &out, "redeemLimitsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RedeemLimitsEnabled is a free data retrieval call binding the contract method 0xef3e0084.
//
// Solidity: function redeemLimitsEnabled() view returns(bool)
func (_StrBtc *StrBtcSession) RedeemLimitsEnabled() (bool, error) {
	return _StrBtc.Contract.RedeemLimitsEnabled(&_StrBtc.CallOpts)
}

// RedeemLimitsEnabled is a free data retrieval call binding the contract method 0xef3e0084.
//
// Solidity: function redeemLimitsEnabled() view returns(bool)
func (_StrBtc *StrBtcCallerSession) RedeemLimitsEnabled() (bool, error) {
	return _StrBtc.Contract.RedeemLimitsEnabled(&_StrBtc.CallOpts)
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

// ReinitializeV2 is a paid mutator transaction binding the contract method 0xc4115874.
//
// Solidity: function reinitializeV2() returns()
func (_StrBtc *StrBtcTransactor) ReinitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "reinitializeV2")
}

// ReinitializeV2 is a paid mutator transaction binding the contract method 0xc4115874.
//
// Solidity: function reinitializeV2() returns()
func (_StrBtc *StrBtcSession) ReinitializeV2() (*types.Transaction, error) {
	return _StrBtc.Contract.ReinitializeV2(&_StrBtc.TransactOpts)
}

// ReinitializeV2 is a paid mutator transaction binding the contract method 0xc4115874.
//
// Solidity: function reinitializeV2() returns()
func (_StrBtc *StrBtcTransactorSession) ReinitializeV2() (*types.Transaction, error) {
	return _StrBtc.Contract.ReinitializeV2(&_StrBtc.TransactOpts)
}

// ReinitializeV3 is a paid mutator transaction binding the contract method 0xbac22bb8.
//
// Solidity: function reinitializeV3() returns()
func (_StrBtc *StrBtcTransactor) ReinitializeV3(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "reinitializeV3")
}

// ReinitializeV3 is a paid mutator transaction binding the contract method 0xbac22bb8.
//
// Solidity: function reinitializeV3() returns()
func (_StrBtc *StrBtcSession) ReinitializeV3() (*types.Transaction, error) {
	return _StrBtc.Contract.ReinitializeV3(&_StrBtc.TransactOpts)
}

// ReinitializeV3 is a paid mutator transaction binding the contract method 0xbac22bb8.
//
// Solidity: function reinitializeV3() returns()
func (_StrBtc *StrBtcTransactorSession) ReinitializeV3() (*types.Transaction, error) {
	return _StrBtc.Contract.ReinitializeV3(&_StrBtc.TransactOpts)
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

// SetDailyGlobalRedeemLimit is a paid mutator transaction binding the contract method 0x587234dd.
//
// Solidity: function setDailyGlobalRedeemLimit(uint256 _limit) returns()
func (_StrBtc *StrBtcTransactor) SetDailyGlobalRedeemLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "setDailyGlobalRedeemLimit", _limit)
}

// SetDailyGlobalRedeemLimit is a paid mutator transaction binding the contract method 0x587234dd.
//
// Solidity: function setDailyGlobalRedeemLimit(uint256 _limit) returns()
func (_StrBtc *StrBtcSession) SetDailyGlobalRedeemLimit(_limit *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetDailyGlobalRedeemLimit(&_StrBtc.TransactOpts, _limit)
}

// SetDailyGlobalRedeemLimit is a paid mutator transaction binding the contract method 0x587234dd.
//
// Solidity: function setDailyGlobalRedeemLimit(uint256 _limit) returns()
func (_StrBtc *StrBtcTransactorSession) SetDailyGlobalRedeemLimit(_limit *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetDailyGlobalRedeemLimit(&_StrBtc.TransactOpts, _limit)
}

// SetDailyRedeemLimitPerAccount is a paid mutator transaction binding the contract method 0x1a7b3f92.
//
// Solidity: function setDailyRedeemLimitPerAccount(uint256 _limit) returns()
func (_StrBtc *StrBtcTransactor) SetDailyRedeemLimitPerAccount(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "setDailyRedeemLimitPerAccount", _limit)
}

// SetDailyRedeemLimitPerAccount is a paid mutator transaction binding the contract method 0x1a7b3f92.
//
// Solidity: function setDailyRedeemLimitPerAccount(uint256 _limit) returns()
func (_StrBtc *StrBtcSession) SetDailyRedeemLimitPerAccount(_limit *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetDailyRedeemLimitPerAccount(&_StrBtc.TransactOpts, _limit)
}

// SetDailyRedeemLimitPerAccount is a paid mutator transaction binding the contract method 0x1a7b3f92.
//
// Solidity: function setDailyRedeemLimitPerAccount(uint256 _limit) returns()
func (_StrBtc *StrBtcTransactorSession) SetDailyRedeemLimitPerAccount(_limit *big.Int) (*types.Transaction, error) {
	return _StrBtc.Contract.SetDailyRedeemLimitPerAccount(&_StrBtc.TransactOpts, _limit)
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

// SetRedeemLimitsEnabled is a paid mutator transaction binding the contract method 0x1e57686e.
//
// Solidity: function setRedeemLimitsEnabled(bool _enabled) returns()
func (_StrBtc *StrBtcTransactor) SetRedeemLimitsEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _StrBtc.contract.Transact(opts, "setRedeemLimitsEnabled", _enabled)
}

// SetRedeemLimitsEnabled is a paid mutator transaction binding the contract method 0x1e57686e.
//
// Solidity: function setRedeemLimitsEnabled(bool _enabled) returns()
func (_StrBtc *StrBtcSession) SetRedeemLimitsEnabled(_enabled bool) (*types.Transaction, error) {
	return _StrBtc.Contract.SetRedeemLimitsEnabled(&_StrBtc.TransactOpts, _enabled)
}

// SetRedeemLimitsEnabled is a paid mutator transaction binding the contract method 0x1e57686e.
//
// Solidity: function setRedeemLimitsEnabled(bool _enabled) returns()
func (_StrBtc *StrBtcTransactorSession) SetRedeemLimitsEnabled(_enabled bool) (*types.Transaction, error) {
	return _StrBtc.Contract.SetRedeemLimitsEnabled(&_StrBtc.TransactOpts, _enabled)
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

// StrBtcDailyGlobalRedeemLimitUpdatedIterator is returned from FilterDailyGlobalRedeemLimitUpdated and is used to iterate over the raw logs and unpacked data for DailyGlobalRedeemLimitUpdated events raised by the StrBtc contract.
type StrBtcDailyGlobalRedeemLimitUpdatedIterator struct {
	Event *StrBtcDailyGlobalRedeemLimitUpdated // Event containing the contract specifics and raw log

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
func (it *StrBtcDailyGlobalRedeemLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcDailyGlobalRedeemLimitUpdated)
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
		it.Event = new(StrBtcDailyGlobalRedeemLimitUpdated)
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
func (it *StrBtcDailyGlobalRedeemLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcDailyGlobalRedeemLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcDailyGlobalRedeemLimitUpdated represents a DailyGlobalRedeemLimitUpdated event raised by the StrBtc contract.
type StrBtcDailyGlobalRedeemLimitUpdated struct {
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDailyGlobalRedeemLimitUpdated is a free log retrieval operation binding the contract event 0x087142b64e709fbcc7d1f7097cf49027c4c377dfc0c9f78d9a6cf2d2c6274034.
//
// Solidity: event DailyGlobalRedeemLimitUpdated(uint256 newLimit)
func (_StrBtc *StrBtcFilterer) FilterDailyGlobalRedeemLimitUpdated(opts *bind.FilterOpts) (*StrBtcDailyGlobalRedeemLimitUpdatedIterator, error) {

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "DailyGlobalRedeemLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &StrBtcDailyGlobalRedeemLimitUpdatedIterator{contract: _StrBtc.contract, event: "DailyGlobalRedeemLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyGlobalRedeemLimitUpdated is a free log subscription operation binding the contract event 0x087142b64e709fbcc7d1f7097cf49027c4c377dfc0c9f78d9a6cf2d2c6274034.
//
// Solidity: event DailyGlobalRedeemLimitUpdated(uint256 newLimit)
func (_StrBtc *StrBtcFilterer) WatchDailyGlobalRedeemLimitUpdated(opts *bind.WatchOpts, sink chan<- *StrBtcDailyGlobalRedeemLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "DailyGlobalRedeemLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcDailyGlobalRedeemLimitUpdated)
				if err := _StrBtc.contract.UnpackLog(event, "DailyGlobalRedeemLimitUpdated", log); err != nil {
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

// ParseDailyGlobalRedeemLimitUpdated is a log parse operation binding the contract event 0x087142b64e709fbcc7d1f7097cf49027c4c377dfc0c9f78d9a6cf2d2c6274034.
//
// Solidity: event DailyGlobalRedeemLimitUpdated(uint256 newLimit)
func (_StrBtc *StrBtcFilterer) ParseDailyGlobalRedeemLimitUpdated(log types.Log) (*StrBtcDailyGlobalRedeemLimitUpdated, error) {
	event := new(StrBtcDailyGlobalRedeemLimitUpdated)
	if err := _StrBtc.contract.UnpackLog(event, "DailyGlobalRedeemLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtcDailyRedeemLimitPerAccountUpdatedIterator is returned from FilterDailyRedeemLimitPerAccountUpdated and is used to iterate over the raw logs and unpacked data for DailyRedeemLimitPerAccountUpdated events raised by the StrBtc contract.
type StrBtcDailyRedeemLimitPerAccountUpdatedIterator struct {
	Event *StrBtcDailyRedeemLimitPerAccountUpdated // Event containing the contract specifics and raw log

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
func (it *StrBtcDailyRedeemLimitPerAccountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcDailyRedeemLimitPerAccountUpdated)
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
		it.Event = new(StrBtcDailyRedeemLimitPerAccountUpdated)
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
func (it *StrBtcDailyRedeemLimitPerAccountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcDailyRedeemLimitPerAccountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcDailyRedeemLimitPerAccountUpdated represents a DailyRedeemLimitPerAccountUpdated event raised by the StrBtc contract.
type StrBtcDailyRedeemLimitPerAccountUpdated struct {
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDailyRedeemLimitPerAccountUpdated is a free log retrieval operation binding the contract event 0xdd7c282f04e637ea678511c395b45078fae3d4226e171c4af2307546f2f3be43.
//
// Solidity: event DailyRedeemLimitPerAccountUpdated(uint256 newLimit)
func (_StrBtc *StrBtcFilterer) FilterDailyRedeemLimitPerAccountUpdated(opts *bind.FilterOpts) (*StrBtcDailyRedeemLimitPerAccountUpdatedIterator, error) {

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "DailyRedeemLimitPerAccountUpdated")
	if err != nil {
		return nil, err
	}
	return &StrBtcDailyRedeemLimitPerAccountUpdatedIterator{contract: _StrBtc.contract, event: "DailyRedeemLimitPerAccountUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyRedeemLimitPerAccountUpdated is a free log subscription operation binding the contract event 0xdd7c282f04e637ea678511c395b45078fae3d4226e171c4af2307546f2f3be43.
//
// Solidity: event DailyRedeemLimitPerAccountUpdated(uint256 newLimit)
func (_StrBtc *StrBtcFilterer) WatchDailyRedeemLimitPerAccountUpdated(opts *bind.WatchOpts, sink chan<- *StrBtcDailyRedeemLimitPerAccountUpdated) (event.Subscription, error) {

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "DailyRedeemLimitPerAccountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcDailyRedeemLimitPerAccountUpdated)
				if err := _StrBtc.contract.UnpackLog(event, "DailyRedeemLimitPerAccountUpdated", log); err != nil {
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

// ParseDailyRedeemLimitPerAccountUpdated is a log parse operation binding the contract event 0xdd7c282f04e637ea678511c395b45078fae3d4226e171c4af2307546f2f3be43.
//
// Solidity: event DailyRedeemLimitPerAccountUpdated(uint256 newLimit)
func (_StrBtc *StrBtcFilterer) ParseDailyRedeemLimitPerAccountUpdated(log types.Log) (*StrBtcDailyRedeemLimitPerAccountUpdated, error) {
	event := new(StrBtcDailyRedeemLimitPerAccountUpdated)
	if err := _StrBtc.contract.UnpackLog(event, "DailyRedeemLimitPerAccountUpdated", log); err != nil {
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

// StrBtcRedeemLimitsEnabledUpdatedIterator is returned from FilterRedeemLimitsEnabledUpdated and is used to iterate over the raw logs and unpacked data for RedeemLimitsEnabledUpdated events raised by the StrBtc contract.
type StrBtcRedeemLimitsEnabledUpdatedIterator struct {
	Event *StrBtcRedeemLimitsEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *StrBtcRedeemLimitsEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtcRedeemLimitsEnabledUpdated)
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
		it.Event = new(StrBtcRedeemLimitsEnabledUpdated)
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
func (it *StrBtcRedeemLimitsEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtcRedeemLimitsEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtcRedeemLimitsEnabledUpdated represents a RedeemLimitsEnabledUpdated event raised by the StrBtc contract.
type StrBtcRedeemLimitsEnabledUpdated struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRedeemLimitsEnabledUpdated is a free log retrieval operation binding the contract event 0xe55d1ef19fc4263b518e1eecbe36f3cfd4f3192c442fc2b28122e779727f97f1.
//
// Solidity: event RedeemLimitsEnabledUpdated(bool enabled)
func (_StrBtc *StrBtcFilterer) FilterRedeemLimitsEnabledUpdated(opts *bind.FilterOpts) (*StrBtcRedeemLimitsEnabledUpdatedIterator, error) {

	logs, sub, err := _StrBtc.contract.FilterLogs(opts, "RedeemLimitsEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &StrBtcRedeemLimitsEnabledUpdatedIterator{contract: _StrBtc.contract, event: "RedeemLimitsEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchRedeemLimitsEnabledUpdated is a free log subscription operation binding the contract event 0xe55d1ef19fc4263b518e1eecbe36f3cfd4f3192c442fc2b28122e779727f97f1.
//
// Solidity: event RedeemLimitsEnabledUpdated(bool enabled)
func (_StrBtc *StrBtcFilterer) WatchRedeemLimitsEnabledUpdated(opts *bind.WatchOpts, sink chan<- *StrBtcRedeemLimitsEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _StrBtc.contract.WatchLogs(opts, "RedeemLimitsEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtcRedeemLimitsEnabledUpdated)
				if err := _StrBtc.contract.UnpackLog(event, "RedeemLimitsEnabledUpdated", log); err != nil {
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

// ParseRedeemLimitsEnabledUpdated is a log parse operation binding the contract event 0xe55d1ef19fc4263b518e1eecbe36f3cfd4f3192c442fc2b28122e779727f97f1.
//
// Solidity: event RedeemLimitsEnabledUpdated(bool enabled)
func (_StrBtc *StrBtcFilterer) ParseRedeemLimitsEnabledUpdated(log types.Log) (*StrBtcRedeemLimitsEnabledUpdated, error) {
	event := new(StrBtcRedeemLimitsEnabledUpdated)
	if err := _StrBtc.contract.UnpackLog(event, "RedeemLimitsEnabledUpdated", log); err != nil {
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
