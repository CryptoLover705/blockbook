package blackcoin

import (
	"bytes"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
	"github.com/trezor/blockbook/bchain/coins/utils"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0xf9beb4d9
	TestnetMagic wire.BitcoinNet = 0x0b110907 // See https://github.com/CoinBlack/blackcoin-more/blob/master/qa/rpc-tests/test_framework/mininode.py#L1404
)

// chain parameters
var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{25}
	MainNetParams.ScriptHashAddrID = []byte{85}

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{111} 
	TestNetParams.ScriptHashAddrID = []byte{196} 
}

// BlackCoinParser handle
type BlackCoinParser struct {
	*btc.BitcoinLikeParser
}

// NewBlackCoinParser returns new BlackCoinParser instance
func NewBlackCoinParser(params *chaincfg.Params, c *btc.Configuration) *BlackCoinParser {
	return &BlackCoinParser{BitcoinLikeParser: btc.NewBitcoinLikeParser(params, c)}
}

// GetChainParams contains network parameters for the main BlackCoin network,
// and the test BlackCoin network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}


