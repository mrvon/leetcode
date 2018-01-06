package main

import "fmt"

func canDecode(bits []int) bool {
	l := len(bits)
	if l == 0 {
		return true
	} else if l == 1 {
		return bits[0] == 0
	} else {
		if bits[l-1] == 1 {
			if bits[l-2] == 1 {
				return canDecode(bits[:l-2])
			} else {
				return false
			}
		} else {
			return canDecode(bits[:l-1]) || // one bit character
				canDecode(bits[:l-2]) // two bit character
		}
	}
}

func isOneBitCharacter(bits []int) bool {
	l := len(bits)
	if l <= 1 {
		return true
	}
	// assert(bits[l-1] == 0)
	if bits[l-2] == 0 {
		return true
	} else {
		return canDecode(bits[:l-1]) // must be one bit character
	}
}

func assert(expect bool, result bool) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %b, Get %b", expect, result))
	}
}

func main() {
	assert(true, canDecode([]int{1, 1}))
	assert(true, canDecode([]int{1, 0}))
	assert(false, canDecode([]int{0, 1}))
	assert(true, canDecode([]int{0, 0}))
	assert(false, canDecode([]int{1, 1, 1}))
	assert(false, canDecode([]int{1}))

	assert(true, isOneBitCharacter([]int{1, 0, 0}))
	assert(false, isOneBitCharacter([]int{1, 1, 1, 0}))
}
