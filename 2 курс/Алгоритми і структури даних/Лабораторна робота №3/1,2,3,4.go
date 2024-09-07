package main

import (
	"fmt"
	"math/rand/v2"
)

// Функція для вставки елементу в геш таблицю
func insert(hashTable map[int]int, key, value int) {
	hashTable[key] = value
}

// Функція для видалення елементу з геш таблиці
func remove(hashTable map[int]int, key int) {
	delete(hashTable, key)
}

// Функція для пошуку елементу в геш таблиці
func get(hashTable map[int]int, key int) int {
	return hashTable[key]
}

// Функція для заповнення геш таблиці випадковими числами
func fillWithRandomNumbers(hashTable map[int]int, size int) {
	for i := 0; i <= size; i++ {
		num := rand.IntN(1000-1) + 1 // генерування випадкового числа від 1 до 1000
		insert(hashTable, i, num)
	}
}

// Функція для видалення парних чисел з мапи
func deleteEvenNumbers(hashTable map[int]int) {
	for key, value := range hashTable {
		if value%2 == 0 {
			remove(hashTable, key)
		}
	}
}

func main() {
	N := 9*5 + 50     // кількість елементів геш таблиці
	S := N * 75 / 100 // розмір геш таблиці (75% від N)

	hashTable := make(map[int]int, S) // створення пустої геш таблиці розміром S

	fillWithRandomNumbers(hashTable, S)
	deleteEvenNumbers(hashTable)

	fmt.Println("Елементи, що залишилися: ", hashTable) // вивід елементів, що залишилися
}
