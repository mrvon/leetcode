package main

import "fmt"

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if len(A) == 0 {
		return true
	}

	pivot := byte(A[0])

	for i, r := range B {
		match := true
		c := byte(r)
		if pivot == c {
			// try match
			k := i
			for j := 0; j < len(A); j++ {
				if A[j] != B[k] {
					// no match, try next
					match = false
					break
				}
				k++
				if k >= len(B) {
					k = 0
				}
			}
			if match {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(rotateString("", ""))
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("abcde", "abced"))
	fmt.Println(rotateString("abcde", "cdeaab"))
	fmt.Println(rotateString(
		"bqqutquvbtgouklsayfvzewpnrbwfcdmwctusunasdbpbmhnvy",
		"wpnrbwfcdmwctusunasdbpbmhnvybqqutquvbtgouklsayfvze"))
}
