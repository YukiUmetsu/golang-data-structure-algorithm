package dsalgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSorted(t *testing.T) {
	testCases := []struct {
		name        string
		unsortedArr []int
		expectedArr []int
	}{
		{"normal sorting should work", []int{9, 8, 7, 6, 5, 4, 3, 1, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"sorting with negative numbers should work", []int{-9, -8, 7, 6, 5, 4, 3, 1, 2}, []int{-9, -8, 1, 2, 3, 4, 5, 6, 7}},
		{"empty arr should get empty arr", []int{}, []int{}},
	}

	for _, tc := range testCases {
		var arr []int = tc.unsortedArr
		res := BubbleSort(arr)
		assert.Equal(t, res, tc.expectedArr)
	}
}

func TestSelectionSorted(t *testing.T) {
	testCases := []struct {
		name        string
		unsortedArr []int
		expectedArr []int
	}{
		{"normal sorting should work", []int{9, 8, 7, 6, 5, 4, 3, 1, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"sorting with negative numbers should work", []int{-9, -8, 7, 6, 5, 4, 3, 1, 2}, []int{-9, -8, 1, 2, 3, 4, 5, 6, 7}},
		{"empty arr should get empty arr", []int{}, []int{}},
	}

	for _, tc := range testCases {
		var arr []int = tc.unsortedArr
		res := SelectionSort(arr)
		assert.Equal(t, res, tc.expectedArr)
	}
}

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name        string
		unsortedArr []int
		expectedArr []int
	}{
		{"normal sorting should work", []int{9, 8, 7, 6, 5, 4, 3, 1, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"sorting with negative numbers should work", []int{-9, -8, 7, 6, 5, 4, 3, 1, 2}, []int{-9, -8, 1, 2, 3, 4, 5, 6, 7}},
		{"empty arr should get empty arr", []int{}, []int{}},
	}

	for _, tc := range testCases {
		var arr []int = tc.unsortedArr
		res := MergeSort(arr)
		assert.Equal(t, tc.expectedArr, res)
	}
}
