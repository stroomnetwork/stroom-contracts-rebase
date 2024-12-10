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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stBTC\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stBTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStBTC\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stBTCPerToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokensPerStBTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unwrap\",\"inputs\":[{\"name\":\"wstBTCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wrap\",\"inputs\":[{\"name\":\"stBTCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2612ExpiredSignature\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2612InvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x61016060405234801561001157600080fd5b50604051611a31380380611a3183398101604081905261003091610247565b6040518060400160405280601681526020017f57726170706564205374726f6f6d20426974636f696e0000000000000000000081525080604051806040016040528060018152602001603160f81b8152506040518060400160405280601681526020017f57726170706564205374726f6f6d20426974636f696e000000000000000000008152506040518060400160405280600681526020016577737442544360d01b81525081600390816100e59190610316565b5060046100f28282610316565b50610102915083905060056101cd565b610120526101118160066101cd565b61014052815160208084019190912060e052815190820120610100524660a05261019e60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60805250503060c05250600880546001600160a01b0319166001600160a01b0392909216919091179055610446565b60006020835110156101e9576101e283610200565b90506101fa565b816101f48482610316565b5060ff90505b92915050565b600080829050601f81511115610234578260405163305a27a960e01b815260040161022b91906103d4565b60405180910390fd5b805161023f82610422565b179392505050565b60006020828403121561025957600080fd5b81516001600160a01b038116811461027057600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806102a157607f821691505b6020821081036102c157634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561031157806000526020600020601f840160051c810160208510156102ee5750805b601f840160051c820191505b8181101561030e57600081556001016102fa565b50505b505050565b81516001600160401b0381111561032f5761032f610277565b6103438161033d845461028d565b846102c7565b6020601f821160018114610377576000831561035f5750848201515b600019600385901b1c1916600184901b17845561030e565b600084815260208120601f198516915b828110156103a75787850151825560209485019460019092019101610387565b50848210156103c55786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b602081526000825180602084015260005b8181101561040257602081860181015160408684010152016103e5565b506000604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102c15760001960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516115916104a06000396000610c8f01526000610c6201526000610c0a01526000610be201526000610b3d01526000610b6701526000610b9101526115916000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806370a08231116100ad578063a9059cbb11610071578063a9059cbb14610248578063d505accf1461025b578063dd62ed3e14610270578063de0e9a3e146102a9578063ea598cb0146102bc57600080fd5b806370a08231146101e15780637ecebe001461020a57806384b0196e1461021d57806393211cd61461023857806395d89b411461024057600080fd5b80632792949d116100f45780632792949d1461018c578063313ce567146101975780633644e515146101a65780635ae7daa5146101ae5780635b93a8b6146101d957600080fd5b806306fdde0314610126578063095ea7b31461014457806318160ddd1461016757806323b872dd14610179575b600080fd5b61012e6102cf565b60405161013b919061125f565b60405180910390f35b610157610152366004611295565b610361565b604051901515815260200161013b565b6002545b60405190815260200161013b565b6101576101873660046112bf565b61037b565b61016b6305f5e10081565b6040516012815260200161013b565b61016b61039f565b6008546101c1906001600160a01b031681565b6040516001600160a01b03909116815260200161013b565b61016b6103ae565b61016b6101ef3660046112fc565b6001600160a01b031660009081526020819052604090205490565b61016b6102183660046112fc565b61049e565b6102256104bc565b60405161013b9796959493929190611317565b61016b610502565b61012e6105e1565b610157610256366004611295565b6105f0565b61026e6102693660046113af565b6105fe565b005b61016b61027e366004611422565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b61016b6102b7366004611455565b610738565b61016b6102ca366004611455565b6108fc565b6060600380546102de9061146e565b80601f016020809104026020016040519081016040528092919081815260200182805461030a9061146e565b80156103575780601f1061032c57610100808354040283529160200191610357565b820191906000526020600020905b81548152906001019060200180831161033a57829003601f168201915b5050505050905090565b60003361036f818585610a41565b60019150505b92915050565b600033610389858285610a53565b610394858585610ad1565b506001949350505050565b60006103a9610b30565b905090565b6008546040516370a0823160e01b815230600482015260009182916001600160a01b03909116906370a0823190602401602060405180830381865afa1580156103fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061041f91906114a8565b9050600061042c60025490565b90506000811161047e5760405162461bcd60e51b81526020600482015260186024820152777773744254433a204e6f2077737442544320737570706c7960401b60448201526064015b60405180910390fd5b8061048d6305f5e100846114d7565b61049791906114ee565b9250505090565b6001600160a01b038116600090815260076020526040812054610375565b6000606080600080600060606104d0610c5b565b6104d8610c88565b60408051600080825260208201909252600f60f81b9b939a50919850469750309650945092509050565b6008546040516370a0823160e01b815230600482015260009182916001600160a01b03909116906370a0823190602401602060405180830381865afa15801561054f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057391906114a8565b9050600061058060025490565b9050600082116105d25760405162461bcd60e51b815260206004820152601860248201527f7773744254433a204e6f2073744254432062616c616e636500000000000000006044820152606401610475565b8161048d6305f5e100836114d7565b6060600480546102de9061146e565b60003361036f818585610ad1565b834211156106225760405163313c898160e11b815260048101859052602401610475565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861066f8c6001600160a01b0316600090815260076020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e00160405160208183030381529060405280519060200120905060006106ca82610cb5565b905060006106da82878787610ce2565b9050896001600160a01b0316816001600160a01b031614610721576040516325c0072360e11b81526001600160a01b0380831660048301528b166024820152604401610475565b61072c8a8a8a610a41565b50505050505050505050565b60008082116107935760405162461bcd60e51b815260206004820152602160248201527f7773744254433a2043616e6e6f7420756e77726170207a65726f2077737442546044820152604360f81b6064820152608401610475565b6008546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa1580156107dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080091906114a8565b9050600061080d60025490565b90506000811161085a5760405162461bcd60e51b81526020600482015260186024820152777773744254433a204e6f2077737442544320737570706c7960401b6044820152606401610475565b60008161086784876114d7565b61087191906114ee565b905061087d3386610d10565b60085460405163a9059cbb60e01b8152336004820152602481018390526001600160a01b039091169063a9059cbb906044015b6020604051808303816000875af11580156108cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f39190611510565b50949350505050565b600080821161094d5760405162461bcd60e51b815260206004820152601e60248201527f7773744254433a2043616e6e6f742077726170207a65726f20737442544300006044820152606401610475565b6008546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa158015610996573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ba91906114a8565b905060006109c760025490565b905060008115806109d6575082155b156109e25750836109fa565b826109ed83876114d7565b6109f791906114ee565b90505b610a043382610d4a565b6008546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016108b0565b610a4e8383836001610d80565b505050565b6001600160a01b038381166000908152600160209081526040808320938616835292905220546000198114610acb5781811015610abc57604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610475565b610acb84848484036000610d80565b50505050565b6001600160a01b038316610afb57604051634b637e8f60e11b815260006004820152602401610475565b6001600160a01b038216610b255760405163ec442f0560e01b815260006004820152602401610475565b610a4e838383610e55565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015610b8957507f000000000000000000000000000000000000000000000000000000000000000046145b15610bb357507f000000000000000000000000000000000000000000000000000000000000000090565b6103a9604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60606103a97f00000000000000000000000000000000000000000000000000000000000000006005610f7f565b60606103a97f00000000000000000000000000000000000000000000000000000000000000006006610f7f565b6000610375610cc2610b30565b8360405161190160f01b8152600281019290925260228201526042902090565b600080600080610cf48888888861102a565b925092509250610d0482826110f9565b50909695505050505050565b6001600160a01b038216610d3a57604051634b637e8f60e11b815260006004820152602401610475565b610d4682600083610e55565b5050565b6001600160a01b038216610d745760405163ec442f0560e01b815260006004820152602401610475565b610d4660008383610e55565b6001600160a01b038416610daa5760405163e602df0560e01b815260006004820152602401610475565b6001600160a01b038316610dd457604051634a1406b160e11b815260006004820152602401610475565b6001600160a01b0380851660009081526001602090815260408083209387168352929052208290558015610acb57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610e4791815260200190565b60405180910390a350505050565b6001600160a01b038316610e80578060026000828254610e759190611532565b90915550610ef29050565b6001600160a01b03831660009081526020819052604090205481811015610ed35760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610475565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b038216610f0e57600280548290039055610f2d565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f7291815260200190565b60405180910390a3505050565b606060ff8314610f9957610f92836111b2565b9050610375565b818054610fa59061146e565b80601f0160208091040260200160405190810160405280929190818152602001828054610fd19061146e565b801561101e5780601f10610ff35761010080835404028352916020019161101e565b820191906000526020600020905b81548152906001019060200180831161100157829003601f168201915b50505050509050610375565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561106557506000915060039050826110ef565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156110b9573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166110e5575060009250600191508290506110ef565b9250600091508190505b9450945094915050565b600082600381111561110d5761110d611545565b03611116575050565b600182600381111561112a5761112a611545565b036111485760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561115c5761115c611545565b0361117d5760405163fce698f760e01b815260048101829052602401610475565b600382600381111561119157611191611545565b03610d46576040516335e2f38360e21b815260048101829052602401610475565b606060006111bf836111f1565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b600060ff8216601f81111561037557604051632cd44ac360e21b815260040160405180910390fd5b6000815180845260005b8181101561123f57602081850181015186830182015201611223565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006112726020830184611219565b9392505050565b80356001600160a01b038116811461129057600080fd5b919050565b600080604083850312156112a857600080fd5b6112b183611279565b946020939093013593505050565b6000806000606084860312156112d457600080fd5b6112dd84611279565b92506112eb60208501611279565b929592945050506040919091013590565b60006020828403121561130e57600080fd5b61127282611279565b60ff60f81b8816815260e06020820152600061133660e0830189611219565b82810360408401526113488189611219565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501528451808252602080870193509091019060005b8181101561139e578351835260209384019390920191600101611380565b50909b9a5050505050505050505050565b600080600080600080600060e0888a0312156113ca57600080fd5b6113d388611279565b96506113e160208901611279565b95506040880135945060608801359350608088013560ff8116811461140557600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561143557600080fd5b61143e83611279565b915061144c60208401611279565b90509250929050565b60006020828403121561146757600080fd5b5035919050565b600181811c9082168061148257607f821691505b6020821081036114a257634e487b7160e01b600052602260045260246000fd5b50919050565b6000602082840312156114ba57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610375576103756114c1565b60008261150b57634e487b7160e01b600052601260045260246000fd5b500490565b60006020828403121561152257600080fd5b8151801515811461127257600080fd5b80820180821115610375576103756114c1565b634e487b7160e01b600052602160045260246000fdfea26469706673582212207f859fb534b9af676615775302e7c163a2aa1e4054ed6c500983904d3b992de964736f6c634300081b0033",
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

// StBTCPerToken is a free data retrieval call binding the contract method 0x5b93a8b6.
//
// Solidity: function stBTCPerToken() view returns(uint256)
func (_WStBtc *WStBtcCaller) StBTCPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "stBTCPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StBTCPerToken is a free data retrieval call binding the contract method 0x5b93a8b6.
//
// Solidity: function stBTCPerToken() view returns(uint256)
func (_WStBtc *WStBtcSession) StBTCPerToken() (*big.Int, error) {
	return _WStBtc.Contract.StBTCPerToken(&_WStBtc.CallOpts)
}

// StBTCPerToken is a free data retrieval call binding the contract method 0x5b93a8b6.
//
// Solidity: function stBTCPerToken() view returns(uint256)
func (_WStBtc *WStBtcCallerSession) StBTCPerToken() (*big.Int, error) {
	return _WStBtc.Contract.StBTCPerToken(&_WStBtc.CallOpts)
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

// TokensPerStBTC is a free data retrieval call binding the contract method 0x93211cd6.
//
// Solidity: function tokensPerStBTC() view returns(uint256)
func (_WStBtc *WStBtcCaller) TokensPerStBTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WStBtc.contract.Call(opts, &out, "tokensPerStBTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerStBTC is a free data retrieval call binding the contract method 0x93211cd6.
//
// Solidity: function tokensPerStBTC() view returns(uint256)
func (_WStBtc *WStBtcSession) TokensPerStBTC() (*big.Int, error) {
	return _WStBtc.Contract.TokensPerStBTC(&_WStBtc.CallOpts)
}

// TokensPerStBTC is a free data retrieval call binding the contract method 0x93211cd6.
//
// Solidity: function tokensPerStBTC() view returns(uint256)
func (_WStBtc *WStBtcCallerSession) TokensPerStBTC() (*big.Int, error) {
	return _WStBtc.Contract.TokensPerStBTC(&_WStBtc.CallOpts)
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
