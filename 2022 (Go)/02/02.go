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
	var c int = 0
	content := strings.Split(strings.Trim(string(f), " \n"), "\n")
	for _, v := range content {
		t := []rune(v)
		l := int(t[0] - 65)
		switch t[2] - 88 {
		case 0:
			c += (l+2)%3 + 1
		case 1:
			c += l + 4
		case 2:
			c += (l+1)%3 + 7
		}
	}
	fmt.Println(c)
	fmt.Println("Took:", time.Since(start))
}
