// O(N)
//
// first count the number of 1 or 0 grouped consecutively,
// so transform "01100" to [1,2,2]
// count of substrings will be min(1,2) + min(2,2) = 3
package main

import "fmt"

func transform(s string) []int {
	n := []int{}

	curr := s[0]
	count := 1

	for i := 1; i < len(s); i++ {
		if curr == s[i] {
			count++
		} else {
			n = append(n, count)
			curr = s[i]
			count = 1
		}
	}
	n = append(n, count)

	return n
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func countBinarySubstrings(s string) int {
	n := transform(s)

	count := 0

	for i := 1; i < len(n); i++ {
		count += min(n[i-1], n[i])
	}

	return count
}

func main() {
	fmt.Println(countBinarySubstrings("1"))
	fmt.Println(countBinarySubstrings("10"))
	fmt.Println(countBinarySubstrings("101"))
	fmt.Println(countBinarySubstrings("1010"))
	fmt.Println(countBinarySubstrings("0011"))
	fmt.Println(countBinarySubstrings("1011"))
	fmt.Println(countBinarySubstrings("10101"))
	fmt.Println(countBinarySubstrings("101010"))
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("100111001"))
}
