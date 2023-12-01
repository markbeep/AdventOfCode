package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Read in a file and trim spaces and newlines
func Read(fp string) string {
	f, err := os.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(f))
}

// Read in a file, trim and then split
func ReadS(fp, split string) []string {
	return strings.Split(Read(fp), split)
}

// Makes a 2D array
func Array2[T any](y, x int) [][]T {
	p := make([][]T, y)
	for i := range p {
		p[i] = make([]T, x)
	}
	return p
}

// Makes a 3D array
func Array3[T any](z, y, x int) [][][]T {
	p := make([][][]T, z)
	for i := 0; i < z; i++ {
		p[i] = make([][]T, y)
		for j := 0; j < y; j++ {
			p[i][j] = make([]T, x)
		}
	}
	return p
}

func CopyMap[T comparable, E any](m map[T]E) map[T]E {
	cpy := map[T]E{}
	for k, v := range m {
		cpy[k] = v
	}
	return cpy
}

func CopyToClipboard(answer int) {
	cmd := exec.Command("xclip", "-selection", "c")
	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	go func() {
		defer in.Close()
		_, err := in.Write([]byte(fmt.Sprintf("%d", answer)))
		if err != nil {
			panic(err)
		}
	}()

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
