package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// tspGreedy - The traveling salesman problem using greedy.
// TC: O(N^2 * log(N))
// SC: O(N)
// Sourced from: https://www.interviewbit.com/blog/travelling-salesman-problem/
// Translated to Go by Andres Cruz on February 7, 2023
func tspGreedy(size int, grid [][]int) int {
	sum := 0
	counter := 0
	i, j := 0, 0
	min := math.MaxInt64
	visited := make(map[int]int)
	visited[0] = 1

	route := make([]int, size)

	for i < size && j < size {
		if i != j && visited[j] == 0 && grid[i][j] < min {
			min = grid[i][j]
			route[counter] = j + 1
		}

		j++

		if j == size {
			sum += min
			min = math.MaxInt64
			visited[route[counter]-1] = 1
			i = route[counter] - 1
			j = 0
			counter++

			if counter >= size-1 {
				break
			}
		}
	}

	i = route[counter-1] - 1

	for k := 0; k < size; k++ {
		if i != k && grid[i][k] < min {
			min = grid[i][k]
			route[counter] = k + 1
		}
	}

	sum += min
	return sum
}

// tspDynamicProgramming - The traveling salesman problem using dynamic programming.
// TC: O(N^2 * 2^N)
// SC: O(N^2)
// Sourced from: https://www.geeksforgeeks.org/travelling-salesman-problem-using-dynamic-programming/
// Translated to Go by Andres Cruz on February 12, 2023
func tspDynamicProgramming(size int, grid [][]int) int {
	memoHeight := size + 1
	memoDepth := 1 << (size + 1)
	// Create memo 2D array
	memo := make([][]int, memoHeight)
	for i := 0; i < memoHeight; i++ {
		memo[i] = make([]int, memoDepth)
	}

	min := math.MaxInt16
	for i := 1; i <= size; i++ {
		mask := (1 << (size + 1)) - 1
		res := tspDPHelper(i, mask, size, memo, grid) + grid[i][1]
		min = int(math.Min(float64(min), float64(res)))
	}

	return min
}

// tspDPHelper - Helper recursive method for the above function.
// See the tspDynamicProgramming function above for attribution and other information.
func tspDPHelper(i, mask, size int, memo [][]int, grid [][]int) int {
	if mask == ((1 << i) | 3) {
		return grid[1][i]
	}

	if memo[i][mask] != 0 {
		return memo[i][mask]
	}

	min := math.MaxInt16
	for j := 1; j <= size; j++ {
		if (mask&(1<<j)) != 0 && j != i && j != 1 {
			res := tspDPHelper(j, mask&(^(1 << i)), size, memo, grid)
			min = int(math.Min(float64(min), float64(res+grid[j][i])))
		}
	}

	memo[i][mask] = min
	return min
}

func testCustomInput() {
	// Correct answer is 80. Input credit InterviewBit.com.
	sampleSize := 4
	custom := [][]int{
		{-1, 10, 15, 20},
		{10, -1, 35, 25},
		{15, 35, -1, 30},
		{20, 25, 30, -1},
	}

	answer := tspGreedy(sampleSize, custom)
	fmt.Println("Custom input (Greedy):", answer)

	oneIndexed := convertTo1IndexBased(sampleSize, custom)
	answer = tspDynamicProgramming(sampleSize, oneIndexed)
	fmt.Println("Custom input (Dynamic Programming):", answer)
}

func testFixedInputGreedy(size int) {
	grid := loadSample(size)
	answer := tspGreedy(size, grid)
	fmt.Println("Fixed Input (Greedy):", answer)
}

func testFixedInputDynamicProgramming(size int) {
	grid := loadSample(size)
	oneIndexed := convertTo1IndexBased(size, grid)
	answer := tspDynamicProgramming(size, oneIndexed)
	fmt.Println("Fixed Input (Dynamic Programming):", answer)
}

func main() {
	testCustomInput()
	testFixedInputGreedy(5)
	testFixedInputDynamicProgramming(5)
}

// Utility methods
// loadSample - Load the grid for a given sample size
func loadSample(size int) [][]int {
	b, err := os.ReadFile(fmt.Sprintf("./data/n_%d.txt", size))
	if err != nil {
		panic(err)
	}

	r := strings.NewReader(string(b))

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	slice := make([][]int, size)

	// Make size*size slice
	for i := 0; i < size; i++ {
		slice[i] = make([]int, size)
	}

	i, j := 0, 0

	for scanner.Scan() {
		x, errConv := strconv.Atoi(scanner.Text())
		if errConv != nil {
			panic(errConv)
		}

		if x < 0 {
			slice[i][j] = math.MaxInt64
		} else {
			slice[i][j] = x
		}

		j++

		if j == size {
			j = 0
			i++
		}
	}

	return slice
}

// convertTo1IndexBased - Add a row and column of zeroes
func convertTo1IndexBased(size int, grid [][]int) [][]int {
	newGrid := make([][]int, size+1)
	newGrid[0] = make([]int, size+1) // Fill top row with zeroes
	for i := 1; i <= size; i++ {
		newGrid[i] = make([]int, size+1)
		newGrid[i][0] = 0 // Column 0 initializes to zero
		// Fill in the grid
		for j := 0; j < size; j++ {
			newGrid[i][j+1] = grid[i-1][j]
		}
	}
	return newGrid
}
