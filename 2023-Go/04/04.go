package main

import (
	"aoc/util"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
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

func getWinners(input string) [][]string {
	f := strings.Split(input, "\n")
	matches := [][]string{}
	r := regexp.MustCompile(`\d+`)
	for _, line := range f {
		halves := strings.Split(line, "|")
		winning := map[string]bool{}
		for _, n := range r.FindAllString(halves[0], -1)[1:] {
			winning[n] = true
		}
		lineMatches := []string{}
		for _, n := range r.FindAllString(halves[1], -1) {
			if winning[n] {
				lineMatches = append(lineMatches, n)
			}
		}
		matches = append(matches, lineMatches)
	}
	return matches
}

func part1(input string) int {
	c := 0
	for _, line := range getWinners(input) {
		c += int(math.Pow(2, float64(len(line)-1)))
	}
	return c
}

func part2(input string) int {
	c := 0
	winners := getWinners(input)
	copies := map[int]int{}
	for gameId := range winners {
		copies[gameId] = 1
	}
	for gameId, line := range winners {
		for i := 0; i < len(line); i++ {
			copies[i+1+gameId] += copies[gameId]
		}
		c += copies[gameId]
	}
	return c
}
