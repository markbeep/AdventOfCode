package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	var c, m int
	for i := range cont {
		for j := range cont[i] {
			l, r, t, b, see := check(cont, i, j)
			if see {
				c++
			}
			if l*r*t*b > m {
				m = l * r * t * b
			}
		}
	}
	fmt.Println(c, m)
	fmt.Println("Took", time.Since(start))
}

// returns left, right, top, bottom, isVisible
func check(vis []string, x, y int) (int, int, int, int, bool) {
	var l, r, b, t int
	see := y*x*(y-len(vis)+1)*(x-len(vis[y])+1) == 0 // edges
	// top
	for i := y - 1; i >= 0; i-- {
		if vis[i][x] >= vis[y][x] || i == 0 {
			t = y - i
			see = see || vis[i][x] < vis[y][x]
			break
		}
	}
	// bottom
	for i := y + 1; i < len(vis); i++ {
		if vis[i][x] >= vis[y][x] || i == len(vis)-1 {
			b = i - y
			see = see || vis[i][x] < vis[y][x]
			break
		}
	}
	// left
	for i := x - 1; i >= 0; i-- {
		if vis[y][i] >= vis[y][x] || i == 0 {
			l = x - i
			see = see || vis[y][i] < vis[y][x]
			break
		}
	}
	// right
	for i := x + 1; i < len(vis[0]); i++ {
		if vis[y][i] >= vis[y][x] || i == len(vis[0])-1 {
			r = i - x
			see = see || vis[y][i] < vis[y][x]
			break
		}
	}
	return l, r, t, b, see
}
