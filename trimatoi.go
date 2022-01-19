package piscine

func TrimAtoi(s string) int {
	new_slice := []rune{}
	firstelement := true
	for _, element := range s {
		if (element >= '0' && element <= '9') || (element == '-' && firstelement) {
			new_slice = append(new_slice, element)
			firstelement = false
		}
	}
	s = string(new_slice)
	n := atoi(s)
	return n
}

func atoi(s string) int {
	x := 0
	firstdigit := true
	negative := false
	for _, n := range s {
		if n == '-' && firstdigit {
			negative = true
		} else if n == '+' && firstdigit {
			negative = false
		} else if n < '0' && n > '9' {
			return 0
		} else {
			y := 0
			for i := '1'; i <= n; i++ {
				y++
			}
			x = x*10 + y
		}
		firstdigit = false
	}
	if negative {
		x = x * -1
	}
	return x
}
