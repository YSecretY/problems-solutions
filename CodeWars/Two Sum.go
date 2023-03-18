func TwoSum(numbers []int, target int) [2]int {
	mp := make(map[int]int)

	for i, el := range numbers {
		goal := target - el
		if _, ok := mp[goal]; ok {
			return [2]int{mp[goal], i}
		}
		mp[el] = i
	}

	return [2]int{}
}
