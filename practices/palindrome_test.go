package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{"oweifjwoifa", false},
		{"aba", true},
		{" aba", false},
		{"greetings", false},
		{"10000001", true},
		{"pennep", true},
	}

	for _, tc := range testCases {
		res := IsPalindrome(tc.str)
		assert.Equal(t, tc.expected, res)
	}
}
