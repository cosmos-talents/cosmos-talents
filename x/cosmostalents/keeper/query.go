package keeper

import (
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/types"
)

var _ types.QueryServer = Keeper{}
