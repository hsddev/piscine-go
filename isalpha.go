package piscine

func IsAlpha(s string) bool {
	var res bool
	for i := 0; i < len(s); i++ {
		if ('a' <= s[i] && s[i] <= 'z') || ('0' <= s[i] && s[i] <= '9') || ('A' <= s[i] && s[i] <= 'Z') || len(s) == 0 {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}
