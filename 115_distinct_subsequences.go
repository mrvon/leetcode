package main

import "fmt"

type M struct {
	s int
	t int
}

// Recursive solution with memoized, 262ms
func distinct(memoized map[M]int, s string, s_beg int, t string, t_beg int) int {
	if t_beg == len(t) {
		// t is empty, this is a solution
		return 1
	} else if s_beg == len(s) {
		// s is empty, t is not empty, this is not a solution
		return 0
	}

	if m, has := memoized[M{s_beg, t_beg}]; has {
		return m
	}

	select_case := 0
	ignore_case := 0
	if s[s_beg] == t[t_beg] {
		select_case = distinct(memoized, s, s_beg+1, t, t_beg+1)
		ignore_case = distinct(memoized, s, s_beg+1, t, t_beg)
	} else {
		ignore_case = distinct(memoized, s, s_beg+1, t, t_beg)
	}
	memoized[M{s_beg, t_beg}] = select_case + ignore_case
	return select_case + ignore_case
}

func recursiveDistinct(s string, t string) int {
	return distinct(make(map[M]int), s, 0, t, 0)
}

func numDistinct(s string, t string) int {
	optimal := [][]int{}

	for i := 0; i <= len(s); i++ {
		optimal = append(optimal, make([]int, len(t)+1))
	}

	// t is empty, this is a solution
	for si := 0; si < len(s); si++ {
		optimal[si][len(t)] = 1
	}

	// s is empty, t is not empty, this is not a solution
	for ti := 0; ti < len(t)-1; ti++ {
		optimal[len(si)][ti] = 0
	}

	for si := 1; si < len(s); si++ {
		for ti := 1; ti < len(t); ti++ {
			// if s[]
			// optimal[si][ti] =
		}
	}

	return optimal[0][0]
}

func main() {
	fmt.Println(numDistinct("", "a"))
	fmt.Println(numDistinct("rabbit", "rabbit"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("rabbbbit", "rabbit"))
	fmt.Println(numDistinct("ccc", "c"))
}
