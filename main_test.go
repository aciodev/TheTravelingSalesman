package main

import (
	"fmt"
	"math"
	"testing"
)

var (
	testCases = []int{5, 6, 7, 8, 9, 10}
)

func BenchmarkGreedy(b *testing.B) {
	inputsZeroIndexed := createZeroIndexedInput()
	b.ResetTimer()

	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspGreedy(v, inputsZeroIndexed[v])
			}
		})
	}
}

func BenchmarkDynamicProgramming(b *testing.B) {
	inputsOneIndexed := createOneIndexedInput()
	b.ResetTimer()

	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspDynamicProgramming(v, inputsOneIndexed[v])
			}
		})
	}
}

// Helper function to create isolated data (zero indexed)
func createZeroIndexedInput() map[int][][]int {
	inputsZeroIndexed := make(map[int][][]int)
	for _, testCase := range testCases {
		inputsZeroIndexed[testCase] = loadSample(testCase, math.MaxInt32, false)
	}
	return inputsZeroIndexed
}

// Helper function to create isolated data (zero indexed)
func createOneIndexedInput() map[int][][]int {
	inputsOneIndexed := make(map[int][][]int)
	for _, testCase := range testCases {
		inputsOneIndexed[testCase] = loadSample(testCase, 0, true)
	}
	return inputsOneIndexed
}
