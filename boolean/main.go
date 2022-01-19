package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	var res bool
	if nbr%2 == 0 {
		res = true
	} else {
		res = false
	}
	return res
}

func main() {
	arguments := os.Args
	params := arguments[1:]
	numOfArgs := 0
	for i := 0; i <= len(params); i++ {
		numOfArgs = i
	}
	if isEven(numOfArgs) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
