package fairblock

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/pememoni/FairBlock/testutil/sample"
	fairblocksimulation "github.com/pememoni/FairBlock/x/fairblock/simulation"
	"github.com/pememoni/FairBlock/x/fairblock/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = fairblocksimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSubmitEncrypted = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitEncrypted int = 100

	opWeightMsgCommitDecryption = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCommitDecryption int = 100

	opWeightMsgRevealDecryption = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevealDecryption int = 100

	opWeightMsgSubmitShare = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitShare int = 100

	opWeightMsgSubmitTarget = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitTarget int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	fairblockGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&fairblockGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitEncrypted int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitEncrypted, &weightMsgSubmitEncrypted, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitEncrypted = defaultWeightMsgSubmitEncrypted
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitEncrypted,
		fairblocksimulation.SimulateMsgSubmitEncrypted(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCommitDecryption int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCommitDecryption, &weightMsgCommitDecryption, nil,
		func(_ *rand.Rand) {
			weightMsgCommitDecryption = defaultWeightMsgCommitDecryption
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCommitDecryption,
		fairblocksimulation.SimulateMsgCommitDecryption(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRevealDecryption int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevealDecryption, &weightMsgRevealDecryption, nil,
		func(_ *rand.Rand) {
			weightMsgRevealDecryption = defaultWeightMsgRevealDecryption
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevealDecryption,
		fairblocksimulation.SimulateMsgRevealDecryption(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitShare int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitShare, &weightMsgSubmitShare, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitShare = defaultWeightMsgSubmitShare
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitShare,
		fairblocksimulation.SimulateMsgSubmitShare(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitTarget int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitTarget, &weightMsgSubmitTarget, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitTarget = defaultWeightMsgSubmitTarget
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitTarget,
		fairblocksimulation.SimulateMsgSubmitTarget(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
