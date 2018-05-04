// TODO
// very slow but accept anyway.
// should use trie to speed up
package main

import "fmt"

func mark(b byte) byte {
	return b | (1 << 7)
}

func clear(b byte) byte {
	return b & ((1 << 7) - 1)
}

func ismark(b byte) bool {
	return (b & (1 << 7)) != 0
}

func findWords(board [][]byte, words []string) []string {
	set := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		set[words[i]] = true
	}

	var list []string
	for word := range set {
		if findWord(board, word) {
			list = append(list, word)
		}
	}
	return list
}

func findWord(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if find(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func find(board [][]byte, i int, j int, word string, k int) bool {
	if k >= len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	if ismark(board[i][j]) || board[i][j] != word[k] {
		return false
	}

	board[i][j] = mark(board[i][j])

	is_exist := find(board, i+1, j, word, k+1) ||
		find(board, i-1, j, word, k+1) ||
		find(board, i, j+1, word, k+1) ||
		find(board, i, j-1, word, k+1)

	board[i][j] = clear(board[i][j])

	return is_exist
}

func main() {
	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}

	words := []string{"oath", "oath", "pea", "eat", "rain"}

	fmt.Println(findWords(board, words))
}
