package example_tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "school/project/examples/ptclenforcer"
)

func M1_End(t *testing.T) E00 {
	start, finalState := NewS00()
	defer finalState.FinalizeChecker()()

	end, res := start.M01()
	if res != "M1" {
		t.Errorf("Expected M1, got %s", res)
	}
	if end.Prtcl == nil {
		t.Error("Expected non-nil End state")
	}

	// Try to reuse "start"
	expectPanic(t, func() {
		start.M01()
	}, "S00.M01 reuse")

	return end
}

func Test_M1_End(t *testing.T) {
	end := M1_End(t)
	assert.IsType(t, E00{}, end)
}

func Test_M1_NotToEnd_ShouldPanic(t *testing.T) {
	s00, finalState := NewS00()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to not reaching final state, but no panic occurred")
		}
	}()

	// should panic if M01 is not called
	defer finalState.FinalizeChecker()() 

	// intentionally do NOT call s00.M01()
	_ = s00
}

func M11_M11_End(t *testing.T) E01 {
	start, finalState := NewS01()
	defer finalState.FinalizeChecker()()

	mid, res1 := start.M11()
	end, res2 := mid.M11()

	if res1 != "M1" || res2 != "M1" {
		t.Errorf("Expected M1 -> M1, got %s -> %s", res1, res2)
	}
	if end.Prtcl == nil {
		t.Error("Expected non-nil End state")
	}
	
	// Try to reuse "start"
	expectPanic(t, func() {
		start.M11()
	}, "S01.M11 reuse")

	// Try to reuse "mid"
	expectPanic(t, func() {
		mid.M11()
	}, "mid.M11 reuse")

	return end
}

func Test_M11_M11_End(t *testing.T) {
	end := M11_M11_End(t)
	assert.IsType(t, E01{}, end)
}

func M12_M21_End(t *testing.T) E02 {
	start, finalState := NewS02()
	defer finalState.FinalizeChecker()()

	mid, res1 := start.M12()
	end, res2 := mid.M21()

	if res1 != "M1" || res2 != "M2" {
		t.Errorf("Expected M1 -> M2, got %s -> %s", res1, res2)
	}
	if end.Prtcl == nil {
		t.Error("Expected non-nil End state")
	}

	// Try to reuse "start"
	expectPanic(t, func() {
		start.M12()
	}, "S02.M12 reuse")

	// Try to reuse "mid"
	expectPanic(t, func() {
		mid.M21()
	}, "mid.M21 reuse")

	return end
}

func Test_M12_M21_End(t *testing.T) {
	end := M12_M21_End(t)
	assert.IsType(t, E02{}, end)
}

func M13_or_M31_End(t *testing.T) (E03, E03) {
	start1, final1 := NewS03()
	defer final1.FinalizeChecker()()
	start2, final2 := NewS03()
	defer final2.FinalizeChecker()()

	end1, res1 := start1.M13()
	end2, res2 := start2.M31()

	if res1 != "M1" || res2 != "M2" {
		t.Errorf("Expected M1 -> M2, got %s -> %s", res1, res2)
	}
	if end1.Prtcl == nil || end2.Prtcl == nil {
		t.Error("Expected non-nil End states")
	}

	// Try to reuse "start"
	expectPanic(t, func() {
		start1.M31()
	}, "S03.M31 reuse")

	// Try to reuse "start2"
	expectPanic(t, func() {
		start2.M13()
	}, "S03.M13 reuse")

	return end1, end2
}

func Test_M13_or_M31_End(t *testing.T) {
	e1, e2 := M13_or_M31_End(t)
	assert.IsType(t, E03{}, e1)
	assert.IsType(t, E03{}, e2)
}


func M14_M41_or_M41_M14_End(t *testing.T) (E04, E04) {
	start1, final1 := NewS04()
	defer final1.FinalizeChecker()()
	start2, final2 := NewS04()
	defer final2.FinalizeChecker()()

	mid1, res1 := start1.M14()
	end1, res2 := mid1.M41()

	mid2, res3 := start2.M41()
	end2, res4 := mid2.M14()

	if res1 != "M1" || res2 != "M2" || res3 != "M2" || res4 != "M1" {
		t.Errorf("Unexpected transition outputs: %s%s, %s%s", res1, res2, res3, res4)
	}
	if end1.Prtcl == nil || end2.Prtcl == nil {
		t.Error("Expected non-nil End state")
	}

	// Try to reuse "start"
	expectPanic(t, func() {
		start1.M41()
	}, "S04.M41 reuse")

	// Try to reuse "mid1"
	expectPanic(t, func() {
		mid1.M41()
	}, "mid1.M41 reuse")

	// Try to reuse "start2"
	expectPanic(t, func() {
		start2.M14()
	}, "S04.M14 reuse")

	// Try to reuse "mid2"
	expectPanic(t, func() {
		mid2.M14()
	}, "mid2.M14 reuse")

	return end1, end2
}