package piscine

func IsNumeric(s string) bool {
	var b bool
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			b = true
		} else {
			b = false
			break
		}
	}
	return b
}
