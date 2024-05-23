package piscine

func Join(strs []string, sep string) string {
	var new string
	for i := 0; i < len(strs); i++ {
		new = new + string(strs[i])
		if i != len(strs)-1 {
			new = new + string(sep)
		}
	}
	return new
}
