# Getting Started

This project is a **typestate API generator for Go**. It reads a small DSL that
describes a protocol (states, transitions, and the underlying “raw” methods on a
real Go type), and generates a safe, typed wrapper that enforces correct
state transitions at compile-time and run-time.

Typical use cases:
- Wrap `encoding/json.Decoder`, `bufio.Scanner`, `os.File`, etc. with
  a typestate API.
- Generate code from `.dsl` files via a command-line tool.
- Or call the generator directly from another Go program.

---

## Quick start

From the repository root, run everything with:

```bash
make all
```

This builds `cmd/tsgen`, runs all tests, generates the bundled file protocol,
and tests the generated wrappers. Individual steps are available as
`make build`, `make test`, and `make run-example`.

## Before and after: reject invalid operations at compile time

The ordinary `os.File` API allows this code to compile. The write fails only
at runtime because the file has already been closed:

```go
// f is an open *os.File; error handling omitted for focus.
f.Close()
f.Write([]byte("too late")) // Compiles; returns an error at runtime.
```

The [bundled file protocol](examples/cli/dsl/fileio.dsl) permits only
`Opened → Write → Written → Close → Closed`. Its generated API makes the
same invalid operation a compile-time error:

```go
// Import ts "github.com/graceajibade/ptclenforcer/examples/cli/generatedfiles/ptclenforcer".
opened, checker := ts.NewOpened(f)
defer checker.FinalizeChecker()()
written, _, err := opened.Write([]byte("hello"))
if err != nil {
    panic(err)
}
closed, err := written.Close()
if err != nil {
    panic(err)
}
closed.Write([]byte("too late")) // Compile error: Closed has no method Write.
```

Each state exposes only its permitted operations. Runtime guards detect reuse
of a consumed state, and the deferred checker verifies that the final state was
reached. Go does not enforce linear ownership: these runtime checks complement
the method restrictions, and callers must avoid bypassing the wrapper through
its exported `Prtcl` field.

## Requirements

- Go 1.24.3+
- ANTLR 4 (only needed if you edit `TSgen.g4` and regenerate the parser)
  - Install guide: https://www.antlr.org/
  - Regenerate command (from repo root):
    ```bash
    (cd antlr && antlr4 -Dlanguage=Go -visitor -o ../parser/antlr TSgen.g4)
    ```

> If you are **not** modifying the grammar, you can skip ANTLR.

## Part A — Run the Tool on a `.dsl` File (CLI)

### 1) Write a DSL file

Create a file like `protocol.dsl`:

```text
// --- Configuration ---
pkg json
import "encoding/json"
protocol Decoder
file ./dsl_generated.go

// --- Initial & final states ---
s0 Check
fstate DecoderDone

// --- States ---
state Check
state StartDecode
state NoMoreData
state DecoderDone

// --- Transitions with explicit outcomes ---
Check -> More: More() -> StartDecode when true
Check -> More: More() -> NoMoreData when false
StartDecode -> Decode: Decode(v any) -> Check
NoMoreData -> Close: Buffered() -> DecoderDone

// --- Raw methods (the real Go calls) ---
raw Check More name More() returns bool
raw StartDecode Decode name Decode(v any) returns error
raw NoMoreData Close name Buffered() returns json.Token
```

### 2) Build the CLI executable

From the repo root (where go.mod lives):
```bash
go build -o tsgen ./cmd/tsgen
```

This produces a local ./tsgen binary.

### 3) Generate code from the DSL

```bash
./tsgen -dsl ./examples/cli/dsl/fileio.dsl
```

The tool writes the Go file declared in the DSL’s `file` line and prints its
path. The bundled example writes `examples/cli/generatedfiles/ptclenforcer/file_io.go`.
To use your own protocol, pass `-dsl ./protocol.dsl`. Paths in the DSL are
relative to your current working directory.

### 4) Use the generated API

Import the generated file in your project and use the safe, stateful API. For example,
you’ll get constructors and methods like:
```go
check, done := NewCheck()             // constructor for the start state
next, ok := check.More()              // branches based on More()'s bool
defer done.FinalizeChecker()()        // panic if final state is never reached
```

## Part B — Call the Generator From Another Go Program (Library)

Instead of using the CLI, you can create code from a DSL string at runtime:
```go
package main

import (
    "fmt"
    dsl "github.com/graceajibade/ptclenforcer/dslgen"
)

func main() {
    const proto = `
pkg json
import "encoding/json"
protocol Decoder
file ./dsl_generated.go

s0 Check
fstate DecoderDone

state Check
state StartDecode
state NoMoreData
state DecoderDone

Check -> More: More() -> StartDecode when true
Check -> More: More() -> NoMoreData when false
StartDecode -> Decode: Decode(v any) -> Check
NoMoreData -> Close: Buffered() -> DecoderDone

raw Check More name More() returns bool
raw StartDecode Decode name Decode(v any) returns error
raw NoMoreData Close name Buffered() returns json.Token
`
    out, err := dsl.GenerateFromDSL(proto)
    if err != nil {
        panic(err)
    }
    fmt.Println("Generated:", out) // e.g., ./dsl_generated.go
}
```

Under the hood this calls your pipeline:

* Parse DSL → build FSM/spec
* Convert to API model
* Render Go source
* Write the file to disk

You can also bypass the DSL entirely and construct the FSM programmatically,
then call the generator (GenerateTSAPI → GenerateTSString → GenerateTSFile)
from ptclenforcer.

## ptclenforcer
This is the core engine that takes a finite state machine (FSM) and method definitions (RawTransitions), and generates Go code that enforces protocol rules using the type system.

### Testing
To test the ptclenforcer, the FSM and RawTransition have to defined as follows:

```go
fsm := FSM{
		Labels: SetFromLabels("Write", "Read", "Close"),
		States: SetFromStates("Opened", "Written", "Read", "Closed"),
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
tsString, err := GenerateTSString(api, "os")
err = GenerateTSFile(tsString, "../../examples/ptclenforcer/file_io.go")
```
See more examples [here](tests/realPkgs/realpkgs_test.go).

## dslgen
A tool that allows for protocols to be written in a custom DSL (domain-specific language), parses and converts them into structured FSMs and method definitions and then calls `ptclenforcer` for code generation.

### Testing
To test the dslgen, the protocol has to defined as follows:

```text
pkg json
import "encoding/json"
protocol Decoder
file ../../examples/dslgenerated/dsl_generated.go

s0 Check
fstate DecoderDone

state Check
state StartDecode
state NoMoreData
state DecoderDone

Check -> More: More() -> StartDecode when true
Check -> More: More() -> NoMoreData when false
StartDecode -> Decode: Decode(v any) -> Check
NoMoreData -> Close: Buffered() -> DecoderDone

raw Check More name More() returns bool
raw StartDecode Decode name Decode(v any) returns error
raw NoMoreData Close name Buffered() returns json.Token
```
The `GenerateFromDSL` function in `dslgen` can then be called to:
* Generate the DTO and
* Use ptclenforcer to generate the files.

### tsgen
A command-line tool that combines `dslgen` and `ptclenforcer`,accepting a written custom DSL as file and automating the entire process from protocol specification to generated Go code.

### Testing
To test, simply run the following commands:

```bash
go build -o tsgen ./cmd/tsgen
./tsgen -dsl ./examples/cli/dsl/fileio.dsl
go test ./...
```

### Generated files
Testing the generated files should be done the usual way, see [here](examples) for examples of generated files are their tests.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
