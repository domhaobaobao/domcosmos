package keeper_test

import (
	"testing"

	testkeeper "domcosmos/testutil/keeper"
	"domcosmos/x/domcosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DomcosmosKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
