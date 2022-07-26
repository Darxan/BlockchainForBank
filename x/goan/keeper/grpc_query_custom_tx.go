package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"goan/x/goan/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CustomTxAll(c context.Context, req *types.QueryAllCustomTxRequest) (*types.QueryAllCustomTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var customTxs []types.CustomTx
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	customTxStore := prefix.NewStore(store, types.KeyPrefix(types.CustomTxKey))

	pageRes, err := query.Paginate(customTxStore, req.Pagination, func(key []byte, value []byte) error {
		var customTx types.CustomTx
		if err := k.cdc.Unmarshal(value, &customTx); err != nil {
			return err
		}

		customTxs = append(customTxs, customTx)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCustomTxResponse{CustomTx: customTxs, Pagination: pageRes}, nil
}

func (k Keeper) CustomTx(c context.Context, req *types.QueryGetCustomTxRequest) (*types.QueryGetCustomTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	customTx, found := k.GetCustomTx(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCustomTxResponse{CustomTx: customTx}, nil
}
