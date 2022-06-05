// big number addition
// big number multiplication
package main

import "fmt"

func add(num1 string, num2 string) string {
	toint := func(b byte) int {
		return int(b - '0')
	}

	tobyte := func(i int) byte {
		return byte(i + '0')
	}

	max := func(i int, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}

	a := []byte(num1)
	b := []byte(num2)
	base := 10

	// reverse
	i := 0
	j := len(a) - 1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	// reverse
	i = 0
	j = len(b) - 1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	m := max(len(a), len(b))
	p := make([]int, m+1)

	carry := 0
	for i := 0; i <= m; i++ {
		x := 0
		y := 0
		if i < len(a) {
			x = toint(a[i])
		}
		if i < len(b) {
			y = toint(b[i])
		}
		p[i] = carry + x + y
		carry = p[i] / base
		p[i] = p[i] % base
	}
	p[m] += carry

	var r []byte

	i = len(p) - 1
	// skip leading 0
	for i > 0 {
		if p[i] != 0 {
			break
		} else {
			i--
		}
	}

	for ; i >= 0; i-- {
		r = append(r, tobyte(p[i]))
	}

	return string(r)
}

func multiply(num1 string, num2 string) string {
	toint := func(b byte) int {
		return int(b - '0')
	}

	tobyte := func(i int) byte {
		return byte(i + '0')
	}

	a := []byte(num1)
	b := []byte(num2)
	base := 10

	// reverse
	i := 0
	j := len(a) - 1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	// reverse
	i = 0
	j = len(b) - 1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	p := make([]int, len(a)+len(b))

	for bi := 0; bi < len(b); bi++ {
		carry := 0
		for ai := 0; ai < len(a); ai++ {
			p[bi+ai] += carry + toint(b[bi])*toint(a[ai])
			carry = p[bi+ai] / base
			p[bi+ai] = p[bi+ai] % base
		}
		p[bi+len(a)] += carry
	}

	var r []byte

	i = len(p) - 1
	// skip leading 0
	for i > 0 {
		if p[i] != 0 {
			break
		} else {
			i--
		}
	}

	for ; i >= 0; i-- {
		r = append(r, tobyte(p[i]))
	}

	return string(r)
}

func assert(expect string, result string) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %s, Get %s", expect, result))
	}
}

func main() {
	assert("0", add("0", "0"))
	assert("2", add("1", "1"))
	assert("11", add("10", "1"))
	assert("24", add("12", "12"))
	assert("198", add("99", "99"))

	assert("0", multiply("0", "0"))
	assert("1", multiply("1", "1"))
	assert("10", multiply("10", "1"))
	assert("144", multiply("12", "12"))
	assert("9801", multiply("99", "99"))
}
