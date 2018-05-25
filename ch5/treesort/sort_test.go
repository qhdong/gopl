package treesort

import (
	"testing"
)

func TestSort(t *testing.T) {
	input := []int{3, 1, 2, 4}
	want := []int{1, 2, 3, 4}
	Sort(input)
	if !equal(input, want) {
		t.Errorf("Sort(%v) = %v", input, want)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}
