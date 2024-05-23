package piscine

func LastRune(s string) rune {
	str := []rune(s)
	i := len(s) - 1

	return str[i]
}
