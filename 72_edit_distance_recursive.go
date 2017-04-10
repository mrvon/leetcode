/*
The idea is process all characters one by one staring from either from left or
right sides of both strings.  Let we traverse from right corner, there are two
possibilities for every pair of character being traversed.

	m: Length of str1 (first string)
	n: Length of str2 (second string)

If last characters of two strings are same, nothing much to do. Ignore last
characters and get count for remaining strings. So we recur for lengths m-1 and
n-1.

Else (If last characters are not same), we consider all operations on ‘str1’,
consider all three operations on last character of first string, recursively
compute minimum cost for all three operations and take minimum of three values.

	Insert: Recur for m and n-1
	Remove: Recur for m-1 and n
	Replace: Recur for m-1 and n-1
*/

package main

import "fmt"

func min(x, y, z int) int {
	if x < y && x < z {
		return x
	} else if y < z {
		return y
	} else {
		return z
	}
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	// If first string is empty, the only option is to
	// insert all characters of second string into first
	if m == 0 {
		return n
	}

	// If second string is empty, the only option is to
	// remove all characters of first string
	if n == 0 {
		return m
	}

	// If last characters of two strings are same, nothing
	// much to do. Ignore last characters and get count for
	// remaining strings.
	if word1[m-1] == word2[n-1] {
		return minDistance(word1[0:m-1], word2[0:n-1])
	}

	// If last characters are not same, consider all three
	// operations on last character of first string, recursively
	// compute minimum cost for all three operations and take
	// minimum of three values.
	return 1 + min(
		minDistance(word1, word2[0:n-1]),        // Insert
		minDistance(word1[0:m-1], word2),        // Remove
		minDistance(word1[0:m-1], word2[0:n-1]), // Replace
	)
}

func main() {
	fmt.Println(minDistance("hello", "hlo"))
	fmt.Println(minDistance("hello", "helo"))
}
