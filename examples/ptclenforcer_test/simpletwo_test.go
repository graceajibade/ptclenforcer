package example_tests

import (
	"testing"

	. "school/project/examples/ptclenforcer"
	"github.com/stretchr/testify/assert"
)

func MM1_True_Ends(t *testing.T) EE00 {
	start, finalState := NewSS00()
	defer finalState.FinalizeChecker()()

	state1, ok := start.MM10(2)
	assert.True(t, ok, "Expected MM10(2) to return true")

	ss10, ok := state1.(SS10)
	assert.True(t, ok, "Expected SS10 after MM10(2)")
	newSS00, res := ss10.MM20()
	assert.Equal(t, "MM2", res)

	// reuse panic for SS10
	expectPanic(t, func() {
		ss10.MM20()
	}, "SS10 reused after MM20, use the new state instead")

	state2, ok := newSS00.MM10(1)
	assert.True(t, ok, "Expected MM10(1) to return true")

	ss10, ok = state2.(SS10)
	assert.True(t, ok, "Expected SS10 after MM10(1)")

	newSS00, res = ss10.MM20()
	assert.Equal(t, "MM2", res)

	state3, ok := newSS00.MM10(0)
	assert.False(t, ok, "Expected MM10(0) to return false")

	ss20, ok := state3.(SS20)
	assert.True(t, ok, "Expected SS20 after MM10(0)")

	end, res := ss20.MM30()
	assert.Equal(t, "MM3", res)
	assert.NotNil(t, end)

	// reuse panic for SS20
	expectPanic(t, func() {
		ss20.MM30()
	}, "SS20 reused after MM30")

	return end
}

func Test_MM1_True_Ends(t *testing.T) {
	end := MM1_True_Ends(t)
	assert.IsType(t, EE00{}, end)
}

func MM1_False_Ends(t *testing.T) EE00 {
	start, finalState := NewSS00()
	defer finalState.FinalizeChecker()()

	state, ok := start.MM10(0)
	assert.False(t, ok, "Expected MM10(0) to return false")

	ss20, ok := state.(SS20)
	assert.True(t, ok, "Expected SS20 after MM10(0)")

	end, res := ss20.MM30()
	assert.Equal(t, "MM3", res)
	assert.NotNil(t, end)

	// reuse panic for SS20
	expectPanic(t, func() {
		ss20.MM30()
	}, "SS20 reused after MM30")

	return end
}

func Test_MM1_False_Ends(t *testing.T) {
	end := MM1_False_Ends(t)
	assert.IsType(t, EE00{}, end)
}
