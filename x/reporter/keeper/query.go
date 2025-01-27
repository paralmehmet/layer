package keeper

import (
	"context"

	"github.com/tellor-io/layer/x/reporter/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

// Reporters queries all the reporters
func (k Querier) Reporters(ctx context.Context, req *types.QueryReportersRequest) (*types.QueryReportersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	repStore := prefix.NewStore(store, types.ReportersKey)
	reporters := make([]*types.Reporter, 0)
	pageRes, err := query.Paginate(repStore, req.Pagination, func(repAddr, value []byte) error {
		var reporterMeta types.OracleReporter
		err := k.cdc.Unmarshal(value, &reporterMeta)
		if err != nil {
			return err
		}

		reporters = append(reporters, &types.Reporter{
			Address:  sdk.AccAddress(repAddr).String(),
			Metadata: &reporterMeta,
		})
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryReportersResponse{Reporters: reporters, Pagination: pageRes}, nil
}

// SelectorReporter queries the reporter of a selector
func (k Querier) SelectorReporter(ctx context.Context, req *types.QuerySelectorReporterRequest) (*types.QuerySelectorReporterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	delAddr := sdk.MustAccAddressFromBech32(req.SelectorAddress)

	selector, err := k.Keeper.Selectors.Get(ctx, delAddr)
	if err != nil {
		return nil, err
	}

	return &types.QuerySelectorReporterResponse{Reporter: sdk.AccAddress(selector.GetReporter()).String()}, nil
}
