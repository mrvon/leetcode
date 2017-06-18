/*
Dijkstra Algorithm with Heap

find all shortest paths
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
	prev     []*Item
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

func remove_begin(word_list []string, begin string) []string {
	word_map := make(map[string]bool)
	for i := 0; i < len(word_list); i++ {
		word_map[word_list[i]] = true
	}
	delete(word_map, begin)
	new_list := []string{}
	for word := range word_map {
		new_list = append(new_list, word)
	}
	return new_list
}

func dijkstra(begin string, end string, word_list []string) [][]string {
	adjacency := make(map[string][]string)
	item_map := make(map[string]*Item)
	open_set := &Heap{}

	word_list = remove_begin(word_list, begin)

	item := &Item{
		word:     begin,
		distance: 1,
	}
	open_set.Push(item)
	item_map[item.word] = item

	for i := 0; i < len(word_list); i++ {
		word := word_list[i]
		item := &Item{
			word:     word,
			distance: math.MaxInt32,
		}
		open_set.Push(item)
		item_map[word] = item
	}

	word_list = append(word_list, begin)
	for i := 0; i < len(word_list); i++ {
		for j := i + 1; j < len(word_list); j++ {
			x := word_list[i]
			y := word_list[j]

			if !is_adjacent(x, y) {
				continue
			}

			adjacency[x] = append(adjacency[x], y)
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
				// throw away old prev
				item.prev = []*Item{min_item}
			} else if item.distance == min_item.distance+1 {
				// same distance just append
				item.prev = append(item.prev, min_item)
			}
		}
	}

	path := []string{}
	paths := [][]string{}
	end_item := item_map[end]

	if end_item != nil && end_item.distance < math.MaxInt32 {
		buildPath(&paths, &path, begin, end_item)
	}
	return paths
}

func buildPath(paths *[][]string, path *[]string, begin string, end_item *Item) {
	if len(end_item.prev) == 0 {
		np := []string{}
		np = append(np, begin)
		for i := len(*path) - 1; i >= 0; i-- {
			np = append(np, (*path)[i])
		}
		(*paths) = append(*paths, np)
	} else {
		(*path) = append(*path, end_item.word)
		for i := 0; i < len(end_item.prev); i++ {
			item := end_item.prev[i]
			buildPath(paths, path, begin, item)
		}
		(*path) = (*path)[:len(*path)-1]
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	return dijkstra(beginWord, endWord, wordList)
}

func main() {
	fmt.Println(findLadders("hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}),
	)

	fmt.Println(findLadders("kit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}),
	)

	fmt.Println(findLadders("hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log"}),
	)

	fmt.Println(findLadders("a", "c",
		[]string{"a", "b", "c"}),
	)
}
