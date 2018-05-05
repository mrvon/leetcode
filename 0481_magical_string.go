// Draw it on the paper and summarize the rule. it's not very hard.
package main

import "fmt"

const (
	B2N = 0
	N2B = 1
)

func magicalString(n int) int {
	b := generateMagic(n)
	one := 0
	for i := 0; i < len(b); i++ {
		if b[i] == '1' {
			one++
		}
	}
	return one
}

func generateMagic(l int) []byte {
	base := []byte{'1', '2', '2'}
	next := []byte{'1', '2'}
	status := B2N
	b := 2
	n := 2

	for {
		if len(base) >= l {
			break
		}

		if status == N2B { // next to base
			for ; n < len(next); n++ {
				if next[n] == '1' {
					if base[len(base)-1] == '1' {
						base = append(base, '2')
					} else {
						base = append(base, '1')
					}
				} else if next[n] == '2' {
					if base[len(base)-1] == '1' {
						base = append(base, '2')
						base = append(base, '2')
					} else {
						base = append(base, '1')
						base = append(base, '1')
					}
				}
			}
			status = B2N
		} else { // base to next
			for ; b < len(base); b++ {
				next = append(next, base[b])
			}
			status = N2B
		}
	}

	// maybe len(base) > l, so truncated it.
	return base[0:l]
}

func main() {
	fmt.Println(string(generateMagic(6)))
	fmt.Println(magicalString(6))
}
