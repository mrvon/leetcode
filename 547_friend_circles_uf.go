// Union find
package main

import "fmt"

type UF struct {
	id    []int // parent link (site indexed)
	sz    []int // size of component for roots (site indexed)
	count int   // number of components
}

func new_uf(n int) *UF {
	uf := &UF{}
	uf.id = make([]int, n)
	uf.sz = make([]int, n)
	for i := 0; i < len(uf.id); i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}
	uf.count = n
	return uf
}

func (uf *UF) connected(p int, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UF) find(p int) int {
	// Follow links to find a root.
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *UF) union(p int, q int) {
	// Give p and q the same root.
	p_root := uf.find(p)
	q_root := uf.find(q)

	if p_root == q_root {
		return
	}

	// Make smaller root point to larger one.
	if uf.sz[p_root] < uf.sz[q_root] {
		// p links to q
		uf.id[p_root] = q_root
		uf.sz[q_root] += uf.sz[p_root]
	} else {
		// q links to p
		uf.id[q_root] = p_root
		uf.sz[p_root] += uf.sz[q_root]
	}

	uf.count--
}

func findCircleNum(M [][]int) int {
	n := len(M)
	uf := new_uf(n)

	for i := 0; i < len(M); i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}

	return uf.count
}

func main() {
	matrix := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 1},
	}
	fmt.Println(findCircleNum(matrix))

	matrix = [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{0, 1, 1},
	}
	fmt.Println(findCircleNum(matrix))
}
