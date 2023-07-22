package main

import (
	"fmt"
	"hw2/util"
	"os"
)

func main() {
	input := os.Args[1:]
	typeOfInput := input[0]
	sortElements := input[1:]

	switch typeOfInput {
	case "-string":
		fmt.Println(util.SortString(sortElements))
	}
}