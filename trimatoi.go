package piscine

func TrimAtoi(s string) int {
	var Numbers int
	var yy int = 1
	var booll bool
	for _, charr := range s {
		if charr >= '0' && charr <= '9' {

			rr := int(charr - '0')
			Numbers = Numbers*10 + rr
			booll = true
		} else if charr == '-' && !booll {
			yy = -1
		}
	}
	return Numbers * yy
}
