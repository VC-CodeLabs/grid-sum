package main

import "fmt"

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

// main function to test the findMaxSum function
func main() {
	// Example grids for testing
	grids := []Grid{
		{ // Grid 1 - Max Sum Expected: 53
			{3, 6, 7},
			{5, 1, 42},
			{9, 14, 4},
		},
		{ // Grid 2 - Max Sum Expected 59
			{1, 1, 3, 1},
			{40, 6, 3, 10},
			{7, 3, 11, 10},
			{2, 0, 0, 16},
		},
		{ // Grid 3 - Max Sum Expected - 7
			{1, 2},
			{3, 4},
		},
	}

	// Run the function for each grid and measure the time
	for i, grid := range grids {
		result := getBestPaths(grid)
		fmt.Printf("Max sum for Grid %d: %d\n", i+1, result)
	}
}
