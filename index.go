package piscine

func Index(s string, toFind string) int {
	lenString := len(s)
	sublen := len(toFind)
	count := 0
	if sublen > lenString {
		return -1
	}
	if toFind == "" || s == "" {
		return 0
	}
	sRune := []rune(s)
	toFindRune := []rune(toFind)
	for i := 0; i < lenString-sublen+1; i++ {
		for j := 0; j < sublen; j++ {
			if sRune[j+i] == toFindRune[j] {
				count++
				if count == sublen {
					return i
				}
			} else {
				count = 0
				break
			}
		}
	}
	return -1
}
