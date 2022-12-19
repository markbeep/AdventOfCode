package main

import (
	"aoc/util"
	"aoc/util/ints"
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

	y, x, cur, c := 3, 2, 0, 0
	p1 := false
	oldHeight := 0
	extra := len(f) % len(pieces)
	lm := 0
	fmt.Println("Extra:", extra)
	var p2Height int64
	path := [][]int{}
	p2 := 0
	_, _, _, _ = p1, p2Height, lm, path
	for {
		moved := false
		dir := f[cur%len(f)]
		pie := pieces[c%len(pieces)]
		if dir == '>' && canMove(board, pie, '>', x, y) {
			x++
		}
		if dir == '<' && canMove(board, pie, '<', x, y) {
			x--
		}
		if canMove(board, pie, 'v', x, y) {
			y--
			moved = true
		}
		if cur%1000000 == 0 {
			fmt.Println(cur)
		}
		cur++
		if !moved {
			if lm < 15 {
				path = append(path, []int{x, measure(board)})
				for start := 0; start < len(path); start++ {
					for i := 1; i < len(path); i++ {
						for j := 0; j < len(path)-i-start; j++ {
							if path[start+j][0] != path[start+j+i][0] || (start+j > 0 && start+j+i > 0 && (path[start+j][1]-path[start+j-1][1]) != (path[start+j+i][1]-path[start+j+i-1][1])) {
								lm = ints.Max(lm, j)
								break
							} else {
								if j >= 14 {
									fmt.Println("START:", start+j)
								}
								lm = ints.Max(lm, j+1)
							}
						}
					}
				}
				fmt.Println(c, lm)
			} else if p2 == 0 {
				p2 = c
			}

			// place piece
			for i := 0; i < len(pie.grid); i++ {
				for j := 0; j < len(pie.grid[i]); j++ {
					if pie.grid[len(pie.grid)-i-1][j] == 1 {
						board[y+i][x+j] = 1
					}
				}
			}
			y, x = measure(board)+4, 2 // reset starting values
			_ = oldHeight
			oldHeight = y
			// make sure field is big enough
			if len(board) < y+int(pie.height) {
				for i := 0; i < y+int(pie.height)-len(board); i++ {
					board = append(board, make([]int8, 7))
				}
			}
			c++
			if c == 2022 {
				fmt.Println("Part 1:", measure(board)+1)
				p1 = true
			}
			if p1 && lm >= 15 {
				fmt.Printf("Range: [%d:%d]\n", lm, path[p2-1][1]-path[p2-lm-1][1])
				for i, v := range path[20:] {
					fmt.Printf("%d-%d ", i+20, v[0])
				}
				fmt.Println("\nHeight:", path[46][1]-path[22][1])
				fmt.Println("H-before:", path[22][1])
				break
			}
		}

	}
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
