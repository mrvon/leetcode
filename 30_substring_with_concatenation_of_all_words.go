package main

import "fmt"

// Slide window
func findSubstring(s string, words []string) []int {
	results := []int{}
	length := len(words[0])
	total := length * len(words)

	counts := make(map[string]int)
	for i := 0; i < len(words); i++ {
		counts[words[i]]++
	}

	for start := 0; start < length; start++ {
		seen := make(map[string]int)

		// window[prev, next]
		// tricky is that when prev is equal than next
		// seen[ss] may be minus value to constrain next never great than prev

		prev := start
		for next := prev + length; next <= len(s); {
			ss := s[next-length : next]
			if seen[ss] < counts[ss] {
				// OK, go on.
				seen[ss]++
				if next-prev == total {
					results = append(results, prev)
				}
				next += length
			} else {
				// FAILED, slide window.
				head := s[prev : prev+length]
				seen[head]--
				prev += length
			}
		}
	}

	return results
}

// Naive solution, very slow.
func naiveFindSubstring(s string, words []string) []int {
	results := []int{}
	counts := make(map[string]int)

	for i := 0; i < len(words); i++ {
		counts[words[i]]++
	}

	l := len(words[0])
	t := len(words) * l

	for i := 0; i < len(s); i++ {
		seen := make(map[string]int)

		j := 0
		for i+j+l <= len(s) {
			ss := s[i+j : i+j+l]
			if counts[ss] == 0 {
				break
			}
			seen[ss]++
			if seen[ss] > counts[ss] {
				break
			}
			j += l
		}

		if j == t {
			results = append(results, i)
		}
	}

	return results
}

func main() {
	fmt.Println(findSubstring("barfoothethethethethefkkfoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
	fmt.Println(naiveFindSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
}
