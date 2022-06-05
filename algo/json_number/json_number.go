// draw json number DFA and translate to code
package main

import "fmt"

const (
	_ = iota
	Blank
	Zero
	OneNight
	Minus
	Plus
	E
	Dot
)

type Pair struct {
	curr_state int
	char_type  int
}

func valid_json_number(s string) bool {
	transition := make(map[Pair]int)
	transition[Pair{1, Zero}] = 3
	transition[Pair{1, Minus}] = 2
	transition[Pair{1, OneNight}] = 4
	transition[Pair{2, Zero}] = 3
	transition[Pair{2, OneNight}] = 4
	transition[Pair{3, E}] = 8
	transition[Pair{3, Dot}] = 6
	transition[Pair{4, Zero}] = 5
	transition[Pair{4, OneNight}] = 5
	transition[Pair{4, Dot}] = 6
	transition[Pair{4, E}] = 8
	transition[Pair{5, Zero}] = 5
	transition[Pair{5, OneNight}] = 5
	transition[Pair{5, Dot}] = 6
	transition[Pair{6, Zero}] = 7
	transition[Pair{6, OneNight}] = 7
	transition[Pair{7, Zero}] = 7
	transition[Pair{7, OneNight}] = 7
	transition[Pair{7, E}] = 8
	transition[Pair{8, Zero}] = 10
	transition[Pair{8, OneNight}] = 10
	transition[Pair{8, Plus}] = 9
	transition[Pair{8, Minus}] = 9
	transition[Pair{9, Zero}] = 10
	transition[Pair{9, OneNight}] = 10
	transition[Pair{10, Zero}] = 10
	transition[Pair{10, OneNight}] = 10

	accept_state := []int{
		3, 4, 5, 7, 10,
	}

	curr_state := 1

	for i := 0; i < len(s); i++ {
		char_type := 0
		c := s[i]
		if c == ' ' {
			char_type = Blank
		} else if c == '0' {
			char_type = Zero
		} else if c >= '1' && c <= '9' {
			char_type = OneNight
		} else if c == '-' {
			char_type = Minus
		} else if c == '+' {
			char_type = Plus
		} else if c == 'e' || c == 'E' {
			char_type = E
		} else if c == '.' {
			char_type = Dot
		}

		next_state := transition[Pair{curr_state, char_type}]
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

func assert(expect bool, result bool) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %b, Get %b", expect, result))
	}
}

func main() {
	assert(false, valid_json_number("00"))
	assert(true, valid_json_number("0"))
	assert(true, valid_json_number("1"))
	assert(true, valid_json_number("1.1"))
	assert(true, valid_json_number("1.123456"))
	assert(true, valid_json_number("0.123456"))
	assert(true, valid_json_number("10000.123456"))
	assert(true, valid_json_number("0e10"))
	assert(true, valid_json_number("0.1e10"))
	assert(true, valid_json_number("1e10"))
	assert(true, valid_json_number("1.1e10"))
	assert(false, valid_json_number("1.1e10.1"))
	assert(true, valid_json_number("1.1e+10"))
	assert(true, valid_json_number("1.1e-10"))
	assert(false, valid_json_number("+1.1e-10"))
	assert(true, valid_json_number("-1.1e-10"))
	assert(true, valid_json_number("-1.1e+10"))
	assert(true, valid_json_number("-0.1e+10"))
	assert(true, valid_json_number("-0.1e+01"))
	assert(true, valid_json_number("-0.1e-01"))
	assert(true, valid_json_number("-0.1e01"))
	assert(false, valid_json_number("a"))
	assert(false, valid_json_number("0a"))
	assert(false, valid_json_number("0a1"))
	assert(false, valid_json_number("01"))
	assert(false, valid_json_number("0."))
	assert(false, valid_json_number("0.1."))
}
