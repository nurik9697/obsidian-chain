package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
        "github.com/cosmos/cosmos-sdk/codec"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDocument{},
		&MsgUpdateDocument{},
		&MsgDeleteDocument{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDocument{}, "obsidian.notary.MsgCreateDocument", nil)
	cdc.RegisterConcrete(&MsgUpdateDocument{}, "obsidian.notary.MsgUpdateDocument", nil)
	cdc.RegisterConcrete(&MsgDeleteDocument{}, "obsidian.notary.MsgDeleteDocument", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "obsidian.notary.MsgUpdateParams", nil)
}

