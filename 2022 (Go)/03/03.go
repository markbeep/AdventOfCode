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
	var c int = 0
	content := strings.Split(strings.Trim(string(f), " \n"), "\n")

	for i := 0; i < len(content); i += 3 {
		mar := false
		for _, k1 := range content[i] {
			for _, k2 := range content[i+1] {
				for _, k3 := range content[i+2] {
					if !mar && k1 == k2 && k2 == k3 {
						mar = true
						p := int(k1)
						if p >= 97 {
							c += p - 96
						} else {
							c += p - 38
						}
					}
				}
			}
		}
	}
	fmt.Println(c)
}
