/* Recursive solution */
package main

import "fmt"

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

func partition(s string) [][]string {
	results := [][]string{}

	for i := len(s) - 1; i >= 0; i-- {
		if isPalindrome(s[i:]) {
			sub_results := partition(s[:i])
			if len(sub_results) == 0 {
				new := []string{}
				new = append(new, s[i:])
				results = append(results, new)
			} else {
				for j := 0; j < len(sub_results); j++ {
					new := append(sub_results[j], s[i:])
					results = append(results, new)
				}
			}
		}
	}

	return results
}

func main() {
	fmt.Println(partition("a"))
	fmt.Println(partition("aa"))
	fmt.Println(partition("aab"))
	fmt.Println(partition("aabb"))
}
