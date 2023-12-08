package main

import (
	"aoc/util"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed inp.txt
var input string

func init() {
	input = strings.TrimSpace(input)
}

var part = flag.Int("part", 1, "the part to execute the code for")

func main() {
	flag.Parse()
	if *part == 1 {
		ans := part1(input)
		util.CopyToClipboard(ans)
		fmt.Printf("P1: %d\n", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(ans)
		fmt.Printf("P2: %d\n", ans)
	}
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}
