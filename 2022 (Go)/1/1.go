package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	helper "aoc/Helper"
)

func main() {
	log.Print(helper.Hint("2"))
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	var c int = 0
	content := strings.Split(string(f), "\n")

	for i := range content {
		if i == 0 {
			continue
		}
		a, _ := strconv.ParseInt(content[i], 10, 64)
		b, _ := strconv.ParseInt(content[i-1], 10, 64)
		if a > b {
			c += 1
		}
	}
	fmt.Print(c)
}
