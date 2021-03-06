package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitEncrypted{}, "fairblock/SubmitEncrypted", nil)
	cdc.RegisterConcrete(&MsgCommitDecryption{}, "fairblock/CommitDecryption", nil)
	cdc.RegisterConcrete(&MsgRevealDecryption{}, "fairblock/RevealDecryption", nil)
	cdc.RegisterConcrete(&MsgSubmitShare{}, "fairblock/SubmitShare", nil)
	cdc.RegisterConcrete(&MsgSubmitTarget{}, "fairblock/SubmitTarget", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitEncrypted{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCommitDecryption{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevealDecryption{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitShare{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitTarget{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
