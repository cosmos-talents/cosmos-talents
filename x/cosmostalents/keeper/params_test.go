package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmos-talents/cosmos-talents/testutil/keeper"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmostalentsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
