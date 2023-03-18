func CountBits(val uint) int {
	binVal := ""

	for val > 0 {
		binVal += strconv.FormatInt(int64(val%2), 10)
		val /= 2
	}

	count := 0
	for _, el := range binVal {
		if el == '1' {
			count++
		}
	}
	return count
}
