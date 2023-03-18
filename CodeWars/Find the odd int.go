func FindOdd(seq []int) int {
	mp := make(map[int]int)

	for _, el := range seq {
		mp[el]++
	}

	for key, el := range mp {
		if el%2 != 0 {
			return key
		}
	}
	return 0
}
