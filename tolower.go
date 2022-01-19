package piscine

func ToLower(s string) string {
	b := []byte(s)
	var res string
	for i := 0; i < len(b); i++ {
		if b[i] >= 65 && b[i] <= 90 {
			b[i] = b[i] + 32
			res += string(b[i])
		} else {
			res += string(b[i])
		}
	}
	return res
}
