package main_test

import (
	"context"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TestChainStart spins up a gaia chain and waits for 5 blocks.
func TestChainStart(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	t.Parallel()

	const (
		denom = "uatom"
	)
	validators := 1
	fullNodes := 0
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          "gaia",
			ChainName:     "gaia",
			Version:       "v15.2.0",
			NumValidators: &validators,
			NumFullNodes:  &fullNodes,
			ChainConfig: ibc.ChainConfig{
				Denom: denom,
			},
		},
	})

	coin := sdk.NewCoin(denom, sdk.NewInt(1000000000))
	t.Log(coin.String())

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	client, network := interchaintest.DockerSetup(t)

	chain := chains[0].(*cosmos.CosmosChain)

	ic := interchaintest.NewInterchain().
		AddChain(chain)
	rep := testreporter.NewNopReporter()

	ctx := context.Background()

	err = ic.Build(ctx, rep.RelayerExecReporter(t), interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,
		// BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
		SkipPathCreation: false,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = ic.Close()
	})

	// wait for 5 blocks
	testutil.WaitForBlocks(ctx, 5, chain)
}
