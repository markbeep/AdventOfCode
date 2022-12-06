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
	ch := make(chan int, 2)
	go ff(4, cont, ch)
	go ff(14, cont, ch)
	fmt.Printf("P1: %d\nP2: %d\n", <-ch, <-ch)
	fmt.Println("Took:", time.Since(start))
}

func ff(p int, cont string, ch chan int) {
	t := make([]int, 26+int('a'))
	c := 0
	for i, v := range cont {
		if t[int(v)] > c {
			c = t[int(v)]
		}
		t[int(v)] = i
		if i-c == p {
			ch <- i + 1
			return
		}
	}
}
