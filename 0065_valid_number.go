/*
Draw DFA diagram before you write program.

DFA structure is that:
	(source_state, current_input_char) -> target_state

The process is easy, just feed dfa machine with input character, if error occurs
in the middle of process or end with a unacceptable state, the string is a
invalid number, otherwise is a valid number.
*/
package main

import "fmt"

type Pair struct {
	source_state int
	char_type    string
}

func isNumber(s string) bool {
	// build transition table
	transition_table := make(map[Pair]int)
	transition_table[Pair{1, "blank"}] = 1
	transition_table[Pair{1, "sign"}] = 2
	transition_table[Pair{1, "digit"}] = 3
	transition_table[Pair{1, "dot"}] = 4
	transition_table[Pair{2, "digit"}] = 3
	transition_table[Pair{2, "dot"}] = 4
	transition_table[Pair{3, "digit"}] = 3
	transition_table[Pair{3, "dot"}] = 5
	transition_table[Pair{3, "e"}] = 6
	transition_table[Pair{3, "blank"}] = 9
	transition_table[Pair{4, "digit"}] = 5
	transition_table[Pair{5, "digit"}] = 5
	transition_table[Pair{5, "e"}] = 6
	transition_table[Pair{5, "blank"}] = 9
	transition_table[Pair{6, "sign"}] = 7
	transition_table[Pair{6, "digit"}] = 8
	transition_table[Pair{7, "digit"}] = 8
	transition_table[Pair{8, "digit"}] = 8
	transition_table[Pair{8, "blank"}] = 9
	transition_table[Pair{9, "blank"}] = 9

	accept_table := []int{
		3, 5, 8, 9,
	}

	// initial state
	current_state := 1

	// feed character for dfa machine
	for i := 0; i < len(s); i++ {
		c := s[i]

		char_type := "unknown"
		if c >= '0' && c <= '9' {
			char_type = "digit"
		} else if c == ' ' {
			char_type = "blank"
		} else if c == '+' || c == '-' {
			char_type = "sign"
		} else if c == '.' {
			char_type = "dot"
		} else if c == 'e' {
			char_type = "e"
		}

		target_state := transition_table[Pair{current_state, char_type}]
		if target_state == 0 { // error occurs
			return false
		}

		// state transition
		current_state = target_state
	}

	// check is it accept?
	for i := 0; i < len(accept_table); i++ {
		if accept_table[i] == current_state {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isNumber("1"))
	fmt.Println(isNumber("a1"))
	fmt.Println(isNumber("1a"))
	fmt.Println(isNumber("7j1"))
}
