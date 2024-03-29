package main

import "fmt"

func romanToInt(s string) int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	i := len(s) - 1
	sum := m[s[i]]
	i--
	for ; i >= 0; i-- {
		if m[s[i]] < m[s[i+1]] {
			sum -= m[s[i]]
		} else {
			sum += m[s[i]]
		}

	}
	return sum
}

func assert(expect int, result int) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %d, Get %d", expect, result))
	}
}

func main() {
	assert(romanToInt("III"), 3)
	assert(romanToInt("IV"), 4)
}
