/*
Using flood-fill to find every isolated region, and record it's surrounded flag.
*/
package main

import "fmt"

type List struct {
	list []int // grid in the region
	flag bool  // is this a surrounded region
}

var direction = []int{
	-1, 0,
	1, 0,
	0, -1,
	0, 1,
}

func flood_fill(board [][]byte, i int, j int, l *List) {
	m := len(board)
	n := len(board[0])

	// already visited, marked it
	board[i][j] = 'Z'
	l.list = append(l.list, i, j)

	for k := 0; k < len(direction); k += 2 {
		x := i + direction[k]
		y := j + direction[k+1]

		if x < 0 || x > m-1 || y < 0 || y > n-1 {
			// not surrounded
			l.flag = false
			continue
		}

		if board[x][y] == 'O' {
			flood_fill(board, x, y, l)
		}
	}
}

func solve(board [][]byte) {
	ll := []List{}

	// find every region
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				l := List{flag: true}

				flood_fill(board, i, j, &l)

				ll = append(ll, l)
			}
		}
	}

	// set
	for i := 0; i < len(ll); i++ {
		l := ll[i]
		for j := 0; j < len(l.list); j += 2 {
			x := l.list[j]
			y := l.list[j+1]
			if l.flag {
				board[x][y] = 'X'
			} else {
				board[x][y] = 'O'
			}
		}
	}
}

func main() {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}
	solve(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}
