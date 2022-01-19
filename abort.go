package piscine

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	var minArr []int
	for i := 0; i < len(arg); i++ {
		for j := 1; j < len(arg)-i; j++ {
			if arg[j] < arg[j-1] {
				arg[j], arg[j-1] = arg[j-1], arg[j]
			}
		}
		minArr = append(minArr, arg[i])
	}
	return minArr[2]
}
