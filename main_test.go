package main

import (
	"fmt"
	"testing"
)

func TestGreedy(t *testing.T) {
	result := tspGreedy(4, knownInputZI)
	if result != knownInputSol {
		t.Fatalf("Greedy calculated %d but expected %d", knownInputSol, result)
	}
}

func TestDynamicProgramming(t *testing.T) {
	result := tspDynamicProgramming(4, knownInput1I)
	if result != knownInputSol {
		t.Fatalf("DynamicProgramming calculated %d but expected %d", knownInputSol, result)
	}
}

func TestGeneticParallelProgramming(t *testing.T) {
	result := tspGeneticParallel(knownInputZI, 4, false)
	if result != knownInputSol {
		t.Fatalf("ParallelProgramming calculated %d but expected %d", knownInputSol, result)
	}
}

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

func BenchmarkGeneticParallelNoRoutines(b *testing.B) {
	inputsZeroIndexed := createZeroIndexedInput()
	b.ResetTimer()
	tspGeneticParallelRoutineHelper(b, inputsZeroIndexed, false)
}

func BenchmarkGeneticParallelWithRoutines(b *testing.B) {
	inputsZeroIndexed := createZeroIndexedInput()
	b.ResetTimer()
	tspGeneticParallelRoutineHelper(b, inputsZeroIndexed, true)
}

func tspGeneticParallelRoutineHelper(b *testing.B, inputsZeroIndexed map[int][][]int, useRoutines bool) {
	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspGeneticParallel(inputsZeroIndexed[v], v, useRoutines)
			}
		})
	}
}
