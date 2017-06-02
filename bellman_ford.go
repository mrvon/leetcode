/*
The Bellman-Ford algorithm solves the single-source shortest-paths problem in
the general case in which edge weights may be negative. Give a weighted,
directed graph G = (V, E) with source s and weight function w: E -> R, the
Bellman-Ford algorithm returns a boolean value indicating whether or not there
is a negative-weight cycle that is reachable from the source. If there is such a
cycle, the algorithm indicates that no solution exists. If there is no such
cycle, the algorithm produces the shortest paths and their weights.

The algorithm relaxes edges, progressively decreasing an estimate v.d on the
weight of a shortest path from the source s to each vertex v in V until it
archieves the actual shortest-path weight. The algorithm return TRUE if and only
if the graph contains no negative-weight cycles that are reachable from the
source.
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	id       int
	estimate int
	parent   *Vertex
}

type Edge struct {
	v_id   int
	weight int
}

type Graph struct {
	adjacency   map[int][]Edge
	vertex_list map[int]*Vertex
}

func (graph *Graph) add_vertex(id int) {
	if graph.vertex_list == nil {
		graph.vertex_list = make(map[int]*Vertex)
	}
	graph.vertex_list[id] = &Vertex{
		id:       id,
		parent:   nil,
		estimate: math.MaxInt32,
	}
}

func (graph *Graph) add_edge(from_id int, to_id int, weight int) {
	if graph.adjacency == nil {
		graph.adjacency = make(map[int][]Edge)
	}
	graph.adjacency[from_id] = append(graph.adjacency[from_id], Edge{
		v_id:   to_id,
		weight: weight,
	})
}

func bellman_ford(source_id int, graph *Graph) bool {
	graph.vertex_list[source_id].estimate = 0

	for i := 0; i < len(graph.vertex_list)-1; i++ {
		for u_id, adj := range graph.adjacency {
			u := graph.vertex_list[u_id]
			for _, edge := range adj {
				v := graph.vertex_list[edge.v_id]
				// Relax
				if u.estimate+edge.weight < v.estimate {
					v.estimate = u.estimate + edge.weight
				}
			}
		}
	}

	// Detect negative weight cycles
	for u_id, adj := range graph.adjacency {
		u := graph.vertex_list[u_id]
		for _, edge := range adj {
			v := graph.vertex_list[edge.v_id]
			if v.estimate > u.estimate+edge.weight {
				// Has negative cycles
				return false
			}
		}
	}

	// Print result
	for _, vertex := range graph.vertex_list {
		fmt.Printf("Vectex(%d) Estimate(%d)\n", vertex.id, vertex.estimate)
	}
	return true
}

func main() {
	graph := &Graph{}

	graph.add_edge(1, 2, 1)
	graph.add_edge(1, 3, 2)
	graph.add_edge(1, 4, 3)
	// graph.add_edge(4, 1, -4)

	graph.add_vertex(1)
	graph.add_vertex(2)
	graph.add_vertex(3)
	graph.add_vertex(4)

	fmt.Printf("Has Negative Cycles: %v\n", !bellman_ford(1, graph))
}
