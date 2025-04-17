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
	Bin: "0x61016080604052346104ae576020816119fb803803809161002082856104b3565b8339810103126104ae57516001600160a01b038116908190036104ae5760405161004b6040826104b3565b60168152602081017f57726170706564205374726f6f6d20426974636f696e000000000000000000008152604051906100856040836104b3565b601682527f57726170706564205374726f6f6d20426974636f696e000000000000000000006020830152604051926100be6040856104b3565b60078452667773747242544360c81b6020850152604051936100e16040866104b3565b60018552603160f81b60208601908152845190946001600160401b0382116103ab5760035490600182811c921680156104a4575b602083101461038b5781601f849311610434575b50602090601f83116001146103cc576000926103c1575b50508160011b916000199060031b1c1916176003555b8051906001600160401b0382116103ab5760045490600182811c921680156103a1575b602083101461038b5781601f84931161031b575b50602090601f83116001146102b3576000926102a8575b50508160011b916000199060031b1c1916176004555b6101c3816104d6565b610120526101d084610684565b61014052519020918260e05251902080610100524660a0526040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261023960c0826104b3565b5190206080523060c052600880546001600160a01b0319169190911790556040516111d89081610823823960805181610ea9015260a05181610f66015260c05181610e73015260e05181610ef801526101005181610f1e01526101205181610800015261014051816108290152f35b0151905038806101a4565b600460009081528281209350601f198516905b81811061030357509084600195949392106102ea575b505050811b016004556101ba565b015160001960f88460031b161c191690553880806102dc565b929360206001819287860151815501950193016102c6565b60046000529091507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c81019160208510610381575b90601f859493920160051c01905b818110610372575061018d565b60008155849350600101610365565b9091508190610357565b634e487b7160e01b600052602260045260246000fd5b91607f1691610179565b634e487b7160e01b600052604160045260246000fd5b015190503880610140565b600360009081528281209350601f198516905b81811061041c5750908460019594939210610403575b505050811b01600355610156565b015160001960f88460031b161c191690553880806103f5565b929360206001819287860151815501950193016103df565b60036000529091507fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c8101916020851061049a575b90601f859493920160051c01905b81811061048b5750610129565b6000815584935060010161047e565b9091508190610470565b91607f1691610115565b600080fd5b601f909101601f19168101906001600160401b038211908210176103ab57604052565b9081516020811060001461056e575090601f815111610512576020815191015160208210610502571790565b6000198260200360031b1b161790565b6040519063305a27a960e01b8252602060048301528181519182602483015260005b8381106105565750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610534565b6001600160401b0381116103ab57600554600181811c9116801561067a575b602082101461038b57601f8111610644575b50602092601f82116001146105df57928192936000926105d4575b50508160011b916000199060031b1c19161760055560ff90565b0151905038806105ba565b601f198216936005600052806000209160005b86811061062c5750836001959610610613575b505050811b0160055560ff90565b015160001960f88460031b161c19169055388080610605565b919260206001819286850151815501940192016105f2565b6005600052601f6020600020910160051c810190601f830160051c015b81811061066e575061059f565b60008155600101610661565b90607f169061058d565b9081516020811060001461070c575090601f8151116106b0576020815191015160208210610502571790565b6040519063305a27a960e01b8252602060048301528181519182602483015260005b8381106106f45750508160006044809484010152601f80199101168101030190fd5b602082820181015160448784010152859350016106d2565b6001600160401b0381116103ab57600654600181811c91168015610818575b602082101461038b57601f81116107e2575b50602092601f821160011461077d5792819293600092610772575b50508160011b916000199060031b1c19161760065560ff90565b015190503880610758565b601f198216936006600052806000209160005b8681106107ca57508360019596106107b1575b505050811b0160065560ff90565b015160001960f88460031b161c191690553880806107a3565b91926020600181928685015181550194019201610790565b6006600052601f6020600020910160051c810190601f830160051c015b81811061080c575061073d565b600081556001016107ff565b90607f169061072b56fe608080604052600436101561001357600080fd5b60003560e01c90816306fdde0314610b8557508063095ea7b314610b5f57806318160ddd14610b4157806323b872dd14610a545780632ba8c1c3146109b8578063313ce5671461099c5780633644e515146109815780636a390c1c1461095857806370a082311461091e5780637ecebe00146108e457806384b0196e146107e757806395d89b4114610702578063a70c2a4a1461064e578063a9059cbb1461061d578063d505accf146104d2578063dd62ed3e14610481578063de0e9a3e146102bd5763ea598cb0146100e557600080fd5b346102b85760203660031901126102b85760043580156102a7576008546040516370a0823160e01b81523060048201526001600160a01b0390911690602081602481855afa90811561024257600091610271575b5060009160209160025480158015610269575b851461024e5750506064845b60405194859384926323b872dd60e01b845233600485015230602485015260448401525af1801561024257610215575b5033156101ff57600254908082018092116101e957602091600255600033815280835260408120828154019055604051908282527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef843393a3604051908152f35b634e487b7160e01b600052601160045260246000fd5b63ec442f0560e01b600052600060045260246000fd5b6102369060203d60201161023b575b61022e8183610d54565b810190610dbf565b610188565b503d610224565b6040513d6000823e3d90fd5b60649161025e6102639288610d8c565b610d9f565b94610158565b50811561014c565b906020823d60201161029f575b8161028b60209383610d54565b8101031261029c5750516000610139565b80fd5b3d915061027e565b63067f9ec160e11b60005260046000fd5b600080fd5b346102b85760203660031901126102b8576004358015610470576008546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa9081156102425760009161043e575b50600254801561042d5761025e6103299284610d8c565b3315610417576000913383528260205260408320548181106103fd57908084923384528360205203604083205580600254036002556040519081527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203392a360085460405163a9059cbb60e01b815233600482015260248101839052906020908290604490829087906001600160a01b03165af180156103f257602093506103d7575b50604051908152f35b6103ed90833d851161023b5761022e8183610d54565b6103ce565b6040513d85823e3d90fd5b63391434e360e21b84523360045260245260445250606490fd5b634b637e8f60e11b600052600060045260246000fd5b63bd7cbed760e01b60005260046000fd5b90506020813d602011610468575b8161045960209383610d54565b810103126102b8575182610312565b3d915061044c565b6323e31f9560e11b60005260046000fd5b346102b85760403660031901126102b85761049a610c6b565b6104a2610c81565b6001600160a01b039182166000908152600160209081526040808320949093168252928352819020549051908152f35b346102b85760e03660031901126102b8576104eb610c6b565b6104f3610c81565b604435906064359260843560ff811681036102b857844211610608576105c96105d29160018060a01b038416968760005260076020526040600020908154916001830190556040519060208201927f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c984528a604084015260018060a01b038916606084015289608084015260a083015260c082015260c0815261059760e082610d54565b5190206105a2610e70565b906040519161190160f01b83526002830152602282015260c43591604260a4359220611091565b9092919261111a565b6001600160a01b03168481036105ef57506105ed9350610f8c565b005b84906325c0072360e11b60005260045260245260446000fd5b8463313c898160e11b60005260045260246000fd5b346102b85760403660031901126102b857610643610639610c6b565b6024359033610dd7565b602060405160018152f35b346102b85760203660031901126102b8576008546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa908115610242576000916106d0575b5060025481156106bf5760209161025e6106b79260043590610d8c565b604051908152f35b63081e49d760e21b60005260046000fd5b906020823d6020116106fa575b816106ea60209383610d54565b8101031261029c5750518161069a565b3d91506106dd565b346102b85760003660031901126102b857604051600060045461072481610c97565b80845290600181169081156107c35750600114610764575b6107608361074c81850382610d54565b604051918291602083526020830190610c2a565b0390f35b600460009081527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b939250905b8082106107a95750909150810160200161074c61073c565b919260018160209254838588010152019101909291610791565b60ff191660208086019190915291151560051b8401909101915061074c905061073c565b346102b85760003660031901126102b8576108866108247f0000000000000000000000000000000000000000000000000000000000000000610ff3565b61084d7f000000000000000000000000000000000000000000000000000000000000000061105a565b6020610894604051926108608385610d54565b600084526000368137604051958695600f60f81b875260e08588015260e0870190610c2a565b908582036040870152610c2a565b466060850152306080850152600060a085015283810360c085015281808451928381520193019160005b8281106108cd57505050500390f35b8351855286955093810193928101926001016108be565b346102b85760203660031901126102b8576001600160a01b03610905610c6b565b1660005260076020526020604060002054604051908152f35b346102b85760203660031901126102b8576001600160a01b0361093f610c6b565b1660005260006020526020604060002054604051908152f35b346102b85760003660031901126102b8576008546040516001600160a01b039091168152602090f35b346102b85760003660031901126102b85760206106b7610e70565b346102b85760003660031901126102b857602060405160128152f35b346102b85760203660031901126102b8576008546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa90811561024257600091610a22575b5060025490811561042d5760209161025e6106b79260043590610d8c565b90506020813d602011610a4c575b81610a3d60209383610d54565b810103126102b8575181610a04565b3d9150610a30565b346102b85760603660031901126102b857610a6d610c6b565b610a75610c81565b6001600160a01b0382166000818152600160208181526040808420338552909152909120549193604435939290918101610ab5575b506106439350610dd7565b838110610b24578415610b0e573315610af857610643946000526001602052604060002060018060a01b0333166000526020528360406000209103905584610aaa565b634a1406b160e11b600052600060045260246000fd5b63e602df0560e01b600052600060045260246000fd5b8390637dc7a0d960e11b6000523360045260245260445260646000fd5b346102b85760003660031901126102b8576020600254604051908152f35b346102b85760403660031901126102b857610643610b7b610c6b565b6024359033610f8c565b346102b85760003660031901126102b8576000600354610ba481610c97565b80845290600181169081156107c35750600114610bcb576107608361074c81850382610d54565b600360009081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b939250905b808210610c105750909150810160200161074c61073c565b919260018160209254838588010152019101909291610bf8565b919082519283825260005b848110610c56575050826000602080949584010152601f8019910116010190565b80602080928401015182828601015201610c35565b600435906001600160a01b03821682036102b857565b602435906001600160a01b03821682036102b857565b90600182811c92168015610cc7575b6020831014610cb157565b634e487b7160e01b600052602260045260246000fd5b91607f1691610ca6565b60009291815491610ce183610c97565b8083529260018116908115610d375750600114610cfd57505050565b60009081526020812093945091925b838310610d1d575060209250010190565b600181602092949394548385870101520191019190610d0c565b915050602093945060ff929192191683830152151560051b010190565b90601f8019910116810190811067ffffffffffffffff821117610d7657604052565b634e487b7160e01b600052604160045260246000fd5b818102929181159184041417156101e957565b8115610da9570490565b634e487b7160e01b600052601260045260246000fd5b908160209103126102b8575180151581036102b85790565b6001600160a01b0316908115610417576001600160a01b03169182156101ff576000828152806020526040812054828110610e565791604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815280845220818154019055604051908152a3565b916064928463391434e360e21b8452600452602452604452fd5b307f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161480610f63575b15610ecb577f000000000000000000000000000000000000000000000000000000000000000090565b60405160208101907f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f82527f000000000000000000000000000000000000000000000000000000000000000060408201527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260a08152610f5d60c082610d54565b51902090565b507f00000000000000000000000000000000000000000000000000000000000000004614610ea2565b6001600160a01b0316908115610b0e576001600160a01b0316918215610af85760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a3565b60ff811461103d5760ff811690601f821161102c5760408051926110178285610d54565b6020808552840191601f190136833783525290565b632cd44ac360e21b60005260046000fd5b5060405161105781611050816005610cd1565b0382610d54565b90565b60ff811461107e5760ff811690601f821161102c5760408051926110178285610d54565b5060405161105781611050816006610cd1565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161110e579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa15610242576000516001600160a01b038116156111025790600090600090565b50600090600190600090565b50505060009160039190565b919091600481101561118c578061113057509050565b60006001820361114b5763f645eedf60e01b60005260046000fd5b5060028103611169578263fce698f760e01b60005260045260246000fd5b9091600360009214611179575050565b6335e2f38360e21b825260045260249150fd5b634e487b7160e01b600052602160045260246000fdfea264697066735822122059c736fe595882f8457f5096c981ff170aec2b912505d2157105727268b3927d64736f6c634300081b0033",
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
