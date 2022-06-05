/* Dynamic Programming

optimal[i][j] is the cost for s1.substr(0, i) and s2.substr(0, j).
Note s1[i], s2[j] not included in the substring.

Base case: optimal[0][0] = 0
Target: optimal[m][n]

if s1[i-1] = s2[j-1] // no deletion
    optimal[i][j] = optimal[i-1][j-1];
else // delete either s1[i-1] or s2[j-1]
    optimal[i][j] = min(optimal[i-1][j]+s1[i-1], optimal[i][j-1]+s2[j-1]);
*/
package main

import "fmt"

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minimumDeleteSum(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)

	optimal := [][]int{}
	for i := 0; i < m+1; i++ {
		optimal = append(optimal, make([]int, n+1))
	}

	optimal[0][0] = 0

	for i := 1; i < m+1; i++ {
		optimal[i][0] = optimal[i-1][0] + int(s1[i-1])
	}

	for j := 1; j < n+1; j++ {
		optimal[0][j] = optimal[0][j-1] + int(s2[j-1])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				optimal[i][j] = optimal[i-1][j-1]
			} else {
				optimal[i][j] = min(optimal[i-1][j]+int(s1[i-1]), optimal[i][j-1]+int(s2[j-1]))
			}
		}
	}

	return optimal[m][n]
}

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
}
