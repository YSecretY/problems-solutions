func reverse(x int) int {
    val := int(math.Abs(float64(x)))
    res := 0
    for val > 0 && res <= math.MaxInt32 {
        res = res * 10 + val % 10
        val /= 10
    }
    if res > math.MaxInt32 {
        return 0
    }
    if x < 0 {
        res = -res
    }
    return res
}
