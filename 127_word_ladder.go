/*
Dijkstra Algorithm with Heap
*/
package main

import (
	"fmt"
	"math"
)

// An Item is something we manage in a priority queue.
type Item struct {
	word     string
	distance int
	index    int // The index of the item in the heap.
}

type Heap []*Item

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].distance <= h[j].distance
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

func (h *Heap) Fix(i int) {
	h.Down(i, h.Len())
	h.Up(i)
}

func (h *Heap) Update(item *Item) {
	h.Fix(item.index)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	return dijkstra(beginWord, endWord, wordList)
}

func is_adjacent(word1 string, word2 string) bool {
	// word1, word2 always has exact same number of characters.
	// word1 will not equal to word2
	counter := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			counter++
			if counter > 1 {
				return false
			}
		}
	}
	return counter == 1
}

func dijkstra(begin string, end string, word_list []string) int {
	adjacency := make(map[string][]string)
	item_map := make(map[string]*Item)
	open_set := &Heap{}

	item := &Item{
		word:     begin,
		distance: 1,
	}
	open_set.Push(item)
	item_map[item.word] = item

	for i := 0; i < len(word_list); i++ {
		item := &Item{
			word:     word_list[i],
			distance: math.MaxInt32,
		}
		open_set.Push(item)
		item_map[item.word] = item
	}

	word_list = append(word_list, begin)
	for i := 0; i < len(word_list); i++ {
		for j := i + 1; j < len(word_list); j++ {
			x := word_list[i]
			y := word_list[j]

			if !is_adjacent(x, y) {
				continue
			}

			if _, is_exist := adjacency[x]; !is_exist {
				adjacency[x] = []string{}
			}
			adjacency[x] = append(adjacency[x], y)

			if _, is_exist := adjacency[y]; !is_exist {
				adjacency[y] = []string{}
			}
			adjacency[y] = append(adjacency[y], x)
		}
	}

	for open_set.Len() > 0 {
		min_item := open_set.Pop()

		for _, word := range adjacency[min_item.word] {
			item := item_map[word]

			if item.distance > min_item.distance+1 {
				item.distance = min_item.distance + 1
				open_set.Update(item)
			}
		}
	}

	if item_map[end] == nil {
		return 0
	} else {
		d := item_map[end].distance
		if d == math.MaxInt32 {
			return 0
		} else {
			return d
		}
	}
}

func main() {
	fmt.Println(ladderLength("hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}))

	fmt.Println(ladderLength("kit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}))

	fmt.Println(ladderLength("hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log"}))
}
