func maxSubArray(nums []int) int {
    maxSum, sum := nums[0], 0

    for _, val := range nums {
        sum += val
        if sum > maxSum {
            maxSum = sum
        }
        if sum < 0 {
            sum = 0
        }
    }

    return maxSum
}
