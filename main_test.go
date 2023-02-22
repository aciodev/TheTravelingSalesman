package main

import (
	"fmt"
	"testing"
)

func TestGreedy(t *testing.T) {
	solution := 80
	zeroIndexed := [][]int{
		{-1, 10, 15, 20},
		{10, -1, 35, 25},
		{15, 35, -1, 30},
		{20, 25, 30, -1},
	}

	result := tspGreedy(4, zeroIndexed)
	if result != solution {
		t.Fatalf("Greedy calculated %d but expected %d", solution, result)
	}
}

func TestDynamicProgramming(t *testing.T) {
	solution := 80
	oneIndexed := [][]int{
		{0, 0, 0, 0, 0},
		{0, -1, 10, 15, 20},
		{0, 10, -1, 35, 25},
		{0, 15, 35, -1, 30},
		{0, 20, 25, 30, -1},
	}

	result := tspDynamicProgramming(4, oneIndexed)
	if result != solution {
		t.Fatalf("DynamicProgramming calculated %d but expected %d", solution, result)
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
