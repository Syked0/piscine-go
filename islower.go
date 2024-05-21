package piscine

func IsLower(s string) bool {
	var b bool
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			b = true
		} else {
			b = false
			break
		}
	}
	return b
}
