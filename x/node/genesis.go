package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v1/types"
	"github.com/sentinel-official/hub/v1/x/node/keeper"
	"github.com/sentinel-official/hub/v1/x/node/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Nodes {
		k.SetNode(ctx, item)
		if item.Status.Equal(hubtypes.StatusActive) {
			addr := item.GetAddress()
			k.SetNodeForInactiveAt(ctx, item.InactiveAt, addr)
		}
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetNodes(ctx),
		k.GetParams(ctx),
	)
}
