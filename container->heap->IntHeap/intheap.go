package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receiver because they modify the slice's length
// not just it's contents.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	fmt.Println(old)
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
This example inserts several ints into an IntHeap, checks the minimum, and
removes them in order of priority

Another way of arranging int's in golang
*/
func main() {
	h := &IntHeap{100, 20, 50, 70, 22}
	heap.Init(h)
	heap.Push(h, 9)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
