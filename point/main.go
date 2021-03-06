package main

import "github.com/01-edu/z01"

type point struct {
	x string
	y string
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func printStr(s string) {
	sRune := []rune(s)
	for i := 0; i < len(sRune); i++ {
		z01.PrintRune(sRune[i])
	}
}

func main() {
	points := point{}
	setPoint(&points)
	newStr := "x = " + points.x + ", y = " + points.y
	printStr(newStr)
	z01.PrintRune('\n')
}
