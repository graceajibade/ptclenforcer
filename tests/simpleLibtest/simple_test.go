package simpletest

import (
	"testing"
	. "school/project/cmd"

	"github.com/stretchr/testify/assert"
)

func SetFromStates(states ...string) Set[State] {
	s := make(Set[State])
	for _, st := range states {
		s[State(st)] = struct{}{}
	}
	return s
}

func SetFromLabels(labels ...string) Set[Label] {
	s := make(Set[Label])
	for _, l := range labels {
		s[Label(l)] = struct{}{}
	}
	return s
}

//M0S1.End
func TestGenerateTSSM1_Simple(t *testing.T) {
	fsm := FSM{
		Labels: SetFromLabels("M01"),
		States: SetFromStates("S00", "E00"),
		S0:     "S00",
		F:      "E00",
		D: Transition{
			{"S00", "M01"}: []Pair{
				{TransitionState: "E00"},
			},
		},
	}

	rawLogic := RawTransition{
		{State: "S00", Label: "M01"}: {
			Name:    "M1",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "simplelib", "SimpleStruct")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "school/project/simplelib")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/M1_end.go")
	if err != nil {
		t.Fatalf("Failed to generated api file\n");
	}
}

//M11.M11.End
func TestGenerateTSSM2_Simple(t *testing.T) {
	fsm := FSM{
		Labels: SetFromLabels("M11"),
		States: SetFromStates("S01", "S11", "E01"),
		S0:     "S01",
		F:      "E01",
		D: Transition{
			{"S01", "M11"}: []Pair{
				{TransitionState: "S11"},
			},
			{"S11", "M11"}: []Pair{
				{TransitionState: "E01"},
			},
		},
	}

	rawLogic := RawTransition{
		{State: "S01", Label: "M11"}: {
			Name:    "M1",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
		{State: "S11", Label: "M11"}: {
			Name:    "M1",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "simplelib", "SimpleStruct")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "school/project/simplelib")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/M11_M11_end.go")
	if err != nil {
		t.Fatalf("Failed to generated api file\n");
	}
}

// M12.M21.End
func TestGenerateTSSM12_Simple(t *testing.T) {
	fsm := FSM{
		Labels: SetFromLabels("M12", "M21"),
		States: SetFromStates("S02", "S12", "E02"),
		S0:     "S02",
		F:      "E02",
		D: Transition{
			{"S02", "M12"}: []Pair{
				{TransitionState: "S12"},
			},
			{"S12", "M21"}: []Pair{
				{TransitionState: "E02"},
			},
		},
	}

	rawLogic := RawTransition{
		{State: "S02", Label: "M12"}: {
			Name:    "M1",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
		{State: "S12", Label: "M21"}: {
			Name:    "M2",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "simplelib", "SimpleStruct")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "school/project/simplelib")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/M12_M21_end.go")
	if err != nil {
		t.Fatalf("Failed to generated api file\n");
	}
}

// M13.End or M31.End
func TestGenerateTSSM13or31_Simple(t *testing.T) {
	fsm := FSM{
		Labels: SetFromLabels("M13", "M31"),
		States: SetFromStates("S03", "E03"),
		S0:     "S03",
		F:      "E03",
		D: Transition{
			{"S03", "M13"}: []Pair{
				{TransitionState: "E03"},
			},
			{"S03", "M31"}: []Pair{
				{TransitionState: "E03"},
			},
		},
	} 

	rawLogic := RawTransition{
		{State: "S03", Label: "M13"}: {
			Name:    "M1",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
		{State: "S03", Label: "M31"}: {
			Name:    "M2",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "simplelib", "SimpleStruct")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "school/project/simplelib")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/M13_or_M31_end.go")
	if err != nil {
		t.Fatalf("Failed to generated api file\n");
	}
}

// M14.M41.End or M41.M14.End
func TestGenerateTSSM12or21_Simple(t *testing.T) {
	fsm := FSM{
		Labels: SetFromLabels("M14", "M41"),
		States: SetFromStates("S04", "S14", "S41", "E04"),
		S0:     "S04",
		F:      "E04",
		D: Transition{
			{"S04", "M14"}: []Pair{
				{TransitionState: "S14"},
			},
			{"S14", "M41"}: []Pair{
				{TransitionState: "E04"},
			},
			{"S04", "M41"}: []Pair{
				{TransitionState: "S41"},
			},
			{"S41", "M14"}: []Pair{
				{TransitionState: "E04"},
			},
		},
	}

	rawLogic := RawTransition{
		{State: "S04", Label: "M14"}: {
			Name:    "M1",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
		{State: "S04", Label: "M41"}: {
			Name:    "M2",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
		{State: "S41", Label: "M14"}: {
			Name:    "M1",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
		{State: "S14", Label: "M41"}: {
			Name:    "M2",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "simplelib", "SimpleStruct")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "school/project/simplelib")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/M14_M41_or_M41_M14_end.go")
	if err != nil {
		t.Fatalf("Failed to generated api file\n");
	}
}