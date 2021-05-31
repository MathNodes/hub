package simulation

import (
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func getRandomInactiveDuration(r *rand.Rand) time.Duration {
	return time.Duration(sdksimulation.RandIntBetween(r, 60, 60<<13))
}

func getRandomProofVerificationEnabled(r *rand.Rand) bool {
	return rand.Int31n(2) == int32(0)
}

func RandomizedGenState(simState *module.SimulationState) *types.GenesisState {

	var (
		inactiveDuration          time.Duration
		proofVerificationEnabled bool
	)

	inactiveDurationSim := func(r *rand.Rand) {
		inactiveDuration = getRandomInactiveDuration(r)
	}

	proofVerificationEnabledSim := func(r *rand.Rand) {
		proofVerificationEnabled = getRandomProofVerificationEnabled(r)
	}

	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyInactiveDuration), inactiveDuration, nil, inactiveDurationSim)
	simState.AppParams.GetOrGenerate(simState.Cdc, string(types.KeyProofVerificationEnabled), proofVerificationEnabled, nil, proofVerificationEnabledSim)

	params := types.NewParams(inactiveDuration, proofVerificationEnabled)
	sessions := getRandomSessions(simState.Rand)

	state := types.NewGenesisState(sessions, params)
	bz := simState.Cdc.MustMarshalJSON(&state.Params)

	fmt.Printf("selected randomly generated sessions parameters: %s\n", bz)
	return state
}

func getRandomSessions(r *rand.Rand) types.Sessions {
	sessions := make(types.Sessions, r.Int31n(28)+4)

	for range sessions {
		sessions = append(sessions, types.Session{
			Id:           rand.Uint64(),
			Subscription: 0,
			Node:         hubtypes.NodeAddress("node-address").String(),
			Address:      sdk.AccAddress("address").String(),
			Duration:     time.Duration(r.Int31n(1800)+1800),
			Bandwidth:    hubtypes.NewBandwidthFromInt64(int64(1024<<20 * r.Intn(10)), int64(1024<<20 * r.Intn(10))),
			Status:       hubtypes.Status(r.Int31n(3)),
			StatusAt:     sdksimulation.RandTimestamp(r),
		})
	}

	return sessions
}
