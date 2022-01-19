package piscine

func Rot14(s string) string {
	sRune := []rune(s)
	var newStr string
	for i := 0; i < len(sRune); i++ {
		if sRune[i] == 'z' {
			sRune[i] = 'n'
		} else if sRune[i] == 'Z' {
			sRune[i] = 'N'
		} else if sRune[i] >= 'a' && sRune[i] <= 'z' {
			if sRune[i] >= 'm' {
				sRune[i] -= 12
			} else {
				sRune[i] += 14
			}
		} else if sRune[i] >= 'A' && sRune[i] <= 'Z' {
			if sRune[i] >= 'M' {
				sRune[i] -= 12
			} else {
				sRune[i] += 14
			}
		}
		newStr += string(sRune[i])
	}
	return newStr
}
