package piscine

func IsLower(s string) bool {
	var res bool
	for i := 0; i < len(s); i++ {
		if 'a' <= s[i] && s[i] <= 'z' {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}
