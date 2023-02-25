package main

import "math"

// tspParDPHelper - The traveling salesman problem using dynamic programming.
// TC: O(N^2 * 2^N)
// SC: O(N^2)
// Sourced from: https://www.geeksforgeeks.org/travelling-salesman-problem-using-dynamic-programming/
// Translated to Go by Andres Cruz on February 12, 2023
func tspParDP(size int, grid [][]int) int {
	memoHeight := size + 1
	memoDepth := 1 << (size + 1)
	// Create memo 3D array
	// - The first dimension represents size many goroutines
	// - The second dimension represents the memoHeight many rows
	// - The third dimension represents the memoDepth many bit flags
	memo := make([][][]int, size)
	for workerId := 0; workerId < size; workerId++ {
		memo[workerId] = make([][]int, memoHeight)
		for j := 0; j < memoHeight; j++ {
			memo[workerId][j] = make([]int, memoDepth)
		}
	}

	// Create a channel with size many workers
	channel := make(chan int, size)

	// Create size many goroutines where 'workerId' is the node
	for i := 1; i <= size; i++ {
		go func(i int) {
			mask := (1 << (size + 1)) - 1
			channel <- tspParDPHelper(i-1, i, mask, size, memo, grid) + grid[i][1]
		}(i)
	}

	// Wait for the threads to all return
	min := math.MaxInt32
	for i := 0; i < size; i += 1 {
		min = intMin(min, <-channel)
	}

	return min
}

// tspSeqDPHelper - Helper recursive method for the above function.
// See the tspSeqDP function above for attribution and other information.
func tspParDPHelper(worker, i, mask, size int, memo [][][]int, grid [][]int) int {
	if mask == ((1 << i) | 3) {
		return grid[1][i]
	}

	if memo[worker][i][mask] != 0 {
		return memo[worker][i][mask]
	}

	min := math.MaxInt32
	for j := 1; j <= size; j++ {
		if (mask&(1<<j)) != 0 && j != i && j != 1 {
			res := tspParDPHelper(worker, j, mask&(^(1 << i)), size, memo, grid)
			min = intMin(min, res+grid[j][i])
		}
	}

	memo[worker][i][mask] = min
	return min
}
