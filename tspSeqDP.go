package main

import "math"

// tspSeqDP - The traveling salesman problem using dynamic programming.
// TC: O(N^2 * 2^N)
// SC: O(N^2)
// Sourced from: https://www.geeksforgeeks.org/travelling-salesman-problem-using-dynamic-programming/
// Translated to Go by Andres Cruz on February 12, 2023
func tspSeqDP(size int, grid [][]int) int {
	memoHeight := size + 1
	memoDepth := 1 << (size + 1)
	// Create memo 2D array
	memo := make([][]int, memoHeight)
	for i := 0; i < memoHeight; i++ {
		memo[i] = make([]int, memoDepth)
	}

	min := math.MaxInt32
	for i := 1; i <= size; i++ {
		mask := (1 << (size + 1)) - 1
		res := tspDPHelper(i, mask, size, memo, grid) + grid[i][1]
		min = int(math.Min(float64(min), float64(res)))
	}

	return min
}

// tspDPHelper - Helper recursive method for the above function.
// See the tspSeqDP function above for attribution and other information.
func tspDPHelper(i, mask, size int, memo [][]int, grid [][]int) int {
	if mask == ((1 << i) | 3) {
		return grid[1][i]
	}

	if memo[i][mask] != 0 {
		return memo[i][mask]
	}

	min := math.MaxInt32
	for j := 1; j <= size; j++ {
		if (mask&(1<<j)) != 0 && j != i && j != 1 {
			res := tspDPHelper(j, mask&(^(1 << i)), size, memo, grid)
			min = int(math.Min(float64(min), float64(res+grid[j][i])))
		}
	}

	memo[i][mask] = min
	return min
}
