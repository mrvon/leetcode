package main

import (
	"bufio"
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {
	scanner := bufio.NewScanner(strings.NewReader(S))
	scanner.Split(bufio.ScanWords)
	buffer := []byte{}
	c := 0
	for scanner.Scan() {
		s := scanner.Text()
		c++
		if len(buffer) > 0 {
			// append space
			buffer = append(buffer, ' ')
		}
		if !(s[0] == 'a' || s[0] == 'A' ||
			s[0] == 'e' || s[0] == 'E' ||
			s[0] == 'i' || s[0] == 'I' ||
			s[0] == 'o' || s[0] == 'O' ||
			s[0] == 'u' || s[0] == 'U') {
			// consonant
			buffer = append(buffer, s[1:]...)
			buffer = append(buffer, s[0])
			buffer = append(buffer, 'm')
		} else {
			// vowel
			buffer = append(buffer, []byte(s)...)
			buffer = append(buffer, 'm')
		}
		for i := 0; i <= c; i++ {
			buffer = append(buffer, 'a')
		}
	}
	return string(buffer)
}

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(toGoatLatin("Each word consists of lowercase and uppercase letters only"))
}
