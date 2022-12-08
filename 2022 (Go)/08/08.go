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
				l, r, t, b, vis := check(cont, i, j)
				if vis {
					c++
				}
				if l*r*t*b > m {
					m = l * r * t * b
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
