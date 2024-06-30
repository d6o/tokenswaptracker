package handlers

import (
	"context"
	"github.com/d6o/tokenswaptracker/appcontext"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	Block struct {
		client             client
		transactionHandler transactionHandler
	}

	transactionHandler interface {
		Handle(ctx context.Context, receipt *types.Receipt) error
	}

	client interface {
		TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	}
)

func NewBlock(client client, transactionHandler transactionHandler) *Block {
	return &Block{client: client, transactionHandler: transactionHandler}
}

func (b Block) Handle(ctx context.Context, block *types.Block) error {
	logger := appcontext.Logger(ctx)
	logger.Info("Found new block")

	for _, tx := range block.Transactions() {
		ctx = appcontext.WithTransactionHash(ctx, tx.Hash().String())

		receipt, err := b.client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return err
		}

		if err := b.transactionHandler.Handle(ctx, receipt); err != nil {
			return err
		}
	}

	return nil
}
