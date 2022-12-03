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
	var c int = 0
	content := strings.Split(strings.Trim(string(f), " \n"), "\n")
	ch := make(chan int, len(content)/3)
	for i := 0; i < len(content); i += 3 {
		go calc(content[i], content[i+1], content[i+2], ch)
	}
	for i := 0; i < len(content)/3; i += 1 {
		c += <-ch
	}
	fmt.Println(c)
}

func calc(l1, l2, l3 string, ch chan int) {
	mar := false
	c := 0
	for _, k1 := range l1 {
		for _, k2 := range l2 {
			for _, k3 := range l3 {
				if !mar && k1 == k2 && k2 == k3 {
					mar = true
					p := int(k1)
					if p >= 97 {
						c += p - 96
					} else {
						c += p - 38
					}
				}
			}
		}
	}
	ch <- c
}
