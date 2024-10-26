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

func mergeSort(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2

		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)

		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left, mid, right int) {
	n1, n2 := mid-left+1, right-mid
	leftArr, rightArr := make([]int, n1), make([]int, n2)

	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
		assignments++
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = arr[mid+1+j]
		assignments++
	}

	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		comparisons++
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		assignments++
		k++
	}

	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
		assignments++
	}

	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
		assignments++
	}
}

func measureSort(arr []int) (time.Duration, int, int) {
	assignments, comparisons = 0, 0
	start := time.Now()
	mergeSort(arr, 0, len(arr)-1)
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
