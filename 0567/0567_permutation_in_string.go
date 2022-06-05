// Slide window
package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	m := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		m[s1[i]]--
	}

	for i := 0; i < len(s2); i++ {
		if i < len(s1) {
			m[s2[i]]++
			if m[s2[i]] == 0 {
				delete(m, s2[i])
			}
		} else {
			m[s2[i]]++
			if m[s2[i]] == 0 {
				delete(m, s2[i])
			}
			m[s2[i-len(s1)]]--
			if m[s2[i-len(s1)]] == 0 {
				delete(m, s2[i-len(s1)])
			}
		}
		if len(m) == 0 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}
