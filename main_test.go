package main

import (
	"fmt"
	"testing"
)

// ------------- Tests -------------

func TestSeqGreedy(t *testing.T) {
	result := tspSeqGreedy(4, knownInputZI)
	if result != knownInputSol {
		t.Fatalf("Greedy calculated %d but expected %d", knownInputSol, result)
	}
}

func TestSeqDP(t *testing.T) {
	result := tspSeqDP(4, knownInput1I)
	if result != knownInputSol {
		t.Fatalf("DynamicProgramming calculated %d but expected %d", knownInputSol, result)
	}
}

func TestParGenetic(t *testing.T) {
	result := tspParGenetic(knownInputZI, 4, false)
	if result != knownInputSol {
		t.Fatalf("ParallelProgramming calculated %d but expected %d", knownInputSol, result)
	}
}

// ------------- Benchmarks -------------

func BenchmarkSeqGreedy(b *testing.B) {
	inputsZeroIndexed := createZeroIndexedInput()
	b.ResetTimer()

	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspSeqGreedy(v, inputsZeroIndexed[v])
			}
		})
	}
}

func BenchmarkSeqDP(b *testing.B) {
	inputsOneIndexed := createOneIndexedInput()
	b.ResetTimer()

	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspSeqDP(v, inputsOneIndexed[v])
			}
		})
	}
}

func BenchmarkParDP(b *testing.B) {
	inputsOneIndexed := createOneIndexedInput()
	b.ResetTimer()

	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspParDP(v, inputsOneIndexed[v])
			}
		})
	}
}

func BenchmarkParGeneticNoRoutines(b *testing.B) {
	inputsZeroIndexed := createZeroIndexedInput()
	b.ResetTimer()
	tspGeneticRoutineHelper(b, inputsZeroIndexed, false)
}

func BenchmarkParGeneticWithRoutines(b *testing.B) {
	inputsZeroIndexed := createZeroIndexedInput()
	b.ResetTimer()
	tspGeneticRoutineHelper(b, inputsZeroIndexed, true)
}

func tspGeneticRoutineHelper(b *testing.B, inputsZeroIndexed map[int][][]int, useRoutines bool) {
	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspParGenetic(inputsZeroIndexed[v], v, useRoutines)
			}
		})
	}
}
