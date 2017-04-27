/*
Dynamic programming

LPS(S[b:e]) ->
	1                                 when b == e
	1                                 when b+1 == e && S[b] != S[e]
	2                                 when b+1 == e && S[b] == S[e]
	2 + LPS(S[b+1:e-1])               when b+1 < e && S[b] == S[e]
	Max(LPS(S[b:e-1]), LPS(S[b+1:e])) when b+1 < e && S[b] != S[e]
*/
package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func longestPalindromeSubseq(s string) int {
	n := len(s)

	optimal := [][]int{}
	for i := 0; i < n; i++ {
		optimal = append(optimal, make([]int, n))
	}

	for l := 1; l <= n; l++ {
		for b := 0; b < n-l+1; b++ {
			e := b + l - 1
			if l == 1 {
				optimal[b][e] = 1
			} else if l == 2 {
				if s[b] == s[e] {
					optimal[b][e] = 2
				} else {
					optimal[b][e] = 1
				}
			} else {
				if s[b] == s[e] {
					optimal[b][e] = 2 + optimal[b+1][e-1]
				} else {
					optimal[b][e] = max(optimal[b][e-1], optimal[b+1][e])
				}
			}
		}
	}

	return optimal[0][n-1]
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
	fmt.Println(longestPalindromeSubseq("jk"))
}
