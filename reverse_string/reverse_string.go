package reverse

func Reverse(input string) string {
	var result = ""
	var runes = []rune(input)

	for i := 0; i < len(runes); i++ {
		idx := len(runes) - 1 - i
		result = result + string(runes[idx])
	}

	return result
}
