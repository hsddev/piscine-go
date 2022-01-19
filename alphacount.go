package piscine

func AlphaCount(str string) int {
	n := 0
	s := []rune(str)
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') {
			s[n] = b
			n++
		}
	}
	return len(s[:n])
}
