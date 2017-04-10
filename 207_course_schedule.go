// Topological sort using Kahn's algorithm
package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	is_exist, _ := tps(numCourses, prerequisites)
	return is_exist
}

func tps(V int, edges [][]int) (bool, []int) {
	L := []int{}             // topological sort result
	S := []int{}             // 0 in-degrees vertex set
	A := make(map[int][]int) // adjacency list
	D := make(map[int]int)   // in-degrees

	// build adjacency list, calculate in-degrees
	for i := 0; i < len(edges); i++ {
		e := edges[i]
		// first take course e[1], then e[0], a.k.a, first x then y
		x := e[1]
		y := e[0]

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
	is_exist := canFinish(6, [][]int{
		{5, 2},
		{5, 0},
		{4, 0},
		{4, 1},
		{2, 3},
		{3, 1},
	})
	fmt.Println(is_exist)

	is_exist = canFinish(2, [][]int{
		{1, 0},
		{0, 1},
	})
	fmt.Println(is_exist)
}
