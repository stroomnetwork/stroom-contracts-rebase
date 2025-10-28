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

// UserActivatorMetaData contains all meta data concerning the UserActivator contract.
var UserActivatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateUser\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activationLimitsEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"btcDeriver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBTCDepositAddressDeriver\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dailyActivationLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBTCDepositAddress\",\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemainingActivationLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUsedActivationCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserTokenId\",\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_btcDeriver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActivated\",\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setActivationLimitsEnabled\",\"inputs\":[{\"name\":\"_enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBTCDeriver\",\"inputs\":[{\"name\":\"_btcDeriver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDailyActivationLimit\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActivationLimitsEnabled\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BTCDeriverUpdated\",\"inputs\":[{\"name\":\"newDeriver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DailyActivationLimitUpdated\",\"inputs\":[{\"name\":\"newLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserAddressActivated\",\"inputs\":[{\"name\":\"userETHAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DailyActivationLimitExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidBTCDeriverAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SeedNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenIsNonTransferable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserAlreadyActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserNotActivated\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611c30806100d65f395ff3fe608060405234801561000f575f5ffd5b50600436106101c6575f3560e01c80638da5cb5b116100fe578063c9faa7c51161009e578063f2fde38b1161006e578063f2fde38b146103ba578063fefd4e4f146103cd578063ff608c29146103f5578063ff698fa414610408575f5ffd5b8063c9faa7c514610383578063d90d693814610396578063e71c76a91461039f578063e985e9c5146103a7575f5ffd5b8063a0ddd62f116100d9578063a0ddd62f14610337578063a22cb4651461034a578063b88d4fde1461035d578063c87b56dd14610370575f5ffd5b80638da5cb5b146102ec57806395d89b411461031c578063a0a1c72214610324575f5ffd5b806342842e0e11610169578063640c40d211610144578063640c40d2146102b357806370a08231146102c9578063715018a6146102dc57806374d99703146102e4575f5ffd5b806342842e0e1461027a578063485cc9551461028d5780636352211e146102a0575f5ffd5b8063095ea7b3116101a4578063095ea7b3146102325780630e155a841461024757806323b872dd1461025457806325d4d8fa14610267575f5ffd5b806301ffc9a7146101ca57806306fdde03146101f2578063081812fc14610207575b5f5ffd5b6101dd6101d836600461167a565b61041a565b60405190151581526020015b60405180910390f35b6101fa61046b565b6040516101e991906116c3565b61021a6102153660046116d5565b61050c565b6040516001600160a01b0390911681526020016101e9565b610245610240366004611707565b610520565b005b6002546101dd9060ff1681565b61024561026236600461172f565b61052f565b610245610275366004611776565b6105bd565b61024561028836600461172f565b61060d565b61024561029b366004611791565b61062c565b61021a6102ae3660046116d5565b6107e7565b6102bb6107f1565b6040519081526020016101e9565b6102bb6102d73660046117c2565b61081e565b610245610876565b6102bb610889565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661021a565b6101fa6108e8565b6102456103323660046116d5565b610926565b6101fa6103453660046117c2565b610963565b6102456103583660046117db565b610a19565b61024561036b36600461187c565b610a24565b6101fa61037e3660046116d5565b610a3c565b6101dd6103913660046117c2565b610aac565b6102bb60015481565b610245610abe565b6101dd6103b5366004611791565b610bef565b6102456103c83660046117c2565b610c3b565b6102bb6103db3660046117c2565b6001600160a01b03165f9081526006602052604090205490565b6102456104033660046117c2565b610c78565b5f5461021a906001600160a01b031681565b5f6001600160e01b031982166380ac58cd60e01b148061044a57506001600160e01b03198216635b5e139f60e01b145b8061046557506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f516020611bdb5f395f51905f52805460609190819061048a90611920565b80601f01602080910402602001604051908101604052809291908181526020018280546104b690611920565b80156105015780601f106104d857610100808354040283529160200191610501565b820191905f5260205f20905b8154815290600101906020018083116104e457829003601f168201915b505050505091505090565b5f61051682610cee565b5061046582610d25565b61052b828233610d5e565b5050565b6001600160a01b03821661055d57604051633250574960e11b81525f60048201526024015b60405180910390fd5b5f610569838333610d6b565b9050836001600160a01b0316816001600160a01b0316146105b7576040516364283d7b60e01b81526001600160a01b0380861660048301526024820184905282166044820152606401610554565b50505050565b6105c5610daa565b6002805460ff19168215159081179091556040519081527f89465a2de545d032569cd4407e855fa5ac2aa4ab39e33a564fffaf8efaf73601906020015b60405180910390a150565b61062783838360405180602001604052805f815250610a24565b505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156106715750825b90505f8267ffffffffffffffff16600114801561068d5750303b155b90508115801561069b575080155b156106b95760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156106e357845460ff60401b1916600160401b1785555b6107386040518060400160405280601581526020017414dd1c9bdbdb481058dd1a5d985d1a5bdb88139195605a1b815250604051806040016040528060068152602001651cdd1c93919560d21b815250610e05565b61074187610e17565b6001600160a01b03861661076857604051630916d7ff60e31b815260040160405180910390fd5b5f80546001600160a01b0319166001600160a01b038816179055606460019081556002805460ff19168217905560055583156107de57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b5f61046582610cee565b5f80610800620151804261196c565b9050806003600101541015610816575f91505090565b505060035490565b5f5f516020611bdb5f395f51905f526001600160a01b038316610856576040516322718ad960e21b81525f6004820152602401610554565b6001600160a01b039092165f908152600390920160205250604090205490565b61087e610daa565b6108875f610e28565b565b6002545f9060ff1661089b57505f1990565b5f6108a9620151804261196c565b90508060036001015410156108c057505060015490565b600154600354106108d2575f91505090565b6003546001546108e2919061198b565b91505090565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930180546060915f516020611bdb5f395f51905f529161048a90611920565b61092e610daa565b60018190556040518181527f0784593f0545d2d36b4be2ef6a4b8d6f71aa885e97ef066a89b17eb39517268690602001610602565b60605f610984836001600160a01b03165f9081526006602052604090205490565b9050805f036109a6576040516332e4c63160e01b815260040160405180910390fd5b5f546040516321613a0b60e11b8152600481018390526001600160a01b03909116906342c27416906024015f60405180830381865afa1580156109eb573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610a12919081019061199e565b9392505050565b61052b338383610e98565b610a2f84848461052f565b6105b73385858585610f47565b6060610a4782610cee565b505f610a5d60408051602081019091525f815290565b90505f815111610a7b5760405180602001604052805f815250610a12565b80610a858461106f565b604051602001610a96929190611a2a565b6040516020818303038152906040529392505050565b5f5f610ab78361081e565b1192915050565b610ac73361081e565b15610ae557604051636adcde4b60e11b815260040160405180910390fd5b5f5f9054906101000a90046001600160a01b03166001600160a01b0316639fa4989c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b34573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b589190611a3e565b610b7557604051635b4f9e2b60e01b815260040160405180910390fd5b60025460ff1615610b8857610b886110ff565b600580549081905f610b9983611a59565b9091555050335f818152600660205260409020829055610bb99082611166565b60408051338152602081018390527f41af0deb3995bbc8a8f0c329239a2e6b4598a89e4c3e9c0993c20d42829eafed9101610602565b6001600160a01b039182165f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793056020908152604080832093909416825291909152205460ff1690565b610c43610daa565b6001600160a01b038116610c6c57604051631e4fbdf760e01b81525f6004820152602401610554565b610c7581610e28565b50565b610c80610daa565b6001600160a01b038116610ca757604051630916d7ff60e31b815260040160405180910390fd5b5f80546001600160a01b0319166001600160a01b038316908117825560405190917ffaf355d7cd4e6595d06548dd2198e0f8d64ec681a95c3415d4b0da6202b78a1991a250565b5f5f610cf98361117f565b90506001600160a01b03811661046557604051637e27328960e01b815260048101849052602401610554565b5f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260409020546001600160a01b031690565b61062783838360016111b8565b5f5f610d788585856112cb565b90506001600160a01b03811615610da257604051637524113960e01b815260040160405180910390fd5b949350505050565b33610ddc7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146108875760405163118cdaa760e01b8152336004820152602401610554565b610e0d6113cd565b61052b8282611416565b610e1f6113cd565b610c7581611446565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f516020611bdb5f395f51905f526001600160a01b038316610ed857604051630b61174360e31b81526001600160a01b0384166004820152602401610554565b6001600160a01b038481165f818152600584016020908152604080832094881680845294825291829020805460ff191687151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a350505050565b6001600160a01b0383163b1561106857604051630a85bd0160e11b81526001600160a01b0384169063150b7a0290610f89908890889087908790600401611a71565b6020604051808303815f875af1925050508015610fc3575060408051601f3d908101601f19168201909252610fc091810190611aad565b60015b61102a573d808015610ff0576040519150601f19603f3d011682016040523d82523d5f602084013e610ff5565b606091505b5080515f0361102257604051633250574960e11b81526001600160a01b0385166004820152602401610554565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b1461106657604051633250574960e11b81526001600160a01b0385166004820152602401610554565b505b5050505050565b60605f61107b8361144e565b60010190505f8167ffffffffffffffff81111561109a5761109a611810565b6040519080825280601f01601f1916602001820160405280156110c4576020820181803683370190505b5090508181016020015b5f19016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a85049450846110ce57509392505050565b5f61110d620151804261196c565b9050806003600101541015611126575f60035560048190555b6001546003541061114a576040516302d6666160e11b815260040160405180910390fd5b600160035f015f82825461115e9190611ac8565b909155505050565b61052b828260405180602001604052805f815250611525565b5f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260409020546001600160a01b031690565b5f516020611bdb5f395f51905f5281806111da57506001600160a01b03831615155b1561129b575f6111e985610cee565b90506001600160a01b038416158015906112155750836001600160a01b0316816001600160a01b031614155b801561122857506112268185610bef565b155b156112515760405163a9fbf51f60e01b81526001600160a01b0385166004820152602401610554565b82156112995784866001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b5f93845260040160205250506040902080546001600160a01b0319166001600160a01b0392909216919091179055565b5f5f516020611bdb5f395f51905f52816112e48561117f565b90506001600160a01b038416156113005761130081858761153c565b6001600160a01b0381161561133c5761131b5f865f5f6111b8565b6001600160a01b0381165f908152600383016020526040902080545f190190555b6001600160a01b0386161561136c576001600160a01b0386165f9081526003830160205260409020805460010190555b5f85815260028301602052604080822080546001600160a01b0319166001600160a01b038a811691821790925591518893918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a495945050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661088757604051631afcd79f60e31b815260040160405180910390fd5b61141e6113cd565b5f516020611bdb5f395f51905f52806114378482611b1f565b50600181016105b78382611b1f565b610c436113cd565b5f8072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b831061148c5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106114b8576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106114d657662386f26fc10000830492506010015b6305f5e10083106114ee576305f5e100830492506008015b612710831061150257612710830492506004015b60648310611514576064830492506002015b600a83106104655760010192915050565b61152f83836115a0565b610627335f858585610f47565b611547838383611601565b610627576001600160a01b03831661157557604051637e27328960e01b815260048101829052602401610554565b60405163177e802f60e01b81526001600160a01b038316600482015260248101829052604401610554565b6001600160a01b0382166115c957604051633250574960e11b81525f6004820152602401610554565b5f6115d583835f610d6b565b90506001600160a01b03811615610627576040516339e3563760e11b81525f6004820152602401610554565b5f6001600160a01b03831615801590610da25750826001600160a01b0316846001600160a01b0316148061163a575061163a8484610bef565b80610da25750826001600160a01b031661165383610d25565b6001600160a01b031614949350505050565b6001600160e01b031981168114610c75575f5ffd5b5f6020828403121561168a575f5ffd5b8135610a1281611665565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610a126020830184611695565b5f602082840312156116e5575f5ffd5b5035919050565b80356001600160a01b0381168114611702575f5ffd5b919050565b5f5f60408385031215611718575f5ffd5b611721836116ec565b946020939093013593505050565b5f5f5f60608486031215611741575f5ffd5b61174a846116ec565b9250611758602085016116ec565b929592945050506040919091013590565b8015158114610c75575f5ffd5b5f60208284031215611786575f5ffd5b8135610a1281611769565b5f5f604083850312156117a2575f5ffd5b6117ab836116ec565b91506117b9602084016116ec565b90509250929050565b5f602082840312156117d2575f5ffd5b610a12826116ec565b5f5f604083850312156117ec575f5ffd5b6117f5836116ec565b9150602083013561180581611769565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561184d5761184d611810565b604052919050565b5f67ffffffffffffffff82111561186e5761186e611810565b50601f01601f191660200190565b5f5f5f5f6080858703121561188f575f5ffd5b611898856116ec565b93506118a6602086016116ec565b925060408501359150606085013567ffffffffffffffff8111156118c8575f5ffd5b8501601f810187136118d8575f5ffd5b80356118eb6118e682611855565b611824565b8181528860208385010111156118ff575f5ffd5b816020840160208301375f6020838301015280935050505092959194509250565b600181811c9082168061193457607f821691505b60208210810361195257634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b5f8261198657634e487b7160e01b5f52601260045260245ffd5b500490565b8181038181111561046557610465611958565b5f602082840312156119ae575f5ffd5b815167ffffffffffffffff8111156119c4575f5ffd5b8201601f810184136119d4575f5ffd5b80516119e26118e682611855565b8181528560208385010111156119f6575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b5f81518060208401855e5f93019283525090919050565b5f610da2611a388386611a13565b84611a13565b5f60208284031215611a4e575f5ffd5b8151610a1281611769565b5f60018201611a6a57611a6a611958565b5060010190565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f90611aa390830184611695565b9695505050505050565b5f60208284031215611abd575f5ffd5b8151610a1281611665565b8082018082111561046557610465611958565b601f82111561062757805f5260205f20601f840160051c81016020851015611b005750805b601f840160051c820191505b81811015611068575f8155600101611b0c565b815167ffffffffffffffff811115611b3957611b39611810565b611b4d81611b478454611920565b84611adb565b6020601f821160018114611b7f575f8315611b685750848201515b5f19600385901b1c1916600184901b178455611068565b5f84815260208120601f198516915b82811015611bae5787850151825560209485019460019092019101611b8e565b5084821015611bcb57868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fe80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300a2646970667358221220a9ebbf36ab5aa6fb6518f3f1a03566e4ec5ef4e43bdc9d865a6ab2ef45f0f16664736f6c634300081b0033",
}

// UserActivatorABI is the input ABI used to generate the binding from.
// Deprecated: Use UserActivatorMetaData.ABI instead.
var UserActivatorABI = UserActivatorMetaData.ABI

// UserActivatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserActivatorMetaData.Bin instead.
var UserActivatorBin = UserActivatorMetaData.Bin

// DeployUserActivator deploys a new Ethereum contract, binding an instance of UserActivator to it.
func DeployUserActivator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserActivator, error) {
	parsed, err := UserActivatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserActivatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserActivator{UserActivatorCaller: UserActivatorCaller{contract: contract}, UserActivatorTransactor: UserActivatorTransactor{contract: contract}, UserActivatorFilterer: UserActivatorFilterer{contract: contract}}, nil
}

// UserActivator is an auto generated Go binding around an Ethereum contract.
type UserActivator struct {
	UserActivatorCaller     // Read-only binding to the contract
	UserActivatorTransactor // Write-only binding to the contract
	UserActivatorFilterer   // Log filterer for contract events
}

// UserActivatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserActivatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserActivatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserActivatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserActivatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserActivatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserActivatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserActivatorSession struct {
	Contract     *UserActivator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserActivatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserActivatorCallerSession struct {
	Contract *UserActivatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UserActivatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserActivatorTransactorSession struct {
	Contract     *UserActivatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UserActivatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserActivatorRaw struct {
	Contract *UserActivator // Generic contract binding to access the raw methods on
}

// UserActivatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserActivatorCallerRaw struct {
	Contract *UserActivatorCaller // Generic read-only contract binding to access the raw methods on
}

// UserActivatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserActivatorTransactorRaw struct {
	Contract *UserActivatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserActivator creates a new instance of UserActivator, bound to a specific deployed contract.
func NewUserActivator(address common.Address, backend bind.ContractBackend) (*UserActivator, error) {
	contract, err := bindUserActivator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserActivator{UserActivatorCaller: UserActivatorCaller{contract: contract}, UserActivatorTransactor: UserActivatorTransactor{contract: contract}, UserActivatorFilterer: UserActivatorFilterer{contract: contract}}, nil
}

// NewUserActivatorCaller creates a new read-only instance of UserActivator, bound to a specific deployed contract.
func NewUserActivatorCaller(address common.Address, caller bind.ContractCaller) (*UserActivatorCaller, error) {
	contract, err := bindUserActivator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserActivatorCaller{contract: contract}, nil
}

// NewUserActivatorTransactor creates a new write-only instance of UserActivator, bound to a specific deployed contract.
func NewUserActivatorTransactor(address common.Address, transactor bind.ContractTransactor) (*UserActivatorTransactor, error) {
	contract, err := bindUserActivator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserActivatorTransactor{contract: contract}, nil
}

// NewUserActivatorFilterer creates a new log filterer instance of UserActivator, bound to a specific deployed contract.
func NewUserActivatorFilterer(address common.Address, filterer bind.ContractFilterer) (*UserActivatorFilterer, error) {
	contract, err := bindUserActivator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserActivatorFilterer{contract: contract}, nil
}

// bindUserActivator binds a generic wrapper to an already deployed contract.
func bindUserActivator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserActivatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserActivator *UserActivatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserActivator.Contract.UserActivatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserActivator *UserActivatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserActivator.Contract.UserActivatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserActivator *UserActivatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserActivator.Contract.UserActivatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserActivator *UserActivatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserActivator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserActivator *UserActivatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserActivator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserActivator *UserActivatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserActivator.Contract.contract.Transact(opts, method, params...)
}

// ActivationLimitsEnabled is a free data retrieval call binding the contract method 0x0e155a84.
//
// Solidity: function activationLimitsEnabled() view returns(bool)
func (_UserActivator *UserActivatorCaller) ActivationLimitsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "activationLimitsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActivationLimitsEnabled is a free data retrieval call binding the contract method 0x0e155a84.
//
// Solidity: function activationLimitsEnabled() view returns(bool)
func (_UserActivator *UserActivatorSession) ActivationLimitsEnabled() (bool, error) {
	return _UserActivator.Contract.ActivationLimitsEnabled(&_UserActivator.CallOpts)
}

// ActivationLimitsEnabled is a free data retrieval call binding the contract method 0x0e155a84.
//
// Solidity: function activationLimitsEnabled() view returns(bool)
func (_UserActivator *UserActivatorCallerSession) ActivationLimitsEnabled() (bool, error) {
	return _UserActivator.Contract.ActivationLimitsEnabled(&_UserActivator.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_UserActivator *UserActivatorCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_UserActivator *UserActivatorSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _UserActivator.Contract.BalanceOf(&_UserActivator.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_UserActivator *UserActivatorCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _UserActivator.Contract.BalanceOf(&_UserActivator.CallOpts, owner)
}

// BtcDeriver is a free data retrieval call binding the contract method 0xff698fa4.
//
// Solidity: function btcDeriver() view returns(address)
func (_UserActivator *UserActivatorCaller) BtcDeriver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "btcDeriver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BtcDeriver is a free data retrieval call binding the contract method 0xff698fa4.
//
// Solidity: function btcDeriver() view returns(address)
func (_UserActivator *UserActivatorSession) BtcDeriver() (common.Address, error) {
	return _UserActivator.Contract.BtcDeriver(&_UserActivator.CallOpts)
}

// BtcDeriver is a free data retrieval call binding the contract method 0xff698fa4.
//
// Solidity: function btcDeriver() view returns(address)
func (_UserActivator *UserActivatorCallerSession) BtcDeriver() (common.Address, error) {
	return _UserActivator.Contract.BtcDeriver(&_UserActivator.CallOpts)
}

// DailyActivationLimit is a free data retrieval call binding the contract method 0xd90d6938.
//
// Solidity: function dailyActivationLimit() view returns(uint256)
func (_UserActivator *UserActivatorCaller) DailyActivationLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "dailyActivationLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyActivationLimit is a free data retrieval call binding the contract method 0xd90d6938.
//
// Solidity: function dailyActivationLimit() view returns(uint256)
func (_UserActivator *UserActivatorSession) DailyActivationLimit() (*big.Int, error) {
	return _UserActivator.Contract.DailyActivationLimit(&_UserActivator.CallOpts)
}

// DailyActivationLimit is a free data retrieval call binding the contract method 0xd90d6938.
//
// Solidity: function dailyActivationLimit() view returns(uint256)
func (_UserActivator *UserActivatorCallerSession) DailyActivationLimit() (*big.Int, error) {
	return _UserActivator.Contract.DailyActivationLimit(&_UserActivator.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_UserActivator *UserActivatorCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_UserActivator *UserActivatorSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _UserActivator.Contract.GetApproved(&_UserActivator.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_UserActivator *UserActivatorCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _UserActivator.Contract.GetApproved(&_UserActivator.CallOpts, tokenId)
}

// GetBTCDepositAddress is a free data retrieval call binding the contract method 0xa0ddd62f.
//
// Solidity: function getBTCDepositAddress(address userAddress) view returns(string)
func (_UserActivator *UserActivatorCaller) GetBTCDepositAddress(opts *bind.CallOpts, userAddress common.Address) (string, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "getBTCDepositAddress", userAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBTCDepositAddress is a free data retrieval call binding the contract method 0xa0ddd62f.
//
// Solidity: function getBTCDepositAddress(address userAddress) view returns(string)
func (_UserActivator *UserActivatorSession) GetBTCDepositAddress(userAddress common.Address) (string, error) {
	return _UserActivator.Contract.GetBTCDepositAddress(&_UserActivator.CallOpts, userAddress)
}

// GetBTCDepositAddress is a free data retrieval call binding the contract method 0xa0ddd62f.
//
// Solidity: function getBTCDepositAddress(address userAddress) view returns(string)
func (_UserActivator *UserActivatorCallerSession) GetBTCDepositAddress(userAddress common.Address) (string, error) {
	return _UserActivator.Contract.GetBTCDepositAddress(&_UserActivator.CallOpts, userAddress)
}

// GetRemainingActivationLimit is a free data retrieval call binding the contract method 0x74d99703.
//
// Solidity: function getRemainingActivationLimit() view returns(uint256)
func (_UserActivator *UserActivatorCaller) GetRemainingActivationLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "getRemainingActivationLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingActivationLimit is a free data retrieval call binding the contract method 0x74d99703.
//
// Solidity: function getRemainingActivationLimit() view returns(uint256)
func (_UserActivator *UserActivatorSession) GetRemainingActivationLimit() (*big.Int, error) {
	return _UserActivator.Contract.GetRemainingActivationLimit(&_UserActivator.CallOpts)
}

// GetRemainingActivationLimit is a free data retrieval call binding the contract method 0x74d99703.
//
// Solidity: function getRemainingActivationLimit() view returns(uint256)
func (_UserActivator *UserActivatorCallerSession) GetRemainingActivationLimit() (*big.Int, error) {
	return _UserActivator.Contract.GetRemainingActivationLimit(&_UserActivator.CallOpts)
}

// GetUsedActivationCount is a free data retrieval call binding the contract method 0x640c40d2.
//
// Solidity: function getUsedActivationCount() view returns(uint256)
func (_UserActivator *UserActivatorCaller) GetUsedActivationCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "getUsedActivationCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsedActivationCount is a free data retrieval call binding the contract method 0x640c40d2.
//
// Solidity: function getUsedActivationCount() view returns(uint256)
func (_UserActivator *UserActivatorSession) GetUsedActivationCount() (*big.Int, error) {
	return _UserActivator.Contract.GetUsedActivationCount(&_UserActivator.CallOpts)
}

// GetUsedActivationCount is a free data retrieval call binding the contract method 0x640c40d2.
//
// Solidity: function getUsedActivationCount() view returns(uint256)
func (_UserActivator *UserActivatorCallerSession) GetUsedActivationCount() (*big.Int, error) {
	return _UserActivator.Contract.GetUsedActivationCount(&_UserActivator.CallOpts)
}

// GetUserTokenId is a free data retrieval call binding the contract method 0xfefd4e4f.
//
// Solidity: function getUserTokenId(address userAddress) view returns(uint256)
func (_UserActivator *UserActivatorCaller) GetUserTokenId(opts *bind.CallOpts, userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "getUserTokenId", userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTokenId is a free data retrieval call binding the contract method 0xfefd4e4f.
//
// Solidity: function getUserTokenId(address userAddress) view returns(uint256)
func (_UserActivator *UserActivatorSession) GetUserTokenId(userAddress common.Address) (*big.Int, error) {
	return _UserActivator.Contract.GetUserTokenId(&_UserActivator.CallOpts, userAddress)
}

// GetUserTokenId is a free data retrieval call binding the contract method 0xfefd4e4f.
//
// Solidity: function getUserTokenId(address userAddress) view returns(uint256)
func (_UserActivator *UserActivatorCallerSession) GetUserTokenId(userAddress common.Address) (*big.Int, error) {
	return _UserActivator.Contract.GetUserTokenId(&_UserActivator.CallOpts, userAddress)
}

// IsActivated is a free data retrieval call binding the contract method 0xc9faa7c5.
//
// Solidity: function isActivated(address userAddress) view returns(bool)
func (_UserActivator *UserActivatorCaller) IsActivated(opts *bind.CallOpts, userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "isActivated", userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActivated is a free data retrieval call binding the contract method 0xc9faa7c5.
//
// Solidity: function isActivated(address userAddress) view returns(bool)
func (_UserActivator *UserActivatorSession) IsActivated(userAddress common.Address) (bool, error) {
	return _UserActivator.Contract.IsActivated(&_UserActivator.CallOpts, userAddress)
}

// IsActivated is a free data retrieval call binding the contract method 0xc9faa7c5.
//
// Solidity: function isActivated(address userAddress) view returns(bool)
func (_UserActivator *UserActivatorCallerSession) IsActivated(userAddress common.Address) (bool, error) {
	return _UserActivator.Contract.IsActivated(&_UserActivator.CallOpts, userAddress)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_UserActivator *UserActivatorCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_UserActivator *UserActivatorSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _UserActivator.Contract.IsApprovedForAll(&_UserActivator.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_UserActivator *UserActivatorCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _UserActivator.Contract.IsApprovedForAll(&_UserActivator.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UserActivator *UserActivatorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UserActivator *UserActivatorSession) Name() (string, error) {
	return _UserActivator.Contract.Name(&_UserActivator.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UserActivator *UserActivatorCallerSession) Name() (string, error) {
	return _UserActivator.Contract.Name(&_UserActivator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserActivator *UserActivatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserActivator *UserActivatorSession) Owner() (common.Address, error) {
	return _UserActivator.Contract.Owner(&_UserActivator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserActivator *UserActivatorCallerSession) Owner() (common.Address, error) {
	return _UserActivator.Contract.Owner(&_UserActivator.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_UserActivator *UserActivatorCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_UserActivator *UserActivatorSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _UserActivator.Contract.OwnerOf(&_UserActivator.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_UserActivator *UserActivatorCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _UserActivator.Contract.OwnerOf(&_UserActivator.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UserActivator *UserActivatorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UserActivator *UserActivatorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UserActivator.Contract.SupportsInterface(&_UserActivator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UserActivator *UserActivatorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UserActivator.Contract.SupportsInterface(&_UserActivator.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UserActivator *UserActivatorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UserActivator *UserActivatorSession) Symbol() (string, error) {
	return _UserActivator.Contract.Symbol(&_UserActivator.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UserActivator *UserActivatorCallerSession) Symbol() (string, error) {
	return _UserActivator.Contract.Symbol(&_UserActivator.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_UserActivator *UserActivatorCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _UserActivator.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_UserActivator *UserActivatorSession) TokenURI(tokenId *big.Int) (string, error) {
	return _UserActivator.Contract.TokenURI(&_UserActivator.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_UserActivator *UserActivatorCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _UserActivator.Contract.TokenURI(&_UserActivator.CallOpts, tokenId)
}

// ActivateUser is a paid mutator transaction binding the contract method 0xe71c76a9.
//
// Solidity: function activateUser() returns()
func (_UserActivator *UserActivatorTransactor) ActivateUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "activateUser")
}

// ActivateUser is a paid mutator transaction binding the contract method 0xe71c76a9.
//
// Solidity: function activateUser() returns()
func (_UserActivator *UserActivatorSession) ActivateUser() (*types.Transaction, error) {
	return _UserActivator.Contract.ActivateUser(&_UserActivator.TransactOpts)
}

// ActivateUser is a paid mutator transaction binding the contract method 0xe71c76a9.
//
// Solidity: function activateUser() returns()
func (_UserActivator *UserActivatorTransactorSession) ActivateUser() (*types.Transaction, error) {
	return _UserActivator.Contract.ActivateUser(&_UserActivator.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_UserActivator *UserActivatorTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_UserActivator *UserActivatorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UserActivator.Contract.Approve(&_UserActivator.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_UserActivator *UserActivatorTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UserActivator.Contract.Approve(&_UserActivator.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _btcDeriver) returns()
func (_UserActivator *UserActivatorTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _btcDeriver common.Address) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "initialize", _owner, _btcDeriver)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _btcDeriver) returns()
func (_UserActivator *UserActivatorSession) Initialize(_owner common.Address, _btcDeriver common.Address) (*types.Transaction, error) {
	return _UserActivator.Contract.Initialize(&_UserActivator.TransactOpts, _owner, _btcDeriver)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _btcDeriver) returns()
func (_UserActivator *UserActivatorTransactorSession) Initialize(_owner common.Address, _btcDeriver common.Address) (*types.Transaction, error) {
	return _UserActivator.Contract.Initialize(&_UserActivator.TransactOpts, _owner, _btcDeriver)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UserActivator *UserActivatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UserActivator *UserActivatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UserActivator.Contract.RenounceOwnership(&_UserActivator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UserActivator *UserActivatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UserActivator.Contract.RenounceOwnership(&_UserActivator.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_UserActivator *UserActivatorTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_UserActivator *UserActivatorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UserActivator.Contract.SafeTransferFrom(&_UserActivator.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_UserActivator *UserActivatorTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UserActivator.Contract.SafeTransferFrom(&_UserActivator.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_UserActivator *UserActivatorTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_UserActivator *UserActivatorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _UserActivator.Contract.SafeTransferFrom0(&_UserActivator.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_UserActivator *UserActivatorTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _UserActivator.Contract.SafeTransferFrom0(&_UserActivator.TransactOpts, from, to, tokenId, data)
}

// SetActivationLimitsEnabled is a paid mutator transaction binding the contract method 0x25d4d8fa.
//
// Solidity: function setActivationLimitsEnabled(bool _enabled) returns()
func (_UserActivator *UserActivatorTransactor) SetActivationLimitsEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "setActivationLimitsEnabled", _enabled)
}

// SetActivationLimitsEnabled is a paid mutator transaction binding the contract method 0x25d4d8fa.
//
// Solidity: function setActivationLimitsEnabled(bool _enabled) returns()
func (_UserActivator *UserActivatorSession) SetActivationLimitsEnabled(_enabled bool) (*types.Transaction, error) {
	return _UserActivator.Contract.SetActivationLimitsEnabled(&_UserActivator.TransactOpts, _enabled)
}

// SetActivationLimitsEnabled is a paid mutator transaction binding the contract method 0x25d4d8fa.
//
// Solidity: function setActivationLimitsEnabled(bool _enabled) returns()
func (_UserActivator *UserActivatorTransactorSession) SetActivationLimitsEnabled(_enabled bool) (*types.Transaction, error) {
	return _UserActivator.Contract.SetActivationLimitsEnabled(&_UserActivator.TransactOpts, _enabled)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_UserActivator *UserActivatorTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_UserActivator *UserActivatorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _UserActivator.Contract.SetApprovalForAll(&_UserActivator.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_UserActivator *UserActivatorTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _UserActivator.Contract.SetApprovalForAll(&_UserActivator.TransactOpts, operator, approved)
}

// SetBTCDeriver is a paid mutator transaction binding the contract method 0xff608c29.
//
// Solidity: function setBTCDeriver(address _btcDeriver) returns()
func (_UserActivator *UserActivatorTransactor) SetBTCDeriver(opts *bind.TransactOpts, _btcDeriver common.Address) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "setBTCDeriver", _btcDeriver)
}

// SetBTCDeriver is a paid mutator transaction binding the contract method 0xff608c29.
//
// Solidity: function setBTCDeriver(address _btcDeriver) returns()
func (_UserActivator *UserActivatorSession) SetBTCDeriver(_btcDeriver common.Address) (*types.Transaction, error) {
	return _UserActivator.Contract.SetBTCDeriver(&_UserActivator.TransactOpts, _btcDeriver)
}

// SetBTCDeriver is a paid mutator transaction binding the contract method 0xff608c29.
//
// Solidity: function setBTCDeriver(address _btcDeriver) returns()
func (_UserActivator *UserActivatorTransactorSession) SetBTCDeriver(_btcDeriver common.Address) (*types.Transaction, error) {
	return _UserActivator.Contract.SetBTCDeriver(&_UserActivator.TransactOpts, _btcDeriver)
}

// SetDailyActivationLimit is a paid mutator transaction binding the contract method 0xa0a1c722.
//
// Solidity: function setDailyActivationLimit(uint256 _limit) returns()
func (_UserActivator *UserActivatorTransactor) SetDailyActivationLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "setDailyActivationLimit", _limit)
}

// SetDailyActivationLimit is a paid mutator transaction binding the contract method 0xa0a1c722.
//
// Solidity: function setDailyActivationLimit(uint256 _limit) returns()
func (_UserActivator *UserActivatorSession) SetDailyActivationLimit(_limit *big.Int) (*types.Transaction, error) {
	return _UserActivator.Contract.SetDailyActivationLimit(&_UserActivator.TransactOpts, _limit)
}

// SetDailyActivationLimit is a paid mutator transaction binding the contract method 0xa0a1c722.
//
// Solidity: function setDailyActivationLimit(uint256 _limit) returns()
func (_UserActivator *UserActivatorTransactorSession) SetDailyActivationLimit(_limit *big.Int) (*types.Transaction, error) {
	return _UserActivator.Contract.SetDailyActivationLimit(&_UserActivator.TransactOpts, _limit)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_UserActivator *UserActivatorTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_UserActivator *UserActivatorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UserActivator.Contract.TransferFrom(&_UserActivator.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_UserActivator *UserActivatorTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _UserActivator.Contract.TransferFrom(&_UserActivator.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UserActivator *UserActivatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UserActivator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UserActivator *UserActivatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UserActivator.Contract.TransferOwnership(&_UserActivator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UserActivator *UserActivatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UserActivator.Contract.TransferOwnership(&_UserActivator.TransactOpts, newOwner)
}

// UserActivatorActivationLimitsEnabledIterator is returned from FilterActivationLimitsEnabled and is used to iterate over the raw logs and unpacked data for ActivationLimitsEnabled events raised by the UserActivator contract.
type UserActivatorActivationLimitsEnabledIterator struct {
	Event *UserActivatorActivationLimitsEnabled // Event containing the contract specifics and raw log

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
func (it *UserActivatorActivationLimitsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserActivatorActivationLimitsEnabled)
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
		it.Event = new(UserActivatorActivationLimitsEnabled)
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
func (it *UserActivatorActivationLimitsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserActivatorActivationLimitsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserActivatorActivationLimitsEnabled represents a ActivationLimitsEnabled event raised by the UserActivator contract.
type UserActivatorActivationLimitsEnabled struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterActivationLimitsEnabled is a free log retrieval operation binding the contract event 0x89465a2de545d032569cd4407e855fa5ac2aa4ab39e33a564fffaf8efaf73601.
//
// Solidity: event ActivationLimitsEnabled(bool enabled)
func (_UserActivator *UserActivatorFilterer) FilterActivationLimitsEnabled(opts *bind.FilterOpts) (*UserActivatorActivationLimitsEnabledIterator, error) {

	logs, sub, err := _UserActivator.contract.FilterLogs(opts, "ActivationLimitsEnabled")
	if err != nil {
		return nil, err
	}
	return &UserActivatorActivationLimitsEnabledIterator{contract: _UserActivator.contract, event: "ActivationLimitsEnabled", logs: logs, sub: sub}, nil
}

// WatchActivationLimitsEnabled is a free log subscription operation binding the contract event 0x89465a2de545d032569cd4407e855fa5ac2aa4ab39e33a564fffaf8efaf73601.
//
// Solidity: event ActivationLimitsEnabled(bool enabled)
func (_UserActivator *UserActivatorFilterer) WatchActivationLimitsEnabled(opts *bind.WatchOpts, sink chan<- *UserActivatorActivationLimitsEnabled) (event.Subscription, error) {

	logs, sub, err := _UserActivator.contract.WatchLogs(opts, "ActivationLimitsEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserActivatorActivationLimitsEnabled)
				if err := _UserActivator.contract.UnpackLog(event, "ActivationLimitsEnabled", log); err != nil {
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

// ParseActivationLimitsEnabled is a log parse operation binding the contract event 0x89465a2de545d032569cd4407e855fa5ac2aa4ab39e33a564fffaf8efaf73601.
//
// Solidity: event ActivationLimitsEnabled(bool enabled)
func (_UserActivator *UserActivatorFilterer) ParseActivationLimitsEnabled(log types.Log) (*UserActivatorActivationLimitsEnabled, error) {
	event := new(UserActivatorActivationLimitsEnabled)
	if err := _UserActivator.contract.UnpackLog(event, "ActivationLimitsEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserActivatorApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UserActivator contract.
type UserActivatorApprovalIterator struct {
	Event *UserActivatorApproval // Event containing the contract specifics and raw log

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
func (it *UserActivatorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserActivatorApproval)
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
		it.Event = new(UserActivatorApproval)
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
func (it *UserActivatorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserActivatorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserActivatorApproval represents a Approval event raised by the UserActivator contract.
type UserActivatorApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_UserActivator *UserActivatorFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*UserActivatorApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UserActivator.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &UserActivatorApprovalIterator{contract: _UserActivator.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_UserActivator *UserActivatorFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UserActivatorApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UserActivator.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserActivatorApproval)
				if err := _UserActivator.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_UserActivator *UserActivatorFilterer) ParseApproval(log types.Log) (*UserActivatorApproval, error) {
	event := new(UserActivatorApproval)
	if err := _UserActivator.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserActivatorApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the UserActivator contract.
type UserActivatorApprovalForAllIterator struct {
	Event *UserActivatorApprovalForAll // Event containing the contract specifics and raw log

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
func (it *UserActivatorApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserActivatorApprovalForAll)
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
		it.Event = new(UserActivatorApprovalForAll)
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
func (it *UserActivatorApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserActivatorApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserActivatorApprovalForAll represents a ApprovalForAll event raised by the UserActivator contract.
type UserActivatorApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_UserActivator *UserActivatorFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*UserActivatorApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _UserActivator.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &UserActivatorApprovalForAllIterator{contract: _UserActivator.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_UserActivator *UserActivatorFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *UserActivatorApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _UserActivator.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserActivatorApprovalForAll)
				if err := _UserActivator.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_UserActivator *UserActivatorFilterer) ParseApprovalForAll(log types.Log) (*UserActivatorApprovalForAll, error) {
	event := new(UserActivatorApprovalForAll)
	if err := _UserActivator.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserActivatorBTCDeriverUpdatedIterator is returned from FilterBTCDeriverUpdated and is used to iterate over the raw logs and unpacked data for BTCDeriverUpdated events raised by the UserActivator contract.
type UserActivatorBTCDeriverUpdatedIterator struct {
	Event *UserActivatorBTCDeriverUpdated // Event containing the contract specifics and raw log

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
func (it *UserActivatorBTCDeriverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserActivatorBTCDeriverUpdated)
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
		it.Event = new(UserActivatorBTCDeriverUpdated)
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
func (it *UserActivatorBTCDeriverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserActivatorBTCDeriverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserActivatorBTCDeriverUpdated represents a BTCDeriverUpdated event raised by the UserActivator contract.
type UserActivatorBTCDeriverUpdated struct {
	NewDeriver common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBTCDeriverUpdated is a free log retrieval operation binding the contract event 0xfaf355d7cd4e6595d06548dd2198e0f8d64ec681a95c3415d4b0da6202b78a19.
//
// Solidity: event BTCDeriverUpdated(address indexed newDeriver)
func (_UserActivator *UserActivatorFilterer) FilterBTCDeriverUpdated(opts *bind.FilterOpts, newDeriver []common.Address) (*UserActivatorBTCDeriverUpdatedIterator, error) {

	var newDeriverRule []interface{}
	for _, newDeriverItem := range newDeriver {
		newDeriverRule = append(newDeriverRule, newDeriverItem)
	}

	logs, sub, err := _UserActivator.contract.FilterLogs(opts, "BTCDeriverUpdated", newDeriverRule)
	if err != nil {
		return nil, err
	}
	return &UserActivatorBTCDeriverUpdatedIterator{contract: _UserActivator.contract, event: "BTCDeriverUpdated", logs: logs, sub: sub}, nil
}

// WatchBTCDeriverUpdated is a free log subscription operation binding the contract event 0xfaf355d7cd4e6595d06548dd2198e0f8d64ec681a95c3415d4b0da6202b78a19.
//
// Solidity: event BTCDeriverUpdated(address indexed newDeriver)
func (_UserActivator *UserActivatorFilterer) WatchBTCDeriverUpdated(opts *bind.WatchOpts, sink chan<- *UserActivatorBTCDeriverUpdated, newDeriver []common.Address) (event.Subscription, error) {

	var newDeriverRule []interface{}
	for _, newDeriverItem := range newDeriver {
		newDeriverRule = append(newDeriverRule, newDeriverItem)
	}

	logs, sub, err := _UserActivator.contract.WatchLogs(opts, "BTCDeriverUpdated", newDeriverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserActivatorBTCDeriverUpdated)
				if err := _UserActivator.contract.UnpackLog(event, "BTCDeriverUpdated", log); err != nil {
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

// ParseBTCDeriverUpdated is a log parse operation binding the contract event 0xfaf355d7cd4e6595d06548dd2198e0f8d64ec681a95c3415d4b0da6202b78a19.
//
// Solidity: event BTCDeriverUpdated(address indexed newDeriver)
func (_UserActivator *UserActivatorFilterer) ParseBTCDeriverUpdated(log types.Log) (*UserActivatorBTCDeriverUpdated, error) {
	event := new(UserActivatorBTCDeriverUpdated)
	if err := _UserActivator.contract.UnpackLog(event, "BTCDeriverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserActivatorDailyActivationLimitUpdatedIterator is returned from FilterDailyActivationLimitUpdated and is used to iterate over the raw logs and unpacked data for DailyActivationLimitUpdated events raised by the UserActivator contract.
type UserActivatorDailyActivationLimitUpdatedIterator struct {
	Event *UserActivatorDailyActivationLimitUpdated // Event containing the contract specifics and raw log

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
func (it *UserActivatorDailyActivationLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserActivatorDailyActivationLimitUpdated)
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
		it.Event = new(UserActivatorDailyActivationLimitUpdated)
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
func (it *UserActivatorDailyActivationLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserActivatorDailyActivationLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserActivatorDailyActivationLimitUpdated represents a DailyActivationLimitUpdated event raised by the UserActivator contract.
type UserActivatorDailyActivationLimitUpdated struct {
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDailyActivationLimitUpdated is a free log retrieval operation binding the contract event 0x0784593f0545d2d36b4be2ef6a4b8d6f71aa885e97ef066a89b17eb395172686.
//
// Solidity: event DailyActivationLimitUpdated(uint256 newLimit)
func (_UserActivator *UserActivatorFilterer) FilterDailyActivationLimitUpdated(opts *bind.FilterOpts) (*UserActivatorDailyActivationLimitUpdatedIterator, error) {

	logs, sub, err := _UserActivator.contract.FilterLogs(opts, "DailyActivationLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &UserActivatorDailyActivationLimitUpdatedIterator{contract: _UserActivator.contract, event: "DailyActivationLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyActivationLimitUpdated is a free log subscription operation binding the contract event 0x0784593f0545d2d36b4be2ef6a4b8d6f71aa885e97ef066a89b17eb395172686.
//
// Solidity: event DailyActivationLimitUpdated(uint256 newLimit)
func (_UserActivator *UserActivatorFilterer) WatchDailyActivationLimitUpdated(opts *bind.WatchOpts, sink chan<- *UserActivatorDailyActivationLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _UserActivator.contract.WatchLogs(opts, "DailyActivationLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserActivatorDailyActivationLimitUpdated)
				if err := _UserActivator.contract.UnpackLog(event, "DailyActivationLimitUpdated", log); err != nil {
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

// ParseDailyActivationLimitUpdated is a log parse operation binding the contract event 0x0784593f0545d2d36b4be2ef6a4b8d6f71aa885e97ef066a89b17eb395172686.
//
// Solidity: event DailyActivationLimitUpdated(uint256 newLimit)
func (_UserActivator *UserActivatorFilterer) ParseDailyActivationLimitUpdated(log types.Log) (*UserActivatorDailyActivationLimitUpdated, error) {
	event := new(UserActivatorDailyActivationLimitUpdated)
	if err := _UserActivator.contract.UnpackLog(event, "DailyActivationLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserActivatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UserActivator contract.
type UserActivatorInitializedIterator struct {
	Event *UserActivatorInitialized // Event containing the contract specifics and raw log

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
func (it *UserActivatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserActivatorInitialized)
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
		it.Event = new(UserActivatorInitialized)
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
func (it *UserActivatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserActivatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserActivatorInitialized represents a Initialized event raised by the UserActivator contract.
type UserActivatorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UserActivator *UserActivatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*UserActivatorInitializedIterator, error) {

	logs, sub, err := _UserActivator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UserActivatorInitializedIterator{contract: _UserActivator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UserActivator *UserActivatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UserActivatorInitialized) (event.Subscription, error) {

	logs, sub, err := _UserActivator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserActivatorInitialized)
				if err := _UserActivator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_UserActivator *UserActivatorFilterer) ParseInitialized(log types.Log) (*UserActivatorInitialized, error) {
	event := new(UserActivatorInitialized)
	if err := _UserActivator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserActivatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UserActivator contract.
type UserActivatorOwnershipTransferredIterator struct {
	Event *UserActivatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UserActivatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserActivatorOwnershipTransferred)
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
		it.Event = new(UserActivatorOwnershipTransferred)
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
func (it *UserActivatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserActivatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserActivatorOwnershipTransferred represents a OwnershipTransferred event raised by the UserActivator contract.
type UserActivatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UserActivator *UserActivatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UserActivatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UserActivator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UserActivatorOwnershipTransferredIterator{contract: _UserActivator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UserActivator *UserActivatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UserActivatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UserActivator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserActivatorOwnershipTransferred)
				if err := _UserActivator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UserActivator *UserActivatorFilterer) ParseOwnershipTransferred(log types.Log) (*UserActivatorOwnershipTransferred, error) {
	event := new(UserActivatorOwnershipTransferred)
	if err := _UserActivator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserActivatorTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UserActivator contract.
type UserActivatorTransferIterator struct {
	Event *UserActivatorTransfer // Event containing the contract specifics and raw log

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
func (it *UserActivatorTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserActivatorTransfer)
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
		it.Event = new(UserActivatorTransfer)
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
func (it *UserActivatorTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserActivatorTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserActivatorTransfer represents a Transfer event raised by the UserActivator contract.
type UserActivatorTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_UserActivator *UserActivatorFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*UserActivatorTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UserActivator.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &UserActivatorTransferIterator{contract: _UserActivator.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_UserActivator *UserActivatorFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UserActivatorTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UserActivator.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserActivatorTransfer)
				if err := _UserActivator.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_UserActivator *UserActivatorFilterer) ParseTransfer(log types.Log) (*UserActivatorTransfer, error) {
	event := new(UserActivatorTransfer)
	if err := _UserActivator.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserActivatorUserAddressActivatedIterator is returned from FilterUserAddressActivated and is used to iterate over the raw logs and unpacked data for UserAddressActivated events raised by the UserActivator contract.
type UserActivatorUserAddressActivatedIterator struct {
	Event *UserActivatorUserAddressActivated // Event containing the contract specifics and raw log

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
func (it *UserActivatorUserAddressActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserActivatorUserAddressActivated)
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
		it.Event = new(UserActivatorUserAddressActivated)
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
func (it *UserActivatorUserAddressActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserActivatorUserAddressActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserActivatorUserAddressActivated represents a UserAddressActivated event raised by the UserActivator contract.
type UserActivatorUserAddressActivated struct {
	UserETHAddress common.Address
	TokenId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUserAddressActivated is a free log retrieval operation binding the contract event 0x41af0deb3995bbc8a8f0c329239a2e6b4598a89e4c3e9c0993c20d42829eafed.
//
// Solidity: event UserAddressActivated(address userETHAddress, uint256 tokenId)
func (_UserActivator *UserActivatorFilterer) FilterUserAddressActivated(opts *bind.FilterOpts) (*UserActivatorUserAddressActivatedIterator, error) {

	logs, sub, err := _UserActivator.contract.FilterLogs(opts, "UserAddressActivated")
	if err != nil {
		return nil, err
	}
	return &UserActivatorUserAddressActivatedIterator{contract: _UserActivator.contract, event: "UserAddressActivated", logs: logs, sub: sub}, nil
}

// WatchUserAddressActivated is a free log subscription operation binding the contract event 0x41af0deb3995bbc8a8f0c329239a2e6b4598a89e4c3e9c0993c20d42829eafed.
//
// Solidity: event UserAddressActivated(address userETHAddress, uint256 tokenId)
func (_UserActivator *UserActivatorFilterer) WatchUserAddressActivated(opts *bind.WatchOpts, sink chan<- *UserActivatorUserAddressActivated) (event.Subscription, error) {

	logs, sub, err := _UserActivator.contract.WatchLogs(opts, "UserAddressActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserActivatorUserAddressActivated)
				if err := _UserActivator.contract.UnpackLog(event, "UserAddressActivated", log); err != nil {
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

// ParseUserAddressActivated is a log parse operation binding the contract event 0x41af0deb3995bbc8a8f0c329239a2e6b4598a89e4c3e9c0993c20d42829eafed.
//
// Solidity: event UserAddressActivated(address userETHAddress, uint256 tokenId)
func (_UserActivator *UserActivatorFilterer) ParseUserAddressActivated(log types.Log) (*UserActivatorUserAddressActivated, error) {
	event := new(UserActivatorUserAddressActivated)
	if err := _UserActivator.contract.UnpackLog(event, "UserAddressActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
