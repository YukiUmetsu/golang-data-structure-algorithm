package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfVowels(t *testing.T) {
	testCases := []struct {
		str      string
		expected int
	}{
		{"aeiou", 5},
		{"AEIOU", 5},
		{"abcdefghijklmnopqrstuvwxyz", 5},
		{"bcdfghjkl", 0},
	}

	for _, tc := range testCases {
		res := NumberOfVowels(tc.str)
		assert.Equal(t, tc.expected, res)
	}
}
