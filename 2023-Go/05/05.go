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

func translateRange(seedRanges []seedRange, locRange ...[]int) [][]int {
	newLocs := [][]int{}
	for _, loc := range locRange {
		start := loc[0]
		length := loc[1]
		added := false
		for _, r := range seedRanges {
			offset := r.sourceRange - r.destRange
			if start >= r.sourceRange && start < r.sourceRange+r.length {
				rightCutoff := min(r.sourceRange+r.length, start+length)
				newLocation := []int{start - offset, min(length, rightCutoff-r.sourceRange)}
				newLocs = append(newLocs, newLocation)
				added = true
				break
			} else if start < r.sourceRange && start+length >= r.sourceRange {
				rightCutoff := min(r.sourceRange+r.length, start+length)
				newLocation := []int{r.destRange, min(length, rightCutoff-r.sourceRange)}
				newLocs = append(newLocs, newLocation)
				// add the lower half
				newLocs = append(newLocs, []int{start, r.sourceRange - start})
				added = true
				break
			}
		}
		if !added {
			newLocs = append(newLocs, loc)
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
		seeds = append(seeds, []int{ints.SInt(res[i]), ints.SInt(res[i+1])})
	}

	for _, v := range f[1:] {
		seedRanges := []seedRange{}
		lines := strings.Split(v, "\n")
		for _, line := range lines[1:] {
			var d, s, r int
			fmt.Sscanf(line, "%d %d %d", &d, &s, &r)
			seedRanges = append(seedRanges, seedRange{destRange: d, sourceRange: s, length: r})
		}
		seeds = translateRange(seedRanges, seeds...)
	}

	mx := math.MaxInt
	for _, v := range seeds {
		mx = min(mx, v[0])
	}
	return mx
}
