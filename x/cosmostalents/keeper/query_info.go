package keeper

import (
	"context"

	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Info(goCtx context.Context, req *types.QueryGetInfoRequest) (*types.QueryGetInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInfoResponse{Info: val}, nil
}
