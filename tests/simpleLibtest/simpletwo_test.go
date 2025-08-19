package simpletest

import (
	"testing"
	. "school/project/cmd"

	"github.com/stretchr/testify/assert"
)

// FSM: StartTwo -> MM1(bool) -> MM2|MM3 depending on result
func TestGenerateTSS_SimpleTwo(t *testing.T) {
	fsm := FSM{
		Labels: SetFromLabels("MM10", "MM20", "MM30"),
		States: SetFromStates( "SS10", "SS20"),
		S0:     "SS00",
		F:      "EE00",
		D: Transition{
			{"SS00", "MM10"}: []Pair{
				{Outcome: true, TransitionState: "SS10"},
				{Outcome: false, TransitionState: "SS20"},
			},
			{"SS10", "MM20"}: []Pair{
				{TransitionState: "SS00"},
			},
			{"SS20", "MM30"}: []Pair{
				{TransitionState: "EE00"},
			},
		},
	}

	rawLogic := RawTransition{
		{State: "SS00", Label: "MM10"}: {
			Name:    "MM1",
			Inputs:  []string{"choice int"},
			Outputs: []string{"bool"},
		},
		{State: "SS10", Label: "MM20"}: {
			Name:    "MM2",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
		{State: "SS20", Label: "MM30"}: {
			Name:    "MM3",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "simplelib", "SimpleStructTwo")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "school/project/simplelib")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/MM10_to_SS10_or_SS20.go")
	if err != nil {
		t.Fatalf("Failed to generated api file\n");
	}
}
