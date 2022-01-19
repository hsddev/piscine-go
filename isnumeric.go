package piscine

func IsNumeric(s string) bool {
	var res bool
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}
