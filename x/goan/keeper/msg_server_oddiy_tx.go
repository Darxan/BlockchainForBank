package keeper

import (
	"context"
	// fmt "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"goan/x/goan/types"
)

func (k msgServer) OddiyTx(goCtx context.Context, msg *types.MsgOddiyTx) (*types.MsgOddiyTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, _ := sdk.AccAddressFromBech32(msg.Creator)
	receiver, _ := sdk.AccAddressFromBech32(msg.Receiver)
	feeReceiver, _ :=  sdk.AccAddressFromBech32(msg.FeeReceiver)
	amount, _ := sdk.ParseCoinsNormalized(msg.Amount)
	fee, _ := sdk.ParseCoinsNormalized(msg.Fee)

	k.bankKeeper.SendCoins(ctx, sender, receiver, amount)

	k.bankKeeper.SendCoins(ctx, sender, feeReceiver, fee)
	var newData = types.CustomTx {
		Sender: sender.String(),
		Receiver: receiver.String(), 
		Amount: msg.Amount,
		Fee: msg.Fee,
		FeeReceiver: msg.FeeReceiver,
		TxType: msg.TxType,
		ServiceName: msg.ServiceName,
	}

	k.AppendCustomTx(ctx, newData)
	return &types.MsgOddiyTxResponse{}, nil
}
