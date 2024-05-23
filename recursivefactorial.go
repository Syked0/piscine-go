package piscine

func RecursiveFactorial(nb int) int {
	r := 0
	if nb == 0 {
		return 1
	} else if nb < 0 || nb > 20 {
		return 0
	} else {
		r = RecursiveFactorial(nb-1) * nb
	}
	return r
}
