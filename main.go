package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	src := os.Args[1]
	wordsCount := len(strings.Fields(src))
	fmt.Println(wordsCount)
}
