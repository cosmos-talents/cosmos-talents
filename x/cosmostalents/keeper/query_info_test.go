package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/cosmos-talents/cosmos-talents/testutil/keeper"
	"github.com/cosmos-talents/cosmos-talents/testutil/nullify"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/types"
)

func TestInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.CosmostalentsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestInfo(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetInfoRequest
		response *types.QueryGetInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetInfoRequest{},
			response: &types.QueryGetInfoResponse{Info: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Info(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
