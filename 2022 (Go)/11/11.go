package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

type monke struct {
	items []int
	op    []string
	div   int
	ift   int
	iff   int
	c     int
}

func main() {
	f := util.ReadS("inp.txt", "\n")
	re_start := regexp.MustCompile(`\d+`)
	re_op := regexp.MustCompile(`([\w\d\*\+]+)`)
	monkes := map[int]*monke{}
	for i := 0; i < len(f); i += 7 {
		var div, monkey, ift, iff int
		fmt.Sscanf(f[i], "Monkey %d:", &monkey)
		start := re_start.FindAllString(f[i+1], -1)
		op := re_op.FindAllString(f[i+2], -1)
		fmt.Sscanf(f[i+3], "  Test: divisible by %d", &div)
		fmt.Sscanf(f[i+4], "    If true: throw to monkey %d", &ift)
		fmt.Sscanf(f[i+5], "    If false: throw to monkey %d", &iff)
		monkes[monkey] = &monke{items: []int{}, op: op, div: div, ift: ift, iff: iff}
		for _, v := range start {
			val, _ := strconv.Atoi(v)
			monkes[monkey].items = append(monkes[monkey].items, val)
		}
	}

	// round
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkes); j++ {
			v := monkes[j]
			for len(v.items) > 0 {
				it := v.items[0]
				v.items = v.items[1:]
				it = oop(it, v.op) / 3
				if it%v.div == 0 {
					monkes[v.ift].items = append(monkes[v.ift].items, it)
				} else {
					monkes[v.iff].items = append(monkes[v.iff].items, it)
				}
				v.c++
			}
		}
	}
	m1 := 0
	m2 := 0
	for _, v := range monkes {
		if v.c > m1 {
			m2 = m1
			m1 = v.c
		} else if v.c > m2 {
			m2 = v.c
		}
	}
	fmt.Println(m1 * m2)

}

func oop(x int, ops []string) int {
	val := 0
	v2 := 0
	if ops[2] == "old" {
		val = x
	} else {
		v, _ := strconv.Atoi(ops[2])
		val = v
	}
	if ops[4] == "old" {
		v2 = x
	} else {
		v, _ := strconv.Atoi(ops[4])
		v2 = v
	}
	if ops[3] == "*" {
		return val * v2
	}
	return val + v2
}
