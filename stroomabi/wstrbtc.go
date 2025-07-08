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

// WStrBtcMetaData contains all meta data concerning the WStrBtc contract.
var WStrBtcMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strBTC\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertToAssets\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertToShares\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxDeposit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxMint\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewMint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewWithdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2612ExpiredSignature\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2612InvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxDeposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxMint\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x6101a0604052348015610010575f5ffd5b50604051611e9a380380611e9a83398101604081905261002f91610335565b6040518060400160405280601681526020017f57726170706564205374726f6f6d20426974636f696e0000000000000000000081525080604051806040016040528060018152602001603160f81b815250836040518060400160405280601681526020017f57726170706564205374726f6f6d20426974636f696e00000000000000000000815250604051806040016040528060078152602001667773747242544360c81b81525081600390816100e691906103fa565b5060046100f382826103fa565b5050505f5f610107836101e760201b60201c565b9150915081610117576012610119565b805b60ff1660a05250506001600160a01b03166080526101388260056102bd565b610160526101478160066102bd565b61018052815160208084019190912061012052815190820120610140524660e0526101d56101205161014051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60c05250503061010052506105399050565b60408051600481526024810182526020810180516001600160e01b031663313ce56760e01b17905290515f918291829182916001600160a01b0387169161022d916104b4565b5f60405180830381855afa9150503d805f8114610265576040519150601f19603f3d011682016040523d82523d5f602084013e61026a565b606091505b509150915081801561027e57506020815110155b156102b1575f8180602001905181019061029891906104ca565b905060ff81116102af576001969095509350505050565b505b505f9485945092505050565b5f6020835110156102d8576102d1836102ef565b90506102e9565b816102e384826103fa565b5060ff90505b92915050565b5f5f829050601f81511115610322578260405163305a27a960e01b815260040161031991906104e1565b60405180910390fd5b805161032d82610516565b179392505050565b5f60208284031215610345575f5ffd5b81516001600160a01b038116811461035b575f5ffd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061038a57607f821691505b6020821081036103a857634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156103f557805f5260205f20601f840160051c810160208510156103d35750805b601f840160051c820191505b818110156103f2575f81556001016103df565b50505b505050565b81516001600160401b0381111561041357610413610362565b610427816104218454610376565b846103ae565b6020601f821160018114610459575f83156104425750848201515b5f19600385901b1c1916600184901b1784556103f2565b5f84815260208120601f198516915b828110156104885787850151825560209485019460019092019101610468565b50848210156104a557868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f82518060208501845e5f920191825250919050565b5f602082840312156104da575f5ffd5b5051919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156103a8575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051610180516118e76105b35f395f610bf901525f610bcc01525f610ae301525f610abb01525f610a1601525f610a4001525f610a6a01525f6109e601525f818161028d0152818161042401528181610b380152610c5201526118e75ff3fe608060405234801561000f575f5ffd5b50600436106101c6575f3560e01c80637ecebe00116100fe578063ba0876521161009e578063d505accf1161006e578063d505accf146103c1578063d905777e146103d6578063dd62ed3e146103e9578063ef8b30f71461039b575f5ffd5b8063ba08765214610388578063c63d75b6146102b7578063c6e6f5921461039b578063ce96cb77146103ae575f5ffd5b806395d89b41116100d957806395d89b4114610347578063a9059cbb1461034f578063b3d7f6b914610362578063b460af9414610375575f5ffd5b80637ecebe001461030657806384b0196e1461031957806394bf804d14610334575f5ffd5b8063313ce56711610169578063402d267d11610144578063402d267d146102b75780634cdad506146101fa5780636e553f65146102cb57806370a08231146102de575f5ffd5b8063313ce5671461025e5780633644e5151461027857806338d52e0f14610280575f5ffd5b8063095ea7b3116101a4578063095ea7b31461020d5780630a28a4771461023057806318160ddd1461024357806323b872dd1461024b575f5ffd5b806301e1d114146101ca57806306fdde03146101e557806307a2d13a146101fa575b5f5ffd5b6101d2610421565b6040519081526020015b60405180910390f35b6101ed6104b0565b6040516101dc919061146e565b6101d2610208366004611480565b610540565b61022061021b3660046114b2565b610551565b60405190151581526020016101dc565b6101d261023e366004611480565b610568565b6002546101d2565b6102206102593660046114da565b610574565b610266610599565b60405160ff90911681526020016101dc565b6101d26105a2565b6040516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001681526020016101dc565b6101d26102c5366004611514565b505f1990565b6101d26102d936600461152d565b6105ab565b6101d26102ec366004611514565b6001600160a01b03165f9081526020819052604090205490565b6101d2610314366004611514565b6105dc565b6103216105f9565b6040516101dc9796959493929190611557565b6101d261034236600461152d565b61063b565b6101ed610656565b61022061035d3660046114b2565b610665565b6101d2610370366004611480565b610672565b6101d26103833660046115ed565b61067e565b6101d26103963660046115ed565b6106d4565b6101d26103a9366004611480565b610721565b6101d26103bc366004611514565b61072c565b6103d46103cf366004611626565b61074e565b005b6101d26103e4366004611514565b610884565b6101d26103f7366004611693565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b5f7f00000000000000000000000000000000000000000000000000000000000000006040516370a0823160e01b81523060048201526001600160a01b0391909116906370a0823190602401602060405180830381865afa158015610487573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104ab91906116bb565b905090565b6060600380546104bf906116d2565b80601f01602080910402602001604051908101604052809291908181526020018280546104eb906116d2565b80156105365780601f1061050d57610100808354040283529160200191610536565b820191905f5260205f20905b81548152906001019060200180831161051957829003601f168201915b5050505050905090565b5f61054b825f6108a1565b92915050565b5f3361055e8185856108d9565b5060019392505050565b5f61054b8260016108eb565b5f3361058185828561091a565b61058c858585610982565b60019150505b9392505050565b5f6104ab6109df565b5f6104ab610a0a565b5f5f196105bc565b60405180910390fd5b5f6105c685610721565b90506105d433858784610b33565b949350505050565b6001600160a01b0381165f9081526007602052604081205461054b565b5f6060805f5f5f606061060a610bc5565b610612610bf2565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f5f195f61064885610672565b90506105d433858388610b33565b6060600480546104bf906116d2565b5f3361055e818585610982565b5f61054b8260016108a1565b5f5f6106898361072c565b9050808511156106b257828582604051633fa733bb60e21b81526004016105b39392919061170a565b5f6106bc86610568565b90506106cb3386868985610c1f565b95945050505050565b5f5f6106df83610884565b90508085111561070857828582604051632e52afbb60e21b81526004016105b39392919061170a565b5f61071286610540565b90506106cb338686848a610c1f565b5f61054b825f6108eb565b6001600160a01b0381165f9081526020819052604081205461054b905f6108a1565b834211156107725760405163313c898160e11b8152600481018590526024016105b3565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886107bd8c6001600160a01b03165f90815260076020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61081782610cdf565b90505f61082682878787610d0b565b9050896001600160a01b0316816001600160a01b03161461086d576040516325c0072360e11b81526001600160a01b0380831660048301528b1660248201526044016105b3565b6108788a8a8a6108d9565b50505050505050505050565b6001600160a01b0381165f9081526020819052604081205461054b565b5f6105926108ad610421565b6108b890600161173f565b6108c35f600a611835565b6002546108d0919061173f565b85919085610d37565b6108e68383836001610d79565b505050565b5f6105926108fa82600a611835565b600254610907919061173f565b61090f610421565b6108d090600161173f565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811461097c578181101561096e57828183604051637dc7a0d960e11b81526004016105b39392919061170a565b61097c84848484035f610d79565b50505050565b6001600160a01b0383166109ab57604051634b637e8f60e11b81525f60048201526024016105b3565b6001600160a01b0382166109d45760405163ec442f0560e01b81525f60048201526024016105b3565b6108e6838383610e3d565b5f6104ab817f0000000000000000000000000000000000000000000000000000000000000000611843565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015610a6257507f000000000000000000000000000000000000000000000000000000000000000046145b15610a8c57507f000000000000000000000000000000000000000000000000000000000000000090565b6104ab604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b610b5f7f0000000000000000000000000000000000000000000000000000000000000000853085610f50565b610b698382610fb7565b826001600160a01b0316846001600160a01b03167fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d78484604051610bb7929190918252602082015260400190565b60405180910390a350505050565b60606104ab7f00000000000000000000000000000000000000000000000000000000000000006005610fef565b60606104ab7f00000000000000000000000000000000000000000000000000000000000000006006610fef565b826001600160a01b0316856001600160a01b031614610c4357610c4383868361091a565b610c4d8382611098565b610c787f000000000000000000000000000000000000000000000000000000000000000085846110cc565b826001600160a01b0316846001600160a01b0316866001600160a01b03167ffbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db8585604051610cd0929190918252602082015260400190565b60405180910390a45050505050565b5f61054b610ceb610a0a565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610d1b888888886110fd565b925092509250610d2b82826111c5565b50909695505050505050565b5f610d64610d448361127d565b8015610d5f57505f8480610d5a57610d5a61185c565b868809115b151590565b610d6f8686866112a9565b6106cb919061173f565b6001600160a01b038416610da25760405163e602df0560e01b81525f60048201526024016105b3565b6001600160a01b038316610dcb57604051634a1406b160e11b81525f60048201526024016105b3565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561097c57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bb791815260200190565b6001600160a01b038316610e67578060025f828254610e5c919061173f565b90915550610ec49050565b6001600160a01b0383165f9081526020819052604090205481811015610ea65783818360405163391434e360e21b81526004016105b39392919061170a565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216610ee057600280548290039055610efe565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f4391815260200190565b60405180910390a3505050565b6040516001600160a01b03848116602483015283811660448301526064820183905261097c9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061135f565b6001600160a01b038216610fe05760405163ec442f0560e01b81525f60048201526024016105b3565b610feb5f8383610e3d565b5050565b606060ff831461100957611002836113cb565b905061054b565b818054611015906116d2565b80601f0160208091040260200160405190810160405280929190818152602001828054611041906116d2565b801561108c5780601f106110635761010080835404028352916020019161108c565b820191905f5260205f20905b81548152906001019060200180831161106f57829003601f168201915b5050505050905061054b565b6001600160a01b0382166110c157604051634b637e8f60e11b81525f60048201526024016105b3565b610feb825f83610e3d565b6040516001600160a01b038381166024830152604482018390526108e691859182169063a9059cbb90606401610f85565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561113657505f915060039050826111bb565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611187573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166111b257505f9250600191508290506111bb565b92505f91508190505b9450945094915050565b5f8260038111156111d8576111d8611870565b036111e1575050565b60018260038111156111f5576111f5611870565b036112135760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561122757611227611870565b036112485760405163fce698f760e01b8152600481018290526024016105b3565b600382600381111561125c5761125c611870565b03610feb576040516335e2f38360e21b8152600481018290526024016105b3565b5f600282600381111561129257611292611870565b61129c9190611884565b60ff166001149050919050565b5f838302815f1985870982811083820303915050805f036112dd578382816112d3576112d361185c565b0492505050610592565b8084116112f4576112f46003851502601118611408565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b5f5f60205f8451602086015f885af18061137e576040513d5f823e3d81fd5b50505f513d915081156113955780600114156113a2565b6001600160a01b0384163b155b1561097c57604051635274afe760e01b81526001600160a01b03851660048201526024016105b3565b60605f6113d783611419565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b634e487b715f52806020526024601cfd5b5f60ff8216601f81111561054b57604051632cd44ac360e21b815260040160405180910390fd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6105926020830184611440565b5f60208284031215611490575f5ffd5b5035919050565b80356001600160a01b03811681146114ad575f5ffd5b919050565b5f5f604083850312156114c3575f5ffd5b6114cc83611497565b946020939093013593505050565b5f5f5f606084860312156114ec575f5ffd5b6114f584611497565b925061150360208501611497565b929592945050506040919091013590565b5f60208284031215611524575f5ffd5b61059282611497565b5f5f6040838503121561153e575f5ffd5b8235915061154e60208401611497565b90509250929050565b60ff60f81b8816815260e060208201525f61157560e0830189611440565b82810360408401526115878189611440565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156115dc5783518352602093840193909201916001016115be565b50909b9a5050505050505050505050565b5f5f5f606084860312156115ff575f5ffd5b8335925061160f60208501611497565b915061161d60408501611497565b90509250925092565b5f5f5f5f5f5f5f60e0888a03121561163c575f5ffd5b61164588611497565b965061165360208901611497565b95506040880135945060608801359350608088013560ff81168114611676575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156116a4575f5ffd5b6116ad83611497565b915061154e60208401611497565b5f602082840312156116cb575f5ffd5b5051919050565b600181811c908216806116e657607f821691505b60208210810361170457634e487b7160e01b5f52602260045260245ffd5b50919050565b6001600160a01b039390931683526020830191909152604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561054b5761054b61172b565b6001815b600184111561178d578085048111156117715761177161172b565b600184161561177f57908102905b60019390931c928002611756565b935093915050565b5f826117a35750600161054b565b816117af57505f61054b565b81600181146117c557600281146117cf576117eb565b600191505061054b565b60ff8411156117e0576117e061172b565b50506001821b61054b565b5060208310610133831016604e8410600b841016171561180e575081810a61054b565b61181a5f198484611752565b805f190482111561182d5761182d61172b565b029392505050565b5f61059260ff841683611795565b60ff818116838216019081111561054b5761054b61172b565b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f60ff8316806118a257634e487b7160e01b5f52601260045260245ffd5b8060ff8416069150509291505056fea2646970667358221220c10ef49dc692c351d40e24d7697a257c7c92954eea538dd7e409543c7b25d7b864736f6c634300081b0033",
}

// WStrBtcABI is the input ABI used to generate the binding from.
// Deprecated: Use WStrBtcMetaData.ABI instead.
var WStrBtcABI = WStrBtcMetaData.ABI

// WStrBtcBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WStrBtcMetaData.Bin instead.
var WStrBtcBin = WStrBtcMetaData.Bin

// DeployWStrBtc deploys a new Ethereum contract, binding an instance of WStrBtc to it.
func DeployWStrBtc(auth *bind.TransactOpts, backend bind.ContractBackend, _strBTC common.Address) (common.Address, *types.Transaction, *WStrBtc, error) {
	parsed, err := WStrBtcMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WStrBtcBin), backend, _strBTC)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WStrBtc{WStrBtcCaller: WStrBtcCaller{contract: contract}, WStrBtcTransactor: WStrBtcTransactor{contract: contract}, WStrBtcFilterer: WStrBtcFilterer{contract: contract}}, nil
}

// WStrBtc is an auto generated Go binding around an Ethereum contract.
type WStrBtc struct {
	WStrBtcCaller     // Read-only binding to the contract
	WStrBtcTransactor // Write-only binding to the contract
	WStrBtcFilterer   // Log filterer for contract events
}

// WStrBtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type WStrBtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WStrBtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WStrBtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WStrBtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WStrBtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WStrBtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WStrBtcSession struct {
	Contract     *WStrBtc          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WStrBtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WStrBtcCallerSession struct {
	Contract *WStrBtcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WStrBtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WStrBtcTransactorSession struct {
	Contract     *WStrBtcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WStrBtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type WStrBtcRaw struct {
	Contract *WStrBtc // Generic contract binding to access the raw methods on
}

// WStrBtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WStrBtcCallerRaw struct {
	Contract *WStrBtcCaller // Generic read-only contract binding to access the raw methods on
}

// WStrBtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WStrBtcTransactorRaw struct {
	Contract *WStrBtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWStrBtc creates a new instance of WStrBtc, bound to a specific deployed contract.
func NewWStrBtc(address common.Address, backend bind.ContractBackend) (*WStrBtc, error) {
	contract, err := bindWStrBtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WStrBtc{WStrBtcCaller: WStrBtcCaller{contract: contract}, WStrBtcTransactor: WStrBtcTransactor{contract: contract}, WStrBtcFilterer: WStrBtcFilterer{contract: contract}}, nil
}

// NewWStrBtcCaller creates a new read-only instance of WStrBtc, bound to a specific deployed contract.
func NewWStrBtcCaller(address common.Address, caller bind.ContractCaller) (*WStrBtcCaller, error) {
	contract, err := bindWStrBtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WStrBtcCaller{contract: contract}, nil
}

// NewWStrBtcTransactor creates a new write-only instance of WStrBtc, bound to a specific deployed contract.
func NewWStrBtcTransactor(address common.Address, transactor bind.ContractTransactor) (*WStrBtcTransactor, error) {
	contract, err := bindWStrBtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WStrBtcTransactor{contract: contract}, nil
}

// NewWStrBtcFilterer creates a new log filterer instance of WStrBtc, bound to a specific deployed contract.
func NewWStrBtcFilterer(address common.Address, filterer bind.ContractFilterer) (*WStrBtcFilterer, error) {
	contract, err := bindWStrBtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WStrBtcFilterer{contract: contract}, nil
}

// bindWStrBtc binds a generic wrapper to an already deployed contract.
func bindWStrBtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WStrBtcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WStrBtc *WStrBtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WStrBtc.Contract.WStrBtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WStrBtc *WStrBtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WStrBtc.Contract.WStrBtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WStrBtc *WStrBtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WStrBtc.Contract.WStrBtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WStrBtc *WStrBtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WStrBtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WStrBtc *WStrBtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WStrBtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WStrBtc *WStrBtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WStrBtc.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WStrBtc *WStrBtcCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WStrBtc *WStrBtcSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WStrBtc.Contract.DOMAINSEPARATOR(&_WStrBtc.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WStrBtc *WStrBtcCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WStrBtc.Contract.DOMAINSEPARATOR(&_WStrBtc.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WStrBtc *WStrBtcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.Allowance(&_WStrBtc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.Allowance(&_WStrBtc.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_WStrBtc *WStrBtcCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_WStrBtc *WStrBtcSession) Asset() (common.Address, error) {
	return _WStrBtc.Contract.Asset(&_WStrBtc.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_WStrBtc *WStrBtcCallerSession) Asset() (common.Address, error) {
	return _WStrBtc.Contract.Asset(&_WStrBtc.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WStrBtc *WStrBtcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.BalanceOf(&_WStrBtc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.BalanceOf(&_WStrBtc.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_WStrBtc *WStrBtcSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.ConvertToAssets(&_WStrBtc.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.ConvertToAssets(&_WStrBtc.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_WStrBtc *WStrBtcSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.ConvertToShares(&_WStrBtc.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.ConvertToShares(&_WStrBtc.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WStrBtc *WStrBtcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WStrBtc *WStrBtcSession) Decimals() (uint8, error) {
	return _WStrBtc.Contract.Decimals(&_WStrBtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WStrBtc *WStrBtcCallerSession) Decimals() (uint8, error) {
	return _WStrBtc.Contract.Decimals(&_WStrBtc.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WStrBtc *WStrBtcCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WStrBtc *WStrBtcSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WStrBtc.Contract.Eip712Domain(&_WStrBtc.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WStrBtc *WStrBtcCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WStrBtc.Contract.Eip712Domain(&_WStrBtc.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_WStrBtc *WStrBtcSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.MaxDeposit(&_WStrBtc.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.MaxDeposit(&_WStrBtc.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_WStrBtc *WStrBtcSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.MaxMint(&_WStrBtc.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.MaxMint(&_WStrBtc.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_WStrBtc *WStrBtcSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.MaxRedeem(&_WStrBtc.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.MaxRedeem(&_WStrBtc.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_WStrBtc *WStrBtcSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.MaxWithdraw(&_WStrBtc.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.MaxWithdraw(&_WStrBtc.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WStrBtc *WStrBtcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WStrBtc *WStrBtcSession) Name() (string, error) {
	return _WStrBtc.Contract.Name(&_WStrBtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WStrBtc *WStrBtcCallerSession) Name() (string, error) {
	return _WStrBtc.Contract.Name(&_WStrBtc.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WStrBtc *WStrBtcSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.Nonces(&_WStrBtc.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WStrBtc.Contract.Nonces(&_WStrBtc.CallOpts, owner)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_WStrBtc *WStrBtcSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.PreviewDeposit(&_WStrBtc.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.PreviewDeposit(&_WStrBtc.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_WStrBtc *WStrBtcSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.PreviewMint(&_WStrBtc.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.PreviewMint(&_WStrBtc.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_WStrBtc *WStrBtcSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.PreviewRedeem(&_WStrBtc.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.PreviewRedeem(&_WStrBtc.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_WStrBtc *WStrBtcSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.PreviewWithdraw(&_WStrBtc.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.PreviewWithdraw(&_WStrBtc.CallOpts, assets)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WStrBtc *WStrBtcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WStrBtc *WStrBtcSession) Symbol() (string, error) {
	return _WStrBtc.Contract.Symbol(&_WStrBtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WStrBtc *WStrBtcCallerSession) Symbol() (string, error) {
	return _WStrBtc.Contract.Symbol(&_WStrBtc.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_WStrBtc *WStrBtcCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_WStrBtc *WStrBtcSession) TotalAssets() (*big.Int, error) {
	return _WStrBtc.Contract.TotalAssets(&_WStrBtc.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) TotalAssets() (*big.Int, error) {
	return _WStrBtc.Contract.TotalAssets(&_WStrBtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WStrBtc *WStrBtcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WStrBtc *WStrBtcSession) TotalSupply() (*big.Int, error) {
	return _WStrBtc.Contract.TotalSupply(&_WStrBtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) TotalSupply() (*big.Int, error) {
	return _WStrBtc.Contract.TotalSupply(&_WStrBtc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WStrBtc *WStrBtcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WStrBtc *WStrBtcSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.Approve(&_WStrBtc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WStrBtc *WStrBtcTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.Approve(&_WStrBtc.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_WStrBtc *WStrBtcTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_WStrBtc *WStrBtcSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WStrBtc.Contract.Deposit(&_WStrBtc.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_WStrBtc *WStrBtcTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WStrBtc.Contract.Deposit(&_WStrBtc.TransactOpts, assets, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_WStrBtc *WStrBtcTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_WStrBtc *WStrBtcSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WStrBtc.Contract.Mint(&_WStrBtc.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_WStrBtc *WStrBtcTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WStrBtc.Contract.Mint(&_WStrBtc.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WStrBtc *WStrBtcTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WStrBtc *WStrBtcSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WStrBtc.Contract.Permit(&_WStrBtc.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WStrBtc *WStrBtcTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WStrBtc.Contract.Permit(&_WStrBtc.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_WStrBtc *WStrBtcTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_WStrBtc *WStrBtcSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WStrBtc.Contract.Redeem(&_WStrBtc.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_WStrBtc *WStrBtcTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WStrBtc.Contract.Redeem(&_WStrBtc.TransactOpts, shares, receiver, owner)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WStrBtc *WStrBtcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WStrBtc *WStrBtcSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.Transfer(&_WStrBtc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WStrBtc *WStrBtcTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.Transfer(&_WStrBtc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WStrBtc *WStrBtcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WStrBtc *WStrBtcSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.TransferFrom(&_WStrBtc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WStrBtc *WStrBtcTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.TransferFrom(&_WStrBtc.TransactOpts, from, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_WStrBtc *WStrBtcTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_WStrBtc *WStrBtcSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WStrBtc.Contract.Withdraw(&_WStrBtc.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_WStrBtc *WStrBtcTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WStrBtc.Contract.Withdraw(&_WStrBtc.TransactOpts, assets, receiver, owner)
}

// WStrBtcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WStrBtc contract.
type WStrBtcApprovalIterator struct {
	Event *WStrBtcApproval // Event containing the contract specifics and raw log

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
func (it *WStrBtcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WStrBtcApproval)
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
		it.Event = new(WStrBtcApproval)
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
func (it *WStrBtcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WStrBtcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WStrBtcApproval represents a Approval event raised by the WStrBtc contract.
type WStrBtcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WStrBtc *WStrBtcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WStrBtcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WStrBtc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WStrBtcApprovalIterator{contract: _WStrBtc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WStrBtc *WStrBtcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WStrBtcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WStrBtc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WStrBtcApproval)
				if err := _WStrBtc.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WStrBtc *WStrBtcFilterer) ParseApproval(log types.Log) (*WStrBtcApproval, error) {
	event := new(WStrBtcApproval)
	if err := _WStrBtc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WStrBtcDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WStrBtc contract.
type WStrBtcDepositIterator struct {
	Event *WStrBtcDeposit // Event containing the contract specifics and raw log

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
func (it *WStrBtcDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WStrBtcDeposit)
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
		it.Event = new(WStrBtcDeposit)
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
func (it *WStrBtcDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WStrBtcDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WStrBtcDeposit represents a Deposit event raised by the WStrBtc contract.
type WStrBtcDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_WStrBtc *WStrBtcFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*WStrBtcDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WStrBtc.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WStrBtcDepositIterator{contract: _WStrBtc.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_WStrBtc *WStrBtcFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WStrBtcDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WStrBtc.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WStrBtcDeposit)
				if err := _WStrBtc.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_WStrBtc *WStrBtcFilterer) ParseDeposit(log types.Log) (*WStrBtcDeposit, error) {
	event := new(WStrBtcDeposit)
	if err := _WStrBtc.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WStrBtcEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the WStrBtc contract.
type WStrBtcEIP712DomainChangedIterator struct {
	Event *WStrBtcEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *WStrBtcEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WStrBtcEIP712DomainChanged)
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
		it.Event = new(WStrBtcEIP712DomainChanged)
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
func (it *WStrBtcEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WStrBtcEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WStrBtcEIP712DomainChanged represents a EIP712DomainChanged event raised by the WStrBtc contract.
type WStrBtcEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WStrBtc *WStrBtcFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*WStrBtcEIP712DomainChangedIterator, error) {

	logs, sub, err := _WStrBtc.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &WStrBtcEIP712DomainChangedIterator{contract: _WStrBtc.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WStrBtc *WStrBtcFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *WStrBtcEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _WStrBtc.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WStrBtcEIP712DomainChanged)
				if err := _WStrBtc.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WStrBtc *WStrBtcFilterer) ParseEIP712DomainChanged(log types.Log) (*WStrBtcEIP712DomainChanged, error) {
	event := new(WStrBtcEIP712DomainChanged)
	if err := _WStrBtc.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WStrBtcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WStrBtc contract.
type WStrBtcTransferIterator struct {
	Event *WStrBtcTransfer // Event containing the contract specifics and raw log

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
func (it *WStrBtcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WStrBtcTransfer)
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
		it.Event = new(WStrBtcTransfer)
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
func (it *WStrBtcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WStrBtcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WStrBtcTransfer represents a Transfer event raised by the WStrBtc contract.
type WStrBtcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WStrBtc *WStrBtcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WStrBtcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WStrBtc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WStrBtcTransferIterator{contract: _WStrBtc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WStrBtc *WStrBtcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WStrBtcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WStrBtc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WStrBtcTransfer)
				if err := _WStrBtc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WStrBtc *WStrBtcFilterer) ParseTransfer(log types.Log) (*WStrBtcTransfer, error) {
	event := new(WStrBtcTransfer)
	if err := _WStrBtc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WStrBtcWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the WStrBtc contract.
type WStrBtcWithdrawIterator struct {
	Event *WStrBtcWithdraw // Event containing the contract specifics and raw log

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
func (it *WStrBtcWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WStrBtcWithdraw)
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
		it.Event = new(WStrBtcWithdraw)
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
func (it *WStrBtcWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WStrBtcWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WStrBtcWithdraw represents a Withdraw event raised by the WStrBtc contract.
type WStrBtcWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_WStrBtc *WStrBtcFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*WStrBtcWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WStrBtc.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WStrBtcWithdrawIterator{contract: _WStrBtc.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_WStrBtc *WStrBtcFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *WStrBtcWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WStrBtc.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WStrBtcWithdraw)
				if err := _WStrBtc.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_WStrBtc *WStrBtcFilterer) ParseWithdraw(log types.Log) (*WStrBtcWithdraw, error) {
	event := new(WStrBtcWithdraw)
	if err := _WStrBtc.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
