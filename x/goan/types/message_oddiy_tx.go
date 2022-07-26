package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOddiyTx = "oddiy_tx"

var _ sdk.Msg = &MsgOddiyTx{}

func NewMsgOddiyTx(creator string, sender string, receiver string, amount string, fee string, feeReceiver string, txType string, serviceName string) *MsgOddiyTx {
	return &MsgOddiyTx{
		Creator:     creator,
		Sender:      sender,
		Receiver:    receiver,
		Amount:      amount,
		Fee:         fee,
		FeeReceiver: feeReceiver,
		TxType:      txType,
		ServiceName: serviceName,
	}
}

func (msg *MsgOddiyTx) Route() string {
	return RouterKey
}

func (msg *MsgOddiyTx) Type() string {
	return TypeMsgOddiyTx
}

func (msg *MsgOddiyTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOddiyTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOddiyTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
