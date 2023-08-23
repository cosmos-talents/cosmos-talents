package keeper

import (
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetInfo set info in the store
func (k Keeper) SetInfo(ctx sdk.Context, info types.Info) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InfoKey))
	b := k.cdc.MustMarshal(&info)
	store.Set([]byte{0}, b)
}

// GetInfo returns info
func (k Keeper) GetInfo(ctx sdk.Context) (val types.Info, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInfo removes info from the store
func (k Keeper) RemoveInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InfoKey))
	store.Delete([]byte{0})
}
