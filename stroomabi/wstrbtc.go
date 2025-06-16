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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strBTC\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strBTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrBTC\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strBTCPerToken\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokensPerStrBTC\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unwrap\",\"inputs\":[{\"name\":\"wstrBTCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wrap\",\"inputs\":[{\"name\":\"strBTCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotUnwrapZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotWrapZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2612ExpiredSignature\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2612InvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoStrBTCBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoWstrBTCSupply\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610160604052348015610010575f5ffd5b506040516118b83803806118b883398101604081905261002f91610244565b6040518060400160405280601681526020017f57726170706564205374726f6f6d20426974636f696e0000000000000000000081525080604051806040016040528060018152602001603160f81b8152506040518060400160405280601681526020017f57726170706564205374726f6f6d20426974636f696e00000000000000000000815250604051806040016040528060078152602001667773747242544360c81b81525081600390816100e59190610309565b5060046100f28282610309565b50610102915083905060056101cc565b610120526101118160066101cc565b61014052815160208084019190912060e052815190820120610100524660a05261019d60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05250600880546001600160a01b0319166001600160a01b039290921691909117905561041b565b5f6020835110156101e7576101e0836101fe565b90506101f8565b816101f28482610309565b5060ff90505b92915050565b5f5f829050601f81511115610231578260405163305a27a960e01b815260040161022891906103c3565b60405180910390fd5b805161023c826103f8565b179392505050565b5f60208284031215610254575f5ffd5b81516001600160a01b038116811461026a575f5ffd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061029957607f821691505b6020821081036102b757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561030457805f5260205f20601f840160051c810160208510156102e25750805b601f840160051c820191505b81811015610301575f81556001016102ee565b50505b505050565b81516001600160401b0381111561032257610322610271565b610336816103308454610285565b846102bd565b6020601f821160018114610368575f83156103515750848201515b5f19600385901b1c1916600184901b178455610301565b5f84815260208120601f198516915b828110156103975787850151825560209485019460019092019101610377565b50848210156103b457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102b7575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161144c61046c5f395f610ba001525f610b7301525f610b1c01525f610af401525f610a4f01525f610a7901525f610aa3015261144c5ff3fe608060405234801561000f575f5ffd5b5060043610610111575f3560e01c80637ecebe001161009e578063a9059cbb1161006e578063a9059cbb14610241578063d505accf14610254578063dd62ed3e14610269578063de0e9a3e146102a1578063ea598cb0146102b4575f5ffd5b80637ecebe00146101f857806384b0196e1461020b57806395d89b4114610226578063a70c2a4a1461022e575f5ffd5b80632ba8c1c3116100e45780632ba8c1c31461017b578063313ce5671461018e5780633644e5151461019d5780636a390c1c146101a557806370a08231146101d0575f5ffd5b806306fdde0314610115578063095ea7b31461013357806318160ddd1461015657806323b872dd14610168575b5f5ffd5b61011d6102c7565b60405161012a919061113d565b60405180910390f35b610146610141366004611171565b610357565b604051901515815260200161012a565b6002545b60405190815260200161012a565b610146610176366004611199565b610370565b61015a6101893660046111d3565b610393565b6040516012815260200161012a565b61015a61044c565b6008546101b8906001600160a01b031681565b6040516001600160a01b03909116815260200161012a565b61015a6101de3660046111ea565b6001600160a01b03165f9081526020819052604090205490565b61015a6102063660046111ea565b61045a565b610213610477565b60405161012a9796959493929190611203565b61011d6104b9565b61015a61023c3660046111d3565b6104c8565b61014661024f366004611171565b61056f565b610267610262366004611299565b61057c565b005b61015a610277366004611306565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b61015a6102af3660046111d3565b6106b7565b61015a6102c23660046111d3565b61080d565b6060600380546102d690611337565b80601f016020809104026020016040519081016040528092919081815260200182805461030290611337565b801561034d5780601f106103245761010080835404028352916020019161034d565b820191905f5260205f20905b81548152906001019060200180831161033057829003601f168201915b5050505050905090565b5f33610364818585610959565b60019150505b92915050565b5f3361037d85828561096b565b6103888585856109e6565b506001949350505050565b6008546040516370a0823160e01b81523060048201525f9182916001600160a01b03909116906370a0823190602401602060405180830381865afa1580156103dd573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610401919061136f565b90505f61040d60025490565b9050805f0361042f5760405163bd7cbed760e01b815260040160405180910390fd5b8061043a858461139a565b61044491906113b1565b949350505050565b5f610455610a43565b905090565b6001600160a01b0381165f9081526007602052604081205461036a565b5f6060805f5f5f6060610488610b6c565b610490610b99565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b6060600480546102d690611337565b6008546040516370a0823160e01b81523060048201525f9182916001600160a01b03909116906370a0823190602401602060405180830381865afa158015610512573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610536919061136f565b90505f61054260025490565b9050815f036105645760405163081e49d760e21b815260040160405180910390fd5b8161043a858361139a565b5f336103648185856109e6565b834211156105a55760405163313c898160e11b8152600481018590526024015b60405180910390fd5b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886105f08c6001600160a01b03165f90815260076020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61064a82610bc6565b90505f61065982878787610bf2565b9050896001600160a01b0316816001600160a01b0316146106a0576040516325c0072360e11b81526001600160a01b0380831660048301528b16602482015260440161059c565b6106ab8a8a8a610959565b50505050505050505050565b5f815f036106d8576040516323e31f9560e11b815260040160405180910390fd5b6008546040516370a0823160e01b81523060048201525f916001600160a01b0316906370a0823190602401602060405180830381865afa15801561071e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610742919061136f565b90505f61074e60025490565b9050805f036107705760405163bd7cbed760e01b815260040160405180910390fd5b5f8161077c848761139a565b61078691906113b1565b90506107923386610c1e565b60085460405163a9059cbb60e01b8152336004820152602481018390526001600160a01b039091169063a9059cbb906044016020604051808303815f875af11580156107e0573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080491906113d0565b50949350505050565b5f815f0361082e5760405163067f9ec160e11b815260040160405180910390fd5b6008546040516370a0823160e01b81523060048201525f916001600160a01b0316906370a0823190602401602060405180830381865afa158015610874573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610898919061136f565b90505f6108a460025490565b90505f8115806108b2575082155b156108be5750836108d6565b826108c9838761139a565b6108d391906113b1565b90505b6008546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af115801561092a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061094e91906113d0565b506104443382610c56565b6109668383836001610c8a565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f1981146109e057818110156109d257604051637dc7a0d960e11b81526001600160a01b0384166004820152602481018290526044810183905260640161059c565b6109e084848484035f610c8a565b50505050565b6001600160a01b038316610a0f57604051634b637e8f60e11b81525f600482015260240161059c565b6001600160a01b038216610a385760405163ec442f0560e01b81525f600482015260240161059c565b610966838383610d5c565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015610a9b57507f000000000000000000000000000000000000000000000000000000000000000046145b15610ac557507f000000000000000000000000000000000000000000000000000000000000000090565b610455604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60606104557f00000000000000000000000000000000000000000000000000000000000000006005610e82565b60606104557f00000000000000000000000000000000000000000000000000000000000000006006610e82565b5f61036a610bd2610a43565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610c0288888888610f2b565b925092509250610c128282610ff3565b50909695505050505050565b6001600160a01b038216610c4757604051634b637e8f60e11b81525f600482015260240161059c565b610c52825f83610d5c565b5050565b6001600160a01b038216610c7f5760405163ec442f0560e01b81525f600482015260240161059c565b610c525f8383610d5c565b6001600160a01b038416610cb35760405163e602df0560e01b81525f600482015260240161059c565b6001600160a01b038316610cdc57604051634a1406b160e11b81525f600482015260240161059c565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156109e057826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610d4e91815260200190565b60405180910390a350505050565b6001600160a01b038316610d86578060025f828254610d7b91906113ef565b90915550610df69050565b6001600160a01b0383165f9081526020819052604090205481811015610dd85760405163391434e360e21b81526001600160a01b0385166004820152602481018290526044810183905260640161059c565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216610e1257600280548290039055610e30565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e7591815260200190565b60405180910390a3505050565b606060ff8314610e9c57610e95836110ab565b905061036a565b818054610ea890611337565b80601f0160208091040260200160405190810160405280929190818152602001828054610ed490611337565b8015610f1f5780601f10610ef657610100808354040283529160200191610f1f565b820191905f5260205f20905b815481529060010190602001808311610f0257829003601f168201915b5050505050905061036a565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610f6457505f91506003905082610fe9565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610fb5573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610fe057505f925060019150829050610fe9565b92505f91508190505b9450945094915050565b5f82600381111561100657611006611402565b0361100f575050565b600182600381111561102357611023611402565b036110415760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561105557611055611402565b036110765760405163fce698f760e01b81526004810182905260240161059c565b600382600381111561108a5761108a611402565b03610c52576040516335e2f38360e21b81526004810182905260240161059c565b60605f6110b7836110e8565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f81111561036a57604051632cd44ac360e21b815260040160405180910390fd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61114f602083018461110f565b9392505050565b80356001600160a01b038116811461116c575f5ffd5b919050565b5f5f60408385031215611182575f5ffd5b61118b83611156565b946020939093013593505050565b5f5f5f606084860312156111ab575f5ffd5b6111b484611156565b92506111c260208501611156565b929592945050506040919091013590565b5f602082840312156111e3575f5ffd5b5035919050565b5f602082840312156111fa575f5ffd5b61114f82611156565b60ff60f81b8816815260e060208201525f61122160e083018961110f565b8281036040840152611233818961110f565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b8181101561128857835183526020938401939092019160010161126a565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a0312156112af575f5ffd5b6112b888611156565b96506112c660208901611156565b95506040880135945060608801359350608088013560ff811681146112e9575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215611317575f5ffd5b61132083611156565b915061132e60208401611156565b90509250929050565b600181811c9082168061134b57607f821691505b60208210810361136957634e487b7160e01b5f52602260045260245ffd5b50919050565b5f6020828403121561137f575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761036a5761036a611386565b5f826113cb57634e487b7160e01b5f52601260045260245ffd5b500490565b5f602082840312156113e0575f5ffd5b8151801515811461114f575f5ffd5b8082018082111561036a5761036a611386565b634e487b7160e01b5f52602160045260245ffdfea26469706673582212208fa3f19142a11385683f029270091c8b220d40ed4f0ce42fcf084f32f2eba9ad64736f6c634300081b0033",
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

// StrBTC is a free data retrieval call binding the contract method 0x6a390c1c.
//
// Solidity: function strBTC() view returns(address)
func (_WStrBtc *WStrBtcCaller) StrBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "strBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrBTC is a free data retrieval call binding the contract method 0x6a390c1c.
//
// Solidity: function strBTC() view returns(address)
func (_WStrBtc *WStrBtcSession) StrBTC() (common.Address, error) {
	return _WStrBtc.Contract.StrBTC(&_WStrBtc.CallOpts)
}

// StrBTC is a free data retrieval call binding the contract method 0x6a390c1c.
//
// Solidity: function strBTC() view returns(address)
func (_WStrBtc *WStrBtcCallerSession) StrBTC() (common.Address, error) {
	return _WStrBtc.Contract.StrBTC(&_WStrBtc.CallOpts)
}

// StrBTCPerToken is a free data retrieval call binding the contract method 0x2ba8c1c3.
//
// Solidity: function strBTCPerToken(uint256 amount) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) StrBTCPerToken(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "strBTCPerToken", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrBTCPerToken is a free data retrieval call binding the contract method 0x2ba8c1c3.
//
// Solidity: function strBTCPerToken(uint256 amount) view returns(uint256)
func (_WStrBtc *WStrBtcSession) StrBTCPerToken(amount *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.StrBTCPerToken(&_WStrBtc.CallOpts, amount)
}

// StrBTCPerToken is a free data retrieval call binding the contract method 0x2ba8c1c3.
//
// Solidity: function strBTCPerToken(uint256 amount) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) StrBTCPerToken(amount *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.StrBTCPerToken(&_WStrBtc.CallOpts, amount)
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

// TokensPerStrBTC is a free data retrieval call binding the contract method 0xa70c2a4a.
//
// Solidity: function tokensPerStrBTC(uint256 amount) view returns(uint256)
func (_WStrBtc *WStrBtcCaller) TokensPerStrBTC(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WStrBtc.contract.Call(opts, &out, "tokensPerStrBTC", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerStrBTC is a free data retrieval call binding the contract method 0xa70c2a4a.
//
// Solidity: function tokensPerStrBTC(uint256 amount) view returns(uint256)
func (_WStrBtc *WStrBtcSession) TokensPerStrBTC(amount *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.TokensPerStrBTC(&_WStrBtc.CallOpts, amount)
}

// TokensPerStrBTC is a free data retrieval call binding the contract method 0xa70c2a4a.
//
// Solidity: function tokensPerStrBTC(uint256 amount) view returns(uint256)
func (_WStrBtc *WStrBtcCallerSession) TokensPerStrBTC(amount *big.Int) (*big.Int, error) {
	return _WStrBtc.Contract.TokensPerStrBTC(&_WStrBtc.CallOpts, amount)
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

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 wstrBTCAmount) returns(uint256)
func (_WStrBtc *WStrBtcTransactor) Unwrap(opts *bind.TransactOpts, wstrBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "unwrap", wstrBTCAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 wstrBTCAmount) returns(uint256)
func (_WStrBtc *WStrBtcSession) Unwrap(wstrBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.Unwrap(&_WStrBtc.TransactOpts, wstrBTCAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 wstrBTCAmount) returns(uint256)
func (_WStrBtc *WStrBtcTransactorSession) Unwrap(wstrBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.Unwrap(&_WStrBtc.TransactOpts, wstrBTCAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 strBTCAmount) returns(uint256)
func (_WStrBtc *WStrBtcTransactor) Wrap(opts *bind.TransactOpts, strBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStrBtc.contract.Transact(opts, "wrap", strBTCAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 strBTCAmount) returns(uint256)
func (_WStrBtc *WStrBtcSession) Wrap(strBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.Wrap(&_WStrBtc.TransactOpts, strBTCAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 strBTCAmount) returns(uint256)
func (_WStrBtc *WStrBtcTransactorSession) Wrap(strBTCAmount *big.Int) (*types.Transaction, error) {
	return _WStrBtc.Contract.Wrap(&_WStrBtc.TransactOpts, strBTCAmount)
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
