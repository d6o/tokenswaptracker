package handlers

import (
	"context"
	"log/slog"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/d6o/tokenswaptracker/appcontext"
	"github.com/d6o/tokenswaptracker/contract/uniswapv2"
)

type (
	Transaction struct {
		swapParser  swapParser
		swapHandler swapHandler
	}

	swapParser interface {
		ParseSwap(log types.Log) (*uniswapv2.ContractSwap, error)
	}

	swapHandler interface {
		Handle(ctx context.Context, swap *uniswapv2.ContractSwap) error
	}
)

func NewTransaction(swapParser swapParser, swapHandler swapHandler) *Transaction {
	return &Transaction{swapParser: swapParser, swapHandler: swapHandler}
}

func (t Transaction) Handle(ctx context.Context, receipt *types.Receipt) error {
	appcontext.Logger(ctx).Debug("Found new transaction")

	for i, log := range receipt.Logs {
		ctx = appcontext.WithSwapIndex(ctx, i)

		swapEvent, err := t.swapParser.ParseSwap(*log)
		if err != nil {
			appcontext.Logger(ctx).Debug("Log isn't a swap event", slog.Any("error", err))
			continue
		}

		if err := t.swapHandler.Handle(ctx, swapEvent); err != nil {
			return err
		}
	}

	return nil
}
