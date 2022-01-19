package piscine

func IsPrintable(s string) bool {
	b := []byte(s)
	var res bool
	for i := 0; i < len(b); i++ {
		if b[i] >= 7 && b[i] <= 13 {
			res = false
			break
		} else {
			res = true
		}
	}
	return res
}
