/*
Two pass, straightforward solution.

TODO - One pass solution?
*/

package main

import "fmt"

func findLonelyPixel(picture [][]byte) int {
	row := make(map[int]int)
	col := make(map[int]int)

	for i := 0; i < len(picture); i++ {
		for j := 0; j < len(picture[i]); j++ {
			if picture[i][j] == 'B' {
				row[i]++
				col[j]++
			}
		}
	}

	count := 0
	for i := 0; i < len(picture); i++ {
		for j := 0; j < len(picture[i]); j++ {
			if picture[i][j] == 'B' {
				if row[i] > 1 {
					break
				}
				if col[j] > 1 {
					break
				}
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(findLonelyPixel([][]byte{
		{'W', 'B', 'B'},
		{'W', 'B', 'W'},
		{'B', 'W', 'W'}}))
	fmt.Println(findLonelyPixel([][]byte{
		{'W', 'W', 'B'},
		{'W', 'B', 'W'},
		{'B', 'W', 'W'}}))
}
