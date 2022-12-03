package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var hash = map[string]bool{}

func main() {
	f, err := os.ReadFile("inp.txt")
	if err != nil {
		log.Fatal(err)
	}
	cont := strings.Split(strings.Trim(string(f), " \n"), "\n")
	c := 0
	re := regexp.MustCompile(`(\w+)`)
	var mat [][]string // group matches. mat[0][1] is the actual captured word

	for _, v := range cont {
		mat = re.FindAllStringSubmatch(v, -1)
	}

	_ = cont
	_ = mat
	fmt.Println(c)
}
