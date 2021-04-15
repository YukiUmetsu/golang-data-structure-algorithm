package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrix(t *testing.T) {
	two := [][]int{[]int{1, 2}, []int{4, 3}}
	three := [][]int{[]int{1, 2, 3}, []int{8, 9, 4}, []int{7, 6, 5}}
	four := [][]int{[]int{1, 2, 3, 4}, []int{12, 13, 14, 5}, []int{11, 16, 15, 6}, []int{10, 9, 8, 7}}

	testCases := []struct {
		name        string
		size        int
		expectedArr [][]int
	}{
		{"1", 2, two},
		{"2", 3, three},
		{"3", 4, four},
	}

	for _, tc := range testCases {
		res := Matrix(tc.size)
		assert.Equal(t, TwoDimensionalEqual(tc.expectedArr, res), true)
	}
}
