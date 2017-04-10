/*
*Graph* Depth-first search solution

For equation A / B = K, just like it has a edge from vertex A to B, this edge's
weight is k, and also an reverse edge from vertex B to A, it's weight is 1.0/k.

For query A / C = ?, what we need to do exactly is find a path from vertex A to
C.
*/
package main

import "fmt"

type Edge struct {
	next   string
	weight float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	adjacency_list := make(map[string][]Edge)

	// build adjacency list
	for i := 0; i < len(equations); i++ {
		e := equations[i]
		x := e[0]
		y := e[1]
		k := values[i]

		if adjacency_list[x] == nil {
			adjacency_list[x] = []Edge{}
		}
		// From X to Y
		adjacency_list[x] = append(adjacency_list[x], Edge{
			next:   y,
			weight: k,
		})

		if adjacency_list[y] == nil {
			adjacency_list[y] = []Edge{}
		}
		// From Y to X
		adjacency_list[y] = append(adjacency_list[y], Edge{
			next:   x,
			weight: 1.0 / k,
		})
	}

	// fmt.Println(adjacency_list)

	results := []float64{}
	for i := 0; i < len(queries); i++ {
		q := queries[i]

		// find a path from x to y
		x := q[0]
		y := q[1]

		visited := make(map[string]bool)
		parent := make(map[string]Edge)

		if x == y { // same vertex
			if _, is_exist := adjacency_list[x]; is_exist {
				results = append(results, 1.0)
			} else { // vertex doesn't exist
				results = append(results, -1.0)
			}
		} else {
			is_find := dfs(x, y, visited, parent, adjacency_list)
			if is_find {
				// have found a path from x to y
				// traversal this path reverse using "parent" data
				k := float64(1.0)
				v := y
				for {
					p := parent[v]
					k *= p.weight
					if p.next == x {
						break
					}
					v = p.next
				}
				results = append(results, k)
			} else {
				// path not found
				results = append(results, -1.0)
			}
		}

	}
	return results
}

func dfs(source string, target string,
	visited map[string]bool, parent map[string]Edge,
	adjacency_list map[string][]Edge) bool {

	visited[source] = true

	for _, edge := range adjacency_list[source] {
		if visited[edge.next] {
			continue
		}

		parent[edge.next] = Edge{
			source,
			edge.weight,
		}

		if edge.next == target {
			// find it
			return true
		}

		is_find := dfs(edge.next, target, visited, parent, adjacency_list)
		if is_find {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(calcEquation(
		[][]string{{"x", "y"}},
		[]float64{2.0},
		[][]string{{"x", "y"}, {"y", "x"}, {"x", "x"}, {"w", "w"}},
	))
}
