package ints

import "math"

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
