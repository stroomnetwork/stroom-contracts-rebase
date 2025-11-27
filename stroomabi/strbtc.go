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
	Bin: "0x6080604052348015600e575f5ffd5b50614a408061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106103eb575f3560e01c80636739afca1161020b578063b03903251161011f578063ddd0c0b0116100b4578063e737557f11610084578063e737557f14610839578063ef3e008414610872578063f04da65b1461087f578063f376ebbb146108a7578063f8077fae146108d1575f5ffd5b8063ddd0c0b0146107f6578063e07fbd00146107ff578063e0a0461014610812578063e63ab1e914610825575f5ffd5b8063c87b7a23116100ef578063c87b7a23146107b4578063cd897c71146107c7578063d547741f146107d0578063dd62ed3e146107e3575f5ffd5b8063b039032514610788578063bac22bb814610791578063bdf446de14610799578063c4115874146107ac575f5ffd5b806391d14854116101a05780639fd95687116101705780639fd9568714610749578063a217fddf1461075c578063a7ce456514610763578063a9059cbb1461076c578063af7f131f1461077f575f5ffd5b806391d148541461070857806395d89b411461071b5780639cee9492146107235780639e76a00714610736575f5ffd5b806380239a04116101db57806380239a04146106c757806382dc1ec4146106da5780638456cb59146106ed5780638c659bf2146106f5575f5ffd5b80636739afca1461066e5780636b2c0f551461068e5780636ce1c4dc146106a157806370a08231146106b4575f5ffd5b8063313ce567116103025780634977305011610297578063587234dd11610267578063587234dd146105ec5780635abdb0dc146105ff5780635c975abb14610612578063603c6a6714610629578063607375a11461065b575f5ffd5b806349773050146105ab5780634c5bb881146105be57806355573c77146105d157806357e90cab146105e4575f5ffd5b80633a98ef39116102d25780633a98ef39146105895780633f4ba83a146105915780634219347314610599578063457e1a49146105a2575f5ffd5b8063313ce5671461054b5780633357325e1461055a5780633395b75a1461056d57806336568abe14610576575f5ffd5b806321172f3b1161038357806325e0a33f1161035357806325e0a33f146104ff5780632792949d14610512578063279b91cb1461051d5780632f2ff15d146105255780632f4cfd4714610538575f5ffd5b806321172f3b146104b357806323b872dd146104c6578063248a9ca3146104d957806324b76fd5146104ec575f5ffd5b806313776a8d116103be57806313776a8d1461046157806318160ddd146104835780631a7b3f921461048b5780631e57686e146104a0575f5ffd5b806301ffc9a7146103ef578063026034f01461041757806306fdde0314610439578063095ea7b31461044e575b5f5ffd5b6104026103fd366004614254565b6108da565b60405190151581526020015b60405180910390f35b61040261042536600461427b565b60096020525f908152604090205460ff1681565b610441610910565b60405161040e91906142c0565b61040261045c3660046142e6565b6109d0565b6104755f5160206149eb5f395f51905f5281565b60405190815260200161040e565b600454610475565b61049e61049936600461427b565b6109e7565b005b61049e6104ae36600461431d565b610a2e565b61049e6104c1366004614375565b610a79565b6104026104d43660046143c3565b610cde565b6104756104e736600461427b565b610d03565b61049e6104fa366004614401565b610d23565b61049e61050d36600461427b565b610e13565b6104756305f5e10081565b610475610e46565b61049e610533366004614448565b610ea5565b61049e6105463660046142e6565b610ec7565b6040516008815260200161040e565b61049e6105683660046142e6565b610f5d565b61047560065481565b61049e610584366004614448565b610fea565b600554610475565b61049e611022565b61047561022281565b61047560015481565b61049e6105b9366004614476565b611044565b6104756105cc366004614476565b611097565b6104756105df36600461427b565b6110de565b610475611127565b61049e6105fa36600461427b565b611154565b61049e61060d36600461427b565b611193565b5f5160206149715f395f51905f525460ff16610402565b610441604051806040016040528060138152602001725354524f4f4d5f4d494e545f494e564f49434560681b81525081565b61049e610669366004614491565b6111c6565b5f5461068190600160a01b900460ff1681565b60405161040e9190614501565b61049e61069c366004614476565b611382565b61049e6106af366004614476565b6113a7565b6104756106c2366004614476565b6113c8565b6104756106d5366004614527565b611401565b61049e6106e8366004614476565b6114ad565b61049e6114ce565b610441610703366004614527565b6114ed565b610402610716366004614448565b611532565b610441611568565b61049e61073136600461427b565b6115a6565b61049e610744366004614476565b6115d8565b61049e61075736600461455d565b6115f9565b6104755f81565b61047560025481565b61040261077a3660046142e6565b6116f2565b61047560035481565b61047560075481565b61049e6116ff565b6104756107a7366004614476565b6117ca565b61049e611848565b6104756107c236600461427b565b6118f6565b610475600b5481565b61049e6107de366004614448565b611935565b6104756107f1366004614593565b611951565b610475600c5481565b61047561080d3660046145bf565b61199a565b6104416108203660046145bf565b611a3e565b6104755f5160206149315f395f51905f5281565b6104416040518060400160405280601a8152602001795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b81525081565b600d546104029060ff1681565b61047561088d366004614476565b6001600160a01b03165f908152600a602052604090205490565b5f546108b9906001600160a01b031681565b6040516001600160a01b03909116815260200161040e565b61047560085481565b5f6001600160e01b03198216637965db0b60e01b148061090a57506301ffc9a760e01b6001600160e01b03198316145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060915f5160206149115f395f51905f529161094e906145d9565b80601f016020809104026020016040519081016040528092919081815260200182805461097a906145d9565b80156109c55780601f1061099c576101008083540402835291602001916109c5565b820191905f5260205f20905b8154815290600101906020018083116109a857829003601f168201915b505050505091505090565b5f336109dd818585611aa1565b5060019392505050565b5f6109f181611aae565b600b8290556040518281527fdd7c282f04e637ea678511c395b45078fae3d4226e171c4af2307546f2f3be43906020015b60405180910390a15050565b5f610a3881611aae565b600d805460ff19168315159081179091556040519081527fe55d1ef19fc4263b518e1eecbe36f3cfd4f3192c442fc2b28122e779727f97f190602001610a22565b610a81611ab8565b6040518060400160405280601a8152602001795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b815250610abe85856114ed565b5f546040516311ee58a960e01b8152859185916001600160a01b03909116906311ee58a990610af7908790879087908790600401614633565b602060405180830381865afa158015610b12573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b369190614677565b610b5357604051636227817160e01b815260040160405180910390fd5b6003548814610b75576040516336b118d760e11b815260040160405180910390fd5b600160035f828254610b8791906146a6565b90915550505f879003610bad57604051630aa7fcbb60e21b815260040160405180910390fd5b600754600854610bbd91906146a6565b421015610bdc576040516283cb6b60e51b815260040160405180910390fd5b5f612710600654600454610bf091906146b9565b610bfa91906146d0565b905080881115610c1d57604051630779a2a360e31b815260040160405180910390fd5b5f610c288a8a611401565b5f8181526009602052604090205490915060ff1615610c5a5760405163ce199d1760e01b815260040160405180910390fd5b5f818152600960205260408120805460ff19166001179055600480548b9290610c849084906146a6565b909155505042600855600454600554604080518d815260208101939093528201527f339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b934409060600160405180910390a150505050505050505050565b5f33610ceb858285611aea565b610cf6858585611b4c565b60019150505b9392505050565b5f9081525f5160206149515f395f51905f52602052604090206001015490565b610d2b611ab8565b600154831015610d4e57604051633813eacd60e21b815260040160405180910390fd5b610d7a82825f60149054906101000a900460ff166003811115610d7357610d736144ed565b9190611ba9565b610d97576040516312d58f2760e21b815260040160405180910390fd5b600d5460ff1615610dac57610dac3384611ec0565b610db63384611fa8565b600160025f828254610dc891906146a6565b909155505060025460405133917f83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e91610e06918691869189916146ef565b60405180910390a2505050565b5f610e1d81611aae565b61a8c0821015610e405760405163a7bbe41d60e01b815260040160405180910390fd5b50600755565b600d545f9060ff16610e5857505f1990565b5f610e6662015180426146d0565b905080600e600101541015610e7d575050600c5490565b600c54600e5410610e8f575f91505090565b600e54600c54610e9f9190614715565b91505090565b610eae82610d03565b610eb781611aae565b610ec18383611fdc565b50505050565b610ecf611ab8565b5f5160206149eb5f395f51905f52610ee681611aae565b6001600160a01b038316610f0d576040516389a4ea1960e01b815260040160405180910390fd5b610f17838361207d565b6040518281526001600160a01b0384169033907f76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e906020015b60405180910390a3505050565b610f65611ab8565b5f5160206149eb5f395f51905f52610f7c81611aae565b6001600160a01b038316610fa35760405163a30d2d8760e01b815260040160405180910390fd5b610fad8383611fa8565b6040518281526001600160a01b0384169033907f25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad56790602001610f50565b6001600160a01b03811633146110135760405163334bd91960e11b815260040160405180910390fd5b61101d82826120b1565b505050565b5f5160206149315f395f51905f5261103981611aae565b61104161212a565b50565b5f61104e81611aae565b6001600160a01b0382166110755760405163e6c4247b60e01b815260040160405180910390fd5b505f80546001600160a01b0319166001600160a01b0392909216919091179055565b5f806110a662015180426146d0565b6001600160a01b0384165f9081526010602052604090206001810154919250908211156110d657505f9392505050565b549392505050565b5f6005545f14806110ef5750600454155b1561110d576040516340a3daff60e11b815260040160405180910390fd5b60055460045461111d90846146b9565b61090a91906146d0565b5f8061113662015180426146d0565b905080600e60010154101561114c575f91505090565b5050600e5490565b5f61115e81611aae565b600c8290556040518281527f087142b64e709fbcc7d1f7097cf49027c4c377dfc0c9f78d9a6cf2d2c627403490602001610a22565b5f61119d81611aae565b6102228210156111c0576040516343a87e9960e11b815260040160405180910390fd5b50600155565b5f5160206149915f395f51905f528054600160401b810460ff1615906001600160401b03165f811580156111f75750825b90505f826001600160401b031660011480156112125750303b155b905081158015611220575080155b1561123e5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561126857845460ff60401b1916600160401b1785555b6112b66040518060400160405280600e81526020016d29ba3937b7b6902134ba31b7b4b760911b8152506040518060400160405280600681526020016573747242544360d01b815250612189565b6112be61219b565b6112c66121ab565b6112cf886121b3565b611b586001555f80548a919060ff60a01b1916600160a01b8360038111156112f9576112f96144ed565b02179055506113085f88611fdc565b506113205f5160206149315f395f51905f5287611fdc565b506064600655620151806007555f600855831561137757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f61138c81611aae565b6113a35f5160206149315f395f51905f5283611935565b5050565b5f6113b181611aae565b6113a35f5160206149eb5f395f51905f5283610ea5565b5f6005545f036113d957505f919050565b6005546004546001600160a01b0384165f908152600a602052604090205461111d91906146b9565b5f805460408051808201909152601a8152795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b60208201526001600160a01b039091169063bca0ac069061145186866114ed565b6040518363ffffffff1660e01b815260040161146e929190614728565b602060405180830381865afa158015611489573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cfc9190614755565b5f6114b781611aae565b6113a35f5160206149315f395f51905f5283610ea5565b5f5160206149315f395f51905f526114e581611aae565b6110416121dc565b604080516020810184905290810182905230606090811b6bffffffffffffffffffffffff19168183015290607401604051602081830303815290604052905092915050565b5f9182525f5160206149515f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060915f5160206149115f395f51905f529161094e906145d9565b5f6115b081611aae565b60648211156115d257604051630737692960e31b815260040160405180910390fd5b50600655565b5f6115e281611aae565b6113a35f5160206149eb5f395f51905f5283611935565b611601611ab8565b604051806040016040528060138152602001725354524f4f4d5f4d494e545f494e564f49434560681b81525061163684611a3e565b5f546040516311ee58a960e01b8152859185916001600160a01b03909116906311ee58a99061166f908790879087908790600401614633565b602060405180830381865afa15801561168a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116ae9190614677565b6116cb57604051636227817160e01b815260040160405180910390fd5b6116e9604088018035906116e29060208b01614476565b8935612224565b50505050505050565b5f336109dd818585611b4c565b5f5160206149915f395f51905f52805460039190600160401b900460ff1680611735575080546001600160401b03808416911610155b156117535760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b038316908117600160401b178255620f4240600b556305f5e100600c55600d805460ff19169055815460ff60401b191682556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602001610a22565b600d545f9060ff166117de57505f19919050565b5f6117ec62015180426146d0565b6001600160a01b0384165f90815260106020526040902060018101549192509082111561181e575050600b5492915050565b600b5481541061183157505f9392505050565b8054600b546118409190614715565b949350505050565b5f5160206149915f395f51905f52805460029190600160401b900460ff168061187e575080546001600160401b03808416911610155b1561189c5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b038316908117600160401b1760ff60401b191682556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602001610a22565b5f6005545f14806119075750600454155b15611925576040516340a3daff60e11b815260040160405180910390fd5b60045460055461111d90846146b9565b61193e82610d03565b61194781611aae565b610ec183836120b1565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b5f80546040805180820190915260138152725354524f4f4d5f4d494e545f494e564f49434560681b60208201526001600160a01b039091169063bca0ac06906119e285611a3e565b6040518363ffffffff1660e01b81526004016119ff929190614728565b602060405180830381865afa158015611a1a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061090a9190614755565b6060611a506040830160208401614476565b604080516bffffffffffffffffffffffff19606093841b81166020830152918501356034820152843560548201523090921b1660748201526088016040516020818303038152906040529050919050565b61101d8383836001612330565b6110418133612414565b5f5160206149715f395f51905f525460ff1615611ae85760405163d93c066560e01b815260040160405180910390fd5b565b5f611af58484611951565b90505f198114610ec15781811015611b3e57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b610ec184848484035f612330565b6001600160a01b038316611b7557604051634b637e8f60e11b81525f6004820152602401611b35565b6001600160a01b038216611b9e5760405163ec442f0560e01b81525f6004820152602401611b35565b61101d83838361244d565b5f6060611bec84848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508592506125e0915050565b15611bfa575f915050610cfc565b611c2c604051806040016040528060118152602001700a7261772061646472657373206461746160781b81525061265a565b611c6a84848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061269d92505050565b5f611c74866126e0565b90505f611c80876127f8565b90505f611c8c886128d1565b9050611cd9611c9e60015f898b61476c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508792506125e0915050565b80611d805750611d2a611cef60015f898b61476c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508692506125e0915050565b8015611d8057508051611d7e90611d43905f898b61476c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508592506125e0915050565b155b15611e0057601a861080611d945750602386115b80611dda5750611dd887878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506129b492505050565b155b15611deb575f945050505050610cfc565b611df58787612a05565b945050505050610cfc565b8051611e1290611d43905f898b61476c565b15611eb3576002886003811115611e2b57611e2b6144ed565b03611e5557602b861080611e3f5750604086115b15611e50575f945050505050610cfc565b611e75565b602a861080611e645750603e86115b15611e75575f945050505050610cfc565b611df587878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250612f2b92505050565b505f979650505050505050565b5f611ece62015180426146d0565b6001600160a01b0384165f908152601060205260409020600181015491925090821115611f00575f8155600181018290555b600b548154611f109085906146a6565b1115611f2f5760405163e55017a760e01b815260040160405180910390fd5b82815f015f828254611f4191906146a6565b9091555050600f54821115611f5a575f600e55600f8290555b600c54600e54611f6b9085906146a6565b1115611f8a5760405163c998883360e01b815260040160405180910390fd5b82600e5f015f828254611f9d91906146a6565b909155505050505050565b6001600160a01b038216611fd157604051634b637e8f60e11b81525f6004820152602401611b35565b6113a3825f8361244d565b5f5f5160206149515f395f51905f52611ff58484611532565b612074575f848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561202a3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061090a565b5f91505061090a565b6001600160a01b0382166120a65760405163ec442f0560e01b81525f6004820152602401611b35565b6113a35f838361244d565b5f5f5160206149515f395f51905f526120ca8484611532565b15612074575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061090a565b61213261355d565b5f5160206149715f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b61219161358c565b6113a382826135c2565b6121a361358c565b611ae8613612565b611ae861358c565b6121bb61358c565b5f80546001600160a01b0319166001600160a01b0392909216919091179055565b6121e4611ab8565b5f5160206149715f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361216b565b825f0361224457604051631cebf66f60e11b815260040160405180910390fd5b6122566305f5e1006301406f406146b9565b831061227557604051636006659960e11b815260040160405180910390fd5b306001600160a01b0383160361229e5760405163603acaa960e11b815260040160405180910390fd5b5f8181526009602052604090205460ff16156122cd57604051637bdb87d160e01b815260040160405180910390fd5b5f818152600960205260409020805460ff191660011790556122ef828461207d565b60408051848152602081018390526001600160a01b038416917fb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa4699101610e06565b5f5160206149115f395f51905f526001600160a01b0385166123675760405163e602df0560e01b81525f6004820152602401611b35565b6001600160a01b03841661239057604051634a1406b160e11b81525f6004820152602401611b35565b6001600160a01b038086165f9081526001830160209081526040808320938816835292905220839055811561240d57836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161240491815260200190565b60405180910390a35b5050505050565b61241e8282611532565b6113a35760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401611b35565b6001600160a01b0383166124d3575f6005545f148061246c5750600454155b61247e57612479826118f6565b612480565b815b90508160045f82825461249391906146a6565b925050819055508060055f8282546124ab91906146a6565b90915550506001600160a01b0383165f908152600a602052604090208054909101905561259b565b5f6124dd846113c8565b90508181101561250057604051631e9acf1760e31b815260040160405180910390fd5b5f61250a836118f6565b90506001600160a01b03841661256b578260045f82825461252b9190614715565b925050819055508060055f8282546125439190614715565b90915550506001600160a01b0385165f908152600a6020526040902080548290039055612598565b6001600160a01b038086165f908152600a6020526040808220805485900390559186168152208054820190555b50505b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f5091815260200190565b5f81518351146125f157505f61090a565b5f5b83518110156109dd5782818151811061260e5761260e614793565b602001015160f81c60f81b6001600160f81b03191684828151811061263557612635614793565b01602001516001600160f81b03191614612652575f91505061090a565b6001016125f3565b6110418160405160240161266e91906142c0565b60408051601f198184030181529190526020810180516001600160e01b031663104c13eb60e21b179052613632565b611041816040516024016126b191906142c0565b60408051601f198184030181529190526020810180516001600160e01b03166305f3bfab60e11b179052613632565b60605f8260038111156126f5576126f56144ed565b036127175750506040805180820190915260018152603160f81b602082015290565b600282600381111561272b5761272b6144ed565b0361274d5750506040805180820190915260018152601960f91b602082015290565b6001826003811115612761576127616144ed565b036127835750506040805180820190915260018152601960f91b602082015290565b6003826003811115612797576127976144ed565b036127b95750506040805180820190915260018152605360f81b602082015290565b60405162461bcd60e51b8152602060048201526014602482015273556e6b6e6f776e206e6574776f726b207479706560601b6044820152606401611b35565b60605f82600381111561280d5761280d6144ed565b0361282f5750506040805180820190915260018152603360f81b602082015290565b6002826003811115612843576128436144ed565b036128655750506040805180820190915260018152606d60f81b602082015290565b6001826003811115612879576128796144ed565b0361289b5750506040805180820190915260018152606d60f81b602082015290565b60038260038111156128af576128af6144ed565b036127b95750506040805180820190915260018152607360f81b602082015290565b60605f8260038111156128e6576128e66144ed565b0361290a57505060408051808201909152600381526262633160e81b602082015290565b600282600381111561291e5761291e6144ed565b03612944575050604080518082019091526005815264626372743160d81b602082015290565b6001826003811115612958576129586144ed565b0361297c57505060408051808201909152600381526274623160e81b602082015290565b6003826003811115612990576129906144ed565b036127b957505060408051808201909152600381526273623160e81b602082015290565b80515f90815b818110156109dd575f8482815181106129d5576129d5614793565b016020015160f81c90505f6129e98261363b565b9050806129fb57505f95945050505050565b50506001016129ba565b5f5f612a4584848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506136e192505050565b9050612a7e6040518060400160405280601681526020017576616c6964617465426173653538436865636b73756d60501b81525061265a565b612aa6604051806040016040528060078152602001661c185e5b1bd85960ca1b81525061265a565b612aaf8161269d565b8051601914612ac1575f91505061090a565b6040805160018082528183019092525f91602082018180368337019050509050815f81518110612af357612af3614793565b602001015160f81c60f81b815f81518110612b1057612b10614793565b60200101906001600160f81b03191690815f1a905350612b4e604051806040016040528060078152602001663b32b939b4b7b760c91b81525061265a565b612b578161269d565b5f600460018451612b689190614715565b612b729190614715565b6001600160401b03811115612b8957612b896147a7565b6040519080825280601f01601f191660200182016040528015612bb3576020820181803683370190505b5090505f600460018551612bc79190614715565b612bd19190614715565b90505f5b81811015612c345784612be98260016146a6565b81518110612bf957612bf9614793565b602001015160f81c60f81b838281518110612c1657612c16614793565b60200101906001600160f81b03191690815f1a905350600101612bd5565b50612c5d604051806040016040528060078152602001661c185e5b1bd85960ca1b81525061265a565b612c668261269d565b8151601414612c7b575f94505050505061090a565b6040805160048082528183019092525f916020820181803683370190505090505f5b6004811015612d0b57858160048851612cb69190614715565b612cc091906146a6565b81518110612cd057612cd0614793565b602001015160f81c60f81b828281518110612ced57612ced614793565b60200101906001600160f81b03191690815f1a905350600101612c9d565b50612d3560405180604001604052806008815260200167636865636b73756d60c01b81525061265a565b612d3e8161269d565b5f6002808686604051602001612d559291906147d2565b60408051601f1981840301815290829052612d6f916147e6565b602060405180830381855afa158015612d8a573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190612dad9190614755565b604051602001612dbf91815260200190565b60408051601f1981840301815290829052612dd9916147e6565b602060405180830381855afa158015612df4573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190612e179190614755565b9050612e4d6040518060400160405280601381526020017263616c63756c6174656420636865636b73756d60681b81525061265a565b612e56816136ec565b8060031a60f81b82600381518110612e7057612e70614793565b01602001516001600160f81b031916188160021a60f81b83600281518110612e9a57612e9a614793565b01602001516001600160f81b031916188260011a60f81b84600181518110612ec457612ec4614793565b01602001516001600160f81b03191618835f1a60f81b855f81518110612eec57612eec614793565b602001015160f81c60f81b181717176001600160f81b0319165f60f81b14612f1c575f965050505050505061090a565b50600198975050505050505050565b5f612f6a6040518060400160405280601981526020017f0a76616c69646174652062656368333220636865636b73756d0000000000000081525061265a565b612f92604051806040016040528060078152602001666164647265737360c81b81525061265a565b612f9b8261265a565b8151829060081115612fd957612fd1604051806040016040528060098152602001681d1bdbc81cda1bdc9d60ba1b81525061265a565b505f92915050565b605a8151111561300c57612fd160405180604001604052806008815260200167746f6f206c6f6e6760c01b81525061265a565b50815182905f90815b818110156130585783818151811061302f5761302f614793565b01602001516001600160f81b031916603160f81b0361305057809250613058565b600101613015565b50815f036130975761308d6040518060400160405280600c81526020016b37379039b2b830b930ba37b960a11b81525061265a565b505f949350505050565b816001036130ce5761308d6040518060400160405280600e81526020016d0dad2e6e6d2dcce40e0e4caccd2f60931b81525061265a565b5f826001600160401b038111156130e7576130e76147a7565b6040519080825280601f01601f191660200182016040528015613111576020820181803683370190505b5090505f60018486516131249190614715565b61312e9190614715565b90505f816001600160401b03811115613149576131496147a7565b6040519080825280601f01601f191660200182016040528015613173576020820181803683370190505b5090505f5b858110156131cd5786818151811061319257613192614793565b602001015160f81c60f81b8482815181106131af576131af614793565b60200101906001600160f81b03191690815f1a905350600101613178565b505f5b8281101561323957866131e387836146a6565b6131ee9060016146a6565b815181106131fe576131fe614793565b602001015160f81c60f81b82828151811061321b5761321b614793565b60200101906001600160f81b03191690815f1a9053506001016131d0565b50613261604051806040016040528060068152602001650e0e4caccd2f60d31b81525061265a565b61326a8361269d565b6006815110156132a357611eb36040518060400160405280600e81526020016d19185d18481d1bdbc81cda1bdc9d60921b81525061265a565b5f6132ad84613731565b9050805f036132f3576132e56040518060400160405280600e81526020016d0d2dcecc2d8d2c840e0e4caccd2f60931b81525061265a565b505f98975050505050505050565b5f82516001600160401b0381111561330d5761330d6147a7565b6040519080825280601f01601f191660200182016040528015613337576020820181803683370190505b5080519091505f5b81811015613490575f85828151811061335a5761335a614793565b01602001516001600160f81b03191690505f61337582613820565b905060fe1960ff82160161342f576133b5604051806040016040528060118152602001703ab735b737bbb71031b430b930b1ba32b960791b81525061265a565b6133be83613c03565b6133c782613c48565b61341c6040518060400160405280600481526020016331b430b960e11b8152508360405160200161340891906001600160f81b031991909116815260010190565b604051602081830303815290604052613c91565b505f9d9c50505050505050505050505050565b8060ff1661343c87613cd6565b88519118965061344d8460066146a6565b10613459575050613488565b8060f81b85848151811061346f5761346f614793565b60200101906001600160f81b03191690815f1a90535050505b60010161333f565b506134b760405180604001604052806005815260200164776f72647360d81b81525061265a565b6134c08261269d565b632bc830a383141580156134d5575082600114155b1561351c5761350c6040518060400160405280601081526020016f696e76616c696420636865636b73756d60801b81525084613d90565b505f9a9950505050505050505050565b61354c6040518060400160405280600e81526020016d76616c696420636865636b73756d60901b81525084613d90565b5060019a9950505050505050505050565b5f5160206149715f395f51905f525460ff16611ae857604051638dfc202b60e01b815260040160405180910390fd5b5f5160206149915f395f51905f5254600160401b900460ff16611ae857604051631afcd79f60e31b815260040160405180910390fd5b6135ca61358c565b5f5160206149115f395f51905f527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036136038482614835565b5060048101610ec18382614835565b61361a61358c565b5f5160206149715f395f51905f52805460ff19169055565b61104181613dd5565b5f8160ff166049148061365157508160ff16604f145b8061365f57508160ff16606c145b1561366b57505f919050565b60318260ff1610158015613683575060398260ff1611155b1561369057506001919050565b60418260ff16101580156136a85750605a8260ff1611155b156136b557506001919050565b60618260ff16101580156136cd5750607a8260ff1611155b156136da57506001919050565b505f919050565b606061090a82613df4565b6110418160405160240161370291815260200190565b60408051601f198184030181529190526020810180516001600160e01b03166327b7cf8560e01b179052613632565b80515f90600190825b818110156137c6575f85828151811061375557613755614793565b016020015160f81c9050602181108061376e5750607e81115b156137ac5760405162461bcd60e51b815260206004820152600e60248201526d092dcecc2d8d2c840e0e4caccd2f60931b6044820152606401611b35565b600581901c6137ba85613cd6565b1893505060010161373a565b506137d082613cd6565b845190925090505f5b81811015613817575f8582815181106137f4576137f4614793565b016020015160f81c9050601f811661380b85613cd6565b189350506001016137d9565b50909392505050565b5f600d60fc1b6001600160f81b031983160161383e5750600f919050565b606760f91b6001600160f81b031983160161385b5750600a919050565b60cd60f81b6001600160f81b031983160161387857506011919050565b603360fa1b6001600160f81b031983160161389557506015919050565b60cb60f81b6001600160f81b03198316016138b257506014919050565b606560f91b6001600160f81b03198316016138cf5750601a919050565b601960fb1b6001600160f81b03198316016138ec57506007919050565b60c960f81b6001600160f81b03198316016139095750601e919050565b60c760f81b6001600160f81b031983160161392657506005919050565b608f60f81b6001600160f81b031983160161394257505f919050565b600960fc1b6001600160f81b031983160161395f57506001919050565b604360f91b6001600160f81b031983160161397c57506002919050565b604760f91b6001600160f81b031983160161399957506003919050565b608760f81b6001600160f81b03198316016139b657506004919050565b601160fb1b6001600160f81b03198316016139d357506006919050565b609960f81b6001600160f81b03198316016139f057506008919050565b604d60f91b6001600160f81b0319831601613a0d57506009919050565b602360fa1b6001600160f81b0319831601613a2a5750600b919050565b604560f91b6001600160f81b0319831601613a475750600c919050565b602760fa1b6001600160f81b0319831601613a645750600d919050565b608960f81b6001600160f81b0319831601613a815750600e919050565b608d60f81b6001600160f81b0319831601613a9e57506010919050565b604b60f91b6001600160f81b0319831601613abb57506012919050565b604960f91b6001600160f81b0319831601613ad857506013919050565b609560f81b6001600160f81b0319831601613af557506016919050565b601360fb1b6001600160f81b0319831601613b1257506017919050565b609d60f81b6001600160f81b0319831601613b2f57506018919050565b609b60f81b6001600160f81b0319831601613b4c57506019919050565b609360f81b6001600160f81b0319831601613b695750601b919050565b608b60f81b6001600160f81b0319831601613b865750601c919050565b609f60f81b6001600160f81b0319831601613ba35750601d919050565b602560fa1b6001600160f81b0319831601613bc05750601f919050565b613bf26040518060400160405280601181526020017024b73b30b634b21031b430b930b1ba32b960791b81525061265a565b613bfb82613c48565b5060ff919050565b61104181604051602401613c1991815260200190565b60408051601f198184030181529190526020810180516001600160e01b031663f82c50f160e01b179052613632565b6040516001600160f81b0319821660248201526110419060440160408051601f198184030181529190526020810180516001600160e01b0316630dc3142560e31b179052613632565b6113a38282604051602401613ca7929190614728565b60408051601f198184030181529190526020810180516001600160e01b0316634b5c427760e01b179052613632565b5f601982901c6001601d84901c811614613cf0575f613cf6565b632a1462b35b63ffffffff16600382901c600116600114613d11575f613d17565b633d4233dd5b63ffffffff16600283901c600116600114613d32575f613d38565b631ea119fa5b63ffffffff16600184811c811614613d50575f613d56565b6326508e6d5b63ffffffff16600180861614613d6c575f613d72565b633b6a57b25b63ffffffff166005886301ffffff16901b1818181818915050919050565b6113a38282604051602401613da69291906148ef565b60408051601f198184030181529190526020810180516001600160e01b0316632d839cb360e21b179052613632565b5f6a636f6e736f6c652e6c6f6790505f5f835160208501845afa505050565b80516060906031905f805b8281108015613e26575083868281518110613e1c57613e1c614793565b016020015160f81c145b15613e375760019182019101613dff565b505f8080806117e361209f8702046001016002026001600160401b03811115613e6257613e626147a7565b6040519080825280601f01601f191660200182016040528015613e8c576020820181803683370190505b5090505f600460038801046001600160401b03811115613eae57613eae6147a7565b604051908082528060200260200182016040528015613ed7578160200160208202803683370190505b508a519091505f5b81811015613ff6575f8c8281518110613efa57613efa614793565b602001015160f81c60f81b9050613f296040518060600160405280603a81526020016149b1603a913982614140565b909750955085613f725760405162461bcd60e51b81526020600482015260146024820152731a5b9d985b1a590818985cd94d4e08191a59da5d60621b6044820152606401611b35565b83515f19015b5f8112613fec5787858281518110613f9257613f92614793565b602002602001015163ffffffff16603a026001600160401b0316019850602089901c97508863ffffffff16858281518110613fcf57613fcf614793565b63ffffffff909216602092830291909101909101525f1901613f78565b5050600101613edf565b50600860038916026001600160401b0381165f03614012575060205b8251600719909101905f90815b818110156140bf575b6020846001600160401b031610156140b357836001600160401b031686828151811061405657614056614793565b602002602001015163ffffffff16901c60f81b87848151811061407b5761407b614793565b60200101906001600160f81b03191690815f1a90535060019092019160086001600160401b038516106140b357600884039350614028565b6018935060010161401f565b5085518a5b81811015614121575f60f81b8882815181106140e2576140e2614793565b01602001516001600160f81b031916111561411957614104888d8303866141a6565b9e505050505050505050505050505050919050565b6001016140c4565b5061412d875f856141a6565b9f9e505050505050505050505050505050565b81515f908190815b8181101561419657846001600160f81b03191686828151811061416d5761416d614793565b01602001516001600160f81b0319160361418e5792506001915061419f9050565b600101614148565b505f5f92509250505b9250929050565b60605f8383036001600160401b038111156141c3576141c36147a7565b6040519080825280601f01601f1916602001820160405280156141ed576020820181803683370190505b5090505f5b84840381101561424b57858582018151811061421057614210614793565b602001015160f81c60f81b82828151811061422d5761422d614793565b60200101906001600160f81b03191690815f1a9053506001016141f2565b50949350505050565b5f60208284031215614264575f5ffd5b81356001600160e01b031981168114610cfc575f5ffd5b5f6020828403121561428b575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610cfc6020830184614292565b6001600160a01b0381168114611041575f5ffd5b5f5f604083850312156142f7575f5ffd5b8235614302816142d2565b946020939093013593505050565b8015158114611041575f5ffd5b5f6020828403121561432d575f5ffd5b8135610cfc81614310565b5f5f83601f840112614348575f5ffd5b5081356001600160401b0381111561435e575f5ffd5b60208301915083602082850101111561419f575f5ffd5b5f5f5f5f60608587031215614388575f5ffd5b843593506020850135925060408501356001600160401b038111156143ab575f5ffd5b6143b787828801614338565b95989497509550505050565b5f5f5f606084860312156143d5575f5ffd5b83356143e0816142d2565b925060208401356143f0816142d2565b929592945050506040919091013590565b5f5f5f60408486031215614413575f5ffd5b8335925060208401356001600160401b0381111561442f575f5ffd5b61443b86828701614338565b9497909650939450505050565b5f5f60408385031215614459575f5ffd5b82359150602083013561446b816142d2565b809150509250929050565b5f60208284031215614486575f5ffd5b8135610cfc816142d2565b5f5f5f5f608085870312156144a4575f5ffd5b8435600481106144b2575f5ffd5b935060208501356144c2816142d2565b925060408501356144d2816142d2565b915060608501356144e2816142d2565b939692955090935050565b634e487b7160e01b5f52602160045260245ffd5b602081016004831061452157634e487b7160e01b5f52602160045260245ffd5b91905290565b5f5f60408385031215614538575f5ffd5b50508035926020909101359150565b5f60608284031215614557575f5ffd5b50919050565b5f5f5f6080848603121561456f575f5ffd5b6145798585614547565b925060608401356001600160401b0381111561442f575f5ffd5b5f5f604083850312156145a4575f5ffd5b82356145af816142d2565b9150602083013561446b816142d2565b5f606082840312156145cf575f5ffd5b610cfc8383614547565b600181811c908216806145ed57607f821691505b60208210810361455757634e487b7160e01b5f52602260045260245ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606081525f6146456060830187614292565b82810360208401526146578187614292565b9050828103604084015261466c81858761460b565b979650505050505050565b5f60208284031215614687575f5ffd5b8151610cfc81614310565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561090a5761090a614692565b808202811582820484141761090a5761090a614692565b5f826146ea57634e487b7160e01b5f52601260045260245ffd5b500490565b606081525f61470260608301868861460b565b6020830194909452506040015292915050565b8181038181111561090a5761090a614692565b604081525f61473a6040830185614292565b828103602084015261474c8185614292565b95945050505050565b5f60208284031215614765575f5ffd5b5051919050565b5f5f8585111561477a575f5ffd5b83861115614786575f5ffd5b5050820193919092039150565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b5f81518060208401855e5f93019283525090919050565b5f6118406147e083866147bb565b846147bb565b5f610cfc82846147bb565b601f82111561101d57805f5260205f20601f840160051c810160208510156148165750805b601f840160051c820191505b8181101561240d575f8155600101614822565b81516001600160401b0381111561484e5761484e6147a7565b6148628161485c84546145d9565b846147f1565b6020601f821160018114614894575f831561487d5750848201515b5f19600385901b1c1916600184901b17845561240d565b5f84815260208120601f198516915b828110156148c357878501518255602094850194600190920191016148a3565b50848210156148e057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b604081525f6149016040830185614292565b9050826020830152939250505056fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0065d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0031323334353637383941424344454647484a4b4c4d4e505152535455565758595a6162636465666768696a6b6d6e6f707172737475767778797a1cf336fddcc7dc48127faf7a5b80ee54fce73ef647eecd31c24bb6cce3ac3eefa2646970667358221220b0ea28e28dcdd16ab59a8229a8af12f391c65203d532121af5cdb4efde49689f64736f6c634300081b0033",
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
