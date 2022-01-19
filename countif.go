package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int
	for index := range tab {
		if f(tab[index]) == true {
			count += 1
		}
	}
	return count
}
