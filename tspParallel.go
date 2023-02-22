package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type offspring struct {
	genome  []int
	fitness int
}

// tspParallel - The traveling salesman problem using parallel programming.
// TC: O(N^2)
// SC: O(N)
// Sourced from: https://www.geeksforgeeks.org/traveling-salesman-problem-using-genetic-algorithm/
// Translated to Go by Andres Cruz on February 21, 2023
// 3 major changes are noted:
//  1. Introduce multi-threading option via go-routines
//  2. Change genome to []int rather than string
//  3. Major code refactor
func tspParallel(grid [][]int, n int, useGoRoutines bool) int {
	maxPopSize := 10

	// Create the initial population
	population := make([]offspring, maxPopSize)
	for i := 0; i < maxPopSize; i += 1 {
		genome := randomGenome(n)
		population[i] = offspring{
			genome:  genome,
			fitness: calculateFitness(genome, grid, n),
		}
	}

	//printGrid(population, maxPopSize, 0)

	currentGeneration := 1
	maxGenerations := 5
	temperature := 10000

	c := make(chan offspring)
	for temperature > 1000 && currentGeneration <= maxGenerations {
		sort.Slice(population, func(i, j int) bool {
			return population[i].fitness < population[j].fitness
		})

		evolvedPopulation := make([]offspring, maxPopSize)

		if useGoRoutines {
			for i := 0; i < maxPopSize; i += 1 {
				go func(c chan offspring, s offspring, g [][]int, n, t int) {
					c <- generateOffSpring(s, g, n, t)
				}(c, population[i], grid, n, temperature)
			}

			for i := 0; i < maxPopSize; i += 1 {
				evolvedPopulation[i] = <-c
			}
		} else {
			for i := 0; i < maxPopSize; i += 1 {
				evolvedPopulation[i] = generateOffSpring(population[i], grid, n, temperature)
			}
		}

		temperature = calculateCoolDown(temperature)
		population = evolvedPopulation
		//printGrid(population, maxPopSize, currentGeneration)

		currentGeneration += 1
	}

	solution := math.MaxInt32
	for i := 0; i < maxPopSize; i += 1 {
		if population[i].fitness < solution {
			solution = population[i].fitness
		}
	}
	return solution
}

// mutate - see tspParallel for citation
func mutate(genome []int, n int) {
	for {
		x := generateRandomNumber(1, n)
		y := generateRandomNumber(1, n)
		if x == y {
			continue
		}

		genome[x], genome[y] = genome[y], genome[x]
		break
	}
}

// randomGenome - see tspParallel for citation
func randomGenome(n int) []int {
	size := n + 1
	slice := make([]int, size)
	for i := 1; i < n; i += 1 {
		slice[i] = -1 // Fill in middle values with -1
	}

	for i := 1; i < n; {
		randomNumber := generateRandomNumber(1, n)
		if contains(slice, randomNumber) {
			continue
		}

		slice[i] = randomNumber
		i += 1
	}

	return slice
}

// calculateFitness - see tspParallel for citation
func calculateFitness(genome []int, grid [][]int, n int) int {
	fitness := 0
	for i := 0; i < n; i += 1 {
		fitness += grid[genome[i]][genome[i+1]]
	}
	return fitness
}

// calculateCoolDown - see tspParallel for citation
func calculateCoolDown(temp int) int {
	return (90 * temp) / 100
}

// generateOffSpring - see tspParallel for citation
func generateOffSpring(tempSalesman offspring, grid [][]int, n, temperature int) offspring {
	for {
		mutate(tempSalesman.genome, n)
		offspring := offspring{
			genome:  tempSalesman.genome,
			fitness: calculateFitness(tempSalesman.genome, grid, n),
		}

		if offspring.fitness <= tempSalesman.fitness {
			return offspring
		}

		probability := math.Pow(2.7, float64(-1*(offspring.fitness-tempSalesman.fitness)/temperature))
		if probability > 0.5 {
			return offspring
		}
	}
}

// Helper functions
func generateRandomNumber(lo, hi int) int {
	return rand.Intn(hi-lo) + lo
}

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

func printGrid(population []offspring, maxPopSize, currentGen int) {
	fmt.Println("Current generation:", currentGen)
	for i := 0; i < maxPopSize; i += 1 {
		fmt.Println(population[i].genome, "\t", population[i].fitness)
	}
}
