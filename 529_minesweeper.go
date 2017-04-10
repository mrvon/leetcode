package main

import "fmt"

var adjacent []int = []int{
	-1, -1,
	-1, 0,
	-1, 1,
	0, -1,
	0, 1,
	1, -1,
	1, 0,
	1, 1,
}

func adjacentMines(board [][]byte, x int, y int) int {
	count := 0
	for i := 0; i < len(adjacent); i += 2 {
		cx := x + adjacent[i]
		cy := y + adjacent[i+1]
		if cx < 0 || cx >= len(board) || cy < 0 || cy >= len(board[0]) {
			continue
		}
		if board[cx][cy] == 'M' {
			count++
		}
	}
	return count
}

func adjacentProcess(board [][]byte, x int, y int) {
	for i := 0; i < len(adjacent); i += 2 {
		cx := x + adjacent[i]
		cy := y + adjacent[i+1]
		if cx < 0 || cx >= len(board) || cy < 0 || cy >= len(board[0]) {
			continue
		}
		if board[cx][cy] == 'E' {
			updateBoard(board, []int{
				cx, cy,
			})
		}
	}
}

func updateBoard(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]

	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else if board[x][y] == 'E' {
		count := adjacentMines(board, x, y)
		if count > 0 {
			board[x][y] = byte('0' + count)
		} else {
			board[x][y] = 'B'
			adjacentProcess(board, x, y)
		}
	} else {
		panic("error")
	}

	return board
}

func printBoard(board [][]byte) {
	for i := 0; i < len(board); i++ {
		fmt.Println(string(board[i]))
	}
	fmt.Println("")
}

func main() {
	board := [][]byte{
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'M', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
	}
	printBoard(board)
	click := []int{3, 0}
	printBoard(updateBoard(board, click))

	board = [][]byte{
		[]byte{'B', '1', 'E', '1', 'B'},
		[]byte{'B', '1', 'M', '1', 'B'},
		[]byte{'B', '1', '1', '1', 'B'},
		[]byte{'B', 'B', 'B', 'B', 'B'},
	}
	printBoard(board)
	click = []int{1, 2}
	printBoard(updateBoard(board, click))
}
