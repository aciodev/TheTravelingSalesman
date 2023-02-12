package main

import (
	"testing"
)

var (
	sizeFixed          = 5
	sizeCustom         = 4
	gridFixed          = loadSample(sizeFixed)
	gridCustom0Indexed = [][]int{
		{-1, 10, 15, 20},
		{10, -1, 35, 25},
		{15, 35, -1, 30},
		{20, 25, 30, -1},
	}
	gridCustom1Indexed = convertTo1IndexBased(sizeCustom, gridCustom0Indexed)
)

func Benchmark_TSP_GreedyFixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tspGreedy(sizeFixed, gridFixed)
	}
}

func Benchmark_TSP_GreedyCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tspGreedy(sizeCustom, gridCustom0Indexed)
	}
}

func Benchmark_TSP_DPCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tspDynamicProgramming(sizeCustom, gridCustom1Indexed)
	}
}
