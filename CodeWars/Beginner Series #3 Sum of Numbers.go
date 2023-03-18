func GetSum(a, b int) int {
	if a == b {
		return a
	}

	var i, big int
	if a < b {
		i = a
		big = b
	} else {
		i = b
		big = a
	}

	s := 0
	for ; i <= big; i++ {
		s += i
	}

	return s
}
