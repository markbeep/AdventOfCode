package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	c := 0
	trees := make([][]int, len(cont))
	for i := range trees {
		trees[i] = make([]int, len(cont[0]))
		for j := range trees[i] {
			val, _ := strconv.Atoi(string(cont[i][j]))
			trees[i][j] = int(val)
		}
	}

	for i := range trees {
		for j := range trees[i] {
			if check(trees, i, j) {
				c += 1
			}
		}
	}
	fmt.Println(c)
}

func check(vis [][]int, x, y int) bool {
	if x == 0 || x == len(vis)-1 || y == 0 || y == len(vis[0])-1 {
		return true
	}

	see := false
	// top down
	for i := 0; i < y; i++ {
		if vis[i][x] >= vis[y][x] {
			see = false || see
			break
		}
		if i == y-1 {
			see = true
		}
	}
	// bottom
	for i := len(vis) - 1; i > y; i-- {
		if vis[i][x] >= vis[y][x] {
			see = false || see
			break
		}
		if i == y+1 {
			see = true
		}
	}
	// left
	for i := 0; i < x; i++ {
		if vis[y][i] >= vis[y][x] {
			see = false || see
			break
		}
		if i == x-1 {
			see = true
		}
	}
	for i := len(vis[0]) - 1; i > x; i-- {
		if vis[y][i] >= vis[y][x] {
			see = false || see
			break
		}
		if i == x+1 {
			see = true
		}
	}
	return see
}
