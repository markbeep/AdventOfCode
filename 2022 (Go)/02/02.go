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

	for _, v := range content {
		t := strings.Split(v, " ")
		switch t[1] {
		case "X":
			switch t[0] {
			case "A":
				c += 3
			case "B":
				c += 1
			case "C":
				c += 2
			}
		case "Y":
			c += 3
			switch t[0] {
			case "A":
				c += 1
			case "B":
				c += 2
			case "C":
				c += 3
			}
		case "Z":
			c += 6
			switch t[0] {
			case "A":
				c += 2
			case "B":
				c += 3
			case "C":
				c += 1
			}
		}
	}
	fmt.Println(c)
}
