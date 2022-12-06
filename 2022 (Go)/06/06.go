package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, _ := os.ReadFile("inp.txt")
	cont := strings.Trim(string(f), " \n")
	ch1 := make(chan int)
	ch2 := make(chan int)
	// routes := 10
	go ff(4, cont, ch1)
	go ff(14, cont, ch2)
	fmt.Printf("P1: %d\nP2: %d\n", <-ch1, <-ch2)
	fmt.Println("Took:", time.Since(start))
}

func ff(p int, cont string, ch chan int) {
	t := make([]int, 26+int('a'))
	c := 0
	last := 0
	for i, v := range cont {
		b := t[int(v)]
		if i-b <= p {
			if b > last {
				c = i - b - 1
				last = b
			} else {
				c = 0
			}
		}
		t[int(v)] = i
		c++
		if c == p {
			ch <- i + 1
			return
		}
	}
}
