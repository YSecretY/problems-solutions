func containsDuplicate(nums []int) bool {
    stack := make(map[int]bool)

    for _, val := range nums {
        if !stack[val] {
            stack[val] = true
        } else {
            return true
        }
    }

    return false
}
