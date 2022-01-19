package piscine

func NRune(s string, n int) rune {
	myRune := []rune(s)
	if n > len(myRune) || n < 1 {
		return 0
	}
	return myRune[n-1]
}
