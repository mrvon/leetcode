/*
Dynamic programming solution.

longest(s) = max(longest(s-1), longestWithRight(s))

longestWithRight(s) is the longest valid parentheses include right end character.
*/
package main

type Stack struct {
	S []byte
}

func (S *Stack) len() int {
	return len(S.S)
}

func (S *Stack) push(t byte) {
	S.S = append(S.S, t)
}

func (S *Stack) pop() (t byte) {
	t = S.S[len(S.S)-1]
	S.S = S.S[:len(S.S)-1]
	return
}

func (S *Stack) peek() (t byte) {
	t = S.S[len(S.S)-1]
	return
}

func longestWithRight(ss string) int {
	var s Stack
	maxlen := 0
	counter := 0

	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == ')' {
			s.push(ss[i])
			counter++
		} else { // '('
			if s.len() == 0 {
				return maxlen
			}

			b := s.pop()
			if b == ')' {
				// match
				if s.len() == 0 {
					maxlen = counter * 2
				}
			} else {
				// mismatch
				return maxlen
			}
		}
	}

	return maxlen
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func longestValidParentheses(s string) int {
	optimal := make(map[int]int)

	for length := 2; length <= len(s); length++ {
		optimal[length] = max(optimal[length-1], longestWithRight(s[0:length]))
	}

	return optimal[len(s)]
}

func main() {
	// fmt.Println(longestWithRight("()"))
	// fmt.Println(longestWithRight("()()"))
	// fmt.Println(longestWithRight("(())"))
	// fmt.Println(longestWithRight(")()"))
	// fmt.Println(longestWithRight(")"))
	// fmt.Println(longestWithRight("())"))
	// fmt.Println(longestWithRight("(()"))
	// fmt.Println(longestWithRight(")()())()"))

	// fmt.Println(longestValidParentheses("(()"))
	// fmt.Println(longestValidParentheses("())"))
	// fmt.Println(longestValidParentheses(")()())"))
	// fmt.Println(longestValidParentheses(")(((((()())()()))()(()))("))
	// fmt.Println(longestValidParentheses(")()())()()("))
}
