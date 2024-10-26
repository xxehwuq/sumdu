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

var assignments, comparisons int

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		comparisons++
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			assignments += 3
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	assignments += 3
	return i + 1
}

func measureSort(arr []int) (time.Duration, int, int) {
	assignments, comparisons = 0, 0
	start := time.Now()
	quickSort(arr, 0, len(arr)-1)
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

		duration, assignments, comparisons = measureSort(randomArr)
		printStats("Випадковий", duration, assignments, comparisons)

		duration, assignments, comparisons = measureSort(ascendingArr)
		printStats("Зростаючий", duration, assignments, comparisons)

		duration, assignments, comparisons = measureSort(descendingArr)
		printStats("Спадаючий", duration, assignments, comparisons)
	}
}
