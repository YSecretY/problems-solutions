#include <iostream>
#include <vector>

void sort(std::vector<int>& nums)
{
    for (int i = 1; i < nums.size(); ++i) {
        for (int j = i; j > 0 && nums[j - 1] > nums[j]; --j) {
            std::swap(nums[j], nums[j - 1]);
        }
    }
}

void merge(std::vector<int>& nums1, int m, std::vector<int>& nums2, int n)
{
    if (nums1.empty())
        return;
    if (nums1.size() == 1 && !nums2.empty()) {
        nums1[0] = nums2[0];
        return;
    }
    int k = 0;
    for (int i = m; i < m + n; ++i) {
        nums1[i] = nums2[k++];
    }
    sort(nums1);
}
