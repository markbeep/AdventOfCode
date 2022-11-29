package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
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
		a, _ := strconv.Atoi(content[i])
		b, _ := strconv.Atoi(content[i-1])
		if a > b {
			c += 1
		}
	}
	fmt.Print(c)
}
