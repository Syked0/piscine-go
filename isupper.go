package piscine

func IsUpper(s string) bool {
	var b bool
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			b = true
		} else {
			b = false
			break
		}
	}
	return b
}
