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
		if v[1] == '1' {
			lines = j + 1
			break
		}
		for i := 1; i < len(v); i += 4 {
			if v[i] != ' ' {
				hash[i/4] = append(hash[i/4], rune(v[i]))
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
		// task 1
		for i := 0; i < k1; i++ {
			last := t1[k2-1][len(t1[k2-1])-1]
			t1[k2-1] = t1[k2-1][:len(t1[k2-1])-1]
			t1[k3-1] = append(t1[k3-1], last)
		}
		// task 2
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
