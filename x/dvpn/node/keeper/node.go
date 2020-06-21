package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

func (k Keeper) SetNode(ctx sdk.Context, node types.Node) {
	key := types.NodeKey(node.Address)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(node)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetNode(ctx sdk.Context, address hub.NodeAddress) (node types.Node, found bool) {
	store := k.Store(ctx)

	key := types.NodeKey(address)
	value := store.Get(key)
	if value == nil {
		return node, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &node)
	return node, true
}

func (k Keeper) SetNodeAddressForProvider(ctx sdk.Context, pa hub.ProvAddress, na hub.NodeAddress) {
	key := types.NodeAddressForProviderKey(pa, na)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(na)

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k Keeper) GetNodes(ctx sdk.Context) (nodes types.Nodes) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var node types.Node
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &node)
		nodes = append(nodes, node)
	}

	return nodes
}

func (k Keeper) GetNodesOfProvider(ctx sdk.Context, address hub.ProvAddress) (nodes types.Nodes) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.NodeAddressForProviderKeyPrefix(address))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var na hub.NodeAddress
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &na)

		node, _ := k.GetNode(ctx, na)
		nodes = append(nodes, node)
	}

	return nodes
}

func (k Keeper) IterateNodes(ctx sdk.Context, f func(i int, node types.Node) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var node types.Node
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &node)

		if stop := f(i, node); stop {
			break
		}
		i++
	}
}
