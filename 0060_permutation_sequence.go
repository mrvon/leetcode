/*
Size of all permutations is Factor(N), We split whole permutation into N parts,
and size of every part is Factor(N-1). We can determined by (k - 1) / size to
select the part we want.

OK, now we consider this part is a n-1 sub-problem, solve the problem
recursively.
*/
package main

import "fmt"

var factorial []int = []int{
	1,
	1,
	2,
	6,
	24,
	120,
	720,
	5040,
	40320,
	362880,
}

func getPermutation(n int, k int) string {
	perm := []byte{}
	dict := []byte("123456789")

	for i := n; i >= 1; i-- {
		size := factorial[i-1]
		index := int((k - 1) / size)
		k = k - index*size
		perm = append(perm, dict[index])

		// remove dict[index]
		copy(dict[index:len(dict)-1], dict[index+1:len(dict)])
		dict = dict[0 : len(dict)-1]
	}

	return string(perm)
}

func main() {
	fmt.Println(getPermutation(3, 1))
	fmt.Println(getPermutation(3, 2))
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(3, 4))
	fmt.Println(getPermutation(3, 5))
	fmt.Println(getPermutation(3, 6))
}
