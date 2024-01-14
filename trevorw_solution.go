package main

import "fmt"

// Grid represents the 2D grid of integers
type Grid [][]int

// findMaxSum finds the maximum sum of a single row or column
func findMaxSum(grid Grid) int {
	currentMax := 0
	totalSlice := make([]int, len(grid[0]))
	for _, row := range grid {
		for rindex, num := range row {
			totalSlice[rindex] = totalSlice[rindex] + num
		}

		rowSum := 0
		for i := 0; i < len(row); i++ {
			rowSum = rowSum + row[i]
		}
		if rowSum > currentMax {
			currentMax = rowSum
		}

		//fmt.Printf("Index: %d, Row: %c\n", index, row)
	}

	for j := 0; j < len(totalSlice); j++ {
		if totalSlice[j] > currentMax {
			currentMax = totalSlice[j]
		}
	}
	return currentMax
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
		result := findMaxSum(grid)
		fmt.Printf("Max sum for Grid %d: %d\n", i+1, result)
	}
}
