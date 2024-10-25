package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func GenerateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.IntN(1000)
	}
	return arr
}

func GenerateAscendingArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

func GenerateDescendingArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i - 1
	}
	return arr
}

func insertionSort(arr []int) (int, int) {
	assignments := 0
	comparisons := 0

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		assignments++
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			assignments++
			comparisons++
			j--
		}
		arr[j+1] = key
		assignments++
	}
	return assignments, comparisons
}

func measureSort(arr []int) (time.Duration, int, int) {
	start := time.Now()
	assignments, comparisons := insertionSort(arr)
	duration := time.Since(start)
	return duration, assignments, comparisons
}

func printStats(arrType string, duration time.Duration, assignments, comparisons int) {
	fmt.Printf("%s масив - Час: %v, Присвоювання: %d, Порівняння: %d\n", arrType, duration, assignments, comparisons)
}

func main() {
	sizes := []int{10, 100, 1000, 5000, 10000}

	for _, size := range sizes {
		fmt.Printf("\nРозмір масиву: %d\n", size)

		randomArr := GenerateRandomArray(size)
		ascendingArr := GenerateAscendingArray(size)
		descendingArr := GenerateDescendingArray(size)

		var duration time.Duration
		var assignments int
		var comparisons int

		duration, assignments, comparisons = measureSort(randomArr)
		printStats("Випадковий", duration, assignments, comparisons)

		duration, assignments, comparisons = measureSort(ascendingArr)
		printStats("Зростаючий", duration, assignments, comparisons)

		duration, assignments, comparisons = measureSort(descendingArr)
		printStats("Спадаючий", duration, assignments, comparisons)
	}
}
