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

// ValidatorRegistryMetaData contains all meta data concerning the ValidatorRegistry contract.
var ValidatorRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSAGE_UPDATE_JOINT_PUBLIC_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMessageHash\",\"inputs\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"jointPublicKey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setJointPublicKey\",\"inputs\":[{\"name\":\"_jointPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setJointPublicKeySigned\",\"inputs\":[{\"name\":\"_jointPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateMessage\",\"inputs\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateMessageHash\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"px\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"s\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"m\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"JointPublicKeyUpdated\",\"inputs\":[{\"name\":\"newjointPublicKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b503380603557604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b603c816041565b506091565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610d69806100a06000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b146101195780639c846f7814610134578063ad75d1fd1461014b578063b4dcbbb714610194578063bca0ac06146101a7578063f2fde38b146101ba57600080fd5b806301b4167e146100ae5780630fde6e55146100c357806311ee58a9146100eb5780631bb449bd146100fe578063715018a614610111575b600080fd5b6100c16100bc366004610911565b6101cd565b005b6100d66100d136600461092a565b610210565b60405190151581526020015b60405180910390f35b6100d66100f9366004610a4a565b610378565b6100d661010c366004610ae1565b61039c565b6100c1610455565b6000546040516001600160a01b0390911681526020016100e2565b61013d60015481565b6040519081526020016100e2565b6101876040518060400160405280601e81526020017f5354524f4f4d5f5550444154455f4a4f494e545f5055424c49435f4b4559000081525081565b6040516100e29190610b51565b6100c16101a2366004610ae1565b610469565b61013d6101b5366004610b84565b610560565b6100c16101c8366004610bed565b61059b565b6101d56105d9565b60018190556040518181527f8c2c4ca5558a22e5df2161c7ba30c7c1da4b0dfd74176f4ae8aa625a146957619060200160405180910390a150565b60006401000003d0198510158061022d57506401000003d0198410155b8061024a575070014551231950b75fc4402da1732fc9bebe198310155b1561025757506000610370565b60008061026386610606565b915091508061027757600092505050610370565b6000610284878987610661565b9050600070014551231950b75fc4402da1732fc9bebe198988096102ba9070014551231950b75fc4402da1732fc9bebe19610c42565b9050600070014551231950b75fc4402da1732fc9bebe198a84096102f09070014551231950b75fc4402da1732fc9bebe19610c42565b60408051600080825260208201808452869052601b92820192909252606081018d9052608081018390529192509060019060a0016020604051602081039080840390855afa158015610346573d6000803e3d6000fd5b505050602060405103519050856001600160a01b0316816001600160a01b03161496505050505050505b949350505050565b6000806103858686610560565b905061039281858561039c565b9695505050505050565b60015460009081036104045760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f7252656769737472793a204e4f5f4a4f494e545f5055424c60448201526549435f4b455960d01b60648201526084015b60405180910390fd5b60006104136020828587610c55565b61041c91610c7f565b9050600061042e604060208688610c55565b61043791610c7f565b60015490915061044990838389610210565b925050505b9392505050565b61045d6105d9565b6104676000610733565b565b6104cb6040518060400160405280601e81526020017f5354524f4f4d5f5550444154455f4a4f494e545f5055424c49435f4b45590000815250846040516020016104b591815260200190565b6040516020818303038152906040528484610378565b6105235760405162461bcd60e51b8152602060048201526024808201527f56616c696461746f7252656769737472793a20494e56414c49445f5349474e416044820152635455524560e01b60648201526084016103fb565b60018390556040518381527f8c2c4ca5558a22e5df2161c7ba30c7c1da4b0dfd74176f4ae8aa625a146957619060200160405180910390a1505050565b600082805190602001208260405160200161057c929190610c9d565b6040516020818303038152906040528051906020012090505b92915050565b6105a36105d9565b6001600160a01b0381166105cd57604051631e4fbdf760e01b8152600060048201526024016103fb565b6105d681610733565b50565b6000546001600160a01b031633146104675760405163118cdaa760e01b81523360048201526024016103fb565b60008060008061061585610783565b915091508061062b575060009485945092505050565b50604080516020808201969096528082019290925280518083038201815260609092019052805193019290922092600192509050565b604080517f7bb52d7a9fef58323eb1bf7a407db382d2f3f2d81bb1224f49fe518f6d48d37c60208201819052918101829052606081018590526080810184905260a081018390526000919070014551231950b75fc4402da1732fc9bebe199060029060c00160408051601f19818403018152908290526106e091610cc3565b602060405180830381855afa1580156106fd573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906107209190610cdf565b61072a9190610cf8565b95945050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806401000003d0198160078286106107a557506000958695509350505050565b600083806107b5576107b5610c16565b84806107c3576107c3610c16565b8386806107d2576107d2610c16565b868b090885806107e4576107e4610c16565b86806107f2576107f2610c16565b8a8b098a0908905061081b81600461080b876001610d0c565b6108159190610d1f565b86610848565b905060006001821615610837576108328286610c42565b610839565b815b98600198509650505050505050565b60008160000361088c5760405162461bcd60e51b815260206004820152600f60248201526e4d6f64756c7573206973207a65726f60881b60448201526064016103fb565b8360000361089c5750600061044e565b826000036108ac5750600161044e565b6001600160ff1b5b801561090857838186161515870a85848509099150836002820486161515870a85848509099150836004820486161515870a85848509099150836008820486161515870a85848509099150601090046108b4565b50949350505050565b60006020828403121561092357600080fd5b5035919050565b6000806000806080858703121561094057600080fd5b5050823594602084013594506040840135936060013592509050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261098357600080fd5b813567ffffffffffffffff81111561099d5761099d61095c565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156109cc576109cc61095c565b6040528181528382016020018510156109e457600080fd5b816020850160208301376000918101602001919091529392505050565b60008083601f840112610a1357600080fd5b50813567ffffffffffffffff811115610a2b57600080fd5b602083019150836020828501011115610a4357600080fd5b9250929050565b60008060008060608587031215610a6057600080fd5b843567ffffffffffffffff811115610a7757600080fd5b610a8387828801610972565b945050602085013567ffffffffffffffff811115610aa057600080fd5b610aac87828801610972565b935050604085013567ffffffffffffffff811115610ac957600080fd5b610ad587828801610a01565b95989497509550505050565b600080600060408486031215610af657600080fd5b83359250602084013567ffffffffffffffff811115610b1457600080fd5b610b2086828701610a01565b9497909650939450505050565b60005b83811015610b48578181015183820152602001610b30565b50506000910152565b6020815260008251806020840152610b70816040850160208701610b2d565b601f01601f19169190910160400192915050565b60008060408385031215610b9757600080fd5b823567ffffffffffffffff811115610bae57600080fd5b610bba85828601610972565b925050602083013567ffffffffffffffff811115610bd757600080fd5b610be385828601610972565b9150509250929050565b600060208284031215610bff57600080fd5b81356001600160a01b038116811461044e57600080fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561059557610595610c2c565b60008085851115610c6557600080fd5b83861115610c7257600080fd5b5050820193919092039150565b8035602083101561059557600019602084900360031b1b1692915050565b82815260008251610cb5816020850160208701610b2d565b919091016020019392505050565b60008251610cd5818460208701610b2d565b9190910192915050565b600060208284031215610cf157600080fd5b5051919050565b600082610d0757610d07610c16565b500690565b8082018082111561059557610595610c2c565b600082610d2e57610d2e610c16565b50049056fea2646970667358221220b77ef08c80ff47649663271ccfe0c3043952a2037794ecd5390507171c41696d64736f6c634300081b0033",
}

// ValidatorRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorRegistryMetaData.ABI instead.
var ValidatorRegistryABI = ValidatorRegistryMetaData.ABI

// ValidatorRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorRegistryMetaData.Bin instead.
var ValidatorRegistryBin = ValidatorRegistryMetaData.Bin

// DeployValidatorRegistry deploys a new Ethereum contract, binding an instance of ValidatorRegistry to it.
func DeployValidatorRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorRegistry, error) {
	parsed, err := ValidatorRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorRegistry{ValidatorRegistryCaller: ValidatorRegistryCaller{contract: contract}, ValidatorRegistryTransactor: ValidatorRegistryTransactor{contract: contract}, ValidatorRegistryFilterer: ValidatorRegistryFilterer{contract: contract}}, nil
}

// ValidatorRegistry is an auto generated Go binding around an Ethereum contract.
type ValidatorRegistry struct {
	ValidatorRegistryCaller     // Read-only binding to the contract
	ValidatorRegistryTransactor // Write-only binding to the contract
	ValidatorRegistryFilterer   // Log filterer for contract events
}

// ValidatorRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorRegistrySession struct {
	Contract     *ValidatorRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorRegistryCallerSession struct {
	Contract *ValidatorRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorRegistryTransactorSession struct {
	Contract     *ValidatorRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRegistryRaw struct {
	Contract *ValidatorRegistry // Generic contract binding to access the raw methods on
}

// ValidatorRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorRegistryCallerRaw struct {
	Contract *ValidatorRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorRegistryTransactorRaw struct {
	Contract *ValidatorRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorRegistry creates a new instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistry(address common.Address, backend bind.ContractBackend) (*ValidatorRegistry, error) {
	contract, err := bindValidatorRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistry{ValidatorRegistryCaller: ValidatorRegistryCaller{contract: contract}, ValidatorRegistryTransactor: ValidatorRegistryTransactor{contract: contract}, ValidatorRegistryFilterer: ValidatorRegistryFilterer{contract: contract}}, nil
}

// NewValidatorRegistryCaller creates a new read-only instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistryCaller(address common.Address, caller bind.ContractCaller) (*ValidatorRegistryCaller, error) {
	contract, err := bindValidatorRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryCaller{contract: contract}, nil
}

// NewValidatorRegistryTransactor creates a new write-only instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorRegistryTransactor, error) {
	contract, err := bindValidatorRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryTransactor{contract: contract}, nil
}

// NewValidatorRegistryFilterer creates a new log filterer instance of ValidatorRegistry, bound to a specific deployed contract.
func NewValidatorRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorRegistryFilterer, error) {
	contract, err := bindValidatorRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryFilterer{contract: contract}, nil
}

// bindValidatorRegistry binds a generic wrapper to an already deployed contract.
func bindValidatorRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRegistry *ValidatorRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorRegistry.Contract.ValidatorRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRegistry *ValidatorRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.ValidatorRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRegistry *ValidatorRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.ValidatorRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRegistry *ValidatorRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRegistry *ValidatorRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRegistry *ValidatorRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.contract.Transact(opts, method, params...)
}

// MESSAGEUPDATEJOINTPUBLICKEY is a free data retrieval call binding the contract method 0xad75d1fd.
//
// Solidity: function MESSAGE_UPDATE_JOINT_PUBLIC_KEY() view returns(bytes)
func (_ValidatorRegistry *ValidatorRegistryCaller) MESSAGEUPDATEJOINTPUBLICKEY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "MESSAGE_UPDATE_JOINT_PUBLIC_KEY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MESSAGEUPDATEJOINTPUBLICKEY is a free data retrieval call binding the contract method 0xad75d1fd.
//
// Solidity: function MESSAGE_UPDATE_JOINT_PUBLIC_KEY() view returns(bytes)
func (_ValidatorRegistry *ValidatorRegistrySession) MESSAGEUPDATEJOINTPUBLICKEY() ([]byte, error) {
	return _ValidatorRegistry.Contract.MESSAGEUPDATEJOINTPUBLICKEY(&_ValidatorRegistry.CallOpts)
}

// MESSAGEUPDATEJOINTPUBLICKEY is a free data retrieval call binding the contract method 0xad75d1fd.
//
// Solidity: function MESSAGE_UPDATE_JOINT_PUBLIC_KEY() view returns(bytes)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) MESSAGEUPDATEJOINTPUBLICKEY() ([]byte, error) {
	return _ValidatorRegistry.Contract.MESSAGEUPDATEJOINTPUBLICKEY(&_ValidatorRegistry.CallOpts)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xbca0ac06.
//
// Solidity: function getMessageHash(bytes prefix, bytes data) pure returns(bytes32)
func (_ValidatorRegistry *ValidatorRegistryCaller) GetMessageHash(opts *bind.CallOpts, prefix []byte, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "getMessageHash", prefix, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xbca0ac06.
//
// Solidity: function getMessageHash(bytes prefix, bytes data) pure returns(bytes32)
func (_ValidatorRegistry *ValidatorRegistrySession) GetMessageHash(prefix []byte, data []byte) ([32]byte, error) {
	return _ValidatorRegistry.Contract.GetMessageHash(&_ValidatorRegistry.CallOpts, prefix, data)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xbca0ac06.
//
// Solidity: function getMessageHash(bytes prefix, bytes data) pure returns(bytes32)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) GetMessageHash(prefix []byte, data []byte) ([32]byte, error) {
	return _ValidatorRegistry.Contract.GetMessageHash(&_ValidatorRegistry.CallOpts, prefix, data)
}

// JointPublicKey is a free data retrieval call binding the contract method 0x9c846f78.
//
// Solidity: function jointPublicKey() view returns(bytes32)
func (_ValidatorRegistry *ValidatorRegistryCaller) JointPublicKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "jointPublicKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// JointPublicKey is a free data retrieval call binding the contract method 0x9c846f78.
//
// Solidity: function jointPublicKey() view returns(bytes32)
func (_ValidatorRegistry *ValidatorRegistrySession) JointPublicKey() ([32]byte, error) {
	return _ValidatorRegistry.Contract.JointPublicKey(&_ValidatorRegistry.CallOpts)
}

// JointPublicKey is a free data retrieval call binding the contract method 0x9c846f78.
//
// Solidity: function jointPublicKey() view returns(bytes32)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) JointPublicKey() ([32]byte, error) {
	return _ValidatorRegistry.Contract.JointPublicKey(&_ValidatorRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorRegistry *ValidatorRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorRegistry *ValidatorRegistrySession) Owner() (common.Address, error) {
	return _ValidatorRegistry.Contract.Owner(&_ValidatorRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) Owner() (common.Address, error) {
	return _ValidatorRegistry.Contract.Owner(&_ValidatorRegistry.CallOpts)
}

// ValidateMessage is a free data retrieval call binding the contract method 0x11ee58a9.
//
// Solidity: function validateMessage(bytes prefix, bytes data, bytes signature) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCaller) ValidateMessage(opts *bind.CallOpts, prefix []byte, data []byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "validateMessage", prefix, data, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateMessage is a free data retrieval call binding the contract method 0x11ee58a9.
//
// Solidity: function validateMessage(bytes prefix, bytes data, bytes signature) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistrySession) ValidateMessage(prefix []byte, data []byte, signature []byte) (bool, error) {
	return _ValidatorRegistry.Contract.ValidateMessage(&_ValidatorRegistry.CallOpts, prefix, data, signature)
}

// ValidateMessage is a free data retrieval call binding the contract method 0x11ee58a9.
//
// Solidity: function validateMessage(bytes prefix, bytes data, bytes signature) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) ValidateMessage(prefix []byte, data []byte, signature []byte) (bool, error) {
	return _ValidatorRegistry.Contract.ValidateMessage(&_ValidatorRegistry.CallOpts, prefix, data, signature)
}

// ValidateMessageHash is a free data retrieval call binding the contract method 0x1bb449bd.
//
// Solidity: function validateMessageHash(bytes32 hash, bytes signature) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCaller) ValidateMessageHash(opts *bind.CallOpts, hash [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "validateMessageHash", hash, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateMessageHash is a free data retrieval call binding the contract method 0x1bb449bd.
//
// Solidity: function validateMessageHash(bytes32 hash, bytes signature) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistrySession) ValidateMessageHash(hash [32]byte, signature []byte) (bool, error) {
	return _ValidatorRegistry.Contract.ValidateMessageHash(&_ValidatorRegistry.CallOpts, hash, signature)
}

// ValidateMessageHash is a free data retrieval call binding the contract method 0x1bb449bd.
//
// Solidity: function validateMessageHash(bytes32 hash, bytes signature) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) ValidateMessageHash(hash [32]byte, signature []byte) (bool, error) {
	return _ValidatorRegistry.Contract.ValidateMessageHash(&_ValidatorRegistry.CallOpts, hash, signature)
}

// Verify is a free data retrieval call binding the contract method 0x0fde6e55.
//
// Solidity: function verify(uint256 px, uint256 rx, uint256 s, bytes32 m) pure returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCaller) Verify(opts *bind.CallOpts, px *big.Int, rx *big.Int, s *big.Int, m [32]byte) (bool, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "verify", px, rx, s, m)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x0fde6e55.
//
// Solidity: function verify(uint256 px, uint256 rx, uint256 s, bytes32 m) pure returns(bool)
func (_ValidatorRegistry *ValidatorRegistrySession) Verify(px *big.Int, rx *big.Int, s *big.Int, m [32]byte) (bool, error) {
	return _ValidatorRegistry.Contract.Verify(&_ValidatorRegistry.CallOpts, px, rx, s, m)
}

// Verify is a free data retrieval call binding the contract method 0x0fde6e55.
//
// Solidity: function verify(uint256 px, uint256 rx, uint256 s, bytes32 m) pure returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) Verify(px *big.Int, rx *big.Int, s *big.Int, m [32]byte) (bool, error) {
	return _ValidatorRegistry.Contract.Verify(&_ValidatorRegistry.CallOpts, px, rx, s, m)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorRegistry *ValidatorRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.RenounceOwnership(&_ValidatorRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.RenounceOwnership(&_ValidatorRegistry.TransactOpts)
}

// SetJointPublicKey is a paid mutator transaction binding the contract method 0x01b4167e.
//
// Solidity: function setJointPublicKey(bytes32 _jointPublicKey) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) SetJointPublicKey(opts *bind.TransactOpts, _jointPublicKey [32]byte) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "setJointPublicKey", _jointPublicKey)
}

// SetJointPublicKey is a paid mutator transaction binding the contract method 0x01b4167e.
//
// Solidity: function setJointPublicKey(bytes32 _jointPublicKey) returns()
func (_ValidatorRegistry *ValidatorRegistrySession) SetJointPublicKey(_jointPublicKey [32]byte) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetJointPublicKey(&_ValidatorRegistry.TransactOpts, _jointPublicKey)
}

// SetJointPublicKey is a paid mutator transaction binding the contract method 0x01b4167e.
//
// Solidity: function setJointPublicKey(bytes32 _jointPublicKey) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) SetJointPublicKey(_jointPublicKey [32]byte) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetJointPublicKey(&_ValidatorRegistry.TransactOpts, _jointPublicKey)
}

// SetJointPublicKeySigned is a paid mutator transaction binding the contract method 0xb4dcbbb7.
//
// Solidity: function setJointPublicKeySigned(bytes32 _jointPublicKey, bytes signature) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) SetJointPublicKeySigned(opts *bind.TransactOpts, _jointPublicKey [32]byte, signature []byte) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "setJointPublicKeySigned", _jointPublicKey, signature)
}

// SetJointPublicKeySigned is a paid mutator transaction binding the contract method 0xb4dcbbb7.
//
// Solidity: function setJointPublicKeySigned(bytes32 _jointPublicKey, bytes signature) returns()
func (_ValidatorRegistry *ValidatorRegistrySession) SetJointPublicKeySigned(_jointPublicKey [32]byte, signature []byte) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetJointPublicKeySigned(&_ValidatorRegistry.TransactOpts, _jointPublicKey, signature)
}

// SetJointPublicKeySigned is a paid mutator transaction binding the contract method 0xb4dcbbb7.
//
// Solidity: function setJointPublicKeySigned(bytes32 _jointPublicKey, bytes signature) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) SetJointPublicKeySigned(_jointPublicKey [32]byte, signature []byte) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetJointPublicKeySigned(&_ValidatorRegistry.TransactOpts, _jointPublicKey, signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorRegistry *ValidatorRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.TransferOwnership(&_ValidatorRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.TransferOwnership(&_ValidatorRegistry.TransactOpts, newOwner)
}

// ValidatorRegistryJointPublicKeyUpdatedIterator is returned from FilterJointPublicKeyUpdated and is used to iterate over the raw logs and unpacked data for JointPublicKeyUpdated events raised by the ValidatorRegistry contract.
type ValidatorRegistryJointPublicKeyUpdatedIterator struct {
	Event *ValidatorRegistryJointPublicKeyUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorRegistryJointPublicKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRegistryJointPublicKeyUpdated)
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
		it.Event = new(ValidatorRegistryJointPublicKeyUpdated)
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
func (it *ValidatorRegistryJointPublicKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRegistryJointPublicKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRegistryJointPublicKeyUpdated represents a JointPublicKeyUpdated event raised by the ValidatorRegistry contract.
type ValidatorRegistryJointPublicKeyUpdated struct {
	NewjointPublicKey [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterJointPublicKeyUpdated is a free log retrieval operation binding the contract event 0x8c2c4ca5558a22e5df2161c7ba30c7c1da4b0dfd74176f4ae8aa625a14695761.
//
// Solidity: event JointPublicKeyUpdated(bytes32 newjointPublicKey)
func (_ValidatorRegistry *ValidatorRegistryFilterer) FilterJointPublicKeyUpdated(opts *bind.FilterOpts) (*ValidatorRegistryJointPublicKeyUpdatedIterator, error) {

	logs, sub, err := _ValidatorRegistry.contract.FilterLogs(opts, "JointPublicKeyUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryJointPublicKeyUpdatedIterator{contract: _ValidatorRegistry.contract, event: "JointPublicKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchJointPublicKeyUpdated is a free log subscription operation binding the contract event 0x8c2c4ca5558a22e5df2161c7ba30c7c1da4b0dfd74176f4ae8aa625a14695761.
//
// Solidity: event JointPublicKeyUpdated(bytes32 newjointPublicKey)
func (_ValidatorRegistry *ValidatorRegistryFilterer) WatchJointPublicKeyUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorRegistryJointPublicKeyUpdated) (event.Subscription, error) {

	logs, sub, err := _ValidatorRegistry.contract.WatchLogs(opts, "JointPublicKeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRegistryJointPublicKeyUpdated)
				if err := _ValidatorRegistry.contract.UnpackLog(event, "JointPublicKeyUpdated", log); err != nil {
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

// ParseJointPublicKeyUpdated is a log parse operation binding the contract event 0x8c2c4ca5558a22e5df2161c7ba30c7c1da4b0dfd74176f4ae8aa625a14695761.
//
// Solidity: event JointPublicKeyUpdated(bytes32 newjointPublicKey)
func (_ValidatorRegistry *ValidatorRegistryFilterer) ParseJointPublicKeyUpdated(log types.Log) (*ValidatorRegistryJointPublicKeyUpdated, error) {
	event := new(ValidatorRegistryJointPublicKeyUpdated)
	if err := _ValidatorRegistry.contract.UnpackLog(event, "JointPublicKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValidatorRegistry contract.
type ValidatorRegistryOwnershipTransferredIterator struct {
	Event *ValidatorRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRegistryOwnershipTransferred)
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
		it.Event = new(ValidatorRegistryOwnershipTransferred)
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
func (it *ValidatorRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ValidatorRegistry contract.
type ValidatorRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorRegistry *ValidatorRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRegistryOwnershipTransferredIterator{contract: _ValidatorRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorRegistry *ValidatorRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRegistryOwnershipTransferred)
				if err := _ValidatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ValidatorRegistry *ValidatorRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorRegistryOwnershipTransferred, error) {
	event := new(ValidatorRegistryOwnershipTransferred)
	if err := _ValidatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
