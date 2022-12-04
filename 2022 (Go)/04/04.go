package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	c1 := 0 // part 1
	c2 := 0 // part 2
	re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	for _, v := range cont {
		m := re.FindAllStringSubmatch(v, -1)[0][1:]
		k1, _ := strconv.Atoi(m[0])
		k2, _ := strconv.Atoi(m[1])
		k3, _ := strconv.Atoi(m[2])
		k4, _ := strconv.Atoi(m[3])
		if k1 <= k3 && k2 >= k4 || k3 <= k1 && k4 >= k2 {
			c1++
		}
		if k1 <= k3 && k2 >= k3 || k3 <= k1 && k4 >= k1 {
			c2++
		}
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", c1, c2)
}
