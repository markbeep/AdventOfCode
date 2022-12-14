package main

import (
	"aoc/util"
	"aoc/util/ints"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	f := util.ReadS("inp.txt", "\n")
	c := 0
	re := regexp.MustCompile(`(\d+,\d+)`)
	field := util.Array2[bool](1000, 1000)
	minX, maxX, minY, maxY := 999, 0, 999, 0
	for _, v := range f {
		p := re.FindAllString(v, -1)
		for i, coord := range p {
			tmp := strings.Split(coord, ",")
			x, y := ints.SInt(tmp[0]), ints.SInt(tmp[1])
			minX = ints.Min(minX, x)
			minY = ints.Min(minY, y)
			maxX = ints.Max(maxX, x)
			maxY = ints.Max(maxY, y)
			if i == 0 {
				continue
			}
			tmp = strings.Split(p[i-1], ",")
			deltaX := x - ints.SInt(tmp[0])
			deltaY := y - ints.SInt(tmp[1])
			curX, curY := ints.SInt(tmp[0]), ints.SInt(tmp[1])
			for j := 0; j <= ints.Max(ints.Abs(deltaX), ints.Abs(deltaY)); j++ {
				field[curY][curX] = true
				curX += ints.Sgn(deltaX)
				curY += ints.Sgn(deltaY)
			}
		}
		fmt.Println(p)
	}
	for i := 0; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if field[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	curX, curY := 500, 0
	for {
		if field[0][500] {
			break
		}
		if !field[curY+1][curX] {
			curY++
		} else if !field[curY+1][curX-1] {
			curY++
			curX--
		} else if !field[curY+1][curX+1] {
			curY++
			curX++
		} else {
			c++
			field[curY][curX] = true
			curX, curY = 500, 0
		}
		if curY > maxY || curX > maxX || curX < minX {
			break
		}

	}

	_ = f
	fmt.Println(c)
}
