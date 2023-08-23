package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/cosmos-talents/cosmos-talents/testutil/keeper"
	"github.com/cosmos-talents/cosmos-talents/testutil/nullify"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/keeper"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/types"
)

func createTestInfo(keeper *keeper.Keeper, ctx sdk.Context) types.Info {
	item := types.Info{}
	keeper.SetInfo(ctx, item)
	return item
}

func TestInfoGet(t *testing.T) {
	keeper, ctx := keepertest.CosmostalentsKeeper(t)
	item := createTestInfo(keeper, ctx)
	rst, found := keeper.GetInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.CosmostalentsKeeper(t)
	createTestInfo(keeper, ctx)
	keeper.RemoveInfo(ctx)
	_, found := keeper.GetInfo(ctx)
	require.False(t, found)
}
