package dsalgo

import (
	"math"
)

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < (len(arr) - 1 - i); j++ {
			// if next one is smaller, swap
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		indexOfMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[indexOfMin] > arr[j] {
				indexOfMin = j
			}
		}
		// if indexOfMin changed, swap
		if indexOfMin != i {
			arr[indexOfMin], arr[i] = arr[i], arr[indexOfMin]
		}
	}
	return arr
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// divide the array into two pieces
	length := float64(len(arr))
	middle := int(math.Floor(length / 2))
	left := arr[:middle]
	right := arr[middle:]

	// sort 2 divided arrays and merge
	return merge(MergeSort(left), MergeSort(right))
}

// merge 2 sorted array
func merge(arr1, arr2 []int) []int {
	var resArr []int
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			resArr = append(resArr, arr1[0])
			arr1 = arr1[1:]
		} else {
			resArr = append(resArr, arr2[0])
			arr2 = arr2[1:]
		}
	}
	resArr = append(resArr, arr1...)
	resArr = append(resArr, arr2...)
	return resArr
}
