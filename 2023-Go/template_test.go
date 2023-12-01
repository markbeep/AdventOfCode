package main

import (
	"testing"
)

const (
	expected1 = 0
	expected2 = 0
)

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}

func TestPart1(t *testing.T) {
	ans := part1(input)
	if ans != expected1 {
		t.Errorf("expected=%d, got=%d", expected1, ans)
	}
}

func TestPart2(t *testing.T) {
	ans := part2(input)
	if ans != expected1 {
		t.Errorf("expected=%d, got=%d", expected1, ans)
	}
}
