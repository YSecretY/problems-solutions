func duplicate_count(s string) int {
	mp := make(map[rune]int)

	for _, r := range s {
		mp[unicode.ToLower(r)]++
	}

	count := 0
	for _, el := range mp {
		if el > 1 {
			count++
		}
	}
	return count
}
