package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmos-talents/cosmos-talents/testutil/keeper"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/keeper"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmostalentsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
