package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{"1", 1, 1},
		{"2", 2, 1},
		{"3", 3, 2},
		{"4", 4, 3},
		{"39", 39, 63245986},
	}

	for _, tc := range testCases {
		res := Fib(tc.n)
		assert.Equal(t, tc.expected, res)
	}
}
