package main

import (
	"math"
	"sync"
)

var mutex = sync.RWMutex{}

// tspParDPv2 - The traveling salesman problem using dynamic programming.
// TC: O(N^2 * 2^N)
// SC: O(N^2)
// Sourced from: https://www.geeksforgeeks.org/travelling-salesman-problem-using-dynamic-programming/
// Translated to Go by Andres Cruz on February 12, 2023
func tspParDPv2(size int, grid [][]int) int {
	memoHeight := size + 1
	memoDepth := 1 << (size + 1)
	memo := make([][]int, memoHeight)
	for i := 0; i < memoHeight; i += 1 {
		memo[i] = make([]int, memoDepth)
	}

	// Create a channel with size many workers
	channel := make(chan int, size)

	// Create size many goroutines where 'workerId' is the node
	for i := 1; i <= size; i++ {
		go func(i int) {
			mask := (1 << (size + 1)) - 1
			channel <- tspParDPv2Helper(i-1, i, mask, size, memo, grid) + grid[i][1]
		}(i)
	}

	// Wait for the threads to all return
	min := math.MaxInt32
	for i := 0; i < size; i += 1 {
		min = intMin(min, <-channel)
	}

	return min
}

// tspParDPv2Helper - Helper recursive method for the above function.
// See the tspParDPv2 function above for attribution and other information.
func tspParDPv2Helper(worker, i, mask, size int, memo [][]int, grid [][]int) int {
	if mask == ((1 << i) | 3) {
		return grid[1][i]
	}

	mutex.RLock()
	cached := memo[i][mask]
	mutex.RUnlock()

	if cached != 0 {
		return cached
	}

	min := math.MaxInt32
	for j := 1; j <= size; j++ {
		if (mask&(1<<j)) != 0 && j != i && j != 1 {
			res := tspParDPv2Helper(worker, j, mask&(^(1 << i)), size, memo, grid)
			min = intMin(min, res+grid[j][i])
		}
	}

	mutex.Lock()
	memo[i][mask] = min
	mutex.Unlock()
	return min
}
