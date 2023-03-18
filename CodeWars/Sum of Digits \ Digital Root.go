func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	n = DigitalRoot(n/10) + n%10
	return DigitalRoot(n)
}
