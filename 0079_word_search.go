package main

import "fmt"

func ismark(b byte) bool {
	return b&(1<<7) != 0
}

func mark(b byte) byte {
	return b | (1 << 7)
}

func clear(b byte) byte {
	return b & ((1 << 7) - 1)
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if __exist(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func __exist(board [][]byte, i int, j int, word string, k int) bool {
	if k >= len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	if ismark(board[i][j]) {
		return false
	}

	if board[i][j] != word[k] {
		return false
	}

	board[i][j] = mark(board[i][j])

	is_exist := __exist(board, i+1, j, word, k+1) ||
		__exist(board, i-1, j, word, k+1) ||
		__exist(board, i, j+1, word, k+1) ||
		__exist(board, i, j-1, word, k+1)

	board[i][j] = clear(board[i][j])

	return is_exist
}

func assert(expect bool, result bool) {
	if result != expect {
		panic(fmt.Sprintf("Assert failed!, Expect %b, Get %b", expect, result))
	}
}

func main() {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
}
