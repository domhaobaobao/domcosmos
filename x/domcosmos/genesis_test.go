package domcosmos_test

import (
	"testing"

	keepertest "domcosmos/testutil/keeper"
	"domcosmos/testutil/nullify"
	"domcosmos/x/domcosmos"
	"domcosmos/x/domcosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DomcosmosKeeper(t)
	domcosmos.InitGenesis(ctx, *k, genesisState)
	got := domcosmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
