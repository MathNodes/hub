package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hubtypes "github.com/sentinel-official/hub/v1/types"
	"github.com/sentinel-official/hub/v1/x/node/types"
)

var (
	_ types.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	Keeper
}

func NewQueryServiceServer(k Keeper) types.QueryServiceServer {
	return &queryServer{k}
}

func (q *queryServer) QueryNode(c context.Context, req *types.QueryNodeRequest) (*types.QueryNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := hubtypes.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetNode(ctx, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "node does not exist for address %s", req.Address)
	}

	return &types.QueryNodeResponse{Node: item}, nil
}

func (q *queryServer) QueryNodes(c context.Context, req *types.QueryNodesRequest) (res *types.QueryNodesResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items     types.Nodes
		keyPrefix []byte
		ctx       = sdk.UnwrapSDKContext(c)
	)

	switch req.Status {
	case hubtypes.StatusActive:
		keyPrefix = types.ActiveNodeKeyPrefix
	case hubtypes.StatusInactive:
		keyPrefix = types.InactiveNodeKeyPrefix
	default:
		keyPrefix = types.NodeKeyPrefix
	}

	store := prefix.NewStore(q.Store(ctx), keyPrefix)
	pagination, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item types.Node
		if err := q.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryNodesResponse{Nodes: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryNodesForPlan(c context.Context, req *types.QueryNodesForPlanRequest) (*types.QueryNodesForPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items types.Nodes
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetNodeForPlanKeyPrefix(req.Id))
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		if !accumulate {
			return false, nil
		}

		item, found := q.GetNode(ctx, key[1:])
		if !found {
			return false, fmt.Errorf("node for key %X does not exist", key)
		}

		if req.Status.Equal(hubtypes.StatusUnspecified) || item.Status.Equal(req.Status) {
			items = append(items, item)
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryNodesForPlanResponse{Nodes: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryParams(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = q.GetParams(ctx)
	)

	return &types.QueryParamsResponse{Params: params}, nil
}
