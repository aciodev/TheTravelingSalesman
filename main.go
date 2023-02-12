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

func main() {
	sampleSize := 4
	custom := [][]int{
		{-1, 10, 15, 20},
		{10, -1, 35, 25},
		{15, 35, -1, 30},
		{20, 25, 30, -1},
	} // Answer 80. Input credit InterviewBit.com.

	//grid := loadSample(sampleSize)
	answer := tspGreedy(sampleSize, custom)
	fmt.Println(answer)
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
