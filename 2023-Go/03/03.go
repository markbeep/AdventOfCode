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

func getNumStart(y, x int, f []string) int {
	for {
		x--
		if x < 0 {
			x = 0
			break
		}
		if !isInt(rune(f[y][x])) {
			x += 1
			break
		}
	}
	return x
}

func getNum(y, x int, f []string) int {
	current := ""
	// find start
	for {
		x--
		if x < 0 {
			x = 0
			break
		}
		if !isInt(rune(f[y][x])) {
			x += 1
			break
		}
	}
	// get whole int
	for {
		if x >= len(f[0]) {
			break
		}
		if !isInt(rune(f[y][x])) {
			break
		}
		current += string(f[y][x])
		x++
	}
	return ints.SInt(current)
}

func part2(input string) int {
	f := strings.Split(input, "\n")
	var added = map[string]bool{}
	c := 0

	for y, line := range f {
		for x, char := range line {
			if char == '*' {
				numbers := []int{}
				// go around star and check for numbers
				for yi := y - 1; yi <= y+1; yi++ {
					for xi := x - 1; xi <= x+1; xi++ {
						if yi < 0 || yi >= len(f) || xi < 0 || xi >= len(f[0]) {
							continue
						}
						if isInt(rune(f[yi][xi])) {
							num := getNum(yi, xi, f)
							start := getNumStart(yi, xi, f)
							if !added[fmt.Sprintf("%d,%d", yi, start)] {
								numbers = append(numbers, num)
								added[fmt.Sprintf("%d,%d", yi, start)] = true
							}
						}
					}
				}
				// multiply
				if len(numbers) == 2 {
					c += numbers[0] * numbers[1]
				}
			}
		}
	}

	return c
}
