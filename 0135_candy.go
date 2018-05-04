/*
Heap solution, Accepted
*/
package main

import "fmt"

// An Item is something we manage in a priority queue.
type Item struct {
	queue_index int
	candy       int
	rate        int
	index       int // The index of the item in the heap.
}

type Heap []*Item

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].rate <= h[j].rate
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *Heap) Up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func (h *Heap) Down(i int, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // left child
		j2 := j1 + 1
		if j2 < n && !h.Less(j1, j2) {
			j = j2 // = 2*i + 2 // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

func (h *Heap) Push(item *Item) {
	n := h.Len()
	item.index = n
	*h = append(*h, item)

	h.Up(h.Len() - 1)
}

func (h *Heap) Pop() *Item {
	n := h.Len() - 1

	h.Swap(0, n)
	h.Down(0, n)

	item := (*h)[n]
	item.index = -1 // for safety
	*h = (*h)[0:n]

	return item
}

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func candy(ratings []int) int {
	queue := []*Item{}
	heap := &Heap{}
	for i := 0; i < len(ratings); i++ {
		item := &Item{
			queue_index: i,
			candy:       0,
			rate:        ratings[i],
		}
		queue = append(queue, item)
		heap.Push(item)
	}

	total := 0
	for heap.Len() > 0 {
		item := heap.Pop()
		candy := 1
		if item.queue_index-1 >= 0 {
			left := queue[item.queue_index-1]
			if left.rate < item.rate {
				candy = max(candy, left.candy+1)
			}
		}
		if item.queue_index+1 < len(queue) {
			right := queue[item.queue_index+1]
			if right.rate < item.rate {
				candy = max(candy, right.candy+1)
			}
		}
		item.candy = candy
		total += candy
	}
	return total
}

func main() {
	fmt.Println(candy([]int{0}))
	fmt.Println(candy([]int{12, 14, 12}))
	fmt.Println(candy([]int{12, 14, 12, 13}))
	fmt.Println(candy([]int{12, 12, 12, 13}))
	fmt.Println(candy([]int{12, 12, 12, 12}))
}
