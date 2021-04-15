package dsalgo

func Chunk(arr []int, size int) [][]int {
	arrLen := len(arr)
	index := 0
	res := make([][]int, 0)
	for index < arrLen {
		var endIndex int
		if index+size > arrLen-1 {
			endIndex = arrLen
		} else {
			endIndex = index + size
		}

		res = append(res, arr[index:endIndex])
		index += size
	}
	return res
}
