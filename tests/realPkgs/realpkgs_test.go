package realpkgstest

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

func TestGenerateTSSExecCmdLifecycle(t *testing.T) {
	fsm := FSM{
		Labels: SetFromLabels("Start", "Wait"),
		States: SetFromStates("Started"),
		S0:     "Initialized",
		F:      "Waited",
		D: Transition{
			{"Initialized", "Start"}: []Pair{
				{TransitionState: "Started"},
			},
			{"Started", "Wait"}: []Pair{
				{TransitionState: "Waited"},
			},
		},
	}

	rawLogic := RawTransition{
		{State: "Initialized", Label: "Start"}: {
			Name:    "Start",
			Inputs:  []string{},
			Outputs: []string{"error"},
		},
		{State: "Started", Label: "Wait"}: {
			Name:    "Wait",
			Inputs:  []string{},
			Outputs: []string{"error"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "exec", "Cmd")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "os/exec")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/cmd_start_wait.go")
	if err != nil {
		t.Fatalf("Failed to generate api file\n")
	}
}

func TestGenerateTSSTransactionFSM(t *testing.T) {
	fsm := FSM{
		Labels: SetFromLabels("Exec", "Commit", "Rollback"),
		States: SetFromStates( "Queried"),
		S0:     "TxStarted",
		F:      "TxFinished",
		D: Transition{
			{"TxStarted", "Exec"}:    []Pair{{TransitionState: "Queried"}},
			{"Queried", "Commit"}:    []Pair{{TransitionState: "TxFinished"}},
			{"Queried", "Rollback"}:  []Pair{{TransitionState: "TxFinished"}},
		},
	}

	rawLogic := RawTransition{
		{State: "TxStarted", Label: "Exec"}: {
			Name:    "Exec",
			Inputs:  []string{"query string", "args ...any"},
			Outputs: []string{"sql.Result", "error"},
		},
		{State: "Queried", Label: "Commit"}: {
			Name:    "Commit",
			Inputs:  []string{},
			Outputs: []string{"error"},
		},
		{State: "Queried", Label: "Rollback"}: {
			Name:    "Rollback",
			Inputs:  []string{},
			Outputs: []string{"error"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "sql", "Tx")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "database/sql")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/sql_transaction.go")
	assert.NoError(t, err)
}

func TestGenerateTSFileIOFSM(t *testing.T) {
	fsm := FSM{
		Labels: SetFromLabels("Write", "Read", "Close"),
		States: SetFromStates("Written", "Read"),
		S0:     "Opened",
		F:      "Closed",
		D: Transition{
			{"Opened", "Write"}: []Pair{{TransitionState: "Written"}},
			{"Opened", "Read"}:  []Pair{{TransitionState: "Read"}},
			{"Written", "Close"}: []Pair{{TransitionState: "Closed"}},
			{"Read", "Close"}:    []Pair{{TransitionState: "Closed"}},
		},
	}

	rawLogic := RawTransition{
		{State: "Opened", Label: "Write"}: {
			Name:    "Write",
			Inputs:  []string{"data []byte"},
			Outputs: []string{"int", "error"},
		},
		{State: "Opened", Label: "Read"}: {
			Name:    "Read",
			Inputs:  []string{"buf []byte"},
			Outputs: []string{"int", "error"},
		},
		{State: "Written", Label: "Close"}: {
			Name:    "Close",
			Inputs:  []string{},
			Outputs: []string{"error"},
		},
		{State: "Read", Label: "Close"}: {
			Name:    "Close",
			Inputs:  []string{},
			Outputs: []string{"error"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "os", "File")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "os")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/file_io.go")
	assert.NoError(t, err)
}


func TestGenerateTSBufioScannerFSM(t *testing.T) {
	fsm := FSM{
		States: SetFromStates("HasLine", "NoMoreLines"),
		Labels: SetFromLabels("Scan", "Text", "Close"),
		S0:     "StartScan",
		F:      "ScannerDone",
		D: Transition{
			{"StartScan", "Scan"}: []Pair{
				{Outcome: true, TransitionState: "HasLine"},
				{Outcome: false, TransitionState: "NoMoreLines"},
			},
			{"HasLine", "Text"}: []Pair{{TransitionState: "StartScan"}},
			{"NoMoreLines", "Close"}: []Pair{{TransitionState: "ScannerDone"}},
		},
	}

	rawLogic := RawTransition{
		{State: "StartScan", Label: "Scan"}: {
			Name:    "Scan",
			Inputs:  []string{},
			Outputs: []string{"bool"},
		},
		{State: "HasLine", Label: "Text"}: {
			Name:    "Text",
			Inputs:  []string{},
			Outputs: []string{"string"},
		},
		{State: "NoMoreLines", Label: "Close"}: {
			Name:    "Err",
			Inputs:  []string{},
			Outputs: []string{"error"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "bufio", "Scanner")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "bufio")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/bufio_scanner.go")
	assert.NoError(t, err)
}

func TestGenerateTSJsonDecoderFSM(t *testing.T) {
	fsm := FSM{
		States: SetFromStates("StartDecode", "NoMoreData"),
		Labels: SetFromLabels("More", "Decode", "Close"),
		S0:     "Check",
		F:      "DecoderDone",
		D: Transition{
			{"Check", "More"}: []Pair{
				{Outcome: true, TransitionState: "StartDecode"},
				{Outcome: false, TransitionState: "NoMoreData"},
			},
			{"StartDecode", "Decode"}: []Pair{{TransitionState: "Check"}},
			{"NoMoreData", "Close"}: []Pair{{TransitionState: "DecoderDone"}},
		},
	}

	rawLogic := RawTransition{
		{State: "Check", Label: "More"}: {
			Name:    "More",
			Inputs:  []string{},
			Outputs: []string{"bool"},
		},
		{State: "StartDecode", Label: "Decode"}: {
			Name:    "Decode",
			Inputs:  []string{"v any"},
			Outputs: []string{"error"},
		},
		{State: "NoMoreData", Label: "Close"}: {
			Name:    "Buffered",
			Inputs:  []string{},
			Outputs: []string{"json.Token"},
		},
	}

	api, err := GenerateTSAPI(fsm, rawLogic, "json", "Decoder")
	assert.NoError(t, err)

	tsString, err := GenerateTSString(api, "encoding/json")
	assert.NoError(t, err)

	err = GenerateTSFile(tsString, "../../examples/ptclenforcer/json_decoder.go")
	assert.NoError(t, err)
}
