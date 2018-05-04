package main

import "fmt"

func checkRecord(s string) bool {
	late := 0
	absent := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			late = 0
			absent++
			if absent > 1 {
				return false
			}
		} else if s[i] == 'L' {
			late++
			if late > 2 {
				return false
			}
		} else {
			late = 0
		}
	}

	return true
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
}
