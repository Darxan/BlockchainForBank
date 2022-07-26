package goan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"goan/x/goan/keeper"
	"goan/x/goan/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the customTx
	for _, elem := range genState.CustomTxList {
		k.SetCustomTx(ctx, elem)
	}

	// Set customTx count
	k.SetCustomTxCount(ctx, genState.CustomTxCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CustomTxList = k.GetAllCustomTx(ctx)
	genesis.CustomTxCount = k.GetCustomTxCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
