package main

import (
	"aoc/util"
	"fmt"
)

type Piece struct {
	height int8
	grid   [][]int
}

func main() {
	f := util.Read("inp.txt")

	// y=0 is the bottom, x=0 left, x=6 right
	// rocks start x=2
	// empty=0, block=1
	board := util.Array2[int8](7, 7)

	pieces := []Piece{
		{1, [][]int{{1, 1, 1, 1}}},
		{3, [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}},
		{3, [][]int{{0, 0, 1}, {0, 0, 1}, {1, 1, 1}}},
		{4, [][]int{{1}, {1}, {1}, {1}}},
		{2, [][]int{{1, 1}, {1, 1}}},
	}

	y, x, pi, startHeight, cur := 3, 2, 0, 3, 0
	c := 0
	for {
		moved := false
		dir := f[cur%len(f)]
		// pBoard(board, x, y)
		if dir == '>' && canMove(board, pieces[pi], '>', x, y) {
			x++
		}
		if dir == '<' && canMove(board, pieces[pi], '<', x, y) {
			x--
		}
		if canMove(board, pieces[pi], 'v', x, y) {
			y--
			moved = true
		}
		// fmt.Println("DIR:", string(dir), cur)
		cur++
		if !moved {
			// place piece
			for i := 0; i < len(pieces[pi].grid); i++ {
				for j := 0; j < len(pieces[pi].grid[i]); j++ {
					if pieces[pi].grid[len(pieces[pi].grid)-i-1][j] == 1 {
						board[y+i][x+j] = 1
					}
				}
			}
			startHeight = y + int(pieces[pi].height) + 3
			pi = (pi + 1) % len(pieces)
			// make sure field is big enough
			y, x = startHeight, 2 // reset starting values
			if len(board) < y+int(pieces[pi].height) {
				for i := 0; i < y+int(pieces[pi].height)-len(board); i++ {
					board = append(board, make([]int8, 7))
				}
			}
			c++
			// pBoard(board, 0, 0)
			if c == 2022 {
				break
			}
		}
	}
	// find last line
	fmt.Println(measure(board))
}

func measure(board [][]int8) int {
	for i := len(board) - 1; i >= 0; i-- {
		for j := 0; j < 7; j++ {
			if board[i][j] == 1 {
				return i
			}
		}
	}
	return 0
}

func canMove(board [][]int8, piece Piece, dir rune, x, y int) bool {
	switch dir {
	case 'v':
		if y == 0 {
			return false
		}
		for i := 0; i < len(piece.grid); i++ {
			for j := 0; j < len(piece.grid[i]); j++ {
				if y-1+i < len(board) && piece.grid[len(piece.grid)-1-i][j] == 1 && board[y-1+i][x+j] == 1 {
					return false
				}
			}
		}
	case '<':
		if x == 0 {
			return false
		}
		for i := 0; i < len(piece.grid); i++ {
			for j := 0; j < len(piece.grid[i]); j++ {
				if y+i < len(board) && piece.grid[len(piece.grid)-1-i][j] == 1 && board[y+i][x+j-1] == 1 {
					return false
				}
			}
		}
	case '>':
		if x+len(piece.grid[0]) >= 7 {
			return false
		}
		for i := 0; i < len(piece.grid); i++ {
			for j := 0; j < len(piece.grid[i]); j++ {
				if y+i < len(board) && piece.grid[len(piece.grid)-1-i][j] == 1 && board[y+i][x+j+1] == 1 {
					return false
				}
			}
		}
	}
	return true
}

func pBoard(board [][]int8, x, y int) {
	fmt.Println("----------------")
	fmt.Println("    0123456    ")
	for i := len(board) - 1; i >= 0; i-- {
		if y == i {
			fmt.Printf("--->")
		} else {
			fmt.Printf("%4d", i)
		}
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 1 {
				fmt.Print("#")
			} else if j == x && i == y {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("%-4d\n", i)
	}
	fmt.Println("    0123456    ")
	fmt.Println("----------------")
}
