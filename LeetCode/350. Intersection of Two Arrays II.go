func intersect(nums1 []int, nums2 []int) []int {
    mp := make(map[int]int)
    var min *[]int
    var max *[]int
    if len(nums1) <= len(nums2) {
        min = &nums1
        max = &nums2
    } else {
        min = &nums2
        max = &nums1
    }

    for _, el := range *min {
        mp[el]++
    }

    var res []int
    for _, el := range *max {
        if val, ok := mp[el]; ok && val > 0 {
            res = append(res, el)
            mp[el]--
        }
    }

    return res
}
