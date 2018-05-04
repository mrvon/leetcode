package main

import (
	"fmt"
	"strconv"
)

func one(s string) int {
	if s[0] == '0' {
		return 0
	} else {
		return 1
	}
}

func two(s string) int {
	if s[0] == '0' {
		return 0
	} else {
		if i, err := strconv.Atoi(s); err == nil && i <= 26 {
			return 1
		} else {
			return 0
		}
	}
}

func ways(s string, m map[string]int) int {
	n := len(s)
	if n == 0 {
		return 1
	}

	if r, is_exist := m[s]; is_exist {
		return r
	}

	r := 0
	if n == 1 {
		r = one(s[0:1]) * ways(s[1:n], m)
	} else {
		r = one(s[0:1])*ways(s[1:n], m) + two(s[0:2])*ways(s[2:n], m)
	}

	m[s] = r
	return r
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	} else {
		return ways(s, make(map[string]int))
	}
}

func main() {
	fmt.Println(numDecodings(""))    //
	fmt.Println(numDecodings("12"))  // 1,2; 12
	fmt.Println(numDecodings("10"))  // 10
	fmt.Println(numDecodings("181")) // 1,8,1; 18,1
	fmt.Println(numDecodings("101")) // 10,1
	fmt.Println(numDecodings("110")) // 1,10
	fmt.Println(numDecodings("121")) // 1,2,1; 12,1; 1,21
	fmt.Println(numDecodings("6065812287883668764831544958683283296479682877898293612168136334983851946827579555449329483852397155"))

}
