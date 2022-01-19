package piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for index, el := range a {
		if f(a[index]) == true {
			result = f(el)
			break
		} else {
			result = false
		}
	}
	return result
}
