package provider

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v1/x/provider/keeper"
	"github.com/sentinel-official/hub/v1/x/provider/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Providers {
		k.SetProvider(ctx, item)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetProviders(ctx),
		k.GetParams(ctx),
	)
}
