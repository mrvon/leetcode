package main

import "fmt"

func complexNumberMultiply(a string, b string) string {
	x := 0
	y := 0
	z := 0
	w := 0

	fmt.Sscanf(a, "%d+%di", &x, &y)
	fmt.Sscanf(b, "%d+%di", &z, &w)
	return fmt.Sprintf("%d+%di", x*z-y*w, x*w+y*z)
}

func main() {
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
	// Output: "0+2i"
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
	// Output: "0+-2i"
}
