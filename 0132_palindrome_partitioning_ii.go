/* Dynamic programming */
package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}

func minCut(s string) int {
	optimal := make([]int, len(s))
	for i := 0; i < len(optimal); i++ {
		optimal[i] = math.MaxInt32
	}

	for end := 0; end < len(s); end++ {
		for start := 0; start <= end; start++ {
			if isPalindrome(s[start : end+1]) {
				if start == 0 {
					optimal[end] = 0
				} else {
					optimal[end] = min(optimal[end], optimal[start-1]+1)
				}
			}
		}
	}

	return optimal[len(s)-1]
}

func main() {
	fmt.Println(minCut("a"))
	fmt.Println(minCut("aa"))
	fmt.Println(minCut("aab"))
	fmt.Println(minCut("aabb"))
	fmt.Println(minCut("aacaa"))
}
