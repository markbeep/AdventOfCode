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

func translateR(seedRanges []seedRange, loc []int) [][]int {
	newLocs := [][]int{}
	for i := 0; i < len(seedRanges); i++ {
		if len(loc) == 0 {
			break
		}
		r := seedRanges[i]
		start := loc[0]
		length := loc[1]
		offset := r.sourceRange - r.destRange
		rightMost := start + length - 1

		// inbetween [1, {2, 3}, 4]
		if start >= r.sourceRange && rightMost < r.sourceRange+r.length {
			newLocation := []int{start - offset, length}
			newLocs = append(newLocs, newLocation)
			loc = []int{}
			break // TODO: maybe not needed
		} else if start >= r.sourceRange && start < r.sourceRange+r.length {
			// right part overlaps [1, 2, {3, 4], 5}
			newLength := r.sourceRange + r.length - start
			newLocation := []int{start - offset, newLength}
			newLocs = append(newLocs, newLocation)

			newLocs = append(newLocs, translateR(seedRanges, []int{r.sourceRange + r.length, length - newLength})...)
			loc = []int{}
			break
		} else if start < r.sourceRange && rightMost >= r.sourceRange+length {
			// completely overlaps
			newLocation := []int{r.destRange, r.length}
			newLocs = append(newLocs, newLocation)

			newLocs = append(newLocs, translateR(seedRanges, []int{start, r.sourceRange - start})...)
			newLocs = append(newLocs, translateR(seedRanges, []int{r.sourceRange + r.length, rightMost - (r.sourceRange + r.length)})...)
			loc = []int{}
			break
		} else if start < r.sourceRange && rightMost >= r.sourceRange && rightMost < r.sourceRange+length {
			// left is outside, rightside is inside {0, [1, 2, 3}, 4]
			newLength := rightMost - r.sourceRange
			fmt.Println(start, rightMost, newLength)
			newLocation := []int{r.sourceRange, newLength}
			newLocs = append(newLocs, newLocation)

			newLocs = append(newLocs, translateR(seedRanges, []int{start, length - newLength})...)
			loc = []int{}
			break
		}

	}
	if len(loc) == 2 {
		// 0 overlap
		newLocs = append(newLocs, loc)
	}

	return newLocs
}

func translateRange(seedRanges []seedRange, locRange ...[]int) [][]int {
	newLocs := [][]int{}
	for _, loc := range locRange {
		newLocs = append(newLocs, translateR(seedRanges, loc)...)
	}
	return newLocs
}

func part2(input string) int {

	fmt.Println(translateR([]seedRange{{destRange: 10, sourceRange: 1, length: 10}}, []int{0, 20}))
	// [1, 2, 3, 4, 5, 6, 7, 8, 9, {10], 11, 12}, 13, 14, 15, 16, 17, 18, 19
	fmt.Println()

	return 0
	// f := strings.Split(input, "\n\n")
	// re := regexp.MustCompile(`\d+`)
	// res := re.FindAllString(f[0], -1)
	// seeds := [][]int{}
	// for i := 0; i < len(res); i += 2 {
	// 	seeds = append(seeds, []int{ints.SInt(res[i]), ints.SInt(res[i+1])})
	// }

	// for _, v := range f[1:] {
	// 	seedRanges := []seedRange{}
	// 	lines := strings.Split(v, "\n")
	// 	for _, line := range lines[1:] {
	// 		var d, s, r int
	// 		fmt.Sscanf(line, "%d %d %d", &d, &s, &r)
	// 		seedRanges = append(seedRanges, seedRange{destRange: d, sourceRange: s, length: r})
	// 	}
	// 	fmt.Println(seeds)
	// 	seeds = translateRange(seedRanges, seeds...)
	// }

	// mx := math.MaxInt
	// for _, v := range seeds {
	// 	mx = min(mx, v[0])
	// }
	// return mx
}
