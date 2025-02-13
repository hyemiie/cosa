package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAuction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveAuction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatBid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCloseAuction{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
