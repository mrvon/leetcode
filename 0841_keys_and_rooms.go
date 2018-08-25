package main

import "fmt"

type state struct {
	rooms   [][]int
	visited map[int]bool
}

func visit(s *state, curr int) {
	if s.visited[curr] {
		return
	}
	s.visited[curr] = true
	for _, r := range s.rooms[curr] {
		visit(s, r)
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	s := &state{
		rooms:   rooms,
		visited: make(map[int]bool),
	}
	visit(s, 0)
	return len(s.visited) == len(rooms)
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
