package main

import (
	"fmt"
	"strconv"
)

func valid(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	} else if n == 2 || n == 3 {
		if s[0] == '0' {
			return false
		}
		if i, err := strconv.Atoi(s); err == nil && i <= 255 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func restoreIpAddresses(s string) []string {
	n := len(s)
	var results []string

	for i := 1; i < n && i <= 3; i++ {
		if !valid(s[0:i]) {
			continue
		}
		for j := i + 1; j < n && j <= i+3; j++ {
			if !valid(s[i:j]) {
				continue
			}
			for k := j + 1; k < n && k <= j+3; k++ {
				if valid(s[j:k]) && valid(s[k:n]) {
					results = append(results, s[0:i]+"."+s[i:j]+"."+s[j:k]+"."+s[k:n])
				}
			}
		}
	}

	return results
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("01010101"))
}
