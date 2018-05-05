package main

import "fmt"

func countArrangement(N int) int {
	Set := make([]int, N+1)
	return count(1, N, Set)
}

func count(S int, N int, Set []int) int {
	if S > N {
		// found a beautiful arrangement
		return 1
	}

	bc := 0
	for i := 1; i <= N; i++ {
		// try to place number i in S-th slot
		if Set[i] == 1 { // number i have been use
			continue
		} else if i > S && i%S != 0 { // check divisible
			continue
		} else if S > i && S%i != 0 { // check divisible
			continue
		} else {
			// ok, place it
			Set[i] = 1
			bc += count(S+1, N, Set)
			Set[i] = 0
		}
	}
	return bc
}

func main() {
	arr := []int{}
	for i := 1; i <= 15; i++ {
		arr = append(arr, countArrangement(i))
	}
	fmt.Println(arr)
}
