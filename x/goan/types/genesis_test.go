package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"goan/x/goan/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				CustomTxList: []types.CustomTx{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				CustomTxCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated customTx",
			genState: &types.GenesisState{
				CustomTxList: []types.CustomTx{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid customTx count",
			genState: &types.GenesisState{
				CustomTxList: []types.CustomTx{
					{
						Id: 1,
					},
				},
				CustomTxCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
