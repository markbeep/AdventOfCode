package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := strings.Split(strings.Trim(string(f), " \n"), "\n")
	a := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ch := make(chan int, len(content)/3)
	for i := 0; i < len(content); i += 3 {
		go func(l1, l2, l3 string, ch chan int) {
			for _, k1 := range l1 {
				for _, k2 := range l2 {
					for _, k3 := range l3 {
						if k1 == k2 && k2 == k3 {
							ch <- strings.IndexRune(a, k1) + 1
							return
						}
					}
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
