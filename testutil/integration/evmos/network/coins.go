package network

import (
	"github.com/evmos/evmos/v20/utils"
	evmtypes "github.com/evmos/evmos/v20/x/evm/types"
)

type CoinInfo struct {
	Denom    string
	Decimals evmtypes.Decimals
}

// ChainCoins information for the coins required from the chian to operate:
// - baseCoin: represents the base coin used to pay gas fees and staking in the
// Cosmos context.
// - evmCoin: represents the evm coin used to pay Ethereum
// transactions fees.
type ChainCoins struct {
	// TODO: not sure if this is an overkill. Do we want to customize the
	// decimals of the base denom? Maybe not..
	baseCoin *CoinInfo
	evmCoin  *CoinInfo
}

func (cc ChainCoins) BaseCoin() CoinInfo {
	return *cc.baseCoin
}

func (cc ChainCoins) EVMCoin() CoinInfo {
	return *cc.evmCoin
}

func (cc ChainCoins) BaseDenom() string {
	return cc.baseCoin.Denom
}

func (cc ChainCoins) EVMDenom() string {
	return cc.evmCoin.Denom
}

func (cc ChainCoins) BaseDecimals() evmtypes.Decimals {
	return cc.baseCoin.Decimals
}

func (cc ChainCoins) EVMDecimals() evmtypes.Decimals {
	return cc.evmCoin.Decimals
}

func (cc ChainCoins) IsBaseEqualToEVM() bool {
	return cc.BaseDenom() == cc.EVMDenom()
}

// DefaultChainCoins returns the default values used for the ChainCoins in which
// base and evm denom are the same.
func DefaultChainCoins() ChainCoins {
	baseCoinInfo := evmtypes.ChainsCoinInfo[utils.MainnetChainID]
	// baseCoin is used for both base and evm coin as default..
	baseCoin := CoinInfo{
		Denom:    baseCoinInfo.Denom,
		Decimals: baseCoinInfo.Decimals,
	}
	evmCoin := baseCoin
	return ChainCoins{
		baseCoin: &baseCoin,
		evmCoin:  &evmCoin,
	}
}
