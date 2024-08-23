package main

import (
	"fmt"
	"math/rand/v2"
)

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

// Заповнення стека випадковими числами від 1 до 1000
func fillStackWithRandomNumbers(stack *Stack, n int) {
	for i := 0; i < n; i++ {
		num := rand.IntN(1000-1) + 1 // Генерування випадкового числа від 1 до 1000
		stack.push(num)
	}
}

// Розподіл чисел на парні та непарні в окремі стеки
func splitStackByParity(original *Stack) (*Stack, *Stack) {
	evenStack := NewStack(original.sizeStack()) //
	oddStack := NewStack(original.sizeStack())  //

	for !original.isEmpty() {
		value := original.pop()
		if value%2 == 0 {
			evenStack.push(value)
		} else {
			oddStack.push(value)
		}
	}

	return evenStack, oddStack
}

// Виведення елементів стека
func (s *Stack) print() {
	for i := 0; i < s.sizeStack(); i++ {
		fmt.Print(s.data[i], " ")
	}

	fmt.Println()
}

func main() {
	// Завдання 1: Визначення розміру стека за формулою S = 9 * 5 + 50
	stackSize := 9*5 + 50
	stack := NewStack(stackSize)

	// Завдання 2: Заповнення стека випадковими числами
	fillStackWithRandomNumbers(stack, stackSize)

	// Завдання 3: Розподіл чисел на парні та непарні в окремі стеки
	evenStack, oddStack := splitStackByParity(stack)

	// Завдання 4: Виведення елементів парного і непарного стеків
	fmt.Println("Стек з парними числами: ")
	evenStack.print()

	fmt.Println("Стек з непарними числами: ")
	oddStack.print()
}
