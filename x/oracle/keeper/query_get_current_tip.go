package keeper

import (
	"context"

	"github.com/tellor-io/layer/utils"
	"github.com/tellor-io/layer/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q Querier) GetCurrentTip(ctx context.Context, req *types.QueryGetCurrentTipRequest) (*types.QueryGetCurrentTipResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	qdata, err := utils.QueryBytesFromString(req.QueryData)
	if err != nil {
		return nil, err
	}

	queryId := utils.QueryIDFromData(qdata)

	tips, err := q.keeper.GetQueryTip(ctx, queryId)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetCurrentTipResponse{
		Tips: &types.Tips{
			QueryData: qdata, // TODO: avoid returning the same data as the request
			Amount:    tips,
		},
	}, nil
}
