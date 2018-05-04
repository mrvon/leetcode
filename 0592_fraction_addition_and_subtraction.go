package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Fraction struct {
	n int // numerator
	d int // denominator
}

func parse(expression string) []Fraction {
	reader := strings.NewReader(expression)
	list := []Fraction{}
	c, _ := reader.ReadByte()
	for reader.Len() > 0 {
		ns := []byte{}
		if c == '+' || c == '-' {
			ns = append(ns, c)
			c, _ = reader.ReadByte()
		}
		for c >= '0' && c <= '9' {
			ns = append(ns, c)
			c, _ = reader.ReadByte()
		}
		// skip '/'
		ds := []byte{}
		c, _ = reader.ReadByte()
		if c == '+' || c == '-' {
			ds = append(ds, c)
			c, _ = reader.ReadByte()
		}
		for c >= '0' && c <= '9' {
			ds = append(ds, c)
			if reader.Len() == 0 {
				break
			}
			c, _ = reader.ReadByte()
		}
		n, _ := strconv.Atoi(string(ns))
		d, _ := strconv.Atoi(string(ds))
		list = append(list, Fraction{n, d})
	}

	return list
}

// Greatest Common Divisor
func gcd(x int, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

// Lowest Common Multiple
func lcm(x int, y int) int {
	return x * y / gcd(x, y)
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func addition(x Fraction, y Fraction) Fraction {
	d := lcm(x.d, y.d)
	n := (x.n * d / x.d) + (y.n * d / y.d)
	return Fraction{n, d}
}

func reduce(x Fraction) Fraction {
	g := gcd(abs(x.n), abs(x.d))
	return Fraction{x.n / g, x.d / g}
}

func fractionAddition(expression string) string {
	list := parse(expression)
	result := Fraction{0, 1}
	for i := 0; i < len(list); i++ {
		result = addition(result, list[i])
	}
	result = reduce(result)
	return fmt.Sprintf("%d/%d", result.n, result.d)
}

func main() {
	fmt.Println(fractionAddition("-1/2+1/2"))
	fmt.Println(fractionAddition("-1/2+1/2+1/3"))
	fmt.Println(fractionAddition("1/3-1/2"))
	fmt.Println(fractionAddition("5/3+1/3"))
}
