package piscine

func ToUpper(s string) string {
	str := []rune(s)

	for i := 0; i < len(s); i++ {
		if s[i] <= 'z' && s[i] >= 'a' {
			str[i] = str[i] - 32
		}
	}
	return string(str)
}
