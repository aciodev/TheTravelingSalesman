package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
