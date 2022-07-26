package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "goan/testutil/keeper"
	"goan/testutil/nullify"
	"goan/x/goan/keeper"
	"goan/x/goan/types"
)

func createNCustomTx(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CustomTx {
	items := make([]types.CustomTx, n)
	for i := range items {
		items[i].Id = keeper.AppendCustomTx(ctx, items[i])
	}
	return items
}

func TestCustomTxGet(t *testing.T) {
	keeper, ctx := keepertest.GoanKeeper(t)
	items := createNCustomTx(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCustomTx(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCustomTxRemove(t *testing.T) {
	keeper, ctx := keepertest.GoanKeeper(t)
	items := createNCustomTx(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCustomTx(ctx, item.Id)
		_, found := keeper.GetCustomTx(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCustomTxGetAll(t *testing.T) {
	keeper, ctx := keepertest.GoanKeeper(t)
	items := createNCustomTx(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCustomTx(ctx)),
	)
}

func TestCustomTxCount(t *testing.T) {
	keeper, ctx := keepertest.GoanKeeper(t)
	items := createNCustomTx(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCustomTxCount(ctx))
}
