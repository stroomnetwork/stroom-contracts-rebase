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

// BitcoinUtilsWrapperMetaData contains all meta data concerning the BitcoinUtilsWrapper contract.
var BitcoinUtilsWrapperMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BECH32_ALPHABET_MAP\",\"inputs\":[{\"name\":\"char\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"alphabetCheck\",\"inputs\":[{\"name\":\"BTCAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"equalBytes\",\"inputs\":[{\"name\":\"one\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"two\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"polymodStep\",\"inputs\":[{\"name\":\"pre\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"prefixChk\",\"inputs\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"validateBase58Checksum\",\"inputs\":[{\"name\":\"btcAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateBech32Checksum\",\"inputs\":[{\"name\":\"btcAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateBitcoinAddress\",\"inputs\":[{\"name\":\"network\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"},{\"name\":\"BTCAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506108a08061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c80639647ea37116100585780639647ea371461010a578063af7b81701461011d578063d354c65d14610130578063d3b6d7f514610143575f5ffd5b80630b2aeb6c146100895780634cac70ff146100b15780637b8d3cb4146100c457806380cf79c8146100e9575b5f5ffd5b61009c6100973660046104cd565b610156565b60405190151581526020015b60405180910390f35b61009c6100bf3660046105ca565b6101ea565b6100d76100d236600461062f565b61026b565b60405160ff90911681526020016100a8565b6100fc6100f7366004610656565b6102ee565b6040519081526020016100a8565b61009c61011836600461066d565b610363565b61009c61012b3660046106ac565b61039e565b61009c61013e3660046106de565b610416565b6100fc6101513660046106ac565b61044f565b5f83600381111561016957610169610723565b604051632ce1e0f560e01b815273__$113774c24a411e022512c75ffd4c4add5e$__91632ce1e0f5916101a391908790879060040161075f565b602060405180830381865af41580156101be573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e2919061079d565b949350505050565b604051634cac70ff60e01b81525f9073__$113774c24a411e022512c75ffd4c4add5e$__90634cac70ff9061022590869086906004016107ea565b602060405180830381865af4158015610240573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610264919061079d565b9392505050565b604051631ee34f2d60e21b81526001600160f81b0319821660048201525f9073__$113774c24a411e022512c75ffd4c4add5e$__90637b8d3cb490602401602060405180830381865af41580156102c4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102e8919061080e565b92915050565b604051631019ef3960e31b8152600481018290525f9073__$113774c24a411e022512c75ffd4c4add5e$__906380cf79c8906024015b602060405180830381865af415801561033f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102e8919061082e565b604051639647ea3760e01b81525f9073__$113774c24a411e022512c75ffd4c4add5e$__90639647ea37906102259086908690600401610845565b604051630af7b81760e41b81525f9073__$113774c24a411e022512c75ffd4c4add5e$__9063af7b8170906103d7908590600401610858565b602060405180830381865af41580156103f2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102e8919061079d565b60405163d354c65d60e01b81525f9073__$113774c24a411e022512c75ffd4c4add5e$__9063d354c65d906103d7908590600401610858565b60405163d3b6d7f560e01b81525f9073__$113774c24a411e022512c75ffd4c4add5e$__9063d3b6d7f590610324908590600401610858565b5f5f83601f840112610498575f5ffd5b50813567ffffffffffffffff8111156104af575f5ffd5b6020830191508360208285010111156104c6575f5ffd5b9250929050565b5f5f5f604084860312156104df575f5ffd5b8335600481106104ed575f5ffd5b9250602084013567ffffffffffffffff811115610508575f5ffd5b61051486828701610488565b9497909650939450505050565b634e487b7160e01b5f52604160045260245ffd5b5f5f67ffffffffffffffff84111561054f5761054f610521565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561057e5761057e610521565b604052838152905080828401851015610595575f5ffd5b838360208301375f60208583010152509392505050565b5f82601f8301126105bb575f5ffd5b61026483833560208501610535565b5f5f604083850312156105db575f5ffd5b823567ffffffffffffffff8111156105f1575f5ffd5b6105fd858286016105ac565b925050602083013567ffffffffffffffff811115610619575f5ffd5b610625858286016105ac565b9150509250929050565b5f6020828403121561063f575f5ffd5b81356001600160f81b031981168114610264575f5ffd5b5f60208284031215610666575f5ffd5b5035919050565b5f5f6020838503121561067e575f5ffd5b823567ffffffffffffffff811115610694575f5ffd5b6106a085828601610488565b90969095509350505050565b5f602082840312156106bc575f5ffd5b813567ffffffffffffffff8111156106d2575f5ffd5b6101e2848285016105ac565b5f602082840312156106ee575f5ffd5b813567ffffffffffffffff811115610704575f5ffd5b8201601f81018413610714575f5ffd5b6101e284823560208401610535565b634e487b7160e01b5f52602160045260245ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f6004851061077c57634e487b7160e01b5f52602160045260245ffd5b84825260406020830152610794604083018486610737565b95945050505050565b5f602082840312156107ad575f5ffd5b81518015158114610264575f5ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b604081525f6107fc60408301856107bc565b828103602084015261079481856107bc565b5f6020828403121561081e575f5ffd5b815160ff81168114610264575f5ffd5b5f6020828403121561083e575f5ffd5b5051919050565b602081525f6101e2602083018486610737565b602081525f61026460208301846107bc56fea264697066735822122097ca26081426d4ecd5b75c6827dc96e8c9d20889ab108661fad7358d8a0c42fe64736f6c634300081b0033",
}

// BitcoinUtilsWrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use BitcoinUtilsWrapperMetaData.ABI instead.
var BitcoinUtilsWrapperABI = BitcoinUtilsWrapperMetaData.ABI

// BitcoinUtilsWrapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BitcoinUtilsWrapperMetaData.Bin instead.
var BitcoinUtilsWrapperBin = BitcoinUtilsWrapperMetaData.Bin

// DeployBitcoinUtilsWrapper deploys a new Ethereum contract, binding an instance of BitcoinUtilsWrapper to it.
func DeployBitcoinUtilsWrapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BitcoinUtilsWrapper, error) {
	parsed, err := BitcoinUtilsWrapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BitcoinUtilsWrapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BitcoinUtilsWrapper{BitcoinUtilsWrapperCaller: BitcoinUtilsWrapperCaller{contract: contract}, BitcoinUtilsWrapperTransactor: BitcoinUtilsWrapperTransactor{contract: contract}, BitcoinUtilsWrapperFilterer: BitcoinUtilsWrapperFilterer{contract: contract}}, nil
}

// BitcoinUtilsWrapper is an auto generated Go binding around an Ethereum contract.
type BitcoinUtilsWrapper struct {
	BitcoinUtilsWrapperCaller     // Read-only binding to the contract
	BitcoinUtilsWrapperTransactor // Write-only binding to the contract
	BitcoinUtilsWrapperFilterer   // Log filterer for contract events
}

// BitcoinUtilsWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitcoinUtilsWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitcoinUtilsWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitcoinUtilsWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitcoinUtilsWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitcoinUtilsWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitcoinUtilsWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitcoinUtilsWrapperSession struct {
	Contract     *BitcoinUtilsWrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BitcoinUtilsWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitcoinUtilsWrapperCallerSession struct {
	Contract *BitcoinUtilsWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BitcoinUtilsWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitcoinUtilsWrapperTransactorSession struct {
	Contract     *BitcoinUtilsWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BitcoinUtilsWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitcoinUtilsWrapperRaw struct {
	Contract *BitcoinUtilsWrapper // Generic contract binding to access the raw methods on
}

// BitcoinUtilsWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitcoinUtilsWrapperCallerRaw struct {
	Contract *BitcoinUtilsWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// BitcoinUtilsWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitcoinUtilsWrapperTransactorRaw struct {
	Contract *BitcoinUtilsWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBitcoinUtilsWrapper creates a new instance of BitcoinUtilsWrapper, bound to a specific deployed contract.
func NewBitcoinUtilsWrapper(address common.Address, backend bind.ContractBackend) (*BitcoinUtilsWrapper, error) {
	contract, err := bindBitcoinUtilsWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BitcoinUtilsWrapper{BitcoinUtilsWrapperCaller: BitcoinUtilsWrapperCaller{contract: contract}, BitcoinUtilsWrapperTransactor: BitcoinUtilsWrapperTransactor{contract: contract}, BitcoinUtilsWrapperFilterer: BitcoinUtilsWrapperFilterer{contract: contract}}, nil
}

// NewBitcoinUtilsWrapperCaller creates a new read-only instance of BitcoinUtilsWrapper, bound to a specific deployed contract.
func NewBitcoinUtilsWrapperCaller(address common.Address, caller bind.ContractCaller) (*BitcoinUtilsWrapperCaller, error) {
	contract, err := bindBitcoinUtilsWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitcoinUtilsWrapperCaller{contract: contract}, nil
}

// NewBitcoinUtilsWrapperTransactor creates a new write-only instance of BitcoinUtilsWrapper, bound to a specific deployed contract.
func NewBitcoinUtilsWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*BitcoinUtilsWrapperTransactor, error) {
	contract, err := bindBitcoinUtilsWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitcoinUtilsWrapperTransactor{contract: contract}, nil
}

// NewBitcoinUtilsWrapperFilterer creates a new log filterer instance of BitcoinUtilsWrapper, bound to a specific deployed contract.
func NewBitcoinUtilsWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*BitcoinUtilsWrapperFilterer, error) {
	contract, err := bindBitcoinUtilsWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitcoinUtilsWrapperFilterer{contract: contract}, nil
}

// bindBitcoinUtilsWrapper binds a generic wrapper to an already deployed contract.
func bindBitcoinUtilsWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BitcoinUtilsWrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitcoinUtilsWrapper.Contract.BitcoinUtilsWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitcoinUtilsWrapper.Contract.BitcoinUtilsWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitcoinUtilsWrapper.Contract.BitcoinUtilsWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitcoinUtilsWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitcoinUtilsWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitcoinUtilsWrapper.Contract.contract.Transact(opts, method, params...)
}

// BECH32ALPHABETMAP is a free data retrieval call binding the contract method 0x7b8d3cb4.
//
// Solidity: function BECH32_ALPHABET_MAP(bytes1 char) view returns(uint8)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCaller) BECH32ALPHABETMAP(opts *bind.CallOpts, char [1]byte) (uint8, error) {
	var out []interface{}
	err := _BitcoinUtilsWrapper.contract.Call(opts, &out, "BECH32_ALPHABET_MAP", char)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BECH32ALPHABETMAP is a free data retrieval call binding the contract method 0x7b8d3cb4.
//
// Solidity: function BECH32_ALPHABET_MAP(bytes1 char) view returns(uint8)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperSession) BECH32ALPHABETMAP(char [1]byte) (uint8, error) {
	return _BitcoinUtilsWrapper.Contract.BECH32ALPHABETMAP(&_BitcoinUtilsWrapper.CallOpts, char)
}

// BECH32ALPHABETMAP is a free data retrieval call binding the contract method 0x7b8d3cb4.
//
// Solidity: function BECH32_ALPHABET_MAP(bytes1 char) view returns(uint8)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCallerSession) BECH32ALPHABETMAP(char [1]byte) (uint8, error) {
	return _BitcoinUtilsWrapper.Contract.BECH32ALPHABETMAP(&_BitcoinUtilsWrapper.CallOpts, char)
}

// AlphabetCheck is a free data retrieval call binding the contract method 0xaf7b8170.
//
// Solidity: function alphabetCheck(bytes BTCAddress) pure returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCaller) AlphabetCheck(opts *bind.CallOpts, BTCAddress []byte) (bool, error) {
	var out []interface{}
	err := _BitcoinUtilsWrapper.contract.Call(opts, &out, "alphabetCheck", BTCAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlphabetCheck is a free data retrieval call binding the contract method 0xaf7b8170.
//
// Solidity: function alphabetCheck(bytes BTCAddress) pure returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperSession) AlphabetCheck(BTCAddress []byte) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.AlphabetCheck(&_BitcoinUtilsWrapper.CallOpts, BTCAddress)
}

// AlphabetCheck is a free data retrieval call binding the contract method 0xaf7b8170.
//
// Solidity: function alphabetCheck(bytes BTCAddress) pure returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCallerSession) AlphabetCheck(BTCAddress []byte) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.AlphabetCheck(&_BitcoinUtilsWrapper.CallOpts, BTCAddress)
}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes one, bytes two) pure returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCaller) EqualBytes(opts *bind.CallOpts, one []byte, two []byte) (bool, error) {
	var out []interface{}
	err := _BitcoinUtilsWrapper.contract.Call(opts, &out, "equalBytes", one, two)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes one, bytes two) pure returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperSession) EqualBytes(one []byte, two []byte) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.EqualBytes(&_BitcoinUtilsWrapper.CallOpts, one, two)
}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes one, bytes two) pure returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCallerSession) EqualBytes(one []byte, two []byte) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.EqualBytes(&_BitcoinUtilsWrapper.CallOpts, one, two)
}

// PolymodStep is a free data retrieval call binding the contract method 0x80cf79c8.
//
// Solidity: function polymodStep(uint256 pre) pure returns(uint256)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCaller) PolymodStep(opts *bind.CallOpts, pre *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BitcoinUtilsWrapper.contract.Call(opts, &out, "polymodStep", pre)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PolymodStep is a free data retrieval call binding the contract method 0x80cf79c8.
//
// Solidity: function polymodStep(uint256 pre) pure returns(uint256)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperSession) PolymodStep(pre *big.Int) (*big.Int, error) {
	return _BitcoinUtilsWrapper.Contract.PolymodStep(&_BitcoinUtilsWrapper.CallOpts, pre)
}

// PolymodStep is a free data retrieval call binding the contract method 0x80cf79c8.
//
// Solidity: function polymodStep(uint256 pre) pure returns(uint256)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCallerSession) PolymodStep(pre *big.Int) (*big.Int, error) {
	return _BitcoinUtilsWrapper.Contract.PolymodStep(&_BitcoinUtilsWrapper.CallOpts, pre)
}

// PrefixChk is a free data retrieval call binding the contract method 0xd3b6d7f5.
//
// Solidity: function prefixChk(bytes prefix) pure returns(uint256)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCaller) PrefixChk(opts *bind.CallOpts, prefix []byte) (*big.Int, error) {
	var out []interface{}
	err := _BitcoinUtilsWrapper.contract.Call(opts, &out, "prefixChk", prefix)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrefixChk is a free data retrieval call binding the contract method 0xd3b6d7f5.
//
// Solidity: function prefixChk(bytes prefix) pure returns(uint256)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperSession) PrefixChk(prefix []byte) (*big.Int, error) {
	return _BitcoinUtilsWrapper.Contract.PrefixChk(&_BitcoinUtilsWrapper.CallOpts, prefix)
}

// PrefixChk is a free data retrieval call binding the contract method 0xd3b6d7f5.
//
// Solidity: function prefixChk(bytes prefix) pure returns(uint256)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCallerSession) PrefixChk(prefix []byte) (*big.Int, error) {
	return _BitcoinUtilsWrapper.Contract.PrefixChk(&_BitcoinUtilsWrapper.CallOpts, prefix)
}

// ValidateBase58Checksum is a free data retrieval call binding the contract method 0x9647ea37.
//
// Solidity: function validateBase58Checksum(string btcAddress) view returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCaller) ValidateBase58Checksum(opts *bind.CallOpts, btcAddress string) (bool, error) {
	var out []interface{}
	err := _BitcoinUtilsWrapper.contract.Call(opts, &out, "validateBase58Checksum", btcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateBase58Checksum is a free data retrieval call binding the contract method 0x9647ea37.
//
// Solidity: function validateBase58Checksum(string btcAddress) view returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperSession) ValidateBase58Checksum(btcAddress string) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.ValidateBase58Checksum(&_BitcoinUtilsWrapper.CallOpts, btcAddress)
}

// ValidateBase58Checksum is a free data retrieval call binding the contract method 0x9647ea37.
//
// Solidity: function validateBase58Checksum(string btcAddress) view returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCallerSession) ValidateBase58Checksum(btcAddress string) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.ValidateBase58Checksum(&_BitcoinUtilsWrapper.CallOpts, btcAddress)
}

// ValidateBech32Checksum is a free data retrieval call binding the contract method 0xd354c65d.
//
// Solidity: function validateBech32Checksum(string btcAddress) view returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCaller) ValidateBech32Checksum(opts *bind.CallOpts, btcAddress string) (bool, error) {
	var out []interface{}
	err := _BitcoinUtilsWrapper.contract.Call(opts, &out, "validateBech32Checksum", btcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateBech32Checksum is a free data retrieval call binding the contract method 0xd354c65d.
//
// Solidity: function validateBech32Checksum(string btcAddress) view returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperSession) ValidateBech32Checksum(btcAddress string) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.ValidateBech32Checksum(&_BitcoinUtilsWrapper.CallOpts, btcAddress)
}

// ValidateBech32Checksum is a free data retrieval call binding the contract method 0xd354c65d.
//
// Solidity: function validateBech32Checksum(string btcAddress) view returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCallerSession) ValidateBech32Checksum(btcAddress string) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.ValidateBech32Checksum(&_BitcoinUtilsWrapper.CallOpts, btcAddress)
}

// ValidateBitcoinAddress is a free data retrieval call binding the contract method 0x0b2aeb6c.
//
// Solidity: function validateBitcoinAddress(uint8 network, string BTCAddress) view returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCaller) ValidateBitcoinAddress(opts *bind.CallOpts, network uint8, BTCAddress string) (bool, error) {
	var out []interface{}
	err := _BitcoinUtilsWrapper.contract.Call(opts, &out, "validateBitcoinAddress", network, BTCAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateBitcoinAddress is a free data retrieval call binding the contract method 0x0b2aeb6c.
//
// Solidity: function validateBitcoinAddress(uint8 network, string BTCAddress) view returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperSession) ValidateBitcoinAddress(network uint8, BTCAddress string) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.ValidateBitcoinAddress(&_BitcoinUtilsWrapper.CallOpts, network, BTCAddress)
}

// ValidateBitcoinAddress is a free data retrieval call binding the contract method 0x0b2aeb6c.
//
// Solidity: function validateBitcoinAddress(uint8 network, string BTCAddress) view returns(bool)
func (_BitcoinUtilsWrapper *BitcoinUtilsWrapperCallerSession) ValidateBitcoinAddress(network uint8, BTCAddress string) (bool, error) {
	return _BitcoinUtilsWrapper.Contract.ValidateBitcoinAddress(&_BitcoinUtilsWrapper.CallOpts, network, BTCAddress)
}
