package main

import "fmt"

// Index returns the first index substr found in the s.
// function should return same result as `strings.Index` function
func Index(s string, substr string) int {
	d := CalculateSlideTable(substr)
	return IndexWithTable(&d, s, substr)
}

// IndexWithTable returns the first index substr found in the s.
// It needs the slide information of substr
func IndexWithTable(d *[256]int, s string, substr string) int {
	lsub := len(substr)
	ls := len(s)
	switch {
	case lsub == 0:
		return 0
	case lsub > ls:
		return -1
	case lsub == ls:
		if s == substr {
			return 0
		}
		return -1
	}
	i := 0
	for i+lsub-1 < ls {
		j := lsub - 1
		for ; j >= 0 && s[i+j] == substr[j]; j-- {
		}
		if j < 0 {
			return i
		}
		slide := j - d[s[i+j]]
		if slide < 1 {
			slide = 1
		}
		i += slide
	}
	return -1
}

// CalculateSlideTable builds sliding amount per each unique byte in the substring
func CalculateSlideTable(substr string) [256]int {
	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -1
	}
	for i := 0; i < len(substr); i++ {
		d[substr[i]] = i
	}
	return d
}

func InspectSlideTable(d *[256]int) {
	for i := int('A'); i <= int('Z'); i++ {
		if d[i] != -1 {
			fmt.Printf("%c:%d ", rune(i), d[i])
		}
	}
	fmt.Println()
	for i := int('a'); i <= int('z'); i++ {
		if d[i] != -1 {
			fmt.Printf("%c:%d ", rune(i), d[i])
		}
	}
	fmt.Println()
}

func main() {
	s := `Some long string which you need to check. Bla bla ... `
	substr := `long`
	d := CalculateSlideTable(substr)
	InspectSlideTable(&d)
	if pos := IndexWithTable(&d, s, substr); pos > -1 {
		fmt.Println("Found in position: ", pos, s[pos:pos+len(substr)])
	} else {
		fmt.Println("Not found")
	}
}
