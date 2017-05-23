// TODO, judge system, may be too strict
package main

import "fmt"

// DFS, time limit exceeded
func killP(children map[int][]int, kill int, results *[]int) {
	(*results) = append(*results, kill)
	list := children[kill]
	for i := 0; i < len(list); i++ {
		killP(children, list[i], results)
	}
}

func killProcessDFS(pid []int, ppid []int, kill int) []int {
	children := make(map[int][]int)

	for i := 0; i < len(pid); i++ {
		children[ppid[i]] = append(children[ppid[i]], pid[i])
	}

	results := []int{}
	killP(children, kill, &results)
	return results
}

type Stack []int

func (S *Stack) len() int {
	return len(*S)
}

func (S *Stack) push(i int) {
	(*S) = append(*S, i)
}

func (S *Stack) pop() int {
	i := (*S)[len(*S)-1]
	(*S) = (*S)[:len(*S)-1]
	return i
}

func (S *Stack) peek() int {
	return (*S)[len(*S)-1]
}

// BFS time limited exceeded
func killProcess(pid []int, ppid []int, kill int) []int {
	children := make(map[int][]int)
	for i := 0; i < len(pid); i++ {
		children[ppid[i]] = append(children[ppid[i]], pid[i])
	}

	results := []int{}

	var s Stack
	s.push(kill)

	for s.len() > 0 {
		k := s.pop()
		results = append(results, k)
		list := children[k]
		for i := 0; i < len(list); i++ {
			s.push(list[i])
		}
	}

	return results
}

func main() {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}
	fmt.Println(killProcess(pid, ppid, 0))
	fmt.Println(killProcess(pid, ppid, 3))
	fmt.Println(killProcess(pid, ppid, 5))
	fmt.Println(killProcess(pid, ppid, 1))
	fmt.Println(killProcess(pid, ppid, 10))
}
