/*
Just Simulate Division
*/
package main

import (
	"fmt"
	"strconv"
)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func sign(x int) int {
	if x >= 0 {
		return 1
	} else {
		return -1
	}
}

func fractionToDecimal(numerator int, denominator int) string {
	n := abs(numerator)
	d := abs(denominator)

	s := []byte{}
	if numerator != 0 && sign(numerator) != sign(denominator) {
		s = append(s, '-')
	}

	// Integral part
	if n > d {
		r := n % d
		q := strconv.Itoa((n - r) / d)
		n = r
		s = append(s, q...)
	} else {
		s = append(s, '0')
	}

	// Next round
	n *= 10

	if n == 0 {
		return string(s)
	}

	s = append(s, '.')

	m := make(map[int]int) // n -> index

	// Fraction part
	for n != 0 {
		if n < d {
			if index, seen := m[n]; seen {
				// found recurring decimal
				wrap := []byte{}
				wrap = append(wrap, s[:index]...)
				wrap = append(wrap, '(')
				wrap = append(wrap, s[index:]...)
				wrap = append(wrap, ')')
				return string(wrap)
			} else {
				m[n] = len(s)
			}
			s = append(s, '0')
		} else {
			r := n % d
			q := strconv.Itoa((n - r) / d)
			if index, seen := m[n]; seen {
				// found recurring decimal
				wrap := []byte{}
				wrap = append(wrap, s[:index]...)
				wrap = append(wrap, '(')
				wrap = append(wrap, s[index:]...)
				wrap = append(wrap, ')')
				return string(wrap)
			} else {
				m[n] = len(s)
			}
			n = r
			s = append(s, q...)
		}

		// Next round
		n *= 10
	}

	return string(s)
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(3, 2))
	fmt.Println(fractionToDecimal(30, 2))
	fmt.Println(fractionToDecimal(3, 20))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(1, 90))
	fmt.Println(fractionToDecimal(1, 900))
	fmt.Println(fractionToDecimal(1, 99))
	fmt.Println(fractionToDecimal(50, -8))
	fmt.Println(fractionToDecimal(-50, -8))
	fmt.Println(fractionToDecimal(-50, 8))
	fmt.Println(fractionToDecimal(50, 8))
	fmt.Println(fractionToDecimal(0, -5))
}
