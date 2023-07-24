package cosmostalents_test

import (
	"testing"

	keepertest "github.com/cosmos-talents/cosmos-talents/testutil/keeper"
	"github.com/cosmos-talents/cosmos-talents/testutil/nullify"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmostalentsKeeper(t)
	cosmostalents.InitGenesis(ctx, *k, genesisState)
	got := cosmostalents.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
