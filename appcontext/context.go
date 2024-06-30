package appcontext

import (
	"context"
	"log/slog"
)

type (
	ContextKey int
)

const (
	loggerKey = iota
	blockNumberKey
	transactionKey
	swapIndexKey
)

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func WithBlockNumber(ctx context.Context, blockNumber int64) context.Context {
	return context.WithValue(ctx, blockNumberKey, blockNumber)
}

func WithTransactionHash(ctx context.Context, txn string) context.Context {
	return context.WithValue(ctx, transactionKey, txn)
}

func WithSwapIndex(ctx context.Context, swapIndex int) context.Context {
	return context.WithValue(ctx, swapIndexKey, swapIndex)
}

func Logger(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}

	blockNumber, ok := ctx.Value(blockNumberKey).(int64)
	if ok {
		logger = logger.With(slog.Int64("block_number", blockNumber))
	}

	txn, ok := ctx.Value(transactionKey).(string)
	if ok {
		logger = logger.With(slog.String("txn", txn))
	}

	swapIndex, ok := ctx.Value(swapIndexKey).(int)
	if ok {
		logger = logger.With(slog.Int("swapIndex", swapIndex))
	}

	return logger
}
