package piscine

func Sqrt(nb int) int {
	var r int
	for i := 0; i <= nb; i++ {
		if nb == i*i {
			r = i
			break
		} else {
			r = 0
		}
	}
	return r
}
