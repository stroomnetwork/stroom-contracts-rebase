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

// StBTCMintInvoice is an auto generated low-level Go binding around an user-defined struct.
type StBTCMintInvoice struct {
	BtcDepositId [32]byte
	Recipient    common.Address
	Amount       *big.Int
}

// StBtcMetaData contains all meta data concerning the StBtc contract.
var StBtcMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DUST_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_MINT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_UPDATE_TOTAL_SUPPLY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"btcDepositIds\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeInvoice\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"encodeTotalSupplyUpdate\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMintInvoiceHash\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPooledBTCByShares\",\"inputs\":[{\"name\":\"sharesAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShares\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSharesByPooledBTC\",\"inputs\":[{\"name\":\"btcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalSupplyUpdateHash\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_network\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"},{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minWithdrawAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstBTC.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintRewards\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"network\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"BTCAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWithdrawAmount\",\"inputs\":[{\"name\":\"_minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupplyUpdateNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"px\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"s\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"m\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintBtcEvent\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_btcDepositId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedeemBtcEvent\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_BTCAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalSupplyUpdatedEvent\",\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalBTCSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AmountBelowMinWithdraw\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DeltaIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBTCAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSharesOrPooledBTC\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSupplyNonce\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinWithdrawTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAlreadyProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountTooBig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintToContractAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UpdateAlreadyProcessed\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557612191908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c908163026034f0146117be5750806306fdde03146116fc578063095ea7b31461167b5780630fde6e551461164a57806318160ddd1461162c57806321172f3b1461145c57806323b872dd1461137957806324b76fd51461116d5780632792949d1461114e578063313ce567146111325780633a98ef39146111145780633f4ba83a146110915780634219347314611074578063457e1a49146110565780635187599d14610b2d57806355573c7714610ad85780635abdb0dc14610a9d5780635c975abb14610a6d578063603c6a6714610a175780636739afca146109ec57806370a08231146109c9578063715018a61461095f57806380239a04146109465780638456cb59146108d25780638c659bf2146108b85780638da5cb5b1461088257806395d89b411461079b5780639fd956871461050b578063a7ce4565146104ed578063a9059cbb146104bc578063af7f131f1461049e578063c4d66de814610442578063c87b7a231461041c578063dd62ed3e146103d3578063e07fbd00146102fb578063e0a04610146102cb578063e737557f1461026a578063f04da65b14610230578063f2fde38b146102055763f376ebbb146101d757600080fd5b34610200576000366003190112610200576000546040516001600160a01b039091168152602090f35b600080fd5b346102005760203660031901126102005761022e610221611833565b610229611f26565b611d9f565b005b34610200576020366003190112610200576001600160a01b03610251611833565b1660005260076020526020604060002054604051908152f35b34610200576000366003190112610200576102c760405161028c60408261188d565b601a8152795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b602082015260405191829160208352602083019061180e565b0390f35b34610200576060366003190112610200576102c76102e7611d48565b60405191829160208352602083019061180e565b34610200576060366003190112610200576000610371602060018060a01b038354166040519061032c60408361188d565b60138252725354524f4f4d5f4d494e545f494e564f49434560681b83830152610353611d48565b604051635e50560360e11b8152948593849283929160048401611c36565b03915afa9081156103c857829161038e575b602082604051908152f35b90506020813d6020116103c0575b816103a96020938361188d565b810103126103bc57602091505182610383565b5080fd5b3d915061039c565b6040513d84823e3d90fd5b34610200576040366003190112610200576103ec611833565b6103fd6103f7611849565b91611bc7565b9060018060a01b03166000526020526020604060002054604051908152f35b3461020057602036600319011261020057602061043a600435611d17565b604051908152f35b346102005760203660031901126102005761045b611833565b610463611f5c565b61046b611f5c565b610473611f5c565b61047c33611d9f565b600080546001600160a01b0319166001600160a01b0392909216919091179055005b34610200576000366003190112610200576020600354604051908152f35b34610200576040366003190112610200576104e26104d8611833565b6024359033611e8c565b602060405160018152f35b34610200576000366003190112610200576020600254604051908152f35b3461020057366003190160808112610200576060136102005760643567ffffffffffffffff811161020057610546602091369060040161185f565b9190610550611e62565b60405161055e60408261188d565b60138152725354524f4f4d5f4d494e545f494e564f49434560681b83820152610585611d48565b60018060a01b0360005416916105b1604051968795869485946311ee58a960e01b865260048601611b4c565b03915afa90811561078f57600091610760575b501561074f576024356001600160a01b03811690604435908281036102005750600435811561073e57660775f05a07400082101561072d5730831461071c5780600052600660205260ff6040600020541661070b578060005260066020526040600020600160ff1982541617905582156106f5577fb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa46991604091600554801580156106eb575b156106d8576106898380925b61068186600454611b87565b600455611b87565b600555856000526007602052836000209081540190558460007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60208651868152a382519182526020820152a2005b6106896106e484611d17565b8092610675565b5060045415610669565b63ec442f0560e01b600052600060045260246000fd5b637bdb87d160e01b60005260046000fd5b63603acaa960e11b60005260046000fd5b636006659960e11b60005260046000fd5b631cebf66f60e11b60005260046000fd5b636227817160e01b60005260046000fd5b610782915060203d602011610788575b61077a818361188d565b810190611b13565b816105c4565b503d610770565b6040513d6000823e3d90fd5b346102005760003660031901126102005760405160006000805160206120dc833981519152546107ca816118e8565b808452906001811690811561085e57506001146107f2575b6102c7836102e78185038261188d565b6000805160206120dc83398151915260009081527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa939250905b808210610844575090915081016020016102e76107e2565b91926001816020925483858801015201910190929161082c565b60ff191660208086019190915291151560051b840190910191506102e790506107e2565b34610200576000366003190112610200576000805160206120fc833981519152546040516001600160a01b039091168152602090f35b34610200576102c76102e76108cc366118d2565b90611cf0565b34610200576000366003190112610200576108eb611f26565b6108f3611e62565b600160ff1960008051602061211c83398151915254161760008051602061211c833981519152557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b3461020057602061043a610959366118d2565b90611c5b565b3461020057600036600319011261020057610978611f26565b6000805160206120fc83398151915280546001600160a01b031981169091556000906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461020057602036600319011261020057602061043a6109e7611833565b611c00565b3461020057600036600319011261020057602060ff60005460a01c16610a1560405180926118af565bf35b34610200576000366003190112610200576102c7604051610a3960408261188d565b60138152725354524f4f4d5f4d494e545f494e564f49434560681b602082015260405191829160208352602083019061180e565b3461020057600036600319011261020057602060ff60008051602061211c83398151915254166040519015158152f35b3461020057602036600319011261020057600435610ab9611f26565b6102228110610ac757600155005b6343a87e9960e11b60005260046000fd5b346102005760203660031901126102005760055480158015610b23575b610b125761043a602091610b0d600454600435611b94565b611ba7565b6340a3daff60e11b60005260046000fd5b5060045415610af5565b3461020057604036600319011261020057600435600481101561020057610b52611849565b60008051602061213c833981519152549060ff8260401c16159167ffffffffffffffff81168015908161104e575b6001149081611044575b15908161103b575b5061102a5767ffffffffffffffff19811660011760008051602061213c8339815191525582610ffd575b5060405192610bcc60408561188d565b600e84526d29ba3937b7b6902134ba31b7b4b760911b602085015260405193610bf660408661188d565b6005855264737442544360d81b6020860152610c10611f5c565b610c18611f5c565b80519067ffffffffffffffff8211610edd578190610c446000805160206120bc833981519152546118e8565b601f8111610f80575b50602090601f8311600114610efe57600092610ef3575b50508160011b916000199060031b1c1916176000805160206120bc833981519152555b835167ffffffffffffffff8111610edd57610cb06000805160206120dc833981519152546118e8565b601f8111610e6b575b50602094601f8211600114610dea57948192939495600092610ddf575b50508160011b916000199060031b1c1916176000805160206120dc833981519152555b610d01611f5c565b610d09611f5c565b60ff1960008051602061211c833981519152541660008051602061211c83398151915255610d35611f5c565b610d3d611f5c565b610d45611f5c565b610d4e33611d9f565b60008054611b586001556001600160a81b0319166001600160a01b039093169290921760a09190911b60ff60a01b16179055610d8657005b68ff00000000000000001960008051602061213c833981519152541660008051602061213c833981519152557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b015190508580610cd6565b601f198216956000805160206120dc833981519152600052806000209160005b888110610e5357508360019596979810610e3a575b505050811b016000805160206120dc83398151915255610cf9565b015160001960f88460031b161c19169055858080610e1f565b91926020600181928685015181550194019201610e0a565b6000805160206120dc8339815191526000527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa601f830160051c81019160208410610ed3575b601f0160051c01905b818110610ec75750610cb9565b60008155600101610eba565b9091508190610eb1565b634e487b7160e01b600052604160045260246000fd5b015190508680610c64565b6000805160206120bc83398151915260009081528281209350601f198516905b818110610f685750908460019594939210610f4f575b505050811b016000805160206120bc83398151915255610c87565b015160001960f88460031b161c19169055868080610f34565b92936020600181928786015181550195019301610f1e565b6000805160206120bc8339815191526000529091507f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0601f840160051c81019160208510610ff3575b90601f859493920160051c01905b818110610fe45750610c4d565b60008155849350600101610fd7565b9091508190610fc9565b68ffffffffffffffffff1916680100000000000000011760008051602061213c8339815191525583610bbc565b63f92ee8a960e01b60005260046000fd5b90501585610b92565b303b159150610b8a565b849150610b80565b34610200576000366003190112610200576020600154604051908152f35b346102005760003660031901126102005760206040516102228152f35b34610200576000366003190112610200576110aa611f26565b60008051602061211c8339815191525460ff8116156111035760ff191660008051602061211c833981519152557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b60005260046000fd5b34610200576000366003190112610200576020600554604051908152f35b3461020057600036600319011261020057602060405160088152f35b346102005760003660031901126102005760206040516305f5e1008152f35b346102005760403660031901126102005760043560243567ffffffffffffffff8111610200576111a190369060040161185f565b6111ac929192611e62565b600154821061136857600054604051632ce1e0f560e01b8152906111da90600483019060a01c60ff166118af565b60406024820152602081806111f3604482018689611b2b565b038173__$113774c24a411e022512c75ffd4c4add5e$__5af490811561078f57600091611349575b5015611338573315611322576000928261123433611c00565b106113135761124283611d17565b61124e84600454611922565b60045561125d81600554611922565b600555338552600760205260408520908154039055836040518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203392a360025490600182018092116112ff57906112eb7f83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e93928260025560405193606085526060850191611b2b565b93602083015260408201528033930390a280f35b634e487b7160e01b85526011600452602485fd5b631e9acf1760e31b8452600484fd5b634b637e8f60e11b600052600060045260246000fd5b6312d58f2760e21b60005260046000fd5b611362915060203d6020116107885761077a818361188d565b8461121b565b633813eacd60e21b60005260046000fd5b3461020057606036600319011261020057611392611833565b61139a611849565b604435906113a783611bc7565b336000908152602091909152604090205492600184016113cc575b6104e29350611e8c565b82841061143f576001600160a01b03811615611429573315611413576104e2936113f582611bc7565b60018060a01b033316600052602052836040600020910390556113c2565b634a1406b160e11b600052600060045260246000fd5b63e602df0560e01b600052600060045260246000fd5b8284637dc7a0d960e11b6000523360045260245260445260646000fd5b346102005760603660031901126102005760043560243560443567ffffffffffffffff811161020057611495602091369060040161185f565b919061149f611e62565b6040516114ad60408261188d565b601a8152795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b838201526114dd8587611cf0565b60018060a01b036000541691611509604051968795869485946311ee58a960e01b865260048601611b4c565b03915afa90811561078f5760009161160d575b501561074f576003548083036115fc57600181018091116115e65760035580156115d55761154a8183611c5b565b9182600052600660205260ff604060002054166115c4576115a96060927f339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b934409460005260066020526040600020600160ff19825416179055600454611b87565b806004556005549060405192835260208301526040820152a1005b63ce199d1760e01b60005260046000fd5b630aa7fcbb60e21b60005260046000fd5b634e487b7160e01b600052601160045260246000fd5b6336b118d760e11b60005260046000fd5b611626915060203d6020116107885761077a818361188d565b8361151c565b34610200576000366003190112610200576020600454604051908152f35b3461020057608036600319011261020057602061167160643560443560243560043561192f565b6040519015158152f35b3461020057604036600319011261020057611694611833565b602435903315611429576001600160a01b0316908115611413576116b733611bc7565b82600052602052806040600020556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b346102005760003660031901126102005760405160006000805160206120bc8339815191525461172b816118e8565b808452906001811690811561085e5750600114611752576102c7836102e78185038261188d565b6000805160206120bc83398151915260009081527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0939250905b8082106117a4575090915081016020016102e76107e2565b91926001816020925483858801015201910190929161178c565b34610200576020366003190112610200576020906004356000526006825260ff6040600020541615158152f35b60005b8381106117fe5750506000910152565b81810151838201526020016117ee565b90602091611827815180928185528580860191016117eb565b601f01601f1916010190565b600435906001600160a01b038216820361020057565b602435906001600160a01b038216820361020057565b9181601f840112156102005782359167ffffffffffffffff8311610200576020838186019501011161020057565b90601f8019910116810190811067ffffffffffffffff821117610edd57604052565b9060048210156118bc5752565b634e487b7160e01b600052602160045260246000fd5b6040906003190112610200576004359060243590565b90600182811c92168015611918575b602083101461190257565b634e487b7160e01b600052602260045260246000fd5b91607f16916118f7565b919082039182116115e657565b90926401000003d0198210801590611b03575b8015611ae7575b611ade5761195684611e15565b94909415611ad4576119e860009160209360405190858201927f7bb52d7a9fef58323eb1bf7a407db382d2f3f2d81bb1224f49fe518f6d48d37c84527f7bb52d7a9fef58323eb1bf7a407db382d2f3f2d81bb1224f49fe518f6d48d37c6040840152606083015286608083015260a082015260a081526119d760c08261188d565b6040519283928392519283916117eb565b8101039060025afa1561078f5770014551231950b75fc4402da1732fc9bebe1960005106918170014551231950b75fc4402da1732fc9bebe19910970014551231950b75fc4402da1732fc9bebe19039170014551231950b75fc4402da1732fc9bebe1983116115e6578170014551231950b75fc4402da1732fc9bebe19910970014551231950b75fc4402da1732fc9bebe19039070014551231950b75fc4402da1732fc9bebe1982116115e657602092600092608092604051928352601b868401526040830152606082015282805260015afa1561078f576000516001600160a01b0390811691161490565b5050505050600090565b50505050600090565b5070014551231950b75fc4402da1732fc9bebe19831015611949565b506401000003d019841015611942565b90816020910312610200575180151581036102005790565b908060209392818452848401376000828201840152601f01601f1916010190565b9290611b849492611b68611b769260608752606087019061180e565b90858203602087015261180e565b926040818503910152611b2b565b90565b919082018092116115e657565b818102929181159184041417156115e657565b8115611bb1570490565b634e487b7160e01b600052601260045260246000fd5b6001600160a01b031660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b6005548015611c2f57611b849160018060a01b03166000526007602052610b0d60406000205460045490611b94565b5050600090565b9091611c4d611b849360408452604084019061180e565b91602081840391015261180e565b602090611cab9261035360018060a01b03600054169160405193611c8060408661188d565b601a8552795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b86860152611cf0565b03915afa90811561078f57600091611cc1575090565b90506020813d602011611ce8575b81611cdc6020938361188d565b81010312610200575190565b3d9150611ccf565b9060405191602083015260408201523060601b606082015260548152611b8460748261188d565b6005549081158015611d3e575b610b1257611b8491611d3591611b94565b60045490611ba7565b5060045415611d24565b6024356001600160a01b038116810361020057604051906bffffffffffffffffffffffff199060601b166020820152604435603482015260043560548201523060601b607482015260688152611b8460888261188d565b6001600160a01b03168015611dff576000805160206120fc83398151915280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b631e4fbdf760e01b600052600060045260246000fd5b611e1e81611f8a565b91909115611e58576040519160208301918252604083015260408252611e4560608361188d565b905190206001600160a01b031690600190565b5050600090600090565b60ff60008051602061211c8339815191525416611e7b57565b63d93c066560e01b60005260046000fd5b916001600160a01b038316918215611322576001600160a01b03169283156106f55781611eba600092611c00565b10611f1757816020916040611eef7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef95611d17565b91868152600785528181208381540390558781526007855220908154019055604051908152a3565b631e9acf1760e31b8152600490fd5b6000805160206120fc833981519152546001600160a01b03163303611f4757565b63118cdaa760e01b6000523360045260246000fd5b60ff60008051602061213c8339815191525460401c1615611f7957565b631afcd79f60e31b60005260046000fd5b6401000003d019811015611ffd57611fc2906401000003d01990816007816000840908906401000003d0199081818009900908612006565b60018116611fd1575b90600190565b6401000003d019036401000003d019811115611fcb57634e487b7160e01b600052601160045260246000fd5b50600090600090565b80156120b557600190600160ff1b90815b61202057505090565b90916401000003d0199063400000f4600160fe1b0384161515830a9082908009096401000003d0199063400000f4600160fe1b03600185901c161515830a9082908009096401000003d0199063400000f4600160fe1b03600285901c161515830a9082908009096401000003d0199063400000f4600160fe1b03600385901c161515830a9082908009099160041c9081612017565b5060009056fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace049016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212203b1fe4d452a689088862cd391202302f89187bab34bb94b34912ae3bc39394f864736f6c634300081b0033",
}

// StBtcABI is the input ABI used to generate the binding from.
// Deprecated: Use StBtcMetaData.ABI instead.
var StBtcABI = StBtcMetaData.ABI

// StBtcBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StBtcMetaData.Bin instead.
var StBtcBin = StBtcMetaData.Bin

// DeployStBtc deploys a new Ethereum contract, binding an instance of StBtc to it.
func DeployStBtc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StBtc, error) {
	parsed, err := StBtcMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StBtcBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StBtc{StBtcCaller: StBtcCaller{contract: contract}, StBtcTransactor: StBtcTransactor{contract: contract}, StBtcFilterer: StBtcFilterer{contract: contract}}, nil
}

// StBtc is an auto generated Go binding around an Ethereum contract.
type StBtc struct {
	StBtcCaller     // Read-only binding to the contract
	StBtcTransactor // Write-only binding to the contract
	StBtcFilterer   // Log filterer for contract events
}

// StBtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type StBtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StBtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StBtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StBtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StBtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StBtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StBtcSession struct {
	Contract     *StBtc            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StBtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StBtcCallerSession struct {
	Contract *StBtcCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StBtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StBtcTransactorSession struct {
	Contract     *StBtcTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StBtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type StBtcRaw struct {
	Contract *StBtc // Generic contract binding to access the raw methods on
}

// StBtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StBtcCallerRaw struct {
	Contract *StBtcCaller // Generic read-only contract binding to access the raw methods on
}

// StBtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StBtcTransactorRaw struct {
	Contract *StBtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStBtc creates a new instance of StBtc, bound to a specific deployed contract.
func NewStBtc(address common.Address, backend bind.ContractBackend) (*StBtc, error) {
	contract, err := bindStBtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StBtc{StBtcCaller: StBtcCaller{contract: contract}, StBtcTransactor: StBtcTransactor{contract: contract}, StBtcFilterer: StBtcFilterer{contract: contract}}, nil
}

// NewStBtcCaller creates a new read-only instance of StBtc, bound to a specific deployed contract.
func NewStBtcCaller(address common.Address, caller bind.ContractCaller) (*StBtcCaller, error) {
	contract, err := bindStBtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StBtcCaller{contract: contract}, nil
}

// NewStBtcTransactor creates a new write-only instance of StBtc, bound to a specific deployed contract.
func NewStBtcTransactor(address common.Address, transactor bind.ContractTransactor) (*StBtcTransactor, error) {
	contract, err := bindStBtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StBtcTransactor{contract: contract}, nil
}

// NewStBtcFilterer creates a new log filterer instance of StBtc, bound to a specific deployed contract.
func NewStBtcFilterer(address common.Address, filterer bind.ContractFilterer) (*StBtcFilterer, error) {
	contract, err := bindStBtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StBtcFilterer{contract: contract}, nil
}

// bindStBtc binds a generic wrapper to an already deployed contract.
func bindStBtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StBtcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StBtc *StBtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StBtc.Contract.StBtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StBtc *StBtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StBtc.Contract.StBtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StBtc *StBtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StBtc.Contract.StBtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StBtc *StBtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StBtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StBtc *StBtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StBtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StBtc *StBtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StBtc.Contract.contract.Transact(opts, method, params...)
}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_StBtc *StBtcCaller) BTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "BTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_StBtc *StBtcSession) BTC() (*big.Int, error) {
	return _StBtc.Contract.BTC(&_StBtc.CallOpts)
}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_StBtc *StBtcCallerSession) BTC() (*big.Int, error) {
	return _StBtc.Contract.BTC(&_StBtc.CallOpts)
}

// DUSTLIMIT is a free data retrieval call binding the contract method 0x42193473.
//
// Solidity: function DUST_LIMIT() view returns(uint256)
func (_StBtc *StBtcCaller) DUSTLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "DUST_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DUSTLIMIT is a free data retrieval call binding the contract method 0x42193473.
//
// Solidity: function DUST_LIMIT() view returns(uint256)
func (_StBtc *StBtcSession) DUSTLIMIT() (*big.Int, error) {
	return _StBtc.Contract.DUSTLIMIT(&_StBtc.CallOpts)
}

// DUSTLIMIT is a free data retrieval call binding the contract method 0x42193473.
//
// Solidity: function DUST_LIMIT() view returns(uint256)
func (_StBtc *StBtcCallerSession) DUSTLIMIT() (*big.Int, error) {
	return _StBtc.Contract.DUSTLIMIT(&_StBtc.CallOpts)
}

// MESSAGEMINT is a free data retrieval call binding the contract method 0x603c6a67.
//
// Solidity: function MESSAGE_MINT() view returns(bytes)
func (_StBtc *StBtcCaller) MESSAGEMINT(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "MESSAGE_MINT")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MESSAGEMINT is a free data retrieval call binding the contract method 0x603c6a67.
//
// Solidity: function MESSAGE_MINT() view returns(bytes)
func (_StBtc *StBtcSession) MESSAGEMINT() ([]byte, error) {
	return _StBtc.Contract.MESSAGEMINT(&_StBtc.CallOpts)
}

// MESSAGEMINT is a free data retrieval call binding the contract method 0x603c6a67.
//
// Solidity: function MESSAGE_MINT() view returns(bytes)
func (_StBtc *StBtcCallerSession) MESSAGEMINT() ([]byte, error) {
	return _StBtc.Contract.MESSAGEMINT(&_StBtc.CallOpts)
}

// MESSAGEUPDATETOTALSUPPLY is a free data retrieval call binding the contract method 0xe737557f.
//
// Solidity: function MESSAGE_UPDATE_TOTAL_SUPPLY() view returns(bytes)
func (_StBtc *StBtcCaller) MESSAGEUPDATETOTALSUPPLY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "MESSAGE_UPDATE_TOTAL_SUPPLY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MESSAGEUPDATETOTALSUPPLY is a free data retrieval call binding the contract method 0xe737557f.
//
// Solidity: function MESSAGE_UPDATE_TOTAL_SUPPLY() view returns(bytes)
func (_StBtc *StBtcSession) MESSAGEUPDATETOTALSUPPLY() ([]byte, error) {
	return _StBtc.Contract.MESSAGEUPDATETOTALSUPPLY(&_StBtc.CallOpts)
}

// MESSAGEUPDATETOTALSUPPLY is a free data retrieval call binding the contract method 0xe737557f.
//
// Solidity: function MESSAGE_UPDATE_TOTAL_SUPPLY() view returns(bytes)
func (_StBtc *StBtcCallerSession) MESSAGEUPDATETOTALSUPPLY() ([]byte, error) {
	return _StBtc.Contract.MESSAGEUPDATETOTALSUPPLY(&_StBtc.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StBtc *StBtcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StBtc *StBtcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StBtc.Contract.Allowance(&_StBtc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StBtc *StBtcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StBtc.Contract.Allowance(&_StBtc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StBtc *StBtcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StBtc *StBtcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StBtc.Contract.BalanceOf(&_StBtc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StBtc *StBtcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StBtc.Contract.BalanceOf(&_StBtc.CallOpts, account)
}

// BtcDepositIds is a free data retrieval call binding the contract method 0x026034f0.
//
// Solidity: function btcDepositIds(bytes32 ) view returns(bool)
func (_StBtc *StBtcCaller) BtcDepositIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "btcDepositIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BtcDepositIds is a free data retrieval call binding the contract method 0x026034f0.
//
// Solidity: function btcDepositIds(bytes32 ) view returns(bool)
func (_StBtc *StBtcSession) BtcDepositIds(arg0 [32]byte) (bool, error) {
	return _StBtc.Contract.BtcDepositIds(&_StBtc.CallOpts, arg0)
}

// BtcDepositIds is a free data retrieval call binding the contract method 0x026034f0.
//
// Solidity: function btcDepositIds(bytes32 ) view returns(bool)
func (_StBtc *StBtcCallerSession) BtcDepositIds(arg0 [32]byte) (bool, error) {
	return _StBtc.Contract.BtcDepositIds(&_StBtc.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StBtc *StBtcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StBtc *StBtcSession) Decimals() (uint8, error) {
	return _StBtc.Contract.Decimals(&_StBtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StBtc *StBtcCallerSession) Decimals() (uint8, error) {
	return _StBtc.Contract.Decimals(&_StBtc.CallOpts)
}

// EncodeInvoice is a free data retrieval call binding the contract method 0xe0a04610.
//
// Solidity: function encodeInvoice((bytes32,address,uint256) invoice) view returns(bytes)
func (_StBtc *StBtcCaller) EncodeInvoice(opts *bind.CallOpts, invoice StBTCMintInvoice) ([]byte, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "encodeInvoice", invoice)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeInvoice is a free data retrieval call binding the contract method 0xe0a04610.
//
// Solidity: function encodeInvoice((bytes32,address,uint256) invoice) view returns(bytes)
func (_StBtc *StBtcSession) EncodeInvoice(invoice StBTCMintInvoice) ([]byte, error) {
	return _StBtc.Contract.EncodeInvoice(&_StBtc.CallOpts, invoice)
}

// EncodeInvoice is a free data retrieval call binding the contract method 0xe0a04610.
//
// Solidity: function encodeInvoice((bytes32,address,uint256) invoice) view returns(bytes)
func (_StBtc *StBtcCallerSession) EncodeInvoice(invoice StBTCMintInvoice) ([]byte, error) {
	return _StBtc.Contract.EncodeInvoice(&_StBtc.CallOpts, invoice)
}

// EncodeTotalSupplyUpdate is a free data retrieval call binding the contract method 0x8c659bf2.
//
// Solidity: function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) view returns(bytes)
func (_StBtc *StBtcCaller) EncodeTotalSupplyUpdate(opts *bind.CallOpts, nonce *big.Int, delta *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "encodeTotalSupplyUpdate", nonce, delta)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTotalSupplyUpdate is a free data retrieval call binding the contract method 0x8c659bf2.
//
// Solidity: function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) view returns(bytes)
func (_StBtc *StBtcSession) EncodeTotalSupplyUpdate(nonce *big.Int, delta *big.Int) ([]byte, error) {
	return _StBtc.Contract.EncodeTotalSupplyUpdate(&_StBtc.CallOpts, nonce, delta)
}

// EncodeTotalSupplyUpdate is a free data retrieval call binding the contract method 0x8c659bf2.
//
// Solidity: function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) view returns(bytes)
func (_StBtc *StBtcCallerSession) EncodeTotalSupplyUpdate(nonce *big.Int, delta *big.Int) ([]byte, error) {
	return _StBtc.Contract.EncodeTotalSupplyUpdate(&_StBtc.CallOpts, nonce, delta)
}

// GetMintInvoiceHash is a free data retrieval call binding the contract method 0xe07fbd00.
//
// Solidity: function getMintInvoiceHash((bytes32,address,uint256) invoice) view returns(bytes32)
func (_StBtc *StBtcCaller) GetMintInvoiceHash(opts *bind.CallOpts, invoice StBTCMintInvoice) ([32]byte, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "getMintInvoiceHash", invoice)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMintInvoiceHash is a free data retrieval call binding the contract method 0xe07fbd00.
//
// Solidity: function getMintInvoiceHash((bytes32,address,uint256) invoice) view returns(bytes32)
func (_StBtc *StBtcSession) GetMintInvoiceHash(invoice StBTCMintInvoice) ([32]byte, error) {
	return _StBtc.Contract.GetMintInvoiceHash(&_StBtc.CallOpts, invoice)
}

// GetMintInvoiceHash is a free data retrieval call binding the contract method 0xe07fbd00.
//
// Solidity: function getMintInvoiceHash((bytes32,address,uint256) invoice) view returns(bytes32)
func (_StBtc *StBtcCallerSession) GetMintInvoiceHash(invoice StBTCMintInvoice) ([32]byte, error) {
	return _StBtc.Contract.GetMintInvoiceHash(&_StBtc.CallOpts, invoice)
}

// GetPooledBTCByShares is a free data retrieval call binding the contract method 0x55573c77.
//
// Solidity: function getPooledBTCByShares(uint256 sharesAmount) view returns(uint256)
func (_StBtc *StBtcCaller) GetPooledBTCByShares(opts *bind.CallOpts, sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "getPooledBTCByShares", sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledBTCByShares is a free data retrieval call binding the contract method 0x55573c77.
//
// Solidity: function getPooledBTCByShares(uint256 sharesAmount) view returns(uint256)
func (_StBtc *StBtcSession) GetPooledBTCByShares(sharesAmount *big.Int) (*big.Int, error) {
	return _StBtc.Contract.GetPooledBTCByShares(&_StBtc.CallOpts, sharesAmount)
}

// GetPooledBTCByShares is a free data retrieval call binding the contract method 0x55573c77.
//
// Solidity: function getPooledBTCByShares(uint256 sharesAmount) view returns(uint256)
func (_StBtc *StBtcCallerSession) GetPooledBTCByShares(sharesAmount *big.Int) (*big.Int, error) {
	return _StBtc.Contract.GetPooledBTCByShares(&_StBtc.CallOpts, sharesAmount)
}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address account) view returns(uint256)
func (_StBtc *StBtcCaller) GetShares(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "getShares", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address account) view returns(uint256)
func (_StBtc *StBtcSession) GetShares(account common.Address) (*big.Int, error) {
	return _StBtc.Contract.GetShares(&_StBtc.CallOpts, account)
}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address account) view returns(uint256)
func (_StBtc *StBtcCallerSession) GetShares(account common.Address) (*big.Int, error) {
	return _StBtc.Contract.GetShares(&_StBtc.CallOpts, account)
}

// GetSharesByPooledBTC is a free data retrieval call binding the contract method 0xc87b7a23.
//
// Solidity: function getSharesByPooledBTC(uint256 btcAmount) view returns(uint256)
func (_StBtc *StBtcCaller) GetSharesByPooledBTC(opts *bind.CallOpts, btcAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "getSharesByPooledBTC", btcAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledBTC is a free data retrieval call binding the contract method 0xc87b7a23.
//
// Solidity: function getSharesByPooledBTC(uint256 btcAmount) view returns(uint256)
func (_StBtc *StBtcSession) GetSharesByPooledBTC(btcAmount *big.Int) (*big.Int, error) {
	return _StBtc.Contract.GetSharesByPooledBTC(&_StBtc.CallOpts, btcAmount)
}

// GetSharesByPooledBTC is a free data retrieval call binding the contract method 0xc87b7a23.
//
// Solidity: function getSharesByPooledBTC(uint256 btcAmount) view returns(uint256)
func (_StBtc *StBtcCallerSession) GetSharesByPooledBTC(btcAmount *big.Int) (*big.Int, error) {
	return _StBtc.Contract.GetSharesByPooledBTC(&_StBtc.CallOpts, btcAmount)
}

// GetTotalSupplyUpdateHash is a free data retrieval call binding the contract method 0x80239a04.
//
// Solidity: function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) view returns(bytes32)
func (_StBtc *StBtcCaller) GetTotalSupplyUpdateHash(opts *bind.CallOpts, nonce *big.Int, delta *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "getTotalSupplyUpdateHash", nonce, delta)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTotalSupplyUpdateHash is a free data retrieval call binding the contract method 0x80239a04.
//
// Solidity: function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) view returns(bytes32)
func (_StBtc *StBtcSession) GetTotalSupplyUpdateHash(nonce *big.Int, delta *big.Int) ([32]byte, error) {
	return _StBtc.Contract.GetTotalSupplyUpdateHash(&_StBtc.CallOpts, nonce, delta)
}

// GetTotalSupplyUpdateHash is a free data retrieval call binding the contract method 0x80239a04.
//
// Solidity: function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) view returns(bytes32)
func (_StBtc *StBtcCallerSession) GetTotalSupplyUpdateHash(nonce *big.Int, delta *big.Int) ([32]byte, error) {
	return _StBtc.Contract.GetTotalSupplyUpdateHash(&_StBtc.CallOpts, nonce, delta)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_StBtc *StBtcCaller) MinWithdrawAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "minWithdrawAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_StBtc *StBtcSession) MinWithdrawAmount() (*big.Int, error) {
	return _StBtc.Contract.MinWithdrawAmount(&_StBtc.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_StBtc *StBtcCallerSession) MinWithdrawAmount() (*big.Int, error) {
	return _StBtc.Contract.MinWithdrawAmount(&_StBtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StBtc *StBtcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StBtc *StBtcSession) Name() (string, error) {
	return _StBtc.Contract.Name(&_StBtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StBtc *StBtcCallerSession) Name() (string, error) {
	return _StBtc.Contract.Name(&_StBtc.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint8)
func (_StBtc *StBtcCaller) Network(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "network")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint8)
func (_StBtc *StBtcSession) Network() (uint8, error) {
	return _StBtc.Contract.Network(&_StBtc.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint8)
func (_StBtc *StBtcCallerSession) Network() (uint8, error) {
	return _StBtc.Contract.Network(&_StBtc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StBtc *StBtcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StBtc *StBtcSession) Owner() (common.Address, error) {
	return _StBtc.Contract.Owner(&_StBtc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StBtc *StBtcCallerSession) Owner() (common.Address, error) {
	return _StBtc.Contract.Owner(&_StBtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StBtc *StBtcCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StBtc *StBtcSession) Paused() (bool, error) {
	return _StBtc.Contract.Paused(&_StBtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StBtc *StBtcCallerSession) Paused() (bool, error) {
	return _StBtc.Contract.Paused(&_StBtc.CallOpts)
}

// RedeemCounter is a free data retrieval call binding the contract method 0xa7ce4565.
//
// Solidity: function redeemCounter() view returns(uint256)
func (_StBtc *StBtcCaller) RedeemCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "redeemCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemCounter is a free data retrieval call binding the contract method 0xa7ce4565.
//
// Solidity: function redeemCounter() view returns(uint256)
func (_StBtc *StBtcSession) RedeemCounter() (*big.Int, error) {
	return _StBtc.Contract.RedeemCounter(&_StBtc.CallOpts)
}

// RedeemCounter is a free data retrieval call binding the contract method 0xa7ce4565.
//
// Solidity: function redeemCounter() view returns(uint256)
func (_StBtc *StBtcCallerSession) RedeemCounter() (*big.Int, error) {
	return _StBtc.Contract.RedeemCounter(&_StBtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StBtc *StBtcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StBtc *StBtcSession) Symbol() (string, error) {
	return _StBtc.Contract.Symbol(&_StBtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StBtc *StBtcCallerSession) Symbol() (string, error) {
	return _StBtc.Contract.Symbol(&_StBtc.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StBtc *StBtcCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StBtc *StBtcSession) TotalShares() (*big.Int, error) {
	return _StBtc.Contract.TotalShares(&_StBtc.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StBtc *StBtcCallerSession) TotalShares() (*big.Int, error) {
	return _StBtc.Contract.TotalShares(&_StBtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StBtc *StBtcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StBtc *StBtcSession) TotalSupply() (*big.Int, error) {
	return _StBtc.Contract.TotalSupply(&_StBtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StBtc *StBtcCallerSession) TotalSupply() (*big.Int, error) {
	return _StBtc.Contract.TotalSupply(&_StBtc.CallOpts)
}

// TotalSupplyUpdateNonce is a free data retrieval call binding the contract method 0xaf7f131f.
//
// Solidity: function totalSupplyUpdateNonce() view returns(uint256)
func (_StBtc *StBtcCaller) TotalSupplyUpdateNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "totalSupplyUpdateNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyUpdateNonce is a free data retrieval call binding the contract method 0xaf7f131f.
//
// Solidity: function totalSupplyUpdateNonce() view returns(uint256)
func (_StBtc *StBtcSession) TotalSupplyUpdateNonce() (*big.Int, error) {
	return _StBtc.Contract.TotalSupplyUpdateNonce(&_StBtc.CallOpts)
}

// TotalSupplyUpdateNonce is a free data retrieval call binding the contract method 0xaf7f131f.
//
// Solidity: function totalSupplyUpdateNonce() view returns(uint256)
func (_StBtc *StBtcCallerSession) TotalSupplyUpdateNonce() (*big.Int, error) {
	return _StBtc.Contract.TotalSupplyUpdateNonce(&_StBtc.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_StBtc *StBtcCaller) ValidatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "validatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_StBtc *StBtcSession) ValidatorRegistry() (common.Address, error) {
	return _StBtc.Contract.ValidatorRegistry(&_StBtc.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_StBtc *StBtcCallerSession) ValidatorRegistry() (common.Address, error) {
	return _StBtc.Contract.ValidatorRegistry(&_StBtc.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x0fde6e55.
//
// Solidity: function verify(uint256 px, uint256 rx, uint256 s, bytes32 m) pure returns(bool)
func (_StBtc *StBtcCaller) Verify(opts *bind.CallOpts, px *big.Int, rx *big.Int, s *big.Int, m [32]byte) (bool, error) {
	var out []interface{}
	err := _StBtc.contract.Call(opts, &out, "verify", px, rx, s, m)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x0fde6e55.
//
// Solidity: function verify(uint256 px, uint256 rx, uint256 s, bytes32 m) pure returns(bool)
func (_StBtc *StBtcSession) Verify(px *big.Int, rx *big.Int, s *big.Int, m [32]byte) (bool, error) {
	return _StBtc.Contract.Verify(&_StBtc.CallOpts, px, rx, s, m)
}

// Verify is a free data retrieval call binding the contract method 0x0fde6e55.
//
// Solidity: function verify(uint256 px, uint256 rx, uint256 s, bytes32 m) pure returns(bool)
func (_StBtc *StBtcCallerSession) Verify(px *big.Int, rx *big.Int, s *big.Int, m [32]byte) (bool, error) {
	return _StBtc.Contract.Verify(&_StBtc.CallOpts, px, rx, s, m)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StBtc *StBtcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StBtc *StBtcSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StBtc.Contract.Approve(&_StBtc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StBtc *StBtcTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StBtc.Contract.Approve(&_StBtc.TransactOpts, spender, value)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry) returns()
func (_StBtc *StBtcTransactor) Initialize(opts *bind.TransactOpts, _network uint8, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "initialize", _network, _validatorRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry) returns()
func (_StBtc *StBtcSession) Initialize(_network uint8, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StBtc.Contract.Initialize(&_StBtc.TransactOpts, _network, _validatorRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry) returns()
func (_StBtc *StBtcTransactorSession) Initialize(_network uint8, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StBtc.Contract.Initialize(&_StBtc.TransactOpts, _network, _validatorRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorRegistry) returns()
func (_StBtc *StBtcTransactor) Initialize0(opts *bind.TransactOpts, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "initialize0", _validatorRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorRegistry) returns()
func (_StBtc *StBtcSession) Initialize0(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StBtc.Contract.Initialize0(&_StBtc.TransactOpts, _validatorRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorRegistry) returns()
func (_StBtc *StBtcTransactorSession) Initialize0(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StBtc.Contract.Initialize0(&_StBtc.TransactOpts, _validatorRegistry)
}

// Mint is a paid mutator transaction binding the contract method 0x9fd95687.
//
// Solidity: function mint((bytes32,address,uint256) invoice, bytes signature) returns()
func (_StBtc *StBtcTransactor) Mint(opts *bind.TransactOpts, invoice StBTCMintInvoice, signature []byte) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "mint", invoice, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x9fd95687.
//
// Solidity: function mint((bytes32,address,uint256) invoice, bytes signature) returns()
func (_StBtc *StBtcSession) Mint(invoice StBTCMintInvoice, signature []byte) (*types.Transaction, error) {
	return _StBtc.Contract.Mint(&_StBtc.TransactOpts, invoice, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x9fd95687.
//
// Solidity: function mint((bytes32,address,uint256) invoice, bytes signature) returns()
func (_StBtc *StBtcTransactorSession) Mint(invoice StBTCMintInvoice, signature []byte) (*types.Transaction, error) {
	return _StBtc.Contract.Mint(&_StBtc.TransactOpts, invoice, signature)
}

// MintRewards is a paid mutator transaction binding the contract method 0x21172f3b.
//
// Solidity: function mintRewards(uint256 nonce, uint256 delta, bytes signature) returns()
func (_StBtc *StBtcTransactor) MintRewards(opts *bind.TransactOpts, nonce *big.Int, delta *big.Int, signature []byte) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "mintRewards", nonce, delta, signature)
}

// MintRewards is a paid mutator transaction binding the contract method 0x21172f3b.
//
// Solidity: function mintRewards(uint256 nonce, uint256 delta, bytes signature) returns()
func (_StBtc *StBtcSession) MintRewards(nonce *big.Int, delta *big.Int, signature []byte) (*types.Transaction, error) {
	return _StBtc.Contract.MintRewards(&_StBtc.TransactOpts, nonce, delta, signature)
}

// MintRewards is a paid mutator transaction binding the contract method 0x21172f3b.
//
// Solidity: function mintRewards(uint256 nonce, uint256 delta, bytes signature) returns()
func (_StBtc *StBtcTransactorSession) MintRewards(nonce *big.Int, delta *big.Int, signature []byte) (*types.Transaction, error) {
	return _StBtc.Contract.MintRewards(&_StBtc.TransactOpts, nonce, delta, signature)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StBtc *StBtcTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StBtc *StBtcSession) Pause() (*types.Transaction, error) {
	return _StBtc.Contract.Pause(&_StBtc.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StBtc *StBtcTransactorSession) Pause() (*types.Transaction, error) {
	return _StBtc.Contract.Pause(&_StBtc.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 _amount, string BTCAddress) returns()
func (_StBtc *StBtcTransactor) Redeem(opts *bind.TransactOpts, _amount *big.Int, BTCAddress string) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "redeem", _amount, BTCAddress)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 _amount, string BTCAddress) returns()
func (_StBtc *StBtcSession) Redeem(_amount *big.Int, BTCAddress string) (*types.Transaction, error) {
	return _StBtc.Contract.Redeem(&_StBtc.TransactOpts, _amount, BTCAddress)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 _amount, string BTCAddress) returns()
func (_StBtc *StBtcTransactorSession) Redeem(_amount *big.Int, BTCAddress string) (*types.Transaction, error) {
	return _StBtc.Contract.Redeem(&_StBtc.TransactOpts, _amount, BTCAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StBtc *StBtcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StBtc *StBtcSession) RenounceOwnership() (*types.Transaction, error) {
	return _StBtc.Contract.RenounceOwnership(&_StBtc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StBtc *StBtcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StBtc.Contract.RenounceOwnership(&_StBtc.TransactOpts)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x5abdb0dc.
//
// Solidity: function setMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StBtc *StBtcTransactor) SetMinWithdrawAmount(opts *bind.TransactOpts, _minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "setMinWithdrawAmount", _minWithdrawAmount)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x5abdb0dc.
//
// Solidity: function setMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StBtc *StBtcSession) SetMinWithdrawAmount(_minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StBtc.Contract.SetMinWithdrawAmount(&_StBtc.TransactOpts, _minWithdrawAmount)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x5abdb0dc.
//
// Solidity: function setMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StBtc *StBtcTransactorSession) SetMinWithdrawAmount(_minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StBtc.Contract.SetMinWithdrawAmount(&_StBtc.TransactOpts, _minWithdrawAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StBtc *StBtcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StBtc *StBtcSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StBtc.Contract.Transfer(&_StBtc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StBtc *StBtcTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StBtc.Contract.Transfer(&_StBtc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StBtc *StBtcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StBtc *StBtcSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StBtc.Contract.TransferFrom(&_StBtc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StBtc *StBtcTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StBtc.Contract.TransferFrom(&_StBtc.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StBtc *StBtcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StBtc *StBtcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StBtc.Contract.TransferOwnership(&_StBtc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StBtc *StBtcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StBtc.Contract.TransferOwnership(&_StBtc.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StBtc *StBtcTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StBtc.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StBtc *StBtcSession) Unpause() (*types.Transaction, error) {
	return _StBtc.Contract.Unpause(&_StBtc.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StBtc *StBtcTransactorSession) Unpause() (*types.Transaction, error) {
	return _StBtc.Contract.Unpause(&_StBtc.TransactOpts)
}

// StBtcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StBtc contract.
type StBtcApprovalIterator struct {
	Event *StBtcApproval // Event containing the contract specifics and raw log

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
func (it *StBtcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StBtcApproval)
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
		it.Event = new(StBtcApproval)
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
func (it *StBtcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StBtcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StBtcApproval represents a Approval event raised by the StBtc contract.
type StBtcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StBtc *StBtcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StBtcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StBtc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StBtcApprovalIterator{contract: _StBtc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StBtc *StBtcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StBtcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StBtc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StBtcApproval)
				if err := _StBtc.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StBtc *StBtcFilterer) ParseApproval(log types.Log) (*StBtcApproval, error) {
	event := new(StBtcApproval)
	if err := _StBtc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StBtcInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StBtc contract.
type StBtcInitializedIterator struct {
	Event *StBtcInitialized // Event containing the contract specifics and raw log

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
func (it *StBtcInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StBtcInitialized)
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
		it.Event = new(StBtcInitialized)
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
func (it *StBtcInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StBtcInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StBtcInitialized represents a Initialized event raised by the StBtc contract.
type StBtcInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StBtc *StBtcFilterer) FilterInitialized(opts *bind.FilterOpts) (*StBtcInitializedIterator, error) {

	logs, sub, err := _StBtc.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StBtcInitializedIterator{contract: _StBtc.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StBtc *StBtcFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StBtcInitialized) (event.Subscription, error) {

	logs, sub, err := _StBtc.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StBtcInitialized)
				if err := _StBtc.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StBtc *StBtcFilterer) ParseInitialized(log types.Log) (*StBtcInitialized, error) {
	event := new(StBtcInitialized)
	if err := _StBtc.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StBtcMintBtcEventIterator is returned from FilterMintBtcEvent and is used to iterate over the raw logs and unpacked data for MintBtcEvent events raised by the StBtc contract.
type StBtcMintBtcEventIterator struct {
	Event *StBtcMintBtcEvent // Event containing the contract specifics and raw log

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
func (it *StBtcMintBtcEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StBtcMintBtcEvent)
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
		it.Event = new(StBtcMintBtcEvent)
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
func (it *StBtcMintBtcEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StBtcMintBtcEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StBtcMintBtcEvent represents a MintBtcEvent event raised by the StBtc contract.
type StBtcMintBtcEvent struct {
	To           common.Address
	Value        *big.Int
	BtcDepositId [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintBtcEvent is a free log retrieval operation binding the contract event 0xb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa469.
//
// Solidity: event MintBtcEvent(address indexed _to, uint256 _value, bytes32 _btcDepositId)
func (_StBtc *StBtcFilterer) FilterMintBtcEvent(opts *bind.FilterOpts, _to []common.Address) (*StBtcMintBtcEventIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StBtc.contract.FilterLogs(opts, "MintBtcEvent", _toRule)
	if err != nil {
		return nil, err
	}
	return &StBtcMintBtcEventIterator{contract: _StBtc.contract, event: "MintBtcEvent", logs: logs, sub: sub}, nil
}

// WatchMintBtcEvent is a free log subscription operation binding the contract event 0xb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa469.
//
// Solidity: event MintBtcEvent(address indexed _to, uint256 _value, bytes32 _btcDepositId)
func (_StBtc *StBtcFilterer) WatchMintBtcEvent(opts *bind.WatchOpts, sink chan<- *StBtcMintBtcEvent, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StBtc.contract.WatchLogs(opts, "MintBtcEvent", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StBtcMintBtcEvent)
				if err := _StBtc.contract.UnpackLog(event, "MintBtcEvent", log); err != nil {
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
func (_StBtc *StBtcFilterer) ParseMintBtcEvent(log types.Log) (*StBtcMintBtcEvent, error) {
	event := new(StBtcMintBtcEvent)
	if err := _StBtc.contract.UnpackLog(event, "MintBtcEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StBtcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StBtc contract.
type StBtcOwnershipTransferredIterator struct {
	Event *StBtcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StBtcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StBtcOwnershipTransferred)
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
		it.Event = new(StBtcOwnershipTransferred)
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
func (it *StBtcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StBtcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StBtcOwnershipTransferred represents a OwnershipTransferred event raised by the StBtc contract.
type StBtcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StBtc *StBtcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StBtcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StBtc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StBtcOwnershipTransferredIterator{contract: _StBtc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StBtc *StBtcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StBtcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StBtc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StBtcOwnershipTransferred)
				if err := _StBtc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StBtc *StBtcFilterer) ParseOwnershipTransferred(log types.Log) (*StBtcOwnershipTransferred, error) {
	event := new(StBtcOwnershipTransferred)
	if err := _StBtc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StBtcPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StBtc contract.
type StBtcPausedIterator struct {
	Event *StBtcPaused // Event containing the contract specifics and raw log

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
func (it *StBtcPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StBtcPaused)
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
		it.Event = new(StBtcPaused)
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
func (it *StBtcPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StBtcPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StBtcPaused represents a Paused event raised by the StBtc contract.
type StBtcPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StBtc *StBtcFilterer) FilterPaused(opts *bind.FilterOpts) (*StBtcPausedIterator, error) {

	logs, sub, err := _StBtc.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StBtcPausedIterator{contract: _StBtc.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StBtc *StBtcFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StBtcPaused) (event.Subscription, error) {

	logs, sub, err := _StBtc.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StBtcPaused)
				if err := _StBtc.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StBtc *StBtcFilterer) ParsePaused(log types.Log) (*StBtcPaused, error) {
	event := new(StBtcPaused)
	if err := _StBtc.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StBtcRedeemBtcEventIterator is returned from FilterRedeemBtcEvent and is used to iterate over the raw logs and unpacked data for RedeemBtcEvent events raised by the StBtc contract.
type StBtcRedeemBtcEventIterator struct {
	Event *StBtcRedeemBtcEvent // Event containing the contract specifics and raw log

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
func (it *StBtcRedeemBtcEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StBtcRedeemBtcEvent)
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
		it.Event = new(StBtcRedeemBtcEvent)
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
func (it *StBtcRedeemBtcEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StBtcRedeemBtcEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StBtcRedeemBtcEvent represents a RedeemBtcEvent event raised by the StBtc contract.
type StBtcRedeemBtcEvent struct {
	From       common.Address
	BTCAddress string
	Value      *big.Int
	Id         *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedeemBtcEvent is a free log retrieval operation binding the contract event 0x83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e.
//
// Solidity: event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id)
func (_StBtc *StBtcFilterer) FilterRedeemBtcEvent(opts *bind.FilterOpts, _from []common.Address) (*StBtcRedeemBtcEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _StBtc.contract.FilterLogs(opts, "RedeemBtcEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &StBtcRedeemBtcEventIterator{contract: _StBtc.contract, event: "RedeemBtcEvent", logs: logs, sub: sub}, nil
}

// WatchRedeemBtcEvent is a free log subscription operation binding the contract event 0x83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e.
//
// Solidity: event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id)
func (_StBtc *StBtcFilterer) WatchRedeemBtcEvent(opts *bind.WatchOpts, sink chan<- *StBtcRedeemBtcEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _StBtc.contract.WatchLogs(opts, "RedeemBtcEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StBtcRedeemBtcEvent)
				if err := _StBtc.contract.UnpackLog(event, "RedeemBtcEvent", log); err != nil {
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
func (_StBtc *StBtcFilterer) ParseRedeemBtcEvent(log types.Log) (*StBtcRedeemBtcEvent, error) {
	event := new(StBtcRedeemBtcEvent)
	if err := _StBtc.contract.UnpackLog(event, "RedeemBtcEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StBtcTotalSupplyUpdatedEventIterator is returned from FilterTotalSupplyUpdatedEvent and is used to iterate over the raw logs and unpacked data for TotalSupplyUpdatedEvent events raised by the StBtc contract.
type StBtcTotalSupplyUpdatedEventIterator struct {
	Event *StBtcTotalSupplyUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *StBtcTotalSupplyUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StBtcTotalSupplyUpdatedEvent)
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
		it.Event = new(StBtcTotalSupplyUpdatedEvent)
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
func (it *StBtcTotalSupplyUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StBtcTotalSupplyUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StBtcTotalSupplyUpdatedEvent represents a TotalSupplyUpdatedEvent event raised by the StBtc contract.
type StBtcTotalSupplyUpdatedEvent struct {
	Nonce          *big.Int
	TotalBTCSupply *big.Int
	TotalShares    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalSupplyUpdatedEvent is a free log retrieval operation binding the contract event 0x339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b93440.
//
// Solidity: event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares)
func (_StBtc *StBtcFilterer) FilterTotalSupplyUpdatedEvent(opts *bind.FilterOpts) (*StBtcTotalSupplyUpdatedEventIterator, error) {

	logs, sub, err := _StBtc.contract.FilterLogs(opts, "TotalSupplyUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return &StBtcTotalSupplyUpdatedEventIterator{contract: _StBtc.contract, event: "TotalSupplyUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchTotalSupplyUpdatedEvent is a free log subscription operation binding the contract event 0x339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b93440.
//
// Solidity: event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares)
func (_StBtc *StBtcFilterer) WatchTotalSupplyUpdatedEvent(opts *bind.WatchOpts, sink chan<- *StBtcTotalSupplyUpdatedEvent) (event.Subscription, error) {

	logs, sub, err := _StBtc.contract.WatchLogs(opts, "TotalSupplyUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StBtcTotalSupplyUpdatedEvent)
				if err := _StBtc.contract.UnpackLog(event, "TotalSupplyUpdatedEvent", log); err != nil {
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
func (_StBtc *StBtcFilterer) ParseTotalSupplyUpdatedEvent(log types.Log) (*StBtcTotalSupplyUpdatedEvent, error) {
	event := new(StBtcTotalSupplyUpdatedEvent)
	if err := _StBtc.contract.UnpackLog(event, "TotalSupplyUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StBtcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StBtc contract.
type StBtcTransferIterator struct {
	Event *StBtcTransfer // Event containing the contract specifics and raw log

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
func (it *StBtcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StBtcTransfer)
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
		it.Event = new(StBtcTransfer)
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
func (it *StBtcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StBtcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StBtcTransfer represents a Transfer event raised by the StBtc contract.
type StBtcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StBtc *StBtcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StBtcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StBtc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StBtcTransferIterator{contract: _StBtc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StBtc *StBtcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StBtcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StBtc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StBtcTransfer)
				if err := _StBtc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StBtc *StBtcFilterer) ParseTransfer(log types.Log) (*StBtcTransfer, error) {
	event := new(StBtcTransfer)
	if err := _StBtc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StBtcUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StBtc contract.
type StBtcUnpausedIterator struct {
	Event *StBtcUnpaused // Event containing the contract specifics and raw log

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
func (it *StBtcUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StBtcUnpaused)
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
		it.Event = new(StBtcUnpaused)
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
func (it *StBtcUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StBtcUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StBtcUnpaused represents a Unpaused event raised by the StBtc contract.
type StBtcUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StBtc *StBtcFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StBtcUnpausedIterator, error) {

	logs, sub, err := _StBtc.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StBtcUnpausedIterator{contract: _StBtc.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StBtc *StBtcFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StBtcUnpaused) (event.Subscription, error) {

	logs, sub, err := _StBtc.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StBtcUnpaused)
				if err := _StBtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StBtc *StBtcFilterer) ParseUnpaused(log types.Log) (*StBtcUnpaused, error) {
	event := new(StBtcUnpaused)
	if err := _StBtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
