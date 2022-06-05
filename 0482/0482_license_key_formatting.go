package main

import (
	"fmt"
	"unicode"
)

func licenseKeyFormatting(S string, K int) string {
	var format_str []byte
	k := 0
	for i := len(S) - 1; i >= 0; i-- {
		c := S[i]
		if c == '-' {
			continue
		}
		c = byte(unicode.ToUpper(rune(c)))
		format_str = append(format_str, c)
		k++
		if k >= K {
			format_str = append(format_str, '-')
			k = 0
		}
	}

	// trim dash in the end
	if len(format_str) > 0 {
		if format_str[len(format_str)-1] == '-' {
			format_str = format_str[:len(format_str)-1]
		}
	}

	// reverse
	i := 0
	j := len(format_str) - 1
	for i < j {
		format_str[i], format_str[j] = format_str[j], format_str[i]
		i++
		j--
	}

	return string(format_str)
}

func assert(result string, expect string) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %s, Get %s", expect, result))
	}
}

func main() {
	assert(licenseKeyFormatting("2-4A0r7-4k", 4), "24A0-R74K")
	assert(licenseKeyFormatting("2-4A0r7-4k", 3), "24-A0R-74K")
	assert(licenseKeyFormatting("--a-a-a-a--", 2), "AA-AA")
}
