/*
DFS + Memoized(Get rid of TLE)
*/
package main

import "fmt"

func word(s string, wordMap map[string]bool, resultsMap map[string][]string) []string {
	results := []string{}

	if len(s) == 0 {
		return results
	}

	if r, is_exist := resultsMap[s]; is_exist {
		return r
	}

	for start := len(s) - 1; start >= 0; start-- {
		ss := s[start:]
		if wordMap[ss] {
			if start == 0 {
				results = append(results, ss)
			} else {
				sub_results := word(s[:start], wordMap, resultsMap)
				for i := 0; i < len(sub_results); i++ {
					results = append(results, sub_results[i]+" "+ss)
				}
			}
		}
	}

	resultsMap[s] = results

	return results
}

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
	}

	resultsMap := make(map[string][]string)

	return word(s, wordMap, resultsMap)
}

func printList(l []string) {
	for i := 0; i < len(l); i++ {
		fmt.Println(l[i])
	}
}

func main() {
	printList(wordBreak(
		"ab",
		[]string{"a", "b"},
	))

	printList(wordBreak(
		"abc",
		[]string{"a", "b", "c"},
	))

	printList(wordBreak(
		"catsanddog",
		[]string{"cat", "cats", "and", "sand", "dog"},
	))

	printList(wordBreak(
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
	))

	printList(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		[]string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa", "ba"},
	))
}
