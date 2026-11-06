package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

// to find len
func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	h := &IntHeap{5, 2, 8, 1, 9}

	heap.Init(h)

	heap.Push(h, 3)

	fmt.Println("Smallest:", heap.Pop(h))
	fmt.Println("Next:", heap.Pop(h))
}
