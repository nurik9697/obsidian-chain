package notary

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"obsidian/testutil/sample"
	notarysimulation "obsidian/x/notary/simulation"
	"obsidian/x/notary/types"
)

// avoid unused import issue
var (
	_ = notarysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateDocument = "op_weight_msg_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDocument int = 100

	opWeightMsgUpdateDocument = "op_weight_msg_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDocument int = 100

	opWeightMsgDeleteDocument = "op_weight_msg_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDocument int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	notaryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DocumentList: []types.Document{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&notaryGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDocument int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateDocument, &weightMsgCreateDocument, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDocument = defaultWeightMsgCreateDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDocument,
		notarysimulation.SimulateMsgCreateDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDocument int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateDocument, &weightMsgUpdateDocument, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDocument = defaultWeightMsgUpdateDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDocument,
		notarysimulation.SimulateMsgUpdateDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDocument int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteDocument, &weightMsgDeleteDocument, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDocument = defaultWeightMsgDeleteDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDocument,
		notarysimulation.SimulateMsgDeleteDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateDocument,
			defaultWeightMsgCreateDocument,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				notarysimulation.SimulateMsgCreateDocument(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateDocument,
			defaultWeightMsgUpdateDocument,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				notarysimulation.SimulateMsgUpdateDocument(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteDocument,
			defaultWeightMsgDeleteDocument,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				notarysimulation.SimulateMsgDeleteDocument(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
