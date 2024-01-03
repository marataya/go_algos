package sliceutils

func StringReverse(s string) (res string) {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		res += string(runes[i])
	}
	return
}
