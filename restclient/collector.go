package restclient

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/log"
)

type Collector interface {
	HandleGenesis(genesisState map[string]json.RawMessage) error
	HandlePrevBlock(block *tmservice.GetBlockByHeightResponse) error
	HandleBlock(block *tmservice.GetBlockByHeightResponse, txs []*tx.GetTxResponse) error
	Logger() log.Logger
}

// 同步数据
func (c *RestClient) Collect(ctx context.Context, start int64, collector Collector) {
	var errCount int64
	sleepFun := func(err error, key string, value interface{}) {
		errCount++
		collector.Logger().Error("collect failed", key, value, "error", err)
		time.Sleep(time.Duration(errCount) * time.Second)
	}

	// 默认从1开始
	if start == 0 {
		if err := collector.HandleGenesis(nil); err != nil {
			panic(err)
		}
		start = 1
	}

	if start > 1 {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			block, err := c.BlockByHeight(start - 1)
			if err != nil {
				sleepFun(err, "BlockByHeight", start-1)
				continue
			}
			if err := collector.HandlePrevBlock(block); err != nil {
				panic(err)
			}
			break
		}
	}

	for {
	relay:
		select {
		case <-ctx.Done():
			return
		default:
		}

		block, err := c.BlockByHeight(start)
		if err != nil {
			if strings.Contains(err.Error(), "block height is bigger then the chain length") {
				time.Sleep(time.Second)
			} else {
				sleepFun(err, "BlockByHeight", start)
			}
			continue
		}
		txs := make([]*tx.GetTxResponse, len(block.Block.Data.Txs))
		for i, txBytes := range block.Block.Data.Txs {
			hash := hex.EncodeToString(tmhash.Sum(txBytes))
			tx, err := c.Tx(hash)
			if err != nil {
				sleepFun(err, "Tx", hash)
				goto relay
			}
			tx.Tx.UnpackInterfaces(c.encodingConfig.InterfaceRegistry)
			txs[i] = tx
		}

		// 处理区块链
		if err := collector.HandleBlock(block, txs); err != nil {
			sleepFun(err, "HandleBlock", start)
			continue
		}

		collector.Logger().Debug("collect succeed", "height", start)
		errCount = 0
		// 处理世界状态
		start++
	}
}