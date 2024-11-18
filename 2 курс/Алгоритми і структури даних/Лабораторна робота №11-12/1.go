package main

import (
	"container/heap"
	"fmt"
	"sort"
	"time"
)

type Edge struct {
	u, v   int
	weight int
}

type Graph struct {
	vertices int
	edges    []Edge
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)

	for i := range parent {
		parent[i] = i
	}

	return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	return true
}

func Kruskal(graph Graph) ([]Edge, int) {
	operations := 0

	sort.Slice(graph.edges, func(i, j int) bool {
		operations++
		return graph.edges[i].weight < graph.edges[j].weight
	})

	uf := NewUnionFind(graph.vertices)
	var mst []Edge

	for _, edge := range graph.edges {
		operations++
		if uf.Union(edge.u, edge.v) {
			operations++
			mst = append(mst, edge)
		}
	}

	return mst, operations
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func Prim(graph Graph, start int) ([]Edge, int) {
	operations := 0

	visited := make([]bool, graph.vertices)
	pq := &PriorityQueue{}
	heap.Init(pq)
	operations++

	var mst []Edge

	visit := func(v int) {
		visited[v] = true
		operations++

		for _, edge := range graph.edges {
			operations++

			if edge.u == v && !visited[edge.v] {
				operations++
				heap.Push(pq, edge)
			} else if edge.v == v && !visited[edge.u] {
				operations++
				heap.Push(pq, Edge{edge.v, edge.u, edge.weight})
			}
		}
	}

	visit(start)

	for pq.Len() > 0 {
		operations++
		edge := heap.Pop(pq).(Edge)
		operations++

		if visited[edge.v] {
			continue
		}

		mst = append(mst, edge)
		operations++

		visit(edge.v)
	}

	return mst, operations
}

func main() {
	graph := Graph{
		vertices: 10,
		edges: []Edge{
			{0, 1, 2}, {1, 2, 7}, {0, 2, 6}, {1, 3, 9}, {2, 3, 4},
			{3, 4, 5}, {3, 5, 3}, {4, 5, 4}, {4, 6, 2}, {6, 7, 10},
			{6, 8, 5}, {7, 8, 9}, {8, 9, 10},
		},
	}

	start := time.Now()
	kruskalMST, kruskalOperations := Kruskal(graph)
	elapsed := time.Since(start)

	fmt.Println("Алгоритм Крускала")
	fmt.Println("MST:", kruskalMST)
	fmt.Println("Кількість операцій:", kruskalOperations)
	fmt.Println("Час:", elapsed)

	start = time.Now()
	primMST, primOperations := Prim(graph, 0)
	elapsed = time.Since(start)

	fmt.Println("\nАлгоритм Прима")
	fmt.Println("MST:", primMST)
	fmt.Println("Кількість операцій:", primOperations)
	fmt.Println("Час:", elapsed)
}
