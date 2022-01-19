package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := intToDigits(n)
	sortIntArr(digits)
	for _, i := range digits {
		z01.PrintRune(rune(i) + '0')
	}
}

func intToDigits(n int) (digits []int) {
	for n > 0 {
		if n == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, n%10)
		}
		n /= 10
	}
	return
}

func sortIntArr(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
