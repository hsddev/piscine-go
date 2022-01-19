package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func CountDigits(i int) int {
	count := 0
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

func intotstr(num int) string {
	count := CountDigits(num)
	p := num
	n := 0
	var srune []rune
	for i := 0; i < count; i++ {
		n = p % 10
		i = i / 10
		p = i
		srune = append(srune, rune(n)+48)
	}
	return string(srune)
}

func main() {
	arguments := os.Args
	importArgs := arguments[1:]
	validOp := []string{"+", "-", "/", "*", "%"}
	result := ""
	if len(importArgs) == 3 {
		for i := 0; i < len(validOp); i++ {
			if validOp[i] == importArgs[1] {
				if validOp[i] == "/" && importArgs[2] == "0" {
					result = "No division by 0"
				} else if validOp[i] == "%" && importArgs[2] == "0" {
					result = "No modulo by 0"
				} else if validOp[i] == "+" && importArgs[0] >= "0" && importArgs[0] <= "9" && importArgs[2] >= "0" && importArgs[2] <= "9" {
					a, err := strconv.Atoi(importArgs[0])
					if err != nil {
						result = "Error"
					}
					b, err := strconv.Atoi(importArgs[2])
					if err != nil {
						result = "Error"
					}
					result = intotstr(a + b)
				} else if validOp[i] == "*" {
					a, err := strconv.Atoi(importArgs[0])
					if err != nil {
						result = "Error"
					}
					b, err := strconv.Atoi(importArgs[2])
					if err != nil {
						result = "Error"
					}
					resInt := a * b
					fmt.Println(resInt)
				} else if validOp[i] == "-" {
					a, err := strconv.Atoi(importArgs[0])
					if err != nil {
						result = "Error"
					}
					b, err := strconv.Atoi(importArgs[2])
					if err != nil {
						result = "Error"
					}
					result = intotstr(a - b)
				} else if validOp[i] == "/" {
					a, err := strconv.Atoi(importArgs[0])
					if err != nil {
						result = "Error"
					}
					b, err := strconv.Atoi(importArgs[2])
					if err != nil {
						result = "Error"
					}
					result = intotstr(a / b)
				} else if validOp[i] == "%" {
					a, err := strconv.Atoi(importArgs[0])
					if err != nil {
						result = "Error"
					}
					b, err := strconv.Atoi(importArgs[2])
					if err != nil {
						result = "Error"
					}
					result = intotstr(a % b)
				}
			}
		}
	}
	if result != "" {
		for _, el := range result {
			z01.PrintRune(el)
		}
		z01.PrintRune('\n')
	}
}
