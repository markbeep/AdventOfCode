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
	f := strings.Split(input, "\n")
	return count(f)
}

func part2(input string) int {
	repl := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for key, v := range repl {
		input = strings.ReplaceAll(input, key, fmt.Sprintf("%s%s%s", key, v, key))
	}
	f := strings.Split(input, "\n")
	return count(f)
}

func count(f []string) int {
	c := 0
	for _, v := range f {
		s := []rune{}
		for _, b := range v {
			if '0' <= b && b <= '9' {
				s = append(s, b)
			}
		}
		c += ints.SInt(fmt.Sprintf("%c%c", s[0], s[len(s)-1]))
	}
	return c
}
