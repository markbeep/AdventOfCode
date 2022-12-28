package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	c1 := 0 // part 1
	c2 := 0 // part 2
	for _, v := range cont {
		var k1, k2, k3, k4 int
		fmt.Sscanf(v, "%d-%d,%d-%d", &k1, &k2, &k3, &k4)
		if k1 <= k3 && k2 >= k4 || k3 <= k1 && k4 >= k2 {
			c1++
		}
		if k1 <= k3 && k2 >= k3 || k3 <= k1 && k4 >= k1 {
			c2++
		}
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", c1, c2)
	fmt.Println("Took:", time.Since(start))
}
