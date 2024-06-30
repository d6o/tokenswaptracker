package handlers

import (
	"context"
	"github.com/d6o/tokenswaptracker/appcontext"
	"github.com/d6o/tokenswaptracker/contract/erc20"
	"github.com/d6o/tokenswaptracker/contract/uniswapv2"
	"github.com/ethereum/go-ethereum/ethclient"
	"log/slog"
)

type (
	Swap struct {
		client *ethclient.Client
	}
)

func NewSwap(client *ethclient.Client) *Swap {
	return &Swap{
		client: client,
	}
}

func (s Swap) Handle(ctx context.Context, swap *uniswapv2.ContractSwap) error {
	logger := appcontext.Logger(ctx)
	logger.Debug("Handling swap")

	caller, err := uniswapv2.NewContractCaller(swap.Raw.Address, s.client)
	if err != nil {
		return err
	}

	token0, err := caller.Token0(nil)
	if err != nil {
		return err
	}

	token1, err := caller.Token1(nil)
	if err != nil {
		return err
	}

	t0Caller, err := erc20.NewContractCaller(token0, s.client)
	if err != nil {
		return err
	}

	t1Caller, err := erc20.NewContractCaller(token1, s.client)
	if err != nil {
		return err
	}

	t0Symbol, err := t0Caller.Symbol(nil)
	if err != nil {
		return err
	}

	t1Symbol, err := t1Caller.Symbol(nil)
	if err != nil {
		return err
	}

	logger.Info("Swap found",
		slog.String("sender", swap.Sender.String()),
		slog.String("to", swap.To.String()),

		slog.String("token0_address", token0.String()),
		slog.String("token1_address", token1.String()),
		slog.String("token0_symbol", t0Symbol),
		slog.String("token1_symbol", t1Symbol),

		slog.String("amount0_in", swap.Amount0In.String()),
		slog.String("amount1_in", swap.Amount1In.String()),

		slog.String("amount0_out", swap.Amount0Out.String()),
		slog.String("amount1_out", swap.Amount1Out.String()),
	)

	return nil
}
