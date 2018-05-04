/*
After failed with pure math solution and time out with DFS solution, I finally
realized that this is a DP problem...

The idea is, if we know the max number of unique substrings in p ends with 'a',
'b', ..., 'z', then the summary of them is the answer. Why is that?

1. The max number of unique substring ends with a letter equals to the length of
max contiguous substring ends with that letter. Example "abcd", the max number
of unique substring ends with 'd' is 4, apparently they are "abcd", "bcd", "cd"
and "d".

2. If there are overlapping, we only need to consider the longest one because it
covers all the possible substrings. Example: "abcdbcd", the max number of unique
substring ends with 'd' is 4 and all substrings formed by the 2nd "bcd" part are
covered in the 4 substrings already.

3. No matter how long is a contiguous substring in p, it is in s since s has
infinite length.

4. Now we know the max number of unique substrings in p ends with 'a', 'b', ...,
'z' and those substrings are all in s. Summary is the answer, according to the
question.
*/
package main

import (
	"fmt"
)

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func findSubstringInWraproundString(p string) int {
	counter := make([]int, 26)

	length := 0
	for i := 0; i < len(p); {
		if length == 0 {
			length++
			counter[p[i]-'a'] = max(counter[p[i]-'a'], length)
			i++
		} else {
			if p[i] == 'a' && p[i-1] == 'z' || p[i-1]+1 == p[i] {
				length++
				counter[p[i]-'a'] = max(counter[p[i]-'a'], length)
				i++
			} else {
				length = 0
			}
		}
	}

	sum := 0
	for i := 0; i < len(counter); i++ {
		sum += counter[i]
	}
	return sum
}

func main() {
	fmt.Println(findSubstringInWraproundString("a"))
	fmt.Println(findSubstringInWraproundString("ab"))
	fmt.Println(findSubstringInWraproundString("abc"))
	fmt.Println(findSubstringInWraproundString("zabc"))
	fmt.Println(findSubstringInWraproundString("kabc"))
	fmt.Println(findSubstringInWraproundString("kabck"))
}
