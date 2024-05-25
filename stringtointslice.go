package piscine

func StringToIntSlice(str string) []int {
	var rtn []int
	for _, runeValue := range str {
		rtn = append(rtn, int(runeValue))
	}
	return rtn
}
