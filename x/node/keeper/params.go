package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) Deposit(ctx sdk.Context) (v sdk.Coin) {
	k.params.Get(ctx, types.KeyDeposit, &v)
	return
}

func (k *Keeper) InactiveDuration(ctx sdk.Context) (v time.Duration) {
	k.params.Get(ctx, types.KeyInactiveDuration, &v)
	return
}

func (k *Keeper) MaxGigabytePrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMaxGigabytePrices, &v)
	return
}

func (k *Keeper) MaxHourlyPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMaxHourlyPrices, &v)
	return
}

func (k *Keeper) MinGigabytePrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMinGigabytePrices, &v)
	return
}

func (k *Keeper) MinHourlyPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, types.KeyMinHourlyPrices, &v)
	return
}

func (k *Keeper) RevenueShare(ctx sdk.Context) (v sdk.Dec) {
	k.params.Get(ctx, types.KeyRevenueShare, &v)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Deposit(ctx),
		k.InactiveDuration(ctx),
		k.MaxGigabytePrices(ctx),
		k.MaxHourlyPrices(ctx),
		k.MinGigabytePrices(ctx),
		k.MinHourlyPrices(ctx),
		k.RevenueShare(ctx),
	)
}

func (k *Keeper) IsValidGigabytePrices(ctx sdk.Context, prices sdk.Coins) bool {
	maxPrice := k.MaxGigabytePrices(ctx)
	for _, coin := range maxPrice {
		amount := prices.AmountOf(coin.Denom)
		if amount.GT(coin.Amount) {
			return false
		}
	}

	minPrice := k.MinGigabytePrices(ctx)
	for _, coin := range minPrice {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsValidHourlyPrices(ctx sdk.Context, prices sdk.Coins) bool {
	maxPrice := k.MaxHourlyPrices(ctx)
	for _, coin := range maxPrice {
		amount := prices.AmountOf(coin.Denom)
		if amount.GT(coin.Amount) {
			return false
		}
	}

	minPrice := k.MinHourlyPrices(ctx)
	for _, coin := range minPrice {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsMaxGigabytePricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMaxGigabytePrices)
}

func (k *Keeper) IsMaxHourlyPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMaxHourlyPrices)
}

func (k *Keeper) IsMinGigabytePricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMinGigabytePrices)
}

func (k *Keeper) IsMinHourlyPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, types.KeyMinHourlyPrices)
}
