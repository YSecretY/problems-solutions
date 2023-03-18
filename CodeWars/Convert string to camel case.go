func ToCamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	res := ""

	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			res += string(s[i])
		} else {
			res += string(unicode.ToUpper(rune(s[i+1])))
			i++
		}
	}

	return res
}