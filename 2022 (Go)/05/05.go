package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(string(f), "\n")

	// parse
	lines := 0
	hash := map[int][]rune{}
	for j, v := range cont {
		if len(v) <= 1 {
			lines = j
			break
		}
		for i, ch := range v {
			if ch == '1' {
				break
			}
			ind := -1
			if i < 4 && i == 1 {
				ind = 0
			} else if i > 4 && i%4 == 1 {
				ind = i / 4
			}
			if ind >= 0 && !(ch == '[' || ch == ']' || ch == ' ') {
				hash[ind] = append(hash[ind], ch)
			}
		}
	}
	var t1 [][]rune // task 1
	var t2 [][]rune // task 2
	// create and reverse table lol
	for i := 0; i < len(hash); i++ {
		t1 = append(t1, []rune{})
		t2 = append(t2, []rune{})
		for j := range hash[i] {
			t1[i] = append(t1[i], hash[i][len(hash[i])-j-1])
			t2[i] = append(t2[i], hash[i][len(hash[i])-j-1])
		}
	}

	var k1, k2, k3 int
	for _, v := range cont[lines+1:] {
		if len(v) <= 1 { // end of line whitespace or so
			break
		}
		fmt.Sscanf(v, "move %d from %d to %d", &k1, &k2, &k3)

		for i := 0; i < k1; i++ {
			last := t1[k2-1][len(t1[k2-1])-1]
			t1[k2-1] = t1[k2-1][:len(t1[k2-1])-1]
			t1[k3-1] = append(t1[k3-1], last)
		}
		p := len(t2[k2-1])
		slice := t2[k2-1][p-k1:]
		t2[k2-1] = t2[k2-1][:p-k1]
		t2[k3-1] = append(t2[k3-1], slice...)

	}
	fmt.Print("Task 1: ")
	for _, v := range t1 {
		fmt.Printf("%c", v[len(v)-1])
	}
	fmt.Print("\nTask 2: ")
	for _, v := range t2 {
		fmt.Printf("%c", v[len(v)-1])
	}
	fmt.Println()

}
