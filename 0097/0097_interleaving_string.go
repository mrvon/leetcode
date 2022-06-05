package main

import "fmt"

// Naive recursive solution, 1300ms
func isInterleaveR(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 {
		return s2 == s3
	} else if len(s2) == 0 {
		return s1 == s3
	} else if len(s3) == 0 {
		return false
	}

	if s1[0] == s3[0] && isInterleaveR(s1[1:], s2, s3[1:]) {
		return true
	}
	if s2[0] == s3[0] && isInterleaveR(s1, s2[1:], s3[1:]) {
		return true
	}

	return false
}

/*
2D Dynamic Programming

DP table represents if s3 is interleaving at (i+j)-th position when s1 is at
i-th position, and s2 is at j-th position. 0-th position means empty string.

	1. If both s1 and s2 is currently empty, s3 is empty too, and it is
	considered interleaving.

	2. If only s1 is empty, then if previous s2 position is interleaving and
	current s2 position char is equal to s3 current position char, it is
	considered interleaving. similar idea applies to when s2 is empty.

	3. when both s1 and s2 is not empty, then if we arrive i, j from i-1, j,
	then if i-1,j is already interleaving and i and current s3 position equal,
	it is interleaving. If we arrive i, j from i, j-1, then if i, j-1 is already
	interleaving and j and current s3 position equal. it is interleaving.
*/
func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	k := len(s3)

	if m+n != k {
		return false
	}

	optimal := [][]bool{}
	for i := 0; i < m+1; i++ {
		optimal = append(optimal, make([]bool, n+1))
	}

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 && j == 0 {
				optimal[i][j] = true
			} else if i == 0 {
				optimal[i][j] = optimal[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				optimal[i][j] = optimal[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				optimal[i][j] = (optimal[i][j-1] && s2[j-1] == s3[i+j-1]) ||
					(optimal[i-1][j] && s1[i-1] == s3[i+j-1])
			}
		}
	}

	return optimal[m][n]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))

	// Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

	// For example,
	// Given:
	// s1 = "aabcc",
	// s2 = "dbbca",

	// When s3 = "aadbbcbcac", return true.
	// When s3 = "aadbbbaccc", return false.

}
