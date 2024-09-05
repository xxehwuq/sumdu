package main

import (
	"fmt"
	"math/rand/v2"
)

// Структура даних "Черга"
type Queue struct {
	Elements []int
	Size     int
}

// Створення нової черги за заданим розміром
func NewQueue(size int) *Queue {
	return &Queue{
		Elements: []int{},
		Size:     size,
	}
}

// Додавання елементу до черги
func (q *Queue) Enqueue(element int) {
	if q.IsFull() {
		fmt.Println("Черга переповнена")
		return
	}
	q.Elements = append(q.Elements, element)
}

// Отримання та видалення першого елементу черги
// Повернення цього елементу
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Черга пуста")
		return 0
	}

	element := q.Elements[0]
	if q.Length() == 1 {
		q.Elements = nil
		return element
	}

	q.Elements = q.Elements[1:]

	return element
}

// Визначення довжини черги
// Повернення довжини черги
func (q *Queue) Length() int {
	return len(q.Elements)
}

// Перевірка черги на порожність
// Повернення true або false
func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

// Перевірка черги на заповненість
// Повернення true або false
func (q *Queue) IsFull() bool {
	return q.Length() == q.Size
}

// Друк черги
func (q *Queue) Print() {
	fmt.Println(q.Elements)
}

// Заповнення черги випадковими числами від 1 до 1000
func (q *Queue) fillWithRandomNumbers() {
	for i := 0; i < q.Size; i++ {
		num := rand.IntN(1000-1) + 1 // Генерування випадкового числа від 1 до 1000
		q.Enqueue(num)
	}
}

// Перевірка чи є число простим
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// Вивід простих чисел з черги
func (q *Queue) printPrimes() {
	for _, num := range q.Elements {
		if isPrime(num) {
			fmt.Print(num, " ")
		}
	}
}

func main() {
	queueSize := 9*5 + 50
	var queue = NewQueue(queueSize)

	queue.fillWithRandomNumbers()
	fmt.Print("Черга, заповнена випадковими числами: ")
	queue.Print()

	fmt.Print("Прості числа з черги: ")
	queue.printPrimes()
}
