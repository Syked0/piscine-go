package piscine

func UltimateDivMod(a *int, b *int) {
	y := *a / *b

	*b = *a % *b
	*a = y
}
