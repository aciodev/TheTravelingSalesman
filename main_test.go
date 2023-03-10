package main

import (
	"fmt"
	"testing"
)

// ------------- Tests -------------

func TestSeqGreedy(t *testing.T) {
	result := tspSeqGreedy(4, knownInputZI)
	if result != knownInputSol {
		t.Fatalf("TestSeqGreedy calculated %d but expected %d", knownInputSol, result)
	}
}

func TestSeqDP(t *testing.T) {
	result := tspSeqDP(4, knownInput1I)
	if result != knownInputSol {
		t.Fatalf("TestSeqDP calculated %d but expected %d", knownInputSol, result)
	}
}

func TestParGenetic(t *testing.T) {
	result := tspParGenetic(knownInputZI, 4, false)
	if result != knownInputSol {
		t.Fatalf("TestParGenetic calculated %d but expected %d", knownInputSol, result)
	}
}

func TestParDPv1(t *testing.T) {
	result := tspParDPv1(4, knownInput1I)
	if result != knownInputSol {
		t.Fatalf("TestParDPv1 calculated %d but expected %d", knownInputSol, result)
	}
}

func TestParDPv2(t *testing.T) {
	result := tspParDPv2(4, knownInput1I)
	if result != knownInputSol {
		t.Fatalf("TestParDPv2 calculated %d but expected %d", knownInputSol, result)
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

func BenchmarkParDPv1(b *testing.B) {
	inputsOneIndexed := createOneIndexedInput()
	b.ResetTimer()

	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspParDPv1(v, inputsOneIndexed[v])
			}
		})
	}
}

func BenchmarkParDPv2(b *testing.B) {
	inputsOneIndexed := createOneIndexedInput()
	b.ResetTimer()

	for _, v := range testCases {
		b.Run(fmt.Sprintf("n_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tspParDPv2(v, inputsOneIndexed[v])
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
