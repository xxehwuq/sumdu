package main

import (
	"container/heap"
	"fmt"
	"os"
)

type HuffmanNode struct {
	char      rune
	frequency int
	left      *HuffmanNode
	right     *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func buildHuffmanTree(freqMap map[rune]int) *HuffmanNode {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for char, freq := range freqMap {
		heap.Push(&pq, &HuffmanNode{char: char, frequency: freq})
	}

	for len(pq) > 1 {
		left := heap.Pop(&pq).(*HuffmanNode)
		right := heap.Pop(&pq).(*HuffmanNode)
		merged := &HuffmanNode{
			frequency: left.frequency + right.frequency,
			left:      left,
			right:     right,
		}
		heap.Push(&pq, merged)
	}

	return heap.Pop(&pq).(*HuffmanNode)
}

func generateHuffmanCodes(node *HuffmanNode, prefix string, codes map[rune]string) {
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		codes[node.char] = prefix
	}
	generateHuffmanCodes(node.left, prefix+"0", codes)
	generateHuffmanCodes(node.right, prefix+"1", codes)
}

func encodeText(text string, codes map[rune]string) string {
	encodedText := ""
	for _, char := range text {
		encodedText += codes[char]
	}
	return encodedText
}

func saveToFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Помилка запису у файл:", err)
	}
}

func decodeText(encodedText string, root *HuffmanNode) string {
	result := ""
	current := root
	for _, bit := range encodedText {
		if bit == '0' {
			current = current.left
		} else {
			current = current.right
		}

		if current.left == nil && current.right == nil {
			result += string(current.char)
			current = root
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Помилка читання файлу:", err)
		return
	}
	text := string(content)

	freqMap := make(map[rune]int)
	for _, char := range text {
		freqMap[char]++
	}

	root := buildHuffmanTree(freqMap)
	codes := make(map[rune]string)
	generateHuffmanCodes(root, "", codes)

	encodedText := encodeText(text, codes)
	saveToFile("encoded.txt", encodedText)
	fmt.Println("Закодований текст збережено у 'encoded.txt'")

	decodedText := decodeText(encodedText, root)
	fmt.Println("Декодований текст:", decodedText)
}
