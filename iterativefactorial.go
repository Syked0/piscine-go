package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	r := 1
	for i := 1; i < nb+1; i++ {
		r = r * i
		if nb > 20 {
			return 0
		}

	}
	return r
}
