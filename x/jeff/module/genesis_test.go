package jeff_test

import (
	"testing"

	keepertest "jeff/testutil/keeper"
	"jeff/testutil/nullify"
	jeff "jeff/x/jeff/module"
	"jeff/x/jeff/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JeffKeeper(t)
	jeff.InitGenesis(ctx, k, genesisState)
	got := jeff.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
