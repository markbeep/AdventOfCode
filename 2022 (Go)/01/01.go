package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := strings.Split(string(f), "\n")
	m1 := 0
	m2 := 0
	m3 := 0
	t := 0
	for _, v := range content {
		if v != "" {
			a, _ := strconv.Atoi(v)
			t += a
		} else {
			if t > m1 {
				m3 = m2
				m2 = m1
				m1 = t
			} else if t > m2 {
				m3 = m2
				m2 = t
			} else if t > m3 {
				m3 = t
			}
			t = 0
		}
	}
	fmt.Println("Part 1:", m1)
	fmt.Println("Part 2:", m1+m2+m3)
	fmt.Println("Took:", time.Since(start))
}
