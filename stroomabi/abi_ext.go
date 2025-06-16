package stroomabi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// IdHex returns mint id string
func (m *StrBtcMintBtcEvent) IdHex() string {
	return common.BytesToHash(m.BtcDepositId[:]).Hex()
}

// BitcoinNetwork corresponds to the Contract's network enum.
type BitcoinNetwork int

const (
	BitcoinNetworkMainnet BitcoinNetwork = iota
	BitcoinNetworkTestnet
	BitcoinNetworkRegtest
	BitcoinNetworkSimnet
)

var MintInvoicePrefix = crypto.Keccak256Hash([]byte("STROOM_MINT_INVOICE")).Bytes()

var TotalSupplyUpdatePrefix = crypto.Keccak256Hash([]byte("STROOM_UPDATE_TOTAL_SUPPLY")).Bytes()
