package main

import (
	"fmt"
)

func checkPlayer(x rune) string {
	if x == 'X' {
		return "A"
	} else {
		return "B"
	}
}

func tictactoe(moves [][]int) string {
	var grid = [3][3]rune{}
	for i := range grid {
		grid[i] = [3]rune{' ', ' ', ' '}
	}

	var i int
	for ; i < len(moves); i++ {
		if i%2 == 0 {
			grid[moves[i][0]][moves[i][1]] = 'X'
		} else {
			grid[moves[i][0]][moves[i][1]] = 'O'
		}
	}

	for x := 0; x < 3; x++ {
		// check horizontally
		if grid[0][x] != ' ' && grid[0][x] == grid[1][x] && grid[1][x] == grid[2][x] {
			return checkPlayer(grid[0][x])
		}
		// check vertically
		if grid[x][0] != ' ' && grid[x][0] == grid[x][1] && grid[x][1] == grid[x][2] {
			return checkPlayer(grid[x][0])
		}
	}
	// check diagonally from top
	if grid[0][0] != ' ' && grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2] {
		return checkPlayer(grid[0][0])
	}
	// check diagonally from bottom
	if grid[2][0] != ' ' && grid[2][0] == grid[1][1] && grid[1][1] == grid[0][2] {
		return checkPlayer(grid[2][0])
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

func main() {
	fmt.Println(tictactoe([][]int{{0, 2}, {1, 0}, {2, 2}, {1, 2}, {2, 0}, {0, 0}, {0, 1}, {2, 1}, {1, 1}}))
}
