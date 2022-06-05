package main

import "fmt"

func solve(board [][]byte, i int, j int) bool {
	if i == -1 || j == -1 {
		// wo got a solution
		return true
	}

	// try place n
	for k := 1; k <= 9; k++ {
		n := byte(k + '0')
		if !check(board, i, j, n) {
			continue
		}
		board[i][j] = n
		ni, nj := next(board, i, j)
		if solve(board, ni, nj) {
			return true
		}
	}

	// restore it
	board[i][j] = '.'
	return false
}

func check(board [][]byte, i int, j int, n byte) bool {
	// check row
	for c := 0; c < 9; c++ {
		if board[i][c] == n {
			return false
		}
	}
	// check column
	for r := 0; r < 9; r++ {
		if board[r][j] == n {
			return false
		}
	}
	// check in the box
	box_i := i / 3 * 3
	box_j := j / 3 * 3
	for p := 0; p < 3; p++ {
		for q := 0; q < 3; q++ {
			if board[p+box_i][q+box_j] == n {
				return false
			}
		}
	}
	return true
}

func next(board [][]byte, i int, j int) (int, int) {
	_next := func(i int, j int) (int, int) {
		if i == -1 || j == -1 {
			return 0, 0
		} else if j < 8 {
			return i, j + 1
		} else { // j >= 8
			if i < 8 {
				return i + 1, 0
			} else { // i >= 8
				return -1, -1
			}
		}
	}

	ni := i
	nj := j

	for {
		ni, nj = _next(ni, nj)
		if ni == -1 || nj == -1 {
			// can't find next slot
			break
		}
		if board[ni][nj] == '.' {
			// ok, we got one.
			break
		}
		// try to find again
	}

	return ni, nj
}

func solveSudoku(board [][]byte) {
	i := 0
	j := 0
	i, j = next(board, -1, -1)
	solve(board, i, j)
}

func print_board(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Printf("%c", '\n')
	}
	fmt.Printf("%c", '\n')
}

func main() {
	board := [][]byte{
		[]byte("..9748..."),
		[]byte("7........"),
		[]byte(".2.1.9..."),
		[]byte("..7...24."),
		[]byte(".64.1.59."),
		[]byte(".98...3.."),
		[]byte("...8.3.2."),
		[]byte("........6"),
		[]byte("...2759.."),
	}
	solveSudoku(board)
}
