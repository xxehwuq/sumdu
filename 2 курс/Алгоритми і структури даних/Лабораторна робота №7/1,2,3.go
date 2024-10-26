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

func heapSort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		assignments += 3

		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n int, i int) {
	largest := i
	left, right := 2*i+1, 2*i+2

	comparisons++
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	comparisons++
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		assignments += 3

		heapify(arr, n, largest)
	}
}

func measureSort(arr []int) (time.Duration, int, int) {
	assignments, comparisons = 0, 0
	start := time.Now()
	heapSort(arr)
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
