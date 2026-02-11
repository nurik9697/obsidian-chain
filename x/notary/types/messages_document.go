package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDocument{}

func NewMsgCreateDocument(
	creator string,
	index string,
	fileHash string,
	owner string,
	timestamp int32,

) *MsgCreateDocument {
	return &MsgCreateDocument{
		Creator:   creator,
		Index:     index,
		FileHash:  fileHash,
		Owner:     owner,
		Timestamp: timestamp,
	}
}

func (msg *MsgCreateDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDocument{}

func NewMsgUpdateDocument(
	creator string,
	index string,
	fileHash string,
	owner string,
	timestamp int32,

) *MsgUpdateDocument {
	return &MsgUpdateDocument{
		Creator:   creator,
		Index:     index,
		FileHash:  fileHash,
		Owner:     owner,
		Timestamp: timestamp,
	}
}

func (msg *MsgUpdateDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDocument{}

func NewMsgDeleteDocument(
	creator string,
	index string,

) *MsgDeleteDocument {
	return &MsgDeleteDocument{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
