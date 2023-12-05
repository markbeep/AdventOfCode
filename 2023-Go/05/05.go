package main

import (
	"aoc/util"
	"aoc/util/ints"
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

type seedRange struct {
	destRange   int
	sourceRange int
	length      int
}

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

func translate(seedRanges []seedRange, seeds ...int) []int {
	newLocs := []int{}
	for _, seed := range seeds {
		for _, r := range seedRanges {
			if seed >= r.sourceRange && seed < r.sourceRange+r.length {
				newLocs = append(newLocs, seed-(r.sourceRange-r.destRange))
			}
		}
	}
	return newLocs
}

func part1(input string) int {
	f := strings.Split(input, "\n\n")
	re := regexp.MustCompile(`\d+`)
	res := re.FindAllString(f[0], -1)
	seeds := []int{}
	for _, v := range res {
		seeds = append(seeds, ints.SInt(v))
	}

	for _, v := range f[1:] {
		seedRanges := []seedRange{}
		lines := strings.Split(v, "\n")
		for _, line := range lines[1:] {
			var d, s, r int
			fmt.Sscanf(line, "%d %d %d", &d, &s, &r)
			seedRanges = append(seedRanges, seedRange{destRange: d, sourceRange: s, length: r})
		}
		seeds = translate(seedRanges, seeds...)
	}

	mx := math.MaxInt
	for _, v := range seeds {
		mx = min(mx, v)
	}
	return mx
}

type rangeStruct struct {
	left, right, offset int
}

func trans(ranges []rangeStruct, seeds [][]int) [][]int {
	newLocs := [][]int{}
	for _, s := range seeds {
		left := s[0]
		right := s[1] // inclusive
		added := false
		for _, r := range ranges {
			if left >= r.left && right <= r.right {
				// inbetween
				newLocs = append(newLocs, []int{left - r.offset, right - r.offset})
				added = true
				break
			} else if left >= r.left && left <= r.right && right > r.right {
				// right overlap
				newLocs = append(newLocs, []int{left - r.offset, r.right - r.offset})
				newLocs = append(newLocs, []int{r.right + 1, right})
				added = true
				break
			} else if left < r.left && right >= r.left && right <= r.right {
				// left overlap
				newLocs = append(newLocs, []int{left, r.left - 1})

				newLocs = append(newLocs, []int{r.left - r.offset, right - r.offset})
				added = true
				break
			} else if left < r.left && right > r.right {
				// complete overlap
				newLocs = append(newLocs, []int{left, r.left - 1})
				newLocs = append(newLocs, []int{r.right + 1, right})
				newLocs = append(newLocs, []int{r.left, r.right})
				added = true
				break
			}
		}
		if !added {
			newLocs = append(newLocs, []int{left, right})
		}
	}
	return newLocs
}

func part2(input string) int {
	f := strings.Split(input, "\n\n")
	re := regexp.MustCompile(`\d+`)
	res := re.FindAllString(f[0], -1)
	seeds := [][]int{}
	for i := 0; i < len(res); i += 2 {
		left := ints.SInt(res[i])
		right := ints.SInt(res[i+1]) + left - 1
		seeds = append(seeds, []int{left, right})
	}

	for _, v := range f[1:] {
		seedRanges := []rangeStruct{}
		lines := strings.Split(v, "\n")
		for _, line := range lines[1:] {
			var d, s, r int
			fmt.Sscanf(line, "%d %d %d", &d, &s, &r)
			seedRanges = append(seedRanges, rangeStruct{left: s, right: s + r - 1, offset: s - d})
		}
		seeds = trans(seedRanges, seeds)
	}

	mx := math.MaxInt
	for _, v := range seeds {
		mx = min(mx, v[0])
	}
	return mx
}
