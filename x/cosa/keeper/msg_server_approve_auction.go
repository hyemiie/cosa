package keeper

import (
	"context"
	"time"

	"cosa/x/cosa/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (k msgServer) ApproveAuction(goCtx context.Context, msg *types.MsgApproveAuction) (*types.MsgApproveAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction, found := k.GetAuction(ctx, msg.Id)
	if !found {
		return &types.MsgApproveAuctionResponse{}, sdkerrors.Wrapf(sdkerror.ErrKeyNotFound, "auction %d doesnt exist", auction.Id)
	}

	if auction.Creator != msg.Creator {
		return &types.MsgApproveAuctionResponse{}, sdkerrors.Wrap(sdkerror.ErrUnauthorized, "cannot change bid creator")
	}

	endTime := ctx.BlockTime().Add(time.Duration(auction.Duration) * time.Second)

	auction.Status = Approved
	auction.Endtime = timestamppb.New(endTime)

	k.SetAuction(ctx, auction)

	return &types.MsgApproveAuctionResponse{
		Status: auction.Status,
	}, nil
}
