// Astar
package main

import "fmt"

type State struct {
	person [2]int
	box    [2]int
}

type HeapItem struct {
	person   [2]int
	box      [2]int
	moves    int
	priority int // manhattan distance between box and target
	index    int // for heap internal use
}

type Heap []*HeapItem

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].priority < h[j].priority
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

func (h *Heap) Push(item *HeapItem) {
	n := h.Len()
	item.index = n
	*h = append(*h, item)
	h.Up(h.Len() - 1)
}

func (h *Heap) Pop() *HeapItem {
	n := h.Len() - 1
	h.Swap(0, n)
	h.Down(0, n)
	item := (*h)[n]
	item.index = -1 // for safety
	*h = (*h)[0:n]
	return item
}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}

func heuristic(box [2]int, target [2]int) int {
	return abs(target[0]-box[0]) + abs(target[1]-box[1])
}

func minPushBox(grid [][]byte) int {
	outBounds := func(p [2]int) bool {
		rows := len(grid)
		cols := len(grid[0])
		if p[0] < 0 || p[0] >= rows {
			return true
		}
		if p[1] < 0 || p[1] >= cols {
			return true
		}
		return grid[p[0]][p[1]] == '#'
	}
	target := [2]int{}
	start_box := [2]int{}
	start_person := [2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			c := grid[i][j]
			if c == 'T' {
				target = [2]int{i, j}
			} else if c == 'B' {
				start_box = [2]int{i, j}
			} else if c == 'S' {
				start_person = [2]int{i, j}
			}
		}
	}
	openSet := Heap{}
	openSet.Push(&HeapItem{
		person:   start_person,
		box:      start_box,
		moves:    0,
		priority: heuristic(start_box, target),
	})
	closeSet := make(map[State]bool)
	direction := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	for openSet.Len() > 0 {
		item := openSet.Pop()
		box := item.box
		person := item.person
		moves := item.moves
		if box == target {
			return moves
		}
		s := State{person: person, box: box}
		if closeSet[s] {
			continue
		}
		closeSet[s] = true
		for _, d := range direction {
			new_person := [2]int{
				person[0] + d[0],
				person[1] + d[1],
			}
			if outBounds(new_person) {
				continue
			}
			if new_person == box {
				new_box := [2]int{box[0] + d[0], box[1] + d[1]}
				if outBounds(new_box) {
					continue
				}
				item := &HeapItem{
					person:   new_person,
					box:      new_box,
					moves:    moves + 1,
					priority: heuristic(new_box, target) + moves + 1,
				}
				openSet.Push(item)
			} else {
				item := &HeapItem{
					person:   new_person,
					box:      box,
					moves:    moves,
					priority: heuristic(box, target) + moves,
				}
				openSet.Push(item)
			}
		}
	}
	return -1
}

func main() {
	grid := [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', 'B', '.', '#'},
		{'#', '.', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'}}
	fmt.Println(minPushBox(grid))
	grid2 := [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', 'B', '.', '#'},
		{'#', '#', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'}}
	fmt.Println(minPushBox(grid2))
	grid3 := [][]byte{
		{'#', '.', '.', '#', '#', '#', '#', '#'},
		{'#', '.', '.', 'T', '#', '.', '.', '#'},
		{'#', '.', '.', '.', '#', 'B', '.', '#'},
		{'#', '.', '.', '.', '.', '.', '.', '#'},
		{'#', '.', '.', '.', '#', '.', 'S', '#'},
		{'#', '.', '.', '#', '#', '#', '#', '#'},
	}
	fmt.Println(minPushBox(grid3))
}
