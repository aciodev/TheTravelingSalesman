package main

import (
	"fmt"
	"math"
	"testing"
)

var (
	inputsZeroIndexed = map[int][][]int{
		5:  loadSample(5, math.MaxInt32, false),
		6:  loadSample(6, math.MaxInt32, false),
		7:  loadSample(7, math.MaxInt32, false),
		8:  loadSample(8, math.MaxInt32, false),
		9:  loadSample(9, math.MaxInt32, false),
		10: loadSample(10, math.MaxInt32, false),
	}

	inputs1Indexed = map[int][][]int{
		5:  loadSample(5, 0, true),
		6:  loadSample(6, 0, true),
		7:  loadSample(7, 0, true),
		8:  loadSample(8, 0, true),
		9:  loadSample(9, 0, true),
		10: loadSample(10, 0, true),
	}

	testCases = []int{5, 6, 7, 8, 9, 10}
)

func BenchmarkGreedy(b *testing.B) {
	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspGreedy(v, inputsZeroIndexed[v])
			}
		})
	}
}

func BenchmarkDynamicProgramming(b *testing.B) {
	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspDynamicProgramming(v, inputs1Indexed[v])
			}
		})
	}
}
