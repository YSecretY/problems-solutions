func GetCount(str string) (count int) {
	for _, el := range str {
		switch el {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'O', 'U', 'I':
			count++
		}
	}
	return count
}
