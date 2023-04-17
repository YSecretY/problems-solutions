func Parse(data string) []int {
	var res []int
	val := 0
	for _, el := range data {
		switch el {
		case 'i':
			val++
		case 'd':
			val--
		case 's':
			val *= val
		case 'o':
			res = append(res, val)
		}
	}

	return res
}
