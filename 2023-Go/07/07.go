package main

import (
	"aoc/util"
	"aoc/util/ints"
	_ "embed"
	"flag"
	"fmt"
	"slices"
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

const (
	FIVE      = 6
	FOUR      = 5
	FULLHOUSE = 4
	THREE     = 3
	TWOPAIR   = 2
	ONEPAIR   = 1
	HIGH      = 0
)

func ofAKind(hand string, joker bool) int {
	m := map[rune]int{}
	jokers := 0
	for _, ch := range hand {
		if joker && ch == 'J' {
			jokers++
		} else {
			m[ch]++
		}
	}
	mx := 0
	kx := ' '
	for k, v := range m {
		if v > mx {
			kx = k
			mx = v
		}
	}
	m[kx] += jokers
	three := 0
	two := 0
	for _, v := range m {
		if v == 5 {
			return FIVE
		}
		if v == 4 {
			return FOUR
		}
		if v == 3 {
			three++
		}
		if v == 2 {
			two++
		}
	}
	if three > 0 && two > 0 {
		return FULLHOUSE
	}
	if three > 0 {
		return THREE
	}
	if two == 2 {
		return TWOPAIR
	}
	if two == 1 {
		return ONEPAIR
	}
	return HIGH
}

func replace(a byte, joker bool) int {
	switch a {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if joker {
			return -10
		}
		return 11
	case 'T':
		return 10
	default:
		return int(a) - int('0')
	}
}

type camel struct {
	hand     string
	strength int
	bid      int
}

func part1(input string) int {
	f := strings.Split(input, "\n")
	camels := []camel{}
	for _, line := range f {
		splitted := strings.Split(line, " ")
		camels = append(camels, camel{hand: splitted[0], bid: ints.SInt(splitted[1]), strength: ofAKind(splitted[0], false)})
	}
	slices.SortFunc(camels, func(a, b camel) int {
		if a.strength == b.strength {
			for i := range a.hand {
				cmp := replace(a.hand[i], false) - replace(b.hand[i], false)
				if cmp != 0 {
					return cmp
				}
			}
			return 0
		}
		return a.strength - b.strength
	})
	c := 0
	for i, v := range camels {
		c += (i + 1) * v.bid
	}

	return c
}

func part2(input string) int {
	f := strings.Split(input, "\n")
	camels := []camel{}
	for _, line := range f {
		splitted := strings.Split(line, " ")
		camels = append(camels, camel{hand: splitted[0], bid: ints.SInt(splitted[1]), strength: ofAKind(splitted[0], true)})
	}
	slices.SortFunc(camels, func(a, b camel) int {
		if a.strength == b.strength {
			for i := range a.hand {
				cmp := replace(a.hand[i], true) - replace(b.hand[i], true)
				if cmp != 0 {
					return cmp
				}
			}
			return 0
		}
		return a.strength - b.strength
	})
	c := 0
	for i, v := range camels {
		c += (i + 1) * v.bid
	}

	return c
}
