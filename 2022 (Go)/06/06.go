package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("inp.txt")
	cont := strings.Trim(string(f), " \n")
	ch := make(chan int, 2)
	go ff(4, cont, ch)
	go ff(14, cont, ch)
	fmt.Printf("P1: %d\nP2: %d\n", <-ch, <-ch)
}

func ff(p int, cont string, ch chan int) {
	for i := p; i < len(cont); i++ {
		var hash = map[byte]bool{}
		for j := 0; j < p; j++ {
			hash[cont[i-j]] = true
		}
		if len(hash) >= p {
			ch <- i + 1
			return
		}
	}
}
