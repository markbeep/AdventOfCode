package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type tree struct {
	l   int
	r   int
	t   int
	b   int
	vis bool
}

func main() {
	start := time.Now()
	f, _ := os.ReadFile("inp.txt")
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	cch := make(chan int, len(cont))
	mch := make(chan int, len(cont))
	for i := range cont {
		go func(i int, cch, mch chan int) {
			var c, m int
			for j := range cont[i] {
				val, vis := check(cont, i, j)
				if vis {
					c++
				}
				if val > m {
					m = val
				}
			}
			cch <- c
			mch <- m
		}(i, cch, mch)
	}
	var c, m int
	for range cont {
		c += <-cch
		v := <-mch
		if v > m {
			m = v
		}
	}
	fmt.Println(c, m)
	fmt.Println("Took", time.Since(start))
}

// returns left, right, top, bottom, isVisible
func check(vis []string, x, y int) (int, bool) {
	see := y*x*(y-len(vis)+1)*(x-len(vis[y])+1) == 0 // edges
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	p := 1

	for _, d := range dirs {
		for r := range vis {
			i := y + d[0]*(r+1)
			j := x + d[1]*(r+1)
			if (i >= 0 && i <= len(vis)-1 && j >= 0 && j <= len(vis)-1) && vis[i][j] >= vis[y][x] || i == 0 || i == len(vis)-1 || j == 0 || j == len(vis)-1 {
				p *= r + 1
				see = see || vis[i][j] < vis[y][x]
				break
			}
		}
	}
	return p, see
}
