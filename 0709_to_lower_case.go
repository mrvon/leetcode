package main

import "fmt"

func toLowerCase(str string) string {
	var n []rune
	for _, s := range str {
		if s >= 'A' && s <= 'Z' {
			n = append(n, s+('a'-'A'))
		} else {
			n = append(n, s)
		}
	}
	return string(n)
}

func main() {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("LOVELY"))
}
