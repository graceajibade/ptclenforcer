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

## Requirements

- Go 1.21+
- ANTLR 4 (only needed if you edit `TSgen.g4` and regenerate the parser)
  - Install guide: https://www.antlr.org/
  - Regenerate command (from repo root):
    ```bash
    antlr4 -Dlanguage=Go -visitor TSgen.g4
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
```code
go build -o tsgen ./tsgen/tsgen
```

This produces a local ./tsgen binary.

### 3) Generate code from the DSL

```code
./tsgen -dsl path/to/protocol.dsl
```

The tool will write the Go file declared in the DSL’s file line (here
./dsl_generated.go) and print its path.

### 4) Use the generated API

Import the generated file in your project and use the safe, stateful API. For example,
you’ll get constructors and methods like:
```code
check, done := NewCheck()             // constructor for the start state
next, ok := check.More()              // branches based on More()'s bool
defer done.FinalizeChecker()()        // panic if final state is never reached
```

## Part B — Call the Generator From Another Go Program (Library)

Instead of using the CLI, you can create code from a DSL string at runtime:
```code
package main

import (
    "fmt"
    dsl "school/project/dsl"
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