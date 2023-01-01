package dsalgo

// --- Directions
// Write a function that accepts an integer N
// and returns a NxN spiral matrix.
// --- Examples
//   matrix(2)
//     [[1, 2],
//     [4, 3]]
//   matrix(3)
//     [[1, 2, 3],
//     [8, 9, 4],
//     [7, 6, 5]]
//  matrix(4)
//     [[1,   2,  3, 4],
//     [12, 13, 14, 5],
//     [11, 16, 15, 6],
//     [10,  9,  8, 7]]

func Matrix(n int) [][]int {
	// create array with 0 values
	result := createZeroMatrix(n)

	// create a counter variable starting at 1
	counter := 1
	startColumn := 0
	endColumn := n - 1
	startRow := 0
	endRow := n - 1

	// loop as long as (start column <= end column) AND (start row <= end row)
	// it means that start column, end column, start row, end row all gathered in the center of the matrix
	for (startColumn <= endColumn) && (startRow <= endRow) {
		// Top row: loop from start column to end column
		for i := startColumn; i <= endColumn; i++ {
			// results[startRow][i] assign counter variable
			result[startRow][i] = counter
			counter++
		}
		// after loop, increment start row
		startRow++

		// Right side: loop from start row to end row
		for i := startRow; i <= endRow; i++ {
			result[i][endColumn] = counter
			counter++
		}
		// after loop, decrement end column
		endColumn--

		// Bottom row: loop from end column to start column
		for i := endColumn; i >= startColumn; i-- {
			result[endRow][i] = counter
			counter++
		}
		// after loop, decrement end row
		endRow--

		// Left side: loop from end row to start row
		for i := endRow; i >= startRow; i-- {
			result[i][startColumn] = counter
			counter++
		}
		// after loop, increment start column
		startColumn++
	}
	return result
}

func createZeroMatrix(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		result = append(result, []int{})
		row := []int{}
		for j := 0; j < n; j++ {
			row = append(row, 0)
		}
		result[i] = row
	}
	return result
}
