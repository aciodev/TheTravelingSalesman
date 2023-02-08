package main

import (
	"testing"
)

var (
	sampleSize = 100
	grid       = loadSample(sampleSize)
)

func Benchmark_TSP_Greedy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tspGreedy(sampleSize, grid)
	}
}
