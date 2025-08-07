package adapter

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/kitelabs-io/ethrpc/adapter/ethereum"
)

func New(chainID uint, url string, options ...rpc.ClientOption) (EthClientAdapter, error) {
	switch chainID {
	default:
		return ethereum.NewAdapter(url, options...)
	}
}
