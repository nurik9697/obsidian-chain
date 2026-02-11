package notary_test

import (
	"testing"

	keepertest "obsidian/testutil/keeper"
	"obsidian/testutil/nullify"
	notary "obsidian/x/notary/module"
	"obsidian/x/notary/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DocumentList: []types.Document{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NotaryKeeper(t)
	notary.InitGenesis(ctx, k, genesisState)
	got := notary.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DocumentList, got.DocumentList)
	// this line is used by starport scaffolding # genesis/test/assert
}
