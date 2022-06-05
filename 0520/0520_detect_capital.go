// draw DFA and translate to code
package main

import "fmt"

const (
	Upper = 1
	Lower = 2
)

type Pair struct {
	source_state int
	char_type    int
}

func detectCapitalUse(word string) bool {
	transition_table := make(map[Pair]int)
	transition_table[Pair{1, Upper}] = 3
	transition_table[Pair{1, Lower}] = 2
	transition_table[Pair{2, Lower}] = 2
	transition_table[Pair{3, Upper}] = 4
	transition_table[Pair{3, Lower}] = 2
	transition_table[Pair{4, Upper}] = 4

	accept_state := []int{
		2, 3, 4,
	}

	// initial state
	curr_state := 1

	for i := 0; i < len(word); i++ {
		char_type := 0
		if word[i] >= 'A' && word[i] <= 'Z' {
			char_type = Upper
		} else {
			char_type = Lower
		}

		next_state := transition_table[Pair{curr_state, char_type}]
		if next_state == 0 { // error occurs
			return false
		}
		curr_state = next_state
	}

	for i := 0; i < len(accept_state); i++ {
		if curr_state == accept_state[i] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("leetcode"))
	fmt.Println(detectCapitalUse("China"))
	fmt.Println(detectCapitalUse("FlAG"))
}
