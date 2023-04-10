func majorityElement(nums []int ) int {
    count := 0
    res := 0

    for _, el := range nums {
        if count == 0 {
            res = el
        }
        if el == res {
            count++
        } else {
            count--
        }
    }

    return res
}
