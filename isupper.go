package piscine

func IsUpper(s string) bool {
	var res bool
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}
