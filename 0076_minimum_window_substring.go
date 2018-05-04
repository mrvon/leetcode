package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	window := make(map[byte]int)
	tset := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		tset[t[i]]++
	}

	// window range is [slow, fast]
	slow := 0
	fast := 0
	counter := 0
	window_left := 0
	window_length := math.MaxInt32
	for ; fast < len(s); fast++ {
		b := s[fast]
		if tset[b] > 0 { // this byte in string `t`
			if window[b] < tset[b] { // window is not enough
				counter++
			}
			window[b]++
		}

		if counter >= len(t) { // window is enough now
			// try to shrink window
			for {
				b := s[slow]
				if tset[b] == 0 { // b is not in string `t`
					slow++
				} else if window[b] > tset[b] { // window has too many byte `b`
					slow++
					window[b]--
				} else { // can't shrink now
					break
				}
			}
			new_length := fast - slow + 1
			if new_length < window_length {
				window_length = new_length
				window_left = slow
			}
		}
	}

	if window_length == math.MaxInt32 {
		return ""
	} else {
		return s[window_left : window_left+window_length]
	}
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
