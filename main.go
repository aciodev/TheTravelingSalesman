package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func testCustomInput() {
	// Correct answer is 80. Input credit InterviewBit.com.
	sampleSize := 4
	zeroIndexed := [][]int{
		{-1, 10, 15, 20},
		{10, -1, 35, 25},
		{15, 35, -1, 30},
		{20, 25, 30, -1},
	}

	oneIndexed := [][]int{
		{0, 0, 0, 0, 0},
		{0, -1, 10, 15, 20},
		{0, 10, -1, 35, 25},
		{0, 15, 35, -1, 30},
		{0, 20, 25, 30, -1},
	}

	answer := tspGreedy(sampleSize, zeroIndexed)
	fmt.Println("Custom input (Greedy):", answer)

	answer = tspDynamicProgramming(sampleSize, oneIndexed)
	fmt.Println("Custom input (Dynamic Programming):", answer)
}

func testFixedInputGreedy(size int) {
	grid := loadSample(size, math.MaxInt32, false)
	answer := tspGreedy(size, grid)
	fmt.Println("Fixed Input (Greedy):", answer)
}

func testFixedInputDynamicProgramming(size int) {
	grid := loadSample(size, 0, true) // Represent disconnected edges as '0'
	answer := tspDynamicProgramming(size, grid)
	fmt.Println("Fixed Input (Dynamic Programming):", answer)
}

func main() {
	testCustomInput()
	testFixedInputGreedy(5)
	testFixedInputDynamicProgramming(5)
}

// Utility methods
// loadSample - Load the gridFixed for a given sample size
func loadSample(size int, disconnect int, oneIndexed bool) [][]int {
	b, err := os.ReadFile(fmt.Sprintf("./data/n_%d.txt", size))
	if err != nil {
		panic(err)
	}

	r := strings.NewReader(string(b))

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	gridSize := size
	start := 0
	if oneIndexed {
		gridSize += 1
		start += 1
	}

	// Create grid
	slice := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		slice[i] = make([]int, gridSize)
	}

	// Create loop variables
	i, j := start, start

	for scanner.Scan() {
		x, errConv := strconv.Atoi(scanner.Text())
		if errConv != nil {
			panic(errConv)
		}

		if x < 0 {
			slice[i][j] = disconnect
		} else {
			slice[i][j] = x
		}

		j++

		if j == gridSize {
			j = start
			i++
		}
	}

	return slice
}
