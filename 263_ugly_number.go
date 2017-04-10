package main

import "fmt"

func isUgly(num int) bool {
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			if num == 1 {
				return true
			} else {
				return false
			}
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
	assert(true, isUgly(1))
	assert(true, isUgly(2))
	assert(true, isUgly(3))
	assert(true, isUgly(4))
	assert(true, isUgly(5))
	assert(true, isUgly(6))
	assert(false, isUgly(7))
	assert(true, isUgly(8))
	assert(true, isUgly(9))
	assert(true, isUgly(10))
	assert(false, isUgly(11))
	assert(true, isUgly(12))
	assert(false, isUgly(13))
	assert(false, isUgly(14))
}
