package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "goan/testutil/keeper"
	"goan/x/goan/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GoanKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
