package main

import "math"

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
