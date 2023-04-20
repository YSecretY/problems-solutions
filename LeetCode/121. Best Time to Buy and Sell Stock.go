func maxProfit(prices []int) int {
    minVal, maxProf := prices[0], 0

    for _, val := range prices {
        if val < minVal {
            minVal = val
        } else if val - minVal > maxProf {
            maxProf = val - minVal
        }
    }

    return maxProf
}
