package main

import (
	z01 "github.com/01-edu/z01"
)

func main() {
	var str string = "abcdefghijklmnopqrstuvwxyz"

	for i := len(str) - 1; i >= 0; i-- {
		z01.PrintRune(rune(str[i]))
	}

	z01.PrintRune('\n')
}
