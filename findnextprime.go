package piscine

func FindNextPrime(nb int) int {
	var yy int
	check := false
	if nb <= 0 {
		yy = 2
	} else {
		if nb > 0 || nb > 8888888 {
			for i := nb; ; i++ {
				if IsPrime(i) {

					yy = i
					check = true

				}

				if check {
					return yy
				}
			}
		}
	}

	return yy
}
