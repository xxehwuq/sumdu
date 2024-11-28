package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	populationSize = 1000
	maxGenerations = 500
	crossoverRate  = 0.8
	mutationRate   = 0.1
)

type Chromosome struct {
	X1, X2       float64
	Adaptability float64
	Feasible     bool
}

func objectiveFunction(x1, x2 float64) float64 {
	return math.Pow(x1-2, 2) + math.Pow(x2-2, 2)
}

func isFeasible(x1, x2 float64) bool {
	return (2*x1+x2 <= 10) &&
		(-2*x1+3*x2 <= 6) &&
		(2*x1+4*x2 >= 8) &&
		(x1 >= 0) && (x2 >= 0)
}

func generatePopulation() []Chromosome {
	population := make([]Chromosome, populationSize)
	for i := range population {
		x1 := rand.Float64() * 10
		x2 := rand.Float64() * 10
		fitness := objectiveFunction(x1, x2)
		feasible := isFeasible(x1, x2)
		population[i] = Chromosome{x1, x2, fitness, feasible}
	}
	return population
}

func evaluateFitness(population []Chromosome) {
	for i := range population {
		population[i].Adaptability = objectiveFunction(population[i].X1, population[i].X2)
		population[i].Feasible = isFeasible(population[i].X1, population[i].X2)
	}
}

func selectParent(population []Chromosome) Chromosome {
	totalFitness := 0.0

	for _, individual := range population {
		if individual.Feasible {
			totalFitness += 1 / (1 + individual.Adaptability)
		}
	}

	r := rand.Float64() * totalFitness
	accumulated := 0.0

	for _, individual := range population {
		if individual.Feasible {
			accumulated += 1 / (1 + individual.Adaptability)
			if accumulated >= r {
				return individual
			}
		}
	}

	return population[0]
}

func crossover(parent1, parent2 Chromosome) (Chromosome, Chromosome) {
	if rand.Float64() > crossoverRate {
		return parent1, parent2
	}

	alpha := rand.Float64()

	child1 := Chromosome{
		X1: alpha*parent1.X1 + (1-alpha)*parent2.X1,
		X2: alpha*parent1.X2 + (1-alpha)*parent2.X2,
	}

	child2 := Chromosome{
		X1: alpha*parent2.X1 + (1-alpha)*parent1.X1,
		X2: alpha*parent2.X2 + (1-alpha)*parent1.X2,
	}

	return child1, child2
}

func mutate(individual Chromosome) Chromosome {
	if rand.Float64() < mutationRate {
		individual.X1 += rand.NormFloat64()
		individual.X2 += rand.NormFloat64()
	}

	return individual
}

func geneticAlgorithm() Chromosome {
	population := generatePopulation()
	best := population[0]

	for generation := 0; generation < maxGenerations; generation++ {
		newPopulation := make([]Chromosome, 0, populationSize)
		for len(newPopulation) < populationSize {
			parent1 := selectParent(population)
			parent2 := selectParent(population)

			child1, child2 := crossover(parent1, parent2)
			child1 = mutate(child1)
			child2 = mutate(child2)

			newPopulation = append(newPopulation, child1, child2)
		}

		evaluateFitness(newPopulation)

		population = newPopulation
		for _, individual := range population {
			if individual.Feasible && individual.Adaptability < best.Adaptability {
				best = individual
			}
		}

		fmt.Printf("Покоління %d: Найкраща пристосованість = %f\n", generation, best.Adaptability)
	}

	return best
}

func main() {
	bestSolution := geneticAlgorithm()
	fmt.Printf("Найкраще рішення: X1 = %f, X2 = %f, Пристосованість = %f\n",
		bestSolution.X1, bestSolution.X2, bestSolution.Adaptability)
}
