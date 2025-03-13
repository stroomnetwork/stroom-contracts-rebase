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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSAGE_UPDATE_JOINT_PUBLIC_KEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMessageHash\",\"inputs\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"jointPublicKey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setJointPublicKey\",\"inputs\":[{\"name\":\"_jointPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setJointPublicKeySigned\",\"inputs\":[{\"name\":\"_jointPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_nonce\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usedNonces\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateMessage\",\"inputs\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateMessageHash\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"px\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"s\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"m\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"JointPublicKeyUpdated\",\"inputs\":[{\"name\":\"newjointPublicKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoJointPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonceAlreadyUsed\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608080604052346075573315605f5760008054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3610a6c908161007b8239f35b631e4fbdf760e01b600052600060045260246000fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806301b4167e146104a05780630fde6e551461047957806311ee58a9146103f65780631bb449bd146103ad578063715018a6146103545780638da5cb5b1461032b578063996d1b91146102305780639c846f7814610212578063ad75d1fd146101c8578063bca0ac0614610163578063f2fde38b146100d65763feb61724146100a057600080fd5b346100d15760203660031901126100d1576004356000526002602052602060ff604060002054166040519015158152f35b600080fd5b346100d15760203660031901126100d1576004356001600160a01b038116908190036100d15761010461088f565b801561014d57600080546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b600052600060045260246000fd5b346100d15760403660031901126100d15760043567ffffffffffffffff81116100d157610194903690600401610524565b60243567ffffffffffffffff81116100d1576020916101ba6101c0923690600401610524565b9061084d565b604051908152f35b346100d15760003660031901126100d1576101e16105a9565b6040805180926020825261020481518092816020860152602086860191016105e4565b601f01601f19168101030190f35b346100d15760003660031901126100d1576020600154604051908152f35b346100d15760603660031901126100d15760043560243560443567ffffffffffffffff81116100d15761026790369060040161057b565b9082600052600260205260ff6040600020541661031657906102b7916102b261028e6105a9565b60405190876020830152866040830152604082526102ad6060836104ec565b61084d565b61080d565b1561030557816020917f8c2c4ca5558a22e5df2161c7ba30c7c1da4b0dfd74176f4ae8aa625a1469576193600155600052600282526040600020600160ff19825416179055604051908152a1005b638baa579f60e01b60005260046000fd5b82638c24228560e01b60005260045260246000fd5b346100d15760003660031901126100d1576000546040516001600160a01b039091168152602090f35b346100d15760003660031901126100d15761036d61088f565b600080546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100d15760403660031901126100d15760243567ffffffffffffffff81116100d1576103ec6103e3602092369060040161057b565b9060043561080d565b6040519015158152f35b346100d15760603660031901126100d15760043567ffffffffffffffff81116100d157610427903690600401610524565b60243567ffffffffffffffff81116100d157610447903690600401610524565b6044359067ffffffffffffffff82116100d1576020926102b26104716103ec94369060040161057b565b93909261084d565b346100d15760803660031901126100d15760206103ec606435604435602435600435610607565b346100d15760203660031901126100d1577f8c2c4ca5558a22e5df2161c7ba30c7c1da4b0dfd74176f4ae8aa625a1469576160206004356104df61088f565b80600155604051908152a1005b90601f8019910116810190811067ffffffffffffffff82111761050e57604052565b634e487b7160e01b600052604160045260246000fd5b81601f820112156100d15780359067ffffffffffffffff821161050e5760405192610559601f8401601f1916602001856104ec565b828452602083830101116100d157816000926020809301838601378301015290565b9181601f840112156100d15782359167ffffffffffffffff83116100d157602083818601950101116100d157565b604051906105b86040836104ec565b601e82527f5354524f4f4d5f5550444154455f4a4f494e545f5055424c49435f4b455900006020830152565b60005b8381106105f75750506000910152565b81810151838201526020016105e7565b90926401000003d01982108015906107fd575b80156107e1575b6107d85761062e846108b8565b949094156107ce576106c060009160209360405190858201927f7bb52d7a9fef58323eb1bf7a407db382d2f3f2d81bb1224f49fe518f6d48d37c84527f7bb52d7a9fef58323eb1bf7a407db382d2f3f2d81bb1224f49fe518f6d48d37c6040840152606083015286608083015260a082015260a081526106af60c0826104ec565b6040519283928392519283916105e4565b8101039060025afa156107ac5770014551231950b75fc4402da1732fc9bebe1960005106918170014551231950b75fc4402da1732fc9bebe19910970014551231950b75fc4402da1732fc9bebe19039170014551231950b75fc4402da1732fc9bebe1983116107b8578170014551231950b75fc4402da1732fc9bebe19910970014551231950b75fc4402da1732fc9bebe19039070014551231950b75fc4402da1732fc9bebe1982116107b857602092600092608092604051928352601b868401526040830152606082015282805260015afa156107ac576000516001600160a01b0390811691161490565b6040513d6000823e3d90fd5b634e487b7160e01b600052601160045260246000fd5b5050505050600090565b50505050600090565b5070014551231950b75fc4402da1732fc9bebe19831015610621565b506401000003d01984101561061a565b909160015490811561083c57806020116100d1578335906040116100d157602061083994013591610607565b90565b631ecd631160e21b60005260046000fd5b60208151910120906108896040805180936020820195865261087881518092602086860191016105e4565b81010301601f1981018352826104ec565b51902090565b6000546001600160a01b031633036108a357565b63118cdaa760e01b6000523360045260246000fd5b6108c181610905565b919091156108fb5760405191602083019182526040830152604082526108e86060836104ec565b905190206001600160a01b031690600190565b5050600090600090565b6401000003d0198110156109785761093d906401000003d01990816007816000840908906401000003d0199081818009900908610981565b6001811661094c575b90600190565b6401000003d019036401000003d01981111561094657634e487b7160e01b600052601160045260246000fd5b50600090600090565b8015610a3057600190600160ff1b90815b61099b57505090565b90916401000003d0199063400000f4600160fe1b0384161515830a9082908009096401000003d0199063400000f4600160fe1b03600185901c161515830a9082908009096401000003d0199063400000f4600160fe1b03600285901c161515830a9082908009096401000003d0199063400000f4600160fe1b03600385901c161515830a9082908009099160041c9081610992565b5060009056fea2646970667358221220268cba503ad1e66e5164b0d5000926b785a794d8ad80f180b09c5bcd84a550c864736f6c634300081b0033",
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

// UsedNonces is a free data retrieval call binding the contract method 0xfeb61724.
//
// Solidity: function usedNonces(bytes32 ) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCaller) UsedNonces(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ValidatorRegistry.contract.Call(opts, &out, "usedNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedNonces is a free data retrieval call binding the contract method 0xfeb61724.
//
// Solidity: function usedNonces(bytes32 ) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistrySession) UsedNonces(arg0 [32]byte) (bool, error) {
	return _ValidatorRegistry.Contract.UsedNonces(&_ValidatorRegistry.CallOpts, arg0)
}

// UsedNonces is a free data retrieval call binding the contract method 0xfeb61724.
//
// Solidity: function usedNonces(bytes32 ) view returns(bool)
func (_ValidatorRegistry *ValidatorRegistryCallerSession) UsedNonces(arg0 [32]byte) (bool, error) {
	return _ValidatorRegistry.Contract.UsedNonces(&_ValidatorRegistry.CallOpts, arg0)
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

// SetJointPublicKeySigned is a paid mutator transaction binding the contract method 0x996d1b91.
//
// Solidity: function setJointPublicKeySigned(bytes32 _jointPublicKey, bytes32 _nonce, bytes signature) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactor) SetJointPublicKeySigned(opts *bind.TransactOpts, _jointPublicKey [32]byte, _nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _ValidatorRegistry.contract.Transact(opts, "setJointPublicKeySigned", _jointPublicKey, _nonce, signature)
}

// SetJointPublicKeySigned is a paid mutator transaction binding the contract method 0x996d1b91.
//
// Solidity: function setJointPublicKeySigned(bytes32 _jointPublicKey, bytes32 _nonce, bytes signature) returns()
func (_ValidatorRegistry *ValidatorRegistrySession) SetJointPublicKeySigned(_jointPublicKey [32]byte, _nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetJointPublicKeySigned(&_ValidatorRegistry.TransactOpts, _jointPublicKey, _nonce, signature)
}

// SetJointPublicKeySigned is a paid mutator transaction binding the contract method 0x996d1b91.
//
// Solidity: function setJointPublicKeySigned(bytes32 _jointPublicKey, bytes32 _nonce, bytes signature) returns()
func (_ValidatorRegistry *ValidatorRegistryTransactorSession) SetJointPublicKeySigned(_jointPublicKey [32]byte, _nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _ValidatorRegistry.Contract.SetJointPublicKeySigned(&_ValidatorRegistry.TransactOpts, _jointPublicKey, _nonce, signature)
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
