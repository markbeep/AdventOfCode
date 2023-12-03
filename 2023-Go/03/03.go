package main

import (
	"aoc/util"
	"aoc/util/ints"
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

var part = flag.Int("part", 2, "the part to execute the code for")

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

func isInt(input rune) bool {
	return input >= '0' && input <= '9'
}

func touches(y, x int, f []string) bool {
	for yi := y - 1; yi <= y+1; yi++ {
		for xi := x - 1; xi <= x+1; xi++ {
			if yi < 0 || yi >= len(f) || xi < 0 || xi >= len(f[0]) {
				continue
			}
			if !isInt(rune(f[yi][xi])) && f[yi][xi] != '.' {
				return true
			}
		}
	}
	return false
}

func part1(input string) int {
	f := strings.Split(input, "\n")
	c := 0

	for y, line := range f {
		current := ""
		touch := false
		for x, char := range line {
			if isInt(char) {
				current += string(char)
				if touches(y, x, f) {
					touch = true
				}
			} else {
				if len(current) > 0 && touch {
					c += ints.SInt(current)
				}
				touch = false
				current = ""
			}
		}
		if len(current) > 0 && touch {
			c += ints.SInt(current)
		}
	}

	return c
}

type pos struct {
	x, y int
}

func part2(input string) int {
	f := strings.Split(input, "\n")
	gear := map[pos]int{}
	c := 0

	combineCurrent := func(current string, star pos) {
		prev, ok := gear[star]
		if ok {
			c += prev * ints.SInt(current)
		} else {
			gear[star] = ints.SInt(current)
		}
	}

	for y, line := range f {
		current := ""
		var star *pos
		for x, char := range line {
			if isInt(char) {
				current += string(char)
				for yi := y - 1; yi <= y+1; yi++ {
					for xi := x - 1; xi <= x+1; xi++ {
						if yi < 0 || yi >= len(f) || xi < 0 || xi >= len(f[0]) {
							continue
						}
						if f[yi][xi] == '*' {
							star = &pos{x: xi, y: yi}
						}
					}
				}
			} else {
				if len(current) > 0 && star != nil {
					combineCurrent(current, *star)
				}
				star = nil
				current = ""
			}
		}
		if len(current) > 0 && star != nil {
			combineCurrent(current, *star)
		}
	}

	return c
}
