package main

import "fmt"

// Grid represents the 2D grid of integers
type Grid [][]int

// findMaxSum finds the maximum sum of a single row or column
func findMaxSum(grid Grid) int {
    // Remember to consider all rows and columns
	var rows = len(grid)
	var cols = len(grid[0])

	var max = -1;
	var maxrow = -1;
	
	for r := 0; r < rows; r++ {
		var sum = 0;
		for c := 0; c < cols; c++ {
			sum += grid[r][c]
		}

		if sum > max { 
			max = sum
			maxrow = r
		}
	}

	var maxcol = -1;

	for c := 0; c < cols; c++ {
		var sum = 0;
		for r := 0; r < rows; r++ {
			sum += grid[r][c]
		}

		if sum > max { 
			max = sum
			maxcol = c
		}
	}

	// fmt.print( "max value %d in %s %d", max, maxcol >= 0 ? "column" : "row", maxcol >= 0 ? maxcol : maxrow )
	if maxcol >= 0 {
		fmt.Printf( "max value %d in col %d\n", max, maxcol )
	} else {
		fmt.Printf( "max value %d in row %d\n", max, maxrow )
	}

    return max // Placeholder return. Update this with your logic.
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
		{  // Grid 3 - Max Sum Expected - 7
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
