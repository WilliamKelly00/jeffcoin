package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "jeff/testutil/keeper"
	"jeff/x/jeff/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.JeffKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
