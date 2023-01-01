package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagrams(t *testing.T) {
	testCases := []struct {
		name        string
		left        string
		right       string
		expectedRes bool
	}{
		{"hello is anagram of llohe", "hello", "llohe", true},
		{"Whoa! Hi! is anagram of Hi! Whoa!", "Whoa! Hi!", "Hi! Whoa!", true},
		{"One One is not anagram of Two Two", "One One", "Two Two", false},
		{"A tree, a life, a bench of A tree, a fence, a yard", "A tree, a life, a bench", "A tree, a fence, a yard", false},
	}

	for _, tc := range testCases {
		res := IsAnagram(tc.left, tc.right)
		assert.Equal(t, tc.expectedRes, res)
	}
}
