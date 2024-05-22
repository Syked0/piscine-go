package piscine

func IsPrime(nb int) bool {
	b := false

	if nb == 1 {
		b = false
	} else {
		for i := 1; i <= nb; i++ {

			resulta := nb % i

			if resulta == 0 && nb != i && i != 1 {
				b = false
				break
			} else {
				b = true
			}
		}
	}

	return b
}
