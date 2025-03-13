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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stBTC\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stBTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStBTC\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stBTCPerToken\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokensPerStBTC\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unwrap\",\"inputs\":[{\"name\":\"wstBTCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wrap\",\"inputs\":[{\"name\":\"stBTCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotUnwrapZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotWrapZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2612ExpiredSignature\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2612InvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoStBTCBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoWstBTCSupply\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x61016080604052346104ad576020816119fa803803809161002082856104b2565b8339810103126104ad57516001600160a01b038116908190036104ad5760405161004b6040826104b2565b60168152602081017f57726170706564205374726f6f6d20426974636f696e000000000000000000008152604051906100856040836104b2565b601682527f57726170706564205374726f6f6d20426974636f696e000000000000000000006020830152604051926100be6040856104b2565b600684526577737442544360d01b6020850152604051936100e06040866104b2565b60018552603160f81b60208601908152845190946001600160401b0382116103aa5760035490600182811c921680156104a3575b602083101461038a5781601f849311610433575b50602090601f83116001146103cb576000926103c0575b50508160011b916000199060031b1c1916176003555b8051906001600160401b0382116103aa5760045490600182811c921680156103a0575b602083101461038a5781601f84931161031a575b50602090601f83116001146102b2576000926102a7575b50508160011b916000199060031b1c1916176004555b6101c2816104d5565b610120526101cf84610683565b61014052519020918260e05251902080610100524660a0526040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261023860c0826104b2565b5190206080523060c052600880546001600160a01b0319169190911790556040516111d89081610822823960805181610ea9015260a05181610f66015260c05181610e73015260e05181610ef801526101005181610f1e015261012051816107f0015261014051816108190152f35b0151905038806101a3565b600460009081528281209350601f198516905b81811061030257509084600195949392106102e9575b505050811b016004556101b9565b015160001960f88460031b161c191690553880806102db565b929360206001819287860151815501950193016102c5565b60046000529091507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c81019160208510610380575b90601f859493920160051c01905b818110610371575061018c565b60008155849350600101610364565b9091508190610356565b634e487b7160e01b600052602260045260246000fd5b91607f1691610178565b634e487b7160e01b600052604160045260246000fd5b01519050388061013f565b600360009081528281209350601f198516905b81811061041b5750908460019594939210610402575b505050811b01600355610155565b015160001960f88460031b161c191690553880806103f4565b929360206001819287860151815501950193016103de565b60036000529091507fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c81019160208510610499575b90601f859493920160051c01905b81811061048a5750610128565b6000815584935060010161047d565b909150819061046f565b91607f1691610114565b600080fd5b601f909101601f19168101906001600160401b038211908210176103aa57604052565b9081516020811060001461056d575090601f815111610511576020815191015160208210610501571790565b6000198260200360031b1b161790565b6040519063305a27a960e01b8252602060048301528181519182602483015260005b8381106105555750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610533565b6001600160401b0381116103aa57600554600181811c91168015610679575b602082101461038a57601f8111610643575b50602092601f82116001146105de57928192936000926105d3575b50508160011b916000199060031b1c19161760055560ff90565b0151905038806105b9565b601f198216936005600052806000209160005b86811061062b5750836001959610610612575b505050811b0160055560ff90565b015160001960f88460031b161c19169055388080610604565b919260206001819286850151815501940192016105f1565b6005600052601f6020600020910160051c810190601f830160051c015b81811061066d575061059e565b60008155600101610660565b90607f169061058c565b9081516020811060001461070b575090601f8151116106af576020815191015160208210610501571790565b6040519063305a27a960e01b8252602060048301528181519182602483015260005b8381106106f35750508160006044809484010152601f80199101168101030190fd5b602082820181015160448784010152859350016106d1565b6001600160401b0381116103aa57600654600181811c91168015610817575b602082101461038a57601f81116107e1575b50602092601f821160011461077c5792819293600092610771575b50508160011b916000199060031b1c19161760065560ff90565b015190503880610757565b601f198216936006600052806000209160005b8681106107c957508360019596106107b0575b505050811b0160065560ff90565b015160001960f88460031b161c191690553880806107a2565b9192602060018192868501518155019401920161078f565b6006600052601f6020600020910160051c810190601f830160051c015b81811061080b575061073c565b600081556001016107fe565b90607f169061072a56fe608080604052600436101561001357600080fd5b60003560e01c90816306fdde0314610b8557508063095ea7b314610b5f57806318160ddd14610b4157806323b872dd14610a54578063313ce56714610a385780633644e51514610a1d5780633714b060146109715780635ae7daa51461094857806370a082311461090e5780637ecebe00146108d457806384b0196e146107d757806392ebbb2f1461073357806395d89b411461064e578063a9059cbb1461061d578063d505accf146104d2578063dd62ed3e14610481578063de0e9a3e146102bd5763ea598cb0146100e557600080fd5b346102b85760203660031901126102b85760043580156102a7576008546040516370a0823160e01b81523060048201526001600160a01b0390911690602081602481855afa90811561024257600091610271575b5060009160209160025480158015610269575b851461024e5750506064845b60405194859384926323b872dd60e01b845233600485015230602485015260448401525af1801561024257610215575b5033156101ff57600254908082018092116101e957602091600255600033815280835260408120828154019055604051908282527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef843393a3604051908152f35b634e487b7160e01b600052601160045260246000fd5b63ec442f0560e01b600052600060045260246000fd5b6102369060203d60201161023b575b61022e8183610d54565b810190610dbf565b610188565b503d610224565b6040513d6000823e3d90fd5b60649161025e6102639288610d8c565b610d9f565b94610158565b50811561014c565b906020823d60201161029f575b8161028b60209383610d54565b8101031261029c5750516000610139565b80fd5b3d915061027e565b63067f9ec160e11b60005260046000fd5b600080fd5b346102b85760203660031901126102b8576004358015610470576008546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa9081156102425760009161043e575b50600254801561042d5761025e6103299284610d8c565b3315610417576000913383528260205260408320548181106103fd57908084923384528360205203604083205580600254036002556040519081527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203392a360085460405163a9059cbb60e01b815233600482015260248101839052906020908290604490829087906001600160a01b03165af180156103f257602093506103d7575b50604051908152f35b6103ed90833d851161023b5761022e8183610d54565b6103ce565b6040513d85823e3d90fd5b63391434e360e21b84523360045260245260445250606490fd5b634b637e8f60e11b600052600060045260246000fd5b63cb133c8960e01b60005260046000fd5b90506020813d602011610468575b8161045960209383610d54565b810103126102b8575182610312565b3d915061044c565b6323e31f9560e11b60005260046000fd5b346102b85760403660031901126102b85761049a610c6b565b6104a2610c81565b6001600160a01b039182166000908152600160209081526040808320949093168252928352819020549051908152f35b346102b85760e03660031901126102b8576104eb610c6b565b6104f3610c81565b604435906064359260843560ff811681036102b857844211610608576105c96105d29160018060a01b038416968760005260076020526040600020908154916001830190556040519060208201927f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c984528a604084015260018060a01b038916606084015289608084015260a083015260c082015260c0815261059760e082610d54565b5190206105a2610e70565b906040519161190160f01b83526002830152602282015260c43591604260a4359220611091565b9092919261111a565b6001600160a01b03168481036105ef57506105ed9350610f8c565b005b84906325c0072360e11b60005260045260245260446000fd5b8463313c898160e11b60005260045260246000fd5b346102b85760403660031901126102b857610643610639610c6b565b6024359033610dd7565b602060405160018152f35b346102b85760003660031901126102b857604051600060045461067081610c97565b808452906001811690811561070f57506001146106b0575b6106ac8361069881850382610d54565b604051918291602083526020830190610c2a565b0390f35b600460009081527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b939250905b8082106106f557509091508101602001610698610688565b9192600181602092548385880101520191019092916106dd565b60ff191660208086019190915291151560051b840190910191506106989050610688565b346102b85760203660031901126102b8576008546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa908115610242576000916107a5575b5060025490811561042d5760209161025e61079d9260043590610d8c565b604051908152f35b90506020813d6020116107cf575b816107c060209383610d54565b810103126102b857518161077f565b3d91506107b3565b346102b85760003660031901126102b8576108766108147f0000000000000000000000000000000000000000000000000000000000000000610ff3565b61083d7f000000000000000000000000000000000000000000000000000000000000000061105a565b6020610884604051926108508385610d54565b600084526000368137604051958695600f60f81b875260e08588015260e0870190610c2a565b908582036040870152610c2a565b466060850152306080850152600060a085015283810360c085015281808451928381520193019160005b8281106108bd57505050500390f35b8351855286955093810193928101926001016108ae565b346102b85760203660031901126102b8576001600160a01b036108f5610c6b565b1660005260076020526020604060002054604051908152f35b346102b85760203660031901126102b8576001600160a01b0361092f610c6b565b1660005260006020526020604060002054604051908152f35b346102b85760003660031901126102b8576008546040516001600160a01b039091168152602090f35b346102b85760203660031901126102b8576008546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa908115610242576000916109eb575b5060025481156109da5760209161025e61079d9260043590610d8c565b6331d7e6d960e01b60005260046000fd5b906020823d602011610a15575b81610a0560209383610d54565b8101031261029c575051816109bd565b3d91506109f8565b346102b85760003660031901126102b857602061079d610e70565b346102b85760003660031901126102b857602060405160128152f35b346102b85760603660031901126102b857610a6d610c6b565b610a75610c81565b6001600160a01b0382166000818152600160208181526040808420338552909152909120549193604435939290918101610ab5575b506106439350610dd7565b838110610b24578415610b0e573315610af857610643946000526001602052604060002060018060a01b0333166000526020528360406000209103905584610aaa565b634a1406b160e11b600052600060045260246000fd5b63e602df0560e01b600052600060045260246000fd5b8390637dc7a0d960e11b6000523360045260245260445260646000fd5b346102b85760003660031901126102b8576020600254604051908152f35b346102b85760403660031901126102b857610643610b7b610c6b565b6024359033610f8c565b346102b85760003660031901126102b8576000600354610ba481610c97565b808452906001811690811561070f5750600114610bcb576106ac8361069881850382610d54565b600360009081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b939250905b808210610c1057509091508101602001610698610688565b919260018160209254838588010152019101909291610bf8565b919082519283825260005b848110610c56575050826000602080949584010152601f8019910116010190565b80602080928401015182828601015201610c35565b600435906001600160a01b03821682036102b857565b602435906001600160a01b03821682036102b857565b90600182811c92168015610cc7575b6020831014610cb157565b634e487b7160e01b600052602260045260246000fd5b91607f1691610ca6565b60009291815491610ce183610c97565b8083529260018116908115610d375750600114610cfd57505050565b60009081526020812093945091925b838310610d1d575060209250010190565b600181602092949394548385870101520191019190610d0c565b915050602093945060ff929192191683830152151560051b010190565b90601f8019910116810190811067ffffffffffffffff821117610d7657604052565b634e487b7160e01b600052604160045260246000fd5b818102929181159184041417156101e957565b8115610da9570490565b634e487b7160e01b600052601260045260246000fd5b908160209103126102b8575180151581036102b85790565b6001600160a01b0316908115610417576001600160a01b03169182156101ff576000828152806020526040812054828110610e565791604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815280845220818154019055604051908152a3565b916064928463391434e360e21b8452600452602452604452fd5b307f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161480610f63575b15610ecb577f000000000000000000000000000000000000000000000000000000000000000090565b60405160208101907f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f82527f000000000000000000000000000000000000000000000000000000000000000060408201527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260a08152610f5d60c082610d54565b51902090565b507f00000000000000000000000000000000000000000000000000000000000000004614610ea2565b6001600160a01b0316908115610b0e576001600160a01b0316918215610af85760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a3565b60ff811461103d5760ff811690601f821161102c5760408051926110178285610d54565b6020808552840191601f190136833783525290565b632cd44ac360e21b60005260046000fd5b5060405161105781611050816005610cd1565b0382610d54565b90565b60ff811461107e5760ff811690601f821161102c5760408051926110178285610d54565b5060405161105781611050816006610cd1565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161110e579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa15610242576000516001600160a01b038116156111025790600090600090565b50600090600190600090565b50505060009160039190565b919091600481101561118c578061113057509050565b60006001820361114b5763f645eedf60e01b60005260046000fd5b5060028103611169578263fce698f760e01b60005260045260246000fd5b9091600360009214611179575050565b6335e2f38360e21b825260045260249150fd5b634e487b7160e01b600052602160045260246000fdfea2646970667358221220e1a82ea529cad5b5b36a88b826c35ddc8ee65d0b3c66f70c5e4932b10668454964736f6c634300081b0033",
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
