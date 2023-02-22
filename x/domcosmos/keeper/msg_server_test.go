package keeper_test

import (
	"context"
	"testing"

	keepertest "domcosmos/testutil/keeper"
	"domcosmos/x/domcosmos/keeper"
	"domcosmos/x/domcosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DomcosmosKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
