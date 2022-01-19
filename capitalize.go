package piscine

func Capitalize(s string) string {
	capitalize := true
	slice := []rune(s)
	for index, element := range slice {
		if capitalize {
			if element >= 'a' && element <= 'z' {
				slice[index] = rune(element - 32)
				capitalize = false
			} else if (element >= 'A' && element <= 'Z') || (element >= '0' && element <= '9') {
				capitalize = false
			}
		} else if !capitalize {
			if element >= 'A' && element <= 'Z' {
				slice[index] = rune(element + 32)
			} else if (element < 'a' || element > 'z') && (element < '0' || element > '9') {
				capitalize = true
			}
		}
	}
	s = string(slice)
	return s
}
