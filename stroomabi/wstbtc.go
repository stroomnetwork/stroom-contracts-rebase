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

// WStBtcMetaData contains all meta data concerning the WStBtc contract.
var WStBtcMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stBTC\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stBTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStBTC\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stBTCPerToken\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokensPerStBTC\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unwrap\",\"inputs\":[{\"name\":\"wstBTCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wrap\",\"inputs\":[{\"name\":\"stBTCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2612ExpiredSignature\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2612InvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x61016060405234801561001157600080fd5b50604051611a7e380380611a7e83398101604081905261003091610247565b6040518060400160405280601681526020017f57726170706564205374726f6f6d20426974636f696e0000000000000000000081525080604051806040016040528060018152602001603160f81b8152506040518060400160405280601681526020017f57726170706564205374726f6f6d20426974636f696e000000000000000000008152506040518060400160405280600681526020016577737442544360d01b81525081600390816100e59190610316565b5060046100f28282610316565b50610102915083905060056101cd565b610120526101118160066101cd565b61014052815160208084019190912060e052815190820120610100524660a05261019e60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60805250503060c05250600880546001600160a01b0319166001600160a01b0392909216919091179055610446565b60006020835110156101e9576101e283610200565b90506101fa565b816101f48482610316565b5060ff90505b92915050565b600080829050601f81511115610234578260405163305a27a960e01b815260040161022b91906103d4565b60405180910390fd5b805161023f82610422565b179392505050565b60006020828403121561025957600080fd5b81516001600160a01b038116811461027057600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806102a157607f821691505b6020821081036102c157634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561031157806000526020600020601f840160051c810160208510156102ee5750805b601f840160051c820191505b8181101561030e57600081556001016102fa565b50505b505050565b81516001600160401b0381111561032f5761032f610277565b6103438161033d845461028d565b846102c7565b6020601f821160018114610377576000831561035f5750848201515b600019600385901b1c1916600184901b17845561030e565b600084815260208120601f198516915b828110156103a75787850151825560209485019460019092019101610387565b50848210156103c55786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b602081526000825180602084015260005b8181101561040257602081860181015160408684010152016103e5565b506000604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102c15760001960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516115de6104a06000396000610cdc01526000610caf01526000610c5701526000610c2f01526000610b8a01526000610bb401526000610bde01526115de6000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806370a08231116100ad578063a9059cbb11610071578063a9059cbb1461025e578063d505accf14610271578063dd62ed3e14610286578063de0e9a3e146102bf578063ea598cb0146102d257600080fd5b806370a08231146101ec5780637ecebe001461021557806384b0196e1461022857806392ebbb2f1461024357806395d89b411461025657600080fd5b80632792949d116100f45780632792949d1461018c578063313ce567146101975780633644e515146101a65780633714b060146101ae5780635ae7daa5146101c157600080fd5b806306fdde0314610126578063095ea7b31461014457806318160ddd1461016757806323b872dd14610179575b600080fd5b61012e6102e5565b60405161013b91906112ac565b60405180910390f35b6101576101523660046112e2565b610377565b604051901515815260200161013b565b6002545b60405190815260200161013b565b61015761018736600461130c565b610391565b61016b6305f5e10081565b6040516012815260200161013b565b61016b6103b5565b61016b6101bc366004611349565b6103c4565b6008546101d4906001600160a01b031681565b6040516001600160a01b03909116815260200161013b565b61016b6101fa366004611362565b6001600160a01b031660009081526020819052604090205490565b61016b610223366004611362565b6104b6565b6102306104d4565b60405161013b979695949392919061137d565b61016b610251366004611349565b61051a565b61012e6105f0565b61015761026c3660046112e2565b6105ff565b61028461027f366004611415565b61060d565b005b61016b610294366004611488565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b61016b6102cd366004611349565b610747565b61016b6102e0366004611349565b61090a565b6060600380546102f4906114bb565b80601f0160208091040260200160405190810160405280929190818152602001828054610320906114bb565b801561036d5780601f106103425761010080835404028352916020019161036d565b820191906000526020600020905b81548152906001019060200180831161035057829003601f168201915b5050505050905090565b600033610385818585610a8e565b60019150505b92915050565b60003361039f858285610aa0565b6103aa858585610b1e565b506001949350505050565b60006103bf610b7d565b905090565b6008546040516370a0823160e01b815230600482015260009182916001600160a01b03909116906370a0823190602401602060405180830381865afa158015610411573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043591906114f5565b9050600061044260025490565b9050600082116104995760405162461bcd60e51b815260206004820152601860248201527f7773744254433a204e6f2073744254432062616c616e6365000000000000000060448201526064015b60405180910390fd5b816104a48583611524565b6104ae919061153b565b949350505050565b6001600160a01b03811660009081526007602052604081205461038b565b6000606080600080600060606104e8610ca8565b6104f0610cd5565b60408051600080825260208201909252600f60f81b9b939a50919850469750309650945092509050565b6008546040516370a0823160e01b815230600482015260009182916001600160a01b03909116906370a0823190602401602060405180830381865afa158015610567573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058b91906114f5565b9050600061059860025490565b9050600081116105e55760405162461bcd60e51b81526020600482015260186024820152777773744254433a204e6f2077737442544320737570706c7960401b6044820152606401610490565b806104a48584611524565b6060600480546102f4906114bb565b600033610385818585610b1e565b834211156106315760405163313c898160e11b815260048101859052602401610490565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861067e8c6001600160a01b0316600090815260076020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e00160405160208183030381529060405280519060200120905060006106d982610d02565b905060006106e982878787610d2f565b9050896001600160a01b0316816001600160a01b031614610730576040516325c0072360e11b81526001600160a01b0380831660048301528b166024820152604401610490565b61073b8a8a8a610a8e565b50505050505050505050565b60008082116107a25760405162461bcd60e51b815260206004820152602160248201527f7773744254433a2043616e6e6f7420756e77726170207a65726f2077737442546044820152604360f81b6064820152608401610490565b6008546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa1580156107eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080f91906114f5565b9050600061081c60025490565b9050600081116108695760405162461bcd60e51b81526020600482015260186024820152777773744254433a204e6f2077737442544320737570706c7960401b6044820152606401610490565b6000816108768487611524565b610880919061153b565b905061088c3386610d5d565b60085460405163a9059cbb60e01b8152336004820152602481018390526001600160a01b039091169063a9059cbb906044016020604051808303816000875af11580156108dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610901919061155d565b50949350505050565b600080821161095b5760405162461bcd60e51b815260206004820152601e60248201527f7773744254433a2043616e6e6f742077726170207a65726f20737442544300006044820152606401610490565b6008546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa1580156109a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c891906114f5565b905060006109d560025490565b905060008115806109e4575082155b156109f0575083610a08565b826109fb8387611524565b610a05919061153b565b90505b6008546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303816000875af1158015610a5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a83919061155d565b506104ae3382610d97565b610a9b8383836001610dcd565b505050565b6001600160a01b038381166000908152600160209081526040808320938616835292905220546000198114610b185781811015610b0957604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610490565b610b1884848484036000610dcd565b50505050565b6001600160a01b038316610b4857604051634b637e8f60e11b815260006004820152602401610490565b6001600160a01b038216610b725760405163ec442f0560e01b815260006004820152602401610490565b610a9b838383610ea2565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015610bd657507f000000000000000000000000000000000000000000000000000000000000000046145b15610c0057507f000000000000000000000000000000000000000000000000000000000000000090565b6103bf604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60606103bf7f00000000000000000000000000000000000000000000000000000000000000006005610fcc565b60606103bf7f00000000000000000000000000000000000000000000000000000000000000006006610fcc565b600061038b610d0f610b7d565b8360405161190160f01b8152600281019290925260228201526042902090565b600080600080610d4188888888611077565b925092509250610d518282611146565b50909695505050505050565b6001600160a01b038216610d8757604051634b637e8f60e11b815260006004820152602401610490565b610d9382600083610ea2565b5050565b6001600160a01b038216610dc15760405163ec442f0560e01b815260006004820152602401610490565b610d9360008383610ea2565b6001600160a01b038416610df75760405163e602df0560e01b815260006004820152602401610490565b6001600160a01b038316610e2157604051634a1406b160e11b815260006004820152602401610490565b6001600160a01b0380851660009081526001602090815260408083209387168352929052208290558015610b1857826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610e9491815260200190565b60405180910390a350505050565b6001600160a01b038316610ecd578060026000828254610ec2919061157f565b90915550610f3f9050565b6001600160a01b03831660009081526020819052604090205481811015610f205760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610490565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b038216610f5b57600280548290039055610f7a565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610fbf91815260200190565b60405180910390a3505050565b606060ff8314610fe657610fdf836111ff565b905061038b565b818054610ff2906114bb565b80601f016020809104026020016040519081016040528092919081815260200182805461101e906114bb565b801561106b5780601f106110405761010080835404028352916020019161106b565b820191906000526020600020905b81548152906001019060200180831161104e57829003601f168201915b5050505050905061038b565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156110b2575060009150600390508261113c565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611106573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166111325750600092506001915082905061113c565b9250600091508190505b9450945094915050565b600082600381111561115a5761115a611592565b03611163575050565b600182600381111561117757611177611592565b036111955760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156111a9576111a9611592565b036111ca5760405163fce698f760e01b815260048101829052602401610490565b60038260038111156111de576111de611592565b03610d93576040516335e2f38360e21b815260048101829052602401610490565b6060600061120c8361123e565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b600060ff8216601f81111561038b57604051632cd44ac360e21b815260040160405180910390fd5b6000815180845260005b8181101561128c57602081850181015186830182015201611270565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006112bf6020830184611266565b9392505050565b80356001600160a01b03811681146112dd57600080fd5b919050565b600080604083850312156112f557600080fd5b6112fe836112c6565b946020939093013593505050565b60008060006060848603121561132157600080fd5b61132a846112c6565b9250611338602085016112c6565b929592945050506040919091013590565b60006020828403121561135b57600080fd5b5035919050565b60006020828403121561137457600080fd5b6112bf826112c6565b60ff60f81b8816815260e06020820152600061139c60e0830189611266565b82810360408401526113ae8189611266565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501528451808252602080870193509091019060005b818110156114045783518352602093840193909201916001016113e6565b50909b9a5050505050505050505050565b600080600080600080600060e0888a03121561143057600080fd5b611439886112c6565b9650611447602089016112c6565b95506040880135945060608801359350608088013560ff8116811461146b57600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561149b57600080fd5b6114a4836112c6565b91506114b2602084016112c6565b90509250929050565b600181811c908216806114cf57607f821691505b6020821081036114ef57634e487b7160e01b600052602260045260246000fd5b50919050565b60006020828403121561150757600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761038b5761038b61150e565b60008261155857634e487b7160e01b600052601260045260246000fd5b500490565b60006020828403121561156f57600080fd5b815180151581146112bf57600080fd5b8082018082111561038b5761038b61150e565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220a450f50afe6cd7bf8de721df725e49098e70ff423f7a25eb4a2d9cfabe87daf564736f6c634300081b0033",
}

// WStBtcABI is the input ABI used to generate the binding from.
// Deprecated: Use WStBtcMetaData.ABI instead.
var WStBtcABI = WStBtcMetaData.ABI

// WStBtcBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WStBtcMetaData.Bin instead.
var WStBtcBin = WStBtcMetaData.Bin

// DeployWStBtc deploys a new Ethereum contract, binding an instance of WStBtc to it.
func DeployWStBtc(auth *bind.TransactOpts, backend bind.ContractBackend, _stBTC common.Address) (common.Address, *types.Transaction, *WStBtc, error) {
	parsed, err := WStBtcMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WStBtcBin), backend, _stBTC)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WStBtc{WStBtcCaller: WStBtcCaller{contract: contract}, WStBtcTransactor: WStBtcTransactor{contract: contract}, WStBtcFilterer: WStBtcFilterer{contract: contract}}, nil
}

// WStBtc is an auto generated Go binding around an Ethereum contract.
type WStBtc struct {
	WStBtcCaller     // Read-only binding to the contract
	WStBtcTransactor // Write-only binding to the contract
	WStBtcFilterer   // Log filterer for contract events
}

// WStBtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type WStBtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WStBtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WStBtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WStBtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WStBtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WStBtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WStBtcSession struct {
	Contract     *WStBtc           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WStBtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WStBtcCallerSession struct {
	Contract *WStBtcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WStBtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WStBtcTransactorSession struct {
	Contract     *WStBtcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WStBtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type WStBtcRaw struct {
	Contract *WStBtc // Generic contract binding to access the raw methods on
}

// WStBtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WStBtcCallerRaw struct {
	Contract *WStBtcCaller // Generic read-only contract binding to access the raw methods on
}

// WStBtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WStBtcTransactorRaw struct {
	Contract *WStBtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWStBtc creates a new instance of WStBtc, bound to a specific deployed contract.
func NewWStBtc(address common.Address, backend bind.ContractBackend) (*WStBtc, error) {
	contract, err := bindWStBtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WStBtc{WStBtcCaller: WStBtcCaller{contract: contract}, WStBtcTransactor: WStBtcTransactor{contract: contract}, WStBtcFilterer: WStBtcFilterer{contract: contract}}, nil
}

// NewWStBtcCaller creates a new read-only instance of WStBtc, bound to a specific deployed contract.
func NewWStBtcCaller(address common.Address, caller bind.ContractCaller) (*WStBtcCaller, error) {
	contract, err := bindWStBtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WStBtcCaller{contract: contract}, nil
}

// NewWStBtcTransactor creates a new write-only instance of WStBtc, bound to a specific deployed contract.
func NewWStBtcTransactor(address common.Address, transactor bind.ContractTransactor) (*WStBtcTransactor, error) {
	contract, err := bindWStBtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WStBtcTransactor{contract: contract}, nil
}

// NewWStBtcFilterer creates a new log filterer instance of WStBtc, bound to a specific deployed contract.
func NewWStBtcFilterer(address common.Address, filterer bind.ContractFilterer) (*WStBtcFilterer, error) {
	contract, err := bindWStBtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WStBtcFilterer{contract: contract}, nil
}

// bindWStBtc binds a generic wrapper to an already deployed contract.
func bindWStBtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WStBtcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WStBtc *WStBtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WStBtc.Contract.WStBtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WStBtc *WStBtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WStBtc.Contract.WStBtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WStBtc *WStBtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WStBtc.Contract.WStBtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WStBtc *WStBtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WStBtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WStBtc *WStBtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WStBtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WStBtc *WStBtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WStBtc.Contract.contract.Transact(opts, method, params...)
}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_WStBtc *WStBtcCaller) BTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "BTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_WStBtc *WStBtcSession) BTC() (*big.Int, error) {
	return _WStBtc.Contract.BTC(&_WStBtc.CallOpts)
}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_WStBtc *WStBtcCallerSession) BTC() (*big.Int, error) {
	return _WStBtc.Contract.BTC(&_WStBtc.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WStBtc *WStBtcCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WStBtc *WStBtcSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WStBtc.Contract.DOMAINSEPARATOR(&_WStBtc.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WStBtc *WStBtcCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WStBtc.Contract.DOMAINSEPARATOR(&_WStBtc.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WStBtc *WStBtcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WStBtc *WStBtcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WStBtc.Contract.Allowance(&_WStBtc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WStBtc *WStBtcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WStBtc.Contract.Allowance(&_WStBtc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WStBtc *WStBtcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WStBtc *WStBtcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WStBtc.Contract.BalanceOf(&_WStBtc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WStBtc *WStBtcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WStBtc.Contract.BalanceOf(&_WStBtc.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WStBtc *WStBtcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WStBtc *WStBtcSession) Decimals() (uint8, error) {
	return _WStBtc.Contract.Decimals(&_WStBtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WStBtc *WStBtcCallerSession) Decimals() (uint8, error) {
	return _WStBtc.Contract.Decimals(&_WStBtc.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WStBtc *WStBtcCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "eip712Domain")

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
func (_WStBtc *WStBtcSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WStBtc.Contract.Eip712Domain(&_WStBtc.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WStBtc *WStBtcCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WStBtc.Contract.Eip712Domain(&_WStBtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WStBtc *WStBtcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WStBtc *WStBtcSession) Name() (string, error) {
	return _WStBtc.Contract.Name(&_WStBtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WStBtc *WStBtcCallerSession) Name() (string, error) {
	return _WStBtc.Contract.Name(&_WStBtc.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WStBtc *WStBtcCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WStBtc *WStBtcSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WStBtc.Contract.Nonces(&_WStBtc.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WStBtc *WStBtcCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WStBtc.Contract.Nonces(&_WStBtc.CallOpts, owner)
}

// StBTC is a free data retrieval call binding the contract method 0x5ae7daa5.
//
// Solidity: function stBTC() view returns(address)
func (_WStBtc *WStBtcCaller) StBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "stBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StBTC is a free data retrieval call binding the contract method 0x5ae7daa5.
//
// Solidity: function stBTC() view returns(address)
func (_WStBtc *WStBtcSession) StBTC() (common.Address, error) {
	return _WStBtc.Contract.StBTC(&_WStBtc.CallOpts)
}

// StBTC is a free data retrieval call binding the contract method 0x5ae7daa5.
//
// Solidity: function stBTC() view returns(address)
func (_WStBtc *WStBtcCallerSession) StBTC() (common.Address, error) {
	return _WStBtc.Contract.StBTC(&_WStBtc.CallOpts)
}

// StBTCPerToken is a free data retrieval call binding the contract method 0x92ebbb2f.
//
// Solidity: function stBTCPerToken(uint256 amount) view returns(uint256)
func (_WStBtc *WStBtcCaller) StBTCPerToken(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "stBTCPerToken", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StBTCPerToken is a free data retrieval call binding the contract method 0x92ebbb2f.
//
// Solidity: function stBTCPerToken(uint256 amount) view returns(uint256)
func (_WStBtc *WStBtcSession) StBTCPerToken(amount *big.Int) (*big.Int, error) {
	return _WStBtc.Contract.StBTCPerToken(&_WStBtc.CallOpts, amount)
}

// StBTCPerToken is a free data retrieval call binding the contract method 0x92ebbb2f.
//
// Solidity: function stBTCPerToken(uint256 amount) view returns(uint256)
func (_WStBtc *WStBtcCallerSession) StBTCPerToken(amount *big.Int) (*big.Int, error) {
	return _WStBtc.Contract.StBTCPerToken(&_WStBtc.CallOpts, amount)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WStBtc *WStBtcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WStBtc *WStBtcSession) Symbol() (string, error) {
	return _WStBtc.Contract.Symbol(&_WStBtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WStBtc *WStBtcCallerSession) Symbol() (string, error) {
	return _WStBtc.Contract.Symbol(&_WStBtc.CallOpts)
}

// TokensPerStBTC is a free data retrieval call binding the contract method 0x3714b060.
//
// Solidity: function tokensPerStBTC(uint256 amount) view returns(uint256)
func (_WStBtc *WStBtcCaller) TokensPerStBTC(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "tokensPerStBTC", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerStBTC is a free data retrieval call binding the contract method 0x3714b060.
//
// Solidity: function tokensPerStBTC(uint256 amount) view returns(uint256)
func (_WStBtc *WStBtcSession) TokensPerStBTC(amount *big.Int) (*big.Int, error) {
	return _WStBtc.Contract.TokensPerStBTC(&_WStBtc.CallOpts, amount)
}

// TokensPerStBTC is a free data retrieval call binding the contract method 0x3714b060.
//
// Solidity: function tokensPerStBTC(uint256 amount) view returns(uint256)
func (_WStBtc *WStBtcCallerSession) TokensPerStBTC(amount *big.Int) (*big.Int, error) {
	return _WStBtc.Contract.TokensPerStBTC(&_WStBtc.CallOpts, amount)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WStBtc *WStBtcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WStBtc *WStBtcSession) TotalSupply() (*big.Int, error) {
	return _WStBtc.Contract.TotalSupply(&_WStBtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WStBtc *WStBtcCallerSession) TotalSupply() (*big.Int, error) {
	return _WStBtc.Contract.TotalSupply(&_WStBtc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WStBtc *WStBtcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStBtc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WStBtc *WStBtcSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.Approve(&_WStBtc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WStBtc *WStBtcTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.Approve(&_WStBtc.TransactOpts, spender, value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WStBtc *WStBtcTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WStBtc.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WStBtc *WStBtcSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WStBtc.Contract.Permit(&_WStBtc.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WStBtc *WStBtcTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WStBtc.Contract.Permit(&_WStBtc.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WStBtc *WStBtcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStBtc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WStBtc *WStBtcSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.Transfer(&_WStBtc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WStBtc *WStBtcTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.Transfer(&_WStBtc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WStBtc *WStBtcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStBtc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WStBtc *WStBtcSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.TransferFrom(&_WStBtc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WStBtc *WStBtcTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.TransferFrom(&_WStBtc.TransactOpts, from, to, value)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 wstBTCAmount) returns(uint256)
func (_WStBtc *WStBtcTransactor) Unwrap(opts *bind.TransactOpts, wstBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStBtc.contract.Transact(opts, "unwrap", wstBTCAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 wstBTCAmount) returns(uint256)
func (_WStBtc *WStBtcSession) Unwrap(wstBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.Unwrap(&_WStBtc.TransactOpts, wstBTCAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 wstBTCAmount) returns(uint256)
func (_WStBtc *WStBtcTransactorSession) Unwrap(wstBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.Unwrap(&_WStBtc.TransactOpts, wstBTCAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 stBTCAmount) returns(uint256)
func (_WStBtc *WStBtcTransactor) Wrap(opts *bind.TransactOpts, stBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStBtc.contract.Transact(opts, "wrap", stBTCAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 stBTCAmount) returns(uint256)
func (_WStBtc *WStBtcSession) Wrap(stBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.Wrap(&_WStBtc.TransactOpts, stBTCAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 stBTCAmount) returns(uint256)
func (_WStBtc *WStBtcTransactorSession) Wrap(stBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStBtc.Contract.Wrap(&_WStBtc.TransactOpts, stBTCAmount)
}

// WStBtcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WStBtc contract.
type WStBtcApprovalIterator struct {
	Event *WStBtcApproval // Event containing the contract specifics and raw log

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
func (it *WStBtcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WStBtcApproval)
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
		it.Event = new(WStBtcApproval)
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
func (it *WStBtcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WStBtcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WStBtcApproval represents a Approval event raised by the WStBtc contract.
type WStBtcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WStBtc *WStBtcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WStBtcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WStBtc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WStBtcApprovalIterator{contract: _WStBtc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WStBtc *WStBtcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WStBtcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WStBtc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WStBtcApproval)
				if err := _WStBtc.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WStBtc *WStBtcFilterer) ParseApproval(log types.Log) (*WStBtcApproval, error) {
	event := new(WStBtcApproval)
	if err := _WStBtc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WStBtcEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the WStBtc contract.
type WStBtcEIP712DomainChangedIterator struct {
	Event *WStBtcEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *WStBtcEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WStBtcEIP712DomainChanged)
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
		it.Event = new(WStBtcEIP712DomainChanged)
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
func (it *WStBtcEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WStBtcEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WStBtcEIP712DomainChanged represents a EIP712DomainChanged event raised by the WStBtc contract.
type WStBtcEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WStBtc *WStBtcFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*WStBtcEIP712DomainChangedIterator, error) {

	logs, sub, err := _WStBtc.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &WStBtcEIP712DomainChangedIterator{contract: _WStBtc.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WStBtc *WStBtcFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *WStBtcEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _WStBtc.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WStBtcEIP712DomainChanged)
				if err := _WStBtc.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_WStBtc *WStBtcFilterer) ParseEIP712DomainChanged(log types.Log) (*WStBtcEIP712DomainChanged, error) {
	event := new(WStBtcEIP712DomainChanged)
	if err := _WStBtc.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WStBtcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WStBtc contract.
type WStBtcTransferIterator struct {
	Event *WStBtcTransfer // Event containing the contract specifics and raw log

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
func (it *WStBtcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WStBtcTransfer)
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
		it.Event = new(WStBtcTransfer)
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
func (it *WStBtcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WStBtcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WStBtcTransfer represents a Transfer event raised by the WStBtc contract.
type WStBtcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WStBtc *WStBtcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WStBtcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WStBtc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WStBtcTransferIterator{contract: _WStBtc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WStBtc *WStBtcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WStBtcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WStBtc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WStBtcTransfer)
				if err := _WStBtc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WStBtc *WStBtcFilterer) ParseTransfer(log types.Log) (*WStBtcTransfer, error) {
	event := new(WStBtcTransfer)
	if err := _WStBtc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
