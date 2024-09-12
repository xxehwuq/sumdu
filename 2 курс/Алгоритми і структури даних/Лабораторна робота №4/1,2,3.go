package main

import (
	"fmt"
)

// Структура для вузла бінарного дерева пошуку
type Node struct {
	key   int
	name  string // Додамо поле для зберігання імені
	phone string // Поле для зберігання телефону
	left  *Node
	right *Node
}

// Структура для бінарного дерева
type BinaryTree struct {
	root *Node
}

// Функція створення нового пустого дерева
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// Функція для вставки нового елемента
func (t *BinaryTree) Insert(key int, name, phone string) {
	t.root = insertNode(t.root, key, name, phone)
}

// Рекурсивна функція вставки
func insertNode(node *Node, key int, name, phone string) *Node {
	if node == nil {
		return &Node{key: key, name: name, phone: phone}
	}
	if key < node.key {
		node.left = insertNode(node.left, key, name, phone)
	} else if key > node.key {
		node.right = insertNode(node.right, key, name, phone)
	}
	return node
}

// Функція для пошуку елемента за ключем
func (t *BinaryTree) Search(key int) *Node {
	return searchNode(t.root, key)
}

// Рекурсивна функція пошуку
func searchNode(node *Node, key int) *Node {
	if node == nil || node.key == key {
		return node
	}
	if key < node.key {
		return searchNode(node.left, key)
	}
	return searchNode(node.right, key)
}

// Функція для видалення елемента
func (t *BinaryTree) Delete(key int) {
	t.root = deleteNode(t.root, key)
}

// Рекурсивна функція видалення
func deleteNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = deleteNode(node.left, key)
	} else if key > node.key {
		node.right = deleteNode(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		minRight := findMin(node.right)
		node.key = minRight.key
		node.name = minRight.name
		node.phone = minRight.phone
		node.right = deleteNode(node.right, minRight.key)
	}
	return node
}

// Функція для пошуку мінімального елемента в дереві
func findMin(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

// Функція для виведення вузлів дерева (інфіксним обходом)
func (t *BinaryTree) InOrder() {
	inOrderTraversal(t.root)
	fmt.Println()
}

// Рекурсивна функція інфіксного обходу
func inOrderTraversal(node *Node) {
	if node != nil {
		inOrderTraversal(node.left)
		fmt.Printf("Ключ: %d, Ім'я: %s, Номер телефону: %s\n", node.key, node.name, node.phone)
		inOrderTraversal(node.right)
	}
}

func main() {
	// Створюємо нове дерево
	tree := NewBinaryTree()

	// Вставка елементів в довідник телефонів
	tree.Insert(3, "Віталій", "+380(056)34-33-14")
	tree.Insert(1, "Андрій", "+380(048)22-29-34")
	tree.Insert(4, "Олександр", "+380(0552)35-61-01")
	tree.Insert(2, "Ярослав", "+380(048)25-30-18")

	fmt.Println("Довідник телефонів (після вставки):")
	tree.InOrder()

	// Пошук елемента за ключем
	foundNode := tree.Search(2)
	if foundNode != nil {
		fmt.Printf("\nЗнайдено контакт:\nКлюч: %d, Ім'я: %s, Номер телефону: %s\n", foundNode.key, foundNode.name, foundNode.phone)
	} else {
		fmt.Printf("\nКонтакт з ключем %d не знайдено\n", 2)
	}

	// Видалення елемента
	tree.Delete(1)

	fmt.Println("\nДовідник телефонів після видалення елементу з ключем 1:")
	tree.InOrder()
}
