package main

import "fmt"

func min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	optimal := [][]int{}
	for i := 0; i < m+1; i++ {
		optimal = append(optimal, make([]int, n+1))
	}

	for i := 0; i < m+1; i++ {
		optimal[i][0] = i
	}

	for j := 0; j < n+1; j++ {
		optimal[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				optimal[i][j] = optimal[i-1][j-1]
			} else {
				optimal[i][j] = 1 + min(
					optimal[i-1][j], // REMOVE word1
					optimal[i][j-1], // REMOVE word2
				)
			}
		}
	}

	return optimal[m][n]
}

func main() {
	fmt.Println(minDistance("sea", "eat"))
}
