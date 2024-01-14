package main

import "fmt"
import "sync"
import "time"
import "math"

var multithreaded = true
var includeFatGrid = false
var fatDim = 100000

// Grid represents the 2D grid of integers
type Grid [][]int

type IndexType string
var Index = struct {
	ROW IndexType
	COL IndexType
}{
	ROW: "row",
	COL: "col",
}

var (
	mu sync.Mutex
	max int
	maxIndexType IndexType
	maxIndex int
)

var wg sync.WaitGroup

func checkRowSum(grid Grid, r int, cols int) {

	var sum = 0;
	for c := 0; c < cols; c++ {
		sum += grid[r][c]
	}

	if multithreaded {
		mu.Lock()
	}

	if sum > max { 
		max = sum
		maxIndexType = Index.ROW
		maxIndex = r
	}

	if multithreaded {
		mu.Unlock()
		wg.Done()
	}

}

func checkColSum(grid Grid, c int, rows int) {

	var sum = 0;
	for r := 0; r < rows; r++ {
		sum += grid[r][c]
	}

	if multithreaded {
		mu.Lock()
	}

	if sum > max { 
		max = sum
		maxIndexType = Index.COL
		maxIndex = c
	}

	if multithreaded {
		mu.Unlock()
		wg.Done()
	}

}

// findMaxSum finds the maximum sum of a single row or column
func findMaxSum(grid Grid) int {

	if multithreaded {
		mu.Lock()
	}

	max = math.MinInt
	maxIndex = -1

	if multithreaded {
		mu.Unlock()
	}

    // Remember to consider all rows and columns
	var rows = len(grid)
	var cols = len(grid[0])
	
	for r := 0; r < rows; r++ {
		if multithreaded {
			wg.Add(1)
			go checkRowSum(grid, r, cols)
		} else {
			checkRowSum(grid, r, cols)
		}
	}

	for c := 0; c < cols; c++ {
		if multithreaded {
			wg.Add(1)
			go checkColSum(grid, c, rows)
		} else {
			checkColSum(grid, c, rows)
		}
	}

	if multithreaded {
		wg.Wait()
	}

	fmt.Printf( "max sum = %d in %dx%d %s %d\n", max, rows, cols, maxIndexType, maxIndex )
	
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

	fmt.Printf("multithreaded: %t\nincludeFatGrid: %t\n", multithreaded, includeFatGrid)

	if includeFatGrid {

		fmt.Printf("assembling %dx%d fat grid...\n", fatDim, fatDim)
		fgs := time.Now()

		fatGrid := make(Grid,fatDim)
		
		fatValue := 0
		for fr := 0; fr < fatDim; fr++ {
			fatGrid[fr] = make([]int, fatDim)
			for fc := 0; fc < fatDim; fc++ {
				fatGrid[fr][fc] = fatValue
				fatValue++
			}
		}

		grids = append( grids, fatGrid)

		fgd := time.Since(fgs)
		fmt.Printf("...fat grid assembled in %s\n", fgd)
	}

	// Run the function for each grid and measure the time
	ts := time.Now()
	for i, grid := range grids {

		s := time.Now()
		result := findMaxSum(grid)
		d := time.Since(s);

		fmt.Printf("Max sum for Grid %d: %d in %s\n", i+1, result, d)
	}
	td := time.Since(ts)
	fmt.Printf("total time %s\n", td)
}
