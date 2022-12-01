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
	content := strings.Split(strings.Trim(string(f), " \n"), "\n")

	_ = content
	_ = c
}
