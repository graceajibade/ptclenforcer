package example_tests

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	. "school/project/tsgen/examples/generatedfiles/ptclenforcer"
)

func expectPanic(t *testing.T, f func(), msg string) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic: %s", msg)
		}
	}()
	f()
}

func Scanner_ReadAllLines(t *testing.T) ScannerDone {
	text := "line1\nline2\nline3\n"
	scanner := bufio.NewScanner(strings.NewReader(text))

	startPtr, finalState := NewStartScan(scanner)
	defer finalState.FinalizeChecker()()
	start := *startPtr


	var line string
	for {
		state, ok := start.Scan()
		if !ok {
			noMore := state.(NoMoreLines)
			end, err := noMore.Close()
			assert.NoError(t, err)
			assert.NotNil(t, end.Prtcl)
			return end
		}

		lineState := state.(HasLine)
		start, line = lineState.Text()
		assert.NotEmpty(t, line)
	}
}

func Test_Scanner_ReadAllLines(t *testing.T) {
	end := Scanner_ReadAllLines(t)
	assert.IsType(t, ScannerDone{}, end)
}

func Test_Scanner_NotToEnd_ShouldPanic(t *testing.T) {
	text := "x\n"
	scanner := bufio.NewScanner(strings.NewReader(text))

	start, finalState := NewStartScan(scanner)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for not reaching final state")
		}
	}()
	defer finalState.FinalizeChecker()()

	// never call Scan → panic on final check
	_ = start
}