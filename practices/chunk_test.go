package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	testCases := []struct {
		name        string
		arr         []int
		size        int
		expectedArr [][]int
	}{
		{"test1", []int{1, 2, 3, 4}, 2, [][]int{[]int{1, 2}, []int{3, 4}}},
		{"test2", []int{1, 2, 3, 4, 5}, 2, [][]int{[]int{1, 2}, []int{3, 4}, []int{5}}},
		{"test3", []int{1, 2, 3, 4, 5, 6, 7, 8}, 3, [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8}}},
		{"test4", []int{1, 2, 3, 4, 5}, 4, [][]int{[]int{1, 2, 3, 4}, []int{5}}},
	}

	for _, tc := range testCases {
		res := Chunk(tc.arr, tc.size)
		assert.Equal(t, TwoDimensionalEqual(tc.expectedArr, res), true)
	}
}

func TwoDimensionalEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !Equal(v, b[i]) {
			return false
		}
	}
	return true
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
