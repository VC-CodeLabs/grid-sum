package main

import "fmt"

// Grid represents the 2D grid of integers
type Grid [][]int

// findMaxSum finds the maximum sum of a single row, column, or diagonal in the grid
func findMaxSum(grid Grid) int {
    // TODO: Implement the algorithm to find the maximum sum
    // Remember to consider all rows, columns, and diagonals

    return 0 // Placeholder return. Update this with your logic.
}

// main function to test the findMaxSum function
func main() {
    // Example grids for testing

    // Grid 1
    grid1 := Grid{
        {3, 6, 7},
        {5, 1, 42},
        {9, 14, 4},
    }

    // Grid 2
    grid2 := Grid{
        {1, 1, 3, 1},
        {40, 6, 3, 10},
        {7, 3, 11, 10},
        {2, 0, 0, 16},
    }

    // Grid 3
    grid3 := Grid{
        {1, 2},
        {3, 4},
    }

    // Test example each grids
    fmt.Printf("Max sum for Grid 1: %d (Expected: 56)\n", findMaxSum(grid1))
    fmt.Printf("Max sum for Grid 2: %d (Expected: 59)\n", findMaxSum(grid2))
    fmt.Printf("Max sum for Grid 3: %d (Expected: 6)\n", findMaxSum(grid3))
}
