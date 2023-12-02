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

func part1(input string) int {
	f := strings.Split(input, "\n")
	c := 0
	for i, v := range f {
		grabs := parseLine(v)
		valid := true

		for _, g := range grabs {
			oneGrabSplitted := strings.Split(g, " ")
			count := ints.SInt(oneGrabSplitted[0])
			color := oneGrabSplitted[1]
			switch color {
			case "red":
				valid = count <= 12 && valid
			case "green":
				valid = count <= 13 && valid
			case "blue":
				valid = count <= 14 && valid
			}
		}
		if valid {
			c += i + 1
		}
	}
	return c
}

func part2(input string) int {
	f := strings.Split(input, "\n")
	c := 0
	for _, v := range f {
		grabs := parseLine(v)
		red := 0
		green := 0
		blue := 0

		for _, g := range grabs {
			oneGrabSplitted := strings.Split(g, " ")
			count := ints.SInt(oneGrabSplitted[0])
			color := oneGrabSplitted[1]
			switch color {
			case "red":
				red = max(red, count)
			case "blue":
				blue = max(blue, count)
			case "green":
				green = max(green, count)
			}
		}
		c += red * green * blue
	}
	return c
}

func parseLine(input string) []string {
	splitted := strings.Split(input, ": ")
	splitted[1] = strings.ReplaceAll(splitted[1], ", ", "; ")
	return strings.Split(splitted[1], "; ")
}
