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

	m := 0
	for i := range trees {
		for j := range trees[i] {
			if check(trees, i, j) {
				c += 1
			}
			val := check2(trees, i, j)
			if val > m {
				m = val
			}
		}
	}
	fmt.Println(c, m)
}

func check2(vis [][]int, x, y int) int {
	top := 0
	bot := 0
	l := 0
	r := 0

	// top down
	for i := y - 1; i >= 0; i-- {
		if vis[i][x] >= vis[y][x] || i == 0 {
			top = y - i
			break
		}
	}
	// bottom
	for i := y + 1; i < len(vis); i++ {
		if vis[i][x] >= vis[y][x] || i == len(vis)-1 {
			bot = i - y
			break
		}
	}
	// left
	for i := x - 1; i >= 0; i-- {
		if vis[y][i] >= vis[y][x] || i == 0 {
			l = x - i
			break
		}
	}
	// right
	for i := x + 1; i < len(vis[0]); i++ {
		if vis[y][i] >= vis[y][x] || i == len(vis[0])-1 {
			r = i - x
			break
		}
	}
	return l * r * top * bot
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
