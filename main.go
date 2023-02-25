package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	knownInputZI = [][]int{
		{-1, 10, 15, 20},
		{10, -1, 35, 25},
		{15, 35, -1, 30},
		{20, 25, 30, -1},
	}
	knownInput1I = [][]int{
		{0, 0, 0, 0, 0},
		{0, -1, 10, 15, 20},
		{0, 10, -1, 35, 25},
		{0, 15, 35, -1, 30},
		{0, 20, 25, 30, -1},
	}
	knownInputSol = 80
	testCases     = []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

// Okay, I know what you're thinking. Why defined a min
// function when math.Min is already available? The rationale
// is as follows:
//   - We do not care about special cases (inf, NaN, Signbit)
//   - The min function takes float64 as an argument, but we are working
//     with int. This causes unnecessary type conversions to parameterize,
//     and then to cast the result back to an int.
func intMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Helper function to create isolated data (zero indexed)
func createZeroIndexedInput() map[int][][]int {
	inputsZeroIndexed := make(map[int][][]int)
	for _, testCase := range testCases {
		inputsZeroIndexed[testCase] = loadSample(testCase, math.MaxInt32, false)
	}
	return inputsZeroIndexed
}

// Helper function to create isolated data (one indexed)
func createOneIndexedInput() map[int][][]int {
	inputsOneIndexed := make(map[int][][]int)
	for _, testCase := range testCases {
		inputsOneIndexed[testCase] = loadSample(testCase, 0, true)
	}
	return inputsOneIndexed
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
