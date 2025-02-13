package cosa

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"cosa/testutil/sample"
	cosasimulation "cosa/x/cosa/simulation"
	"cosa/x/cosa/types"
)

// avoid unused import issue
var (
	_ = cosasimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateAuction = "op_weight_msg_create_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAuction int = 100

	opWeightMsgApproveAuction = "op_weight_msg_approve_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveAuction int = 100

	opWeightMsgCreatBid = "op_weight_msg_creat_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatBid int = 100

	opWeightMsgCloseAuction = "op_weight_msg_close_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCloseAuction int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cosaGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cosaGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAuction int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateAuction, &weightMsgCreateAuction, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAuction = defaultWeightMsgCreateAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAuction,
		cosasimulation.SimulateMsgCreateAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveAuction int
	simState.AppParams.GetOrGenerate(opWeightMsgApproveAuction, &weightMsgApproveAuction, nil,
		func(_ *rand.Rand) {
			weightMsgApproveAuction = defaultWeightMsgApproveAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveAuction,
		cosasimulation.SimulateMsgApproveAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatBid int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatBid, &weightMsgCreatBid, nil,
		func(_ *rand.Rand) {
			weightMsgCreatBid = defaultWeightMsgCreatBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatBid,
		cosasimulation.SimulateMsgCreatBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCloseAuction int
	simState.AppParams.GetOrGenerate(opWeightMsgCloseAuction, &weightMsgCloseAuction, nil,
		func(_ *rand.Rand) {
			weightMsgCloseAuction = defaultWeightMsgCloseAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCloseAuction,
		cosasimulation.SimulateMsgCloseAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAuction,
			defaultWeightMsgCreateAuction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cosasimulation.SimulateMsgCreateAuction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApproveAuction,
			defaultWeightMsgApproveAuction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cosasimulation.SimulateMsgApproveAuction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatBid,
			defaultWeightMsgCreatBid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cosasimulation.SimulateMsgCreatBid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCloseAuction,
			defaultWeightMsgCloseAuction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cosasimulation.SimulateMsgCloseAuction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
