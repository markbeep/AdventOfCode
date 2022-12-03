package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("inp.txt")
	content := strings.Split(strings.Trim(string(f), " \n"), "\n")
	a := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ch := make(chan int, len(content)/3)
	for i := 0; i < len(content); i += 3 {
		go func(l1, l2, l3 string, ch chan int) {
			for _, k1 := range l1 {
				if strings.ContainsRune(l2, k1) && strings.ContainsRune(l3, k1) {
					ch <- strings.IndexRune(a, k1) + 1
					return
				}
			}
		}(content[i], content[i+1], content[i+2], ch)
	}
	c := 0
	for i := 0; i < len(content)/3; i += 1 {
		c += <-ch
	}
	fmt.Println(c)
}
