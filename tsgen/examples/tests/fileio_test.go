package example_tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	. "school/project/tsgen/examples/generatedfiles/ptclenforcer"
)

func File_Write_Then_Close(t *testing.T) Closed {
	tmpfile, err := os.CreateTemp("", "typestate_example_*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	// --- Write flow ---
	writeStart, writeFinal := NewOpened(tmpfile)
	defer writeFinal.FinalizeChecker()()

	written, n, err := writeStart.Write([]byte("hello, FSM!"))
	assert.NoError(t, err)
	assert.Equal(t, 11, n)

	closedWrite, err := written.Close()
	assert.NoError(t, err)
	assert.NotNil(t, closedWrite.Prtcl)

	expectPanic(t, func() {
		written.Close()
	}, "written reused")
	return closedWrite
}

func Test_File_Write_Then_Close(t *testing.T) {
	closedWrite := File_Write_Then_Close(t)
	assert.IsType(t, Closed{}, closedWrite)
}

func Test_File_NotToEnd_ShouldPanic(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "typestate_example_*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	start, finalState := NewOpened(tmpfile)

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for not reaching final state")
		}
	}()
	defer finalState.FinalizeChecker()()

	// Deliberately skip Write/Close
	_ = start
}