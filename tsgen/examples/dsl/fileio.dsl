// --- Configuration ---
pkg os
import "os"
protocol File
file ./tsgen/examples/generatedfiles/ptclenforcer/file_io.go

// --- Initial & final states ---
s0 Opened
fstate Closed

// --- States ---
state Opened
state Written
state Closed

// --- Transitions ---
Opened -> Write: Write(data []byte) -> Written
Written -> Close: Close() -> Closed

// --- Raw method signatures ---
raw Opened Write name Write(data []byte) returns int, error
raw Written Close name Close()           returns error
