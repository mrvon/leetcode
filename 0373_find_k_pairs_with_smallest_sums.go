package main

import "fmt"

type State struct {
	h      Heap
	i_nums []int
	j_nums []int
}

// An Item is something we manage in a priority queue.
type Item struct {
	i        int
	j        int
	pair_sum int

	// internal
	index int // The index of the item in the heap.
}

type Heap []*Item

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].pair_sum <= h[j].pair_sum
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

func (s *State) nextPairs() []int {
	if s.h.Len() == 0 {
		return []int{}
	}
	item := s.h.Pop()
	pair := []int{s.i_nums[item.i], s.j_nums[item.j]}
	item.j++
	if item.j < len(s.j_nums) {
		item.pair_sum = s.i_nums[item.i] + s.j_nums[item.j]
		s.h.Push(item)
	}
	return pair
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	s := &State{
		i_nums: nums1,
		j_nums: nums2,
	}
	for i := 0; i < len(nums1); i++ {
		j := 0
		if j < len(nums2) {
			pair_sum := nums1[i] + nums2[j]
			s.h.Push(&Item{i: i, j: j, pair_sum: pair_sum})
		}
	}

	results := [][]int{}
	for i := 0; i < k; i++ {
		pair := s.nextPairs()
		if len(pair) > 0 {
			results = append(results, pair)
		}
	}
	return results
}

func main() {
	// Given nums1 = [1,7,11], nums2 = [2,4,6],  k = 3
	// Return: [1,2],[1,4],[1,6]
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 9
	fmt.Println(kSmallestPairs(nums1, nums2, k))

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	fmt.Println(kSmallestPairs(nums1, nums2, k))

	nums1 = []int{1, 2}
	nums2 = []int{3}
	k = 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))

	nums1 = []int{3, 5, 7, 9}
	nums2 = []int{}
	k = 1
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}
