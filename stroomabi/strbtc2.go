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

// StrBTC2MintInvoice is an auto generated low-level Go binding around an user-defined struct.
type StrBTC2MintInvoice struct {
	BtcDepositId [32]byte
	Recipient    common.Address
	Amount       *big.Int
}

// StrBtc2MetaData contains all meta data concerning the StrBtc2 contract.
var StrBtc2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BTC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONVERTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DUST_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_MINT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSAGE_UPDATE_TOTAL_SUPPLY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"btcDepositIds\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"converterBurn\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"converterMint\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"encodeInvoice\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC2.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"encodeTotalSupplyUpdate\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMintInvoiceHash\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC2.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPooledBTCByShares\",\"inputs\":[{\"name\":\"sharesAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShares\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSharesByPooledBTC\",\"inputs\":[{\"name\":\"btcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalSupplyUpdateHash\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_network\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"},{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"},{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRewardTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxRewardPercent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minTimeBetweenRewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"invoice\",\"type\":\"tuple\",\"internalType\":\"structstrBTC2.MintInvoice\",\"components\":[{\"name\":\"btcDepositId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintRewards\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"network\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBitcoinNetworkEncoder.Network\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"BTCAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMaxRewardPercent\",\"inputs\":[{\"name\":\"_maxRewardPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinTimeBetweenRewards\",\"inputs\":[{\"name\":\"_minTimeBetweenRewards\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWithdrawAmount\",\"inputs\":[{\"name\":\"_minWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidatorRegistry\",\"inputs\":[{\"name\":\"_validatorRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupplyUpdateNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractValidatorRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConverterBurn\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConverterMint\",\"inputs\":[{\"name\":\"converter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintBtcEvent\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_btcDepositId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedeemBtcEvent\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_BTCAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalSupplyUpdatedEvent\",\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalBTCSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_totalShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AmountBelowMinWithdraw\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotBurnFromZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotMintToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DeltaIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBTCAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSharesOrPooledBTC\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTotalSupplyNonce\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxRewardPercentTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinTimeBetweenRewardsTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinWithdrawTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAlreadyProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountTooBig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintToContractAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardTooBig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardTooFrequent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UpdateAlreadyProcessed\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50611de38061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610260575f3560e01c8063607375a11161014b578063b0390325116100bf578063e0a0461011610084578063e0a0461014610530578063e63ab1e914610543578063e737557f1461056a578063f04da65b146105a3578063f376ebbb146105cb578063f8077fae146105f5575f5ffd5b8063b0390325146104db578063c4d66de8146104e4578063c87b7a23146104f7578063dd62ed3e1461050a578063e07fbd001461051d575f5ffd5b806395d89b411161011057806395d89b41146104885780639cee9492146104905780639fd95687146104a3578063a7ce4565146104b6578063a9059cbb146104bf578063af7f131f146104d2575f5ffd5b8063607375a11461041c5780636739afca1461042f57806370a082311461044f57806380239a04146104625780638c659bf214610475575f5ffd5b80632f4cfd47116101e257806342193473116101a7578063421934731461039f578063457e1a49146103a857806349773050146103b157806355573c77146103c45780635abdb0dc146103d7578063603c6a67146103ea575f5ffd5b80632f4cfd4714610359578063313ce5671461036c5780633357325e1461037b5780633395b75a1461038e5780633a98ef3914610397575f5ffd5b806321172f3b1161022857806321172f3b1461030057806323b872dd1461031557806324b76fd51461032857806325e0a33f1461033b5780632792949d1461034e575f5ffd5b8063026034f01461026457806306fdde031461029b578063095ea7b3146102b057806313776a8d146102c357806318160ddd146102f8575b5f5ffd5b61028661027236600461179a565b60096020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6102a36105fe565b60405161029291906117df565b6102866102be366004611808565b6106be565b6102ea7f1cf336fddcc7dc48127faf7a5b80ee54fce73ef647eecd31c24bb6cce3ac3eef81565b604051908152602001610292565b6004546102ea565b61031361030e366004611877565b6106d7565b005b6102866103233660046118c6565b610934565b610313610336366004611904565b610957565b61031361034936600461179a565b610a94565b6102ea6305f5e10081565b610313610367366004611808565b610abc565b60405160088152602001610292565b610313610389366004611808565b610b32565b6102ea60065481565b6005546102ea565b6102ea61022281565b6102ea60015481565b6103136103bf36600461194c565b610ba0565b6102ea6103d236600461179a565b610be8565b6103136103e536600461179a565b610c31565b6102a3604051806040016040528060138152602001725354524f4f4d5f4d494e545f494e564f49434560681b81525081565b61031361042a366004611967565b610c59565b5f5461044290600160a01b900460ff1681565b60405161029291906119f7565b6102ea61045d36600461194c565b610df7565b6102ea610470366004611a05565b610e30565b6102a3610483366004611a05565b610ee3565b6102a3610f28565b61031361049e36600461179a565b610f66565b6103136104b1366004611a3b565b610f8d565b6102ea60025481565b6102866104cd366004611808565b61107e565b6102ea60035481565b6102ea60075481565b6103136104f236600461194c565b610bc7565b6102ea61050536600461179a565b61108b565b6102ea610518366004611a72565b6110ca565b6102ea61052b366004611aa9565b611113565b6102a361053e366004611aa9565b6111b7565b6102ea7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102a36040518060400160405280601a8152602001795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b81525081565b6102ea6105b136600461194c565b6001600160a01b03165f908152600a602052604090205490565b5f546105dd906001600160a01b031681565b6040516001600160a01b039091168152602001610292565b6102ea60085481565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060915f516020611d8e5f395f51905f529161063c90611ac3565b80601f016020809104026020016040519081016040528092919081815260200182805461066890611ac3565b80156106b35780601f1061068a576101008083540402835291602001916106b3565b820191905f5260205f20905b81548152906001019060200180831161069657829003601f168201915b505050505091505090565b5f336106cb81858561121a565b60019150505b92915050565b6040518060400160405280601a8152602001795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b8152506107148585610ee3565b5f546040516311ee58a960e01b8152859185916001600160a01b03909116906311ee58a99061074d908790879087908790600401611b1d565b602060405180830381865afa158015610768573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078c9190611b61565b6107a957604051636227817160e01b815260040160405180910390fd5b60035488146107cb576040516336b118d760e11b815260040160405180910390fd5b600160035f8282546107dd9190611b94565b90915550505f87900361080357604051630aa7fcbb60e21b815260040160405180910390fd5b6007546008546108139190611b94565b421015610832576040516283cb6b60e51b815260040160405180910390fd5b5f6127106006546004546108469190611ba7565b6108509190611bbe565b90508088111561087357604051630779a2a360e31b815260040160405180910390fd5b5f61087e8a8a610e30565b5f8181526009602052604090205490915060ff16156108b05760405163ce199d1760e01b815260040160405180910390fd5b5f818152600960205260408120805460ff19166001179055600480548b92906108da908490611b94565b909155505042600855600454600554604080518d815260208101939093528201527f339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b934409060600160405180910390a150505050505050505050565b5f3361094185828561122c565b61094c858585611294565b506001949350505050565b60015483101561097a57604051633813eacd60e21b815260040160405180910390fd5b5f54600160a01b900460ff166003811115610997576109976119c3565b604051632ce1e0f560e01b815273__$113774c24a411e022512c75ffd4c4add5e$__91632ce1e0f5916109d1919086908690600401611bdd565b602060405180830381865af41580156109ec573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a109190611b61565b610a2d576040516312d58f2760e21b815260040160405180910390fd5b610a3733846112f1565b600160025f828254610a499190611b94565b909155505060025460405133917f83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e91610a8791869186918991611c06565b60405180910390a2505050565b61a8c0811015610ab75760405163a7bbe41d60e01b815260040160405180910390fd5b600755565b6001600160a01b038216610ae3576040516389a4ea1960e01b815260040160405180910390fd5b610aed8282611329565b6040518181526001600160a01b0383169033907f76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e906020015b60405180910390a35050565b6001600160a01b038216610b595760405163a30d2d8760e01b815260040160405180910390fd5b610b6382826112f1565b6040518181526001600160a01b0383169033907f25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad56790602001610b26565b6001600160a01b038116610bc75760405163e6c4247b60e01b815260040160405180910390fd5b5f80546001600160a01b0319166001600160a01b0392909216919091179055565b5f6005545f1480610bf95750600454155b15610c17576040516340a3daff60e11b815260040160405180910390fd5b600554600454610c279084611ba7565b6106d19190611bbe565b610222811015610c54576040516343a87e9960e11b815260040160405180910390fd5b600155565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f81158015610c9e5750825b90505f8267ffffffffffffffff166001148015610cba5750303b155b905081158015610cc8575080155b15610ce65760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610d1057845460ff60401b1916600160401b1785555b610d5e6040518060400160405280600e81526020016d29ba3937b7b6902134ba31b7b4b760911b8152506040518060400160405280600681526020016573747242544360d01b81525061135d565b610d6788610bc7565b611b586001555f80548a919060ff60a01b1916600160a01b836003811115610d9157610d916119c3565b02179055506064600655620151806007555f6008558315610dec57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f6005545f03610e0857505f919050565b6005546004546001600160a01b0384165f908152600a6020526040902054610c279190611ba7565b5f805460408051808201909152601a8152795354524f4f4d5f5550444154455f544f54414c5f535550504c5960301b60208201526001600160a01b039091169063bca0ac0690610e808686610ee3565b6040518363ffffffff1660e01b8152600401610e9d929190611c2c565b602060405180830381865afa158015610eb8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610edc9190611c50565b9392505050565b604080516020810184905290810182905230606090811b6bffffffffffffffffffffffff19168183015290607401604051602081830303815290604052905092915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060915f516020611d8e5f395f51905f529161063c90611ac3565b6064811115610f8857604051630737692960e31b815260040160405180910390fd5b600655565b604051806040016040528060138152602001725354524f4f4d5f4d494e545f494e564f49434560681b815250610fc2846111b7565b5f546040516311ee58a960e01b8152859185916001600160a01b03909116906311ee58a990610ffb908790879087908790600401611b1d565b602060405180830381865afa158015611016573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061103a9190611b61565b61105757604051636227817160e01b815260040160405180910390fd5b6110756040880180359061106e9060208b0161194c565b893561136f565b50505050505050565b5f336106cb818585611294565b5f6005545f148061109c5750600454155b156110ba576040516340a3daff60e11b815260040160405180910390fd5b600454600554610c279084611ba7565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b5f80546040805180820190915260138152725354524f4f4d5f4d494e545f494e564f49434560681b60208201526001600160a01b039091169063bca0ac069061115b856111b7565b6040518363ffffffff1660e01b8152600401611178929190611c2c565b602060405180830381865afa158015611193573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106d19190611c50565b60606111c9604083016020840161194c565b604080516bffffffffffffffffffffffff19606093841b81166020830152918501356034820152843560548201523090921b1660748201526088016040516020818303038152906040529050919050565b611227838383600161147b565b505050565b5f61123784846110ca565b90505f19811461128e578181101561128057604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b61128e84848484035f61147b565b50505050565b6001600160a01b0383166112bd57604051634b637e8f60e11b81525f6004820152602401611277565b6001600160a01b0382166112e65760405163ec442f0560e01b81525f6004820152602401611277565b61122783838361155f565b6001600160a01b03821661131a57604051634b637e8f60e11b81525f6004820152602401611277565b611325825f8361155f565b5050565b6001600160a01b0382166113525760405163ec442f0560e01b81525f6004820152602401611277565b6113255f838361155f565b6113656116ff565b611325828261174a565b825f0361138f57604051631cebf66f60e11b815260040160405180910390fd5b6113a16305f5e1006301406f40611ba7565b83106113c057604051636006659960e11b815260040160405180910390fd5b306001600160a01b038316036113e95760405163603acaa960e11b815260040160405180910390fd5b5f8181526009602052604090205460ff161561141857604051637bdb87d160e01b815260040160405180910390fd5b5f818152600960205260409020805460ff1916600117905561143a8284611329565b60408051848152602081018390526001600160a01b038416917fb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa4699101610a87565b5f516020611d8e5f395f51905f526001600160a01b0385166114b25760405163e602df0560e01b81525f6004820152602401611277565b6001600160a01b0384166114db57604051634a1406b160e11b81525f6004820152602401611277565b6001600160a01b038086165f9081526001830160209081526040808320938816835292905220839055811561155857836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161154f91815260200190565b60405180910390a35b5050505050565b6001600160a01b0383166115e5575f6005545f148061157e5750600454155b6115905761158b8261108b565b611592565b815b90508160045f8282546115a59190611b94565b925050819055508060055f8282546115bd9190611b94565b90915550506001600160a01b0383165f908152600a60205260409020805490910190556116ad565b5f6115ef84610df7565b90508181101561161257604051631e9acf1760e31b815260040160405180910390fd5b5f61161c8361108b565b90506001600160a01b03841661167d578260045f82825461163d9190611c67565b925050819055508060055f8282546116559190611c67565b90915550506001600160a01b0385165f908152600a60205260409020805482900390556116aa565b6001600160a01b038086165f908152600a6020526040808220805485900390559186168152208054820190555b50505b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516116f291815260200190565b60405180910390a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661174857604051631afcd79f60e31b815260040160405180910390fd5b565b6117526116ff565b5f516020611d8e5f395f51905f527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0361178b8482611cd2565b506004810161128e8382611cd2565b5f602082840312156117aa575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610edc60208301846117b1565b6001600160a01b0381168114611805575f5ffd5b50565b5f5f60408385031215611819575f5ffd5b8235611824816117f1565b946020939093013593505050565b5f5f83601f840112611842575f5ffd5b50813567ffffffffffffffff811115611859575f5ffd5b602083019150836020828501011115611870575f5ffd5b9250929050565b5f5f5f5f6060858703121561188a575f5ffd5b8435935060208501359250604085013567ffffffffffffffff8111156118ae575f5ffd5b6118ba87828801611832565b95989497509550505050565b5f5f5f606084860312156118d8575f5ffd5b83356118e3816117f1565b925060208401356118f3816117f1565b929592945050506040919091013590565b5f5f5f60408486031215611916575f5ffd5b83359250602084013567ffffffffffffffff811115611933575f5ffd5b61193f86828701611832565b9497909650939450505050565b5f6020828403121561195c575f5ffd5b8135610edc816117f1565b5f5f5f5f6080858703121561197a575f5ffd5b843560048110611988575f5ffd5b93506020850135611998816117f1565b925060408501356119a8816117f1565b915060608501356119b8816117f1565b939692955090935050565b634e487b7160e01b5f52602160045260245ffd5b600481106119f357634e487b7160e01b5f52602160045260245ffd5b9052565b602081016106d182846119d7565b5f5f60408385031215611a16575f5ffd5b50508035926020909101359150565b5f60608284031215611a35575f5ffd5b50919050565b5f5f5f60808486031215611a4d575f5ffd5b611a578585611a25565b9250606084013567ffffffffffffffff811115611933575f5ffd5b5f5f60408385031215611a83575f5ffd5b8235611a8e816117f1565b91506020830135611a9e816117f1565b809150509250929050565b5f60608284031215611ab9575f5ffd5b610edc8383611a25565b600181811c90821680611ad757607f821691505b602082108103611a3557634e487b7160e01b5f52602260045260245ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606081525f611b2f60608301876117b1565b8281036020840152611b4181876117b1565b90508281036040840152611b56818587611af5565b979650505050505050565b5f60208284031215611b71575f5ffd5b81518015158114610edc575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b808201808211156106d1576106d1611b80565b80820281158282048414176106d1576106d1611b80565b5f82611bd857634e487b7160e01b5f52601260045260245ffd5b500490565b611be781856119d7565b604060208201525f611bfd604083018486611af5565b95945050505050565b606081525f611c19606083018688611af5565b6020830194909452506040015292915050565b604081525f611c3e60408301856117b1565b8281036020840152611bfd81856117b1565b5f60208284031215611c60575f5ffd5b5051919050565b818103818111156106d1576106d1611b80565b634e487b7160e01b5f52604160045260245ffd5b601f82111561122757805f5260205f20601f840160051c81016020851015611cb35750805b601f840160051c820191505b81811015611558575f8155600101611cbf565b815167ffffffffffffffff811115611cec57611cec611c7a565b611d0081611cfa8454611ac3565b84611c8e565b6020601f821160018114611d32575f8315611d1b5750848201515b5f19600385901b1c1916600184901b178455611558565b5f84815260208120601f198516915b82811015611d615787850151825560209485019460019092019101611d41565b5084821015611d7e57868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00a26469706673582212207579764aa94f72ff1bbe31f5c050d28bad7d042552824ab217a7bda252bd076c64736f6c634300081b0033",
}

// StrBtc2ABI is the input ABI used to generate the binding from.
// Deprecated: Use StrBtc2MetaData.ABI instead.
var StrBtc2ABI = StrBtc2MetaData.ABI

// StrBtc2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrBtc2MetaData.Bin instead.
var StrBtc2Bin = StrBtc2MetaData.Bin

// DeployStrBtc2 deploys a new Ethereum contract, binding an instance of StrBtc2 to it.
func DeployStrBtc2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StrBtc2, error) {
	parsed, err := StrBtc2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrBtc2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrBtc2{StrBtc2Caller: StrBtc2Caller{contract: contract}, StrBtc2Transactor: StrBtc2Transactor{contract: contract}, StrBtc2Filterer: StrBtc2Filterer{contract: contract}}, nil
}

// StrBtc2 is an auto generated Go binding around an Ethereum contract.
type StrBtc2 struct {
	StrBtc2Caller     // Read-only binding to the contract
	StrBtc2Transactor // Write-only binding to the contract
	StrBtc2Filterer   // Log filterer for contract events
}

// StrBtc2Caller is an auto generated read-only Go binding around an Ethereum contract.
type StrBtc2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrBtc2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type StrBtc2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrBtc2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrBtc2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrBtc2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrBtc2Session struct {
	Contract     *StrBtc2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrBtc2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrBtc2CallerSession struct {
	Contract *StrBtc2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StrBtc2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrBtc2TransactorSession struct {
	Contract     *StrBtc2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StrBtc2Raw is an auto generated low-level Go binding around an Ethereum contract.
type StrBtc2Raw struct {
	Contract *StrBtc2 // Generic contract binding to access the raw methods on
}

// StrBtc2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrBtc2CallerRaw struct {
	Contract *StrBtc2Caller // Generic read-only contract binding to access the raw methods on
}

// StrBtc2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrBtc2TransactorRaw struct {
	Contract *StrBtc2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStrBtc2 creates a new instance of StrBtc2, bound to a specific deployed contract.
func NewStrBtc2(address common.Address, backend bind.ContractBackend) (*StrBtc2, error) {
	contract, err := bindStrBtc2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrBtc2{StrBtc2Caller: StrBtc2Caller{contract: contract}, StrBtc2Transactor: StrBtc2Transactor{contract: contract}, StrBtc2Filterer: StrBtc2Filterer{contract: contract}}, nil
}

// NewStrBtc2Caller creates a new read-only instance of StrBtc2, bound to a specific deployed contract.
func NewStrBtc2Caller(address common.Address, caller bind.ContractCaller) (*StrBtc2Caller, error) {
	contract, err := bindStrBtc2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrBtc2Caller{contract: contract}, nil
}

// NewStrBtc2Transactor creates a new write-only instance of StrBtc2, bound to a specific deployed contract.
func NewStrBtc2Transactor(address common.Address, transactor bind.ContractTransactor) (*StrBtc2Transactor, error) {
	contract, err := bindStrBtc2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrBtc2Transactor{contract: contract}, nil
}

// NewStrBtc2Filterer creates a new log filterer instance of StrBtc2, bound to a specific deployed contract.
func NewStrBtc2Filterer(address common.Address, filterer bind.ContractFilterer) (*StrBtc2Filterer, error) {
	contract, err := bindStrBtc2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrBtc2Filterer{contract: contract}, nil
}

// bindStrBtc2 binds a generic wrapper to an already deployed contract.
func bindStrBtc2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrBtc2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrBtc2 *StrBtc2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrBtc2.Contract.StrBtc2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrBtc2 *StrBtc2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrBtc2.Contract.StrBtc2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrBtc2 *StrBtc2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrBtc2.Contract.StrBtc2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrBtc2 *StrBtc2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrBtc2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrBtc2 *StrBtc2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrBtc2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrBtc2 *StrBtc2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrBtc2.Contract.contract.Transact(opts, method, params...)
}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) BTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "BTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) BTC() (*big.Int, error) {
	return _StrBtc2.Contract.BTC(&_StrBtc2.CallOpts)
}

// BTC is a free data retrieval call binding the contract method 0x2792949d.
//
// Solidity: function BTC() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) BTC() (*big.Int, error) {
	return _StrBtc2.Contract.BTC(&_StrBtc2.CallOpts)
}

// CONVERTERROLE is a free data retrieval call binding the contract method 0x13776a8d.
//
// Solidity: function CONVERTER_ROLE() view returns(bytes32)
func (_StrBtc2 *StrBtc2Caller) CONVERTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "CONVERTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONVERTERROLE is a free data retrieval call binding the contract method 0x13776a8d.
//
// Solidity: function CONVERTER_ROLE() view returns(bytes32)
func (_StrBtc2 *StrBtc2Session) CONVERTERROLE() ([32]byte, error) {
	return _StrBtc2.Contract.CONVERTERROLE(&_StrBtc2.CallOpts)
}

// CONVERTERROLE is a free data retrieval call binding the contract method 0x13776a8d.
//
// Solidity: function CONVERTER_ROLE() view returns(bytes32)
func (_StrBtc2 *StrBtc2CallerSession) CONVERTERROLE() ([32]byte, error) {
	return _StrBtc2.Contract.CONVERTERROLE(&_StrBtc2.CallOpts)
}

// DUSTLIMIT is a free data retrieval call binding the contract method 0x42193473.
//
// Solidity: function DUST_LIMIT() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) DUSTLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "DUST_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DUSTLIMIT is a free data retrieval call binding the contract method 0x42193473.
//
// Solidity: function DUST_LIMIT() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) DUSTLIMIT() (*big.Int, error) {
	return _StrBtc2.Contract.DUSTLIMIT(&_StrBtc2.CallOpts)
}

// DUSTLIMIT is a free data retrieval call binding the contract method 0x42193473.
//
// Solidity: function DUST_LIMIT() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) DUSTLIMIT() (*big.Int, error) {
	return _StrBtc2.Contract.DUSTLIMIT(&_StrBtc2.CallOpts)
}

// MESSAGEMINT is a free data retrieval call binding the contract method 0x603c6a67.
//
// Solidity: function MESSAGE_MINT() view returns(bytes)
func (_StrBtc2 *StrBtc2Caller) MESSAGEMINT(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "MESSAGE_MINT")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MESSAGEMINT is a free data retrieval call binding the contract method 0x603c6a67.
//
// Solidity: function MESSAGE_MINT() view returns(bytes)
func (_StrBtc2 *StrBtc2Session) MESSAGEMINT() ([]byte, error) {
	return _StrBtc2.Contract.MESSAGEMINT(&_StrBtc2.CallOpts)
}

// MESSAGEMINT is a free data retrieval call binding the contract method 0x603c6a67.
//
// Solidity: function MESSAGE_MINT() view returns(bytes)
func (_StrBtc2 *StrBtc2CallerSession) MESSAGEMINT() ([]byte, error) {
	return _StrBtc2.Contract.MESSAGEMINT(&_StrBtc2.CallOpts)
}

// MESSAGEUPDATETOTALSUPPLY is a free data retrieval call binding the contract method 0xe737557f.
//
// Solidity: function MESSAGE_UPDATE_TOTAL_SUPPLY() view returns(bytes)
func (_StrBtc2 *StrBtc2Caller) MESSAGEUPDATETOTALSUPPLY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "MESSAGE_UPDATE_TOTAL_SUPPLY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MESSAGEUPDATETOTALSUPPLY is a free data retrieval call binding the contract method 0xe737557f.
//
// Solidity: function MESSAGE_UPDATE_TOTAL_SUPPLY() view returns(bytes)
func (_StrBtc2 *StrBtc2Session) MESSAGEUPDATETOTALSUPPLY() ([]byte, error) {
	return _StrBtc2.Contract.MESSAGEUPDATETOTALSUPPLY(&_StrBtc2.CallOpts)
}

// MESSAGEUPDATETOTALSUPPLY is a free data retrieval call binding the contract method 0xe737557f.
//
// Solidity: function MESSAGE_UPDATE_TOTAL_SUPPLY() view returns(bytes)
func (_StrBtc2 *StrBtc2CallerSession) MESSAGEUPDATETOTALSUPPLY() ([]byte, error) {
	return _StrBtc2.Contract.MESSAGEUPDATETOTALSUPPLY(&_StrBtc2.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StrBtc2 *StrBtc2Caller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StrBtc2 *StrBtc2Session) PAUSERROLE() ([32]byte, error) {
	return _StrBtc2.Contract.PAUSERROLE(&_StrBtc2.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StrBtc2 *StrBtc2CallerSession) PAUSERROLE() ([32]byte, error) {
	return _StrBtc2.Contract.PAUSERROLE(&_StrBtc2.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StrBtc2 *StrBtc2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StrBtc2.Contract.Allowance(&_StrBtc2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StrBtc2.Contract.Allowance(&_StrBtc2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StrBtc2 *StrBtc2Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _StrBtc2.Contract.BalanceOf(&_StrBtc2.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StrBtc2.Contract.BalanceOf(&_StrBtc2.CallOpts, account)
}

// BtcDepositIds is a free data retrieval call binding the contract method 0x026034f0.
//
// Solidity: function btcDepositIds(bytes32 ) view returns(bool)
func (_StrBtc2 *StrBtc2Caller) BtcDepositIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "btcDepositIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BtcDepositIds is a free data retrieval call binding the contract method 0x026034f0.
//
// Solidity: function btcDepositIds(bytes32 ) view returns(bool)
func (_StrBtc2 *StrBtc2Session) BtcDepositIds(arg0 [32]byte) (bool, error) {
	return _StrBtc2.Contract.BtcDepositIds(&_StrBtc2.CallOpts, arg0)
}

// BtcDepositIds is a free data retrieval call binding the contract method 0x026034f0.
//
// Solidity: function btcDepositIds(bytes32 ) view returns(bool)
func (_StrBtc2 *StrBtc2CallerSession) BtcDepositIds(arg0 [32]byte) (bool, error) {
	return _StrBtc2.Contract.BtcDepositIds(&_StrBtc2.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StrBtc2 *StrBtc2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StrBtc2 *StrBtc2Session) Decimals() (uint8, error) {
	return _StrBtc2.Contract.Decimals(&_StrBtc2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StrBtc2 *StrBtc2CallerSession) Decimals() (uint8, error) {
	return _StrBtc2.Contract.Decimals(&_StrBtc2.CallOpts)
}

// EncodeInvoice is a free data retrieval call binding the contract method 0xe0a04610.
//
// Solidity: function encodeInvoice((bytes32,address,uint256) invoice) view returns(bytes)
func (_StrBtc2 *StrBtc2Caller) EncodeInvoice(opts *bind.CallOpts, invoice StrBTC2MintInvoice) ([]byte, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "encodeInvoice", invoice)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeInvoice is a free data retrieval call binding the contract method 0xe0a04610.
//
// Solidity: function encodeInvoice((bytes32,address,uint256) invoice) view returns(bytes)
func (_StrBtc2 *StrBtc2Session) EncodeInvoice(invoice StrBTC2MintInvoice) ([]byte, error) {
	return _StrBtc2.Contract.EncodeInvoice(&_StrBtc2.CallOpts, invoice)
}

// EncodeInvoice is a free data retrieval call binding the contract method 0xe0a04610.
//
// Solidity: function encodeInvoice((bytes32,address,uint256) invoice) view returns(bytes)
func (_StrBtc2 *StrBtc2CallerSession) EncodeInvoice(invoice StrBTC2MintInvoice) ([]byte, error) {
	return _StrBtc2.Contract.EncodeInvoice(&_StrBtc2.CallOpts, invoice)
}

// EncodeTotalSupplyUpdate is a free data retrieval call binding the contract method 0x8c659bf2.
//
// Solidity: function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) view returns(bytes)
func (_StrBtc2 *StrBtc2Caller) EncodeTotalSupplyUpdate(opts *bind.CallOpts, nonce *big.Int, delta *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "encodeTotalSupplyUpdate", nonce, delta)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTotalSupplyUpdate is a free data retrieval call binding the contract method 0x8c659bf2.
//
// Solidity: function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) view returns(bytes)
func (_StrBtc2 *StrBtc2Session) EncodeTotalSupplyUpdate(nonce *big.Int, delta *big.Int) ([]byte, error) {
	return _StrBtc2.Contract.EncodeTotalSupplyUpdate(&_StrBtc2.CallOpts, nonce, delta)
}

// EncodeTotalSupplyUpdate is a free data retrieval call binding the contract method 0x8c659bf2.
//
// Solidity: function encodeTotalSupplyUpdate(uint256 nonce, uint256 delta) view returns(bytes)
func (_StrBtc2 *StrBtc2CallerSession) EncodeTotalSupplyUpdate(nonce *big.Int, delta *big.Int) ([]byte, error) {
	return _StrBtc2.Contract.EncodeTotalSupplyUpdate(&_StrBtc2.CallOpts, nonce, delta)
}

// GetMintInvoiceHash is a free data retrieval call binding the contract method 0xe07fbd00.
//
// Solidity: function getMintInvoiceHash((bytes32,address,uint256) invoice) view returns(bytes32)
func (_StrBtc2 *StrBtc2Caller) GetMintInvoiceHash(opts *bind.CallOpts, invoice StrBTC2MintInvoice) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "getMintInvoiceHash", invoice)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMintInvoiceHash is a free data retrieval call binding the contract method 0xe07fbd00.
//
// Solidity: function getMintInvoiceHash((bytes32,address,uint256) invoice) view returns(bytes32)
func (_StrBtc2 *StrBtc2Session) GetMintInvoiceHash(invoice StrBTC2MintInvoice) ([32]byte, error) {
	return _StrBtc2.Contract.GetMintInvoiceHash(&_StrBtc2.CallOpts, invoice)
}

// GetMintInvoiceHash is a free data retrieval call binding the contract method 0xe07fbd00.
//
// Solidity: function getMintInvoiceHash((bytes32,address,uint256) invoice) view returns(bytes32)
func (_StrBtc2 *StrBtc2CallerSession) GetMintInvoiceHash(invoice StrBTC2MintInvoice) ([32]byte, error) {
	return _StrBtc2.Contract.GetMintInvoiceHash(&_StrBtc2.CallOpts, invoice)
}

// GetPooledBTCByShares is a free data retrieval call binding the contract method 0x55573c77.
//
// Solidity: function getPooledBTCByShares(uint256 sharesAmount) view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) GetPooledBTCByShares(opts *bind.CallOpts, sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "getPooledBTCByShares", sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledBTCByShares is a free data retrieval call binding the contract method 0x55573c77.
//
// Solidity: function getPooledBTCByShares(uint256 sharesAmount) view returns(uint256)
func (_StrBtc2 *StrBtc2Session) GetPooledBTCByShares(sharesAmount *big.Int) (*big.Int, error) {
	return _StrBtc2.Contract.GetPooledBTCByShares(&_StrBtc2.CallOpts, sharesAmount)
}

// GetPooledBTCByShares is a free data retrieval call binding the contract method 0x55573c77.
//
// Solidity: function getPooledBTCByShares(uint256 sharesAmount) view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) GetPooledBTCByShares(sharesAmount *big.Int) (*big.Int, error) {
	return _StrBtc2.Contract.GetPooledBTCByShares(&_StrBtc2.CallOpts, sharesAmount)
}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address account) view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) GetShares(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "getShares", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address account) view returns(uint256)
func (_StrBtc2 *StrBtc2Session) GetShares(account common.Address) (*big.Int, error) {
	return _StrBtc2.Contract.GetShares(&_StrBtc2.CallOpts, account)
}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address account) view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) GetShares(account common.Address) (*big.Int, error) {
	return _StrBtc2.Contract.GetShares(&_StrBtc2.CallOpts, account)
}

// GetSharesByPooledBTC is a free data retrieval call binding the contract method 0xc87b7a23.
//
// Solidity: function getSharesByPooledBTC(uint256 btcAmount) view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) GetSharesByPooledBTC(opts *bind.CallOpts, btcAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "getSharesByPooledBTC", btcAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledBTC is a free data retrieval call binding the contract method 0xc87b7a23.
//
// Solidity: function getSharesByPooledBTC(uint256 btcAmount) view returns(uint256)
func (_StrBtc2 *StrBtc2Session) GetSharesByPooledBTC(btcAmount *big.Int) (*big.Int, error) {
	return _StrBtc2.Contract.GetSharesByPooledBTC(&_StrBtc2.CallOpts, btcAmount)
}

// GetSharesByPooledBTC is a free data retrieval call binding the contract method 0xc87b7a23.
//
// Solidity: function getSharesByPooledBTC(uint256 btcAmount) view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) GetSharesByPooledBTC(btcAmount *big.Int) (*big.Int, error) {
	return _StrBtc2.Contract.GetSharesByPooledBTC(&_StrBtc2.CallOpts, btcAmount)
}

// GetTotalSupplyUpdateHash is a free data retrieval call binding the contract method 0x80239a04.
//
// Solidity: function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) view returns(bytes32)
func (_StrBtc2 *StrBtc2Caller) GetTotalSupplyUpdateHash(opts *bind.CallOpts, nonce *big.Int, delta *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "getTotalSupplyUpdateHash", nonce, delta)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTotalSupplyUpdateHash is a free data retrieval call binding the contract method 0x80239a04.
//
// Solidity: function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) view returns(bytes32)
func (_StrBtc2 *StrBtc2Session) GetTotalSupplyUpdateHash(nonce *big.Int, delta *big.Int) ([32]byte, error) {
	return _StrBtc2.Contract.GetTotalSupplyUpdateHash(&_StrBtc2.CallOpts, nonce, delta)
}

// GetTotalSupplyUpdateHash is a free data retrieval call binding the contract method 0x80239a04.
//
// Solidity: function getTotalSupplyUpdateHash(uint256 nonce, uint256 delta) view returns(bytes32)
func (_StrBtc2 *StrBtc2CallerSession) GetTotalSupplyUpdateHash(nonce *big.Int, delta *big.Int) ([32]byte, error) {
	return _StrBtc2.Contract.GetTotalSupplyUpdateHash(&_StrBtc2.CallOpts, nonce, delta)
}

// LastRewardTimestamp is a free data retrieval call binding the contract method 0xf8077fae.
//
// Solidity: function lastRewardTimestamp() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) LastRewardTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "lastRewardTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardTimestamp is a free data retrieval call binding the contract method 0xf8077fae.
//
// Solidity: function lastRewardTimestamp() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) LastRewardTimestamp() (*big.Int, error) {
	return _StrBtc2.Contract.LastRewardTimestamp(&_StrBtc2.CallOpts)
}

// LastRewardTimestamp is a free data retrieval call binding the contract method 0xf8077fae.
//
// Solidity: function lastRewardTimestamp() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) LastRewardTimestamp() (*big.Int, error) {
	return _StrBtc2.Contract.LastRewardTimestamp(&_StrBtc2.CallOpts)
}

// MaxRewardPercent is a free data retrieval call binding the contract method 0x3395b75a.
//
// Solidity: function maxRewardPercent() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) MaxRewardPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "maxRewardPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRewardPercent is a free data retrieval call binding the contract method 0x3395b75a.
//
// Solidity: function maxRewardPercent() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) MaxRewardPercent() (*big.Int, error) {
	return _StrBtc2.Contract.MaxRewardPercent(&_StrBtc2.CallOpts)
}

// MaxRewardPercent is a free data retrieval call binding the contract method 0x3395b75a.
//
// Solidity: function maxRewardPercent() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) MaxRewardPercent() (*big.Int, error) {
	return _StrBtc2.Contract.MaxRewardPercent(&_StrBtc2.CallOpts)
}

// MinTimeBetweenRewards is a free data retrieval call binding the contract method 0xb0390325.
//
// Solidity: function minTimeBetweenRewards() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) MinTimeBetweenRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "minTimeBetweenRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTimeBetweenRewards is a free data retrieval call binding the contract method 0xb0390325.
//
// Solidity: function minTimeBetweenRewards() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) MinTimeBetweenRewards() (*big.Int, error) {
	return _StrBtc2.Contract.MinTimeBetweenRewards(&_StrBtc2.CallOpts)
}

// MinTimeBetweenRewards is a free data retrieval call binding the contract method 0xb0390325.
//
// Solidity: function minTimeBetweenRewards() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) MinTimeBetweenRewards() (*big.Int, error) {
	return _StrBtc2.Contract.MinTimeBetweenRewards(&_StrBtc2.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) MinWithdrawAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "minWithdrawAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) MinWithdrawAmount() (*big.Int, error) {
	return _StrBtc2.Contract.MinWithdrawAmount(&_StrBtc2.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) MinWithdrawAmount() (*big.Int, error) {
	return _StrBtc2.Contract.MinWithdrawAmount(&_StrBtc2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrBtc2 *StrBtc2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrBtc2 *StrBtc2Session) Name() (string, error) {
	return _StrBtc2.Contract.Name(&_StrBtc2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StrBtc2 *StrBtc2CallerSession) Name() (string, error) {
	return _StrBtc2.Contract.Name(&_StrBtc2.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint8)
func (_StrBtc2 *StrBtc2Caller) Network(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "network")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint8)
func (_StrBtc2 *StrBtc2Session) Network() (uint8, error) {
	return _StrBtc2.Contract.Network(&_StrBtc2.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(uint8)
func (_StrBtc2 *StrBtc2CallerSession) Network() (uint8, error) {
	return _StrBtc2.Contract.Network(&_StrBtc2.CallOpts)
}

// RedeemCounter is a free data retrieval call binding the contract method 0xa7ce4565.
//
// Solidity: function redeemCounter() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) RedeemCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "redeemCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemCounter is a free data retrieval call binding the contract method 0xa7ce4565.
//
// Solidity: function redeemCounter() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) RedeemCounter() (*big.Int, error) {
	return _StrBtc2.Contract.RedeemCounter(&_StrBtc2.CallOpts)
}

// RedeemCounter is a free data retrieval call binding the contract method 0xa7ce4565.
//
// Solidity: function redeemCounter() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) RedeemCounter() (*big.Int, error) {
	return _StrBtc2.Contract.RedeemCounter(&_StrBtc2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StrBtc2 *StrBtc2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StrBtc2 *StrBtc2Session) Symbol() (string, error) {
	return _StrBtc2.Contract.Symbol(&_StrBtc2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StrBtc2 *StrBtc2CallerSession) Symbol() (string, error) {
	return _StrBtc2.Contract.Symbol(&_StrBtc2.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) TotalShares() (*big.Int, error) {
	return _StrBtc2.Contract.TotalShares(&_StrBtc2.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) TotalShares() (*big.Int, error) {
	return _StrBtc2.Contract.TotalShares(&_StrBtc2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) TotalSupply() (*big.Int, error) {
	return _StrBtc2.Contract.TotalSupply(&_StrBtc2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) TotalSupply() (*big.Int, error) {
	return _StrBtc2.Contract.TotalSupply(&_StrBtc2.CallOpts)
}

// TotalSupplyUpdateNonce is a free data retrieval call binding the contract method 0xaf7f131f.
//
// Solidity: function totalSupplyUpdateNonce() view returns(uint256)
func (_StrBtc2 *StrBtc2Caller) TotalSupplyUpdateNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "totalSupplyUpdateNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyUpdateNonce is a free data retrieval call binding the contract method 0xaf7f131f.
//
// Solidity: function totalSupplyUpdateNonce() view returns(uint256)
func (_StrBtc2 *StrBtc2Session) TotalSupplyUpdateNonce() (*big.Int, error) {
	return _StrBtc2.Contract.TotalSupplyUpdateNonce(&_StrBtc2.CallOpts)
}

// TotalSupplyUpdateNonce is a free data retrieval call binding the contract method 0xaf7f131f.
//
// Solidity: function totalSupplyUpdateNonce() view returns(uint256)
func (_StrBtc2 *StrBtc2CallerSession) TotalSupplyUpdateNonce() (*big.Int, error) {
	return _StrBtc2.Contract.TotalSupplyUpdateNonce(&_StrBtc2.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_StrBtc2 *StrBtc2Caller) ValidatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrBtc2.contract.Call(opts, &out, "validatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_StrBtc2 *StrBtc2Session) ValidatorRegistry() (common.Address, error) {
	return _StrBtc2.Contract.ValidatorRegistry(&_StrBtc2.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_StrBtc2 *StrBtc2CallerSession) ValidatorRegistry() (common.Address, error) {
	return _StrBtc2.Contract.ValidatorRegistry(&_StrBtc2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StrBtc2 *StrBtc2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StrBtc2 *StrBtc2Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.Approve(&_StrBtc2.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StrBtc2 *StrBtc2TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.Approve(&_StrBtc2.TransactOpts, spender, value)
}

// ConverterBurn is a paid mutator transaction binding the contract method 0x3357325e.
//
// Solidity: function converterBurn(address from, uint256 amount) returns()
func (_StrBtc2 *StrBtc2Transactor) ConverterBurn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "converterBurn", from, amount)
}

// ConverterBurn is a paid mutator transaction binding the contract method 0x3357325e.
//
// Solidity: function converterBurn(address from, uint256 amount) returns()
func (_StrBtc2 *StrBtc2Session) ConverterBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.ConverterBurn(&_StrBtc2.TransactOpts, from, amount)
}

// ConverterBurn is a paid mutator transaction binding the contract method 0x3357325e.
//
// Solidity: function converterBurn(address from, uint256 amount) returns()
func (_StrBtc2 *StrBtc2TransactorSession) ConverterBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.ConverterBurn(&_StrBtc2.TransactOpts, from, amount)
}

// ConverterMint is a paid mutator transaction binding the contract method 0x2f4cfd47.
//
// Solidity: function converterMint(address recipient, uint256 amount) returns()
func (_StrBtc2 *StrBtc2Transactor) ConverterMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "converterMint", recipient, amount)
}

// ConverterMint is a paid mutator transaction binding the contract method 0x2f4cfd47.
//
// Solidity: function converterMint(address recipient, uint256 amount) returns()
func (_StrBtc2 *StrBtc2Session) ConverterMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.ConverterMint(&_StrBtc2.TransactOpts, recipient, amount)
}

// ConverterMint is a paid mutator transaction binding the contract method 0x2f4cfd47.
//
// Solidity: function converterMint(address recipient, uint256 amount) returns()
func (_StrBtc2 *StrBtc2TransactorSession) ConverterMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.ConverterMint(&_StrBtc2.TransactOpts, recipient, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x607375a1.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry, address _admin, address _pauser) returns()
func (_StrBtc2 *StrBtc2Transactor) Initialize(opts *bind.TransactOpts, _network uint8, _validatorRegistry common.Address, _admin common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "initialize", _network, _validatorRegistry, _admin, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0x607375a1.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry, address _admin, address _pauser) returns()
func (_StrBtc2 *StrBtc2Session) Initialize(_network uint8, _validatorRegistry common.Address, _admin common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrBtc2.Contract.Initialize(&_StrBtc2.TransactOpts, _network, _validatorRegistry, _admin, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0x607375a1.
//
// Solidity: function initialize(uint8 _network, address _validatorRegistry, address _admin, address _pauser) returns()
func (_StrBtc2 *StrBtc2TransactorSession) Initialize(_network uint8, _validatorRegistry common.Address, _admin common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrBtc2.Contract.Initialize(&_StrBtc2.TransactOpts, _network, _validatorRegistry, _admin, _pauser)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorRegistry) returns()
func (_StrBtc2 *StrBtc2Transactor) Initialize0(opts *bind.TransactOpts, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "initialize0", _validatorRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorRegistry) returns()
func (_StrBtc2 *StrBtc2Session) Initialize0(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc2.Contract.Initialize0(&_StrBtc2.TransactOpts, _validatorRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorRegistry) returns()
func (_StrBtc2 *StrBtc2TransactorSession) Initialize0(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc2.Contract.Initialize0(&_StrBtc2.TransactOpts, _validatorRegistry)
}

// Mint is a paid mutator transaction binding the contract method 0x9fd95687.
//
// Solidity: function mint((bytes32,address,uint256) invoice, bytes signature) returns()
func (_StrBtc2 *StrBtc2Transactor) Mint(opts *bind.TransactOpts, invoice StrBTC2MintInvoice, signature []byte) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "mint", invoice, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x9fd95687.
//
// Solidity: function mint((bytes32,address,uint256) invoice, bytes signature) returns()
func (_StrBtc2 *StrBtc2Session) Mint(invoice StrBTC2MintInvoice, signature []byte) (*types.Transaction, error) {
	return _StrBtc2.Contract.Mint(&_StrBtc2.TransactOpts, invoice, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x9fd95687.
//
// Solidity: function mint((bytes32,address,uint256) invoice, bytes signature) returns()
func (_StrBtc2 *StrBtc2TransactorSession) Mint(invoice StrBTC2MintInvoice, signature []byte) (*types.Transaction, error) {
	return _StrBtc2.Contract.Mint(&_StrBtc2.TransactOpts, invoice, signature)
}

// MintRewards is a paid mutator transaction binding the contract method 0x21172f3b.
//
// Solidity: function mintRewards(uint256 nonce, uint256 delta, bytes signature) returns()
func (_StrBtc2 *StrBtc2Transactor) MintRewards(opts *bind.TransactOpts, nonce *big.Int, delta *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "mintRewards", nonce, delta, signature)
}

// MintRewards is a paid mutator transaction binding the contract method 0x21172f3b.
//
// Solidity: function mintRewards(uint256 nonce, uint256 delta, bytes signature) returns()
func (_StrBtc2 *StrBtc2Session) MintRewards(nonce *big.Int, delta *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrBtc2.Contract.MintRewards(&_StrBtc2.TransactOpts, nonce, delta, signature)
}

// MintRewards is a paid mutator transaction binding the contract method 0x21172f3b.
//
// Solidity: function mintRewards(uint256 nonce, uint256 delta, bytes signature) returns()
func (_StrBtc2 *StrBtc2TransactorSession) MintRewards(nonce *big.Int, delta *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrBtc2.Contract.MintRewards(&_StrBtc2.TransactOpts, nonce, delta, signature)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 _amount, string BTCAddress) returns()
func (_StrBtc2 *StrBtc2Transactor) Redeem(opts *bind.TransactOpts, _amount *big.Int, BTCAddress string) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "redeem", _amount, BTCAddress)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 _amount, string BTCAddress) returns()
func (_StrBtc2 *StrBtc2Session) Redeem(_amount *big.Int, BTCAddress string) (*types.Transaction, error) {
	return _StrBtc2.Contract.Redeem(&_StrBtc2.TransactOpts, _amount, BTCAddress)
}

// Redeem is a paid mutator transaction binding the contract method 0x24b76fd5.
//
// Solidity: function redeem(uint256 _amount, string BTCAddress) returns()
func (_StrBtc2 *StrBtc2TransactorSession) Redeem(_amount *big.Int, BTCAddress string) (*types.Transaction, error) {
	return _StrBtc2.Contract.Redeem(&_StrBtc2.TransactOpts, _amount, BTCAddress)
}

// SetMaxRewardPercent is a paid mutator transaction binding the contract method 0x9cee9492.
//
// Solidity: function setMaxRewardPercent(uint256 _maxRewardPercent) returns()
func (_StrBtc2 *StrBtc2Transactor) SetMaxRewardPercent(opts *bind.TransactOpts, _maxRewardPercent *big.Int) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "setMaxRewardPercent", _maxRewardPercent)
}

// SetMaxRewardPercent is a paid mutator transaction binding the contract method 0x9cee9492.
//
// Solidity: function setMaxRewardPercent(uint256 _maxRewardPercent) returns()
func (_StrBtc2 *StrBtc2Session) SetMaxRewardPercent(_maxRewardPercent *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.SetMaxRewardPercent(&_StrBtc2.TransactOpts, _maxRewardPercent)
}

// SetMaxRewardPercent is a paid mutator transaction binding the contract method 0x9cee9492.
//
// Solidity: function setMaxRewardPercent(uint256 _maxRewardPercent) returns()
func (_StrBtc2 *StrBtc2TransactorSession) SetMaxRewardPercent(_maxRewardPercent *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.SetMaxRewardPercent(&_StrBtc2.TransactOpts, _maxRewardPercent)
}

// SetMinTimeBetweenRewards is a paid mutator transaction binding the contract method 0x25e0a33f.
//
// Solidity: function setMinTimeBetweenRewards(uint256 _minTimeBetweenRewards) returns()
func (_StrBtc2 *StrBtc2Transactor) SetMinTimeBetweenRewards(opts *bind.TransactOpts, _minTimeBetweenRewards *big.Int) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "setMinTimeBetweenRewards", _minTimeBetweenRewards)
}

// SetMinTimeBetweenRewards is a paid mutator transaction binding the contract method 0x25e0a33f.
//
// Solidity: function setMinTimeBetweenRewards(uint256 _minTimeBetweenRewards) returns()
func (_StrBtc2 *StrBtc2Session) SetMinTimeBetweenRewards(_minTimeBetweenRewards *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.SetMinTimeBetweenRewards(&_StrBtc2.TransactOpts, _minTimeBetweenRewards)
}

// SetMinTimeBetweenRewards is a paid mutator transaction binding the contract method 0x25e0a33f.
//
// Solidity: function setMinTimeBetweenRewards(uint256 _minTimeBetweenRewards) returns()
func (_StrBtc2 *StrBtc2TransactorSession) SetMinTimeBetweenRewards(_minTimeBetweenRewards *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.SetMinTimeBetweenRewards(&_StrBtc2.TransactOpts, _minTimeBetweenRewards)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x5abdb0dc.
//
// Solidity: function setMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StrBtc2 *StrBtc2Transactor) SetMinWithdrawAmount(opts *bind.TransactOpts, _minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "setMinWithdrawAmount", _minWithdrawAmount)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x5abdb0dc.
//
// Solidity: function setMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StrBtc2 *StrBtc2Session) SetMinWithdrawAmount(_minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.SetMinWithdrawAmount(&_StrBtc2.TransactOpts, _minWithdrawAmount)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x5abdb0dc.
//
// Solidity: function setMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StrBtc2 *StrBtc2TransactorSession) SetMinWithdrawAmount(_minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.SetMinWithdrawAmount(&_StrBtc2.TransactOpts, _minWithdrawAmount)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_StrBtc2 *StrBtc2Transactor) SetValidatorRegistry(opts *bind.TransactOpts, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "setValidatorRegistry", _validatorRegistry)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_StrBtc2 *StrBtc2Session) SetValidatorRegistry(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc2.Contract.SetValidatorRegistry(&_StrBtc2.TransactOpts, _validatorRegistry)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_StrBtc2 *StrBtc2TransactorSession) SetValidatorRegistry(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _StrBtc2.Contract.SetValidatorRegistry(&_StrBtc2.TransactOpts, _validatorRegistry)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StrBtc2 *StrBtc2Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StrBtc2 *StrBtc2Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.Transfer(&_StrBtc2.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StrBtc2 *StrBtc2TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.Transfer(&_StrBtc2.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StrBtc2 *StrBtc2Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc2.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StrBtc2 *StrBtc2Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.TransferFrom(&_StrBtc2.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StrBtc2 *StrBtc2TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StrBtc2.Contract.TransferFrom(&_StrBtc2.TransactOpts, from, to, value)
}

// StrBtc2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StrBtc2 contract.
type StrBtc2ApprovalIterator struct {
	Event *StrBtc2Approval // Event containing the contract specifics and raw log

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
func (it *StrBtc2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtc2Approval)
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
		it.Event = new(StrBtc2Approval)
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
func (it *StrBtc2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtc2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtc2Approval represents a Approval event raised by the StrBtc2 contract.
type StrBtc2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StrBtc2 *StrBtc2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StrBtc2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StrBtc2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StrBtc2ApprovalIterator{contract: _StrBtc2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StrBtc2 *StrBtc2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StrBtc2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StrBtc2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtc2Approval)
				if err := _StrBtc2.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StrBtc2 *StrBtc2Filterer) ParseApproval(log types.Log) (*StrBtc2Approval, error) {
	event := new(StrBtc2Approval)
	if err := _StrBtc2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtc2ConverterBurnIterator is returned from FilterConverterBurn and is used to iterate over the raw logs and unpacked data for ConverterBurn events raised by the StrBtc2 contract.
type StrBtc2ConverterBurnIterator struct {
	Event *StrBtc2ConverterBurn // Event containing the contract specifics and raw log

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
func (it *StrBtc2ConverterBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtc2ConverterBurn)
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
		it.Event = new(StrBtc2ConverterBurn)
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
func (it *StrBtc2ConverterBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtc2ConverterBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtc2ConverterBurn represents a ConverterBurn event raised by the StrBtc2 contract.
type StrBtc2ConverterBurn struct {
	Converter common.Address
	From      common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConverterBurn is a free log retrieval operation binding the contract event 0x25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad567.
//
// Solidity: event ConverterBurn(address indexed converter, address indexed from, uint256 amount)
func (_StrBtc2 *StrBtc2Filterer) FilterConverterBurn(opts *bind.FilterOpts, converter []common.Address, from []common.Address) (*StrBtc2ConverterBurnIterator, error) {

	var converterRule []interface{}
	for _, converterItem := range converter {
		converterRule = append(converterRule, converterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StrBtc2.contract.FilterLogs(opts, "ConverterBurn", converterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &StrBtc2ConverterBurnIterator{contract: _StrBtc2.contract, event: "ConverterBurn", logs: logs, sub: sub}, nil
}

// WatchConverterBurn is a free log subscription operation binding the contract event 0x25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad567.
//
// Solidity: event ConverterBurn(address indexed converter, address indexed from, uint256 amount)
func (_StrBtc2 *StrBtc2Filterer) WatchConverterBurn(opts *bind.WatchOpts, sink chan<- *StrBtc2ConverterBurn, converter []common.Address, from []common.Address) (event.Subscription, error) {

	var converterRule []interface{}
	for _, converterItem := range converter {
		converterRule = append(converterRule, converterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StrBtc2.contract.WatchLogs(opts, "ConverterBurn", converterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtc2ConverterBurn)
				if err := _StrBtc2.contract.UnpackLog(event, "ConverterBurn", log); err != nil {
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

// ParseConverterBurn is a log parse operation binding the contract event 0x25af8198e0603d11f941e41fdf3659a6cb1f571031869d21c0401cb12e7ad567.
//
// Solidity: event ConverterBurn(address indexed converter, address indexed from, uint256 amount)
func (_StrBtc2 *StrBtc2Filterer) ParseConverterBurn(log types.Log) (*StrBtc2ConverterBurn, error) {
	event := new(StrBtc2ConverterBurn)
	if err := _StrBtc2.contract.UnpackLog(event, "ConverterBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtc2ConverterMintIterator is returned from FilterConverterMint and is used to iterate over the raw logs and unpacked data for ConverterMint events raised by the StrBtc2 contract.
type StrBtc2ConverterMintIterator struct {
	Event *StrBtc2ConverterMint // Event containing the contract specifics and raw log

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
func (it *StrBtc2ConverterMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtc2ConverterMint)
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
		it.Event = new(StrBtc2ConverterMint)
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
func (it *StrBtc2ConverterMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtc2ConverterMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtc2ConverterMint represents a ConverterMint event raised by the StrBtc2 contract.
type StrBtc2ConverterMint struct {
	Converter common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConverterMint is a free log retrieval operation binding the contract event 0x76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e.
//
// Solidity: event ConverterMint(address indexed converter, address indexed recipient, uint256 amount)
func (_StrBtc2 *StrBtc2Filterer) FilterConverterMint(opts *bind.FilterOpts, converter []common.Address, recipient []common.Address) (*StrBtc2ConverterMintIterator, error) {

	var converterRule []interface{}
	for _, converterItem := range converter {
		converterRule = append(converterRule, converterItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StrBtc2.contract.FilterLogs(opts, "ConverterMint", converterRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &StrBtc2ConverterMintIterator{contract: _StrBtc2.contract, event: "ConverterMint", logs: logs, sub: sub}, nil
}

// WatchConverterMint is a free log subscription operation binding the contract event 0x76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e.
//
// Solidity: event ConverterMint(address indexed converter, address indexed recipient, uint256 amount)
func (_StrBtc2 *StrBtc2Filterer) WatchConverterMint(opts *bind.WatchOpts, sink chan<- *StrBtc2ConverterMint, converter []common.Address, recipient []common.Address) (event.Subscription, error) {

	var converterRule []interface{}
	for _, converterItem := range converter {
		converterRule = append(converterRule, converterItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StrBtc2.contract.WatchLogs(opts, "ConverterMint", converterRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtc2ConverterMint)
				if err := _StrBtc2.contract.UnpackLog(event, "ConverterMint", log); err != nil {
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

// ParseConverterMint is a log parse operation binding the contract event 0x76bf2183ddbd3f44507ad7d1989ec6ce8bb5a1974f0862fbf29060dea8431d0e.
//
// Solidity: event ConverterMint(address indexed converter, address indexed recipient, uint256 amount)
func (_StrBtc2 *StrBtc2Filterer) ParseConverterMint(log types.Log) (*StrBtc2ConverterMint, error) {
	event := new(StrBtc2ConverterMint)
	if err := _StrBtc2.contract.UnpackLog(event, "ConverterMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtc2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrBtc2 contract.
type StrBtc2InitializedIterator struct {
	Event *StrBtc2Initialized // Event containing the contract specifics and raw log

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
func (it *StrBtc2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtc2Initialized)
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
		it.Event = new(StrBtc2Initialized)
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
func (it *StrBtc2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtc2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtc2Initialized represents a Initialized event raised by the StrBtc2 contract.
type StrBtc2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrBtc2 *StrBtc2Filterer) FilterInitialized(opts *bind.FilterOpts) (*StrBtc2InitializedIterator, error) {

	logs, sub, err := _StrBtc2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrBtc2InitializedIterator{contract: _StrBtc2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrBtc2 *StrBtc2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrBtc2Initialized) (event.Subscription, error) {

	logs, sub, err := _StrBtc2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtc2Initialized)
				if err := _StrBtc2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StrBtc2 *StrBtc2Filterer) ParseInitialized(log types.Log) (*StrBtc2Initialized, error) {
	event := new(StrBtc2Initialized)
	if err := _StrBtc2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtc2MintBtcEventIterator is returned from FilterMintBtcEvent and is used to iterate over the raw logs and unpacked data for MintBtcEvent events raised by the StrBtc2 contract.
type StrBtc2MintBtcEventIterator struct {
	Event *StrBtc2MintBtcEvent // Event containing the contract specifics and raw log

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
func (it *StrBtc2MintBtcEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtc2MintBtcEvent)
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
		it.Event = new(StrBtc2MintBtcEvent)
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
func (it *StrBtc2MintBtcEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtc2MintBtcEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtc2MintBtcEvent represents a MintBtcEvent event raised by the StrBtc2 contract.
type StrBtc2MintBtcEvent struct {
	To           common.Address
	Value        *big.Int
	BtcDepositId [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintBtcEvent is a free log retrieval operation binding the contract event 0xb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa469.
//
// Solidity: event MintBtcEvent(address indexed _to, uint256 _value, bytes32 _btcDepositId)
func (_StrBtc2 *StrBtc2Filterer) FilterMintBtcEvent(opts *bind.FilterOpts, _to []common.Address) (*StrBtc2MintBtcEventIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StrBtc2.contract.FilterLogs(opts, "MintBtcEvent", _toRule)
	if err != nil {
		return nil, err
	}
	return &StrBtc2MintBtcEventIterator{contract: _StrBtc2.contract, event: "MintBtcEvent", logs: logs, sub: sub}, nil
}

// WatchMintBtcEvent is a free log subscription operation binding the contract event 0xb73f3e96d1e157f064cb3a8d0abed06bcec05e5515bf7486364c027dab6aa469.
//
// Solidity: event MintBtcEvent(address indexed _to, uint256 _value, bytes32 _btcDepositId)
func (_StrBtc2 *StrBtc2Filterer) WatchMintBtcEvent(opts *bind.WatchOpts, sink chan<- *StrBtc2MintBtcEvent, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StrBtc2.contract.WatchLogs(opts, "MintBtcEvent", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtc2MintBtcEvent)
				if err := _StrBtc2.contract.UnpackLog(event, "MintBtcEvent", log); err != nil {
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
func (_StrBtc2 *StrBtc2Filterer) ParseMintBtcEvent(log types.Log) (*StrBtc2MintBtcEvent, error) {
	event := new(StrBtc2MintBtcEvent)
	if err := _StrBtc2.contract.UnpackLog(event, "MintBtcEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtc2RedeemBtcEventIterator is returned from FilterRedeemBtcEvent and is used to iterate over the raw logs and unpacked data for RedeemBtcEvent events raised by the StrBtc2 contract.
type StrBtc2RedeemBtcEventIterator struct {
	Event *StrBtc2RedeemBtcEvent // Event containing the contract specifics and raw log

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
func (it *StrBtc2RedeemBtcEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtc2RedeemBtcEvent)
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
		it.Event = new(StrBtc2RedeemBtcEvent)
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
func (it *StrBtc2RedeemBtcEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtc2RedeemBtcEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtc2RedeemBtcEvent represents a RedeemBtcEvent event raised by the StrBtc2 contract.
type StrBtc2RedeemBtcEvent struct {
	From       common.Address
	BTCAddress string
	Value      *big.Int
	Id         *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedeemBtcEvent is a free log retrieval operation binding the contract event 0x83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e.
//
// Solidity: event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id)
func (_StrBtc2 *StrBtc2Filterer) FilterRedeemBtcEvent(opts *bind.FilterOpts, _from []common.Address) (*StrBtc2RedeemBtcEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _StrBtc2.contract.FilterLogs(opts, "RedeemBtcEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &StrBtc2RedeemBtcEventIterator{contract: _StrBtc2.contract, event: "RedeemBtcEvent", logs: logs, sub: sub}, nil
}

// WatchRedeemBtcEvent is a free log subscription operation binding the contract event 0x83c16822c691a011b471d2653b84faff158a050c4e117390a6c008ecdefcc14e.
//
// Solidity: event RedeemBtcEvent(address indexed _from, string _BTCAddress, uint256 _value, uint256 _id)
func (_StrBtc2 *StrBtc2Filterer) WatchRedeemBtcEvent(opts *bind.WatchOpts, sink chan<- *StrBtc2RedeemBtcEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _StrBtc2.contract.WatchLogs(opts, "RedeemBtcEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtc2RedeemBtcEvent)
				if err := _StrBtc2.contract.UnpackLog(event, "RedeemBtcEvent", log); err != nil {
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
func (_StrBtc2 *StrBtc2Filterer) ParseRedeemBtcEvent(log types.Log) (*StrBtc2RedeemBtcEvent, error) {
	event := new(StrBtc2RedeemBtcEvent)
	if err := _StrBtc2.contract.UnpackLog(event, "RedeemBtcEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtc2TotalSupplyUpdatedEventIterator is returned from FilterTotalSupplyUpdatedEvent and is used to iterate over the raw logs and unpacked data for TotalSupplyUpdatedEvent events raised by the StrBtc2 contract.
type StrBtc2TotalSupplyUpdatedEventIterator struct {
	Event *StrBtc2TotalSupplyUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *StrBtc2TotalSupplyUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtc2TotalSupplyUpdatedEvent)
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
		it.Event = new(StrBtc2TotalSupplyUpdatedEvent)
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
func (it *StrBtc2TotalSupplyUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtc2TotalSupplyUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtc2TotalSupplyUpdatedEvent represents a TotalSupplyUpdatedEvent event raised by the StrBtc2 contract.
type StrBtc2TotalSupplyUpdatedEvent struct {
	Nonce          *big.Int
	TotalBTCSupply *big.Int
	TotalShares    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalSupplyUpdatedEvent is a free log retrieval operation binding the contract event 0x339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b93440.
//
// Solidity: event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares)
func (_StrBtc2 *StrBtc2Filterer) FilterTotalSupplyUpdatedEvent(opts *bind.FilterOpts) (*StrBtc2TotalSupplyUpdatedEventIterator, error) {

	logs, sub, err := _StrBtc2.contract.FilterLogs(opts, "TotalSupplyUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return &StrBtc2TotalSupplyUpdatedEventIterator{contract: _StrBtc2.contract, event: "TotalSupplyUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchTotalSupplyUpdatedEvent is a free log subscription operation binding the contract event 0x339ea31e567d96bc11133446c07d2afa7b1a67accc22bd1b6149fd58d1b93440.
//
// Solidity: event TotalSupplyUpdatedEvent(uint256 _nonce, uint256 _totalBTCSupply, uint256 _totalShares)
func (_StrBtc2 *StrBtc2Filterer) WatchTotalSupplyUpdatedEvent(opts *bind.WatchOpts, sink chan<- *StrBtc2TotalSupplyUpdatedEvent) (event.Subscription, error) {

	logs, sub, err := _StrBtc2.contract.WatchLogs(opts, "TotalSupplyUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtc2TotalSupplyUpdatedEvent)
				if err := _StrBtc2.contract.UnpackLog(event, "TotalSupplyUpdatedEvent", log); err != nil {
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
func (_StrBtc2 *StrBtc2Filterer) ParseTotalSupplyUpdatedEvent(log types.Log) (*StrBtc2TotalSupplyUpdatedEvent, error) {
	event := new(StrBtc2TotalSupplyUpdatedEvent)
	if err := _StrBtc2.contract.UnpackLog(event, "TotalSupplyUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrBtc2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StrBtc2 contract.
type StrBtc2TransferIterator struct {
	Event *StrBtc2Transfer // Event containing the contract specifics and raw log

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
func (it *StrBtc2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrBtc2Transfer)
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
		it.Event = new(StrBtc2Transfer)
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
func (it *StrBtc2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrBtc2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrBtc2Transfer represents a Transfer event raised by the StrBtc2 contract.
type StrBtc2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StrBtc2 *StrBtc2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StrBtc2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StrBtc2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StrBtc2TransferIterator{contract: _StrBtc2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StrBtc2 *StrBtc2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StrBtc2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StrBtc2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrBtc2Transfer)
				if err := _StrBtc2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StrBtc2 *StrBtc2Filterer) ParseTransfer(log types.Log) (*StrBtc2Transfer, error) {
	event := new(StrBtc2Transfer)
	if err := _StrBtc2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
