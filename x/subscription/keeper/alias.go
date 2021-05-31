package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	hubtypes "github.com/sentinel-official/hub/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
)

func (k *Keeper) GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI {
	return k.account.GetAccount(ctx, address)
}

func (k *Keeper) SendCoin(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, coin sdk.Coin) error {
	return k.bank.SendCoins(ctx, from, to, sdk.NewCoins(coin))
}

func (k *Keeper) AddDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) error {
	return k.deposit.Add(ctx, address, sdk.NewCoins(coin))
}

func (k *Keeper) SubtractDeposit(ctx sdk.Context, address sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.Subtract(ctx, address, sdk.NewCoins(coin))
}

func (k *Keeper) GetNode(ctx sdk.Context, address hubtypes.NodeAddress) (nodetypes.Node, bool) {
	return k.node.GetNode(ctx, address)
}

func (k *Keeper) GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool) {
	return k.plan.GetPlan(ctx, id)
}
