package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	populationSize = 500
	maxGenerations = 1000
	crossoverRate  = 0.8
	mutationRate   = 0.1
)

type Chromosome struct {
	X, Y         int
	Adaptability float64
}

func diophantineEquation(x, y int) int {
	return x*x - 4*x - y*y + 2*y + 6
}

func calculateAdaptability(x, y int) float64 {
	return 1.0 / (1.0 + math.Abs(float64(diophantineEquation(x, y))))
}

func generatePopulation() []Chromosome {
	population := make([]Chromosome, populationSize)

	for i := range population {
		x := rand.Intn(21) - 10
		y := rand.Intn(21) - 10
		population[i] = Chromosome{
			X: x, Y: y,
			Adaptability: calculateAdaptability(x, y),
		}
	}

	return population
}

func selectParent(population []Chromosome) Chromosome {
	totalAdaptability := 0.0
	for _, individual := range population {
		totalAdaptability += individual.Adaptability
	}

	r := rand.Float64() * totalAdaptability
	accumulated := 0.0

	for _, individual := range population {
		accumulated += individual.Adaptability
		if accumulated >= r {
			return individual
		}
	}

	return population[0]
}

func crossover(parent1, parent2 Chromosome) (Chromosome, Chromosome) {
	if rand.Float64() > crossoverRate {
		return parent1, parent2
	}

	child1 := Chromosome{
		X: (parent1.X + parent2.X) / 2,
		Y: (parent1.Y + parent2.Y) / 2,
	}

	child2 := Chromosome{
		X: (parent1.Y + parent2.Y) / 2,
		Y: (parent1.X + parent2.X) / 2,
	}

	return child1, child2
}

func mutate(individual Chromosome) Chromosome {
	if rand.Float64() < mutationRate {
		individual.X += rand.Intn(3) - 1
		individual.Y += rand.Intn(3) - 1
	}

	return individual
}

func geneticAlgorithm() []Chromosome {
	population := generatePopulation()
	var bestSolutions []Chromosome

	for generation := 0; generation < maxGenerations; generation++ {
		newPopulation := make([]Chromosome, 0, populationSize)
		for len(newPopulation) < populationSize {
			parent1 := selectParent(population)
			parent2 := selectParent(population)
			child1, child2 := crossover(parent1, parent2)
			child1 = mutate(child1)
			child2 = mutate(child2)
			child1.Adaptability = calculateAdaptability(child1.X, child1.Y)
			child2.Adaptability = calculateAdaptability(child2.X, child2.Y)
			newPopulation = append(newPopulation, child1, child2)
		}

		population = newPopulation
		for _, individual := range population {
			if diophantineEquation(individual.X, individual.Y) == 0 {
				bestSolutions = append(bestSolutions, individual)
			}
		}
	}

	return uniqueSolutions(bestSolutions)
}

func uniqueSolutions(solutions []Chromosome) []Chromosome {
	seen := make(map[string]bool)
	var unique []Chromosome

	for _, s := range solutions {
		key := fmt.Sprintf("%d,%d", s.X, s.Y)
		if !seen[key] {
			seen[key] = true
			unique = append(unique, s)
		}
	}

	return unique
}

func main() {
	solutions := geneticAlgorithm()

	fmt.Println("Розв'язки діафантового рівняння:")
	for _, solution := range solutions {
		fmt.Printf("x = %d, y = %d\n", solution.X, solution.Y)
	}

	if len(solutions) == 0 {
		fmt.Println("Розв'язків не знайдено")
	} else {
		fmt.Printf("Знайдено %d розв'язків.\n", len(solutions))
	}
}
