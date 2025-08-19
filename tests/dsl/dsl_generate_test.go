package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	dsl "school/project/dslgen"
)

func TestDSL_Generates_JSONDecoder(t *testing.T) {
	input := `
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
`

	// Parse + generate in one go
	outFile, err := dsl.GenerateFromDSL(input)
	require.NoError(t, err)
	require.Equal(t, "../../examples/dslgenerated/dsl_generated.go", outFile)

	// Verify file exists and has key content
	data, err := os.ReadFile(outFile)
	require.NoError(t, err)
	code := string(data)

	// spot checks – adjust to your generator's exact output strings
	require.Contains(t, code, "package ptclenforcer")
	require.Contains(t, code, "import (\n\t\"encoding/json\"\n)")
	require.Contains(t, code, "type Check struct {")
	require.True(t, strings.Contains(code, "func (x *Check) More()"))
	// require.True(t, strings.Contains(code, "func (x *StartDecode) Decode("))
	// require.True(t, strings.Contains(code, "func (x *NoMoreData) Close()"))
}
