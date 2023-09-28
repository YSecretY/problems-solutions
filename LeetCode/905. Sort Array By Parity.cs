public class Solution {
    public int[] SortArrayByParity(int[] nums)
    {
        for (int i = 0; i < nums.Count(); ++i)
        {
            for (int j = i; j > 0 && nums[j - 1] % 2 != 0; --j)
            {
                (nums[j - 1], nums[j]) = (nums[j], nums[j - 1]);
            }
        }

        return nums;
    }
}
