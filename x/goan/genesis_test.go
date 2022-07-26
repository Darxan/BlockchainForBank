package goan_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "goan/testutil/keeper"
	"goan/testutil/nullify"
	"goan/x/goan"
	"goan/x/goan/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CustomTxList: []types.CustomTx{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CustomTxCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GoanKeeper(t)
	goan.InitGenesis(ctx, *k, genesisState)
	got := goan.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CustomTxList, got.CustomTxList)
	require.Equal(t, genesisState.CustomTxCount, got.CustomTxCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
