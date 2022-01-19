package piscine

func FindNextPrime(nb int) int {
	if isPrimeEff(nb) {
		return nb
	}
	for !isPrimeEff(nb) {
		nb++
	}
	return nb
}

func isPrimeEff(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb <= 3 {
		return true
	} else if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	for i := 5; i*i < nb; i = i + 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}
