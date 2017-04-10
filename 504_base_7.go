package main

import "fmt"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	var b []byte
	is_negative := false

	if num < 0 {
		num = -num
		is_negative = true
	}

	for num > 0 {
		b = append(b, byte('0'+num%7))
		num = num / 7
	}

	if is_negative {
		b = append(b, '-')
	}

	i := 0
	j := len(b) - 1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}

func main() {
	fmt.Println(convertToBase7(0))
	fmt.Println(convertToBase7(2))
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
}
