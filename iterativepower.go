package piscine

func IterativePower(nb int, power int) int {
	if nb == 0 && power == 0 {
		return 1
	}
	if nb == 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	r := 1
	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		r *= nb
	}
	return r
}
