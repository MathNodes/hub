package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/deposit/types"
)

// SetDeposit stores a deposit in the module's KVStore.
func (k *Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	var (
		store = k.Store(ctx)
		key   = types.DepositKey(deposit.GetAddress())
		value = k.cdc.MustMarshal(&deposit)
	)

	store.Set(key, value)
}

// GetDeposit retrieves a deposit from the module's KVStore based on the account address.
// If the deposit exists, it returns the deposit and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetDeposit(ctx sdk.Context, addr sdk.AccAddress) (deposit types.Deposit, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.DepositKey(addr)
		value = store.Get(key)
	)

	if value == nil {
		return deposit, false
	}

	k.cdc.MustUnmarshal(value, &deposit)
	return deposit, true
}

// GetDeposits retrieves all deposits stored in the module's KVStore.
func (k *Keeper) GetDeposits(ctx sdk.Context) (items types.Deposits) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item types.Deposit
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

// IterateDeposits iterates over all deposits stored in the module's KVStore and calls the provided function for each deposit.
// The iteration stops when the provided function returns 'true'.
func (k *Keeper) IterateDeposits(ctx sdk.Context, fn func(index int, item types.Deposit) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.DepositKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var item types.Deposit
		k.cdc.MustUnmarshal(iter.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

// SendCoinsFromAccountToDeposit transfers coins from an account to a deposit and updates the deposit in the KVStore.
// If the deposit does not exist, a new deposit will be created.
// It returns an error if the account doesn't have enough balance.
func (k *Keeper) SendCoinsFromAccountToDeposit(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error {
	if err := k.bank.SendCoinsFromAccountToModule(ctx, fromAddr, types.ModuleName, coins); err != nil {
		return err
	}

	deposit, found := k.GetDeposit(ctx, toAddr)
	if !found {
		deposit = types.Deposit{
			Address: toAddr.String(),
			Coins:   sdk.NewCoins(),
		}
	}

	deposit.Coins = deposit.Coins.Add(coins...)
	if deposit.Coins.IsAnyNegative() {
		return types.NewErrorInsufficientFunds(fromAddr, deposit.Coins)
	}

	k.SetDeposit(ctx, deposit)
	ctx.EventManager().EmitTypedEvent(
		&types.EventAdd{
			Address: toAddr.String(),
			Coins:   coins,
		},
	)

	return nil
}

// SendCoinsFromDepositToAccount transfers coins from a deposit to an account and updates the deposit in the KVStore.
// It returns an error if the deposit doesn't have enough balance.
func (k *Keeper) SendCoinsFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error {
	deposit, found := k.GetDeposit(ctx, fromAddr)
	if !found {
		return types.NewErrorDepositNotFound(fromAddr)
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.NewErrorInsufficientFunds(fromAddr, coins)
	}

	if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, coins); err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSubtract{
			Address: fromAddr.String(),
			Coins:   coins,
		},
	)

	return nil
}

// SendCoinsFromDepositToModule transfers coins from a deposit to a module and updates the deposit in the KVStore.
// It returns an error if the deposit doesn't have enough balance.
func (k *Keeper) SendCoinsFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coins sdk.Coins) error {
	deposit, found := k.GetDeposit(ctx, fromAddr)
	if !found {
		return types.NewErrorDepositNotFound(fromAddr)
	}

	deposit.Coins, _ = deposit.Coins.SafeSub(coins)
	if deposit.Coins.IsAnyNegative() {
		return types.NewErrorInsufficientFunds(fromAddr, coins)
	}

	if err := k.bank.SendCoinsFromModuleToModule(ctx, types.ModuleName, toModule, coins); err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)
	ctx.EventManager().EmitTypedEvent(
		&types.EventSubtract{
			Address: fromAddr.String(),
			Coins:   coins,
		},
	)

	return nil
}

// Add is a utility function to add coins to a deposit by transferring from the account to the deposit.
func (k *Keeper) Add(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
	return k.SendCoinsFromAccountToDeposit(ctx, addr, addr, coins)
}

// Subtract is a utility function to subtract coins from a deposit by transferring from the deposit to the account.
func (k *Keeper) Subtract(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
	return k.SendCoinsFromDepositToAccount(ctx, addr, addr, coins)
}
