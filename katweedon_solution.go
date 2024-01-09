package main

type GridSlice []int

func getBestPaths(grid Grid) int {
	size := len(grid)
	maxSum := 0

	for i := range grid {
		column := make(GridSlice, size)
		row := make(GridSlice, size)
		for j := range column {
			column[j] = grid[j][i]
			row[j] = grid[i][j]
		}
		colSum := getSliceSum(column)
		rowSum := getSliceSum(row)

		// check column
		if colSum > maxSum {
			maxSum = colSum
		}

		// check row
		if rowSum > maxSum {
			maxSum = rowSum
		}
	}
	return maxSum
}

func getSliceSum(slice GridSlice) int {
	var currentSum = 0
	for i := range slice {
		currentSum = currentSum + slice[i]
	}
	return currentSum
}
