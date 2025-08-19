package example_tests

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	. "school/project/examples/ptclenforcer"
	_ "github.com/mattn/go-sqlite3"
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

/*
Cmd FSM Test

We use exec.Command("echo", "hello") to create a safe and fast-running command 
that ensures the Cmd object is valid and behaves predictably across platforms. 

We also test for proper final state transitions and panic safety. If any typestate 
is reused (e.g., calling Start or Wait twice), the generated code panics as expected.

Additionally, FinalizeChecker ensures that tests verify reaching the intended final 
state — and panic otherwise.
*/

func Cmd_Start_Wait(t *testing.T) Waited {
	cmd := exec.Command("echo", "hello") // "echo" is executable and always safe
	start, finalState := NewInitialized(cmd)
	defer finalState.FinalizeChecker()()

	started, err := start.Start()
	assert.NoError(t, err)

	waited, err := started.Wait()
	assert.NoError(t, err)
	assert.NotNil(t, waited.Prtcl)

	expectPanic(t, func() {
		start.Start()
	}, "Initialized reused")

	expectPanic(t, func() {
		started.Wait()
	}, "Started reused")

	return waited
}

func Test_Cmd_Start_Wait(t *testing.T) {
	end := Cmd_Start_Wait(t)
	assert.IsType(t, Waited{}, end)
}

func Test_CmdNotToEnd_ShouldPanic(t *testing.T) {
	start, finalState := NewInitialized()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to not reaching final state")
		}
	}()

	defer finalState.FinalizeChecker()()
	_ = start
}

/*
SQL FSM Test

We use setupTestDB to initialize a real in-memory SQLite database for testing typestate transitions
involving *sql.Tx. This approach avoids the need for a full external database setup while still
exercising real SQL behavior — including transaction semantics, query execution, and finalization.

Unlike mocking, this setup verifies not just the method calls but also the correct interaction
with the actual SQL engine. This helps catch integration issues and gives confidence that
typestate transitions like Exec, Commit, and Rollback behave correctly in practice.

By using an in-memory database (":memory:"), the test remains isolated, fast, and disposable.
It also supports safe creation of real *sql.Tx values, which can be passed directly to the
generated FSM API.

As with other typestate tests, we validate proper final state transitions and ensure that
reused states panic as expected. FinalizeChecker enforces that tests must reach the terminal
FSM state to be considered complete.
*/

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	_, err = db.Exec(`CREATE TABLE dummy (id INTEGER PRIMARY KEY)`)
	assert.NoError(t, err)

	return db
}

func Tx_Exec_Commit(t *testing.T) TxFinished {
	db := setupTestDB(t)
	tx, err := db.Begin()
	assert.NoError(t, err)

	start, finalState := NewTxStarted(tx)
	defer finalState.FinalizeChecker()()

	q, res, err := start.Exec("INSERT INTO dummy(id) VALUES (?)", 42)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	end, err := q.Commit()
	assert.NoError(t, err)
	assert.NotNil(t, end.Prtcl)

	return end
}

func Test_Tx_Exec_Commit(t *testing.T) {
	end := Tx_Exec_Commit(t)
	assert.IsType(t, TxFinished{}, end)
}

func Tx_Exec_Rollback(t *testing.T) TxFinished {
	db := setupTestDB(t)
	tx, err := db.Begin()
	assert.NoError(t, err)

	start, finalState := NewTxStarted(tx)
	defer finalState.FinalizeChecker()()

	q, res, err := start.Exec("SELECT 1")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	end, err := q.Rollback()
	assert.NoError(t, err)
	assert.NotNil(t, end.Prtcl)

	return end
}

func Test_Tx_Exec_Rollback(t *testing.T) {
	end := Tx_Exec_Rollback(t)
	assert.IsType(t, TxFinished{}, end)
}

func File_Write_Then_Read_Close(t *testing.T) (Closed, Closed) {
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

	// --- Reopen file for reading ---
	reopened, err := os.Open(tmpfile.Name()) // reopen in read-only mode
	assert.NoError(t, err)

	readStart, readFinal := NewOpened(reopened)
	defer readFinal.FinalizeChecker()()

	reader, _, err := readStart.Read(make([]byte, 32))
	assert.NoError(t, err)

	closedRead, err := reader.Close()
	assert.NoError(t, err)
	assert.NotNil(t, closedRead.Prtcl)

	expectPanic(t, func() {
		readStart.Read(make([]byte, 8))
	}, "Opened reused")

	expectPanic(t, func() {
		reader.Close()
	}, "Read reused")
	return closedWrite, closedRead
}

func Test_File_Write_Then_Read_Close(t *testing.T) {
	closedWrite, closedRead := File_Write_Then_Read_Close(t)
	assert.IsType(t, Closed{}, closedWrite)
	assert.IsType(t, Closed{}, closedRead)
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

func Decoder_ReadArrayOfObjects(t *testing.T) DecoderDone {
	input := `[{"id":1}, {"id":2}, {"id":3}]`
	dec := json.NewDecoder(strings.NewReader(input))

	// Read the opening bracket
	open, err := dec.Token()
	assert.NoError(t, err)
	assert.Equal(t, json.Delim('['), open)

	startPtr, finalState := NewCheck(dec)
	defer finalState.FinalizeChecker()()
	start := *startPtr

	for {
		state, ok := start.More()
		if !ok {
			noMore := state.(NoMoreData)
			end, tok := noMore.Close()
			assert.NotNil(t, tok)
			assert.IsType(t, DecoderDone{}, end)
			return end
		}

		decodeState := state.(StartDecode)
		var obj map[string]int
		start, err = decodeState.Decode(&obj)
		assert.NoError(t, err)
		assert.Contains(t, obj, "id")
	}
}

func Test_Decoder_ReadArrayOfObjects(t *testing.T) {
	end := Decoder_ReadArrayOfObjects(t)
	assert.IsType(t, DecoderDone{}, end)
}

func Test_Decoder_NotToEnd_ShouldPanic(t *testing.T) {
	input := `[{"id": 1}]`
	dec := json.NewDecoder(strings.NewReader(input))

	start, finalState := NewCheck(dec)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for not reaching final state")
		}
	}()
	defer finalState.FinalizeChecker()()

	// do not call More/Decode/Close → triggers panic
	_ = start
}