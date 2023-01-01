package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalize(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"a short sentence", "a short sentence", "A Short Sentence"},
		{"Wa lazy fox", "a lazy fox", "A Lazy Fox"},
		{"look, it is working!", "look, it is working!", "Look, It Is Working!"},
	}

	for _, tc := range testCases {
		res := Capitalize(tc.input)
		assert.Equal(t, tc.expected, res)
	}
}
