package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxChar(t *testing.T) {
	testCases := []struct {
		char     string
		expected string
	}{
		{"abcdefghijklmnaaaaa", "a"},
		{"ab1c1d1e1f1g1", "1"},
		{"", ""},
	}

	for _, tc := range testCases {
		res := MaxChar(tc.char)
		assert.Equal(t, tc.expected, res)
	}
}
