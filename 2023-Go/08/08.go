package main

import (
	"aoc/util"
	_ "embed"
	"flag"
	"fmt"
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

func part1(input string) int {
	f := strings.Split(input, "\n")
	directions := f[0]
	m := map[string][]string{}
	for _, line := range f[2:] {
		splitted := strings.Split(line, " = ")
		r := regexp.MustCompile(`\w+`)
		m[splitted[0]] = r.FindAllString(splitted[1], -1)
	}

	repeat := 0
	i := 0
	cur := "AAA"
	for {
		ch := directions[i]
		repeat++
		if ch == 'L' {
			cur = m[cur][0]
		} else {
			cur = m[cur][1]
		}
		// are we done?
		if i == len(directions)-1 {
			if cur == "ZZZ" {
				break
			}
			i = 0
			continue
		}

		i++
	}

	return repeat
}

func compute(s, dirs string, m map[string][]string) string {
	cur := s
	for _, ch := range dirs {
		if ch == 'L' {
			cur = m[cur][0]
		} else {
			cur = m[cur][1]
		}
	}
	return cur
}

func part2(input string) int {
	f := strings.Split(input, "\n")
	directions := f[0]
	m := map[string][]string{}
	for _, line := range f[2:] {
		splitted := strings.Split(line, " = ")
		r := regexp.MustCompile(`\w+`)
		m[splitted[0]] = r.FindAllString(splitted[1], -1)
	}
	comp := map[string]string{}
	for k := range m {
		comp[k] = compute(k, directions, m)
	}

	cur := []string{}
	for k := range m {
		if k[2] == 'A' {
			cur = append(cur, k)
		}
	}
	fmt.Println(cur)

	repeat := 0
	done := false
	for !done {
		repeat++
		if repeat%10000000 == 0 {
			fmt.Println(repeat)
		}
		done := true
		newCur := []string{}
		for _, c := range cur {
			nw := comp[c]
			newCur = append(newCur, nw)
			done = done && nw[2] == 'Z'
		}
		if done {
			break
		}
		cur = newCur
	}

	return repeat * len(directions)
}
