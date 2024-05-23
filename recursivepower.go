package piscine

func RecursivePower(nb int, power int) int {
	var r int

	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	r = RecursivePower(nb, power-1) * nb

	return r
}
