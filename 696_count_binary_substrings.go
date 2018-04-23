// Time limit exceeded
package main

import "fmt"

func pattern(l []byte) int {
	for i := 1; i < len(l); i++ {
		if l[i] != l[0] {
			return -1
		}
	}
	if l[0] == '0' {
		return 0
	} else if l[0] == '1' {
		return 1
	} else {
		return -1
	}
}

func cbs(s []byte) int {
	count := 0
	for size := 1; size*2 <= len(s); size++ {
		for start := 0; start <= len(s)-size*2; start++ {
			l := s[start : start+size]
			r := s[start+size : start+size*2]
			pl := pattern(l)
			if pl == -1 {
				continue
			}
			pr := pattern(r)
			if pr == -1 {
				continue
			}
			if pl == pr {
				continue
			}
			count++
		}
	}
	return count
}

func countBinarySubstrings(s string) int {
	return cbs([]byte(s))
}

func main() {
	fmt.Println(countBinarySubstrings("10"))
	fmt.Println(countBinarySubstrings("1010"))
	fmt.Println(countBinarySubstrings("10101"))
	fmt.Println(countBinarySubstrings("101010"))
	fmt.Println(countBinarySubstrings("00110011"))
}
