/*
Topological sort

Kahn's algorithm

One of these algorithms, first described by Kahn (1962), works by choosing
vertices in the same order as the eventual topological sort. First, find a list
of "start nodes" which have no incoming edges and insert them into a set S; at
least one such node must exist in a non-empty acyclic graph. Then:

	L ← Empty list that will contain the sorted elements
	S ← Set of all nodes with no incoming edges

	while S is non-empty do
		remove a node n from S
		add n to tail of L
	for each node m with an edge e from n to m do
		remove edge e from the graph
		if m has no other incoming edges then
			insert m into S
	if graph has edges then
		return error (graph has at least one cycle)
	else
		return L (a topologically sorted order)

If the graph is a DAG, a solution will be contained in the list L
(the solution is not necessarily unique). Otherwise, the graph must have at
least one cycle and therefore a topological sorting is impossible.
*/
package main

import "fmt"

func tps(V int, edges [][]int) (bool, []int) {
	L := []int{}             // topological sort result
	S := []int{}             // 0 in-degrees vertex set
	A := make(map[int][]int) // adjacency list
	D := make(map[int]int)   // in-degrees

	// build adjacency list, calculate in-degrees
	for i := 0; i < len(edges); i++ {
		e := edges[i]
		x := e[0]
		y := e[1]

		if _, is_exist := A[x]; !is_exist {
			A[x] = []int{}
		}
		A[x] = append(A[x], y)

		D[y]++
	}

	// zero in-degrees
	for i := 0; i < V; i++ {
		if D[i] == 0 {
			S = append(S, i)
		}
	}

	for len(S) > 0 {
		n := S[0]
		S = S[1:len(S)]
		L = append(L, n)

		for _, y := range A[n] {
			D[y]--
			if D[y] <= 0 {
				S = append(S, y)
			}
		}
	}

	if len(L) == V {
		return true, L
	} else {
		return false, L
	}
}

func main() {
	is_exist, list := tps(6, [][]int{
		{5, 2},
		{5, 0},
		{4, 0},
		{4, 1},
		{2, 3},
		{3, 1},
	})
	fmt.Println(is_exist, list)

	is_exist, list = tps(2, [][]int{
		{1, 0},
		{0, 1},
	})
	fmt.Println(is_exist, list)
}
