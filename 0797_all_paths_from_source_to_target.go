// DFS

package main

import "fmt"

type State struct {
	graph   [][]int
	visited map[int]bool
	current int
	target  int
	trace   []int
	paths   [][]int
}

func dfs(s *State) {
	current := s.trace[len(s.trace)-1]

	if current == s.target {
		path := make([]int, len(s.trace))
		copy(path, s.trace)
		s.paths = append(s.paths, path)
		return
	}

	for i := 0; i < len(s.graph[current]); i++ {
		to := s.graph[current][i]
		if !s.visited[to] {
			s.visited[to] = true
			s.trace = append(s.trace, to)
			dfs(s)
			s.trace = s.trace[:len(s.trace)-1]
			s.visited[to] = false
		}
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	s := &State{}

	start := 0

	s.graph = graph
	s.visited = make(map[int]bool)
	s.visited[start] = true
	s.trace = append(s.trace, start)
	s.target = len(graph) - 1

	dfs(s)

	return s.paths
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
