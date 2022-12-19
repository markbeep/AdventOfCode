package main

import (
	"aoc/util"
	"fmt"
)

func main() {
	f := util.ReadS("inp.txt", "\n")
	grid := util.Array3[int](20, 20, 20)
	for _, v := range f {
		var x, y, z int
		fmt.Sscanf(v, "%d,%d,%d", &x, &y, &z)
		grid[z][y][x] = 1
	}
	tot := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for k := 0; k < len(grid[0][0]); k++ {
				if grid[i][j][k] > 0 {
					if i <= 0 || grid[i-1][j][k] == 0 {
						tot++
					}
					if i >= len(grid)-1 || grid[i+1][j][k] == 0 {
						tot++
					}

					if j <= 0 || grid[i][j-1][k] == 0 {
						tot++
					}
					if j >= len(grid[0])-1 || grid[i][j+1][k] == 0 {
						tot++
					}

					if k <= 0 || grid[i][j][k-1] == 0 {
						tot++
					}
					if k >= len(grid[0][0])-1 || grid[i][j][k+1] == 0 {
						tot++
					}
				} else if covered(grid, i, j, k) {
					fmt.Println(i, j, k)
					tot -= 6
				}
			}
		}
	}
	fmt.Println(tot)
}

// make into dfs
func covered(grid [][][]int, i, j, k int) bool {
	if i <= 0 || i >= len(grid)-1 || j <= 0 || j >= len(grid[0])-1 || k <= 0 || k >= len(grid[0][0])-1 {
		return false
	}
	if grid[i-1][j][k] == 1 &&
		grid[i+1][j][k] == 1 &&
		grid[i][j-1][k] == 1 &&
		grid[i][j+1][k] == 1 &&
		grid[i][j][k-1] == 1 &&
		grid[i][j][k+1] == 1 {
		return true
	}

	return false
}
