package piscine

func ToLower(s string) string {
	str := []rune(s)

	for i := 0; i < len(s); i++ {
		if s[i] <= 'Z' && s[i] >= 'A' {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}
