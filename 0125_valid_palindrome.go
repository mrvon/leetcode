package main

import "fmt"

func isPalindrome(s string) bool {
	t := []byte{}

	for i := 0; i < len(s); i++ {
		b := s[i]
		if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') {
			t = append(t, b)
		} else if b >= 'A' && b <= 'Z' {
			t = append(t, b-'A'+'a')
		}
	}

	for i := 0; i < len(t)/2; i++ {
		if t[i] != t[len(t)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))

}
