/* Dynamic programming */
package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	valid := make(map[int]bool)

	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
	}

	for end := 0; end < len(s); end++ {
		for start := end; start >= 0; start-- {
			if wordMap[s[start:end+1]] {
				if start == 0 || valid[start-1] {
					valid[end] = true
					break
				}
			}
		}
	}

	return valid[len(s)-1]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("leetcode", []string{"le", "co", "de", "et"}))
	fmt.Println(wordBreak("leetcode", []string{"le", "co", "de", "ea", "ee"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "codk"}))
	fmt.Println(wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
}
