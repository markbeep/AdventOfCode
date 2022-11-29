package main

import (
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
	content := strings.Split(string(f), "\n")

	_ = content
	_ = c
}
