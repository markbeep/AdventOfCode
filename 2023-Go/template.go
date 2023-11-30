package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inp.txt
var input string

func init() {
	input = strings.TrimSpace(input)
}

func main() {
	f := strings.Split(input, "\n")
	c := 0

	for _, v := range f {
		fmt.Sscanf(v, "")

	}
	fmt.Println(c)
}
