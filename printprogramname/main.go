package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for index, char := range arguments {
		if index > 1 {
			z01.PrintRune(rune(char))
		}
	}
	z01.PrintRune('\n')
}
