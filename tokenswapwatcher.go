package main

import (
	"context"
	"log"
	"log/slog"
	"math/big"
	"os"

	"github.com/d6o/tokenswaptracker/appcontext"
	"github.com/d6o/tokenswaptracker/contract/uniswapv2"
	"github.com/d6o/tokenswaptracker/handlers"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	level := new(slog.LevelVar)
	level.Set(slog.LevelInfo)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       level,
		ReplaceAttr: nil,
	}))

	ctx := appcontext.WithLogger(context.Background(), logger)

	ethWS := os.Getenv("ETH_WS")
	if ethWS == "" {
		log.Panicf("You should set ETH_WS")
	}

	client, err := ethclient.Dial(ethWS)
	if err != nil {
		log.Panicf("Failed to connect to the Ethereum client: %v", err)
	}

	swapParser, err := uniswapv2.NewContractFilterer(common.Address{}, nil)
	if err != nil {
		log.Panicf("Failed to create contract Filterer: %v", err)
	}
	logger.Debug("test")

	swapHandler := handlers.NewSwap(client)
	transactionHandler := handlers.NewTransaction(swapParser, swapHandler)
	blockHandler := handlers.NewBlock(client, transactionHandler)
	tsw := NewTokenSwapWatcher(client, blockHandler)

	if err := tsw.Run(ctx); err != nil {
		log.Panic(err)
	}
}

type (
	TokenSwapWatcher struct {
		client       client
		blockHandler blockHandler
	}

	client interface {
		SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
		BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	}

	blockHandler interface {
		Handle(ctx context.Context, block *types.Block) error
	}
)

func NewTokenSwapWatcher(client client, blockHandler blockHandler) *TokenSwapWatcher {
	return &TokenSwapWatcher{client: client, blockHandler: blockHandler}
}

func (tsw TokenSwapWatcher) Run(ctx context.Context) error {
	headers := make(chan *types.Header)
	sub, err := tsw.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case header := <-headers:
			block, err := tsw.client.BlockByNumber(ctx, header.Number)
			if err != nil {
				return err
			}

			ctx = appcontext.WithBlockNumber(ctx, block.Number().Int64())
			if err := tsw.blockHandler.Handle(ctx, block); err != nil {
				return err
			}
		}
	}

}
