package helper

import (
	"log"
	"strconv"
)

func Hint(x string) int {
	i, err := strconv.ParseInt(x, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}
