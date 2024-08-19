package main

import "fmt"

// Структура даних "Стек"
type Stack struct {
	data []int // масив для зберігання елементів
	top  int   // індекс верхнього елемента
	size int   // максимальний розмір стека
}

// Функція для створення нового стека з заданим розміром
func NewStack(size int) *Stack {
	return &Stack{
		data: make([]int, size),
		top:  -1,
		size: size,
	}
}

// Метод перевіряє, чи порожній стек
func (s *Stack) isEmpty() bool {
	return s.top == -1
}

// Метод перевіряє, чи заповнений стек
func (s *Stack) isFull() bool {
	return s.top == s.size-1
}

// Метод додає елемент у вершину стека
func (s *Stack) push(value int) {
	if s.isFull() {
		fmt.Println("Стек переповнений!")
		return
	}
	s.top++
	s.data[s.top] = value
}

// Метод видаляє верхній елемент стека
func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Println("Стек порожній!")
		return -1
	}
	value := s.data[s.top]
	s.top--
	return value
}

// Метод повертає верхній елемент стека, не видаляючи його
func (s *Stack) peek() int {
	if s.isEmpty() {
		fmt.Println("Стек порожній!")
		return -1
	}
	return s.data[s.top]
}

// Метод повертає розмір стека
func (s *Stack) sizeStack() int {
	return s.top + 1
}

// Метод для отримання посилання на верхній елемент стека
func (s *Stack) topElement() int {
	return s.peek()
}

func main() {
	// Визначення розміру стека за формулою S = 9 * 5 + 50
	stackSize := 9*5 + 50

	NewStack(stackSize)
}
