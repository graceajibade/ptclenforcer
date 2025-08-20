// --- Configuration ---
pkg bufio
import "bufio"
protocol Scanner
file ./tsgen/examples/generatedfiles/ptclenforcer/bufio_scanner.go

// --- Initial & final states ---
s0 StartScan
fstate ScannerDone

// --- States ---
state StartScan
state HasLine
state NoMoreLines
state ScannerDone

// --- Transitions with outcomes ---
StartScan -> Scan: Scan() -> HasLine when true
StartScan -> Scan: Scan() -> NoMoreLines when false
HasLine   -> Text: Text() -> StartScan
NoMoreLines -> Close: Err() -> ScannerDone

// --- Raw method signatures ---
raw StartScan  Scan  name Scan()     returns bool
raw HasLine    Text  name Text()     returns string
raw NoMoreLines Close name Err()      returns error
