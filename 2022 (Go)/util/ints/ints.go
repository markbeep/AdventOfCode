package ints

import (
	"log"
	"math"
	"strconv"
)

func Sgn(x int) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	}
	return 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(a ...int) int {
	m := math.MinInt
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

func Min(a ...int) int {
	m := math.MaxInt
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

func BInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IBool(i int) bool {
	return i != 0
}

func SInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

func CInt(c byte) int {
	return int(c - '0')
}
