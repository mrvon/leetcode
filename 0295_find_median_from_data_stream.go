package main

import (
	"fmt"
)

const (
	MinHeap = 0
	MaxHeap = 1
)

// An Item is something we manage in a priority queue.
type Item struct {
	number int
	index  int // The index of the item in the heap.
}

type Heap struct {
	t int
	h []*Item
}

func (h Heap) Len() int {
	return len(h.h)
}

func (h Heap) Less(i, j int) bool {
	if h.t == MinHeap {
		return h.h[i].number <= h.h[j].number
	} else {
		return h.h[i].number >= h.h[j].number
	}
}

func (h Heap) Swap(i, j int) {
	h.h[i], h.h[j] = h.h[j], h.h[i]
	h.h[i].index = i
	h.h[j].index = j
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
	h.h = append(h.h, item)

	h.Up(h.Len() - 1)
}

func (h *Heap) Pop() *Item {
	n := h.Len() - 1

	h.Swap(0, n)
	h.Down(0, n)

	item := h.h[n]
	item.index = -1 // for safety
	h.h = h.h[0:n]

	return item
}

func (h *Heap) Top() *Item {
	return h.h[0]
}

func (h *Heap) Fix(i int) {
	h.Down(i, h.Len())
	h.Up(i)
}

func (h *Heap) Update(item *Item) {
	h.Fix(item.index)
}

func assert(b bool) {
	if !b {
		panic("assert failed!")
	}
}

type MedianFinder struct {
	left  *Heap
	mid   []int
	right *Heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{
		left:  &Heap{t: MaxHeap, h: []*Item{}},
		mid:   []int{},
		right: &Heap{t: MinHeap, h: []*Item{}},
	}
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	l := this.left.Len()
	r := this.right.Len()

	if l < r {
		rt := this.right.Top().number
		if num <= rt {
			this.left.Push(&Item{number: num})
		} else {
			this.left.Push(this.right.Pop())
			this.right.Push(&Item{number: num})
		}
	} else if l > r {
		lt := this.left.Top().number
		if num >= lt {
			this.right.Push(&Item{number: num})
		} else {
			this.right.Push(this.left.Pop())
			this.left.Push(&Item{number: num})
		}
	} else {
		if l == 0 {
			this.right.Push(&Item{number: num})
		} else {
			lt := this.left.Top().number

			if num <= lt {
				this.left.Push(&Item{number: num})
			} else {
				this.right.Push(&Item{number: num})
			}
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	l := this.left.Len()
	r := this.right.Len()

	if l > r {
		return float64(this.left.Top().number)
	} else if l < r {
		return float64(this.right.Top().number)
	} else {
		if l == 0 {
			return 0
		} else {
			return float64(this.left.Top().number+this.right.Top().number) / 2
		}
	}
}

func main() {
	mf := Constructor()
	mf.AddNum(1)
	fmt.Println(mf.FindMedian())
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(3)
	fmt.Println(mf.FindMedian())
	mf.AddNum(4)
	fmt.Println(mf.FindMedian())
	mf.AddNum(5)
	fmt.Println(mf.FindMedian())
	mf.AddNum(6)
	fmt.Println(mf.FindMedian())
}
